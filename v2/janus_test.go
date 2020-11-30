package janus


import (
    "testing"
)


// -----------------------------------------------------------------------------
// Boolean options.
// -----------------------------------------------------------------------------


func TestBoolOptionEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool")
    parser.ParseArgs([]string{})
    if parser.GetFlag("bool") != false {
        t.Fail()
    }
}


func TestBoolOptionMissing(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool")
    parser.ParseArgs([]string{"foo", "bar"})
    if parser.GetFlag("bool") != false {
        t.Fail()
    }
}


func TestBoolOptionLongform(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool")
    parser.ParseArgs([]string{"--bool"})
    if parser.GetFlag("bool") != true {
        t.Fail()
    }
}


func TestBoolOptionShortform(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool b")
    parser.ParseArgs([]string{"-b"})
    if parser.GetFlag("bool") != true {
        t.Fail()
    }
}


func TestBoolListEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool")
    parser.ParseArgs([]string{})
    if parser.LenList("bool") != 0 {
        t.Fail()
    }
}


func TestBoolListLongform(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool")
    parser.ParseArgs([]string{"--bool", "--bool", "--bool"})
    if parser.LenList("bool") != 3 {
        t.Fail()
    }
}


func TestBoolListShortform(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool b")
    parser.ParseArgs([]string{"-b", "-b", "-b"})
    if parser.LenList("bool") != 3 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// String options.
// -----------------------------------------------------------------------------


func TestStringOptionEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewString("string", "default")
    parser.ParseArgs([]string{})
    if parser.GetString("string") != "default" {
        t.Fail()
    }
}


func TestStringOptionMissing(t *testing.T) {
    parser := NewParser()
    parser.NewString("string", "default")
    parser.ParseArgs([]string{"foo", "bar"})
    if parser.GetString("string") != "default" {
        t.Fail()
    }
}


func TestStringOptionLongform(t *testing.T) {
    parser := NewParser()
    parser.NewString("string", "default")
    parser.ParseArgs([]string{"--string", "value"})
    if parser.GetString("string") != "value" {
        t.Fail()
    }
}


func TestStringOptionShortform(t *testing.T) {
    parser := NewParser()
    parser.NewString("string s", "default")
    parser.ParseArgs([]string{"-s", "value"})
    if parser.GetString("string") != "value" {
        t.Fail()
    }
}


func TestStringListEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewString("str")
    parser.ParseArgs([]string{})
    if parser.LenList("str") != 0 {
        t.Fail()
    }
}


func TestStringListLongform(t *testing.T) {
    parser := NewParser()
    parser.NewString("str")
    parser.ParseArgs([]string{"--str", "a", "b", "--str", "c"})
    if parser.LenList("str") != 2 {
        t.Fail()
    }
    if parser.GetStringList("str")[0] != "a" {
        t.Fail()
    }
    if parser.GetStringList("str")[1] != "c" {
        t.Fail()
    }
}


func TestStringListShortform(t *testing.T) {
    parser := NewParser()
    parser.NewString("str s")
    parser.ParseArgs([]string{"-s", "a", "b", "-s", "c"})
    if parser.LenList("str") != 2 {
        t.Fail()
    }
    if parser.GetStringList("str")[0] != "a" {
        t.Fail()
    }
    if parser.GetStringList("str")[1] != "c" {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Integer options.
// -----------------------------------------------------------------------------


func TestIntOptionEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int", 101)
    parser.ParseArgs([]string{})
    if parser.GetInt("int") != 101 {
        t.Fail()
    }
}


func TestIntOptionMissing(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int", 101)
    parser.ParseArgs([]string{"foo", "bar"})
    if parser.GetInt("int") != 101 {
        t.Fail()
    }
}


func TestIntOptionLongform(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int", 101)
    parser.ParseArgs([]string{"--int", "202"})
    if parser.GetInt("int") != 202 {
        t.Fail()
    }
}


func TestIntOptionShortform(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int i", 101)
    parser.ParseArgs([]string{"-i", "202"})
    if parser.GetInt("int") != 202 {
        t.Fail()
    }
}


func TestIntOptionNegative(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int", 101)
    parser.ParseArgs([]string{"--int", "-202"})
    if parser.GetInt("int") != -202 {
        t.Fail()
    }
}


func TestIntListEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int")
    parser.ParseArgs([]string{})
    if parser.LenList("int") != 0 {
        t.Fail()
    }
}


func TestIntListLongform(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int")
    parser.ParseArgs([]string{"--int", "1", "2", "--int", "3"})
    if parser.LenList("int") != 2 {
        t.Fail()
    }
    if parser.GetIntList("int")[0] != 1 {
        t.Fail()
    }
    if parser.GetIntList("int")[1] != 3 {
        t.Fail()
    }
}


