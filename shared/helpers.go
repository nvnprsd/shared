package shared

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/nvnprsd/base"
	"github.com/nvnprsd/shared/entities"
)

func SendZippedResponse(c *gin.Context, name string) {
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename="+name+".zip")
	c.Writer.Flush()

}
func ConvertToFile(jsonData []byte) *[]entities.File {
	var jsonMap entities.RequestFiles
	json.Unmarshal(jsonData, &jsonMap)
	return &jsonMap.Files
}
func CreateTempFolder(path string) func() error {

	err := os.MkdirAll(path, 0755)
	if err != nil {
		//base.Log.Error("Failed to create temp folder" + err.Error())
	}
	Flush := func() error {
		err = os.RemoveAll(path)
		if err != nil {
			return base.Log.Debug("failed to remove temp folder" + err.Error())
		}
		return nil
	}
	return Flush
}
func PrintStacktrace() {
	var buff [32]uintptr
	n := runtime.Callers(3, buff[:])
	frames := runtime.CallersFrames(buff[:n])

	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}
