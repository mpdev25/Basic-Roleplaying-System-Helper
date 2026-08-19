package main
import (
	"http/template"
	"log"
	"net/http"
)
var tmpl *template.Template

func main(){
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nill {
		log.Fatalf("Error parsing template: %v", err)
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/generate", handleGenerate)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
