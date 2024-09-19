package main

import (
	"encoding/xml"
	"html"
	"strings"
)

type Transcript struct {
	Texts []Text `xml:"text"`
}

type Text struct {
	Content string `xml:",chardata"`
}

func extractText(input string) (string, error) {
	var t Transcript
	err := xml.Unmarshal([]byte(input), &t)
	if err != nil {
		return "", err
	}

	texts := []string{}
	for _, text := range t.Texts {
		content := html.UnescapeString(text.Content)
		texts = append(texts, content)
	}
	return strings.Join(texts, " "), nil
}
