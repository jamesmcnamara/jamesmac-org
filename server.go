package main 

import (
		"bytes"
		"fmt"
		"text/template"
		"net/http"
		"os"
		"strings"
		)

var html_prefix string = "templates"
var html_files = []string{"/about", "/blog", "/code", "/contact", 
"/resume", "/under_construction"}

var webpageData map[string]*Page = make(map[string]*Page)

var mux *http.ServeMux = http.NewServeMux()

//Struct contains basic data about each webpage
type Page struct {
	title string
	content string
}

//consumes a file name, and returns a pointer to the 
//page struct containing it's data, or any errors 
//encountered
func loadPage(name string) (p *Page, err error) {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stats.Size(), stats.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	return &Page{name, string(data)}, nil
}

//Runs through the list of html_files listed in the slice
//above and initializes the map of webpage data with
//Page structs containing the data and returns any errors
func loadHTMLFiles() error {
	for _, file := range html_files {
		p, err := loadPage(html_prefix + file + ".html")
		if err != nil {
			return err
		}
		webpageData[file] = p
	}
	return nil
}

//Applies the given template to each of the files in
//the map of webpages
func templateHTMLFiles(layout *template.Template) {
	var buff bytes.Buffer
	buff_ptr := &buff
	for _, webpage := range webpageData {
		layout.Execute(buff_ptr, webpage.content)
		webpage.content = buff_ptr.String()
		buff.Reset()
	}
}

//serves the correct page or static file as specified by
//the given request pointers url path
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.Contains(path, "static") {
		mux.ServeHTTP(w, r)
	} else if webpage, ok := webpageData[path]; ok {
		fmt.Fprintf(w, webpage.content)	
	} else {
		fmt.Println(path + " was not found")
	}
}

//preloads all pages and begins the reactor loop
func main() {
	err := loadHTMLFiles()
	if err != nil {
		panic(err)
	}
	layout_temp, err := template.ParseFiles(html_prefix + "/layout.html")
	if err != nil {
		panic(err)
	}
	templateHTMLFiles(layout_temp)
	mux.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:80", nil)
}
