package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://rate.bot.com.tw/Pages/Static/UIP003.zh-TW.htm")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("[class=\"titleLeft\"]").Each(func(i int, s *goquery.Selection) {
		currency := s.Text()
		fmt.Printf("%s: ", currency)
		inCashRate := s.Next().Text()
		outCashRate := s.Next().Text()
		inRate := s.Next().Text()
		outRate := s.Next().Text()
		fmt.Printf("%s %s %s %s\n", inCashRate, outCashRate, inRate, outRate)

		//		}
	})

}
