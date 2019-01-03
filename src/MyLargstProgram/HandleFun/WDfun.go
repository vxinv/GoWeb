package HandleFun
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "encoding/json"
	"MyLargstProgram/MgoInit"
	"fmt"
	"net/http"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"time"
	// "encoding/json"
	// "reflect"
	"MyLargstProgram/Mytype"
	"strconv"
)
// var json=jsoniter.ConfigCompatibleWithStandardLibrary
var Counter int
var Si = make(map[string]int)
var dbarray = []string{"jiaocai","tiaozao","kaoyan","shiwu","biaobai","guangbo","blog"}
// var countarray=[]int{}
// int(time.Now().UnixNano()>>8)
func In(i int){
    Counter=i
}
func Out() int{
	return Counter
}
func Loadone(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Form["type"][0])
	session:=MgoInit.Session()
	defer session.Close()
	db:=session.DB("NGDATA").C(r.Form["type"][0])
	var blog Mytype.Blogs
	monney,_:=strconv.Atoi(r.Form["Monney"][0])
	err:=db.Find(bson.M{"monney":monney}).One(&blog)
	data,_:=json.Marshal(blog)
	w.Write(data)
	m:=blog.Monney
	db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
	if(err!=nil){
		fmt.Println(err)
		fmt.Println("error")
	}
}
func notifyuser(){

}

func Uppinglun(w http.ResponseWriter,r *http.Request){
	timeUnix:=time.Now().Unix()    
	times:=strconv.Itoa(int(timeUnix))
	r.ParseForm()
	typename:=r.Form["type"][0]
	session:=MgoInit.Session()
	dbuser:=session.DB("NGDATA").C("user")
	// 找到该评论要插入到那个
	db:=session.DB("NGDATA").C(typename)
	// dbuser:=session.DB("NGDATA").C("user")
	defer session.Close()
	commentnumber,_:=strconv.Atoi(r.Form["commentnumber"][0])
	switch commentnumber{
	case 1:
	var comment Mytype.Comment
	comment.Cusername=r.Form["cusername"][0]
	comment.Cuserid=r.Form["cuserid"][0]
	comment.Comentext=r.Form["commenttext"][0]
	comment.Ctime=r.Form["ctime"][0]
	comment.Cusertouxiang=r.Form["ctouxiang"][0]
	comment.CComents=[]Mytype.CComent{}
	comment.Clistnumber,_=strconv.Atoi(r.Form["clist"][0])
	monney,_:=strconv.Atoi(r.Form["Monney"][0])
	var replay Mytype.Replay
	replay.Reid=r.Form["id"][0]
	replay.Remonney=monney
	replay.Rename=comment.Cusername
	replay.Retouxiang=comment.Cusertouxiang
	replay.Text=comment.Comentext
	replay.Retime=times
	replay.Type=typename
	err:=db.Update(bson.M{"monney":monney},bson.M{"$push":bson.M{"blogcomments":comment}})
	err=db.Update(bson.M{"monney":monney},bson.M{"$inc":bson.M{"blogcommentsnum":1}})
	err=dbuser.Update(bson.M{"userid":r.Form["id"][0]},bson.M{"$push":bson.M{"replay":replay}})
	// err=dbuser.Update(bson.M{"userid":r.Form["id"][0]},bson.M{"$pull":bson.M{"replay":bson.M{"retime":"1535974746"}}})
	if(err!=nil){
		fmt.Println(err)
	}
	case 2:
	fmt.Println("hello world")
	var ccomment Mytype.CComent
	ccomment.CCusername=r.Form["ccusername"][0]
	ccomment.CCuserid=r.Form["ccuserid"][0]
	ccomment.CComenttext=r.Form["ccomenttext"][0]
	cclist:=r.Form["cclist"][0]
	monney,_:=strconv.Atoi(r.Form["Monney"][0])
	fmt.Println(r.Form)
	var replay Mytype.Replay
		replay.Reid=r.Form["id"][0]
		replay.Remonney=monney
		replay.Rename=ccomment.CCusername
		replay.Retouxiang=r.Form["cctouxiang"][0]
		replay.Text=r.Form["ccomenttext"][0]
		replay.Retime=times
		replay.Type=typename
		err:=dbuser.Update(bson.M{"userid":r.Form["id"][0]},bson.M{"$push":bson.M{"replay":replay}})
	checkerr(err,1)
	err=db.Update(bson.M{"monney":monney},bson.M{"$push":bson.M{"blogcomments."+cclist+".ccoments":ccomment}})
	checkerr(err,2)
	err=db.Update(bson.M{"monney":monney},bson.M{"$inc":bson.M{"blogcommentsnum":1}})
	checkerr(err,3)
	if(err!=nil){
		fmt.Println(err)
	}
	w.Write([]byte("ok"))
	case 3:
		fmt.Println("hello 3")
		var cccoment Mytype.CCComent
		monney,_:=strconv.Atoi(r.Form["Monney"][0])
		cccoment.CCComenttext=r.Form["cccomenttext"][0]
		cccoment.CCCusername=r.Form["cccusername"][0]
		cccoment.CCCuserid=r.Form["cccuserid"][0]
		da:=r.Form["da"][0]
		xiao:=r.Form["xiao"][0]
		var replay Mytype.Replay
		replay.Reid=r.Form["id"][0]
		replay.Remonney=monney
		replay.Rename=cccoment.CCCusername
		replay.Retouxiang=r.Form["cccusertouxiang"][0]
		replay.Text=cccoment.CCComenttext
		replay.Retime=times
		replay.Type=typename
		fmt.Println(r.Form["id"][0])
		err:=dbuser.Update(bson.M{"userid":r.Form["id"][0]},bson.M{"$push":bson.M{"replay":replay}})
		err=db.Update(bson.M{"monney":monney},bson.M{"$push":bson.M{"blogcomments."+da+".ccoments."+xiao+".cccoments":cccoment}})
		err=db.Update(bson.M{"monney":monney},bson.M{"$inc":bson.M{"blogcommentsnum":1}})
		// err=dbuser.Update(bson.M{"userid":id},bson.M{"$push":bson.M{"usernewnotify":notify}})
		if(err!=nil){
			fmt.Println(err)
		}else{
			w.Write([]byte("ok"))
		}
	}
}

