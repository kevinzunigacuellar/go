package helloworld

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	const (
		englishHelloPrefix = "Hello, "
		spanishHelloPrefix = "Hola, "
		frenchHelloPrefix  = "Bonjour, "
	)

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
