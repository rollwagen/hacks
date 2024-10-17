package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Check if a URL is provided as the first command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Error: URL argument is missing")
		os.Exit(1)
	}

	url := os.Args[1]

	apiKey := os.Getenv("JINA_API_KEY")
	if apiKey == "" || !strings.HasPrefix(apiKey, "jina_") {
		fmt.Println("Error: JINA_API_KEY env var is not set or invalid")
		os.Exit(1)
	}

	proxyURL := "https://r.jina.ai/" + url
	fmt.Println(proxyURL)
	req, err := http.NewRequest("GET", proxyURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
