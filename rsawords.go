package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	f, err := os.Open("rsa2008exhibitors.html")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("td.bwcellpaddingleft0").Each(func(i int, s *goquery.Selection) {
		span, ok := s.Attr("colspan")
		if !ok || span != "2" {
			return
		}
		desc := strings.TrimSpace(s.Text())
		fmt.Printf("%s\n", desc)
	})

	// 2014, 2015
	// doc.Find("td.description").Each(func(i int, s *goquery.Selection) {
	// 	desc := strings.TrimSpace(s.Text())
	// 	fmt.Printf("%s\n", desc)
	// })
}
