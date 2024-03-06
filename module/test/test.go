package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.ReadFile("./module/test/ico.png")
	if err != nil {
		fmt.Println("read fail", err)
	}
	fmt.Println(hex.EncodeToString(f))
	content := "package main\n\nimport (\n\t\"encoding/hex\"\n\t\"fyne.io/fyne/v2\"\n)\nvar Logo,_ = hex.DecodeString(\"" + hex.EncodeToString(f) + "\")\nvar ResourceLOGOPng = &fyne.StaticResource{\n\tStaticName:    \"LOGO.png\",\n\tStaticContent: Logo,\n}"
	var _ = ioutil.WriteFile("./module/test/pic.go", []byte(content), 0644)
}
