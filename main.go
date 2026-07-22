package main

import (
	"log"
	"net/http"
)

func rootRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/courses.html")
}

func legacyBlogRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", rootRedirect)
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/blog", blogPage)
	http.HandleFunc("/experiences", legacyBlogRedirect)
	http.HandleFunc("/courses", legacyBlogRedirect)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
