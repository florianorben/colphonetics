// Cologne phonetics is an algorithm related to soundex but optimized for the German language
// It calculates a phonetic code for a given sequence of words
package colphonetics

import "strings"

var (
	replacer = strings.NewReplacer(
		"Ä", "a",
		"Ö", "o",
		"Ü", "u",
		"ä", "a",
		"ö", "o",
		"ü", "u",
		"PH", "F",
		"Ph", "F",
		"pH", "f",
		"ph", "f",
		"Ç", "c",
		"É", "e",
		"È", "e",
		"Ê", "e",
		"À", "A",
		"Á", "a",
		"Â", "a",
		"Ë", "e",
		"Í", "i",
		"Ì", "i",
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
		65:  "0",  //A
		66:  "1",  //B
		67:  "4",  //C
		68:  "2",  //D
		69:  "0",  //E
		70:  "3",  //F
		71:  "4",  //G
		72:  "",   //H
		73:  "0",  //I
		74:  "0",  //J
		75:  "4",  //K
		76:  "5",  //L
		77:  "6",  //M
		78:  "6",  //N
		79:  "0",  //O
		80:  "1",  //P
		81:  "4",  //Q
		82:  "7",  //R
		83:  "8",  //S
		84:  "2",  //T
		85:  "0",  //U
		86:  "3",  //V
		87:  "3",  //W
		88:  "48", //X
		89:  "0",  //Y
		90:  "8",  //Z
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
		"SC": "8",
		"ZC": "8",
		"CX": "8",
		"KX": "8",
		"QX": "8",
		"sc": "8",
		"zc": "8",
		"cx": "8",
		"kx": "8",
		"qx": "8",
	}

	exceptionsLeading = map[string]string{
		"CA": "4",
		"CH": "4",
		"CK": "4",
		"CL": "4",
		"CO": "4",
		"CQ": "4",
		"CU": "4",
		"CX": "4",
		"ca": "4",
		"ch": "4",
		"ck": "4",
		"cl": "4",
		"co": "4",
		"cq": "4",
		"cu": "4",
		"cx": "4",
		"DC": "8",
		"DS": "8",
		"DZ": "8",
		"TC": "8",
		"TS": "8",
		"TZ": "8",
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
	s = replacer.Replace(s)
	l := len(s)
	r := make([]string, l, l)

	for i, char := range s {
		if i == 0 && len(s) > 1 && (char == 67 && s[i+1] == 83 || char == 98 && s[i+1] == 115) {
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
