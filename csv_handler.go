package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
)

func process_csv_line_by_line( filePath string ) {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file: " + filePath, err)
    }
    defer f.Close()

    r := csv.NewReader(f)
    var header []string
    pu_count, do_count := 0, 0
    for {
        record, err := r.Read()
        if err == io.EOF {
            fmt.Println("PU Count: " + strconv.Itoa(pu_count))
            fmt.Println("DO Count: " + strconv.Itoa(do_count))
            return
        }
        if err != nil {
            log.Fatal(err)
        }
        if header == nil {
            header = record
        } else {
            csv_record := map[string]string{}
            for i := range header {
                csv_record[header[i]] = record[i]
            }
            payment_type := csv_record["payment_type"]
            pu_location  := csv_record["PULocationID"]
            do_location  := csv_record["DOLocationID"]
            if payment_type == "3" {
                if pu_location == "170" {
                    pu_count += 1
                } else if do_location == "170" {
                    do_count += 1
                }
            }
        }
    }

    return
}

func main() {
    if len(os.Args) < 1 {
        log.Fatal("Too few arguments.")
    }
    filename := os.Args[1]
    process_csv_line_by_line(filename)
}
