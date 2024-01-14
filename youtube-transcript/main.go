package main

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"regexp"
	"strings"
)

func main() {
	c := youtube.Client{}

	//url := "https://www.youtube.com/watch?v=M16BnyhSTsw"
	url := "https://www.youtube.com/watch?v=UTRBVPvzt9w"

	v, err := c.GetVideo(url)
	if err != nil {
		panic(err)
	}

	t, err := c.GetTranscript(v)
	if err != nil {
		panic(err)
	}

	// e.g. '13:30 - so much for watching make sure to like'.
	reg := regexp.MustCompile("^(\\d+:\\d+ - )")
	for _, line := range strings.Split(t.String(), "\n") {
		sp := reg.Split(line, -1)
		if len(sp) > 1 {
			//fmt.Printf("%s ---- %s\n", sp[1], line)
			fmt.Println(sp[1])
		}
	}

}
