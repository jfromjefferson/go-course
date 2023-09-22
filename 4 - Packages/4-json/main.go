package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BankAccount struct {
	Owner     string `json:"o"`
	Balance   int    `json:"b"`
	HasCredit bool   `json:"hc"`
	CreatedAt string `json:"ca"`
}

type Ip struct {
	Ip      string `json:"ip"`
	Country string `json:"country"`
	GeoIp   string `json:"geo-ip"`
	APIHelp string `json:"API help"`
}

func main() {
	bankAccount := BankAccount{
		Owner:     "John Doe",
		Balance:   1000,
		HasCredit: true,
		CreatedAt: "2021-01-01",
	}

	result, error := json.Marshal(bankAccount)

	if error != nil {
		panic(error)
	}

	println(string(result))

	// encoder := json.NewEncoder(os.Stdout)

	// encoder.Encode(bankAccount)

	// We can also use json.Unmarshal to convert a JSON string to a Go struct
	// Using tags, we can map the JSON keys to the struct fields
	accountDataJson := []byte(`{
		"o": "Jane Doe",
		"b": 2000,
		"hc": false,
		"ac": "2021-01-01"
	}`)

	var newBankAccount BankAccount

	err := json.Unmarshal(accountDataJson, &newBankAccount)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", newBankAccount)

	url := "https://jsonip.com"

	response, responseError := http.Get(url)

	if responseError != nil {
		panic(responseError)
	}

	defer response.Body.Close()

	responseBody, bodyError := io.ReadAll(response.Body)

	if bodyError != nil {
		panic(bodyError)
	}

	var ip Ip

	json.Unmarshal(responseBody, &ip)

	fmt.Printf("%+v\n", ip)

	fmt.Println("IP:", ip.Ip)
	fmt.Println("Country:", ip.Country)
	fmt.Println("Geo IP:", ip.GeoIp)
	fmt.Println("API Help:", ip.APIHelp)

}
