package runescape

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// var export = {getUsersOnline,}
func Hello() {
	fmt.Println("hello from runescape")
}

func GetUsersOnline() string { return getUsersOnline() }

func getUsersOnline() string {
	resp, err := http.Get("https://oldschool.runescape.com/")
	if err != nil {
		return "ERR"
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "ERR"
		}
		// fmt.Println(body)
		re := regexp.MustCompile("<p class='player-count'>There are currently ([0-9,]+) people playing!</p>")
		result := re.FindString(string(body))
		players := strings.Split(result, " ")[4]
		return players
	}
}