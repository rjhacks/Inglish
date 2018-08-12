package inglish

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	engToIPADict map[string]string

	regNonAlphaNumeric *regexp.Regexp
	regUselessIPA      *regexp.Regexp

	// Replaces Kirschenbaum IPA phrases with Inglish phrases.
	// Longer IPA phrases come first, to ensure we always replace the longest
	// possible phrases. The Replacer ensures no overlapping replaces.
	ipaToIngReplacer = strings.NewReplacer(
		// Consonants
		"tS", "tsj",
		"dZ", "dj",
		"Ng", "ng",
		"D", "th",
		"N", "ng",
		"R", "r",
		"S", "sh",
		"T", "th",
		"Z", "sj",
		"j", "y",
		"w", "wh",

		// Vowels
		"@U", "ow",
		"eI", "ay",
		"&", "a", // Kirschenbaum calls this: "near-open front unrounded vowel". The "a" in "dad".
		"A", "ah",
		"@", "aa",
		"0", "o",
		"O", "oh",
		"I", "i",
		"3", "ur",
		"V", "uh",
		"U", "u",
		"u", "oo",
	)
)

// LoadDict will load a dictionary of English Received Pronounciation (RP) to IPA.
// The given dictionary must render its IPA in Kirschenbaum ASCII format.
func LoadDict(filePath string) {

	engToIPADict = make(map[string]string)

	dictfile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dictfile.Close()

	scanner := bufio.NewScanner(dictfile)
	for scanner.Scan() {
		wordDef := strings.Split(scanner.Text(), ";")
		engToIPADict[wordDef[0]] = wordDef[1]
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

func EngToIPA(engText string) []string {
	engWords := strings.Split(engText, " ")
	ipaWords := make([]string, len(engWords))
	for i, engWord := range engWords {
		ipaWord, found := engToIPADict[sanitizeEng(engWord, true)]
		if !found {
			// Try again, but leaving capitalization intact.
			ipaWord, found = engToIPADict[sanitizeEng(engWord, false)]
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

func sanitizeIPA(word string) string {
	// Remove all useless IPA markings.
	if regUselessIPA == nil {
		var err error
		// Useless-to-us IPA includes:
		// - "'", denotes "primary stress"
		// - ".", denotes an open-back rounded vowel.
		// - "\", denotes an alveolar approximant.
		regUselessIPA, err = regexp.Compile("['.\\\\]")
		if err != nil {
			log.Fatal(err)
		}
	}
	word = regUselessIPA.ReplaceAllString(word, "")
	return word
}

func IPAToIng(ipaWords []string) []string {
	ingWords := make([]string, len(ipaWords))
	for i, ipaWord := range ipaWords {
		word := sanitizeIPA(ipaWord)
		word = ipaToIngReplacer.Replace(word)
		ingWords[i] = word
	}
	return ingWords
}
