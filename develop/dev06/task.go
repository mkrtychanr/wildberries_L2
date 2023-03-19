package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pborman/getopt"
)

func main() {
	f := getopt.IntLong("fields", 'f', 0, "выбрать поля (колонки)")
	d := getopt.StringLong("delimiter", 'd', "\t", "использовать другой разделитель")
	s := getopt.BoolLong("separated", 's', "использовать другой разделитель")

	getopt.Parse()
	if *f <= 0 {
		log.Fatal("f <= 0")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		splitTxt := strings.Split(txt, *d)
		if *s && !strings.Contains(txt, *d) {
			fmt.Println("")
		} else if len(splitTxt) < *f {
			fmt.Println(txt)
		} else {
			fmt.Println(splitTxt[*f-1])
		}
	}
}
