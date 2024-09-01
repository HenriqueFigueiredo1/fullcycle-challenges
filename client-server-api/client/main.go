package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	File           string = "cotacao.txt"
	URLServerQuote string = "http://localhost:8080/cotacao"
)

type ResponseQuote struct {
	Bid string `json:"bid"`
}

func main() {
	quote, err := GetQuote()
	if err != nil {
		log.Printf("error when requesting quote: %s", err.Error())
		return
	}

	f, err := os.Create(File)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := fmt.Sprintf("Dólar: %s", *quote)

	_, err = f.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
}

func GetQuote() (*string, error) {
	cause := errors.New("timeout error when placing the quote request")
	ctx, cancel := context.WithTimeoutCause(context.Background(), time.Millisecond*300, cause)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", URLServerQuote, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if context.Cause(ctx) != nil {
			return nil, cause
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	quote := ResponseQuote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &quote.Bid, nil
}
