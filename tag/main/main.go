package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	UserId   int    `json:"user_id" bson:"user_ids"`
	UserName string `json:"user_name" bson:"user_name"`
}

func main() {
	u := &User{
		UserId:   123456,
		UserName: "Txiaozhe",
	}

	j, _ := json.Marshal(u)
	fmt.Println(string(j))

	t := reflect.TypeOf(u)
	field1 := t.Elem().Field(0)
	fmt.Println(field1.Tag.Get("json"))
	fmt.Println(field1.Tag.Get("bson"))

	field2 := t.Elem().Field(1)
	fmt.Println("json tag: ", field2.Tag.Get("json"))
	fmt.Println("bson tag: ", field2.Tag.Get("bson"))
}
