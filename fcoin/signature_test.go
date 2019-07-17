package fcoin

import (
	"testing"
)

func TestPOSTSign(t *testing.T) {
	// according to: https://developer.fcoin.com/en.html?python#05e32f581f
	method := "POST"
	uri := "https://api.fcoin.com/v2/orders"
	ts := "1523069544359"
	args := "amount=100.0&price=100.0&side=buy&symbol=btcusdt&type=limit"
	key := "3600d0a74aa3410fb3b1996cca2419c8"
	got := Sign(method, uri, ts, args, key)

	expected := "DeP6oftldIrys06uq3B7Lkh3a0U="
	if got != expected {
		t.Fatalf("expected %s and got %v", expected, got)
	}
}

func TestGETSign(t *testing.T) {
	// according to: https://developer.fcoin.com/en.html?python#05e32f581f
	method := "GET"
	uri := "https://api.fcoin.com/v2/orders"
	ts := "1523069544359"
	args := "amount=100.0&price=100.0&side=buy&symbol=btcusdt&type=limit"
	key := "3600d0a74aa3410fb3b1996cca2419c8"
	got := Sign(method, uri, ts, args, key)

	expected := "We6jlnYGMm0EFI7OZd6n/Q7yyn0="
	if got != expected {
		t.Fatalf("expected %s and got %v", expected, got)
	}
}
