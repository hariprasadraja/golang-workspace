package main

/*
gobrace is a linting tool writern to format my go code.
it will insert line break when it founds a brace followed by text in next line
Example
    if all == "true" {
		baseQuery = allDimensionsQuery
	}
	res, err = d.DBr.Query(baseQuery, metricCode, accountID)


	the above code in the go file will be parsed and transformed into

	if all == "true" {
		baseQuery = allDimensionsQuery
	}

	res, err = d.DBr.Query(baseQuery, metricCode, accountID)
*/
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var (
		filename string
		dir      string
		err      error
	)

	flag.StringVar(&filename, "f", "", "absolute file path which has to be linted")
	flag.StringVar(&dir, "d", "", "directory path where files has to be linted")
	flag.Parse()
	if flag.NFlag() != 1 {
		fmt.Fprintf(os.Stderr, "only one flag is allowed and must be set but got %v flags", flag.NFlag())
		flag.Usage()
		os.Exit(1)
	}

	if dir != "" {
		err = filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !strings.Contains(path, "vendor") && strings.Contains(file.Name(), "go") {
				return updateFileWithRegex(path)
			}

			return nil
		})
	}

	if filename != "" && strings.HasSuffix(filename, ".go") {
		err = updateFileWithRegex(filename)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func updateFileWithRegex(filename string) error {
	exp, err := regexp.Compile("[^{]}\n.\\w+")
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data := exp.ReplaceAllFunc(file, func(line []byte) []byte {
		linestring := string(line)
		data := strings.SplitN(linestring, "}", 2)
		var newdata string
		for i, val := range data {
			newdata += val

			if i == 0 && strings.TrimSpace(val) != "" {
				newdata += "}"
				continue
			}

			if i == 0 {
				newdata += "}" + "\n"
			}
		}

		return []byte(newdata)
	})

	ioutil.WriteFile(filename, data, os.ModePerm)
	return nil
}
