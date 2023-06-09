package main //логическая абст.
import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
func main() {

	http.HandleFunc("/", Handler)

	log.Println("Server Ok 🥛 ")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
