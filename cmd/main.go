package main

import (
	"fmt"
	"strings"

	"../inglish"
)

var (
	userInput string = "loch"
	//"Go provides a familiar syntax for working with maps."
	
	engToIpaDictFile = "../dicts/english-to-ipa.csv"
)

func main() {
	inglish.LoadDict(engToIpaDictFile)
	ipa := inglish.EngToIpa(userInput)
	ing := inglish.IpaToIng(ipa)
	fmt.Printf("%s\n", strings.Join(ing, " "))
}