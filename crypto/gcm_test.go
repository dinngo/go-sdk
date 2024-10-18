package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextEncrypter(t *testing.T) {
	testCases := []struct {
		plaintext string
		key       string
		expected  string
	}{
		{
			plaintext: "c451fb09d44031506f2c55cc3132026a45f47bd85e8b54e8ba2edf38ced3dc56",
			key:       "165eb677f24ba123746082d33f376575",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			ciphertext, err := TextEncrypter([]byte(testCase.plaintext), []byte(testCase.key))
			assert.Nil(t, err)
			_, err = base64.StdEncoding.DecodeString(string(ciphertext))
			assert.Nil(t, err)
		})
	}
}

func TestTextDecrypter(t *testing.T) {
	testCases := []struct {
		ciphertext string
		key        string
		expected   string
	}{
		{
			ciphertext: "96UkQUfT1aelNxzO3FMlUAnFk8X6XIyiQDLNEU0C9bxaGkpJPj5u37Fxlf9pvfjb2bRvWzT2oXw/9oFGL0rh+c/q+8k+qBZI9trr2t+/JGSPR1HPVsyfiOsssLI=",
			key:        "165eb677f24ba123746082d33f376575",
			expected:   "c451fb09d44031506f2c55cc3132026a45f47bd85e8b54e8ba2edf38ced3dc56",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			plaintext, err := TextDecrypter([]byte(testCase.ciphertext), []byte(testCase.key))
			assert.Nil(t, err)
			assert.Equal(t, testCase.expected, string(plaintext))
		})
	}
}
