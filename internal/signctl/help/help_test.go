package help

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	S string `json:"s"`
}

func TestJsonMarshal(t *testing.T) {
	a := &A{
		S: "fff",
	}
	d, _ := json.Marshal(a)
	fmt.Printf("%s\n", d)
}

func TestJsonEncode(t *testing.T) {
	fj := &formatterJson{}
	raw := "abc=def,123=456"
	res, err := fj.Encode(raw)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %s\n", res)

	raw = ",123=456"
	res, err = fj.Encode(raw)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %s\n", res)

	raw = ",123=456,abc=def"
	res, err = fj.Encode(raw)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %s\n", res)

	raw = "123=456,,abc=def"
	res, err = fj.Encode(raw)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %s\n", res)

	raw = "123=456,"
	res, err = fj.Encode(raw)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %s\n", res)
}
