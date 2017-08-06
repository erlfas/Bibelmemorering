package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type indexHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type page struct {
	Tittel string
	Body   template.HTML
}

// ServeHTTP handles the HTTP request.
func (t *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles("public/index.html"))
	})

	p := &page{Tittel: "Bibelmemorering"}

	err := t.templ.Execute(w, p)
	if err != nil {
		fmt.Println("indexHandler: error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type resultHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *resultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles("public/result.html"))
	})

	r.ParseForm()
	tekst := r.Form.Get("tekst")
	valg := r.Form.Get("valg")

	fmt.Println("Tekst:", tekst)
	fmt.Println("Valg:", valg)

	data := map[string]interface{}{
		"original": tekst,
	}

	w.WriteHeader(http.StatusOK)

	switch valg {
	case "0":
		data["valg"] = "alle"
		data["a"] = template.HTML(htmlFilter(makeFirstTransformation(tekst)))
		data["b"] = template.HTML(htmlFilter(makeSecondTransformation(tekst)))
		data["c"] = template.HTML(htmlFilter(makeThirdTransformation(tekst)))
		data["d"] = template.HTML(htmlFilter(makeFourthTransformation(tekst)))
	case "1":
		data["valg"] = "a"
		body := htmlFilter(makeFirstTransformation(tekst))
		data["a"] = template.HTML(body)
	case "2":
		data["valg"] = "b"
		body := htmlFilter(makeSecondTransformation(tekst))
		data["b"] = template.HTML(body)
	case "3":
		data["valg"] = "c"
		body := htmlFilter(makeThirdTransformation(tekst))
		data["c"] = template.HTML(body)
	case "4":
		data["valg"] = "d"
		body := htmlFilter(makeFourthTransformation(tekst))
		data["d"] = template.HTML(body)
	}

	t.templ.Execute(w, data)
}
