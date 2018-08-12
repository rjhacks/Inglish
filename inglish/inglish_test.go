package inglish

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/rjhacks/Inglish/inglish"
	"github.com/stretchr/testify/assert"
)

func engToIng(engText string) string {
	return strings.Join(inglish.IPAToIng(inglish.EngToIPA(engText)), " ")
}

func startsWith(s string) string {
	return "^" + s + ".*"
}

func endsWith(s string) string {
	return ".*" + s + "$"
}

func hasMiddle(s string) string {
	return ".+" + s + ".+"
}

type testCase struct {
	in, out string
}

func assertMatches(t *testing.T, tc testCase) {
	r := tc.out
	if !strings.HasPrefix(r, "^") {
		r = "^" + r
	}
	if !strings.HasSuffix(r, "$") {
		r = r + "$"
	}
	assert.Regexp(t, regexp.MustCompile(r), engToIng(tc.in), "For input: "+tc.in)
}

func TestFunExamples(t *testing.T) {
	for _, tc := range []testCase{
		{"do", "doo"},
		{"daddy", "dadi"},
		{"odd", "od"},
		{"chair", "tsjeaar"},
		{"nature", "naytsjaar"},
		{"teach", "titsj"},
		{"gin", "djin"},
		{"edge", "edj"},
		{"queen", "kwhin"},
		{"enough", "inuhf"},
		{"off", "of"},
		{"photo", "fowtow"},
		{"of", "ov"},
		{"city", "siti"},
		{"pass", "pahs"},
		{"rose", "rowz"},
		{"session", "seshn"},
		{"emotion", "imowshn"},
		{"genre", "sjahnraa"},
		{"pleasure", "plesjaar"},
		{"equation", "ikwhaysjn"},
		{"drink", "dringk"},
		{"bell", "bel"},
		{"chocolate", "tsjoklaat"},
	} {
		assertMatches(t, tc)
	}
}

