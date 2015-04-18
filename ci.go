package main

import (
    "fmt"
    "net/http"
     "os/exec"
     "math/rand"
  )

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var destination string = "/Users/mattivandeweem/Development/go/live/"

func handler(w http.ResponseWriter, r *http.Request) {
  tmpDir := randSeq(16) + "/"
  fmt.Println(" >> Cloning your repository")
  if gitClone("http://github.com/mattivdweem/learning-go",tmpDir) {
    fmt.Println(" >> Deploying your project to "+ destination)
    if rSync(tmpDir, destination) {
      fmt.Println(" >> Deployment succeeded")
    }
  }


  // Remove the dir anyway
  removeDir(tmpDir);


}

func gitClone(repository string, dir string) bool{
  out,err := exec.Command("git","clone",repository,dir).Output()
  if err != nil {
    fmt.Printf("%s", err)
    return false
  }

  fmt.Printf("%s", out)
  return true
}

func removeDir(dir string){
  outputRm, errorsRm := exec.Command("rm","-r",dir).Output()
  if errorsRm != nil {
    fmt.Printf("%s", errorsRm)
  }
  fmt.Printf("%s", outputRm)
}

func rSync(dir string, output string) bool{
  outputRsync, errorsRsync := exec.Command("rsync","-q","-av",dir,output).Output()
  if errorsRsync != nil {
    fmt.Printf("%s", errorsRsync)
    return false
  }

  fmt.Printf("%s", outputRsync)
  return true
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":3768", nil)
}

func randSeq(n int) string {
  b := make([]rune, n)
  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}
