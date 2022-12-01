package main

import (
	"embed"
	"encoding/base64"
	"fmt"
)

//go:embed 99-low-latency-hooks.sh
var f embed.FS

func main() {
	data, _ := f.ReadFile("99-low-latency-hooks.sh")
	s := base64.StdEncoding.EncodeToString(data)
	fmt.Print(s)
}
