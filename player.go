package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// var raw_url string = "https://www.youtube.com/watch?v=5qap5aO4i9A"
var raw_url string = "https://www.youtube.com/watch?v=1ZuldMGkJNY"

func main() {
	// fmt.Println(len(os.Args), os.Args[1])
	fmt.Println("Start player")
	var ready_url string = getYoutubeUrl(raw_url)
	playMusic(ready_url)
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

func playMusic(url string) {
	// delete last character from string
	tmp := strings.Split(url, "")
	tmp = tmp[:len(tmp)-1] 
	url = strings.Join(tmp, "")

	// play using ffplay
	fmt.Println("executing the command 'ffplay -nodisp ", url,"'")
	err := exec.Command("ffplay", "-autoexit", "-nodisp", "-hide_banner", "-loglevel", "error", url).Run()
	if err != nil {
		log.Println(err)
	}
}