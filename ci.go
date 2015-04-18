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
    Script string
}



/*
  handle all requests
*/
func handler(w http.ResponseWriter, r *http.Request) {
  tmpDir := randSeq(16) + "/"
  fmt.Println(" >> Cloning your repository")
  if gitClone("http://github.com/mattivdweem/learning-go",tmpDir) {

    fmt.Println(" >> Looking for a build file")
    build := "./"+tmpDir+buildFile

    if _, err := os.Stat(build); err == nil {
        fmt.Println(" >> processing build file")

        conf := readSettings(build)
        var destination string = conf.Target

        if conf.Script != "" {
          if !runScript("./"+tmpDir+conf.Script){
            os.Exit(1)
          }
        }

        removeFile("./"+tmpDir+".git");
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
  Run the script you've given.
*/
func runScript(script string) bool{
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


// read your build file
func readSettings(file string) settings{
  jsonData, err := ioutil.ReadFile(file)
  if err != nil { }

  var conf settings
  err=json.Unmarshal(jsonData, &conf)
  if err!=nil{
      fmt.Print("Error:",err)
  }
  return conf
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
  Remove a directory recursively
*/
func removeFile(file string){
  outputRm, errorsRm := exec.Command("rm",file).Output()
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
