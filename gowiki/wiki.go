package main

// Import fmt and os pkgs from Go standard library.
// Add additional functionality, add more pkgs to import declaration.

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp" // adding regular expression for title validation
)

// have to remove unused imports, go treats as a compile error

// html template lets us keep the HTML in a separate file
// can change layout of the edit page without modifying underlying Go code

// Data Structures.
// A Wiki consists of a series of interconnected pages.
// Each page has a title and a body.

type Page struct {
	Title string
	Body  []byte // byte slice, not a string only because it is the type expected by io libraries.
}

// slices: https://go.dev/blog/slices-intro
// convenient and efficient way of working with sequences of typed data.
// analagous to arrays in other languages but have unusual properties.
// built on top of Go's array type.
// array type specifies a length and an element type.
// [4]int is an array of 4 integers.
// An array's size is fixed.
// Arrays can be indexed in usual way, s[n] accesses nth element, starting from 0.
// var a [4]int a[0] = 1

// arrays can be inflexible, not seen often in Go code
// slices are everywhere
// []T t is type of elements of the slice
// slice has no specified length, like an array

// Slicing does not copy the slice’s data.
// It creates a new slice value that points to the original array.
// This makes slice operations as efficient as manipulating array indices.

// the page Struct describes how page data wil be stored in memory
// but what about persistent storage
// have to create a save method on Page

// This is a method named save that takes as its receiver p,
// a pointer to Page .
// It takes no parameters,
// and returns a value of type error

// method saves the Page's Body to a text file.
// Title is used as the filename.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// save method returns error value to let app handle
// if anything goes wrong when writing a file
// and because that is the return type of WriteFile.
// standard library function that writes
// a byte slice to a file
// if successful, returns nil, zero value for pointers, interfaces and other types
// octal integer 0600 passed as a param to WriteFile
// indicates file should be created with read-write permissions
// for the current user only

// Loading Pages
// In loadPage, error isn't being handled yet;
// the "blank identifier" represented by the underscore (_) symbol
// is used to throw away the error return value
// (in essence, assigning the value to nothing)

// if ReadFile encounters an error, what happens?
// modify to return *Page and error

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "home.html")
}

// callers of this function can now check 2nd param
// if nil, page successfully loaded
// if not, error can handled by caller

// Write Main function to test code above
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Compile and Execute code
// go build wiki.go
// ./wiki
// A file named TestPage.txt is created, containing contents of p1. File is then read into the struct p2 and its Body element is printed to the screen.

// net/http pkg
// use net/http to serve wiki pages

// create handler that allows users to view a wiki page, handle urls prefixed with /view/

// modift viewhandler after creating view html template

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	// http.Redirect function adds an HTTP status code of http.StatusFound (302)
	// and a Location header to the HTTP response

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

// same templating code in both Handlers,
// lets remove duplicate code by moving templating code into its own function
// and modify handlers, view and edit

// handle errors in renderTemplate function

// http.Error function sends specified HTTP response code and error message

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handle non existent pages!
// error return value from loadPage continues to try and fill out template with no data
// if requested page does not exist, redirect to edit page to create the content
// update viewHandler

// validation
// serious security flaw, user can supply an arbitrary path to be read / written on the server
// to migitgate, we can write a function to validate the title with a regular expression

// to use this function, have to rewrite main function to intialize http using view handler

// visit http://localhost:8080/view/test to see test.txt served by our web server

// Editing Pages
// Have to create 2 new handlers to edit pages
// edit form
// save form
// first add to main and then write the functions

// loads page, or if it does not exist, an empy page struct and displays html form

// The function template.ParseFiles will read the contents of edit.html
// and return a *template
// The method t.Execute executes the template,
// writing the generated HTML to the http.ResponseWriter

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// the html template pkg helps gurantee that only safe and
// correct looking HTML is generated by template actions

// write saveHandler
// saveHandler handles the submission of forms located on the edit pages

// page title in url and form's only field body are stored in a new Page
// save method is called to write the data to a file
// client is redirected to view page
// value returned by form is of type string
// has to be converted to byte to fit into the Page struct

// update errors for saveHandler
// any errors that occur during p.save() will be reported to the user
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// error handling
// do not ignore errors
// when an error occurs, the program will have unintended behaviour
// better to handle errors, and return an error msg to the user
// if something goes wrong, the server will function exactly how we want
// and the user will be notified

// template caching

// remove inefficiency in the code
// renderTemplate and ParseFiles
// rendertemplate calls parsefiles every time a page is rendered
// a better approach is to call parsefiles
// once at program initialization
// parsing all templates into a single Template
// then use execute template method to render a specific template

// create global variable names templates
// intialize with parsefiles

// template.Must is a convenience wrapper that panics when passed a non-nil error value,
// and otherwise returns the *Template unaltered
// A panic is appropriate here;
// if the templates can't be loaded the only sensible thing to do is exit the program
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// modify renderTemplate to call templates.ExecuteTemplate method with specific template

// global variable to store validation expression
// regexp.MustCompile will parse and compile the regular expression
// it will panic if the expression compilation fails

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// now lets use the validpath expression to validate path and extract page title
// if title is valid, it will be returned along with a nil value
// if title is invalid, the function will write a 404 Not Found
// to create new error, have to import errors pkg

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // the title is the second subexpression
}

// function literals and closures
// catching error condition in each handler introduces a lot of repeated code
// can wrap each handler in a function that does this validation and error checking
// use Gos function literals
// rewrite function handlers to accept a title string

// define wrapper that takes function of string type and returns handlerfunc
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// extract page title from the Request
		// call provided handler 'fn'
		// returned function is a closure because it encloses values defined outside of it
		// variable fn is enclosed by the closure
		// variable fn will be one of our save, edit or view handlers
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// this works but hardcoded HTML is ugly and hard to write.
// there is a better way with the html/template pkg
