package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type PriceData struct {
	Price string `json:"price"`
}

type PriceResponse struct {
	HighestPrice float64 `json:"highest_price"`
	LowestPrice  float64 `json:"lowest_price"`
}

var (
	highestPrice float64
	lowestPrice  float64
	mutex        sync.Mutex
	botToken     = "7105924273:AAHqk07jfhQrHyAbk1ppe_A3BrgPJOVaGas"
)

func main() {
	highestPrice = 0
	lowestPrice = 9999999999999

	// En yüksek ve en düşük fiyatları düzenli olarak güncellemek için bir gorutine başlatın
	symbol := "BTCUSDT"
	go updatePricesPeriodically(symbol)

	// API endpoint'ini tanımla
	http.HandleFunc("/price", getPrice)

	// Bot komutlarını dinleyen endpoint'ı tanımla
	http.HandleFunc("/"+botToken, handleTelegramUpdates)

	// HTTP sunucusunu başlat
	fmt.Println("Server started at :8080")
	go http.ListenAndServe(":8080", nil)

	// Sonsuz döngüyü başlat
	select {}
}

func getPrice(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	// En yüksek ve en düşük fiyatları içeren yanıtı oluştur
	priceData := PriceResponse{
		HighestPrice: highestPrice,
		LowestPrice:  lowestPrice,
	}

	// JSON formatına dönüştür
	response, err := json.Marshal(priceData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Yanıtı gönder
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func updatePrice(price float64) {
	mutex.Lock()
	defer mutex.Unlock()

	// En yüksek ve en düşük fiyatları güncelle
	if price > highestPrice {
		highestPrice = price
	}
	if price < lowestPrice {
		lowestPrice = price
	}
}

func getCurrentPrice(symbol string) (float64, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var priceData PriceData
	err = json.NewDecoder(resp.Body).Decode(&priceData)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(priceData.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func updatePricesPeriodically(symbol string) {
	for {
		price, err := getCurrentPrice(symbol)
		if err != nil {
			fmt.Println("Fiyat alınamadı:", err)
			continue
		}

		updatePrice(price)

		time.Sleep(5 * time.Second)
	}
}

func handleTelegramUpdates(w http.ResponseWriter, r *http.Request) {
	var update struct {
		Message struct {
			Text string `json:"text"`
		} `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	message := update.Message.Text
	if message == "/start" {
		mutex.Lock()
		defer mutex.Unlock()

		// En yüksek ve en düşük fiyatları içeren yanıtı oluştur
		priceData := PriceResponse{
			HighestPrice: highestPrice,
			LowestPrice:  lowestPrice,
		}

		// JSON formatına dönüştür
		response, err := json.Marshal(priceData)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Yanıtı gönder
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
