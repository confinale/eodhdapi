# eodhdapi
This is a simple eod historical data api for Golang. 
It does a bunch of conversions to handle the data from eodhdapi in a easier way.


# example
counts all the number of prices for all exchanges.
package main

    import (
        "context"
        "fmt"
        "github.com/confinale/eodhdapi"
        "github.com/confinale/eodhdapi/exchanges"
        "os"
        "time"
    )
    
    func main() {
        d := eodhdapi.NewDefaultEOD(os.Getenv("EODHD_TOKEN"))
    
        for _, e := range exchanges.All() {
    
            prices := make(chan eodhdapi.EODPrice)
            done := make(chan int, 1)
    
            go func(f chan eodhdapi.EODPrice, d chan int) {
                count := 0
                for range f {
                    count++
                }
                d <- count
            }(prices, done)
    
            if err := d.FetchPrices(context.Background(), prices, e, time.Date(2019, 9, 25, 0, 0, 0, 0, time.UTC)); err != nil {
                panic(err)
            }
            close(prices)
    
            count := <-done
    
            fmt.Printf("exchange %s had %d elements\n", e.Code, count)
        }
    }