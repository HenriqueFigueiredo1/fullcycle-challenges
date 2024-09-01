package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func fetchFromAPI(ctx context.Context, url string, ch chan<- *http.Response, chErr chan<- error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		chErr <- err
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		chErr <- err
		return
	}

	ch <- resp
}

func getAddressFromResponse(resp *http.Response) (Address, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Address{}, err
	}

	var address Address
	if err := json.Unmarshal(body, &address); err != nil {
		return Address{}, err
	}
	return address, nil
}

func main() {
	cep := "01153000"
	url1 := "https://brasilapi.com.br/api/cep/v1/" + cep
	url2 := "http://viacep.com.br/ws/" + cep + "/json/"

	ch := make(chan *http.Response)
	chErr := make(chan error)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go fetchFromAPI(ctx, url1, ch, chErr)
	go fetchFromAPI(ctx, url2, ch, chErr)

	select {
	case resp := <-ch:
		address, err := getAddressFromResponse(resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("API: %s\nAddress: %+v\n", resp.Request.URL, address)
	case err := <-chErr:
		log.Fatal("Error fetching from API: ", err)
	case <-ctx.Done():
		fmt.Println("Request timed out")
	}
}
