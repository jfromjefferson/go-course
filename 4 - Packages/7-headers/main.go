package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
)

var allowedPaths = []string{"/", "/about", "/contact"}

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
	fmt.Println("Starting server...")

	http.HandleFunc("/", PostalCodeHandler)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe error: %s\n", err)
		return
	}

}

func PostalCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(`{"error": "Method not allowed"}`))
		if err != nil {
			return
		}
		return
	}

	if slices.Contains(allowedPaths, r.URL.Path) == false {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(`{"error": "Not found"}`))
		if err != nil {
			return
		}
		return
	}

	code := r.URL.Query().Get("code")

	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(`{"error": "Missing postalCode"}`))
		if err != nil {
			return
		}

		return
	}

	postalCode, postalCodeError := GetPostalCode(code)

	if postalCodeError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(`{"error": "Internal server error"}`))
		if err != nil {
			return
		}

		return
	}

	err := json.NewEncoder(w).Encode(postalCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetPostalCode(code string) (*PostalCode, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + code + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get error: %s\n", err)
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Read error: %s\n", err)
		return nil, err
	}

	var postalCode PostalCode
	err = json.Unmarshal(body, &postalCode)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unmarshal error: %s\n", err)
		return nil, err
	}

	return &postalCode, nil
}
