// Package greeting says hello world
package greeting

const testVersion = 3

// HelloWorld returns a greeting to the user or to the World if name is blank
func HelloWorld(name string) string {
	if name == "" {
		return "Hello, World!"
	} else {
		return "Hello, " + name + "!"
	}
}
