package main

import (
  "gopkg.in/ukautz/clif.v1"
  "example.org/auth"
  "fmt"
)


func setStyle(style string, c *clif.Cli) {
	
    for key, _ := range clif.DebugStyles {
      clif.DebugStyles[key] = ""
    }
    clif.DefaultStyles = clif.DebugStyles

    auth.OpenLogin()
}

func deploy(c *clif.Command) {
	fmt.Println("you're in the deploy func")
	name := c.Argument("name").String()
		fmt.Println("my name is ", name)
}

func cb(c *clif.Command, out clif.Output) {
	out.Printf("Called %s\n", c.Name)
	name := c.Argument("name").String()

	if name == "" {
		fmt.Println("i have no name")
	} else {
		fmt.Println("my name is ", name)
	}
	
}

func main() {
  setStyle("debug", nil)
  
    // Create command with callback using the peviously registered instance
    // cli.NewCommand("auth", "Call foo", auth.OpenLogin)

    // build & add a complex command
	// initialize the app with custom registered objects in the injection container
	c := clif.New("Placeholder name", "1.0.0", "CLI for Placeholder name")

	// customize error handler
	clif.Die = func(msg string, args ...interface{}) {
		c.Output().Printf("<error>Everyting went wrong: %s<reset>\n\n", fmt.Sprintf(msg, args...))
		clif.Exit(1)
	}

	cmd := clif.NewCommand("deploy", "A description", cb).
	// Follows the format of name/description/default string/required (true/false)/multiple options (true/false)
	NewArgument("name", "Name for greeting", "", false, false)

	c.Add(cmd)
	// build & add a complex command

	cmd = clif.NewCommand("auth", "It does foo", auth.OpenLogin)
	c.Add(cmd)
	c.Run()
  
}