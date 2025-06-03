package grammar

// These keywords are used by compiler
var ReservedSymbols = map[string][]string{
	"identifier": {
		"ident", "i",
	},
	"string": {
		"str", "s",
	},
}
