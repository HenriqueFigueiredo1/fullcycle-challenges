package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DBFile   string = "quote.db"
	URLQuote string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

type Quote struct {
	ID         string `json:"id" gorm:primaryKey`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type QuoteResponse struct {
	Quote Quote `json:"USDBRL"`
}

func NewQuoteResponse() *QuoteResponse {
	return &QuoteResponse{
		Quote: Quote{
			ID: uuid.New().String(),
		},
	}
}

func main() {
	db := ConnectDB()

	http.HandleFunc("/cotacao", HandleDollarQuote(db))
	http.ListenAndServe(":8080", nil)
}

func HandleDollarQuote(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		quote, err := GetQuote()
		if err != nil {
			log.Printf("error when requesting quote: %s", err.Error())
			http.Error(w, ResponseError(err), http.StatusInternalServerError)
			return
		}

		cause := errors.New("timeout error when performing an action on the database")
		ctx, cancel := context.WithTimeoutCause(context.Background(), time.Millisecond*10, cause)
		defer cancel()

		db.WithContext(ctx).Create(quote)
		causeErr := context.Cause(ctx)
		if ctx.Err() != nil && causeErr != nil {
			log.Printf("error when saving quote: %s", causeErr.Error())
			http.Error(w, ResponseError(causeErr), http.StatusInternalServerError)
			return
		}

		log.Print("quote generated successfully")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(quote)
	}
}

func GetQuote() (*Quote, error) {
	cause := errors.New("timeout error when placing the quote request")
	ctx, cancel := context.WithTimeoutCause(context.Background(), time.Millisecond*200, cause)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", URLQuote, nil)
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

	quoteResponse := NewQuoteResponse()
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		return nil, err
	}

	return &quoteResponse.Quote, nil
}
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBFile), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Quote{})

	return db
}

func ResponseError(err error) string {
	msg := map[string]any{
		"error": err.Error(),
		"time":  time.Now(),
	}

	msgStr, _ := json.Marshal(msg)

	return string(msgStr)
}
