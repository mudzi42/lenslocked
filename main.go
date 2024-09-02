package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:mudzi42@gmail.com\">mudzi42@gmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	<li>Q: Is there a free version? A: Yes! We offer a free trial for 30 days on any paid plans.</li>
	<li>Q: What are your support hours? A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</li>
	<li>Q: How do I contact support? A: Email us - <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>.</li>
	</ul>
	`)

}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		// TODO: handle the page not found error
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>Page Not Found</h1>")
// 	}
// }

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		// TODO: handle the page not found error
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>Page Not Found</h1>")
// 	}
// }

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", r)
}
