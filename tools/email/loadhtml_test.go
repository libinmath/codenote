package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"testing"
)

func TestGetLocalHtml(t *testing.T) {
	res, err := http.Get("https://github.com/libinmath/codenote/blob/master/tools/charts/html/bar.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	h, _ := doc.Html()
	fmt.Printf("%+v\n", h)
}
