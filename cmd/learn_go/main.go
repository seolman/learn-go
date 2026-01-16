package main

func reformat(message string, formatter func(string) string) string {
	for range 3 {
		message = formatter(message)
	}
	const prefix = "TEXTIO: "
	return prefix + message
}
