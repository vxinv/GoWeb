package main
import (
	"fmt"
	"sync"
)
/*单例模式*/
var m *Manager
var lock *sync.Mutex= &sync.Mutex{}
type Manager struct{
   num int
}
func(m *Manager)add(){
  m.num+=1
}
func getInstanceM() *Manager {
	if m==nil{     //带有双重锁 避免单例失败
		lock.Lock()  
		defer lock.Unlock()
        if m==nil{
			m= &Manager{}
		}
	}
	return m
}
var once sync.Once
func getInstance1() *Manager{
     once.Do(func(){
		 m=&Manager{}
	 })
	 return m
}
func(m *Manager) show()int{
	return m.num
}
func main(){
	 m:=getInstanceM();
	 for i:=0;i<100;i++{
		go func(m *Manager){
			 m.add()
			 fmt.Println(m.show())
	      }(m) 
	 }
	 for{
		 if(m.show()==99){
			fmt.Println(m.show())
            break
		 }
	 }
}
















/*代码更简单了，而且有没有发现-漂亮了！Once.Do方法的参数是一个函数
，这里我们给的是一个匿名函数，在这个函数中我们做的工作很简单，就是去赋值m变量，
而且go能保证这个函数中的代码仅仅执行一次！
ok，到现在单例模式我们就介绍完了，
内容并不多，因为单例模式太简单而且太常见了。
我们用单例的目的是为了保证在整个系统中存在唯一的实例，
我们加锁的目的是为了在并发的环境中单例依旧好用。不过虽然单例简单，我们还是不能任性的用，
因为这样做实例会一直存在内存中，一些我们用的不是那么频繁的东西使用了单例是不是就造成了内存的浪费？
大家在用单例的时候还是要多思考思考，这个模块适不适合用单例！*/