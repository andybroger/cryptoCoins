package main

import "testing"

func TestCurrency(t *testing.T) {
	currency := currency{Symbol: "BTC", Price: Price{Value: 1.44}}
	got := currency.PriceString()
	want := "1.44"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
