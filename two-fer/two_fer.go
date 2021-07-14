package twofer

// ShareWith shares something
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	phrase := "One for " + name + ", one for me."

	return phrase
}
