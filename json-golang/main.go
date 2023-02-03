package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// to convert json to struct use google automatically then copy paste below:`)

// type User struct {
// 	Name string `json:name`
// 	Age  int    `json:age`
// }

type Something struct {
	Name    string `json:"name"`
	Married bool   `json:"married"`
	Age     int    `json:"age"`
	Address []struct {
		Street int    `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
	Marks []int `json:"marks"`
}

func main() {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var something Something
	json.Unmarshal(jsonByteValues, &something) // converting json data to struct
	log.Println(something)
	// fmt.Println(string(jsonByteValues)) //converting json data to string

	// Marshall (code) :=
	// newJsonByteValues,err:=json.Marshal(something)
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// os.WriteFile("getting.json"newnewJsonByteValues)
	// (consume rest api in go code)
	// or
	// (using rest api)
	// returned value ? = json
	// how to read json?
	// problem that occurs during reading json data as string
	// sol is := read json in struct format (Marshalling),(Unmarshalling)

}
