package main
import (
	"net/http"
	"../datastructures"
	"../utilities/formparser"
	"../utilities/logger"
	"../utilities/databaseutils"
	"../utilities/servehtml"
	"html/template"
	"fmt"
)

var userLoggedIn bool = false
var purchaseReport datastructures.Report
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/purchase", purchase)
	mux.HandleFunc("/report", report)

	mux.Handle("/", http.FileServer(http.Dir("Client")))

	mux.Handle("/Client/", http.StripPrefix("/Client/",
		http.FileServer(http.Dir("Client"))))

	http.ListenAndServe(":8000", mux)
}

func report (response http.ResponseWriter, request *http.Request) {
	var method string = request.Method
	switch method {

	case "GET":
		if userLoggedIn {
			tpl, err := template.ParseFiles("/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/report.html")

			if err != nil {
				panic(err)
			} else {
				//FIXME
				purchaseReport.Price = 0.0
				err = tpl.Execute(response, purchaseReport)
			}
		} else {
			http.Redirect(response, request, "/login", 302)
		}


	case "POST":
		var html string = `<!DOCTYPE html>
			<html lang="en">

			<head>
			<meta charset="UTF-8">
			<title>Report</title>
			<link href="styles/materialize.css" type="text/css" rel="stylesheet">
			<link href="styles/styles.css" type="text/css" rel="stylesheet">
			<script src="scripts/jquery-2.1.4.min.js" type="application/javascript"></script>
			<script src="scripts/materialize.min.js" type="application/javascript"></script>
			<style>
			p {
			color: #fff;
			}
			</style>
			</head>

			<body class="teal darken-3">

			<div class="card-panel">
				<h6 class="red-text">WE THANK YOU FOR YOUR SUPPORT AND LOOK FORWARD TO SERVING YOU IN THE FUTURE</h6>
			</div>

			</body>
			</html>`

		tpl, err := template.New("response").Parse(html)

		if err != nil {
			panic(err)
		} else {
			err = tpl.Execute(response, "")
		}
	}
}
func signup (response http.ResponseWriter, request *http.Request) {
	var method string = request.Method
	switch method {

	case "GET":
		path := "/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/signup.html"
		servehtml.ServeHtml(path, response)

	case "POST":
		var user datastructures.Signup = formparser.ParseSignUpForm(request)

		if databaseutils.EmailExists(&user) {
			response.Write(
				[]byte("Email already exists so please go back and fill out the form again, this time with a different email address"))
		} else {
			databaseutils.CreateUser(&user)
		}
	}
}

func login (response http.ResponseWriter, request *http.Request) {
	var method string = request.Method

	switch method {

	case "GET":
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

	case "POST":
		var purchase datastructures.Purchase = formparser.ParsePurchaseForm(request)
		purchaseReport.PurchaseInfo = purchase
		fmt.Println(purchaseReport)
		http.Redirect(response, request, "/report", 302)
	}
}
