package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func getURL(path string) *url.URL {
	url, err := url.Parse("https://www.strava.com/api/v3/")
	if err != nil {
		panic(err)
	}
	targetURL, targetErr := url.Parse(path)
	if targetErr != nil {
		panic(targetErr)
	}
	return targetURL
}

func GetAnswer(path string, object interface{}, client *http.Client) {
	url := getURL(path).String()
	println("URL: ", url)
	r, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	bodyBytes, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		panic(err2)
	}
	bodyString := string(bodyBytes)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	println(bodyString)
	println("Antwort-Code: ", r.StatusCode)

	err = json.NewDecoder(r.Body).Decode(object)
	if err != nil {
		panic(err)
	}
}

func WriteGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
