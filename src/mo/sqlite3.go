package main

import (
    "database/sql"

    "os"

    "github.com/fatih/color"
    _ "github.com/mattn/go-sqlite3"
    "github.com/urfave/cli"
)

func getSqlite3Command() cli.Command {
    command := cli.Command{
        Name:        "sqlite3",
        Aliases:     []string{"s3", "db"},
        Category:    "Labthink tools",
        Usage:       "Change Sqlite3 UserPass",
        UsageText:   "Example: mo sqlite3 admin 123456 ",
        Description: "Calculate base64 or debase64.",
        ArgsUsage:   "<DbFullName> <UserToChange> <PasswordToSet>",
        Flags: []cli.Flag{
            cli.BoolFlag{
                Name:   "show,s",
                Usage:  "show current password",
                Hidden: true,
            },
        },
        Action: func(c *cli.Context) error {
            //fmt.Println("base64 task: ", c.Args().First())
            //fmt.Println("base64 task: ", c.Args())
            checkArgResult := checkArgs(c)
            if checkArgResult > 0 {
                return nil
            }

            DbFile := c.Args().First()

            username := c.Args().Get(1)
            password := c.Args().Get(2)
            showflag := c.Bool("show")

            changePassword(DbFile, username, password, showflag)

            //calcBase(c.Args().First(), c.Bool("debase"))
            return nil
        },
    }
    return command
}

/**
    Check if The args satisfied the input protocol
 */
func checkArgs(c *cli.Context) (rtn int) {
    rtn = 0
    if c.NArg() < 3 {
        color.Red("Wrong Input.")
        color.Red("Use mo sqlite3 <dbFullName> <UserToChange> <PasswordToSet>")
        color.Red("Example: mo sqlite3 USER.DB admin 111111 ")
        rtn = 1
        return
    }

    _, err := os.Stat(c.Args().First())
    if err != nil {
        if os.IsNotExist(err) {
            color.Red("File  %s does not exist.", c.Args().First())
            rtn = 2
            return
        }
    }
    return

}

/**
    Change Password
 */
func changePassword(DbFile string, username string, password string, showFlag bool) {


    db, err := sql.Open("sqlite3", "./USER.DB")
    checkErr(err)

    if showFlag {
        //查询数据
        color.Green("Show current users and password in the db")
        rows, err := db.Query("SELECT V_name,V_Pwd,name FROM users")
        checkErr(err)
        color.Blue("%-20s,%-20s,%-20s", "V_name", "V_Pwd", "name")
        for rows.Next() {
            var V_name string
            var V_Pwd string
            var name string
            err = rows.Scan(&V_name, &V_Pwd, &name)
            checkErr(err)
            color.Blue("%-20s,%-20s,%-20s", V_name, V_Pwd, name)
        }
    }else{
        color.Green("Try to change %s 's password to %s", username, password)
        stmt, err := db.Prepare("update users set V_Pwd=? where V_name=?")
        checkErr(err)

        res, err := stmt.Exec(password,username )
        checkErr(err)
        affect, err := res.RowsAffected()
        checkErr(err)
        if affect>0{
            color.Blue("Done! %d rows affected", affect)
        }
    }

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
