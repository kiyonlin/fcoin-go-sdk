package fcoin

import (
	"testing"
)

func TestEncode(t *testing.T) {
	reader, err := encode(&GetOrdersArgs{
		States: "submitted",
		Symbol: "btcusdt",
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"states":"submitted","symbol":"btcusdt"}`
	got := string(reader.Bytes())
	if expected != got {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestStructToMap(t *testing.T) {
	args := &GetOrdersArgs{
		States: "submitted",
		Symbol: "btcusdt",
	}

	argValues := structToMap(args)
	got := argValues.Encode()
	expected := "states=submitted&symbol=btcusdt"
	if got != expected {
		t.Fatalf("expected %s and got %v", expected, got)
	}
}
