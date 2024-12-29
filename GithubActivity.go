package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubActivity struct {
  Type string `json:"type"`
  Repo Repo   `json:"repo"`
  CreatedAt string `json:"created_at"`
  Payload struct {
    Action string `json:"action"`
    Ref    string `json:"ref"`
    RefType string `json:"ref_type"`
    Commits []struct {
      Message string `json:"message"`
    } `json:"commits"`
  }
}

type Repo struct {
  Name string `json:"name"`
}

func getGithubActivity(username string) ([]GithubActivity, error) {

	APIstring := fmt.Sprintf("https://api.github.com/users/%s/events" , username)
	res, err := http.Get(APIstring)

  if err != nil {
    return nil , err
  }

  var activities []GithubActivity
  if err := json.NewDecoder(res.Body).Decode(&activities); err != nil {
    return nil , err
  }

  return activities , nil
}

