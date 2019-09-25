package eodhdapi

import (
	"context"
	"fmt"
	"github.com/gitu/eodhdapi/exchanges"
	"log"
	"strconv"
	"time"
)

func (g *General) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.Code, err = reader.asString(prefix + "Code"); err != nil {
		return err
	}
	if g.Code, err = reader.asString(prefix + "Code"); err != nil {
		return err
	}
	if g.Type, err = reader.asString(prefix + "Type"); err != nil {
		return err
	}
	if g.Name, err = reader.asString(prefix + "Name"); err != nil {
		return err
	}
	if g.Exchange, err = reader.asString(prefix + "Exchange"); err != nil {
		return err
	}
	if g.CurrencyCode, err = reader.asString(prefix + "CurrencyCode"); err != nil {
		return err
	}
	if g.CurrencyName, err = reader.asString(prefix + "CurrencyName"); err != nil {
		return err
	}
	if g.CurrencySymbol, err = reader.asString(prefix + "CurrencySymbol"); err != nil {
		return err
	}
	if g.CountryName, err = reader.asString(prefix + "CountryName"); err != nil {
		return err
	}
	if g.CountryISO, err = reader.asString(prefix + "CountryISO"); err != nil {
		return err
	}
	if g.ISIN, err = reader.asOptionalString(prefix + "ISIN"); err != nil {
		return err
	}
	if g.Sector, err = reader.asString(prefix + "Sector"); err != nil {
		return err
	}
	if g.Industry, err = reader.asString(prefix + "Industry"); err != nil {
		return err
	}
	if g.Description, err = reader.asString(prefix + "Description"); err != nil {
		return err
	}
	if g.FullTimeEmployees, err = reader.asInt(prefix + "FullTimeEmployees"); err != nil {
		return err
	}
	if g.UpdatedAt, err = reader.asString(prefix + "UpdatedAt"); err != nil {
		return err
	}
	return nil
}

func (g *Highlights) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.MarketCapitalization, err = reader.asFloat64(prefix + "MarketCapitalization"); err != nil {
		return err
	}
	if g.MarketCapitalizationMln, err = reader.asString(prefix + "MarketCapitalizationMln"); err != nil {
		return err
	}
	if g.EBITDA, err = reader.asOptionalFloat64(prefix + "EBITDA"); err != nil {
		return err
	}
	if g.PERatio, err = reader.asString(prefix + "PERatio"); err != nil {
		return err
	}
	if g.PEGRatio, err = reader.asOptionalFloat64(prefix + "PEGRatio"); err != nil {
		return err
	}
	if g.WallStreetTargetPrice, err = reader.asOptionalFloat64(prefix + "WallStreetTargetPrice"); err != nil {
		return err
	}
	if g.BookValue, err = reader.asString(prefix + "BookValue"); err != nil {
		return err
	}
	if g.DividendShare, err = reader.asString(prefix + "DividendShare"); err != nil {
		return err
	}
	if g.DividendYield, err = reader.asString(prefix + "DividendYield"); err != nil {
		return err
	}
	if g.EarningsShare, err = reader.asString(prefix + "EarningsShare"); err != nil {
		return err
	}
	if g.EPSEstimateCurrentYear, err = reader.asOptionalFloat64(prefix + "EPSEstimateCurrentYear"); err != nil {
		return err
	}
	if g.EPSEstimateNextYear, err = reader.asOptionalFloat64(prefix + "EPSEstimateNextYear"); err != nil {
		return err
	}
	if g.EPSEstimateNextQuarter, err = reader.asOptionalFloat64(prefix + "EPSEstimateNextQuarter"); err != nil {
		return err
	}
	if g.MostRecentQuarter, err = reader.asString(prefix + "MostRecentQuarter"); err != nil {
		return err
	}
	if g.ProfitMargin, err = reader.asString(prefix + "ProfitMargin"); err != nil {
		return err
	}
	if g.OperatingMarginTTM, err = reader.asString(prefix + "OperatingMarginTTM"); err != nil {
		return err
	}
	if g.ReturnOnAssetsTTM, err = reader.asString(prefix + "ReturnOnAssetsTTM"); err != nil {
		return err
	}
	if g.ReturnOnEquityTTM, err = reader.asString(prefix + "ReturnOnEquityTTM"); err != nil {
		return err
	}
	if g.RevenueTTM, err = reader.asOptionalString(prefix + "RevenueTTM"); err != nil {
		return err
	}
	if g.RevenuePerShareTTM, err = reader.asOptionalString(prefix + "RevenuePerShareTTM"); err != nil {
		return err
	}
	if g.QuarterlyRevenueGrowthYOY, err = reader.asString(prefix + "QuarterlyRevenueGrowthYOY"); err != nil {
		return err
	}
	if g.GrossProfitTTM, err = reader.asString(prefix + "GrossProfitTTM"); err != nil {
		return err
	}
	if g.DilutedEpsTTM, err = reader.asString(prefix + "DilutedEpsTTM"); err != nil {
		return err
	}
	if g.QuarterlyEarningsGrowthYOY, err = reader.asString(prefix + "QuarterlyEarningsGrowthYOY"); err != nil {
		return err
	}
	return nil

}

