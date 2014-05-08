package main 

import (
		"bytes"
		"fmt"
		"github.com/jamesmcnamara/insta_ipsum"
		"log"
		"net/http"
		"os"
		"strings"
		"text/template"
		)
var logger = openLogFile("./.server_log.txt")

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

//Opens the given file for use as a log file for errors in this
//server. If the file does not exist, it creates it, otherwise,
//it opens it with append privileges 
func openLogFile(filename string) *log.Logger {
	log_file, err := os.OpenFile(filename, os.O_RDWR | os.O_APPEND, 0660)
	if err != nil {
		if os.IsNotExist(err) {
			log_file, _ = os.Create(filename)
		} else {
			panic(err)
		}
	}
	return log.New(log_file, "Server Error: ", log.Ldate | log.Ltime)
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
	data := make([]byte, stats.Size())
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
func router(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	if path := r.URL.Path; path == "/" || path == "" {
            fmt.Fprintf(w, webpageData["/resume"].content)
        } else if strings.Contains(path, "static") {
=======
	
	if path := r.URL.Path; path == "/" || path == "/" {
		fmt.Fprintf(w, webpageData['/resume'].content)
	} else if strings.Contains(path, "static") {
>>>>>>> 88f1349148148d3baf4e38112047c78e4793f8e9
		mux.ServeHTTP(w, r)
	} else if strings.Contains(path, "api") {
		query_map := r.URL.Query()
		paragraphs, ok := query_map["p"]
		if ok {
			fmt.Fprintf(w, ipsum.GetIpsum(strings.Join(paragraphs, ""), true))
		}
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
		logger.Println(err)
	}
	layout_temp, err := template.ParseFiles(html_prefix + "/layout.html")
	if err != nil {
		logger.Println(err)
	}
	templateHTMLFiles(layout_temp)
	mux.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", router)
	logger.Println(http.ListenAndServe(":80", nil))
}

