## Fetch Tickers for exchanges


	for _, e := range exchanges.All() {

		err := DownloadFile("pkg/data/eodhistoricaldata/fetcher/data/exchanges/raw/"+e.Code+".csv",
			fmt.Sprintf("%s/exchanges/%s?api_token=%s", config.BaseURL, e.Code, config.Token))

		if err != nil {
			panic(err)
		}
	}
	

## CSV Reader  setup

    
        func newCsvReader(r io.Reader) *csv.Reader {
            reader := csv.NewReader(r)
            reader.Comma = ','
            reader.ReuseRecord = true
            // skip first line
            _, err := reader.Read()
            if err != nil {
                log.Fatal(err)
            }
            return reader
        }