func (g *Valuation) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.TrailingPE, err = reader.asString(prefix + "TrailingPE"); err != nil {
		return err
	}
	if g.ForwardPE, err = reader.asOptionalString(prefix + "ForwardPE"); err != nil {
		return err
	}
	if g.PriceSalesTTM, err = reader.asOptionalString(prefix + "PriceSalesTTM"); err != nil {
		return err
	}
	if g.PriceBookMRQ, err = reader.asString(prefix + "PriceBookMRQ"); err != nil {
		return err
	}
	if g.EnterpriseValueRevenue, err = reader.asOptionalString(prefix + "EnterpriseValueRevenue"); err != nil {
		return err
	}
	if g.EnterpriseValueEbitda, err = reader.asOptionalString(prefix + "EnterpriseValueEbitda"); err != nil {
		return err
	}
	return nil
}

func (g *Technicals) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.Beta, err = reader.asString(prefix + "Beta"); err != nil {
		return err
	}
	if g.FiftyTwoWeekHigh, err = reader.asString(prefix + "52WeekHigh"); err != nil {
		return err
	}
	if g.FiftyTwoWeekLow, err = reader.asString(prefix + "52WeekLow"); err != nil {
		return err
	}
	if g.FiftyDayMA, err = reader.asString(prefix + "50DayMA"); err != nil {
		return err
	}
	if g.TwoHundredDayMA, err = reader.asString(prefix + "200DayMA"); err != nil {
		return err
	}
	if g.SharesShort, err = reader.asOptionalFloat64(prefix + "SharesShort"); err != nil {
		return err
	}
	if g.SharesShortPriorMonth, err = reader.asOptionalFloat64(prefix + "SharesShortPriorMonth"); err != nil {
		return err
	}
	if g.ShortRatio, err = reader.asOptionalFloat64(prefix + "ShortRatio"); err != nil {
		return err
	}
	if g.ShortPercent, err = reader.asOptionalFloat64(prefix + "ShortPercent"); err != nil {
		return err
	}
	return nil
}

func (g *SplitsDividends) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.ForwardAnnualDividendRate, err = reader.asString(prefix + "ForwardAnnualDividendRate"); err != nil {
		return err
	}
	if g.ForwardAnnualDividendYield, err = reader.asString(prefix + "ForwardAnnualDividendYield"); err != nil {
		return err
	}
	if g.PayoutRatio, err = reader.asString(prefix + "PayoutRatio"); err != nil {
		return err
	}
	if g.DividendDate, err = reader.asString(prefix + "DividendDate"); err != nil {
		return err
	}
	if g.ExDividendDate, err = reader.asString(prefix + "ExDividendDate"); err != nil {
		return err
	}
	if g.LastSplitFactor, err = reader.asString(prefix + "LastSplitFactor"); err != nil {
		return err
	}
	if g.LastSplitDate, err = reader.asString(prefix + "LastSplitDate"); err != nil {
		return err
	}
	return nil
}

