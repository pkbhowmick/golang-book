package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

func main() {
	// initialize a flag called flag1 of type int
	var ip *int = flag.IntP("flag1", "f", 1234, "help message for flag1")
	flag.Lookup("flag1").NoOptDefVal = "4321"

	// initialize a flag called flag1 of type int
	var flagvar int
	flag.IntVarP(&flagvar, "flag2", "p", 5678, "help message for flag2")
	flag.Lookup("flag2").NoOptDefVal = "8765"

	// initialize a flag called flag3 of type bool
	var flagbool bool
	flag.BoolVarP(&flagbool, "flag3", "b", true, "help message for flag3")
	flag.Lookup("flag3").NoOptDefVal = "false"

	flag.Parse()
	fmt.Println("ip:", *ip)
	fmt.Println("flagvar", flagvar)
	fmt.Println("flagbool", flagbool)

	// Lookup a flag with flagname
	flag4 := flag.Lookup("flag3")
	fmt.Println(flag4.Value)

	// Lookup a flag with flag shorthand name
	flag5 := flag.ShorthandLookup("f")
	fmt.Println(flag5.Value)

	// Initialized a flagset and add a flag in the flagset
	flags := flag.NewFlagSet("flagset", flag.ExitOnError)
	flags.AddFlag(flag.Lookup("flag1"))

	val, err := flags.GetInt("flag1")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Value from flagset:", val)

	// Deprecating a flag by specifying its name and a usage message
	//flags.MarkDeprecated("flag1", "please use --flag2 instead")

	// Deprecate a flag shorthand by specifying its flag name and a usage message
	//flags.MarkShorthandDeprecated("flag3", "please use --flag3 only")

}
