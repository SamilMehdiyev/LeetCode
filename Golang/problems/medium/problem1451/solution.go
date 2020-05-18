package solutions

import (
	"sort"
	"strings"
)

// ArrangeWords function is a solution for the
// Problem #1451 - Rearrange Words in a Sentence
// from leetcode.com
func ArrangeWords(text string) string {
	if len(text) == 0 {
		return ""
	}

	strs := strings.Split(text, " ")
	strsMap := make(map[int]string)
	shortLen := len(strs[0])

	for i := 0; i < len(strs); i++ {
		strsMap[len(strs[i])] += strings.ToLower(strs[i]) + " "
		if len(strs[i]) < shortLen {
			shortLen = len(strs[i])
		}
	}

	var keys []int
	for k := range strsMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	firstStrs := strings.Split(strsMap[shortLen], " ")
	isFirstCapitalized := false
	for i := 0; i < len(firstStrs)-1; i++ {
		if !isFirstCapitalized {
			strsMap[shortLen] = strings.Title(firstStrs[i]) + " "
			isFirstCapitalized = true
		} else {
			strsMap[shortLen] += firstStrs[i] + " "
		}
	}

	newStr := ""
	for _, k := range keys {
		newStr += strsMap[k]
	}

	newStr = strings.Trim(newStr, " ")

	return newStr
}
