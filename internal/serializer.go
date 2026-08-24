package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func serializer(res *http.Response) ([]activity, error) {
	var result []activity

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []activity{}, err
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return []activity{}, err
	}

	return result, nil
}
