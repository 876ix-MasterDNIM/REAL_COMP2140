package logger
import (
	"../../datastructures"
	"fmt"
)

func LogUser(login *datastructures.Login) {
	fmt.Println("Username: " + login.Username)
	fmt.Println("Password: " + login.Password)
}

func LogSignup(signup *datastructures.Signup) {
	fmt.Println("Firstname: " + signup.Firstname)
	fmt.Println("Lastname: " + signup.Lastname)
	fmt.Println("Email: " + signup.Email)
	fmt.Println("Username: " + signup.Username)
	fmt.Println("Password: " + signup.Password)
	fmt.Println("Credit Card: " + signup.CreditCard)
	fmt.Println("CCV: " + string(signup.CCVNumber))
	fmt.Println("Credit Card Type: " + signup.CreditCardType)
	fmt.Println("Company: " + signup.Company)
	fmt.Println("Telephone: " + signup.Telephone)
}