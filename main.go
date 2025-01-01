package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mudzi42/lenslocked/controllers"
	"github.com/mudzi42/lenslocked/templates"
	"github.com/mudzi42/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

// func main() {
// exercise in errors
// err := B()
// if errors.Is(err, ErrNotFound) {
// 	fmt.Println("ErrNotFound")
// }

//
// r := chi.NewRouter()

// tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
// r.Get("/", controllers.StaticHandler(tpl))

// // Or inline everything...
// r.Get("/contact", controllers.StaticHandler(
// 	views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
// r.Get("/faq", controllers.StaticHandler(
// 	views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))
// r.Get("/page1", controllers.StaticHandler(
// 	views.Must(views.Parse(filepath.Join("templates", "page1.gohtml")))))
// r.Get("/page2", controllers.StaticHandler(
// 	views.Must(views.Parse(filepath.Join("templates", "page2.gohtml")))))

// r.Get("/", controllers.StaticHandler(
// 	views.Must(views.ParseFS(templates.FS, "home.gohtml"))))
// r.Get("/contact", controllers.StaticHandler(
// 	views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
// r.Get("/faq", controllers.FAQ(
// 	views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))
// r.Get("/page1", controllers.StaticHandler(
// 	views.Must(views.ParseFS(templates.FS, "page1.gohtml"))))
// r.Get("/page2", controllers.StaticHandler(
// 	views.Must(views.ParseFS(templates.FS, "page2.gohtml"))))

// 	r.Get("/", controllers.StaticHandler(views.Must(
// 		views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml"))))
// 	r.Get("/contact", controllers.StaticHandler(views.Must(
// 		views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml"))))

// 	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	})
// 	fmt.Println("Starting the server on :3000...")
// 	http.ListenAndServe(":3000", r)
// }

// var ErrNotFound = errors.New("not found")

// func A() error {
// 	return ErrNotFound
// }

// func B() error {
// 	err := A()
// 	if err != nil {
// 		return fmt.Errorf("b: %w", err)
// 	}
// 	return nil
// }
