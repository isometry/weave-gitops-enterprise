package steps

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert"
)

func TestFetchHelmChart(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, `entries:
  mccp:
  - version: 1.0.0
    name: mccp
  - version: 1.1.0
    name: mccp
  - version: 1.2.0
    name: mccp`)
	}))
	defer mockServer.Close()

	versions, err := fetchHelmChartVersions(mockServer.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedVersions := []string{"1.0.0", "1.1.0", "1.2.0"}
	assert.Equal(t, expectedVersions, versions, "versions are not equal")
}
