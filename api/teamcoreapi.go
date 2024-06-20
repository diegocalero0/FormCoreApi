package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
)

// Method that allows to perform a GET request
// to the Teamcore API
func Get[T any]() (*T, error) {
	url := os.Getenv("TEAMCORE_API_URL")
	token := os.Getenv("TEAMCORE_API_TOKEN")
	request, _ := http.NewRequest("GET", url+"/questions", nil)
	request.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var data T
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, err
		}
		return &data, nil
	} else {
		return nil, errors.New("error proccesing " + strconv.Itoa(resp.StatusCode))
	}
}
