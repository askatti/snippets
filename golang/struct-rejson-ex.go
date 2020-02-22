package main

import (
	c "controller/cmd/common"
	"encoding/json"
	"fmt"
	"log"
)

// Name - student name
type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Student - student object
type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

func Example_JSONSet() {

	student := Student{
		Name: Name{
			"Mark",
			"S",
			"Pronto",
		},
		Rank: 1,
	}
	ok := c.Evdb.SetStruct("student", student)
	if !ok {
		log.Printf("Failed to set struct")
		return
	}

	readStudent := Student{}

	studentJSON,err := c.Evdb.GetStruct("student")
	err = json.Unmarshal(studentJSON, &readStudent)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return
	}

	fmt.Printf("Student read from redis : %s\n", readStudent.Name)
	fmt.Printf("Student read from redis : %d\n", readStudent.Rank)
}

func main() {
	serverport := "localhost:6379"
	c.Evdb.Dbinit(serverport, c.EVENTSTOREDB)

	fmt.Println("Executing Example_JSONSET for Redigo Client")
	Example_JSONSet()

}
