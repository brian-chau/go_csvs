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
    payment_type_index := 0
    pu_index, do_index := 0, 0
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
            for i := range header {
                if header[i] == "payment_type" {
                    payment_type_index = i;
                } else if header[i] == "PULocationID" {
                    pu_index           = i;
                } else if header[i] == "DOLocationID" {
                    do_index           = i;
                }
            }
        } else {
            if record[payment_type_index] == "3" {
                if record[pu_index] == "170" {
                    pu_count += 1
                } else if record[do_index] == "170" {
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
    process_csv_line_by_line( filename )
}
