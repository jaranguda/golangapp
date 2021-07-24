package main
 
import (
    "fmt"
    "net/http"
    "os"
)
 
func home(w http.ResponseWriter, req *http.Request) {
 
    hostname, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, "Hostname: %s", hostname)
}
 
func main() {
    http.HandleFunc("/", home)
    http.ListenAndServe(":8000", nil)
}
