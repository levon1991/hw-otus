package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	Name  string
	count int
}

func Top10(text string) []string {
	sliceList := make([]string, 0)

	// delete all symbols like '\n', '\t', double space
	space := regexp.MustCompile(`\s+`)
	res := space.ReplaceAllString(text, " ")

	// check for empty string
	if len(res) == 0 || res == " " {
		return sliceList
	}

	// use map just for calculate words count
	split := strings.Split(res, " ")
	listMap := make(map[string]int)
	for _, v := range split {
		listMap[v]++
	}

	// put map in struct for sorting
	w := make([]Word, 0)
	for key, v := range listMap {
		w = append(w, Word{Name: key, count: v})
	}

	// sort by Name and count
	sort.Slice(w, func(i, j int) bool {
		if w[i].count == w[j].count {
			return w[i].Name < w[j].Name
		}
		return w[i].count > w[j].count
	})

	// get first 10 words
	for i := 0; i < 10 && i < len(w); i++ {
		sliceList = append(sliceList, w[i].Name)
	}

	return sliceList
}
