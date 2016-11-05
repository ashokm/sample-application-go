package main
import (
        "fmt"
        "net/http"
        "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
        h, _ := os.Hostname()
        fmt.Fprintf(w, "<html><head></head><body><h1>sample-application-go!</h1><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.<p><img src=\"https://avatars2.githubusercontent.com/u/1902568?v=3&s=466\"><p>Hi there, I'm served from %s!</body></html>",h)
}

func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8484", nil)
}