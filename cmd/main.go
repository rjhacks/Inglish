package main

import (
	"fmt"
	"strings"

	"../inglish"
)

var (
	userInput string = "loch"
	//"Go provides a familiar syntax for working with maps."

	engToIPADictFile = "../dicts/english-to-ipa.csv"
)

func main() {
	inglish.LoadDict(engToIPADictFile)
	ipa := inglish.EngToIPA(userInput)
	ing := inglish.IPAToIng(ipa)
	fmt.Printf("%s\n", strings.Join(ing, " "))
}
