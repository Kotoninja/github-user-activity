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

	var activityCounter int = 1
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			activityCounter++
			continue
		}

		repoName := data[i].Repo.Name
		ref := data[i].Payload.Ref
		refType := data[i].Payload.RefType

		switch data[i].Type {
		case "PushEvent":
			result = append(result, fmt.Sprintf(getEvent(data[i].Type), activityCounter, repoName))
		case "CreateEvent":
			result = append(result, fmt.Sprintf(getEvent(data[i].Type), refType, ref, repoName))
		default:
			result = append(result, fmt.Sprintf(getEvent(data[i].Type), repoName))
		}
		activityCounter = 1
	}

	//TODO handle last item
	return result
}

func getEvent(eventName string) string {
	switch eventName {
	case "WatchEvent":
		return "⭐ Starred %s"
	case "PushEvent":
		return "Pushed %d commit(s) to %s"
	case "PullRequestEvent":
		return "Opened a pull request in %s"
	case "CreateEvent":
		return "Created %s %s in %s"
	default:
		return "Not found"
	}
}
