package main

// naimportujeme si balíčky, se kterými v apičku asi budeme pracovat.. 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"
)

// posílání a příjem dat.. 
type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// fce pro GET požadavek
func makeGetRequest(url string) {
	client := &http.Client{
		Timeout: time.Second * 10, // nastavení timeoutu
	}

	// Vytvoření GET požadavku
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Nastavení hlaviček
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN")
	req.Header.Set("Cache-Control", "no-cache")


	// Odeslání požadavku
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Zpracování odpovědi
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response:", string(body))
}

// Funkce pro odeslání POST požadavku
func makePostRequest(url string, data Data) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Převod dat do JSON formátu
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// Vytvoření POST požadavku
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Nastavení hlaviček
	req.Header.Set("Content-Type", "application/json")

	// Odeslání požadavku
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Zpracování odpovědi
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	} 
	fmt.Println("Response:", string(body))
}

func main() {
	// URL pro GET a POST požadavek
	url := "https://example.com/api"

	// Příklad GET požadavku
	makeGetRequest(url)

	// Příklad POST požadavku
	data := Data{Name: "Test", Value: "Hello"}
	makePostRequest(url, data)
}