// 生成一个map来保存数句酷的值
func init(){
	session :=MgoInit.Session()
	defer session.Close()
	for _,v:= range dbarray{
	 db:= session.DB("NGDATA").C(v)
	 CounterCopy,err:=db.Count()
	 if err!=nil{
		fmt.Println(err)
		panic(err)
	  }
	//   countarray=append(countarray,CounterCopy)
	Si[v]=CounterCopy
	}
    fmt.Println(Si)
	// Counter=CounterCopy+1
	// fmt.Println(Counter)
}

func Loadmore(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
	fmt.Println(r.Form)
	// i:=getIndex(r.Form["type"][0],dbarray)
	var typename=r.Form["type"][0];
	var db *(mgo.Collection)
	if(typename=="hot"){
		db=session.DB("NGDATA").C("blog")
	}else{
		db=session.DB("NGDATA").C(typename)
	}
	var blogs []Mytype.Blogs
	var err error
	_,ok:=r.Form["id"]
	// 就不是个人动态
	if !ok {
		_,ok:=r.Form["Monney"]
		// 第一次查询
		if !ok{
			if r.Form["type"][0]=="hot"{
				//   fmt.Println(Si[typename])
				  err=db.Find(nil).Sort("-blogcommentsnum").Limit(10).All(&blogs)
				  fmt.Println("allll")
	              if(err!=nil){
		          fmt.Println(err)
	             }
			}else{
				 err=db.Find(bson.M{"monney":bson.M{"$lt":Si[typename]+1}}).Sort("-monney").Limit(10).All(&blogs)
	             if(err!=nil){
		         fmt.Println(err)
	             }
			}
	        data,_:=json.Marshal(blogs)
            fmt.Println(string(data))
            w.Write(data)
          	for _,v:= range blogs{
		     m:=v.Monney
	      	 db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
								}				   
		}else {
			if r.Form["type"][0]=="hot"{
				// Monney,_:=strconv.Atoi( r.Form["Monney"][0])
	    		err=db.Find(bson.M{"blogcommentsnum":bson.M{"$lt":r.Form["commentnumber"][0]}}).Sort("-blogcommentsnum").Limit(10).All(&blogs)
			}else{
				Monney,_:=strconv.Atoi( r.Form["Monney"][0])
	    		err=db.Find(bson.M{"monney":bson.M{"$lt":Monney}}).Sort("-monney").Limit(10).All(&blogs)
			}
	 		data,_:=json.Marshal(blogs)
       		 w.Write(data)
	   		for _,v:= range blogs{
			m:=v.Monney
			db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
	 							 }
			}
	}else{
       _,ok:=r.Form["Monney"]
		if !ok{
            err=db.Find(bson.M{"userid":r.Form["id"][0],"monney":bson.M{"$lt":Si[typename]+1}}).Sort("-monney").Limit(10).All(&blogs)
	        if(err!=nil){
		      fmt.Println(err)
	        }
	        data,_:=json.Marshal(blogs)
           	//  fmt.Println(string(data))
           w.Write(data)
          	for _,v:= range blogs{
		     m:=v.Monney
	      	db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
      	                         }
		}else {
		    Monney,_:=strconv.Atoi( r.Form["Monney"][0])
	        // fmt.Println(Monney)
	    	err=db.Find(bson.M{"userid":r.Form["id"][0],"monney":bson.M{"$lt":Monney}}).Sort("-monney").Limit(10).All(&blogs)
	 		data,_:=json.Marshal(blogs)
       		 w.Write(data)
	   		for _,v:= range blogs{
			m:=v.Monney
			db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
	 							 }
			}
	}
// 	if(r.Form["id"][0]=="undefined"){
// 	 if(r.Form["Monney"]==nil){
//     //  fmt.Println(Counter)		
// 	 err=db.Find(bson.M{"monney":bson.M{"$lte":Si[typename]}}).Sort("-monney").Limit(5).All(&blogs)
// 	 if(err!=nil){
// 		 fmt.Println(err)
// 	 }
// 	 data,_:=json.Marshal(blogs)
// 	//  fmt.Println(string(data))
//      w.Write(data)
// 	 for _,v:= range blogs{
// 		m:=v.Monney
// 		db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
// 	 }

// 	//  err=db.Update(bson.M{"monney":bson.M{"$lt":Counter}},bson.M{"$inc"})
// 	}else{
// 	 Monney,_:=strconv.Atoi( r.Form["Monney"][0])
// 	 fmt.Println(Monney)
// 	    err=db.Find(bson.M{"monney":bson.M{"$lte":Monney}}).Sort("-monney").Limit(5).All(&blogs)
// 	 	data,_:=json.Marshal(blogs)
//         w.Write(data)
// 	   for _,v:= range blogs{
// 		m:=v.Monney
// 		db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
// 	  }
// 	//  fmt.Println(blogs)
//    	}
// 	if err!=nil{
// 		fmt.Println("查询出错了")
// 		panic(err)
// 	}
//  }else{
// 	//  判断是否是个人动态的加载
//     if(r.Form["Monney"]==nil){
//     //  fmt.Println(Counter)		
// 	 err=db.Find(bson.M{"userid":r.Form["id"][0],"monney":bson.M{"$lte":Si[r.Form["type"][0]]}}).Sort("-monney").Limit(5).All(&blogs)
// 	 if(err!=nil){
// 		 fmt.Println(err)
// 	 }
// 	 data,_:=json.Marshal(blogs)
// 	//  fmt.Println(string(data))
//      w.Write(data)
// 	 for _,v:= range blogs{
// 		m:=v.Monney
// 		db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
// 	 }

// 	//  err=db.Update(bson.M{"monney":bson.M{"$lt":Counter}},bson.M{"$inc"})
// 	}else{
// 	 Monney,_:=strconv.Atoi( r.Form["Monney"][0])
// 	//  fmt.Println(Monney)
// 	    err=db.Find(bson.M{"userid":r.Form["id"][0],"monney":bson.M{"$lte":Monney}}).Sort("-monney").Limit(5).All(&blogs)
// 	 	data,_:=json.Marshal(blogs)
//         w.Write(data)
// 	   for _,v:= range blogs{
// 		m:=v.Monney
// 		db.Update(bson.M{"monney":m},bson.M{"$inc":bson.M{"view":1}})
// 	  }
// 	//  fmt.Println(blogs)
//    	}
// 	if err!=nil{
// 		fmt.Println("查询出错了")
// 		panic(err)
// 	}
//  }
	// var json=jsoniter.ConfigCompatibleWithStandardLibrary
}
func LoginLocalServer(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	session:=MgoInit.Session()
	defer session.Close()
	db:=session.DB("NGDATA").C("user")
	var user Mytype.User
	err:=db.Find(bson.M{"userid":r.Form["userid"][0]}).One(&user)
    if(err==nil){
        return
	}
	user.Usertouxiang=r.Form["usertouxiang"][0]
	user.Username=r.Form["username"][0]
	user.Userid=r.Form["userid"][0]
	user.Useradress=r.Form["useradress"][0]
	user.Userinfo=r.Form["userinfo"][0]
	user.Usermonney=[]string{}
	user.Message=make(map[string]Mytype.MessManager)
	user.Replay=[]Mytype.Replay{}
	user.MessageRecorder=[]string{}
	// user.Usernewnotify=[]Mytype.Notify{}
	err=db.Insert(&user)
	if(err!=nil){
		fmt.Println(err)	
	}
	fmt.Println("success")
	w.Write([]byte("ok"))
}
func Zan(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
    session:=MgoInit.Session()
	defer session.Close()
	db:=session.DB("NGDATA").C(r.Form["type"][0])
	int,_:=strconv.Atoi(r.Form["monney"][0])
	err:=db.Update(bson.M{"monney":int},bson.M{"$inc":bson.M{"zan":1}})
	if(err!=nil){
		fmt.Println(err)
	}
}
func Publish(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Counter)
	session:=MgoInit.Session()
	defer session.Close()
	r.ParseForm()
	dbname:=r.Form["typetype"][0]
	fmt.Println(dbname)
	// 到那个数据库
	db:=session.DB("NGDATA").C(dbname)
	switch r.Form["Type"][0]{
		// 只有文字
	case "1":
    // var dbi []string
	// json.Unmarshal([]byte(r.Form["Dbimagurls"][0]),&dbi)
	t := time.Now() 
	str_t := t.Format("01-02 15:04")
	err:=db.Insert(&Mytype.Blogs{
	  Username :r.Form["Username"][0],
	  Userid:r.Form["Openid"][0],
	  Type:1,
	  Userinfo:r.Form["Userinfo"][0],
	  Blogtext:r.Form["Usertext"][0],
	  Usertouxiang:r.Form["Usertouxiang"][0],
	  Blogimage:[]string{},
	  BlogCommentsnum:0,
	  Usersex:r.Form["Usersex"][0],
	  Usercredit:6000,
	  BlogComments:nil,
	  Time:str_t,
	  Monney: Si[dbname]+1,
	  Zan:0,
	  View:0,
	  Unique:0,
	})
	fmt.Println(Si[dbname])
	if err!=nil{
		panic(err)
	}
	// Counter+=1
	Si[dbname]+=1
// 带图片
	case "2":
    var dbi []string
	json.Unmarshal([]byte(r.Form["Dbimagurls"][0]),&dbi)
	t := time.Now() 
	str_t := t.Format("01-02 15:04")
	err:=db.Insert(&Mytype.Blogs{
	  Username :r.Form["Username"][0],
	  Userid:r.Form["Openid"][0],
	  Userinfo:r.Form["Userinfo"][0],
	  Blogtext:r.Form["Usertext"][0],
	  Usertouxiang:r.Form["Usertouxiang"][0],
	  Blogimage:dbi,
	  Type:2,
	  BlogCommentsnum:0,
	  Usersex:r.Form["Usersex"][0],
	  Usercredit:6000,
	  BlogComments:nil,
	  Time:str_t,
	  Monney: Si[dbname]+1,
	  Zan:0,
	  View:0,
	  Unique:0,
	})
	if err!=nil{
		panic(err)
	}
	Si[dbname]+=1
// video
	case "3":
    // var dbi []string
	// json.Unmarshal([]byte(r.Form["Dbimagurls"][0]),&dbi)
	t := time.Now() 
	str_t := t.Format("01-02 15:04")
	err:=db.Insert(&Mytype.Blogs{
	  Username :r.Form["Username"][0],
	  Userid:r.Form["Openid"][0],
	  Userinfo:r.Form["Userinfo"][0],
	  Blogtext:r.Form["Usertext"][0],
	  Usertouxiang:r.Form["Usertouxiang"][0],
	  Blogimage:[]string{},
	  BlogCommentsnum:0,
	  Usersex:r.Form["Usersex"][0],
	  Usercredit:6000,
	  VideoImage:r.Form["VideoImage"][0],
	  Video:r.Form["Video"][0],
	  Type:3,
	  BlogComments:nil,
	  Time:str_t,
	  Monney: Si[dbname]+1,
	  Zan:0,
	  View:0,
	  Unique:0,
	})
	if err!=nil{
		panic(err)
	}
	Si[dbname]+=1
	}
}
func checkerr(err error,flag int){
   fmt.Println(err)
   fmt.Println(flag)
}