func TestIntListShortform(t *testing.T) {
    parser := NewParser()
    parser.NewInt("int i")
    parser.ParseArgs([]string{"-i", "1", "2", "-i", "3"})
    if parser.LenList("int") != 2 {
        t.Fail()
    }
    if parser.GetIntList("int")[0] != 1 {
        t.Fail()
    }
    if parser.GetIntList("int")[1] != 3 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Float options.
// -----------------------------------------------------------------------------


func TestFloatOptionEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("float", 1.1)
    parser.ParseArgs([]string{})
    if parser.GetFloat("float") != 1.1 {
        t.Fail()
    }
}


func TestFloatOptionMissing(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("float", 1.1)
    parser.ParseArgs([]string{"foo", "bar"})
    if parser.GetFloat("float") != 1.1 {
        t.Fail()
    }
}


func TestFloatOptionLongform(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("float", 1.1)
    parser.ParseArgs([]string{"--float", "2.2"})
    if parser.GetFloat("float") != 2.2 {
        t.Fail()
    }
}


func TestFloatOptionShortform(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("float f", 1.1)
    parser.ParseArgs([]string{"-f", "2.2"})
    if parser.GetFloat("float") != 2.2 {
        t.Fail()
    }
}


func TestFloatOptionNegative(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("float", 1.1)
    parser.ParseArgs([]string{"--float", "-2.2"})
    if parser.GetFloat("float") != -2.2 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Float lists.
// -----------------------------------------------------------------------------


func TestFloatListEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("flt")
    parser.ParseArgs([]string{})
    if parser.LenList("flt") != 0 {
        t.Fail()
    }
}


func TestFloatListLongform(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("flt")
    parser.ParseArgs([]string{"--flt", "1", "2", "--flt", "3"})
    if parser.LenList("flt") != 2 {
        t.Fail()
    }
    if parser.GetFloatList("flt")[0] != 1 {
        t.Fail()
    }
    if parser.GetFloatList("flt")[1] != 3 {
        t.Fail()
    }
}


