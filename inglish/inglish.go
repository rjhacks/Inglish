package inglish

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	engToIpaDict map[string]string

	regNonAlphaNumeric *regexp.Regexp
	regUselessIpa      *regexp.Regexp

	// Replaces Kirschenbaum IPA phrases with Inglish phrases.
	// Longer IPA phrases come first, to ensure we always replace the longest
	// possible phrases. The Replacer ensures no overlapping replaces.
	ipaToIngReplacer = strings.NewReplacer(
		"tS", "tsj",
		"dZ", "dj",
		"Ng", "ng",
		"&", "a", // Kirschenbaum calls this: "near-open front unrounded vowel". The "a" in "dad".
		"D", "th",
		"I", "i",
		"N", "ng",
		"R", "r",
		"S", "sh",
		"T", "th",
		"Z", "sj",
		"j", "y",
		"w", "wh",
	)
)

// LoadDict will load a dictionary of English Received Pronounciation (RP) to IPA.
// The given dictionary must render its IPA in Kirschenbaum ASCII format.
func LoadDict(filePath string) {

	engToIpaDict = make(map[string]string)

	dictfile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dictfile.Close()

	scanner := bufio.NewScanner(dictfile)
	for scanner.Scan() {
		wordDef := strings.Split(scanner.Text(), ";")
		engToIpaDict[wordDef[0]] = wordDef[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sanitizeEng(word string, decapitalize bool) string {
	// Remove capitalization.
	if decapitalize {
		word = strings.ToLower(word)
	}

	// Remove all non-alphanumeric characters.
	if regNonAlphaNumeric == nil {
		var err error
		regNonAlphaNumeric, err = regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
	}
	word = regNonAlphaNumeric.ReplaceAllString(word, "")

	return word
}

func EngToIpa(engText string) []string {
	engWords := strings.Split(engText, " ")
	ipaWords := make([]string, len(engWords))
	for i, engWord := range engWords {
		ipaWord, found := engToIpaDict[sanitizeEng(engWord, true)]
		if !found {
			// Try again, but leaving capitalization intact.
			ipaWord, found = engToIpaDict[sanitizeEng(engWord, false)]
		}
		if !found {
			// Give up, leave this word in English, but quote it to show the failure.
			ipaWord = "```" + engWord + "```"
		}
		if ipaWord == "" {
			continue
		}
		ipaWords[i] = ipaWord
	}
	return ipaWords
}

func sanitizeIpa(word string) string {
	// Remove all useless Ipa markings.
	if regUselessIpa == nil {
		var err error
		// Useless-to-us Ipa includes:
		// - "'", denotes "primary stress"
		regUselessIpa, err = regexp.Compile("[']")
		if err != nil {
			log.Fatal(err)
		}
	}
	word = regUselessIpa.ReplaceAllString(word, "")
	return word
}

func IpaToIng(ipaWords []string) []string {
	ingWords := make([]string, len(ipaWords))
	for i, ipaWord := range ipaWords {
		word := sanitizeIpa(ipaWord)
		word = ipaToIngReplacer.Replace(word)
		ingWords[i] = word
	}
	return ingWords
}
