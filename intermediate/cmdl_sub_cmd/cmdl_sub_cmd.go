package main

import (
	"fmt"
	"os"
	"flag"
)

func main(){
	//"e:\coding\golang\basics\intermediate\cmdl_sub_cmd\cmdl_sub_cmd.go" golang -name andro -exp 10
	golang := flag.NewFlagSet("golang",flag.ExitOnError)
	love := flag.NewFlagSet("love",flag.ExitOnError)

	userName := golang.String("name","nil","username")
	experience := golang.Int("exp",0,"username")

	loveTrue := love.Bool("love",false,"do you love")

	if len(os.Args) < 2 {
		fmt.Println("no args passed!")
		return
	}

	switch os.Args[1]{
		case "golang":
			golang.Parse(os.Args[2:])
			fmt.Println("username:", *userName)
			fmt.Println("experience:", *experience)
		case "love":
			love.Parse(os.Args[2:])
			fmt.Println("love:", *loveTrue)

		default:
			fmt.Println("invalid sub cmd entered")
	}

}