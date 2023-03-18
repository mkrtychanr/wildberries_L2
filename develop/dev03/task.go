package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getFile(path string, uniqueRequred bool) ([]string, error) {
	file, err := os.Open(path)
	set := make(map[string]struct{})
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		if uniqueRequred {
			if _, ok := set[scanner.Text()]; !ok {
				set[scanner.Text()] = struct{}{}
			} else {
				continue
			}
		}
		data = append(data, scanner.Text())
	}
	return data, nil
}

func toSort(data []string, n bool) []string {
	if n {
		sort.Slice(data, func(i, j int) bool {
			vi, _ := strconv.Atoi(data[i])
			vj, _ := strconv.Atoi(data[j])
			return vi < vj
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	}
	return data
}

func main() {
	n := *flag.Bool("n", false, "сортировка по числовому значению")
	r := *flag.Bool("r", false, "сортировка в обратном порядке")
	u := *flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	filename := flag.Arg(0)
	file, err := getFile(filename, u)
	if err != nil {
		panic(err)
	}

	file = toSort(file, n)

	if r {
		for i := len(file) - 1; i > 0; i-- {
			fmt.Println(file[i])
		}
	} else {
		for _, value := range file {
			fmt.Println(value)
		}
	}
}
