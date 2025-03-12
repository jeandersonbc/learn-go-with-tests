package main

const (
	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
	frenchGreenting = "Bonjour, "
)

// A friendly hello
func Hello(s, lang string) string {
	name := validateName(s)
	prefix := pickPrefix(lang)
	return prefix + name + "!"
}

func validateName(s string) string {
	if len(s) == 0 {
		return "World"
	}
	return s
}

func pickPrefix(lang string) (prefix string) {
	prefix = englishGreeting
	switch lang {
	case "Spanish":
		prefix = spanishGreeting
	case "French":
		prefix = frenchGreenting
	}
	return
}