func TestFloatListShortform(t *testing.T) {
    parser := NewParser()
    parser.NewFloat("flt f")
    parser.ParseArgs([]string{"-f", "1", "2", "-f", "3"})
    if parser.LenList("flt") != 2 {
        t.Fail()
    }
    if parser.GetFloatList("flt")[0] != 1 {
        t.Fail()
    }
    if parser.GetFloatList("flt")[1] != 3 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Multiple options.
// -----------------------------------------------------------------------------


func TestMultiOptionsEmpty(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool1")
    parser.NewFlag("bool2 b")
    parser.NewString("string1", "default1")
    parser.NewString("string2 s", "default2")
    parser.NewInt("int1", 101)
    parser.NewInt("int2 i", 202)
    parser.NewFloat("float1", 1.1)
    parser.NewFloat("float2 f", 2.2)
    parser.ParseArgs([]string{})
    if parser.GetFlag("bool1") != false {
        t.Fail()
    }
    if parser.GetFlag("bool2") != false {
        t.Fail()
    }
    if parser.GetString("string1") != "default1" {
        t.Fail()
    }
    if parser.GetString("string2") != "default2" {
        t.Fail()
    }
    if parser.GetInt("int1") != 101 {
        t.Fail()
    }
    if parser.GetInt("int2") != 202 {
        t.Fail()
    }
    if parser.GetFloat("float1") != 1.1 {
        t.Fail()
    }
    if parser.GetFloat("float2") != 2.2 {
        t.Fail()
    }
}


func TestMultiOptionsLongform(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool1")
    parser.NewFlag("bool2 b")
    parser.NewString("string1", "default1")
    parser.NewString("string2 s", "default2")
    parser.NewInt("int1", 101)
    parser.NewInt("int2 i", 202)
    parser.NewFloat("float1", 1.1)
    parser.NewFloat("float2 f", 2.2)
    parser.ParseArgs([]string{
        "--bool1",
        "--bool2",
        "--string1", "value1",
        "--string2", "value2",
        "--int1", "303",
        "--int2", "404",
        "--float1", "3.3",
        "--float2", "4.4",
    })
    if parser.GetFlag("bool1") != true {
        t.Fail()
    }
    if parser.GetFlag("bool2") != true {
        t.Fail()
    }
    if parser.GetString("string1") != "value1" {
        t.Fail()
    }
    if parser.GetString("string2") != "value2" {
        t.Fail()
    }
    if parser.GetInt("int1") != 303 {
        t.Fail()
    }
    if parser.GetInt("int2") != 404 {
        t.Fail()
    }
    if parser.GetFloat("float1") != 3.3 {
        t.Fail()
    }
    if parser.GetFloat("float2") != 4.4 {
        t.Fail()
    }
}


func TestMultiOptionsShortform(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool1")
    parser.NewFlag("bool2 b")
    parser.NewString("string1", "default1")
    parser.NewString("string2 s", "default2")
    parser.NewInt("int1", 101)
    parser.NewInt("int2 i", 202)
    parser.NewFloat("float1", 1.1)
    parser.NewFloat("float2 f", 2.2)
    parser.ParseArgs([]string{
        "--bool1",
        "-b",
        "--string1", "value1",
        "-s", "value2",
        "--int1", "303",
        "-i", "404",
        "--float1", "3.3",
        "-f", "4.4",
    })
    if parser.GetFlag("bool1") != true {
        t.Fail()
    }
    if parser.GetFlag("bool2") != true {
        t.Fail()
    }
    if parser.GetString("string1") != "value1" {
        t.Fail()
    }
    if parser.GetString("string2") != "value2" {
        t.Fail()
    }
    if parser.GetInt("int1") != 303 {
        t.Fail()
    }
    if parser.GetInt("int2") != 404 {
        t.Fail()
    }
    if parser.GetFloat("float1") != 3.3 {
        t.Fail()
    }
    if parser.GetFloat("float2") != 4.4 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Condensed short-form options.
// -----------------------------------------------------------------------------


func TestCondensedOptions(t *testing.T) {
    parser := NewParser()
    parser.NewFlag("bool b")
    parser.NewString("string s", "default")
    parser.NewInt("int i", 101)
    parser.NewFloat("float f", 1.1)
    parser.ParseArgs([]string{"-bsif", "value", "202", "2.2"})
    if parser.GetFlag("bool") != true {
        t.Fail()
    }
    if parser.GetString("string") != "value" {
        t.Fail()
    }
    if parser.GetInt("int") != 202 {
        t.Fail()
    }
    if parser.GetFloat("float") != 2.2 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Positional arguments.
// -----------------------------------------------------------------------------


func TestPositionalArgsEmpty(t *testing.T) {
    parser := NewParser()
    parser.ParseArgs([]string{})
    if parser.HasArgs() != false {
        t.Fail()
    }
}


func TestPositionalArgs(t *testing.T) {
    parser := NewParser()
    parser.ParseArgs([]string{"foo", "bar"})
    if parser.HasArgs() != true {
        t.Fail()
    }
    if parser.NumArgs() != 2 {
        t.Fail()
    }
    if parser.GetArg(0) != "foo" {
        t.Fail()
    }
    if parser.GetArg(1) != "bar" {
        t.Fail()
    }
    if parser.GetArgs()[0] != "foo" {
        t.Fail()
    }
    if parser.GetArgs()[1] != "bar" {
        t.Fail()
    }
}


func TestPositionalArgsAsInts(t *testing.T) {
    parser := NewParser()
    parser.ParseArgs([]string{"1", "11"})
    if parser.GetArgsAsInts()[0] != 1 {
        t.Fail()
    }
    if parser.GetArgsAsInts()[1] != 11 {
        t.Fail()
    }
}


func TestPositionalArgsAsFloats(t *testing.T) {
    parser := NewParser()
    parser.ParseArgs([]string{"1.1", "11.1"})
    if parser.GetArgsAsFloats()[0] != 1.1 {
        t.Fail()
    }
    if parser.GetArgsAsFloats()[1] != 11.1 {
        t.Fail()
    }
}


// -----------------------------------------------------------------------------
// Commands
// -----------------------------------------------------------------------------


func callback(parser *ArgParser) {}


func TestCommandAbsent(t *testing.T) {
    parser := NewParser()
    parser.NewCmd("cmd", "helptext", callback)
    parser.ParseArgs([]string{})
    if parser.HasCmd() != false {
        t.Fail()
    }
}


func TestCommandPresent(t *testing.T) {
    parser := NewParser()
    cmdParser := parser.NewCmd("cmd", "helptext", callback)
    parser.ParseArgs([]string{"cmd"})
    if parser.HasCmd() != true {
        t.Fail()
    }
    if parser.GetCmdName() != "cmd" {
        t.Fail()
    }
    if parser.GetCmdParser() != cmdParser {
        t.Fail()
    }
}


func TestCommandWithOptions(t *testing.T) {
    parser := NewParser()
    cmdParser := parser.NewCmd("cmd", "helptext", callback)
    cmdParser.NewFlag("bool")
    cmdParser.NewString("string", "default")
    cmdParser.NewInt("int", 101)
    cmdParser.NewFloat("float", 1.1)
    parser.ParseArgs([]string{
        "cmd",
        "foo", "bar",
        "--string", "value",
        "--int", "202",
        "--float", "2.2",
    })
    if parser.HasCmd() != true {
        t.Fail()
    }
    if parser.GetCmdName() != "cmd" {
        t.Fail()
    }
    if parser.GetCmdParser() != cmdParser {
        t.Fail()
    }
    if cmdParser.HasArgs() != true {
        t.Fail()
    }
    if cmdParser.NumArgs() != 2 {
        t.Fail()
    }
    if cmdParser.GetString("string") != "value" {
        t.Fail()
    }
    if cmdParser.GetInt("int") != 202 {
        t.Fail()
    }
    if cmdParser.GetFloat("float") != 2.2 {
        t.Fail()
    }
}
