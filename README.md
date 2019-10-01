# Go Helper

Go Helper is a Golang library for various needs.

## Installation

Assumption Go is already installed and configured just use go get command.

```bash
go get github.com/ihsankurniawan/go-helper
```

## Usage

### Array
```go
import "github.com/ihsankurniawan/go-helper"

// Can be used for basic types such as int, uint, string, etc
var a int
var b []int
a = 1
b = []int{1, 2, 3}

rs, idx := helper.InArray(a, b) // returns true, 0

// Same like InArray function but not returning index
rs = helper.InArrayNoIndex(a, b) // returns true

// Same like InArray function but only for string, faster than InArray
rs = helper.InArrayContains("foo", []string{"foo", "bar"}) // returns true
```

### File
```go
import "github.com/ihsankurniawan/go-helper"

// Function to save multipart file into desired location
var someFile multipart.File

err := helper.SaveFileToDisk(someFile, "/tmp/someFile.jpg") // returns error if there any
```

### Image
```go
import "github.com/ihsankurniawan/go-helper"

// Convert base64 of image to standard img file
// Currently only support jpeg and png
// returns image and error if there any
img, err := helper.Base64ToImg("/tmp/temporary-file", "base64encoded of image file") 
```

### String
```go
import "github.com/ihsankurniawan/go-helper"

// Randomize string with given length and charset
// You can also use constant in this helper, available:
// STRING_NUMBER        = "0123456789"
// STRING_ALPHABET_LOWER= "abcdefghijklmnopqrstuvwxyz"
// STRING_ALPHABET_UPPER= "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// STRING_ALPHANUMERIC 	= STRING_ALPHABET_LOWER + STRING_ALPHABET_UPPER + STRING_NUMBER
str := helper.StringRandomWithCharset(5, "abcde") // return "aabad"

// Randomize string with lowercase alphabet, uppercase alphabet, number with given length
str = helper.StringRandom(5) // return "ab12F"

// Function to check whether string can be converted to json or not
isJson := helper.StringIsJSON("{}") // return true

// Function to convert int to IDR currency format
str = helper.FormatToIdrCurrency(100000) // return "100.000"

// Remove whitespace from string
str = helper.WhiteSpaceRemove(" lorem ipsum ") // return "loremipsum"
```

### Struct
```go
import "github.com/ihsankurniawan/go-helper"

// Function to check whether given struct is empty or not
isEmpty := helper.IsEmpty(struct{}{}) // return true 
```

### Telco Operator
##### This is only for Indonesian operator only
```go
import "github.com/ihsankurniawan/go-helper"

// Function to check whether given phone number is from Telkomsel operator
isTelkomsel := helper.TselPrefix("6281312341234") // return true 

// Function to check whether given phone number is from Indosat operator
isIndosat := helper.IsatPrefix("6285612341234") // return true 

// Function to check whether given phone number is from XL operator
isIndosat := helper.XlPrefix("6281712341234") // return true 

// Function to check whether given phone number is from Three operator
isIndosat := helper.ThreePrefix("6289512341234") // return true 
```

### Struct
```go
import "github.com/ihsankurniawan/go-helper"

// Function to get Time in Asia/Jakarta GMT+7 location based
now := helper.GetJakartaNowTime() // return time.Time
```

### URL
```go
import "github.com/ihsankurniawan/go-helper"

// Helper to extract json data from given URL, only available for GET method
type SomeStruct struct{
    Id    int    `json:"id"`
    Data  string `json:"data"`
}
var str SomeStruct

// var str will be filled if json data match
err := helper.GetJsonDataFromUrl("https://someurl.com", &str) // return error if there any

// Basically same like above function, only difference is URL returns XML data
err = helper.GetXmlDataFromUrl("https://someurl.com", &str) // return error if there any
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
