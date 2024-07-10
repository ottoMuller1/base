package logger

import (
	"fmt"
	"log"
	"os"
)

// colors
func makeTextColorized(message string, color string) string {
	return color + message + "\033[0m"
}

// rgb type
type RGB struct {
	R int
	G int
	B int
}

// rgb to ansi value
func rgbToAnsi(rgb RGB) string {

	return fmt.Sprintf("\u001b[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B)

}




// logger type
type logger struct {
	color  string
	tagHandler func (string) string
}

// new logger
func New(rgb RGB) logger {

	return logger {
		color: rgbToAnsi(rgb),
		tagHandler: func(s string) string {
			tag := "LOG:"

			if s != "" {
				tag = s
			}

			return tag
		},
	}

}




// basic loger terms
var Error = logger {
	color: rgbToAnsi(RGB {255, 0, 0}),
	tagHandler: func(message string) string {return "ERROR: " + message},
}

var Warning = logger {
	color: rgbToAnsi(RGB {255, 255, 0}),
	tagHandler: func(message string) string {return "WARNING: " + message},
}

var Info = logger {
	color: rgbToAnsi(RGB{0, 255, 0}),
	tagHandler: func(message string) string {return "INFO: " + message},
}




// log function
func (lg logger) Log(tag string, message string) {
	logger := log.New(os.Stdout, makeTextColorized(lg.tagHandler(tag) + " ", lg.color), log.Ldate|log.Ltime)
	logger.Printf(makeTextColorized(message, lg.color))
}
