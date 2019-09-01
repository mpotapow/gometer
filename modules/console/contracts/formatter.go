package contracts

// Formatter ...
type Formatter interface {
	NewLine()
	Text(text string)
	Write(text string)
	Writeln(text string)
}
