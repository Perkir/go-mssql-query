package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"io/ioutil"
	"os"
	"time"
)

func getSQL(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(" Error open sql:", err.Error())
		return ""
	}
	return string(data)
}

func main() {

	// timeout 10 hours.
	if len(os.Args) < 5 {
		fmt.Printf("%s <user> <password> <server> <queryfile>\n", os.Args[0])
		os.Exit(1)
	}
	user := os.Args[1]
	password := os.Args[2]
	server := os.Args[3]
	sqlConn := fmt.Sprintf("sqlserver://%s:%s@%S?connection+timeout=36000", user, password, server)
	db, err := sql.Open("mssql", sqlConn)
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return
	}
	defer db.Close()
	q := getSQL(os.Args[4])

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println("error with sql:", err)
		return
	}
	defer rows.Close()
	cols, err := rows.Columns()
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("|")
		}
		fmt.Print(cols[i])
	}
	fmt.Println("")

	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("|")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()
		//break
	}
}

func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}
