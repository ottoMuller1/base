package logger

import (
	"log"
	"os"
	"github.com/ottoMuller1/base/nullable"
)


// colors
func makeTextColorized(message string, color string) string {
	return color + message + "\033[0m"
}




// log types
type PrefixHandler func(string) string

type logger struct {
	color  string
	prefixHandler PrefixHandler
}

func New(color string, prefixHandlerNullable nullable.Nullable[PrefixHandler]) logger {

	defaultPrefixHandler := func(s string) string {
		return s
	}

	return logger {
		color: color,
		prefixHandler: prefixHandlerNullable.FromNullable(defaultPrefixHandler, true),
	}

}





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




// logger fuction
func (lg logger) Log(message string, prefix string) {
	logger := log.New(os.Stdout, makeTextColorized(lg.prefixHandler(prefix), lg.color), log.Ldate|log.Ltime)
	logger.Printf(makeTextColorized(message, lg.color))
}
