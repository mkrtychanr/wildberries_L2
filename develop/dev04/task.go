package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(data []string) map[string][]string {
	sortedData := make([]string, 0, len(data))
	for _, value := range data {
		runs := []rune(value)
		sort.Slice(runs, func(i, j int) bool {
			return runs[i] < runs[j]
		})
		sortedData = append(sortedData, string(runs))
	}
	setOfStrings := make(map[string][]string)
	for index, value := range sortedData {
		setOfStrings[value] = append(setOfStrings[value], data[index])
	}
	resultWithMaps := make(map[string]map[string]struct{})
	for _, value := range setOfStrings {
		if len(value) == 1 {
			continue
		}
		resultWithMaps[value[0]] = make(map[string]struct{})
		for _, str := range value {
			resultWithMaps[value[0]][str] = struct{}{}
		}
	}
	resultWithSlices := make(map[string][]string)
	for key, value := range resultWithMaps {
		for str := range value {
			resultWithSlices[key] = append(resultWithSlices[key], str)
		}
	}
	result := make(map[string][]string)
	for key, value := range resultWithSlices {
		sort.Slice(value, func(i, j int) bool {
			return strings.Compare(value[i], value[j]) == -1
		})
		result[key] = value
	}

	return result
}

func main() {
	res := groupAnagrams([]string{"тяпка", "пятак", "пятка", "клоун", "локун"})
	for key, value := range res {
		fmt.Println(key, value)
	}
}
