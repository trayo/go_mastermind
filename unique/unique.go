package unique

func RemoveDuplicateLetters(elements string) string {
	encountered := map[rune]bool{}
	result := []rune{}

	for _, element := range elements {
		if encountered[element] != true {
			encountered[element] = true
			result = append(result, element)
		}
	}

	return string(result)
}
