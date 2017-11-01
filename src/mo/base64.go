package main

import (
    "encoding/base64"
    "os"

    "github.com/fatih/color"
    "github.com/urfave/cli"
)

func getBase64Command() cli.Command {
    command := cli.Command{
        Name:        "base64",
        Aliases:     []string{"b64"},
        Category:    "Common tools",
        Usage:       "Calculate base64 for input",
        UsageText:   "Example: mo b64 \"Hello world!\" ",
        Description: "Calculate base64 or debase64.",
        ArgsUsage:   "ArgsUsage",
        Flags: []cli.Flag{
            cli.BoolFlag{
                Name:  "debase,d",
                Usage: "debase input string Example: mo b64 -d SGVsbG8gd29ybGQh",
            },
        },
        Action: func(c *cli.Context) error {
            //fmt.Println("base64 task: ", c.Args().First())
            //fmt.Println("base64 task: ", c.Args())
            calcBase(c.Args().First(), c.Bool("debase"))
            return nil
        },
    }
    return command
}

func calcBase(content string, isDebase bool) {
    padEqual := ""
    calcBasePlus(content, isDebase, padEqual)
}

func calcBasePlus(content string, isDebase bool, padEqual string) {
    if isDebase {
        decodeString, err := base64.StdEncoding.DecodeString(content + padEqual)
        if err != nil {
            if len(padEqual) < 2 {
                padEqual = padEqual + "="
                //fmt.Println("add =")
                calcBasePlus(content, isDebase, padEqual)
            } else {
                //fmt.Println(padEqual)
                color.Red("Error occurred:", err)
                color.Green("%s  Decode result:    %s", content, string(decodeString))
                os.Exit(1)
            }
        } else {
            color.Green("%s  Decode result:    %s", content+padEqual, string(decodeString))
            os.Exit(1)
        }
        //fmt.Println(string(decodeString))
    } else {
        encodeString := base64.StdEncoding.EncodeToString([]byte(content))
        color.HiBlue("%s  encode result:    %s", content, encodeString)
        //fmt.Println(encodeString)
    }
}
