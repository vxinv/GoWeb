package HandleFun

 import(
	"gopkg.in/mgo.v2/bson"
	// "encoding/json"
	"MyLargstProgram/MgoInit"
	// "MyLargstProgram/Mytype"
	"fmt"
 )
func main(){
    session:=MgoInit.Session()
    defer session.Close()
	db:=session.DB("NGDATA").C("blog")
	// var blogs []Mytype.Blogs
	count,err:=db.Find(bson.M{"monney":bson.M{"$lt":20}}).Select(bson.M{"blogimages":1}).Count()
    if(err!=nil){
		panic(err)
	}
	// data,_:=json.Marshal(blogs)
	fmt.Println(count)
}