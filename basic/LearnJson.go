package basic

import (
	"encoding/json"
	"fmt"
)

// Username : belajar encode & decode json
type Username struct {
	FullName string `json:"Name"`
	Age      int
}

// Cetak : method untuk mencetak
func (username Username) Cetak() ([]Username, error) {
	var err error
	var jsonString = `[
		{"Name" : "Al Kautsar", "Age" : 1},
		{"Name" : "Aisyah", "Age" : 1}
	]`

	var jsonData = []byte(jsonString)

	var data []Username

	// decode to object array
	json.Unmarshal(jsonData, &data)

	for i := 0; i < len(data); i++ {
		fmt.Println("Username :", data[i].FullName)
		fmt.Println("Age :", data[i].Age)
	}

	// object to json string
	jsonObject, err := json.Marshal(data)

	fmt.Println(string(jsonObject))
	if err != nil {
		return nil, err
	}

	return data, err
}
