package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struktura pro parsování odpovědi z API
type ApiResponse struct {
	IP        string  `json:"ip"`
	Country   string  `json:"country_name"`
	Region    string  `json:"region_name"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	// API klíč (vlož svůj vlastní klíč zde)
	apiKey := "xxxxxxxxxxxxxxxxxxxxxx" // Změň na svůj klíč z apiip.net
	apiURL := fmt.Sprintf("https://apiip.net/api?key=%s", apiKey)

	// Odeslat HTTP GET požadavek
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Chyba při získávání veřejné IP adresy:", err)
		return
	}
	defer resp.Body.Close()

	// Kontrola stavu odpovědi
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Chyba: server vrátil status %d\n", resp.StatusCode)
		return
	}

	// Parsování JSON odpovědi
	var data ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Chyba při parsování JSON odpovědi:", err)
		return
	}

	// Výstup získaných dat
	fmt.Println("Veřejná IP adresa počítače je:", data.IP)
	fmt.Println("Země:", data.Country)
	fmt.Println("Region:", data.Region)
	fmt.Println("Město:", data.City)
	fmt.Printf("Poloha: %.2f, %.2f\n", data.Latitude, data.Longitude)
}
