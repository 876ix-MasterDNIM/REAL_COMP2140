package databaseutils
import (
	"gopkg.in/mgo.v2"
	"../../datastructures"
	"gopkg.in/mgo.v2/bson"
)

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

func CreateReport(report datastructures.Report) {
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

func Authenticate(user *datastructures.Login) bool {
	dbSession, err := mgo.Dial("127.0.0.1")


	if err != nil {
		panic(err)
	}

	defer dbSession.Close()

	dbSession.SetMode(mgo.Monotonic, true)
	collection := dbSession.DB("comp2140").C("users")

	var result datastructures.Login

	err = collection.Find(bson.M{"username" : user.Username, "password" : user.Password}).One(&result)

	if err != nil {
		return false
	} else {
		return true
	}
}

func EmailExists(user *datastructures.Signup) bool {
	dbSession, err := mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer dbSession.Close()
	dbSession.SetMode(mgo.Monotonic, true)

	collection := dbSession.DB("comp2140").C("users")

	var result datastructures.Signup

	err = collection.Find(bson.M{"email" : user.Email}).One(&result)

	if err != nil {
		return false
	} else {
		return true
	}
}