func (g *Earnings) fill(reader *csvReaderMap, prefix string) error {
	g.Last0 = EarningsInfo{}
	g.Last1 = EarningsInfo{}
	g.Last2 = EarningsInfo{}
	g.Last3 = EarningsInfo{}
	return nil
}

func (g *EarningsInfo) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.Date, err = reader.asString(prefix + "Date"); err != nil {
		return err
	}
	if g.EpsActual, err = reader.asFloat64(prefix + "EpsActual"); err != nil {
		return err
	}
	if g.EpsEstimate, err = reader.asFloat64(prefix + "EpsEstimate"); err != nil {
		return err
	}
	if g.EpsDifference, err = reader.asFloat64(prefix + "EpsDifference"); err != nil {
		return err
	}
	if g.SurprisePercent, err = reader.asFloat64(prefix + "SurprisePercent"); err != nil {
		return err
	}
	return nil
}

// FetchFundamentals Fetches Fundamentals for the exchange
func (d *EODhd) FetchFundamentals(ctx context.Context, fundamentals chan Fundamentals, exchange *exchanges.Exchange, pagesize int, lenient bool) error {

	for _, e := range exchange.ExchangeCodeComponents {

		offset := 0

		newElements := pagesize
		for newElements == pagesize {
			newElements = 0
			res, err := d.readPath("/bulk-fundamentals/"+e,
				urlParam{"fmt", "csv"},
				urlParam{"offset", strconv.Itoa(offset)},
				urlParam{"limit", strconv.Itoa(pagesize)})

			if err != nil {
				return err
			}

			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Printf("body for url: %s - code %d: %v\n", res.Request.URL, res.StatusCode, res.Body)
				return fmt.Errorf("received non 200 error code: %d", res.StatusCode)
			}

			reader, err := newCsvReaderMap(res.Body, lenient, lenient)
			if err != nil {
				return err
			}
			for reader.Next() {

				f := Fundamentals{
					LastUpdate:      time.Now(),
					General:         General{},
					Highlights:      Highlights{},
					Valuation:       Valuation{},
					Technicals:      Technicals{},
					SplitsDividends: SplitsDividends{},
					Earnings:        Earnings{},
					Financials:      Financials{},
				}
				err = f.General.fill(reader, "General_")
				if err != nil {
					return err
				}

				err = f.Highlights.fill(reader, "Highlights_")
				if err != nil {
					return err
				}
				err = f.Valuation.fill(reader, "Valuation_")
				if err != nil {
					return err
				}
				err = f.Technicals.fill(reader, "Technicals_")
				if err != nil {
					return err
				}
				err = f.SplitsDividends.fill(reader, "SplitsDividends_")
				if err != nil {
					return err
				}
				err = f.Earnings.fill(reader, "Earnings_")
				if err != nil {
					return err
				}

				f.Ticker = f.General.Code + "." + exchange.Code

				fundamentals <- f

				newElements++
			}

			if !lenient {
				err = reader.checkAllVisited()
				if err != nil {
					return err
				}
			}
			offset += newElements
		}
	}

	return nil
}

// Fundamentals for a ticker
type Fundamentals struct {
	LastUpdate      time.Time
	Ticker          string
	General         General         `json:"General"`
	Highlights      Highlights      `json:"Highlights"`
	Valuation       Valuation       `json:"Valuation"`
	Technicals      Technicals      `json:"Technicals"`
	SplitsDividends SplitsDividends `json:"SplitsDividends"`
	Earnings        Earnings        `json:"Earnings"`
	Financials      Financials      `json:"Financials"`
}

