package main

import (
    "fmt"
    "os"

    "github.com/fatih/color"
    "github.com/urfave/cli"
    "encoding/binary"
)

func getExe64Command() cli.Command {
    command := cli.Command{
        Name:        "exe64",
        Aliases:     []string{"64", "exe"},
        Category:    "Common tools",
        Usage:       "Check a exe is 32bit or 64 bit",
        UsageText:   "Example: mo exe64 c:\a.exe ",
        Description: "Check a exe is 32bit or 64 bit.",
        ArgsUsage:   "<filename>",
        //Flags: []cli.Flag{
        //    cli.BoolFlag{
        //        Name:   "show,s",
        //        Usage:  "show current password",
        //        Hidden: true,
        //    },
        //},
        Action: func(c *cli.Context) error {
            //fmt.Println("base64 task: ", c.Args().First())
            //fmt.Println("base64 task: ", c.Args())
            color.Yellow("Author:Liu Xuan,last modified at 2017/11/02")
            filepath := c.Args().First()
            fi, err := os.Open(filepath)
            checkErr(err)
            defer fi.Close()

            const NBUF = 4096
            var buf [NBUF]byte

            nr, err := fi.Read(buf[:])
            checkErr(err)
            if nr > 0{
                //fmt.Println(buf)
                //fmt.Println(buf[:2])
                startcode := binary.LittleEndian.Uint16(buf[0:2])
                peOffset :=binary.LittleEndian.Uint16(buf[60:62])
                if startcode==0x5a4d{
                    fmt.Println("head is 0x4d5a,this is a pe file")
                    fmt.Printf("PE offset is %v \r\n",peOffset)
                    machine:= binary.LittleEndian.Uint16(buf[peOffset+4:peOffset+6])
                    var arch string;
                    switch machine {
                    case 0x8664 : arch = "x64"
                    case 0x0200 : arch = "Intel Itanium"
                    case 0x014c : arch = "x86"
                    default:
                        arch = "not know"
                    }
                    //fmt.Printf("%X \r\n",machine)
                    fmt.Printf("ARCH is: %s\r\n",arch)
                }else{
                    fmt.Println("head is not 0x4d5a,this is not a pe file")
                }
            }

            //calcBase(c.Args().First(), c.Bool("debase"))
            return nil
        },
    }
    return command
}
