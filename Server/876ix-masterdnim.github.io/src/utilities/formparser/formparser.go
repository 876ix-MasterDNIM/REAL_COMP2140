package formparser
import (
	"../../datastructures"
	"net/http"
	"strings"
)

func ParseLoginForm(request *http.Request) datastructures.Login {
	request.ParseForm()

	var loginStruct = new (datastructures.Login)
	loginStruct.Username = strings.Join(request.Form["username"], "")
	loginStruct.Password = strings.Join(request.Form["password"], "")

	return *loginStruct
}

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