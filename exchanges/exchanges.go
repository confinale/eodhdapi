package exchanges

// Exchange holds the information for an eod
type Exchange struct {
	Code                   string
	MicCodes               []string
	Name                   string
	FigiExchangeCodes      []string
	ExchangeCodeComponents []string
	ForceLenient           bool
}

func newExchange(code, name string) *Exchange {
	return &Exchange{Code: code, Name: name, ExchangeCodeComponents: []string{code}, FigiExchangeCodes: []string{}}
}

func (e *Exchange) forceLenient() *Exchange {
	e.ForceLenient = true
	return e
}

func (e *Exchange) setMicCodes(codes ...string) *Exchange {
	e.MicCodes = codes
	return e
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

		newExchange("US", "USA Stocks").setMicCodes("XNAS", "XNYS").setExchangeCodeComponents("NASDAQ", "NYSE", "BATS", "AMEX"),
		newExchange("LSE", "London Exchange").setMicCodes("XLON"),
		newExchange("TO", "Toronto Exchange").setMicCodes("XTSX"),
		newExchange("V", "TSX Venture Exchange").setMicCodes("XTSX "),
		newExchange("CN", "Canadian Securities Exchange").setMicCodes("XCNQ"),
		newExchange("BE", "Berlin Exchange").setMicCodes("XBER"),
		newExchange("F", "Frankfurt Exchange").setMicCodes("XFRA"),
		newExchange("STU", "Stuttgart Exchange").setMicCodes("XSTU"),
		newExchange("HM", "Hamburg Exchange").setMicCodes("XHAM"),
		newExchange("HA", "Hanover Exchange").setMicCodes("XHAN"),
		newExchange("MU", "Munich Exchange").setMicCodes("XMUN"),
		newExchange("XETRA", "XETRA Exchange").setMicCodes("XETR"),
		newExchange("DU", "Dusseldorf Exchange").setMicCodes("XDUS"),
		newExchange("MI", "Borsa Italiana").setMicCodes("XMIL"),
		newExchange("VI", "Vienna Exchange").setMicCodes("XWBO"),
		newExchange("LU", "Luxembourg Stock Exchange").setMicCodes("XLUX"),
		newExchange("PA", "Euronext Paris").setMicCodes("XPAR"),
		newExchange("BR", "Euronext Brussels").setMicCodes("XBRU"),
		newExchange("MC", "Madrid Exchange").setMicCodes("BMEX"),
		newExchange("AS", "Euronext Amsterdam").setMicCodes("XAMS"),
		newExchange("VX", "Swiss Exchange").setMicCodes("XSWX"),
		newExchange("LS", "Euronext Lisbon").setMicCodes("XLIS"),
		newExchange("SW", "SIX Swiss Exchange").setMicCodes("XSWX"),
		newExchange("IR", "Irish Exchange").setMicCodes("XDUB"),
		newExchange("ST", "Stockholm Exchange").setMicCodes("XSTO"),
		newExchange("OL", "Oslo Stock Exchange").setMicCodes("XOSL"),
		newExchange("CO", "Coppenhagen Exchange").setMicCodes("XCSE"),
		newExchange("HE", "Helsinki Exchange").setMicCodes("XHEL"),
		newExchange("IC", "Iceland Exchange").setMicCodes("XICE"),
		newExchange("NFN", "Nasdaq First North").setMicCodes("XCSE"),
		newExchange("NB", "Nasdaq Baltic").setMicCodes("XTAL"),
		newExchange("HK", "Hong Kong Exchange").setMicCodes("XHKG"),
		newExchange("TA", "Tel Aviv Exchange").setMicCodes("XTAE"),
		newExchange("KO", "Korea Stock Exchange").setMicCodes("XKRX"),
		newExchange("KQ", "KOSDAQ").setMicCodes("XKOS"),
		newExchange("IS", "Istanbul Stock Exchange").setMicCodes("XIST"),
		newExchange("BUD", "Budapest Stock Exchange").setMicCodes("XBUD"),
		newExchange("WAR", "Warsaw Stock Exchange").setMicCodes("XWAR"),
		newExchange("PSE", "Philippine Stock Exchange").setMicCodes("XPHS").forceLenient(),
		newExchange("SG", "Singapore Exchange").setMicCodes("XSES"),
		newExchange("BSE", "Bombay Exchange").setMicCodes("XBOM"),
		newExchange("NSE", "NSE (India)").setMicCodes("XNSE"),
		newExchange("JSE", "Johannesburg Exchange").setMicCodes("XJSE"),
		newExchange("BK", "Thailand Exchange").setMicCodes("XBKK"),
		newExchange("SR", "Saudi Arabia Exchange").setMicCodes("XSAU"),
		newExchange("JK", "Jakarta Exchange").setMicCodes("XIDX"),
		newExchange("TSE", "Tokyo Stock Exchange").setMicCodes("XJPX"),
		newExchange("AT", "Athens Exchange").setMicCodes("ASEX").forceLenient(),
		newExchange("SHE", "Shenzhen Exchange").setMicCodes("XSHE"),
		newExchange("SN", "Chilean Stock Exchange").setMicCodes("XSGO").forceLenient(),
		newExchange("KAR", "Karachi Stock Exchange").setMicCodes("XKAR"),
		newExchange("SHG", "Shanghai Exchange").setMicCodes("XSHG"),
		newExchange("VN", "Vietnam Stocks").setMicCodes("HSTC"),
		newExchange("KLSE", "Kuala Lumpur Exchange").setMicCodes("XKLS"),
		newExchange("SA", "Sao Paolo Exchange").setMicCodes("BVMF"),
		newExchange("MX", "Mexican Exchange").setMicCodes("XMEX"),
		newExchange("IL", "London IL"),
		newExchange("TWO", "Taiwan OTC Exchange").setMicCodes("ROCO"),
		newExchange("TW", "Taiwan Exchange").setMicCodes("XTAI"),

		newExchange("EUFUND", "Europe Fund Virtual Exchange"),
		newExchange("BOND", "Bonds"),
		newExchange("CC", "Cryptocurrencies").setMicCodes("CRYP"),
		newExchange("COMM", "Commodities"),
		newExchange("FOREX", "FOREX").setMicCodes("CDSL"),
		newExchange("INDX", "Indexes"),
	}
}
