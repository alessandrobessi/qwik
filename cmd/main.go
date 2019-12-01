package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/alessandrobessi/qwik/internal"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func request(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	return body
}

func main() {
	var err error
	var lang string

	flag.StringVar(&lang, "lang", "en", "Wikipedia Language.")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage:\nqwik [-lang] query")
		os.Exit(1)
	}

	var q string
	if len(flag.Args()) > 1 {
		q = strings.Join(flag.Args()[:], "%20")
	} else {
		q = flag.Args()[0]
	}

	url := "https://" + lang + ".wikipedia.org/w/api.php?action=query&list=search&srsearch=" + q + "&format=json"
	body := request(url)
	var searchResult wikistruct.SearchResult
	err = json.Unmarshal(body, &searchResult)
	if err != nil {
		panic(err.Error())
	}

	candidates := make([]string, len(searchResult.Query.Search))
	for i, sr := range searchResult.Query.Search {
		candidates[i] = sr.Title
	}

	prompt := promptui.Select{
		Label: "Select a page",
		Items: candidates,
	}

	_, choice, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	url = "https://" + lang + ".wikipedia.org/api/rest_v1/page/summary/" + choice
	body = request(url)

	var page wikistruct.Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(page.Extract)
	fmt.Println(page.ContentUrls.Desktop.Page)

}
