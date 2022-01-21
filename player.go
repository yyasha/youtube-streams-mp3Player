package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// var raw_url string = "https://www.youtube.com/watch?v=5qap5aO4i9A"
// var raw_url string = "https://www.youtube.com/watch?v=xgSKo9xx0E4"

var raw_url string

func main() {
	useDisplay := flag.Bool("display", false, "show display")
	flag.StringVar(&raw_url, "url", "https://www.youtube.com/watch?v=xgSKo9xx0E4", "Link to the video on YouTube")
	flag.Parse()

	fmt.Println("raw_url:", raw_url)
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

	fmt.Println("executing the command 'ffplay -autoexit -nodisp -hide_banner -loglevel error", url,"'")

	// play using ffplay

	args := []string{"-autoexit", "-hide_banner", "-loglevel", "error", url}
	err := exec.Command("ffplay", args...).Run()
	if err != nil {
		log.Println(err)
	}
	// if *useDisplay {
	// 	err := exec.Command("ffplay", "-autoexit", "-hide_banner", "-loglevel", "error", url).Run()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// } else {
	// 	err := exec.Command("ffplay", "-autoexit", "-nodisp", "-hide_banner", "-loglevel", "error", url).Run() //
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }
}