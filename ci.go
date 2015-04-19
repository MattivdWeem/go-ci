package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
     "encoding/json"
     "github.com/MattivdWeem/stringutil"
     "github.com/MattivdWeem/commands"

  )

var buildFile string = "build.json"

type settings struct {
    Target string
    Script string
}

/*
  When starting the script, run a little webserver on port 3768
*/
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":3768", nil)
}



/*
  handle all requests
*/
func handler(w http.ResponseWriter, r *http.Request) {
  tmpDir := stringutil.RandSeq(16) + "/"
  fmt.Println(" >> Cloning your repository")
  if commands.GitClone("http://github.com/mattivdweem/learning-go",tmpDir) {

    fmt.Println(" >> Looking for a build file")
    build := "./"+tmpDir+buildFile

    if _, err := os.Stat(build); err == nil {
        fmt.Println(" >> processing build file")

        conf := readSettings(build)
        var destination string = conf.Target

        if conf.Script != "" {
          if !commands.RunScript("./"+tmpDir+conf.Script){
            fmt.Println("Script failed")
            os.Exit(1)
          }
        }

        // remove dev assets
        commands.RemoveDir(tmpDir+".git");
        commands.RemoveFile(tmpDir+buildFile);

        // rsync everything
        fmt.Println(" >> Deploying your project to "+ destination)
        if commands.Rsync(tmpDir, destination) {
          fmt.Println(" >> Deployment succeeded")
        }
    } else {
      fmt.Println(" >> No build file found..")
    }

  }

  // Remove the dir anyway
  commands.RemoveDir(tmpDir);
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