func TestConsonants(t *testing.T) {
	for _, tc := range []testCase{
		// p
		{"pen", startsWith("p")},
		{"spin", hasMiddle("p")},
		{"tip", endsWith("p")},
		// b
		{"but", startsWith("b")},
		{"web", endsWith("b")},
		// t
		{"two", startsWith("t")},
		{"sting", startsWith(".t")},
		{"bet", endsWith("t")},
		// d
		{"do", startsWith("d")},
		{"daddy", "d.d."},
		{"odd", ".d"},
		// tsj
		{"chair", startsWith("tsj")},
		{"nature", hasMiddle("tsj")},
		{"teach", endsWith("tsj")},
		// dj
		{"gin", startsWith("dj")},
		{"joy", startsWith("dj")},
		{"edge", endsWith("dj")},
		// k
		{"cat", startsWith("k")},
		{"kill", startsWith("k")},
		{"skin", hasMiddle("k")},
		{"queen", startsWith("k")},
		{"unique", endsWith("k")},
		{"thick", endsWith("k")},
		// g
		{"go", startsWith("g")},
		{"get", startsWith("g")},
		{"beg", endsWith("g")},
		// f
		{"fool", startsWith("f")},
		{"enough", endsWith("f")},
		{"leaf", endsWith("f")},
		{"off", ".f"},
		{"photo", startsWith("f")},
		// v
		{"voice", startsWith("v")},
		{"have", endsWith("v")},
		{"of", ".v"},
		// th
		{"thing", startsWith("th")},
		{"teeth", endsWith("th")},
		{"this", startsWith("th")},
		{"breathe", endsWith("th")},
		{"father", hasMiddle("th")},
		// s
		{"see", startsWith("s")},
		{"city", startsWith("s")},
		{"pass", endsWith("s")},
		// z
		{"zoo", startsWith("z")},
		{"rose", endsWith("z")},
		// sh
		{"she", startsWith("sh")},
		{"sure", startsWith("sh")},
		{"session", hasMiddle("sh")},
		{"emotion", hasMiddle("sh")},
		{"leash", endsWith("sh")},
		// sj
		{"genre", startsWith("sj")},
		{"pleasure", hasMiddle("sj")},
		{"beige", endsWith("sj")},
		{"equation", hasMiddle("sj")},
		{"seizure", hasMiddle("sj")},
		// h
		{"ham", startsWith("h")},
		{"hue", startsWith("h")},
		// m
		{"man", startsWith("m")},
		{"ham", endsWith("m")},
		// n
		{"no", startsWith("n")},
		{"tin", endsWith("n")},
		// ng
		{"ringer", hasMiddle("ng")},
		{"sing", endsWith("ng")},
		{"finger", hasMiddle("ng")},
		{"drink", hasMiddle("ng")},
		// l
		{"left", startsWith("l")},
		{"bell", "..l"},
		{"sable", endsWith("l")},
		{"please", startsWith(".l")},
		// r
		{"run", startsWith("r")},
		{"very", hasMiddle("r")},
		{"probably", startsWith(".r")},
		// w
		{"we", startsWith("w")},
		{"queen", startsWith(".w")},
		// y
		{"yes", startsWith("y")},
		// wh
		{"what", startsWith("wh")},
	} {
		assertMatches(t, tc)
	}
}
func TestVowels(t *testing.T) {
	for _, tc := range []testCase{
		// a
		{"lad", "lad"},
		{"bad", "bad"},
		{"cat", ".at"},
		// ah
		{"pass", "pahs"},
		{"path", "pahth"},
		{"sample", "sahmpl"},
		{"palm", "pahm"},
		{"father", startsWith("fah")},
		{"start", "staht"},
		{"arm", "ahm"},
		// ahr
		{"car", "kahr"},
		// aa
		{"coma", endsWith("aa")},
		{"about", startsWith("aab")},
		// aar
		{"winner", "whinaar"},
		{"letter", "letaar"},
		{"donor", endsWith("aar")},
		{"massacre", endsWith("aar")},
		// iaar
		{"near", "niaar"},
		{"deer", "diaar"},
		{"here", "hiaar"},
		// eaar
		{"square", "skwheaar"},
		{"mare", "meaar"},
		{"there", "theaar"},
		{"bear", "beaar"},
		// uaar
		{"cure", "kyuaar"},
		{"tour", "tuaar"},
		{"moor", "muaar"},
		// yuaar
		{"pure", "pyuaar"},
		// {"Europe", "Yuaarope"},  // TODO: enable when capitalization works.
		// o
		{"lot", "lot"},
		{"not", "not"},
		{"wasp", "whosp"},
		{"off", "of"},
		{"loss", "los"},
		{"cloth", "kloth"},
		{"long", "long"},
		{"dog", "dog"},
		{"chocolate", hasMiddle("ok")},
		// oh
		{"thought", startsWith("thoh")},
		{"law", "loh"},
		{"caught", startsWith("koh")},
		{"all", "ohl"},
		{"halt", "hohlt"},
		{"talk", "tohk"},
		{"north", "nohth"},
		{"force", "fohs"},
		{"sort", "soht"},
		{"warm", "whohm"},
		{"port", "poht"},
		// ohr
		{"tore", "tohr"},
		{"boar", "bohr"},
		// i
		{"kit", "kit"},
		{"sit", "sit"},
		{"happy", "hapi"},
		{"city", "siti"},
		// ii - this sound doesn't actually exist according to our dictionary, and so is "i".
		{"fleece", "flis"},
		{"see", "si"},
		{"meat", "mit"},
		// ay
		{"face", "fays"},
		{"date", "dayt"},
		{"day", "day"},
		{"pain", "payn"},
		{"whey", "whay"},
		{"rein", "rayn"},
		// e
		{"dress", "dres"},
		{"bed", "bed"},
		// ur
		{"nurse", "nurs"},
		{"burn", "burn"},
		{"bird", "burd"},
		{"herd", "hurd"},
		{"earth", "urth"},
		// uh
		{"strut", "struht"},
		{"run", "ruhn"},
		{"won", "whuhn"},
		{"flood", "fluhd"},
		{"put", "puht"}, // Note: homonym, spoken both "puht" and "put".
		// u
		{"hood", "hud"},
		{"foot", "fut"},
		// oo
		{"do", "doo"},
		{"goose", "goos"},
		{"through", "throo"},
		{"you", "yoo"},
		{"threw", "throo"},
		{"yew", "yoo"},
		// yoo
		{"cute", "kyoot"},
		{"dew", "dyoo"},
		{"ewe", "yoo"},
		// ai
		{"price", "prais"},
		{"flight", "flait"},
		{"mice", "mais"},
		{"my", "mai"},
		{"wise", "whaiz"},
		{"high", "hai"},
		// oi
		{"choice", endsWith("ois")},
		{"boy", "boi"},
		{"hoist", "hoist"},
		// ow
		{"goat", "gowt"},
		{"no", "now"},
		{"toe", "tow"},
		{"tow", "tow"},
		{"soap", "sowp"},
		{"folk", "fowk"},
		{"soul", "sowl"},
		{"roll", "rowl"},
		{"cold", "kowld"},
		// au
		{"mouth", "mauth"},
		{"about", "aabaut"},
		{"house", "hauz"},
		{"now", "nau"},
		{"trout", "traut"},
	} {
		assertMatches(t, tc)
	}
}

func TestWordsWithFunnyCharacters(t *testing.T) {
	for _, tc := range []testCase{
		{"pronunciation", "praanuhnsiayshn"},
	} {
		assertMatches(t, tc)
	}
}

func TestMain(m *testing.M) {
	inglish.LoadDict("../dicts/english-to-ipa.csv")
	os.Exit(m.Run())
}
