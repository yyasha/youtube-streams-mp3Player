package main

import (
	"fmt"
	"log"
	"os/exec"
)

var raw_url string = "https://www.youtube.com/watch?v=5qap5aO4i9A"

func main() {
	// fmt.Println(len(os.Args), os.Args[1])
	fmt.Println("Start player")
	var ready_url string = getYoutubeUrl(raw_url)
	playMusic(ready_url)
}

func getYoutubeUrl(raw_url string) string {
	// youtube-dl -f best --get-url url
	fmt.Println("executing the command 'youtube-dl -f best --get-url ", raw_url,"'")
	cmd, err := exec.Command("youtube-dl", "-f", "best", "--get-url", raw_url).Output()
	if err != nil {
		log.Println(err)
	}
	var out_url string = fmt.Sprintf("%s", cmd)
	return out_url
}

func playMusic(url string) {
	// ffplay -nodisp url
	fmt.Println("executing the command 'ffplay -nodisp ", url,"'")
	err := exec.Command("ffplay", "-nodisp", url).Run()
	if err != nil {
		log.Println(err)
	}
}