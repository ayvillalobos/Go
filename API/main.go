package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Copyright       string `json:"copyright"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Title           string `json:"title"`
	Url             string `json:"url"`
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

	fmt.Println(Response.Copyright)
	fmt.Println(Response.Date)
	fmt.Println(Response.Explanation)
	fmt.Println(Response.Hdurl)
	fmt.Println(Response.Media_type)
	fmt.Println(Response.Service_version)
	fmt.Println(Response.Title)
	fmt.Println(Response.Url)

}
