package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"title"   : " BELAJAR BOLANG TRANS 7",
	"isLogin" : true,
}

type Blog struct {
	Title    string
	PostDate string
	NewPostdate int
	Author   string
	Content  string
}

var Blogs = []Blog{
	{
		Title:    "Learning struct",
		PostDate: "20 januari 2023",
		Author:   "intizam",
		Content:  "lorem ipsum amet sirod dorod ampuisi abginaik non shabbs laottnm bs",
	},
}

func main() {

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/project", project).Methods("GET")
	router.HandleFunc("/mainblog/{id}", mainblog).Methods("GET")
	router.HandleFunc("/new-blog", newblog).Methods("POST")
	router.HandleFunc("/delete/{id}", delete).Methods("GET")
	

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", router)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html ; Charset=utf-8")
	w.WriteHeader(http.StatusOK)
	templ, err := template.ParseFiles("html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("message" + err.Error()))
		return
	}
	templ.Execute(w, Data)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html ; Charset=utf-8")
	// w.WriteHeader(http.StatusOK)

	templ, err := template.ParseFiles("html/blog.html")
	// id,_ := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("message  " + err.Error()))
		return
	}
	resp := map[string]interface{}{
		"Data":  Data,
		"Blogs": Blogs,
	}
	templ.Execute(w, resp)
}

func mainblog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html ; Charset=utf-8")
	w.WriteHeader(http.StatusOK)

	templ, err := template.ParseFiles("html/mainblog.html")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {

		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("message  " + err.Error()))
		return
	}
	CatchBlog := Blog{}

	for i, data := range Blogs {
		if i == id {
			CatchBlog = Blog{
				Title:    data.Title,
				PostDate: data.PostDate,
				Author:   data.Author,
				Content:  data.Content,
			}
		}
	}

	var resp = map[string]interface{}{
		"Data": Data,
		"Blogs": CatchBlog,
	}

	templ.Execute(w, resp)
}

func newblog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
		return
	}

	Projectname := r.PostForm.Get("projectname")
	Description := r.PostForm.Get("description")
	StartDate := r.PostForm.Get("startDate")
	EndDate := r.PostForm.Get("endDate")

	
	start, _ := time.Parse("2006-01-02", StartDate)
	end, _ := time.Parse("2006-01-02", EndDate)
    diff := end.Sub(start)



	var refilData = Blog{

		Title:    Projectname,
		NewPostdate: int(diff.Hours() / 24),
		Author:   "intizam",
		Content:  Description,
	}
	Blogs = append(Blogs, refilData)

	http.Redirect(w, r, "/project", http.StatusMovedPermanently)
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "Charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	Blogs = append(Blogs[:id], Blogs[id+1:]...)
	http.Redirect(w, r, "/project", http.StatusMovedPermanently)
}



