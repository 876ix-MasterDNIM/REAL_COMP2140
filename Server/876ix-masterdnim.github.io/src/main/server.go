package main
import (
	"mime"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"net/http"
)


func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/signup", login)

	mux.Handle("/", http.FileServer(http.Dir("Client")))

	mux.Handle("/Client/", http.StripPrefix("/Client/",
		http.FileServer(http.Dir("Client"))))
//
//	mux.Handle("/Client/styles/", http.StripPrefix("/Client/styles/",
//		http.FileServer(http.Dir("Client/styles"))))

	http.ListenAndServe(":8000", mux)
}

func login (response http.ResponseWriter, request *http.Request) {
	//response.Write([]byte(request.Method))
	path := "/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/signup.html"
	fmt.Println(path)

	loginPage, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	} else {
		response.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
		response.Write(loginPage)
	}
}
