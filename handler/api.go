package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Val string `json:"val"`
}

func API(c *gin.Context) {

	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("errr:", err)
	}
	fmt.Println("jsonbytes", string(jsonBytes))
	var data Data
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		fmt.Println("unmarshal err:::", err)
	}

	resp := ReverseString(data.Val)
	fmt.Println("resp", resp)
	c.JSON(200, resp)
}

type Person struct {
	DOB CustomDate `json:"dob"`
}
type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]                   // Remove the quotes from the string
	t, err := time.Parse("2006-01-02", s) // Specify the layout for your date string
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func FindAge(c *gin.Context) {
	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
	}
	fmt.Println("Received JSON data:", string(jsonBytes))

	var person Person
	if err := json.Unmarshal(jsonBytes, &person); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	fmt.Println("Unmarshaled :", person.DOB)
	age := CalculateAge(person.DOB)
	c.JSON(200, age)
}

func CalculateAge(dob CustomDate) int {
	now := time.Now()
	fmt.Println("dob", dob.Year())
	fmt.Println("now", now.Year())
	var age int
	if dob.Year() < now.Year() {
		age = now.Year() - dob.Year() - 1
	}
	return age
}

func ReverseString(str string) string {
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}
