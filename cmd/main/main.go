import (
    "net/http"
    "github.com/stewyb314/api"
func main() {
	http.HandleFunc("/", api.GetRoot)
	http.HandleFunc("/hello", api.GetHello)

	err := http.ListenAndServe(":3333", nil)
}
