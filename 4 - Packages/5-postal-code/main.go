package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type PostalCode struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		request, requestError := http.Get("https://viacep.com.br/ws/" + url + "/json/")

		if requestError != nil {
			fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", url, requestError)
			return
		}

		defer request.Body.Close()

		response, responseError := io.ReadAll(request.Body)

		if responseError != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", url, responseError)
			return
		}

		var postalCode PostalCode
		postalCodeError := json.Unmarshal(response, &postalCode)

		if postalCodeError != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", url, postalCodeError)
			return
		}

		fmt.Printf("CEP: %s\n", postalCode)
	}
}
