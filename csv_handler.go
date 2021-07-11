package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "io"
)

func readCSVFile(filePath string) []map[string]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file: " + filePath, err);
    }
    defer f.Close()

    r := csv.NewReader(f)
    rows := []map[string]string{}
    var header []string
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        if header == nil {
            header = record
        } else {
            dict := map[string]string{}
            for i := range header {
                dict[header[i]] = record[i]
            }
            rows = append(rows, dict)
        }
    }
    return rows
}

func main() {
    if len(os.Args) < 1 {
        log.Fatal("Too few arguments.")
    }
    filename := os.Args[1]
    records := readCSVFile(filename)
    fmt.Println(records)

    //for _, i := range records {
    //    fmt.Println(i["a"])
    //}
}
