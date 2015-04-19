package commands

import(
  "fmt"
  "os/exec"
  )

/*
 Clone a git repository in the given directory
*/
func GitClone(repository string, dir string) bool{
  out,err := exec.Command("git","clone",repository,dir).Output()
  if err != nil {
    fmt.Printf("%s", err)
    return false
  }

  fmt.Printf("%s", out)
  return true
}
