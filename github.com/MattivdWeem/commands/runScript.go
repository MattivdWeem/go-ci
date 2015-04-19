package commands

import(
  "fmt"
  "os/exec"
  "os"
  )

/*
  Run the script you've given.
*/
func RunScript(script string) bool{
  fmt.Println(" >> Scripts started")
  if _, err := os.Stat(script); err == nil {
    _,err := exec.Command("./"+script).Output()
    if err != nil {
      return false
    }
    fmt.Println(" >> Scripts succeeded")
    return true
  }
  return false
}
