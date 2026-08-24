package internal

import (
	"errors"
	"fmt"
	"net/http"
)

func FetchActivity(username string) ([]string, error) {
	res, err := http.Get(fmt.Sprintf(gitHubUrl, username))
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []string{}, errors.New("User not found")
	}

	data, err := serializer(res)
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	result := createOutput(data)

	return result, nil
}

func createOutput(data []activity) []string {
	result := []string{}

	hashMap := map[string]int{}
	for _, act := range data {
		hashMap[fmt.Sprintf("%s - %s", act.Type, act.Repo.Name)] ++
		// result = append(result, fmt.Sprintf("%s - %s\n", act.Type, act.Repo.Name))
	}

	fmt.Println(hashMap)

	return result
}
