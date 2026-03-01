package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Streak int `json:"streak"`
	Name string `json:"name"`
}
type DuolingoResponse struct {
	Users []User `json:"users"`
}

func getStreak(username string)(int, error){
	// Grabbing streak info from Duolingo API
	url := fmt.Sprintf("https://www.duolingo.com/2017-06-30/users?username=%s", username)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil{
		return 0, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200{
		return 0, fmt.Errorf("returned status: %d", resp.StatusCode)
	}

	var data DuolingoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil{
		return 0, err
	}

	if len(data.Users) == 0 {
		return 0, fmt.Errorf("user not found")
	}

	fmt.Println(data.Users[0].Streak)
	return data.Users[0].Streak, nil
}
