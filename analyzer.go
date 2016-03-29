package bullyanalyzer

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Analyzer struct {
	Entries []string
}

type AnalyzerResult struct {
	Post  string
	Value float32
}

// New creates a new profanity analyzer from an input file
func New(path string) (*Analyzer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var a Analyzer

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a.Entries = append(a.Entries, scanner.Text())
	}

	sort.Strings(a.Entries)

	return &a, scanner.Err()
}

// ContainsEntry check is an entry is in the profanity lexicon list
// this function assumes that the list is sorted
func (a *Analyzer) ContainsEntry(entry string) bool {
	index := sort.SearchStrings(a.Entries, entry)
	if index == len(a.Entries) {
		return false
	}
	return a.Entries[index] == entry
}

// AnalyzePost assigns a fuzzy value represnting the "offensifness of a post"
func (a *Analyzer) AnalyzePost(post string) (r AnalyzerResult) {

	r.Post = post
	r.Value = 0

	splitPost := strings.Split(post, " ")
	for _, word := range splitPost {
		if a.ContainsEntry(word) {
			r.Value += float32(1) / float32(len(splitPost))
		}
	}

	return

}
