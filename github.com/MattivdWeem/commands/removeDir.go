package commands
import (
  "fmt"
  "os/exec"
  )

/*
  Remove a directory recursively
*/
func RemoveDir(dir string){
  outputRm, errorsRm := exec.Command("rm","-r",dir).Output()
  if errorsRm != nil {
    fmt.Printf("%s", errorsRm)
  }
  fmt.Printf("%s", outputRm)
}
