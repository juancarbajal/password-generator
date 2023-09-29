package password
import (
	"math/rand"
	"strings"
	"time"
)

// Generate a password with a specific size, consider alphanumeric and special characters
func Generate(size int, includeDigits bool, includeSpecial bool)  string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	digits := "1234567890"
	special := "~!@#$%^&*()_+`-={}|[]:<>?,./"
	totalChars := chars + strings.ToUpper(chars) 
	if includeDigits {
		totalChars = totalChars+ digits 
	}
	if includeSpecial {
		totalChars = totalChars + special
	}
	sizeTotalChars := len(totalChars)
	pass := ""
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		pass = pass + string([]rune(totalChars)[rand.Intn(sizeTotalChars)])
	}
	return pass 
}
