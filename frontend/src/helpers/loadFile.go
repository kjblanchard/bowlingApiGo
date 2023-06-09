package helpers

import "os"

func LoadTemplateFromFile(title string) (string, error) {
	filename := title + ".html"
	body, err := os.ReadFile("./templates/" + filename)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
