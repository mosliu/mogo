package main

import (
    "fmt"
    "os"
    "time"

    "github.com/urfave/cli"
    "github.com/fatih/color"
    "dict"
)



func main() {
    app := cli.NewApp()
    app.Name = "Moses ToolBox"
    app.Version = "0.0.1.2"
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        {
            Name:  "Moses",
            Email: "mogo@liuxuan.net",
        },
    }
    app.Copyright = "(c) 2017 Dreamaker Studio"
    app.Usage = "Tools box for daily usage"
    app.ArgsUsage = "[args and such]"

    app.Commands = []cli.Command{
        {
            Name:        "dict",
            Aliases:     []string{"d"},
            Category:    "Dict tools",
            Usage:       "find the dict on Youdao dict",
            UsageText:   "dict usagetext",
            Description: "dict Description",
            ArgsUsage:   "ArgsUsage",
            Action: func(c *cli.Context) error {
                color.Red("%s is %t",c.Args().First(),mosDict.IsChinese(c.Args().First()))
                //fmt.Println("added task: ", c.Args().First())
                return nil
            },

        },
        getBase64Command(),
        getSqlite3Command(),
    }

    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "lang",
            Value: "english",
            Usage: "language for the greeting",
        },
    }

    app.Action = func(c *cli.Context) error {
        name := "Nefertiti"
        if c.NArg() > 0 {
            name = c.Args().Get(0)
        }
        if c.String("lang") == "spanish" {
            fmt.Println("Hola", name)
        } else {
            fmt.Println("Hello", name)
        }
        return nil
    }

    defer func() {
        if e := recover(); e != nil {
            color.Red("Panicing,error occured: %s\r\n", e)
        }
    }()

    app.Run(os.Args)

}
