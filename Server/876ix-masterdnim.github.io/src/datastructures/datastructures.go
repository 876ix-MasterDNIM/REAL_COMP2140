package datastructures

type Login struct {
	Username string
	Password string
}

type Signup struct {
	Firstname         string
	Lastname          string
	Email             string
	Username          string
	Password          string
	ConfirmedPassword string
	CreditCard        string
	CCVNumber         string
	CreditCardType    string
	Telephone         string
	Company           string
}

type Purchase struct {
	Firstname    string
	Lastname     string
	CompanyName  string
	Contact      string
	StartDate    string
	EndDate      string
	ShowDates    string
	AdColors     string
	Size         string
	DesignForYou string
}

type Report struct {
	PurchaseInfo Purchase
	Price        float64
}