package log

import (
    "fmt"
    "github.com/onsi/ginkgo"
    "github.com/zryfish/framework/framework/ginkgowrapper"
    "time"
)

func nowStamp() string {
    return time.Now().Format(time.StampMilli)
}

func log(level string, format string, args ...interface{}) {
    fmt.Fprintf(ginkgo.GinkgoWriter, nowStamp()+": "+level+": "+format+"\n", args...)
}

// Logf logs the info.
func Logf(format string, args ...interface{})  {
    log("INFO", format, args...)
}

// Failf logs the fail info.
func Failf(format string, args ...interface{}) {

}

// FailfWithOffset calls "Fail" and logs the error at "offset" levels above its caller
// (for example, for call chain f -> g -> FailWithOffset(1, ...) error would be logged for "f").
func FailfWithOffset(offset int, format string, args ...interface{}) {
    msg := fmt.Sprintf(format, args...)
    log("INFO", msg)
    ginkgowrapper.Fail(nowStamp()+": "+msg, 1+offset)
}