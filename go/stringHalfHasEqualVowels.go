package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var vol = map[int32]bool{97: true, 101: true, 105: true, 111: true, 117: true}

func totalVowels(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := totalVowels(s[1:])
	ch := rune(s[0])
	if _, e := vol[ch]; e {
		return res + 1
	}
	return res
}

func halvesAreAlike(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)

	i := 0
	mid := len(s)/2
	j := len(s)

	l := totalVowels(s[i:mid])
	r := totalVowels(s[mid:j])
	return l==r
}

/**
 *	Alternate approach
 */
func halvesAreAlikeAlt(s string) bool {
	couter := 0
	for i, v := range s {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < len(s)/2 {
				couter += 1
			} else {
				couter -= 1
			}
		}
	}
	return couter == 0
}


func main() {
	sc := bufio.NewReader(os.Stdin)
	text, _ := sc.ReadString('\n')
	fmt.Println(halvesAreAlike(text))
}
