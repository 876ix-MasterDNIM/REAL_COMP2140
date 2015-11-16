package databaseutils
import (
	"gopkg.in/mgo.v2"
	"../../datastructures"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func Authenticate(user *datastructures.Login) bool {
	dbSession, err := mgo.Dial("127.0.0.1")


	if err != nil {
		panic(err)
	}

	defer dbSession.Close()

	dbSession.SetMode(mgo.Monotonic, true)
	collection := dbSession.DB("comp2140").C("dummyuser")

	var result datastructures.Login

	err = collection.Find(bson.M{"username" : user.Username, "password" : user.Password}).One(&result)


	fmt.Println(result.Username)
	fmt.Println(result.Password)

	if err != nil {
		return false
	} else {
		return true
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