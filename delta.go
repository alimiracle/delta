package main
import          auth "github.com/abbot/go-http-auth"
import "syscall"

import "bufio"
import          "path/filepath"

import "html/template"
import "strconv"
import "os"
import "io"
import "crypto/md5"
import "time"
import "log"
import "fmt"
import "net/http"
import "strings"
var (
         targetFolder string
         targetFile   string
         searchResult []string
)

func css(w http.ResponseWriter, req *http.Request) {
    path := req.URL.Path

    var contentType string
    if strings.HasSuffix(path, ".css") {
        contentType = "text/css"
    } else if strings.HasSuffix(path, ".svg") {
        contentType = "image/svg+xml"
    } else if strings.HasSuffix(path, ".png") {
        contentType = "image/png"
    } else {
        contentType = "text/plain"
    }

    f, err := os.Open(path)

    if err == nil {
        defer f.Close()
        w.Header().Add("Content-Type", contentType)

        br := bufio.NewReader(f)
        br.WriteTo(w)
    } else {
        w.WriteHeader(404)
    }
} 
        
 func findFile(path string, fileInfo os.FileInfo, err error) error {

         if err != nil {
                 fmt.Println(err)
                 return nil
         }

         // get absolute path of the folder that we are searching
         absolute, err := filepath.Abs(path)

         if err != nil {
                 fmt.Println(err)
                 return nil
         }

         if fileInfo.IsDir() {
                 // correct permission to scan folder?
                 testDir, err := os.Open(absolute)

                 if err != nil {
                         if os.IsPermission(err) {
                                 fmt.Println("No permission to scan ... ", absolute)
                                 fmt.Println(err)
                         }
                 }
                 testDir.Close()
                 return nil
         } else {
                 // ok, we are dealing with a file
                 // is this the target file?

                 // yes, need to support wildcard search as well

                 matched, err := filepath.Match(targetFile, fileInfo.Name())
                 if err != nil {
                         fmt.Println(err)
                 }

                 if matched {
                         // yes, add into our search result
                         add := absolute
                         searchResult = append(searchResult, add)
                 }
         }

         return nil
 }

func search(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        t, _ := template.ParseFiles("serch.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
    name := r.PostFormValue("name")
         targetFolder = "/apps"
         targetFile = name + "*"

         fmt.Fprintf(w, "<p> Searching for [" + targetFile + "] </p>")

         // sanity check
         testFile, err := os.Open(targetFolder)
         if err != nil {
                 fmt.Println(err)
         }
         defer testFile.Close()

         testFileInfo, _ := testFile.Stat()
         if !testFileInfo.IsDir() {
                 fmt.Println(targetFolder, " is not a directory!")
         }

         err = filepath.Walk(targetFolder, findFile)

         if err != nil {
                 fmt.Println(err)
         }

         // display our search result
sys := strconv.Itoa(len(searchResult))
         fmt.Fprintf(w, "\n\n <p> Found " + sys + " hits! </p>\n\n")

         for _, v := range searchResult {
path := v
file := filepath.Base(path)                 
fmt.Fprintf(w, "<a href=\"" + path + "\">" + file + "</a>\n")
         }
}
searchResult = searchResult[:0]

 } 
/* this is your end, this is your end */

func upload(w http.ResponseWriter, r *http.Request) {
       if r.Method == "GET" {
           crutime := time.Now().Unix()
           h := md5.New()
           io.WriteString(h, strconv.FormatInt(crutime, 10))
           token := fmt.Sprintf("%x", h.Sum(nil))

           t, _ := template.ParseFiles("upload.gtpl")
           t.Execute(w, token)
       } else {
           
           file, handler, err := r.FormFile("uploadfile")
           if err != nil {
               fmt.Println(err)
               return
           }
           defer file.Close()
           f, err := os.OpenFile("./apps/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
           if err != nil {
               fmt.Println(err)
               return
           }
           defer f.Close()
           io.Copy(f, file)
    fmt.Fprintf(w, "<p> uplode ok </p>");
       }
}

func home(w http.ResponseWriter, r *http.Request) {
           t, _ := template.ParseFiles("home.gtpl")
           t.Execute(w, nil)
}
func admin(w http.ResponseWriter, r *http.Request) {
           t, _ := template.ParseFiles("admin.gtpl")
           t.Execute(w, nil)
}
func remove(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("remove.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
    name := r.PostFormValue("name")
         err := os.Remove("apps/" + name)
         if err != nil {
                 fmt.Println(err)
                    return
         }
fmt.Fprintf(w, "remove ok")
}
}
func mkdir(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("new.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
    name := r.PostFormValue("name")
         err := os.Mkdir("apps/" + name, 0)
         if err != nil {
                 fmt.Println(err)
                    return
         }
}
                fmt.Fprintf(w,"ok")
 
}
func rename(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("rename.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
    name := r.PostFormValue("name")
    newname := r.PostFormValue("new name")

         err := os.Rename("apps/" + name, "apps/" + newname)
         if err != nil {
                 fmt.Println(err)
                    return
         }
fmt.Fprintf(w, "rename ok")
}
}
func handleFileServer(dir, prefix string) http.HandlerFunc {
    fs := http.FileServer(http.Dir(dir))
    realHandler := http.StripPrefix(prefix, fs).ServeHTTP
    return func(w http.ResponseWriter, req *http.Request) {
        log.Println(req.URL)
        realHandler(w, req)
    }
}
func main() {
syscall.Chroot(".")
         htpasswd := auth.HtpasswdFileProvider("passwd")
         admin_htpasswd := auth.HtpasswdFileProvider("passwd")

         authenticator := auth.NewBasicAuthenticator("Basic Realm", htpasswd)
         admin_authenticator := auth.NewBasicAuthenticator("Basic Realm", admin_htpasswd)
    http.HandleFunc("/css/", css) // setting router rule

 http.HandleFunc("/", home)

    http.HandleFunc("/apps/", auth.JustCheck(authenticator, handleFileServer("apps", "/apps/")))
    http.HandleFunc("/remove", auth.JustCheck(admin_authenticator, remove))
    http.HandleFunc("/rename", auth.JustCheck(admin_authenticator, rename))
    http.HandleFunc("/new", auth.JustCheck(admin_authenticator, mkdir))
    http.HandleFunc("/upload", auth.JustCheck(admin_authenticator, upload))
    http.HandleFunc("/search", auth.JustCheck(authenticator, search))
    http.HandleFunc("/admin", auth.JustCheck(admin_authenticator, admin))
    err := http.ListenAndServe(":80", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

}
