package main

import (
	"strings"
	"testing"
)

func TestAnagram(t *testing.T) {
	got := groupAnagrams([]string{"тяпка", "пятак", "пятка", "клоун", "локун"})
	want := make(map[string][]string)
	want["тяпка"] = []string{"пятак", "пятка", "тяпка"}
	want["клоун"] = []string{"клоун", "локун"}

	if len(want) != len(got) {
		t.Errorf("invalid size")
	}

	for key, valueInGot := range want {
		valueInWant, exist := got[key]
		if !exist {
			t.Errorf("the key %s is missing in got", key)
		}
		if len(valueInGot) != len(valueInWant) {
			t.Errorf("the size of slice in got is %d, and in want is %d", len(valueInGot), len(valueInWant))
		}
		for i := 0; i < len(valueInWant); i++ {
			if strings.Compare(valueInGot[i], valueInWant[i]) != 0 {
				t.Errorf("the strings are different")
			}
		}
	}
}

func TestOneStringCase(t *testing.T) {
	got := groupAnagrams([]string{"привет"})
	if len(got) != 0 {
		t.Errorf("wrong size. one word collection")
	}
}
