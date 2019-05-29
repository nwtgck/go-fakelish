package fakelish

import (
	"math/rand"
	"strings"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateFakeWordWithUnexpectedLength() string {
	maxSeq := 2 // TODO: Hard code
	ch := "^"
	fakeWord := ""
	var chrs []string
	for ch != "END" {
		chrs = append(chrs, ch)
		if len(chrs) > maxSeq {
			chrs = chrs[1:]
		}
		var nextAccumedProbs []AccumedProb = nil
		n := 0
		for {
			str := strings.Join(chrs[n:], "")
			nextAccumedProbs = WordProbability[str]
			n += 1
			if !(nextAccumedProbs == nil && n < len(chrs)) {
				break
			}
		}
		nextCh := ""
		r := random.Float32()
		for _, x := range nextAccumedProbs {
			candidateNextCh := x.Ch
			prob := x.Prob
			if r <= prob {
				nextCh = candidateNextCh
				break
			}
		}
		if nextCh != "END" {
			fakeWord += nextCh
		}
		ch = nextCh
	}
	return fakeWord
}

func GenerateFakeWordByLength(length int) string {
	fakeWord := ""
	for len(fakeWord) != length {
		fakeWord = GenerateFakeWordWithUnexpectedLength()
	}
	return fakeWord
}

func GenerateFakeWord(minLength int, maxLength int) string {
	fakeWord := ""
	for !(minLength <= len(fakeWord) && len(fakeWord) <= maxLength) {
		fakeWord = GenerateFakeWordWithUnexpectedLength()
	}
	return fakeWord
}
