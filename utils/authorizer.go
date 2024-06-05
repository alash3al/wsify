package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func ShouldAcceptPayload(url string, payload any) (bool, error) {
	if url == "" {
		return true, nil
	}

	j, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(j))
	if err != nil {
		return false, err
	}

	defer (func() {
		_ = resp.Body.Close()
	})()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	if resp.StatusCode >= 500 {
		return false, fmt.Errorf("unexpected statusCode (%d) returned from the authorizer", resp.StatusCode)
	}

	return false, nil
}
