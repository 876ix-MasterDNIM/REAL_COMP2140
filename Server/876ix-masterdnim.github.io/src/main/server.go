package main
import (
	"net/http"
	"../datastructures"
	"../utilities/formparser"
	"../utilities/logger"
	"../utilities/databaseutils"
	"../utilities/servehtml"
	"fmt"
)

var userLoggedIn bool = false

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/purchase", purchase)

	mux.Handle("/", http.FileServer(http.Dir("Client")))

	mux.Handle("/Client/", http.StripPrefix("/Client/",
		http.FileServer(http.Dir("Client"))))

	http.ListenAndServe(":8000", mux)
}

func signup (response http.ResponseWriter, request *http.Request) {
	var method string = request.Method
	switch method {

	case "GET":
		//Path of signup.html file to be served on get requests.
		path := "/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/signup.html"
		servehtml.ServeHtml(path, response)

	case "POST":
		var test datastructures.Signup = formparser.ParseSignUpForm(request)
		logger.LogSignup(&test)
	}
}

func login (response http.ResponseWriter, request *http.Request) {
	var method string = request.Method

	switch method {

	case "GET":
		//Path of login.html file to be served on get requests.
		path := "/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/login.html"
		servehtml.ServeHtml(path, response)

	case "POST":
		var test datastructures.Login = formparser.ParseLoginForm(request)
		logger.LogUser(&test)
		if databaseutils.Authenticate(&test) {
			userLoggedIn = true
			http.Redirect(response, request, "/purchase", 302)
		} else {
			userLoggedIn = false
			http.Redirect(response, request, "/login", 302)
		}
	}

}

func purchase (response http.ResponseWriter, request *http.Request) {
	var method string = request.Method
	switch method {
	case "GET":
		if userLoggedIn {
			var path string = "/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/purchase.html"
			servehtml.ServeHtml(path, response)
		} else {
			http.Redirect(response, request, "/login", 302)
		}
		//http.Redirect(response, request, "/login", 302)


	case "POST":
		//TODO
		fmt.Println("todo post")
	}
}

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

//	if request.Method == "GET" {
////		//var path string = string(request.URL.Path[1:])
////		t, _ := ioutil.ReadFile("login.html")
////		//var contentType string = addContentType(path)
////		response.Header().Set("Content-Type", "text/html")
////		fmt.Println(response.Header().Get("Content-Type"))
//
//		fmt.Println(request.URL.Path)
//
////		response.Write(t)
//
//	} else {
//		request.ParseForm()
//		fmt.Println("username:", request.Form["username"])
//		fmt.Println("password:", request.Form["password"])
//
//	}

