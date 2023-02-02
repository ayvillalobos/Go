package main

import (
	"encoding/json"
	"fmt"

	"io"
	"log"
	"net/http"
)

type Response struct {
	Hdurl string `json: hdurl`
	Date  string `json: date`
}

func main() {

	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=SzhFxbRoaCdkuY0aXhQk5j0luq3Kajm5VAqutSpC")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	var Response Response
	jsonErr := json.Unmarshal(body, &Response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(Response.Date)
	fmt.Println(Response.Hdurl)

}
