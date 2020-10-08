// Cologne phonetics is an algorithm related to soundex but optimized for the German language
// It calculates a phonetic code for a given sequence of words
package colphonetics

import "strings"

var (
	replacer = strings.NewReplacer(
		"ä", "a",
		"ö", "o",
		"ü", "u",
		"ph", "f",
		"ç", "c",
		"é", "e",
		"è", "e",
		"ê", "e",
		"à", "a",
		"á", "a",
		"â", "a",
		"ë", "e",
		"í", "i",
		"ì", "i",
		"ß", "ss",
	)

	charToScore = map[int32]string{
		97:  "0",  //a
		98:  "1",  //b
		99:  "4",  //c
		100: "2",  //d
		101: "0",  //e
		102: "3",  //f
		103: "4",  //g
		104: "",   //h
		105: "0",  //i
		106: "0",  //j
		107: "4",  //k
		108: "5",  //l
		109: "6",  //m
		110: "6",  //n
		111: "0",  //o
		112: "1",  //p
		113: "4",  //q
		114: "7",  //r
		115: "8",  //s
		116: "2",  //t
		117: "0",  //u
		118: "3",  //v
		119: "3",  //w
		120: "48", //x
		121: "0",  //y
		122: "8",  //z
	}

	exceptionsFollowing = map[string]string{
		"sc": "8",
		"zc": "8",
		"cx": "8",
		"kx": "8",
		"qx": "8",
	}

	exceptionsLeading = map[string]string{
		"ca": "4",
		"ch": "4",
		"ck": "4",
		"cl": "4",
		"co": "4",
		"cq": "4",
		"cu": "4",
		"cx": "4",
		"dc": "8",
		"ds": "8",
		"dz": "8",
		"tc": "8",
		"ts": "8",
		"tz": "8",
	}
)

// Code returns the phonetic code for a given string.
func Code(s string) string {
	s = strings.ToLower(s)

	s = replacer.Replace(s)
	l := len(s)
	r := make([]string, l, l)

	for i, char := range s {
		if i == 0 && len(s) > 1 && char == 98 && s[i+1] == 115 {
			r[i] = "4"
		}

		if len(s) > i+3 && exceptionsLeading[s[i:i+2]] != "" {
			r[i] = exceptionsLeading[s[i:i+2]]
		}

		if i > 2 && exceptionsFollowing[s[i-1:i+1]] != "" {
			r[i] = exceptionsFollowing[s[i-1:i+1]]
		}

		if r[i] == "" {
			r[i] = charToScore[char]
		}
	}

	// remove duplicate codes and "0"s except for first char
	for j := 1; j < l; j++ {
		if r[j] == "0" || r[j-1] == r[j] {
			r[j] = ""
		}
	}

	return strings.Join(r, "")
}
