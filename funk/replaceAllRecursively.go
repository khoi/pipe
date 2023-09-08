package funk

import "strings"

// ReplaceAllRecursively replaces all copies of a string with the substring and
// keeps going until ths string's length is constant (e.g. if the first iteration
// creates the desired substring for the (n+1)th iteration to remove)
func ReplaceAllRecursively(text string, toReplace string, replaceWith string) string {
	var currentText string = text
	strippedText := strings.ReplaceAll(text, toReplace, replaceWith)
	for len(strippedText) < len(currentText) {
		currentText = strippedText
		strippedText = strings.ReplaceAll(strippedText, toReplace, replaceWith)
	}
	return strippedText
}
