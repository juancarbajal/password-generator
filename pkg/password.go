package password

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

const THREADS = 4

// Generate a password with a specific size, consider alphanumeric and special characters
func generatePartial(size int, totalChars []rune, partChan chan string) {
	defer wg.Done()
	sizeTotalChars := len(totalChars)
	pass := ""
	for i := 0; i < size; i++ {
		pass = pass + string(totalChars[rand.Intn(sizeTotalChars)])
	}
	partChan <- pass
}

func getTotalChars(includeDigits bool, includeSpecial bool) []rune {
	chars := "abcdefghijklmnopqrstuvwxyz"
	digits := "1234567890"
	special := "~!@#$%^&*()_+`-={}|[]:<>?,./"
	totalChars := chars + strings.ToUpper(chars)
	if includeDigits {
		totalChars = totalChars + digits
	}
	if includeSpecial {
		totalChars = totalChars + special
	}
	return []rune(totalChars)
}
func Generate(size int, includeDigits bool, includeSpecial bool) string {
	totalChars := getTotalChars(includeDigits, includeSpecial)
	rand.Seed(time.Now().UnixNano())
	partsChan := make(chan string, THREADS)
	wg.Add(THREADS)
	charsToGenerate, charsToGenerateExtra := int(size/THREADS), size%THREADS
	for i := 1; i <= THREADS; i++ {
		if i == THREADS {
			charsToGenerate += charsToGenerateExtra
		}
		go generatePartial(charsToGenerate, totalChars, partsChan)
	}
	wg.Wait()
	close(partsChan)

	password := ""
	for part := range partsChan {
		password += part
	}
	return password
}
