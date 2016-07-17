package greeting

func Greeting(args []string) string {
	return "Hello" + getName(args)
}

func getName(args []string) string {
	if len(args) == 0 {
		return " World"
	}
	var name string
	for _, n := range args {
		name = name + " " + n
	}
	return name
}

