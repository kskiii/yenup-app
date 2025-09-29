package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Success bool `json:"success"`
	Query   struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Rate float64 `json:"rate"`
	} `json:"info"`
	Result float64 `json:"result"`
}

func main() {
	fmt.Println("💹 YenUp - Check if JPY is stronger today!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//  通貨指定（CAD -> JPY）
	base := os.Getenv("BASE_CURRENCY")
	target := os.Getenv("TARGET_CURRENCY")
	amount := os.Getenv("AMOUNT")
	apiKey := os.Getenv("API_KEY")
	apiUrl := fmt.Sprintf("https://api.exchangerate.host/convert?from=%s&to=%s&access_key=%s&amount=%s", base, target, apiKey, amount)

	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var data Response
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		panic(err)
	}

	if data.Success {
		fmt.Printf("1 %s = %.2f %s\n", data.Query.From, data.Result, data.Query.To)
	} else {
		fmt.Println("API request failed")
	}

}
