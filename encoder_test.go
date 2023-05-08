package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Arif",
		LastName:   "Bukittinggi",
	}

	encoder.Encode(customer)

	fmt.Println(customer)
}
