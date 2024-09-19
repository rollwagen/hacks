package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ytt",
	Short: "Download YouTube video transcripts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("missing required parameter YouTube URL")
			return
		}
		youTubeURL := args[0]
		if !isValidURL(youTubeURL) {
			fmt.Println("given paramater is not a valid URL")
			return
		}

		id, _ := extractVideoID(youTubeURL)
		transcript, err := getTranscript(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		text, _ := extractText(transcript)

		fmt.Println(text)
	},
}

// isValidURL checks if the provided string is a valid URL
func isValidURL(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// extractVideoID extracts the video id from a YouTube url such as "https://www.youtube.com/watch?v=oD-d9B71yLo"
func extractVideoID(url string) (string, error) {
	var err error
	var ret string
	pattern := `(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]{11})`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(url)
	if len(match) > 1 {
		ret = match[1]
	} else {
		err = fmt.Errorf("invalid YouTube URL, can't get video ID")
	}
	return ret, err
}

func getTranscript(videoID string) (string, error) {
	url := "https://www.youtube.com/watch?v=" + videoID
	response, err := soup.Get(url)
	if err != nil {
		return "", err
	}

	doc := soup.HTMLParse(response)
	scriptTags := doc.FindAll("script")
	for _, scriptTag := range scriptTags {
		if strings.Contains(scriptTag.Text(), "captionTracks") {
			regex := regexp.MustCompile(`"captionTracks":(\[.*?\])`)
			match := regex.FindStringSubmatch(scriptTag.Text())
			if len(match) > 1 {
				var captionTracks []struct {
					BaseURL string `json:"baseUrl"`
				}

				err = json.Unmarshal([]byte(match[1]), &captionTracks)
				if err != nil {
					return "", err
				}

				if len(captionTracks) > 0 {
					transcriptURL := captionTracks[0].BaseURL
					ret, err := soup.Get(transcriptURL)
					if err != nil {
						return "", err
					}
					return ret, nil
				}
			}
		}
	}
	err = fmt.Errorf("could not find transcript")
	return "", err
}
