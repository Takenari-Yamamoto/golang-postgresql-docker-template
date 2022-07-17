package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email"`
	Created time.Time `json:"created"`
	A *A `json:"A,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	} {
		Name: "Mr." + u.Name,
	})
	return v, err
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// json to struct
	u2 := new(User)

	if err := json.Unmarshal(bs, &u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)


}
