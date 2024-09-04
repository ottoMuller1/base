package logger



import (
	nl "github.com/ottoMuller1/base/nullable"
)



type DebugInfo struct {
	Tag nl.Nullable[string]    // if we want to mark the exact place at our code
}

type Logger interface {
	Debug(DebugInfo)
	Error()
	Warning()
	Info()
}




// default logger
type DefaultLogger struct {
	Name string
	Message string
}

func (d DefaultLogger) Debug(di DebugInfo) {

	tag := di.Tag.PassError(nil).FromNullable("")

	New(RGB{R: 255, G: 51, B: 255}).Log("DEBUG: " + d.Name, d.Message + " " + tag)

}

func (d DefaultLogger) Error() {

	Error.Log(d.Name, d.Message)

}

func (d DefaultLogger) Info() {

	Info.Log(d.Name, d.Message)

}

func (d DefaultLogger) Warning() {

	Warning.Log(d.Name, d.Message)

}

