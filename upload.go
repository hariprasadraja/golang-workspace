package main

import (
  "log"
  "net/http"
  "io"
  //"fmt"
  "os"
  "bytes"
  "gopkg.in/mgo.v2"
  "github.com/julienschmidt/httprouter"

)


type Files struct {
 File []byte
}

func download(res http.ResponseWriter,req *http.Request, _ httprouter.Params) {
    session,_:= mgo.Dial("localhost:27017")
    Files := &Files{}
    res.Header().Set("Content-Disposition","attachment;")
    session.DB("uploads").C("files").Find(nil).One(Files)
    res.Write(Files.File)
    log.Println("file is downloaded")
  }
func uploadHandler(res http.ResponseWriter,req *http.Request, _ httprouter.Params) {


  file,_,_ := req.FormFile("file")

    buffer := bytes.Buffer{}
    io.Copy(&buffer,file)


    session,err := mgo.Dial("localhost:27017")
    if err!= nil {
      os.Exit(1)
    }
    defer session.Close()
    Files:= &Files {
    File: buffer.Bytes() }
    session.DB("uploads").C("files").Insert(Files)

    session.DB("uploads").C("files").Find(nil).One(Files)
    res.Write(Files.File)
    //fmt.Fprintf(res,Files.File)
  }
func main() {
  router := httprouter.New()
  router.POST("/upload",uploadHandler)
  router.GET("/download",download)
  http.ListenAndServe(":8080",router)
}
