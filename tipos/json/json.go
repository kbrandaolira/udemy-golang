package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct {
	Id     int      `json:"id"`
	Nome   string   `json:"nome"`
	Emails []string `json:"emails"`
}

func main() {
	p1 := pessoa{1, "Kayan", []string{"klira@gmail.com"}}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	var p2 pessoa

	jsonString := `{"id":1,"nome":"Kayan","emails":["klira@gmail.com"]}`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)

}
