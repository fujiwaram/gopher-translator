package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/fujiwaram/gopherTranslator/translator"
)

func main() {
	flag.Parse()
	str := flag.Arg(0)
	message := flag.Arg(1)

	order, err := translate(str, message)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf(order)
}

func translate(str, message string) (string, error) {
	outlines, err := parseParam(str)
	if err != nil {
		return "", err
	}
	return outlines.Translate(message), nil
}

func parseParam(str string) (translator.FontOutlines, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return nil, err
	}
	outlines := translator.ParseFontOutlines(data)

	return outlines, nil
}
