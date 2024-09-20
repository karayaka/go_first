package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*var colors []string

	colors = append(colors, "red")
	colors = append(colors, "blue")
	colors = append(colors, "green")
	colors = append(colors, "deneme")

	for i, valeu := range colors {
		fmt.Println(i, valeu)
	}

	humen := new(Human)

	humen.name = "mehmet"

	h2 := humen //bu bölüm refarens değer haline getiriyor
	h2.name = "deneme"

	fmt.Println(humen.name)
	fmt.Println(h2.name)
	slive := make([]int, 5)
	slive[0] = 111
	fmt.Println(len(slive))*/

	ch := make(chan string)
	go pr(ch)
	say(ch)
}

func say(ch chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		ch <- strings.Replace(text, "\n", "", -1)
	}
}
func pr(ch chan string) {
	for {
		name := <-ch
		println(len(name), name)
	}
}
