package exchanges

// Exchange holds the information for an eod
type Exchange struct {
	Code                   string
	Name                   string
	FigiExchangeCodes      []string
	ExchangeCodeComponents []string
}

func newExchange(code, name string) *Exchange {
	return &Exchange{Code: code, Name: name, ExchangeCodeComponents: []string{code}, FigiExchangeCodes: []string{}}
}

func (e *Exchange) setExchangeCodeComponents(codes ...string) *Exchange {
	e.ExchangeCodeComponents = codes
	return e
}

// Exchanges is a list of exchanges
type Exchanges []*Exchange

// GetByCode returns exchanges by code - returns nil if not found
func (e Exchanges) GetByCode(code string) *Exchange {
	for _, v := range e {
		if v.Code == code {
			return v
		}
	}
	return nil
}

// All gives all known exchanges
func All() Exchanges {
	return []*Exchange{
		newExchange("US", "USA Stocks").setExchangeCodeComponents("NASDAQ", "NYSE", "BATS", "AMEX"),
		newExchange("LSE", "London Exchange"),
		newExchange("V", "TSX Venture Exchange"),
		newExchange("TO", "Toronto Exchange"),
		newExchange("CN", "Canadian Securities Exchange"),
		newExchange("BE", "Berlin Exchange"),
		newExchange("F", "Frankfurt Exchange"),
		newExchange("STU", "Stuttgart Exchange"),
		newExchange("MU", "Munich Exchange"),
		newExchange("HA", "Hanover Exchange"),
		newExchange("HM", "Hamburg Exchange"),
		newExchange("XETRA", "XETRA Exchange"),
		newExchange("DU", "Dusseldorf Exchange"),
		newExchange("MI", "Borsa Italiana"),
		newExchange("LU", "Luxembourg Stock Exchange"),
		newExchange("VI", "Vienna Exchange"),
		newExchange("PA", "Euronext Paris"),
		newExchange("BR", "Euronext Brussels"),
		newExchange("MC", "Madrid Exchange"),
		newExchange("LS", "Euronext Lisbon"),
		newExchange("VX", "Swiss Exchange"),
		newExchange("AS", "Euronext Amsterdam"),
		newExchange("SW", "SIX Swiss Exchange"),
		newExchange("IC", "Iceland Exchange"),
		newExchange("ST", "Stockholm Exchange"),
		newExchange("OL", "Oslo Stock Exchange"),
		newExchange("NFN", "Nasdaq First North"),
		newExchange("CO", "Coppenhagen Exchange"),
		newExchange("HE", "Helsinki Exchange"),
		newExchange("NB", "Nasdaq Baltic"),
		newExchange("IR", "Irish Exchange"),
		newExchange("HK", "Hong Kong Exchange"),
		newExchange("TA", "Tel Aviv Exchange"),
		newExchange("KQ", "KOSDAQ"),
		newExchange("KO", "Korea Stock Exchange"),
		newExchange("MCX", "MICEX Russia"),
		newExchange("AU", "Australia Exchange"),
		newExchange("WAR", "Warsaw Stock Exchange"),
		newExchange("NZ", "New Zealand Exchange"),
		newExchange("BUD", "Budapest Stock Exchange"),
		newExchange("PSE", "Philippine Stock Exchange"),
		newExchange("IS", "Istanbul Stock Exchange"),
		newExchange("SG", "Singapore Exchange"),
		newExchange("BSE", "Bombay Exchange"),
		newExchange("NSE", "NSE (India)"),
		newExchange("SHG", "Shanghai Exchange"),
		newExchange("JSE", "Johannesburg Exchange"),
		newExchange("BK", "Thailand Exchange"),
		newExchange("SR", "Saudi Arabia Exchange"),
		newExchange("KAR", "Karachi Stock Exchange"),
		newExchange("JK", "Jakarta Exchange"),
		newExchange("TSE", "Tokyo Stock Exchange"),
		newExchange("SHE", "Shenzhen Exchange"),
		newExchange("VN", "Vietnam Stocks"),
		newExchange("KLSE", "Kuala Lumpur Exchange"),
		newExchange("SA", "Sao Paolo Exchange"),
		newExchange("MX", "Mexican Exchange"),
		newExchange("IL", "London IL"),
		newExchange("TWO", "Taiwan OTC Exchange"),
		newExchange("TW", "Taiwan Exchange"),

		newExchange("CC", "Cryptocurrencies"),
		newExchange("COMM", "Commodities"),
		newExchange("FOREX", "FOREX"),
		newExchange("BOND", "Bonds"),
		newExchange("INDX", "Indexes"),
	}
}
