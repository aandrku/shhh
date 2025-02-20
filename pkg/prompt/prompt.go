package prompt

import (
	"shhh/pkg/style"
	"strings"
)

// A Get function return a shell prompt as configured by user
func Get() string {
	builder := strings.Builder{}

	builder.WriteString(style.MAGENTA)
	builder.WriteString("aandrku's $:  ")
	builder.WriteString(style.RESET)

	return builder.String()
}

// Get2 function returns a prompt which is displayed when user types \
// for line continuation.
func Get2() string {
	return style.CYAN + ">  " + style.RESET
}
