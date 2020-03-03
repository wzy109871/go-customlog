# go-customlog使用说明
##输出到控制台
      
      package main
      import (
      	"goFive/day06/customLog"
      	"time"
      )
      
      func main() {
      	log := customLog.NewLog("info")
      	//模拟debug错误
      	//模拟info错误
      	//模拟waring错误
      	//模拟error错误
      	//模拟fata错误
      	for {
      		debugError := "这是debugError!"
      		log.Debug(debugError)
      		time.Sleep(time.Second*2)
      		infoError := "这是infoError!"
      		log.Info(infoError)
      		time.Sleep(time.Second*2)
      		waringError := "这是waringError!"
      		log.Waring(waringError)
      		time.Sleep(time.Second*2)
            //errorError := "这是errorError!"
		    id := 100
		    name := "test error log"
		    log.Error("这是一条error日志 %d %s", id, name)
      		time.Sleep(time.Second*2)
      		fataError := "这是fataError!"
      		log.Fata(fataError)
      		time.Sleep(time.Second*2)
      	}
      }
##输出到文件
~~~~
package main
import (
	"goFive/go-customlog"
	"time"
)

func main() {
	log := customLog.NewFileLog("info","./log","wzylog.log")
	//模拟debug错误
	//模拟info错误
	//模拟waring错误
	//模拟error错误
	//模拟fata错误
	for {
		debugError := "这是debugError!"
		log.FileDebug(debugError)
		time.Sleep(time.Second*2)
		infoError := "这是infoError!"
		log.FileInfo(infoError)
		time.Sleep(time.Second*2)
		waringError := "这是waringError!"
		log.FileWaring(waringError)
		time.Sleep(time.Second*2)
		//errorError := "这是errorError!"
		id := 100
		name := "test error log"
		log.FileError("这是一条error日志 %d %s", id, name)
		time.Sleep(time.Second*2)
		fataError := "这是fataError!"
		log.FileFata(fataError)
		time.Sleep(time.Second*2)
	}
}