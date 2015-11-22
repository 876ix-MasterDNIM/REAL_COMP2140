package databaseutils

import (
	"../../datastructures"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/**
 * Retrieves company rates from database
 */
func GetCompanyInfo() datastructures.CompanyInfo {
	dbSession, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer dbSession.Close()

	dbSession.SetMode(mgo.Monotonic, true)
	collection := dbSession.DB("comp2140").C("company")

	var result datastructures.CompanyInfo

	err = collection.Find(nil).One(&result)

	if err != nil {
		panic(err)
	} else {
		return result
	}
}

/**
 * Creates a user in database
 * @param {[type]} user *datastructures.Signup [description]
 */
func CreateUser(user *datastructures.Signup) {
	dbSession, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer dbSession.Close()

	dbSession.SetMode(mgo.Monotonic, true)
	collection := dbSession.DB("comp2140").C("users")

	err = collection.Insert(&user)

	if err != nil {
		panic(err)
	}
}

/**
 * Creates a report in database
 * @param {[type]} report *datastructures.Report [description]
 */
func CreateReport(report *datastructures.Report) {
	dbSession, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer dbSession.Close()

	dbSession.SetMode(mgo.Monotonic, true)
	collection := dbSession.DB("comp2140").C("purchases")

	err = collection.Insert(&report)

	if err != nil {
		panic(err)
	}
}

/**
 * Facilitates authentication of users
 * @param {[type]} user *datastructures.Login [description]
 */
func Authenticate(user *datastructures.Login) bool {
	dbSession, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer dbSession.Close()

	dbSession.SetMode(mgo.Monotonic, true)
	collection := dbSession.DB("comp2140").C("users")

	var result datastructures.Login

	err = collection.Find(bson.M{"username": user.Username, "password": user.Password}).One(&result)

	if err != nil {
		return false
	} else {
		return true
	}
}

/**
 * Checks if email already exists in database
 * @param {[type]} user *datastructures.Signup [description]
 */
func EmailExists(user *datastructures.Signup) bool {
	dbSession, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer dbSession.Close()
	dbSession.SetMode(mgo.Monotonic, true)

	collection := dbSession.DB("comp2140").C("users")

	var result datastructures.Signup

	err = collection.Find(bson.M{"email": user.Email}).One(&result)

	if err != nil {
		return false
	} else {
		return true
	}
}
