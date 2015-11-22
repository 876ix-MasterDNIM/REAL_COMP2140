package serversetup

import (
	"fmt"
	"html/template"
	"net/http"
	"../datastructures"
	"../utilities/databaseutils"
	"../utilities/formparser"
	"../utilities/servehtml"
	"strings"
	"../utilities/pricecalc"
)

/**
 * Maps routes to their respective functions
 */
func RunServer() {

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

var userLoggedIn bool = false //keeps track of whether user is logged in or not
var purchaseInfo datastructures.Purchase //represents info collected on purchases form
var purchaseReport datastructures.Report //represents completed transaction
var companyInfo datastructures.CompanyInfo = databaseutils.GetCompanyInfo() //represents company rates


/**
 * Facilitates all features on the report page
 * @param  {[type]} response http.ResponseWriter [description]
 * @param  {[type]} request  *http.Request       [description]
 * @return {[type]}          [description]
 */
func report(response http.ResponseWriter, request *http.Request) {
	var method string = request.Method
	switch method {

	case "GET":
		if userLoggedIn {
			//Show report page if user is logged in
			tpl, err := template.ParseFiles("/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/report.html")
			if err != nil {
				panic(err)
			} else {
				purchaseReport.Price = pricecalc.AdCost(purchaseInfo, companyInfo, purchaseReport)
				err = tpl.Execute(response, purchaseReport)
			}
		} else {
			//reditrect user to login page if he/she is not logged in
			http.Redirect(response, request, "/login", 302)
		}

	case "POST":
		//Show thank you message after a succesful transaction
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
		databaseutils.CreateReport(&purchaseReport) //Log transaction to database
	}
}

/**
 * Facilitates all features on the sign up page
 * @param  {[type]} response http.ResponseWriter [description]
 * @param  {[type]} request  *http.Request       [description]
 * @return {[type]}          [description]
 */
func signup(response http.ResponseWriter, request *http.Request) {
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

/**
 * Facilitates all features on the login page
 * @param  {[type]} response http.ResponseWriter [description]
 * @param  {[type]} request  *http.Request       [description]
 * @return {[type]}          [description]
 */
func login(response http.ResponseWriter, request *http.Request) {
	var method string = request.Method

	switch method {

	case "GET":
		fmt.Println(companyInfo.SundayRate)
		path := "/Users/LeaderOfTheNewSchool/WebstormProjects/REAL_COMP2140/Client/login.html"
		servehtml.ServeHtml(path, response)

	case "POST":
		var user datastructures.Login = formparser.ParseLoginForm(request)
		if databaseutils.Authenticate(&user) {
			userLoggedIn = true
			http.Redirect(response, request, "/purchase", 302)
		} else {
			userLoggedIn = false
			http.Redirect(response, request, "/login", 302)
		}
	}

}
/**
 * Facilitates all features on the purchase page
 * @param  {[type]} response http.ResponseWriter [description]
 * @param  {[type]} request  *http.Request       [description]
 * @return {[type]}          [description]
 */
func purchase(response http.ResponseWriter, request *http.Request) {
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
		purchaseInfo = formparser.ParsePurchaseForm(request)
		purchaseReport.PurchaseInfo = purchaseInfo
		fmt.Println(purchaseInfo.ShowDates)
		http.Redirect(response, request, "/report", 302)
	}
}