type General struct {
	Code              string  `json:"Code"`
	Type              string  `json:"Type"`
	Name              string  `json:"Name"`
	Exchange          string  `json:"Exchange"`
	CurrencyCode      string  `json:"CurrencyCode"`
	CurrencyName      string  `json:"CurrencyName"`
	CurrencySymbol    string  `json:"CurrencySymbol"`
	CountryName       string  `json:"CountryName"`
	CountryISO        string  `json:"CountryISO"`
	ISIN              *string `json:"ISIN"`
	Sector            string  `json:"Sector"`
	Industry          string  `json:"Industry"`
	Description       string  `json:"Description"`
	FullTimeEmployees int     `json:"FullTimeEmployees"`
	UpdatedAt         string  `json:"UpdatedAt"`
}
type Highlights struct {
	MarketCapitalization       float64  `json:"MarketCapitalization"`
	MarketCapitalizationMln    string   `json:"MarketCapitalizationMln"`
	EBITDA                     *float64 `json:"EBITDA"`
	PERatio                    string   `json:"PERatio"`
	PEGRatio                   *float64 `json:"PEGRatio"`
	WallStreetTargetPrice      *float64 `json:"WallStreetTargetPrice"`
	BookValue                  string   `json:"BookValue"`
	DividendShare              string   `json:"DividendShare"`
	DividendYield              string   `json:"DividendYield"`
	EarningsShare              string   `json:"EarningsShare"`
	EPSEstimateCurrentYear     *float64 `json:"EPSEstimateCurrentYear"`
	EPSEstimateNextYear        *float64 `json:"EPSEstimateNextYear"`
	EPSEstimateNextQuarter     *float64 `json:"EPSEstimateNextQuarter"`
	MostRecentQuarter          string   `json:"MostRecentQuarter"`
	ProfitMargin               string   `json:"ProfitMargin"`
	OperatingMarginTTM         string   `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          string   `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          string   `json:"ReturnOnEquityTTM"`
	RevenueTTM                 *string  `json:"RevenueTTM"`
	RevenuePerShareTTM         *string  `json:"RevenuePerShareTTM"`
	QuarterlyRevenueGrowthYOY  string   `json:"QuarterlyRevenueGrowthYOY"`
	GrossProfitTTM             string   `json:"GrossProfitTTM"`
	DilutedEpsTTM              string   `json:"DilutedEpsTTM"`
	QuarterlyEarningsGrowthYOY string   `json:"QuarterlyEarningsGrowthYOY"`
}
type Valuation struct {
	TrailingPE             string  `json:"TrailingPE"`
	ForwardPE              *string `json:"ForwardPE"`
	PriceSalesTTM          *string `json:"PriceSalesTTM"`
	PriceBookMRQ           string  `json:"PriceBookMRQ"`
	EnterpriseValueRevenue *string `json:"EnterpriseValueRevenue"`
	EnterpriseValueEbitda  *string `json:"EnterpriseValueEbitda"`
}
type Technicals struct {
	Beta                  string   `json:"Beta"`
	FiftyTwoWeekHigh      string   `json:"52WeekHigh"`
	FiftyTwoWeekLow       string   `json:"52WeekLow"`
	FiftyDayMA            string   `json:"50DayMA"`
	TwoHundredDayMA       string   `json:"200DayMA"`
	SharesShort           *float64 `json:"SharesShort"`
	SharesShortPriorMonth *float64 `json:"SharesShortPriorMonth"`
	ShortRatio            *float64 `json:"ShortRatio"`
	ShortPercent          *float64 `json:"ShortPercent"`
}
type SplitsDividends struct {
	ForwardAnnualDividendRate  string `json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield string `json:"ForwardAnnualDividendYield"`
	PayoutRatio                string `json:"PayoutRatio"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
	LastSplitFactor            string `json:"LastSplitFactor"`
	LastSplitDate              string `json:"LastSplitDate"`
}
type EarningsInfo struct {
	Date            string  `json:"date"`
	EpsActual       float64 `json:"epsActual"`
	EpsEstimate     float64 `json:"epsEstimate"`
	EpsDifference   float64 `json:"epsDifference"`
	SurprisePercent float64 `json:"surprisePercent"`
}
type Earnings struct {
	Last0 EarningsInfo `json:"Last_0"`
	Last1 EarningsInfo `json:"Last_1"`
	Last2 EarningsInfo `json:"Last_2"`
	Last3 EarningsInfo `json:"Last_3"`
}
type BalanceSheetInfo struct {
	Date                    string   `json:"date"`
	FilingDate              *string  `json:"filing_date"`
	IntangibleAssets        *float64 `json:"intangibleAssets"`
	TotalLiab               string   `json:"totalLiab"`
	TotalStockholderEquity  string   `json:"totalStockholderEquity"`
	DeferredLongTermLiab    *string  `json:"deferredLongTermLiab"`
	OtherCurrentLiab        string   `json:"otherCurrentLiab"`
	TotalAssets             string   `json:"totalAssets"`
	CommonStock             string   `json:"commonStock"`
	OtherCurrentAssets      string   `json:"otherCurrentAssets"`
	RetainedEarnings        string   `json:"retainedEarnings"`
	OtherLiab               string   `json:"otherLiab"`
	GoodWill                string   `json:"goodWill"`
	OtherAssets             string   `json:"otherAssets"`
	Cash                    string   `json:"cash"`
	TotalCurrentLiabilities string   `json:"totalCurrentLiabilities"`
	ShortLongTermDebt       *string  `json:"shortLongTermDebt"`
	OtherStockholderEquity  string   `json:"otherStockholderEquity"`
	PropertyPlantEquipment  string   `json:"propertyPlantEquipment"`
	TotalCurrentAssets      string   `json:"totalCurrentAssets"`
	LongTermInvestments     string   `json:"longTermInvestments"`
	NetTangibleAssets       string   `json:"netTangibleAssets"`
	ShortTermInvestments    *string  `json:"shortTermInvestments"`
	NetReceivables          string   `json:"netReceivables"`
	LongTermDebt            *string  `json:"longTermDebt"`
	Inventory               string   `json:"inventory"`
	AccountsPayable         *string  `json:"accountsPayable"`
}
type BalanceSheet struct {
	CurrencySymbol string           `json:"currency_symbol"`
	QuarterlyLast0 BalanceSheetInfo `json:"quarterly_last_0"`
	QuarterlyLast1 BalanceSheetInfo `json:"quarterly_last_1"`
	QuarterlyLast2 BalanceSheetInfo `json:"quarterly_last_2"`
	QuarterlyLast3 BalanceSheetInfo `json:"quarterly_last_3"`
	YearlyLast0    BalanceSheetInfo `json:"yearly_last_0"`
	YearlyLast1    BalanceSheetInfo `json:"yearly_last_1"`
	YearlyLast2    BalanceSheetInfo `json:"yearly_last_2"`
	YearlyLast3    BalanceSheetInfo `json:"yearly_last_3"`
}
type CashFlowInfo struct {
	Date                                  string  `json:"date"`
	FilingDate                            *string `json:"filing_date"`
	Investments                           string  `json:"investments"`
	ChangeToLiabilities                   string  `json:"changeToLiabilities"`
	TotalCashflowsFromInvestingActivities string  `json:"totalCashflowsFromInvestingActivities"`
	NetBorrowings                         string  `json:"netBorrowings"`
	TotalCashFromFinancingActivities      string  `json:"totalCashFromFinancingActivities"`
	ChangeToOperatingActivities           string  `json:"changeToOperatingActivities"`
	NetIncome                             string  `json:"netIncome"`
	ChangeInCash                          string  `json:"changeInCash"`
	TotalCashFromOperatingActivities      string  `json:"totalCashFromOperatingActivities"`
	Depreciation                          string  `json:"depreciation"`
	OtherCashflowsFromInvestingActivities string  `json:"otherCashflowsFromInvestingActivities"`
	DividendsPaid                         string  `json:"dividendsPaid"`
	ChangeToInventory                     string  `json:"changeToInventory"`
	ChangeToAccountReceivables            string  `json:"changeToAccountReceivables"`
	SalePurchaseOfStock                   *string `json:"salePurchaseOfStock"`
	OtherCashflowsFromFinancingActivities *string `json:"otherCashflowsFromFinancingActivities"`
	ChangeToNetincome                     string  `json:"changeToNetincome"`
	CapitalExpenditures                   string  `json:"capitalExpenditures"`
}
type CashFlow struct {
	CurrencySymbol string       `json:"currency_symbol"`
	QuarterlyLast0 CashFlowInfo `json:"quarterly_last_0"`
	QuarterlyLast1 CashFlowInfo `json:"quarterly_last_1"`
	QuarterlyLast2 CashFlowInfo `json:"quarterly_last_2"`
	QuarterlyLast3 CashFlowInfo `json:"quarterly_last_3"`
	YearlyLast0    CashFlowInfo `json:"yearly_last_0"`
	YearlyLast1    CashFlowInfo `json:"yearly_last_1"`
	YearlyLast2    CashFlowInfo `json:"yearly_last_2"`
	YearlyLast3    CashFlowInfo `json:"yearly_last_3"`
}

