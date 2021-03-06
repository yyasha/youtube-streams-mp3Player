package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var raw_url string
var volume int

func main() {
	useDisplay := flag.Bool("display", false, "Show display")
	flag.StringVar(&raw_url, "url", "https://www.youtube.com/watch?v=5qap5aO4i9A", "Link to the video on YouTube")
	flag.IntVar(&volume, "volume", 100, "Playback volume")
	flag.Parse()

	fmt.Println("Start player")
	var ready_url string = getYoutubeUrl(raw_url)
	playMusic(ready_url, useDisplay)
}

func getYoutubeUrl(raw_url string) string {
	// get clear url using youtube-dl
	fmt.Println("executing the command 'youtube-dl -f best --get-url ", raw_url,"'")
	cmd, err := exec.Command("youtube-dl", "-f", "best", "--get-url", raw_url).Output()
	if err != nil {
		log.Println(err)
	}
	var out_url string = fmt.Sprintf("%s", cmd)
	return out_url
}

func playMusic(url string, useDisplay *bool) {
	// delete last character from string
	tmp := strings.Split(url, "")
	tmp = tmp[:len(tmp)-1] 
	url = strings.Join(tmp, "")

	// play using ffplay
	args := []string{"-autoexit", "-hide_banner", "-loglevel", "error"}

	// setting args

	// display
	if !*useDisplay {
		args = append(args, "-nodisp")
	}

	// volume
	args = append(args, "-volume")
	args = append(args, strconv.Itoa(volume))

	// url
	args = append(args, url)

	err := exec.Command("ffplay", args...).Run()
	if err != nil {
		log.Println(err)
	}
}