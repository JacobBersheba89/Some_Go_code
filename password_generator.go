package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Nastavení náhodného seedu
	rand.Seed(time.Now().UnixNano())

	// Načtení délky hesla od uživatele
	var length int
	fmt.Print("Zadejte délku hesla: ")
	fmt.Scan(&length)

	// Generování hesla
	password := generatePassword(length)

	fmt.Println("Vygenerované heslo:", password)
}

func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
