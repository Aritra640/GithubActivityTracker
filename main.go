package main

import (
	"flag"
	"fmt"
)

func main() {

	username := flag.String("u" , "" , "<username> github username")
  flag.Parse()

  if *username == "" {
    fmt.Println("username found empty , please enter a valid username in the format :\ngithub-activity <username>")
    return 
  }

  getData,err := getGithubActivity(*username)
  if err != nil {
    fmt.Println("username not found , please enter a valid username")
    return 
  }

  fmt.Println("Output :")
  for _,line := range getData {
    fmt.Print(" - ")
    fmt.Println(line)
  }

  fmt.Println("\n\nThank you")
}
