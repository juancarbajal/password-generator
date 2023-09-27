package password
import (
	"math/rand"
	"strings"
	"time"
)

// Generate a password with a specific size, consider alphanumeric and special characters
func Generate(size int)  string {
	chars := "abcdefghijklmnopqrstuvwxyz"
	digits := "1234567890"
	special := "@#%&/(){}!'[]?"
	totalChars := chars + strings.ToUpper(chars) + special + digits
	sizeTotalChars := len(totalChars)
	pass := ""
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		pass = pass + string([]rune(totalChars)[rand.Intn(sizeTotalChars)])
	}
	return pass 
}
