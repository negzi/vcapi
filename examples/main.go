package main

import (
	"fmt"
	"log"
	"os"

	"github.com/whitby/vcapi"
)

var (
	username = "api.whitby"
	password string
)

func init() {
	password = os.Getenv("VCAPI_PASSWORD")
	if password == "" {
		log.Fatal("VCAPI_PASSWORD not set")
	}

}

func main() {
	config := &vcapi.Config{
		Username:   username,
		Password:   password,
		SchoolID:   "whitby",
		APIVersion: "v2",
	}
	client := vcapi.NewClient(config)
	opt := &vcapi.ListOptions{Params: vcapi.Params{}}

	for {
		students, err := client.Students.List(opt)
		if err != nil {
			log.Fatal(err)
		}

		for _, student := range students {
			fmt.Println(student.FirstName + " " + student.LastName)
			fmt.Println(client.Students.Relationships(student))
		}
		opt.Page++

		if opt.NextPage == 0 {
			break
		}
	}

}
