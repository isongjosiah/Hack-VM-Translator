package codewriter

var segments = map[string]string{
	"local":    "@LCL",
	"stack":    "@SP",
	"argument": "@ARG",
	"this":     "@THIS",
	"that":     "@THAT",
	"static":   "@16",
	"pointer":  "@THIS",
}

// CodeWriter translates the VM commands
// into Hack assembly code
type CodeWriter struct {
}

// Writepush ...
func (c *CodeWriter) Writepush(filename string, segment string, integer int) string {
	switch segments[segment] {
	case "local":

	}
	return "nil"
}

// Writepop ...
func (c *CodeWriter) Writepop(filename string) {

}

// Doarithmetic ...
func (c *CodeWriter) Doarithmetic(filename string) {

	// func (c *CodeWriter) Dogt() {

	// }

	// func (c *CodeWriter) Dolt() {

	// }

	// func (c *CodeWriter) Doeq() {

	// }

	// func (c *CodeWriter) Doadd() {

	// }

	// func (c *CodeWriter) Dosub() {

	// }

	// func (c *CodeWriter) Doneg() {

	// }

	// func (c *CodeWriter) Donot() {

	// }

	// func (c *CodeWriter) Doand() {

	// }

	// func (c *CodeWriter) Door() {

	// }
}

// Writelabel ...
func (c *CodeWriter) Writelabel(filename string) {

}

// Dogoto ...
func (c *CodeWriter) Dogoto(filename string) {

}

// Doifgoto ...
func (c *CodeWriter) Doifgoto(filename string) {

}

// Dofunction ...
func (c *CodeWriter) Dofunction(filename string) {

}

// Docall ...
func (c *CodeWriter) Docall(filename string) {

}

// Doreturn ...
func (c *CodeWriter) Doreturn(filename string) {

}
