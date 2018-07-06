package main

import (
	"fmt"
	"flag"
	"net/http"
	"encoding/json"
)

type ExchangeRate struct {
	Currency string
	Buying string
	Selling string
}

const (
	defautBank = "cbe"
)

func main() {
	
	var port string
	flag.StringVar(&port, "port", "8888", "http port to run the server")
	flag.Parse()

	var exchangeList []ExchangeRate

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		bank := req.URL.Query().Get("bank")
		if bank == "" {
			bank = defautBank
		}

		fmt.Println("Selected Bank", bank)
		switch bank {
			case "cbe":
				exchangeList, _ = GetCBERate()
				js, _ := json.Marshal(exchangeList)
				fmt.Fprintf(w, string(js))
			// TODO: Add more banks here
			default:
				js, _ := json.Marshal([]ExchangeRate{})
				fmt.Fprintf(w, string(js))
		}
	})
	fmt.Println("Running at", port)
	http.ListenAndServe(":" + port, nil)
}