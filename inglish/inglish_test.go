package inglish

import (
	"os"
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

func TestFunExamples(t *testing.T) {
	for _, tt := range []struct {
		in, out string
	}{
		{"do", "doo"},
		{"daddy", "dadi"},
		//{"odd", "od"},
		//{"chair", "tsjair"},
		//{"nature", "natsjur"},
		//{"teach", "teatsj"},
		{"gin", "djin"},
		{"edge", "edj"},
		//{"queen", "kween"},
		//{"enough", "enuf"},
		//{"off", "of"},
		//{"photo", "foto"},
		//{"of", "ov"},
		{"city", "siti"},
		//{"pass", "pas"},
		//{"rose", "roz"},
		//{"session", "seshion"},
		//{"emotion", "emoshion"},
		//{"genre", "sjenre"},
		//{"pleasure", "pleasjur"},
		//{"equation", "equasjn"},
		{"drink", "dringk"},
		{"bell", "bel"},
		//{"chocolate", "tsjoklet"},
	} {
		assert.Regexp(t, tt.out, engToIng(tt.in), "For input: "+tt.in)
	}
}

func TestConsonants(t *testing.T) {
	for _, tt := range []struct {
		in, out string
	}{
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
		{"do", "d."},
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
		{"go", "g."},
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
		{"pass", "..s"},
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
		{"we", "w."},
		{"queen", startsWith(".w")},
		// y
		{"yes", startsWith("y")},
		// wh
		{"what", startsWith("wh")},
	} {
		assert.Regexp(t, tt.out, engToIng(tt.in), "For input: "+tt.in)
	}
}
func TestVowels(t *testing.T) {
	for _, tt := range []struct {
		in, out string
	}{
		// a
		{"lad", "lad"},
		{"bad", "bad"},
		{"cat", ".at"},
		// ah
		{"pass", "pahs"},
		{"path", "pahth"},
		{"sample", "sahmpl"},
		{"palm", "pahm"},
		{"father", "fahth.."},
		// aa
		{"coma", endsWith("aa")},
		{"about", startsWith("aab")},
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
		// ar
		{"letter", "letaa"},
		{"winner", "whinaa"},
		{"donor", endsWith("aar")},
		{"massacre", endsWith("aar")},
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
		{"about", "abaut"},
		{"house", "hauz"},
		{"now", "nau"},
		{"trout", "traut"},
	} {
		assert.Regexp(t, tt.out, engToIng(tt.in), "For input: "+tt.in)
	}
}

func TestMain(m *testing.M) {
	inglish.LoadDict("../dicts/english-to-ipa.csv")
	os.Exit(m.Run())
}
