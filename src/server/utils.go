package server

import (
	"io/ioutil"
	"log"
	"strings"
)

//LoadWordsList Loads a words list from a newline separated text file
func LoadWordsList(fp string) []string {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Println("[SYSTEM] - Couldn't open words list.")
		panic(err)
	}
	return strings.Split(strings.Replace(string(data), "\r\n", "\n", -1), "\n")
}
