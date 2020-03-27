// Wirtten 2018 by Andy Broger.
// MIT License

// Get Crypto price from cryptocompare.com and print it to terminal.
// the first argument the script recives will get fetched.
// exp: `./cryptoPrice ETH` returns: `2018/08/05 21:53:08 - 1 ETH is 406.14 USD`
// this recipes is part of my 52 weeks project.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Price is the struct to unmarshall the json response
type Price struct {
	Value float64 `json:"USD"`
}

// Currency defines a cryptocoin
type currency struct {
	Symbol string
	Price
}

// PriceString returns the Price formatet as String
func (c *currency) PriceString() string {
	return strconv.FormatFloat(c.Value, 'f', 2, 64)
}

func main() {
	c := currency{Symbol: "BTC"}
	clean := false
	// check if a symbol is provided as argument and change symbol to that argument
	if len(os.Args) >= 2 {
		c.Symbol = os.Args[1]
	}
	if len(os.Args) >= 3 {
		if os.Args[2] == "--clean" {
			clean = true
		}
	}
	// http.get the url returns a resp *Response or a error
	resp, err := http.Get("https://min-api.cryptocompare.com/data/price?fsym=" + c.Symbol + "&tsyms=USD")
	// check if error is not nil, else print error message
	if err != nil {
		log.Fatalln("Error fetching price:", err)
	}
	// close the body
	defer resp.Body.Close()
	// ioutil.Readall takes a io.Reader and returns a []byte or err
	b, err := ioutil.ReadAll(resp.Body)
	// check errors return from ioutil
	if err != nil {
		log.Fatalln("Error parsing body:", err)
	}
	// unmarshal the response
	err = json.Unmarshal(b, &c.Price)
	// check if error occured while unmarshalling
	if err != nil {
		log.Fatalln("Error unmarshalling json:", err)
	}
	// Check if clean or pretty print
	if clean {
		log.Println(c.PriceString())
	} else {
		log.Println("- 1 " + c.Symbol + " is " + c.PriceString() + " USD")
	}
}
