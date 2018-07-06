package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const cbeURL = "http://www.combanketh.et/More/CurrencyRate.aspx"


func GetCBERate() ([]ExchangeRate, error) {
	var exchangeRateList []ExchangeRate

	c := colly.NewCollector()

	c.OnHTML("#dnn_ctr535_ModuleContent > table:nth-child(5)", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var row []string
			el.ForEach("td", func(_ int, ell *colly.HTMLElement) {
				row = append(row, ell.ChildText("span"))
			})
			if len(row) > 0 {
				exchangeRateList = append(exchangeRateList, ExchangeRate{row[0], row[1], row[2]})
			}			
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(cbeURL)

	return exchangeRateList, nil

}