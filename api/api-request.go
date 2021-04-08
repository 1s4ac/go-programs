package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=dogecoin&vs_currencies=gbp")

	if err != nil {
		fmt.Printf("error occured, %s", err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
