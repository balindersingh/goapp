package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path/filepath"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  port:=os.Getenv("PORT")
  if port==""{
    port="5000"
  }
  http.ListenAndServe(":"+port, nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", filepath.Clean(r.URL.Path))  
  if r.URL.Path=="/"{
	log.Println("path is root")
	fp = "templates/home.html"
  }
  log.Println("path is :"+fp)
  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}