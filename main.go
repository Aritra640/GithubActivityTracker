package main

import (
	"flag"
	"fmt"
)

func main() {

  var username string 
  flag.StringVar(&username , "username" , "" , "github username")

  flag.Parse()

  args := flag.Args()
  if len(args) > 0 {
    username = args[0]
  }

  if username == "" {
    fmt.Println("username found empty , please enter a valid username in the format :\ngithub-activity <username>")
    return 
  }

  getData,err := getGithubActivity(username)
  if err != nil {
    fmt.Println("username not found , please enter a valid username")
    return 
  }

  if len(getData) == 0 {
    fmt.Println("no activity found")
  }

  fmt.Println()

  for _,events := range getData {
    var event string 
    switch events.Type {
    case "PushEvent":
      count := len(events.Payload.Commits)
      event = fmt.Sprintf("Pushed %d commits to %s" , count , events.Repo.Name)
    case "IssueEvent":
      event = fmt.Sprintf("%s an issue in %s" , events.Payload.Action , events.Repo.Name)
    case "WatchEvent":
      event = fmt.Sprintf("Starred %s" , events.Repo.Name)
    case "ForkEvent":
      event = fmt.Sprintf("Forked %s" , events.Repo.Name)
    case "CreateEvent":
      event = fmt.Sprintf("Created %s in %s" , events.Payload.RefType , events.Repo.Name)
    default:
      event = fmt.Sprintf("%s in %s" , events.Type , events.Repo.Name)
    }

    fmt.Println("# " , event)
  }
  
  fmt.Println("\n\nThank you")
}
