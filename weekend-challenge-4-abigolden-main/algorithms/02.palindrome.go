// A palindrome is a word, phrase, or number that is spelled the same forward and backward. For example, “dad” is a palindrome; “A man, a plan, a canal: Panama” is a palindrome if you take out the spaces and ignore the punctuation; and 1,001 is a numeric palindrome. We can use a stack to determine whether or not a given string is a palindrome.
// Write a function that takes a string of letters and returns true or false to determine whether it is palindromic.
// Input: amazon
// Output: false
//
// Input: dad
// Output: true

package algorithms

import "strings"

func IsPalindrome(text string) bool {
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	cleanedText := strings.ReplaceAll(strings.ToLower(text), " ", "")

	reversedText := reverse(cleanedText)

	return cleanedText == reversedText
}
