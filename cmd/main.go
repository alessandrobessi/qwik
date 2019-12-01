package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alessandrobessi/qwik/internal/wikistruct"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

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

	var url_search = "https://" + lang + ".wikipedia.org/w/api.php?action=query&list=search&srsearch=" + q + "&format=json"

	resp_search, err := http.Get(url_search)
	if err != nil {
		panic(err)
	}
	body_search, err := ioutil.ReadAll(resp_search.Body)
	if err != nil {
		panic(err.Error())
	}

	var searchResult wikistruct.SearchResult
	err = json.Unmarshal(body_search, &searchResult)
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

	url_summary := "https://" + lang + ".wikipedia.org/api/rest_v1/page/summary/" + choice

	resp_summary, err := http.Get(url_summary)
	if err != nil {
		panic(err)
	}
	body_summary, err := ioutil.ReadAll(resp_summary.Body)
	if err != nil {
		panic(err.Error())
	}

	var page wikistruct.Page
	err = json.Unmarshal(body_summary, &page)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(page.Extract)
	fmt.Println(page.ContentUrls.Desktop.Page)

}
