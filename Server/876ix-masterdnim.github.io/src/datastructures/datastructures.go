package datastructures

/**
 * Structure to facilitate log in function
 */
type Login struct {
	Username string
	Password string
}

/**
 * Structure to represent information collected on the sign up form
 */
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

/**
 * Structure to represent information collected on the purchase form
 */
type Purchase struct {
	Firstname    string
	Lastname     string
	CompanyName  string
	Contact      string
	StartDate    string
	EndDate      string
	ShowDates    string
	AdColors     string
	Columns      int8
	Rows         int8
	DesignForYou string
}

/**
 * Structure to represent a completed purchase
 */
type Report struct {
	PurchaseInfo Purchase
	Price        float64
}

/**
 * Structure to represent the rates of various advertisement customisations
 * retrieved from company's database
 */
type CompanyInfo struct {
	Tax          float64
	WeekdayRate  float64
	SaturdayRate float64
	SundayRate   float64
}
