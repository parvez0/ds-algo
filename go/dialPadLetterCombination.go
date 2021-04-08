package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 *  Problem statement
 *  Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
 *
 *  Input: digits = "23"
 *  Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 */

var letters = map[string]string {
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := letterCombinations(digits[1:])
	chars := letters[string(digits[0])]
	if len(res) == 0 {
		for _, ch := range chars {
			res = append(res, string(ch))
		}
		return res
	} else {
		var temp []string
		for _, ch := range chars {
			for _, v := range res {
				temp = append(temp, string(ch) + v)
			}
		}
		return temp
	}
}

func main() {
	sc := bufio.NewReader(os.Stdin)
	text, _ := sc.ReadString('\n')
	fmt.Println(letterCombinations(text))
}
