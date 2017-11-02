package main
import (
      //"fmt"
      "io/ioutil"
      "net/http"
      "html/template"
      "fmt"
      "log"
      "regexp"
  )
/* Creating a file ,updating and displaying text inside the file */
func init() {
log.SetFlags(log.LstdFlags|log.Lshortfile)
}
type Page struct {

  Title string
  Body []byte
}

func (p *Page) save() error {

 filename:= p.Title + ".txt"
 return ioutil.WriteFile(filename,p.Body,0600)

}

func loadPage(title string) (*Page,error){
  filename := title + ".txt"
  body,err := ioutil.ReadFile(filename)
  if err != nil {
    return nil,err
  }
  return &Page{Title: title, Body: body},nil
}

func main() {
  // p1:= &Page{Title: "TestPage",Body: []byte("This is the sample page")}
  // p1.save()
  // p2,_ := loadPage("TestPage")
  // fmt.Println(string(p2.Body))
  // http.HandleFunc("/",handler)
   // http.ListenAndServe(":8080",nil)
   var template = template.Must(template.ParseFiles("edit.html","view.html"))
   var validpath = regexp.MustCompile("^/(edit|save|view)/([a-ZA-Z0-9]+)$")
   http.HandleFunc("/view/", makeHandler(viewHandler))
   http.HandleFunc("/edit/", makeHandler(editHandler))
   http.HandleFunc("/save/" ,makeHandler(saveHandler))
   error:=http.ListenAndServe(":8080",nil)
   log.Println(error)
}
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"Hi, there, I love %s!", r.URL.Path[1:])
  fmt.Println("Inside the println")
}
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
  title:= r.URL.Path[len("/view/"):]
  p,err := loadPage(title)
  if err != nil {
      http.Redirect(w, r, "/edit/"+title, http.StatusFound)
      return
  }

  renderTemplate(w,"view",p)

  }

//   fmt.Fprintf(w,"<h1> %s </h1> <div> %s </div>", p.Title,p.Body)
// }
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
  //  title := r.URL.Path[len("/edit/"):]
   title,err := getTitle(w,r)
   if err != nil {
     return
   }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title,}
    }
    // t, _ := template.ParseFiles("edit.html")
    // t.Execute(w, p)
    renderTemplate(w,"edit",p)
}

func renderTemplate(w http.ResponseWriter, tmpl string,p *Page) {
  t,err:= template.ExceuteTemplate(w,tmpl+".html",p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  err = t.Execute(w,p)
  if err != nil {
    http.Error(w.err.Error(), http.StatusInternalServerError)
  }


}
func saveHandler(w http.ResponseWriter,r *http.Request, title string) {
  title,err := getTitle(w,r)
  if err!=nil {
    return
  }
  body := r.FormValue("body")
  p := &Page{Title: title, Body: []byte(body)}
  err :=p.save()
  if err != nil {
    http.Error(w,err.Error().http.StatusInternalServerError)
    return
  }
  http.Redirect(w,r,"/view/"+ title,http.StatusFound)
}
func getTitle(w http.ResponseWriter,r *http.Request) (string,error){
m := validPath.FindStringSubmatch(r.URL.Path)
if m == nil {
  http.NotFound(w,r)
  retrun "", errors.New("Invalid Page Title")
}
retrun m[2],nil
}

fun makeHandler(fn func(http.ResponseWriter,*http.Request, title string)) http.HandleFunc {
  return func(w http.ResponseWriter,r *http.Request) {
    m := validPath.FindStringSubmatch(r.url.Path)
    if m == nil {
      http.NotFound(w,r)
      return
    }
    fn(w,r,m[2])
  }
}
