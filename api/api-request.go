package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//Get request
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=dogecoin&vs_currencies=gbp")

	if err != nil {
		fmt.Printf("error occured, %s", err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	//Post request
	Data := map[string]string{"username": "wuve", "colour": "pink"}
	json, _ := json.Marshal(Data)                                                                    //Marshal map into json
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(json)) //Create request(post endpoint)/header(content type)/buffer json to byte

	if err != nil {
		fmt.Printf("error occured, %s", err)
	}

	content, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
