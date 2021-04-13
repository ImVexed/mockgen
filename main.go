package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cheggaaa/pb/v3"
	"github.com/mattn/go-isatty"
)

var CLI struct {
	SQL struct {
		Bulk  bool   `short:"b" help:"Bulk insert"`
		Table string `arg help:"Table to insert into"`
		Count uint64 `arg help:"Ammount of rows to generate"`

		Columns map[string]string `arg type:"file:" help:"Columns to create (name={firstname} id={uuid} bio={sentence:30})"`
	} `cmd help:"Generate SQL mock data"`
}

func main() {
	kong.Parse(&CLI)

	columns := []string{}

	for column := range CLI.SQL.Columns {
		columns = append(columns, column)
	}

	insertStmt := fmt.Sprintf("INSERT INTO \"%s\" (\"%s\") VALUES", CLI.SQL.Table, strings.Join(columns, "\",\""))

	if CLI.SQL.Bulk {
		io.WriteString(os.Stdout, insertStmt+"\n")
	}

	bar := pb.New(int(CLI.SQL.Count))

	shouldProgress := !isatty.IsTerminal(os.Stdout.Fd())

	if shouldProgress {
		bar.SetWriter(os.Stderr)
		bar.Start()
	}

	for i := uint64(0); i < CLI.SQL.Count; i++ {
		values := []string{}

		for _, column := range columns {
			value := gofakeit.Generate(CLI.SQL.Columns[column])

			if _, err := strconv.Atoi(value); err != nil {
				value = "'" + strings.Replace(value, "'", "''", -1) + "'"
			}
			values = append(values, value)
		}

		row := ""
		if !CLI.SQL.Bulk {
			row = insertStmt + " "
		}

		row += fmt.Sprintf("(%s)", strings.Join(values, ","))

		if CLI.SQL.Bulk {
			if i == CLI.SQL.Count-1 {
				row += ";\n"
			} else {
				row += ",\n"
			}
		} else {
			row += ";\n"
		}

		io.WriteString(os.Stdout, row)

		if shouldProgress {
			bar.Increment()
		}
	}

	if shouldProgress {
		bar.Finish()
	}
}
