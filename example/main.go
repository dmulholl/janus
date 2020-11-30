package main

import (
    "fmt"
    "github.com/dmulholl/janus/v2"
)

// This sample application will parse its own command-line arguments.
func main() {

    // We instantiate an argument parser, then specify help text and a version
    // string. Specifying help text activates an automatic --help flag;
    // specifying a version string activates an automatic --version flag.
    parser := janus.NewParser()
    parser.Helptext = "App Help"
    parser.Version = "1.2.3"

    // Register a flag, --bool, with a single-character alias, -b. A flag is a
    // boolean option - it's either present (true) or absent (false).
    parser.NewFlag("bool b")

    // Register a string option, --string <arg>, with a single-character
    // alias, -s <arg>. The default fallback value for string options is the
    // empty string. Here we specify a custom fallback value, "foobar".
    parser.NewString("string s", "foobar")

    // Register an integer option, --int <arg>, with a single-character alias,
    // -i <arg>. The default fallback value for integer options is 0.
    parser.NewInt("int i")

    // Register a floating-point option, --float <arg>, with a single-character
    // alias, -f <arg>. The default fallback value for floating-point options
    // is 0.
    parser.NewFloat("float f")

    // Register a command 'foo'. We need to supply the command's help text and
    // callback method.
    cmdParser := parser.NewCmd("foo", "Command Help", callback)

    // Registering a command returns a new ArgParser instance dedicated to
    // parsing the command's arguments. We can register as many flags and
    // options as we like on this sub-parser. Note that the sub-parser can
    // reuse the parent's option names without interference.
    cmdParser.NewFlag("bool b")
    cmdParser.NewInt("int i", 123)

    // Once all our options and commands have been registered we can call the
    // parser's Parse() method to parse the application's command line
    // arguments. Only the root parser's Parse() method should be called -
    // command arguments will be parsed automatically.
    parser.Parse()

    // We can now retrieve our option and argument values from the parser
    // instance. Here we simply dump the parser to stdout.
    fmt.Println(parser)
}

// Callback function for the 'foo' command. The function receives an ArgParser
// instance containing the command's parsed arguments. Here we simply dump it
// to stdout.
func callback(parser *janus.ArgParser) {
    fmt.Println("---------- callback ----------")
    fmt.Println(parser)
    fmt.Println("------------------------------\n")
}
