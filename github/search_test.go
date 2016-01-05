package github

import (
	"testing"
)

func TestSearhc(t *testing.T) {
	result, err := Search("http2")
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range result.Items {
		t.Logf("%s (%d) - %s", item.PackageName(), item.StargazersCount, item.Description)
	}
}
