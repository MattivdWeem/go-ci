package commands

import(
  "fmt"
  "os/exec"
  )

  /*
    Rsync a folder / it's contents into another folder
  */
  func Rsync(dir string, output string) bool{
    outputRsync, errorsRsync := exec.Command("rsync","-q","-av",dir,output).Output()
    if errorsRsync != nil {
      fmt.Printf("%s", errorsRsync)
      return false
    }

    fmt.Printf("%s", outputRsync)
    return true
  }
