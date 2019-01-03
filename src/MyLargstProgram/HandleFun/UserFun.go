package HandleFun
import(
	// "gopkg.in/mgo.v2/bson"
	// "encoding/json"
	"MyLargstProgram/MgoInit"
	"fmt"
	"net/http"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "time"
	// "encoding/json"
	// "reflect"
	"MyLargstProgram/Mytype"
     "github.com/json-iterator/go"
	// "strconv"
)
var json=jsoniter.ConfigCompatibleWithStandardLibrary
func ClearMessage(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
    session:=MgoInit.Session()
	defer session.Close()
	dbuser:=session.DB("NGDATA").C("user")
	// 是自己的id
	err:=dbuser.Update(bson.M{"userid":r.Form["openid"][0]},bson.M{"$set":bson.M{"message."+r.Form["findid"][0]+".counter":0}})
	if(err!=nil){
		fmt.Println(err)
	}else{
		w.Write([]byte("ok"))
	}
	
}
func Test(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	w.Write([]byte("hello android"))
}

func LoadUserinfo(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
	dbuser:=session.DB("NGDATA").C("user")
	var m=Mytype.User{}
	err:=dbuser.Find(bson.M{"userid":r.Form["findid"][0]}).Select(bson.M{"replay":0,"message":0,"merecorder":0,"messagerecorder":0}).One(&m)
	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println(m)
	data,_:=json.Marshal(m)
	w.Write(data)
}
func Pullreplay(w http.ResponseWriter,r *http.Request){
   r.ParseForm()
   session:=MgoInit.Session()
	defer session.Close()
	dbuser:=session.DB("NGDATA").C("user")
	err:=dbuser.Update(bson.M{"userid":r.Form["id"][0]},bson.M{"$pull":bson.M{"replay":bson.M{"retime":r.Form["time"][0]}}})
    if(err!=nil){
		fmt.Println(err)
	}
}
func ReadMessage(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	// fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
	dbuser:=session.DB("NGDATA").C("user")
	var m=Mytype.User{}
	// 应该带着自己deopenid过来
	err:=dbuser.Find(bson.M{"userid":r.Form["id"][0]}).Select(bson.M{"message":1,"replay":1}).One(&m)
	if(err!=nil){
		fmt.Println(err)
	}
	var alldata=make(map[string]interface{})
	alldata["message"]=m.Message
	alldata["replay"]=m.Replay
	data,_:=json.Marshal(alldata)
	// fmt.Println(string(data))
	fmt.Println(string(data))
	w.Write(data)
    fmt.Println("1")
}
func ReadOneMessage(w http.ResponseWriter,r *http.Request){
	r.ParseForm() 
	fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
    dbuser:=session.DB("NGDATA").C("user")
	var m=Mytype.User{}
	// 应该带着自己deopenid过来
	err:=dbuser.Find(bson.M{"userid":"ohWQQ5c6r9vt6aBHY4EdKEV5hrxI"}).Select(bson.M{"message":1}).One(&m)
	if(err!=nil){
		fmt.Println(err)
	}
	// fmt.Println(m.Message)
	data,_:=json.Marshal(m.Message)
	fmt.Println(string(data))
	w.Write(data)
}
func UnsetChat(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	session:=MgoInit.Session()  
	defer session.Close()  
	dbuser:=session.DB("NGDATA").C("user")
	dbuser.Update(bson.M{"userid":r.Form["mid"][0]},bson.M{"$unset":bson.M{"message."+r.Form["sid"][0]:0}})
    dbuser.Update(bson.M{"userid":r.Form["mid"][0]},bson.M{"$pull":bson.M{"messagerecorder":r.Form["sid"][0]}})
}
func WriteMessage(w http.ResponseWriter,r *http.Request){
    r.ParseForm()  
	fmt.Println(r.Form)
	session:=MgoInit.Session()  
	defer session.Close()  
	dbuser:=session.DB("NGDATA").C("user")
	var m Mytype.Message
	m.Sename=r.Form["Rename"][0]
	m.Rename=r.Form["Rename"][0]
	m.Reid=r.Form["Reid"][0]
	m.Seid=r.Form["Seid"][0]
	m.Stouxiang=r.Form["Retouxiang"][0]
	m.Mtext=r.Form["Mtext"][0]
	m.Mimg=r.Form["Mimg"][0]
	var u Mytype.Message
	u.Sename=r.Form["Sename"][0]
	u.Rename=r.Form["Sename"][0]
	u.Reid=r.Form["Reid"][0]
	u.Seid=r.Form["Seid"][0]
	u.Stouxiang=r.Form["Stouxiang"][0]
	u.Mtext=r.Form["Mtext"][0]
	u.Mimg=r.Form["Mimg"][0]
	// 分别去查看对方是否和自己已经有过私信记录
	var user Mytype.User
	err:=dbuser.Find(bson.M{"userid":m.Seid,"messagerecorder":bson.M{"$in":[]string{""+m.Reid}}}).One(&user)
	err1:=dbuser.Find(bson.M{"userid":m.Reid,"messagerecorder":bson.M{"$in":[]string{""+m.Seid}}}).One(&user)
	// 先去写自己的数据库
	//在chatroom的数据清洗中,对方需要知道我的名字,我需要知道对方的名字
	//所以在此定义sename是我的名字,rename是对方的名字
	if (err!=nil){
		// fmt.Println("lalla")
		errs:=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$set":bson.M{"message."+m.Reid:Mytype.MessManager{Counter:0,Messages:[]Mytype.Message{m}}}})
		// errs=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$inc":bson.M{"message."+m.Seid+".counter":1}})
		errs=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$push":bson.M{"messagerecorder":m.Reid}})
		if(errs!=nil){
			fmt.Println(errs)
		}
		w.Write([]byte("success"))
	}else{
		 err=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$push":bson.M{"message."+m.Reid+".messages":m}})
		//  err=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$inc":bson.M{"message."+m.Seid+".counter":1}})
		 if(err!=nil){
			 fmt.Println(err)
		 }
		 w.Write([]byte("success"))
	}
	if (err1!=nil){
		fmt.Println("lalla")
		errs:=dbuser.Update(bson.M{"userid":m.Reid},bson.M{"$set":bson.M{"message."+m.Seid:Mytype.MessManager{Counter:1,Messages:[]Mytype.Message{u}}}})
		// errs=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$inc":bson.M{"message."+m.Seid+".counter":1}})
		errs=dbuser.Update(bson.M{"userid":m.Reid},bson.M{"$push":bson.M{"messagerecorder":m.Seid}})
		if(errs!=nil){
			fmt.Println(errs)
		}
		w.Write([]byte("success"))
	}else{
		 err=dbuser.Update(bson.M{"userid":m.Reid},bson.M{"$push":bson.M{"message."+m.Seid+".messages":u}})
		 err=dbuser.Update(bson.M{"userid":m.Reid},bson.M{"$inc":bson.M{"message."+m.Seid+".counter":1}})
		 if(err!=nil){
			 fmt.Println(err)
		 }
		 w.Write([]byte("success"))
	}
	// err=dbuser.Find(bson.M{"userid":m.Seid}).Select(bson.M{"$exists":bson.M{"message."+m.Seid:true}})
	// if(err!=nil){
	// 	fmt.Println()
	// }
	
	// fmt.Pri// errs:=dbuser.Update(bson.M{"userid":m.Seid},bson.M{"$set":bson.M{"message."+m.Seid:Mytype.MessManager{Counter:1,Messages:[]Mytype.Message{m}}}})
	// if(errs!=nil){
	// 	fmt.Println(err)
	// }ntln("success")
}

func WriteReply (w http.ResponseWriter,r *http.Request){
    r.ParseForm()
	fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
	// db:=session.DB("NGDATA").C("user")
}
func ReadReply(w http.ResponseWriter,r *http.Request){
    r.ParseForm()
	fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
	// db:=session.DB("NGDATA").C("user")
}

