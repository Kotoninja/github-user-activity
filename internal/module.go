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

	// keys := make([]activity, 0, len(data))

	// hashMap := map[activity]int{}
	for _, act := range data {
		result = append(result, fmt.Sprintf(getEvent(act.Type), act.Repo.Name))
	}

	// slices.SortFunc(keys, func(a, b activity) int {
	// 	return b.CreatedAt.Compare(a.CreatedAt)
	// })

	// fmt.Println(keys)
	return result
}

func getEvent(eventName string) string {
	switch eventName {
	case "WatchEvent":
		return "⭐ Starred %s"
	case "PushEvent":
		return "Pushed commit to %s"
	case "PullRequestEvent":
		return "Opened a pull request in %s"
	case "CreateEvent":
		return "Created a %s repository."
	default:
		return "Not found"
	}
}
