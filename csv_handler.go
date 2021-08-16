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
    is_first_line := true
    header := map[string]int{}

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
        if is_first_line {
            is_first_line = false
            for i := range record {
                header[record[i]] = i
            }
        } else {
            if record[header["payment_type"]] == "3" {
                if record[header["PULocationID"]] == "170" {
                    pu_count += 1
                } else if record[header["DOLocationID"]] == "170" {
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
