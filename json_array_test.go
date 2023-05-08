package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Arif",
		LastName:   "Bukittinggi",
		Age:        34,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","MiddleName":"Arif","LastName":"Bukittinggi","Age":34,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)

}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Arif",
		Addresses: []Address{
			{
				Street:     "Jalanin Aja Dulu",
				Country:    "Wakanda",
				PostalCode: "26121",
			},
			{
				Street:     "Jalan Jalan Dulu",
				Country:    "Indocina",
				PostalCode: "8888",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Arif","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalanin Aja Dulu","Country":"Wakanda","PostalCode":"26121"},{"Street":"Jalan Jalan Dulu","Country":"Indocina","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalanin Aja Dulu","Country":"Wakanda","PostalCode":"26121"},{"Street":"Jalan Jalan Dulu","Country":"Indocina","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	Addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, Addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(Addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	Addresses := []Address{
		{
			Street:     "Jalanin Aja Dulu",
			Country:    "Wakanda",
			PostalCode: "26121",
		},
		{
			Street:     "Jalan Jalan Dulu",
			Country:    "Indocina",
			PostalCode: "8888",
		},
	}
	bytes, _ := json.Marshal(Addresses)
	fmt.Println(string(bytes))
}
