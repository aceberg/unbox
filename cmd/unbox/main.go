package main

import (
	"encoding/json"
	"log"

	"github.com/aceberg/unbox/pkg/vless"
)

func main() {
	raw := ""

	res, _ := vless.ParseVLESS(raw)
	log.Println(res)

	data, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(data))
}
