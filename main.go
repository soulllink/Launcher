package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type applist struct {
	applist apps `json:"applist"`
}
type apps struct {
	apps []app_item `json:"apps"`
}
type app_item struct {
	appid int64  `json:"appid"`
	name  string `json:"name"`
}

var games items

const steamapireq = "http://api.steampowered.com/ISteamApps/GetAppList/v0002/?key=STEAMKEY&format=json"

func main() {
	requestapi()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		chatevents(text, games)
	}
}

func startgame(path string) {
	c := exec.Command("cmd", "/C", path)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Runnig")

}

func chatevents(s string, g items) {
	//var state string
	if g.find(s) != -1 {
		GamePath := g.items[g.find(s)].Path
		startgame(GamePath)
	}
	if strings.Contains(s, "addgame") {
		addgame(s)
	}
	if strings.Contains(s, "delgame") {
		delgame(s)
	}
	if strings.Contains(s, "save") {
		games.save()
	}

}

func addgame(s string) {
	s = strings.Replace(s, "addgame ", "", -1)
	split := strings.Split(s, ": ")
	path := strings.Replace(split[1], "\\", "/", -1)
	temp := JSON{
		Name:      split[0],
		AliasName: "",
		Path:      path,
		Args:      []string{},
		TimesRun:  0,
	}
	games.additem(temp)
}

func delgame(s string) {
	s = strings.Replace(s, "delgame ", "", -1)
	games.delete(s)
}

func requestapi() {

	req, err := http.Get(steamapireq)
	if err != nil {

	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	fmt.Println("Request: ", string(body))
}
