package main
import (

	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"html/template"
)

func main() {



//	mux := pat.New()
//	mux.Get("/singup", http.HandlerFunc(signup))
//	mux.Post("Server/github.876ix-masterdnim/src/signup.html", http.HandlerFunc(signupPost))
	//var handle *requestHandler = &requestHandler{}
//	fmt.Println("Server started on port 8000.")
//	http.Handle("/", handle)
//	http.ListenAndServe(":8000", new(requestHandler))

	http.HandleFunc("/Server/github.876ix-masterdnim/src/login.html", signup)
	http.HandleFunc("/Server/github.876ix-masterdnim/src/login", signup)
	http.HandleFunc("/Server/github.876ix-masterdnim/src/signup.html", signup)

	//http.Handle("/Server/github.876ix-masterdnim/src/signup.html", handle)

	http.Handle("/", http.FileServer(http.Dir("/Users/LeaderOfTheNewSchool/WebstormProjects/COMP2140")))
	http.ListenAndServe(":8000", nil)
}

func signupPost(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	request.ParseForm()
	fmt.Println("username:", request.Form["username"])
	fmt.Println("password:", request.Form["password"])
}

func signup(response http.ResponseWriter, request *http.Request) {
	var path string = string(request.URL.Path[1:]) + ".html"
	fmt.Println(path)
	fmt.Println("In signup")
	webpageData, err := ioutil.ReadFile(path)

	if err == nil {
		if request.Method == "GET" {

			var contentType string = addContentType(path)

			response.Header().Set("Content-Type", contentType)
			response.Write(webpageData)
		}
		fmt.Println(request.Method)
		if request.Method == "POST" {
			request.ParseForm()
			fmt.Println("username:", request.Form["username"])
			fmt.Println("password:", request.Form["password"])
		}
	} else {
		panic(err)
	}
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func login(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	if request.Method == "GET" {
//		//var path string = string(request.URL.Path[1:])
//		t, _ := ioutil.ReadFile("login.html")
//		//var contentType string = addContentType(path)
//		response.Header().Set("Content-Type", "text/html")
//		fmt.Println(response.Header().Get("Content-Type"))

		fmt.Println(request.URL.Path)

//		response.Write(t)

	} else {
		request.ParseForm()
		fmt.Println("username:", request.Form["username"])
		fmt.Println("password:", request.Form["password"])

	}
}

type requestHandler struct{}

func (handler *requestHandler) ServeHTTP(response http.ResponseWriter,
request *http.Request) {
//	var path string = string(request.URL.Path[1:])
//
//	//fmt.Println(path)
//	//fmt.Println(request.Method)
//
//	if path == "Server/github.876ix-masterdnim/src/signup" {
//		fmt.Println("Running signupPost now...")
//		signupPost(response, request)
//	}
//	webpageData, err := ioutil.ReadFile(path)
//
//	if err == nil {
//		var contentType string = addContentType(path)
//
//		response.Header().Set("Content-Type", contentType)
//		response.Write(webpageData)
//	} else {
//		response.WriteHeader(404)
//		response.Write([]byte(http.StatusText(404)))
//	}
//	//fmt.Println(request.URL.Path)
}

func addContentType(path string) string {
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/js; charset=utf-8"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpeg; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}
	return contentType
}


//	fmt.Println("Hello World")
//
//	dbSession, err := mgo.Dial("127.0.0.1")
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer dbSession.Close()
//
//	dbSession.SetMode(mgo.Monotonic, true)
//
//	c := dbSession.DB("comp2140").C("dummyuser")
//
////	err = c.Insert(&Dummy{Username: "comp2140", Password: "comp2140also"}, &Dummy{Username: "comp2140",
////	Password:"comp2140yetagain"})
////
////	if err != nil {
////		panic(err)
////	}
//
//	var result []Dummy
//	c.Find(bson.M{"username" : "comp2140"}).All(&result)
//	for _, dummyuser := range result {
//		fmt.Println(dummyuser.Password)
//	}
//}
//
//type Dummy struct {
//	ID bson.ObjectId `bson:"_id,omitempty"`
//	Username string
//	Password string
//}
