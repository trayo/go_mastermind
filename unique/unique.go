package unique

func RemoveDuplicateStrings(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, element := range elements {
		if encountered[element] != true {
			encountered[element] = true
			result = append(result, element)
		}
	}

	return result
}
