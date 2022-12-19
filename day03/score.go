package day03

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func scoreForCharacter(char string) int {
	for i, c := range alphabet {
		if string(c) == char {
			return i + 1
		}
	}
	return 0
}
