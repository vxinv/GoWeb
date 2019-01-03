package Test

func selectmain(){

	select{
		case <-ch1
           //如果ch1成功读到数据.就会执行
		case ch2<-1
		//如果成功向ch2写入数据,就会执行
	}
}
// 与switch 可使用任何相等的条件相比,select 的case 语句必须是一个IO操作
//当任意一个case 没有满足条件,select 会阻塞
//