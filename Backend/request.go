package backend

import (
	"log"
	"net/http"
	"os"
	"text/template"

	functions "ascii-art/Functions"
)

type TemplateData struct {
	Input    string
	AsciiArt []string
	Fontfile string
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form data:", err)
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		input := r.Form.Get("text")
		fontFile := r.Form.Get("File")
		if fontFile == "" {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		asciiArt, err := functions.Ascii(input, fontFile)
		
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				return
			} else {
				http.Error(w, "400 bad request", http.StatusBadRequest)
				return
			}
		}

		data := TemplateData{
			Input:    input,
			AsciiArt: asciiArt,
			Fontfile: fontFile,
		}

		temp := template.Must(template.ParseFiles("Templates/index.html"))
		err = temp.Execute(w, data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}

	}
}
