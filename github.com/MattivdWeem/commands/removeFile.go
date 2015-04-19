package commands
import (
  "fmt"
  "os/exec"
  )

  /*
    delete a file
  */
  func RemoveFile(file string){
    outputRm, errorsRm := exec.Command("rm",file).Output()
    if errorsRm != nil {
      fmt.Printf("%s", errorsRm)
    }
    fmt.Printf("%s", outputRm)
  }
