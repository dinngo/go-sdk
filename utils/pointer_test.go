package utils

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	testCases := []any{
		1, uint64(1), "1", time.Duration(1),
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, reflect.Ptr, reflect.ValueOf(Pointer(testCase)).Kind())
		})
	}
}
