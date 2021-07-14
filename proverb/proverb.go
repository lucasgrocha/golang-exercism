// Provides proverb stuff
package proverb

import "fmt"

// Proverb returns the proverb
func Proverb(rhyme []string) []string {
	rhymes := []string{}
	rhymeLength := len(rhyme)

	for index, word := range rhyme {
		var phrase string
		index++

		if isLastIndex(index, rhymeLength) {
			phrase = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			phrase = fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[index])
		}

		rhymes = append(rhymes, phrase)
	}

	return rhymes
}

func isLastIndex(index int, length int) bool {
	return index == length
}
