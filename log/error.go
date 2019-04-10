package log

import (
	"fmt"
	"github.com/golang/glog"
	"runtime"
)


/*
	Error logs first the function name that calling log.Error and the error message.
*/
func Error(err error) {
	fmt.Println("------------")

	skipFrames := 2

	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	glog.Errorln(frame.Function)
	glog.Errorln(err)
	fmt.Println("------------")
}
