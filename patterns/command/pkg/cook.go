package pkg

// A Cook comes with their list of commands as attributes
type Cook struct {
	Name     string
	Commands []Command
}

// The executeCommands method executes all the commands
// one by one
func (c *Cook) ExecuteCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}
