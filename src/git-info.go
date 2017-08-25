package main

import (
  "./helpers"
  "fmt"
  "path/filepath"
  "github.com/libgit2/git2go"
)

func main() {
  dir, _ := filepath.Abs(filepath.Dir("."))
  repo, _ := git.OpenRepositoryExtended(dir, 0, "0")
  remotes, _ := repo.Remotes.List()
  default_remote, _ := repo.Remotes.Lookup("origin")
  for _,v := range remotes {
    remote, _ := repo.Remotes.Lookup(v)
    if default_remote == nil {
      default_remote = remote
    }
    name := remote.Url()
    fmt.Println(
      helpers.Colorize(v, helpers.Green()),
      helpers.Colorize(name, helpers.Blue()),
    )
  }

  fmt.Println(default_remote.Name())
}
