
package MgoInit
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)
var (
    session   *mgo.Session
)
func Session() *mgo.Session {
   if(session==nil){
	   dialInfo := &mgo.DialInfo{
        Addrs:     []string{"211.159.175.75:28018"},
        Direct:    false,
        Timeout:   time.Second * 1,
		Username:  "lixin",
		Database: "NGDATA",
        Password:  "lixin888",
        PoolLimit: 4096, // Session.SetPoolLimit
                }
	   var err error
	   session,err:= mgo.DialWithInfo(dialInfo)
	   session.SetMode(mgo.Monotonic, true)
	   if err!=nil{
		   fmt.Println("error session")
		   panic(err)
	   }
	   return session.Copy()
   }else{
	   return session.Copy()
   }
}

// func GetDb() *mgo.Collection {
// 	if(db==nil){
// 		session, err := mgo.Dial("211.159.175.75:27017")
// 		if err!=nil{
// 			panic(err)
// 		}
// 		//Optional. Switch the session to a monotonic behavior.
// 		session.SetMode(mgo.Monotonic, true)
// 		db = session.DB("NGDATA").C("blog")
// 		return db
// 	}else {
// 		return db
// 	}
// }