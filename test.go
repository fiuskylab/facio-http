package facio

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestType int

const (
	Equal TestType = iota
	Nil
	Regex
)

type testCases struct {
	name      string
	want      interface{}
	got       interface{}
	regexRule *regexp.Regexp
	testType  TestType
}

func runTests(t *testing.T, tts []testCases) {
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			switch tt.testType {
			case Equal:
				assert.Equal(t, tt.want, tt.got)
			case Regex:
				assert.Regexp(t, tt.regexRule, tt.got)
			case Nil:
				assert.Nil(t, tt.got)
			default:
			}
		})
	}
}
