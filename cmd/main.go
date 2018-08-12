package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../inglish"
)

const (
	engToIPADictFile = "../dicts/english-to-ipa.csv"
)

func main() {
	inglish.LoadDict(engToIPADictFile)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter English: ")
	userInput, _ := reader.ReadString('\n')

	ipa := inglish.EngToIPA(userInput)
	ing := inglish.IPAToIng(ipa)
	fmt.Printf("Thiz iz Inglish: %s\n", strings.Join(ing, " "))
}
