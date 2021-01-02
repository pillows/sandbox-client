package main

import (
  "gopkg.in/ukautz/clif.v1"
  "example.org/auth"
  "fmt"
  "reflect"
)

type exampleInterface interface {
	Foo() string
}

type exampleStruct struct {
	foo string
}

func (this *exampleStruct) Foo() string {
	return this.foo
}

func callHello(out clif.Output) {
	out.Printf("Hello <mine>World<reset>\n")
}

func callStyles(out clif.Output) {
	for token, _ := range clif.DefaultStyles {
		if token == "mine" {
			continue
		}
		out.Printf("Token \\<%s>: <%s>%s<reset>\n", token, token, token)
	}
}

func callFoo(c *clif.Command, out clif.Output, custom1 exampleInterface, custom2 *exampleStruct) {
	out.Printf("Hello %s, how is the\n", c.Argument("name").String(), )
	// if m := c.Argument("more-names").Strings(); m != nil && len(m) > 0 {
	// 	for _, n := range m {
	// 		out.Printf("  Say hello to <info>%s<reset>\n", n)
	// 	}
	// }
	// if c.Option("counter").Int() > 5 {
	// 	out.Printf("  You can count real high!\n")
	// }
	// out.Printf("  <headline>Custom 1: %s<reset>\n", custom1.Foo())
	// out.Printf("  <subline>Custom 2: %s<reset>\n", custom2.foo)
}

func setStyle(style string, c *clif.Cli) {
	
    for key, _ := range clif.DebugStyles {
      clif.DebugStyles[key] = ""
    }
    clif.DefaultStyles = clif.DebugStyles

    auth.OpenLogin()
}

func main() {
  setStyle("debug", nil)
  
    // Create command with callback using the peviously registered instance
    // cli.NewCommand("auth", "Call foo", auth.OpenLogin)

    // build & add a complex command
	// initialize the app with custom registered objects in the injection container
	c := clif.New("Placeholder name", "1.0.0", "CLI for Placeholder name").
		Register(&exampleStruct{"bar1"}).
		RegisterAs(reflect.TypeOf((*exampleInterface)(nil)).Elem().String(), &exampleStruct{"bar2"}).
		New("hello", "The obligatory hello world", callHello)

	styleArg := clif.NewArgument("style", "Name of a style. Available: default, sunburn, winter", "default", true, false).
		SetParse(func(name, value string) (string, error) { setStyle(value, c); return value, nil })
	c.Add(clif.NewCommand("styles", "Print all color style tokens", callStyles).AddArgument(styleArg))

	// customize error handler
	clif.Die = func(msg string, args ...interface{}) {
		c.Output().Printf("<error>Everyting went wrong: %s<reset>\n\n", fmt.Sprintf(msg, args...))
		clif.Exit(1)
	}

	// build & add a complex command
	cmd := clif.NewCommand("foo", "It does foo", callFoo).
		NewArgument("name", "Name for greeting", "", true, false).
		NewArgument("more-names", "And more names for greeting", "", false, true)

	c.Add(cmd)

	cb := func(c *clif.Command, out clif.Output) {
		out.Printf("Called %s\n", c.Name)
	}
	c.New("bar:baz", "A grouped command", cb).
		New("bar:zoing", "Another grouped command", cb).
		New("hmm:huh", "Yet another grouped command", cb).
		New("hmm:uhm", "And yet another grouped command", cb)

	// execute the main loop
	c.Run()
  
}