package main

import (
	"fmt"
	"io"
	"net/http"
)

func getGithubActivity(username string) ([]string, error) {

	APIstring := fmt.Sprintf("https://api.github.com/users/%s/events" , username)
	res, err := http.Get(APIstring)

  if err != nil {
    return []string{} , err
  }

  resData,err := io.ReadAll(res.Body)
  if err != nil {
    return []string{} , err
  }

  return []string{string(resData)},nil

}