type IncomeStatementInfo struct {
	Date                              string      `json:"date"`
	FilingDate                        interface{} `json:"filing_date"`
	ResearchDevelopment               interface{} `json:"researchDevelopment"`
	EffectOfAccountingCharges         interface{} `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   string      `json:"incomeBeforeTax"`
	MinorityInterest                  string      `json:"minorityInterest"`
	NetIncome                         string      `json:"netIncome"`
	SellingGeneralAdministrative      string      `json:"sellingGeneralAdministrative"`
	GrossProfit                       string      `json:"grossProfit"`
	Ebit                              string      `json:"ebit"`
	OperatingIncome                   string      `json:"operatingIncome"`
	OtherOperatingExpenses            *string     `json:"otherOperatingExpenses"`
	InterestExpense                   *string     `json:"interestExpense"`
	ExtraordinaryItems                *string     `json:"extraordinaryItems"`
	NonRecurring                      *string     `json:"nonRecurring"`
	OtherItems                        *string     `json:"otherItems"`
	IncomeTaxExpense                  string      `json:"incomeTaxExpense"`
	TotalRevenue                      string      `json:"totalRevenue"`
	TotalOperatingExpenses            string      `json:"totalOperatingExpenses"`
	CostOfRevenue                     string      `json:"costOfRevenue"`
	TotalOtherIncomeExpenseNet        string      `json:"totalOtherIncomeExpenseNet"`
	DiscontinuedOperations            *string     `json:"discontinuedOperations"`
	NetIncomeFromContinuingOps        string      `json:"netIncomeFromContinuingOps"`
	NetIncomeApplicableToCommonShares string      `json:"netIncomeApplicableToCommonShares"`
}
type IncomeStatement struct {
	CurrencySymbol string              `json:"currency_symbol"`
	QuarterlyLast0 IncomeStatementInfo `json:"quarterly_last_0"`
	QuarterlyLast1 IncomeStatementInfo `json:"quarterly_last_1"`
	QuarterlyLast2 IncomeStatementInfo `json:"quarterly_last_2"`
	QuarterlyLast3 IncomeStatementInfo `json:"quarterly_last_3"`
	YearlyLast0    IncomeStatementInfo `json:"yearly_last_0"`
	YearlyLast1    IncomeStatementInfo `json:"yearly_last_1"`
	YearlyLast2    IncomeStatementInfo `json:"yearly_last_2"`
	YearlyLast3    IncomeStatementInfo `json:"yearly_last_3"`
}
type Financials struct {
	BalanceSheet    BalanceSheet    `json:"Balance_Sheet"`
	CashFlow        CashFlow        `json:"Cash_Flow"`
	IncomeStatement IncomeStatement `json:"Income_Statement"`
}
