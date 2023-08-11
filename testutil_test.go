package main

import (
	"errors"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func repeat(b string, n int) string {
	buf := make([]byte, 0, n*len(b))
	for i := 0; i < n; i++ {
		buf = append(buf, b...)
	}
	return string(buf)
}

func errorIs(expect error) func(t *testing.T, err error) {
	return func(t *testing.T, err error) {
		require.Error(t, err)
		require.True(t, errors.Is(err, expect),
			"actual:   %#v;\nexpected: %#v", err, expect)
	}
}

func MustReadFile(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}

func keysSorted[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
