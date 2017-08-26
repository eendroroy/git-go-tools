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

  bi,_ := repo.NewBranchIterator(git.BranchRemote)

  bi.ForEach(func(branch *git.Branch, branchType git.BranchType) error {
    br_name, _ := branch.Name()
    fmt.Println(br_name)
    return nil
  })

  odb, _ := repo.Odb()

  count := 0

  odb.ForEach(func(id *git.Oid) error {
    commit, err := repo.LookupCommit(id)
    if err == nil {
      fmt.Println(
        helpers.Colorize(commit.Committer().Name, helpers.Blue()),
        helpers.Colorize(commit.Committer().Email, helpers.Green()),
        helpers.Colorize(commit.Message(), helpers.White()),
        commit.ParentCount(),
      )
    }
    count++
    return nil
  })

  fmt.Println(count)
}
