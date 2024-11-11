package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    var length int
    fmt.Print("Zadejte délku hesla: ")
    fmt.Scan(&length)

    password := generatePassword(length)
    fmt.Println("Vygenerované heslo:", password)
}

func generatePassword(length int) string {
    const (
        lowercase    = "abcdefghijklmnopqrstuvwxyz"
        uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        digits       = "0123456789"
        specialChars = "!@#$%^&*()-_=+<>?"
    )

    charset := lowercase // základní sada znaků je malá písmena

    // Zeptáme se uživatele, zda chce použít další typy znaků
    var useUppercase, useDigits, useSpecialChars bool
    fmt.Print("Chcete použít velká písmena? (Y/N): ")
    fmt.Scan(&useUppercase)
    fmt.Print("Chcete použít čísla? (Y/N): ")
    fmt.Scan(&useDigits)
    fmt.Print("Chcete použít speciální znaky? (Y/N): ")
    fmt.Scan(&useSpecialChars)

    // Přidání dalších sad znaků podle výběru uživatele
    if useUppercase {
        charset += uppercase
    }
    if useDigits {
        charset += digits
    }
    if useSpecialChars {
        charset += specialChars
    }

    password := make([]byte, length)
    for i := range password {
        password[i] = charset[rand.Intn(len(charset))]
    }
    return string(password)
}
