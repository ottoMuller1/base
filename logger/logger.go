package logger

import (
	"fmt"
	"log"
	"os"
	"github.com/ottoMuller1/base/nullable"
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




// log types
var NullPrefixHandler = nullable.Null[func (string) string]()

// logger type
type logger struct {
	color  string
	prefixHandler func (string) string
}

// new logger
func New(rgb RGB, prefixHandlerNullable nullable.Nullable[func (string) string]) logger {

	return logger {
		color: rgbToAnsi(rgb),
		prefixHandler: prefixHandlerNullable.FromNullable(
			func(s string) string {return s},
			true,
		),
	}

}




// basic loger terms
var Error = logger {
	color: "\033[31m",
	prefixHandler: func(message string) string {return "ERROR: " + message + " "},
}

var Warning = logger {
	color: "\033[33m",
	prefixHandler: func(message string) string {return "WARNING: " + message + " "},
}

var Info = logger {
	color: "\033[32m",
	prefixHandler: func(message string) string {return "INFO: " + message + " "},
}

var Custom = logger {
	color: "\033[36m",
	prefixHandler: func(message string) string {
		if message != "" {
			return message + " "
		}
		return "LOG: "
	},
}




// log function
func (lg logger) Log(message string, prefix string) {
	logger := log.New(os.Stdout, makeTextColorized(lg.prefixHandler(prefix), lg.color), log.Ldate|log.Ltime)
	logger.Printf(makeTextColorized(message, lg.color))
}
