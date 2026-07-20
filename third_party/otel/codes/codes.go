package codes

type Code uint32

const (
	Unset Code = iota
	Error
	Ok
)

var codeStrings = map[Code]string{
	Unset: "Unset",
	Error: "Error",
	Ok:    "Ok",
}

func (c Code) String() string {
	return codeStrings[c]
}
