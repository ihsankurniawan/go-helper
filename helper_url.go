package helper

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
	"io"
)

/* 
	Helper to extract json data from given URL, should be using GET method
	var str SomeStruct
	Ex: GetJson("https://someurl.com", &str)
*/
func GetJsonDataFromUrl(url string, target interface{}) error {
	myClient := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Connection", "close")
	
	r, err := myClient.Do(req)
	if err != nil {
		return err
	}
	//defer r.Body.Close()
	defer func() {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

/* 
	Helper to extract XML data from given URL, should be using GET method
	var str SomeStruct
	Ex: GetXml("https://someurl.com", &str)
*/
func GetXmlDataFromUrl(url string, target interface{}) error {
	myClient := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Connection", "close")
	
	r, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return xml.NewDecoder(r.Body).Decode(target)
}