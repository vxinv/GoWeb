package Mytype

type Blogs struct { 
	Username     string   `json:"username" bson:"username"`
	Userid       string   ` json:"openid" bson:"userid"`
	Unique       int      `json:"unique" bson:"unique"`
	Usertouxiang string   `json:"usertouxiang" bson:"usertouxiang"`
	Usersex      string   `json:"usersex" bson:"sex"`
	Userinfo     string   `json:"userinfo" bson:"userinfo"`
	Blogtext     string   `json:"usertext" bson:"blogtext"`
	Blogimage    []string `json:"blogimages" bson:"blogimages"`
	BlogComments []Comment `json:"blogcomments" bson:"blogcomments"`
	BlogCommentsnum int `json:"blogcommentsnum"`
	Time string `json:"publishtime" bson:"publishtime"`
	Type int  `json:"type" bson:"type"`
	VideoImage string 
	Video string
	Usercredit int 
	Monney int
	View int
	Zan int
}
type Comment struct{
	 Clistnumber int   `json:"listnumber"`
	 Cuserid string     `json:"cuserid"`
	 Cusername  string `json:"cusername"`
	 Cusertouxiang string `json:"cusertouxiang"`
	 Comentext string  `json:"comenttext"`
	 Ctime string  `json:"time"`
	 CComents []CComent `json:"ccoments"`
 }
type CComent struct{
	CCComents []CCComent `json:"cccoments"`
	CCuserid string  `json:"ccuserid"`
	CCusername string `json:"ccusername"`
	CComenttext string `json:"ccomenttext"`

}
type CCComent struct{
	CCCuserid string  `json:"cccuserid"`
	CCCusername string `json:"cccusername"`
	CCComenttext string `json:"cccomenttext"`
}
type User struct{
	Usertouxiang string `json:"usertouxiang"`
	Username string    `json:"username"`
	Userid string       `json:"userid"`
	Userinfo string      `json:"userUserinfoname"`
	Useradress string       `json:"useradress"`
	Usermonney []string    `json:"usermonney"`
	Message map[string]MessManager   `json:"message"  bson:"message"`
	Replay  []Replay  `json:"replay" bson:"replay" `  
	MessageRecorder []string   `json:"merecorder" bson:"merecorder"`        
	Monneys  []string `json:"monneys" bson:"monneys"`    
}
type MessManager struct {
   Counter int   `json:"counter" bson:"counter"`  
   Messages []Message `json:"messages" bson:"messages" `  
}

type Replay struct{
   Retime string `jsin:"retime" bson:"retime"`	
   Rename string    `json:"rename" bson:"rename"`  
   Reid string 	`json:"reid" bson:"reid"`  
   Text string	`json:"text" bson:"text"`  
   Retouxiang string	`json:"retouxiang" bson:"retouxiang"`  
   Remonney int	`json:"remonney" bson:"remonney"`  
   Type string `json:"type" bson:"type"`
}
type ReplayManager struct{
	Replays []Replay  `json:"replays" bson:"replays"`  
    Counter  int     `json:"counter" bson:"counter"`  
}

type Message struct{
	Sename string    `json:"sename" bson:"sename" `
	Rename string    `json:"rename" bson:"rename" `
	Seid string      `json:"seid" bson:"seid" `
	Reid string      `json:"reid" bson:"reid" `
	Mtext string	 `json:"mtext" bson:"mtext" `
	Mimg string	     `json:"mimg" bson:"mimg" `
	Stouxiang string `json:"stouxiang" bson:"stouxiang" `
}



