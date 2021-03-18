package util

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

const (
	LevelDebug = (iota + 1) * 100 //100
	LevelInfo //200
	LevelWarning //300
	LevelDatalog //400
	LevelError //500
	LevelReport //600
)

var (
	levels = map[int]string{
		LevelDebug:   "DEBUG",
		LevelInfo:    "INFO",
		LevelWarning: "WARNING",
		LevelDatalog: "DATALOG",
		LevelError:   "ERROR",
		LevelReport:  "REPORT",
	}
)
var loglevel int

func Setloglevel(level string) {
	for k, v := range levels {
		if v == strings.ToUpper(level) {
			loglevel = k
		}
	}
}
func init() {
	log.SetFlags(log.LstdFlags)
	//log.SetPrefix("[security-url]\t")
}

func logFile() (sb string) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		file_arr := strings.Split(file, "/")
		sb = file_arr[len(file_arr)-1] + ":" + fmt.Sprintf("%d", line)
	} else {
		sb = ""
	}
	return
}

func Debugln(v ...interface{}) {
	if LevelDebug < loglevel {
		return
	}
	s := fmt.Sprint(v...)
	log.Println(logFile()+" DEBUG\t", s)
}

func Debugf(format string, v ...interface{}) {
	if LevelDebug < loglevel {
		return
	}
	format = logFile() + " DEBUG\t" + format
	log.Printf(format, v...)
}

func Infoln(v ...interface{}) {
	if LevelInfo < loglevel {
		return
	}
	s := fmt.Sprint(v...)
	log.Println(logFile()+" INFO\t", s)

}

func Infof(format string, v ...interface{}) {
	if LevelInfo < loglevel {
		return
	}
	format = logFile() + " INFO\t" + format
	log.Printf(format, v...)
}

func Warningln(v ...interface{}) {
	if LevelWarning < loglevel {
		return
	}
	s := fmt.Sprint(v...)
	log.Println(logFile()+" WARNING\t", s)
}

func Warningf(format string, v ...interface{}) {
	if LevelWarning < loglevel {
		return
	}
	format = logFile() + " WARNING\t" + format
	log.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	if LevelError < loglevel {
		return
	}
	s := fmt.Sprint(v...)
	log.Println(logFile()+" ERROR\t", s)
}

func Errorf(format string, v ...interface{}) {
	if LevelError < loglevel {
		return
	}
	format = logFile() + " ERROR\t" + format
	log.Printf(format, v...)
}

func DataLogln(v ...interface{}) {
	s := fmt.Sprint(v...)
	log.Println(logFile()+" DATALOG\t", s)
}

func DataLogf(format string, v ...interface{}) {
	format = logFile() + " DATALOG\t" + format
	log.Printf(format, v...)
}

func ReportLogln(v ...interface{}) {
	s := fmt.Sprint(v...)
	log.Println(logFile()+" REPORTLOG\t", s)
}

func ReportLogf(format string, v ...interface{}) {
	format = logFile() + " REPORTLOG\t" + format
	log.Printf(format, v...)
}

func PanicPosition() string {
	_, file, line, _ := runtime.Caller(1)
	li := strings.Split(file, "/")
	if len(li) >= 3 {
		file = strings.Join(li[len(li)-3:], "/")
	}
	s := fmt.Sprintf("%s +%d: ", file, line)
	return s
}
func PanicRecover(pos string) {
	if err := recover(); err != nil {
		log.Printf(logFile() + " PANIC! Pay Attention at " + pos + fmt.Sprintf(" err %v", err))
		SendAlarm(" PANIC! Pay Attention at "+pos, " err ", err)
	}
}
