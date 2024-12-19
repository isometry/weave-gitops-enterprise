package steps

import (
	"fmt"
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
)

// getWgeVersions gets the latest 3 available WGE versions from the helm repository
func getWgeVersions() ([]string, error) {
	chartUrl := fmt.Sprintf("%s/index.yaml", wgeChartUrl)
	versions, err := fetchHelmChartVersions(chartUrl)
	if err != nil {
		return []string{}, err
	}
	return versions, nil
}

// fetchHelmChartVersions helper method to fetch wge helm chart versions.
func fetchHelmChartVersions(chartUrl string) ([]string, error) {
	bodyBytes, err := doGetRequest(chartUrl)
	if err != nil {
		return []string{}, err
	}

	var chart helmChartResponse
	err = yaml.Unmarshal(bodyBytes, &chart)
	if err != nil {
		return []string{}, err
	}
	entries := chart.Entries[wgeChartName]
	var versions []string
	for _, entry := range entries {
		if entry.Name == wgeChartName {
			versions = append(versions, entry.Version)
			if len(versions) == 3 {
				break
			}
		}
	}

	return versions, nil
}

func doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err

	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyBytes, err
}
