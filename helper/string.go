package helper

import (
	"encoding/json"
	"math/rand"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/ihsankurniawan/go-helper/constant"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

/*
	Randomize string with given length and charset
	Ex: StringRandomWithCharset(5, "abcde") // return "aabad"
*/
func StringRandomWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

/*
	Randomize string with lowercase alphabet, uppercase alphabet, number with given length
	Ex: StringRandom(5) // return "ab12F"
*/
func StringRandom(length int) string {
	return StringRandomWithCharset(length, constant.STRING_ALPHANUMERIC)
}

/*
	Function to check whether string can be converted to json or not
*/
func StringIsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

/*
	Function to convert int to IDR currency format
	Ex: FormatToIdrCurrency(100000) // return "100.000"
*/
func FormatToIdrCurrency(n int) string {
	var s []string
	is := fmt.Sprintf("%d", n)

	for i := len(is); i > 0; i -= 3 {
		switch {
		case i >= 3:
			s = append([]string{is[i-3 : i]}, s...)
		case i < 3:
			s = append([]string{is[:i]}, s...)
		}
	}

	return strings.Join(s, ".")
}

/*
	Remove whitespace from string
	Ex: WhiteSpaceRemove(" lorem ipsum ") // return "loremipsum"
*/
func WhiteSpaceRemove(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}