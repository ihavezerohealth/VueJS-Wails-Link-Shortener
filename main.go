package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

//go:embed frontend/src/main.js
var js string

//go:embed frontend/src/assets/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Learning Wails",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(shortenLink)
	app.Run()
}

func shortenLink(r *http.Request, longURL string) string {
	fmt.Println("Function called")
	client := &http.Client{}
	var data = strings.NewReader(`{ "` + longURL + `": "https://dev.bitly.com", "domain": "bit.ly", "group_guid": "Ba1bc23dE4F" }`)
	req, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer {TOKEN}")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	return string("works ig")
	/*
		form := r.ParseForm()
		resp, err := http.Post(link)
		if err != nil {
			panic(err.Error())
		}
		resp.Header.Set("Authorization", "Bearer"+apiToken)
	*/
}
