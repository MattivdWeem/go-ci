package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
     "os/exec"
     "math/rand"
     "encoding/json"
  )

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var buildFile string = "build.json"

type settings struct {
    Target string
}



/*
  handle all requests
*/
func handler(w http.ResponseWriter, r *http.Request) {
  tmpDir := randSeq(16) + "/"
  fmt.Println(" >> Cloning your repository")
  if gitClone("http://github.com/mattivdweem/learning-go",tmpDir) {

    fmt.Println(" >> Looking for a build file")
    if _, err := os.Stat("./"+tmpDir+buildFile); err == nil {
        fmt.Println(" >> processing build file")

        jsonData, err := ioutil.ReadFile("./"+tmpDir+buildFile)
        fmt.Println(jsonData)
        if err != nil { }

        var conf settings
        err=json.Unmarshal(jsonData, &conf)
        if err!=nil{
            fmt.Print("Error:",err)
        }

        var destination string = conf.Target

        // check if there are scripts to run yes? run them



        // no scripts/  if succeeded continue : 

        fmt.Println(" >> Deploying your project to "+ destination)
        if rSync(tmpDir, destination) {
          fmt.Println(" >> Deployment succeeded")
        }
    } else {
      fmt.Println(" >> No build file found..")
    }

  }

  // Remove the dir anyway
  removeDir(tmpDir);
}


/*
 Clone a git repository in the given directory
*/
func gitClone(repository string, dir string) bool{
  out,err := exec.Command("git","clone",repository,dir).Output()
  if err != nil {
    fmt.Printf("%s", err)
    return false
  }

  fmt.Printf("%s", out)
  return true
}


/*
  Remove a directory recursively
*/
func removeDir(dir string){
  outputRm, errorsRm := exec.Command("rm","-r",dir).Output()
  if errorsRm != nil {
    fmt.Printf("%s", errorsRm)
  }
  fmt.Printf("%s", outputRm)
}


/*
  Rsync a folder / it's contents into another folder
*/
func rSync(dir string, output string) bool{
  outputRsync, errorsRsync := exec.Command("rsync","-q","-av",dir,output).Output()
  if errorsRsync != nil {
    fmt.Printf("%s", errorsRsync)
    return false
  }

  fmt.Printf("%s", outputRsync)
  return true
}


/*
  When starting the script, run a little webserver on port 3768
*/
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":3768", nil)
}


/*
  Create a randomized string
*/
func randSeq(n int) string {
  b := make([]rune, n)
  for i := range b {
    b[i] = letters[rand.Intn(len(letters))]
  }
  return string(b)
}
