package formparser

import (
	"net/http"
	"strconv"
	"strings"

	"../../datastructures"
)

/**
 * Parses login form
 * @param {[type]} request *http.Request [description]
 */
func ParseLoginForm(request *http.Request) datastructures.Login {
	request.ParseForm()

	var loginStruct = new(datastructures.Login)
	loginStruct.Username = strings.Join(request.Form["username"], "")
	loginStruct.Password = strings.Join(request.Form["password"], "")

	return *loginStruct
}

/**
 * Parses sign up form
 * @param {[type]} request *http.Request [description]
 */
func ParseSignUpForm(request *http.Request) datastructures.Signup {
	request.ParseForm()

	var signupStruct = new(datastructures.Signup)
	signupStruct.Firstname = strings.Join(request.Form["first"], "")
	signupStruct.Lastname = strings.Join(request.Form["last"], "")
	signupStruct.Email = strings.Join(request.Form["email"], "")
	signupStruct.Username = strings.Join(request.Form["username"], "")
	signupStruct.Password = strings.Join(request.Form["password"], "")
	signupStruct.ConfirmedPassword = strings.Join(request.Form["cpassword"], "")
	signupStruct.CreditCard = strings.Join(request.Form["cc"], "")
	signupStruct.CCVNumber = strings.Join(request.Form["ccv"], "")
	signupStruct.CreditCardType = strings.Join(request.Form["ctype"], "")
	signupStruct.Telephone = strings.Join(request.Form["telephone"], "")
	signupStruct.Company = strings.Join(request.Form["company"], "")

	return *signupStruct
}

/**
 * Parses purchase form
 * @param {[type]} request *http.Request [description]
 */
func ParsePurchaseForm(request *http.Request) datastructures.Purchase {
	request.ParseForm()

	var purchaseStruct = new(datastructures.Purchase)
	purchaseStruct.Firstname = strings.Join(request.Form["first"], "")
	purchaseStruct.Lastname = strings.Join(request.Form["last"], "")
	purchaseStruct.CompanyName = strings.Join(request.Form["cname"], "")
	purchaseStruct.Contact = strings.Join(request.Form["cnumber"], "")
	purchaseStruct.StartDate = strings.Join(request.Form["startdate"], "")
	purchaseStruct.EndDate = strings.Join(request.Form["enddate"], "")
	purchaseStruct.ShowDates = days(request)
	purchaseStruct.AdColors = colors(request)
	purchaseStruct.Columns, _ = strconv.ParseFloat(strings.Join(request.Form["columns"], ""), 10)
	purchaseStruct.Rows, _ = strconv.ParseFloat(strings.Join(request.Form["rows"], ""), 10)
	purchaseStruct.DesignForYou = strings.Join(request.Form["ownad"], "")

	return *purchaseStruct
}

/**
 * Retrieves colors requested on purchase form
 * @param  {[type]} request *http.Request [description]
 * @return {[type]}         [description]
 */
func colors(request *http.Request) string {
	colors := []string{"Red", "Yellow", "Blue"}
	trueColors := make([]string, 3)

	for i := 0; i < len(colors); i++ {
		temp := strings.Join(request.Form[colors[i]], "")
		if temp == "on" {
			trueColors = append(trueColors, colors[i])
		}
	}

	return strings.TrimSpace(strings.Join(trueColors, " "))
}

/**
 * Retrieves days advertisement is requested on purchase form
 * @param  {[type]} request *http.Request [description]
 * @return {[type]}         [description]
 */
func days(request *http.Request) string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	trueDays := make([]string, 6)

	for i := 0; i < len(days); i++ {
		temp := strings.Join(request.Form[days[i]], "")
		if temp == "on" {
			trueDays = append(trueDays, days[i])
		}
	}

	return strings.TrimSpace(strings.Join(trueDays, " "))
}
