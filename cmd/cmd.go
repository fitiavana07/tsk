// cmd package, for cmd related functions
package cmd

// CLIParam represents tsk CLI params, composed of action and the main arg
type Param struct {
	Action string
	Arg    string
}

// ParseArgs parses args array returned by flag.Args(),
// and returns a Param value
func ParseArgs(args []string) Param {
	switch len(args) {
	case 0:
		return Param{"", ""}
	case 1:
		return Param{args[0], ""}
	case 2:
		return Param{args[0], args[1]}
	default:
		return Param{"", ""}
	}
}
