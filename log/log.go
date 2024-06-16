package log


import (
		"github.com/zerodha/logf"
		"time"
		"os"
	)


var Log logf.Logger

func init() {
	os.Mkdir("./log",0777)
	file,err:=os.OpenFile("./log/log.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0777);
	if err!=nil{panic("Unable to create log file")}
	
	Log = logf.New(logf.Opts{
		Writer:               file,
		EnableColor:          false,
		Level:                logf.DebugLevel,
		CallerSkipFrameCount: 3,
		EnableCaller:         true,
		TimestampFormat:      time.RFC3339Nano,
	})
		
}


