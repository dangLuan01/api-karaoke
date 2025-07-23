package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/Luzifer/go-openssl"
)

type APIResponse struct {
    Status     string `json:"status"`
    Data       string `json:"data"`
    Pagination struct {
        Total int `json:"total"`
        Page  int `json:"page"`
        Limit int `json:"limit"`
    } `json:"pagination"`
}
// Map of Vietnamese diacritics to non-diacritic equivalents
var diacriticMap = map[rune]rune{
	// A vowels
	'À': 'A', 'Á': 'A', 'Ả': 'A', 'Ã': 'A', 'Ạ': 'A',
	'à': 'a', 'á': 'a', 'ả': 'a', 'ã': 'a', 'ạ': 'a',
	// Ă vowels
	'Ă': 'A', 'Ắ': 'A', 'Ằ': 'A', 'Ẳ': 'A', 'Ẵ': 'A', 'Ặ': 'A',
	'ă': 'a', 'ắ': 'a', 'ằ': 'a', 'ẳ': 'a', 'ẵ': 'a', 'ặ': 'a',
	// Â vowels
	'Â': 'A', 'Ấ': 'A', 'Ầ': 'A', 'Ẩ': 'A', 'Ẫ': 'A', 'Ậ': 'A',
	'â': 'a', 'ấ': 'a', 'ầ': 'a', 'ẩ': 'a', 'ẫ': 'a', 'ậ': 'a',
	// E vowels
	'È': 'E', 'É': 'E', 'Ẻ': 'E', 'Ẽ': 'E', 'Ẹ': 'E',
	'è': 'e', 'é': 'e', 'ẻ': 'e', 'ẽ': 'e', 'ẹ': 'e',
	// Ê vowels
	'Ê': 'E', 'Ế': 'E', 'Ề': 'E', 'Ể': 'E', 'Ễ': 'E', 'Ệ': 'E',
	'ê': 'e', 'ế': 'e', 'ề': 'e', 'ể': 'e', 'ễ': 'e', 'ệ': 'e',
	// I vowels
	'Ì': 'I', 'Í': 'I', 'Ỉ': 'I', 'Ĩ': 'I', 'Ị': 'I',
	'ì': 'i', 'í': 'i', 'ỉ': 'i', 'ĩ': 'i', 'ị': 'i',
	// O vowels
	'Ò': 'O', 'Ó': 'O', 'Ỏ': 'O', 'Õ': 'O', 'Ọ': 'O',
	'ò': 'o', 'ó': 'o', 'ỏ': 'o', 'õ': 'o', 'ọ': 'o',
	// Ô vowels
	'Ô': 'O', 'Ố': 'O', 'Ồ': 'O', 'Ổ': 'O', 'Ỗ': 'O', 'Ộ': 'O',
	'ô': 'o', 'ố': 'o', 'ồ': 'o', 'ổ': 'o', 'ỗ': 'o', 'ộ': 'o',
	// Ơ vowels
	'Ơ': 'O', 'Ớ': 'O', 'Ờ': 'O', 'Ở': 'O', 'Ỡ': 'O', 'Ợ': 'O',
	'ơ': 'o', 'ớ': 'o', 'ờ': 'o', 'ở': 'o', 'ỡ': 'o', 'ợ': 'o',
	// U vowels
	'Ù': 'U', 'Ú': 'U', 'Ủ': 'U', 'Ũ': 'U', 'Ụ': 'U',
	'ù': 'u', 'ú': 'u', 'ủ': 'u', 'ũ': 'u', 'ụ': 'u',
	// Ư vowels
	'Ư': 'U', 'Ứ': 'U', 'Ừ': 'U', 'Ử': 'U', 'Ữ': 'U', 'Ự': 'U',
	'ư': 'u', 'ứ': 'u', 'ừ': 'u', 'ử': 'u', 'ữ': 'u', 'ự': 'u',
	// Y vowels
	'Ỳ': 'Y', 'Ý': 'Y', 'Ỷ': 'Y', 'Ỹ': 'Y', 'Ỵ': 'Y',
	'ỳ': 'y', 'ý': 'y', 'ỷ': 'y', 'ỹ': 'y', 'ỵ': 'y',
	// Đ character
	'Đ': 'D', 'đ': 'd',
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}

func GetHttpAndDecrypto(url string ) ([]byte, error) {
	resp, err := http.Get(url)
	
	if err != nil {
		fmt.Println("Error fetching data:", err)

		return nil, err
	}
	defer resp.Body.Close()
	
	data , err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println("Error reading response body:", err)

		return nil, err
	}

	var apiResponse APIResponse
    if err := json.Unmarshal(data, &apiResponse); err != nil {
        return nil, fmt.Errorf("error parsing JSON response: %v", err)
    }

	o := openssl.New()
	secret := GetEnv("SECRET", "123456")
	dec, err := o.DecryptBytes(secret, []byte(apiResponse.Data))
	if err != nil {
		fmt.Println("Error decrypting data:", err)

		return nil, err
	}

	return dec, nil
}

func removeDiacritics(r rune) rune {
	if replacement, exists := diacriticMap[r]; exists {
		return replacement
	}
	return r
}

func ToAbbreviation(input string) string {
	// Split the input string into words
	words := strings.Fields(input)
	
	var abbr strings.Builder

	// Take the first letter of each word, ignoring diacritics
	for _, word := range words {
		if len(word) > 0 {
			// Convert to runes to handle Vietnamese characters
			runes := []rune(word)
			// Get the first rune and convert to lowercase
			firstLetter := removeDiacritics(runes[0])
			abbr.WriteRune(unicode.ToLower(firstLetter))
		}
	}
	return abbr.String()
}