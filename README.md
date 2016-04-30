# colphonetics

Cologne phonetics (Kölner Phonetik) is an algorithm related to soundex but optimized for the German language. It calculates a phonetic code for a given sequence of words.

Master: [![Build Status](https://travis-ci.org/florianorben/colphonetics.svg?branch=master)](https://travis-ci.org/florianorben/colphonetics/)

## Installation

    $ go get github.com/florianorben/colphonetics
    
## Docs

https://godoc.org/github.com/florianorben/colphonetics
    
## Usage

```go
import (
	"fmt"

	"github.com/florianorben/colphonetics"
)

func main() {
	fmt.Println(colphonetics.Code("Müller-Lüdenscheidt"))
    // Output: 65752682	
}
```