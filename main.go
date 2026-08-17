package main

var tmpl *template.Template

func main(){
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nill {
		log.Fatalf("Error parsing template: %v", err)
	}
}
