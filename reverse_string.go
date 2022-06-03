package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here

	for _, s := range input {
		output = string(s) + output
	}

	return output
}
