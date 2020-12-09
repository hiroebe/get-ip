package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)

func extract(ip []byte) (string, bool) {
	if ip == nil || string(ip) == "127.0.0.1" {
		return "", false
	}
	splitIP := bytes.SplitN(ip, []byte("."), 4)
	for _, sp := range splitIP {
		i, err := strconv.ParseInt(string(sp), 10, 32)
		if err != nil {
			return "", false
		}
		if i < 0 || i > 255 {
			return "", false
		}
	}
	return string(ip), true
}

func main() {
	out, err := exec.Command("ifconfig").Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(out, []byte("\n"))

	for _, l := range lines {
		if !bytes.Contains(l, []byte("inet")) {
			continue
		}
		ip, ok := extract(re.Find(l))
		if ok {
			fmt.Print(ip)
			break
		}
	}
}
