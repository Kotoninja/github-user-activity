package internal

import (
	"encoding/json"
	"net/http"
)

func serializer(res *http.Response) ([]activity, error) {
	var result []activity

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
