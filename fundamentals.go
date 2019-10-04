package eodhdapi

//go:generate go run github.com/mailru/easyjson/easyjson -all fundamentals.go

import (
	"context"
	"fmt"
	"github.com/gitu/eodhdapi/exchanges"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
	"time"
)

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
	Sector            string  `json:"Sector"`
	Industry          string  `json:"Industry"`
	Description       string  `json:"Description"`
	ISIN              *string `json:"ISIN"`
	FullTimeEmployees *int    `json:"FullTimeEmployees"`
	UpdatedAt         *string `json:"UpdatedAt"`
	Cusip             *string `json:"CUSIP"`
	LogoURL           *string `json:"LogoURL"`

	CIK                   *string `json:"CIK"`
	EmployerIDNumber      *string `json:"EmployerIdNumber"`
	FiscalYearEnd         *string `json:"FiscalYearEnd"`
	IPODate               *string `json:"IPODate"`
	InternationalDomestic *string `json:"InternationalDomestic"`
	GicSector             *string `json:"GicSector"`
	GicGroup              *string `json:"GicGroup"`
	GicIndustry           *string `json:"GicIndustry"`
	GicSubIndustry        *string `json:"GicSubIndustry"`
	Address               *string `json:"Address"`
	Phone                 *string `json:"Phone"`
	WebURL                *string `json:"WebURL"`
}
type Highlights struct {
	MarketCapitalization       *float64 `json:"MarketCapitalization"`
	MarketCapitalizationMln    string   `json:"MarketCapitalizationMln"`
	EBITDA                     *float64 `json:"EBITDA"`
	PERatio                    *float64 `json:"PERatio"`
	PEGRatio                   *float64 `json:"PEGRatio"`
	WallStreetTargetPrice      *float64 `json:"WallStreetTargetPrice"`
	BookValue                  *float64 `json:"BookValue"`
	DividendShare              *float64 `json:"DividendShare"`
	DividendYield              *float64 `json:"DividendYield"`
	EarningsShare              *float64 `json:"EarningsShare"`
	EPSEstimateCurrentYear     *float64 `json:"EPSEstimateCurrentYear"`
	EPSEstimateNextYear        *float64 `json:"EPSEstimateNextYear"`
	EPSEstimateNextQuarter     *float64 `json:"EPSEstimateNextQuarter"`
	MostRecentQuarter          string   `json:"MostRecentQuarter"`
	ProfitMargin               *float64 `json:"ProfitMargin"`
	OperatingMarginTTM         *float64 `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          *float64 `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          *float64 `json:"ReturnOnEquityTTM"`
	RevenueTTM                 *float64 `json:"RevenueTTM"`
	RevenuePerShareTTM         *float64 `json:"RevenuePerShareTTM"`
	QuarterlyRevenueGrowthYOY  *float64 `json:"QuarterlyRevenueGrowthYOY"`
	GrossProfitTTM             *float64 `json:"GrossProfitTTM"`
	DilutedEpsTTM              *float64 `json:"DilutedEpsTTM"`
	QuarterlyEarningsGrowthYOY *float64 `json:"QuarterlyEarningsGrowthYOY"`
}

type ETFData struct {
	ISIN                    string               `json:"ISIN"`
	CompanyName             string               `json:"Company_Name"`
	CompanyURL              string               `json:"Company_URL"`
	ETFURL                  string               `json:"ETF_URL"`
	Yield                   string               `json:"Yield"`
	DividendPayingFrequency string               `json:"Dividend_Paying_Frequency"`
	InceptionDate           string               `json:"Inception_Date"`
	MaxAnnualMgmtCharge     string               `json:"Max_Annual_Mgmt_Charge"`
	OngoingCharge           string               `json:"Ongoing_Charge"`
	DateOngoingCharge       string               `json:"Date_Ongoing_Charge"`
	NetExpenseRatio         string               `json:"NetExpenseRatio"`
	AnnualHoldingsTurnover  string               `json:"AnnualHoldingsTurnover"`
	TotalAssets             string               `json:"TotalAssets"`
	AverageMktCapMil        string               `json:"Average_Mkt_Cap_Mil"`
	MarketCapitalisation    MarketCapitalisation `json:"Market_Capitalisation"`
	AssetAllocation         AssetAllocation      `json:"Asset_Allocation"`
	WorldRegions            WorldRegions         `json:"World_Regions"`
	SectorWeights           SectorWeights        `json:"Sector_Weights"`
	Top10Holdings           Top10Holdings        `json:"Top_10_Holdings"`
	Holdings                Holdings             `json:"Holdings"`
	ValuationsGrowth        ValuationsGrowth     `json:"Valuations_Growth"`
	MorningStar             MorningStar          `json:"MorningStar"`
	Performance             Performance          `json:"Performance"`
}
type Valuation struct {
	TrailingPE             *float64 `json:"TrailingPE"`
	ForwardPE              *float64 `json:"ForwardPE"`
	PriceSalesTTM          *float64 `json:"PriceSalesTTM"`
	PriceBookMRQ           *float64 `json:"PriceBookMRQ"`
	EnterpriseValueRevenue *float64 `json:"EnterpriseValueRevenue"`
	EnterpriseValueEbitda  *float64 `json:"EnterpriseValueEbitda"`
}
type Technicals struct {
	Beta                  *float64 `json:"Beta"`
	FiftyTwoWeekHigh      *float64 `json:"52WeekHigh"`
	FiftyTwoWeekLow       *float64 `json:"52WeekLow"`
	FiftyDayMA            *float64 `json:"50DayMA"`
	TwoHundredDayMA       *float64 `json:"200DayMA"`
	SharesShort           *float64 `json:"SharesShort"`
	SharesShortPriorMonth *float64 `json:"SharesShortPriorMonth"`
	ShortRatio            *float64 `json:"ShortRatio"`
	ShortPercent          *float64 `json:"ShortPercent"`
}
type SplitsDividends struct {
	ForwardAnnualDividendRate  *float64 `json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield *float64 `json:"ForwardAnnualDividendYield"`
	PayoutRatio                *float64 `json:"PayoutRatio"`
	DividendDate               string   `json:"DividendDate"`
	ExDividendDate             string   `json:"ExDividendDate"`
	LastSplitFactor            string   `json:"LastSplitFactor"`
	LastSplitDate              string   `json:"LastSplitDate"`
}
type EarningsInfo struct {
	Date            string   `json:"date"`
	EpsActual       *float64 `json:"epsActual"`
	EpsEstimate     *float64 `json:"epsEstimate"`
	EpsDifference   *float64 `json:"epsDifference"`
	SurprisePercent *float64 `json:"surprisePercent"`
}
type Earnings struct {
	Last0 EarningsInfo `json:"Last_0"`
	Last1 EarningsInfo `json:"Last_1"`
	Last2 EarningsInfo `json:"Last_2"`
	Last3 EarningsInfo `json:"Last_3"`
}
type BalanceSheetInfo struct {
	Date                         string   `json:"date"`
	FilingDate                   *string  `json:"filing_date"`
	IntangibleAssets             *float64 `json:"intangibleAssets"`
	TotalLiab                    *float64 `json:"totalLiab"`
	TotalStockholderEquity       *float64 `json:"totalStockholderEquity"`
	DeferredLongTermLiab         *float64 `json:"deferredLongTermLiab"`
	OtherCurrentLiab             *float64 `json:"otherCurrentLiab"`
	TotalAssets                  *float64 `json:"totalAssets"`
	CommonStock                  *float64 `json:"commonStock"`
	CommonStockSharesOutstanding *float64 `json:"commonStockSharesOutStanding"`
	OtherCurrentAssets           *float64 `json:"otherCurrentAssets"`
	RetainedEarnings             *float64 `json:"retainedEarnings"`
	OtherLiab                    *float64 `json:"otherLiab"`
	GoodWill                     *float64 `json:"goodWill"`
	OtherAssets                  *float64 `json:"otherAssets"`
	Cash                         *float64 `json:"cash"`
	TotalCurrentLiabilities      *float64 `json:"totalCurrentLiabilities"`
	ShortLongTermDebt            *float64 `json:"shortLongTermDebt"`
	OtherStockholderEquity       *float64 `json:"otherStockholderEquity"`
	PropertyPlantEquipment       *float64 `json:"propertyPlantEquipment"`
	TotalCurrentAssets           *float64 `json:"totalCurrentAssets"`
	LongTermInvestments          *float64 `json:"longTermInvestments"`
	NetTangibleAssets            *float64 `json:"netTangibleAssets"`
	ShortTermInvestments         *float64 `json:"shortTermInvestments"`
	NetReceivables               *float64 `json:"netReceivables"`
	LongTermDebt                 *float64 `json:"longTermDebt"`
	Inventory                    *float64 `json:"inventory"`
	AccountsPayable              *float64 `json:"accountsPayable"`
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
	Date                                  string   `json:"date"`
	FilingDate                            *string  `json:"filing_date"`
	Investments                           *float64 `json:"investments"`
	ChangeToLiabilities                   *float64 `json:"changeToLiabilities"`
	TotalCashflowsFromInvestingActivities *float64 `json:"totalCashflowsFromInvestingActivities"`
	NetBorrowings                         *float64 `json:"netBorrowings"`
	TotalCashFromFinancingActivities      *float64 `json:"totalCashFromFinancingActivities"`
	ChangeToOperatingActivities           *float64 `json:"changeToOperatingActivities"`
	NetIncome                             *float64 `json:"netIncome"`
	ChangeInCash                          *float64 `json:"changeInCash"`
	TotalCashFromOperatingActivities      *float64 `json:"totalCashFromOperatingActivities"`
	Depreciation                          *float64 `json:"depreciation"`
	OtherCashflowsFromInvestingActivities *float64 `json:"otherCashflowsFromInvestingActivities"`
	DividendsPaid                         *float64 `json:"dividendsPaid"`
	ChangeToInventory                     *float64 `json:"changeToInventory"`
	ChangeToAccountReceivables            *float64 `json:"changeToAccountReceivables"`
	SalePurchaseOfStock                   *float64 `json:"salePurchaseOfStock"`
	OtherCashflowsFromFinancingActivities *float64 `json:"otherCashflowsFromFinancingActivities"`
	ChangeToNetincome                     *float64 `json:"changeToNetincome"`
	CapitalExpenditures                   *float64 `json:"capitalExpenditures"`
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
	Date                              string   `json:"date"`
	FilingDate                        *string  `json:"filing_date"`
	ResearchDevelopment               *float64 `json:"researchDevelopment"`
	EffectOfAccountingCharges         *float64 `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   *float64 `json:"incomeBeforeTax"`
	MinorityInterest                  *float64 `json:"minorityInterest"`
	NetIncome                         *float64 `json:"netIncome"`
	SellingGeneralAdministrative      *float64 `json:"sellingGeneralAdministrative"`
	GrossProfit                       *float64 `json:"grossProfit"`
	Ebit                              *float64 `json:"ebit"`
	OperatingIncome                   *float64 `json:"operatingIncome"`
	OtherOperatingExpenses            *float64 `json:"otherOperatingExpenses"`
	InterestExpense                   *float64 `json:"interestExpense"`
	ExtraordinaryItems                *float64 `json:"extraordinaryItems"`
	NonRecurring                      *float64 `json:"nonRecurring"`
	OtherItems                        *float64 `json:"otherItems"`
	IncomeTaxExpense                  *float64 `json:"incomeTaxExpense"`
	TotalRevenue                      *float64 `json:"totalRevenue"`
	TotalOperatingExpenses            *float64 `json:"totalOperatingExpenses"`
	CostOfRevenue                     *float64 `json:"costOfRevenue"`
	TotalOtherIncomeExpenseNet        *float64 `json:"totalOtherIncomeExpenseNet"`
	DiscontinuedOperations            *float64 `json:"discontinuedOperations"`
	NetIncomeFromContinuingOps        *float64 `json:"netIncomeFromContinuingOps"`
	NetIncomeApplicableToCommonShares *float64 `json:"netIncomeApplicableToCommonShares"`
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

// FetchFundamentals Fetches Fundamentals for the exchange
func (d *EODhd) FetchFundamentals(ctx context.Context, fundamentals chan Fundamentals, exchange *exchanges.Exchange, pagesize int, lenient bool) error {

	if exchange.ForceLenient {
		lenient = true
	}
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
				log.Printf("body for url: %s - code %d: %v\n", strings.ReplaceAll(res.Request.URL.String(), d.token, "******"), res.StatusCode, res.Body)
				return fmt.Errorf("received non 200 error code: %d", res.StatusCode)
			}

			reader, err := newCsvReaderMap(res.Body, lenient, !lenient)
			if err != nil {
				return err
			}
			for reader.Next() {
				f, err := buildFundamental(reader, exchange)
				if err != nil {
					if !lenient {
						return errors.Wrap(err, fmt.Sprintf("while parsing line: %.50s", strings.Join(reader.current, ",")))
					}
					log.Println(err, strings.Join(reader.current, ","))
					continue
				}

				fundamentals <- f

				if reader.trackVisits {
					// skip tracking after first visit
					reader.trackVisits = false
				}

				newElements++
			}

			if !lenient && newElements > 0 {
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

func buildFundamental(reader *csvReaderMap, exchange *exchanges.Exchange) (Fundamentals, error) {
	var err error
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
		return Fundamentals{}, err
	}
	err = f.Highlights.fill(reader, "Highlights_")
	if err != nil {
		return Fundamentals{}, err
	}
	err = f.Valuation.fill(reader, "Valuation_")
	if err != nil {
		return Fundamentals{}, err
	}
	err = f.Technicals.fill(reader, "Technicals_")
	if err != nil {
		return Fundamentals{}, err
	}
	err = f.SplitsDividends.fill(reader, "SplitsDividends_")
	if err != nil {
		return Fundamentals{}, err
	}
	err = f.Earnings.fill(reader, "Earnings_")
	if err != nil {
		return Fundamentals{}, err
	}
	err = f.Financials.fill(reader, "Financials_")
	if err != nil {
		return Fundamentals{}, err
	}
	f.Ticker = f.General.Code + "." + exchange.Code
	return f, err
}

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
	if g.FullTimeEmployees, err = reader.asOptionalInt(prefix + "FullTimeEmployees"); err != nil {
		return err
	}
	if g.UpdatedAt, err = reader.asOptionalStringLenient(prefix + "UpdatedAt"); err != nil {
		return err
	}

	if g.Cusip, err = reader.asOptionalStringLenient(prefix + "CUSIP"); err != nil {
		return err
	}
	return nil
}
func (g *Highlights) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.MarketCapitalization, err = reader.asOptionalFloat64(prefix + "MarketCapitalization"); err != nil {
		return err
	}
	if g.MarketCapitalizationMln, err = reader.asString(prefix + "MarketCapitalizationMln"); err != nil {
		return err
	}
	if g.EBITDA, err = reader.asOptionalFloat64(prefix + "EBITDA"); err != nil {
		return err
	}
	if g.PERatio, err = reader.asOptionalFloat64(prefix + "PERatio"); err != nil {
		return err
	}
	if g.PEGRatio, err = reader.asOptionalFloat64(prefix + "PEGRatio"); err != nil {
		return err
	}
	if g.WallStreetTargetPrice, err = reader.asOptionalFloat64(prefix + "WallStreetTargetPrice"); err != nil {
		return err
	}
	if g.BookValue, err = reader.asOptionalFloat64(prefix + "BookValue"); err != nil {
		return err
	}
	if g.DividendShare, err = reader.asOptionalFloat64(prefix + "DividendShare"); err != nil {
		return err
	}
	if g.DividendYield, err = reader.asOptionalFloat64(prefix + "DividendYield"); err != nil {
		return err
	}
	if g.EarningsShare, err = reader.asOptionalFloat64(prefix + "EarningsShare"); err != nil {
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
	if g.ProfitMargin, err = reader.asOptionalFloat64(prefix + "ProfitMargin"); err != nil {
		return err
	}
	if g.OperatingMarginTTM, err = reader.asOptionalFloat64(prefix + "OperatingMarginTTM"); err != nil {
		return err
	}
	if g.ReturnOnAssetsTTM, err = reader.asOptionalFloat64(prefix + "ReturnOnAssetsTTM"); err != nil {
		return err
	}
	if g.ReturnOnEquityTTM, err = reader.asOptionalFloat64(prefix + "ReturnOnEquityTTM"); err != nil {
		return err
	}
	if g.RevenueTTM, err = reader.asOptionalFloat64(prefix + "RevenueTTM"); err != nil {
		return err
	}
	if g.RevenuePerShareTTM, err = reader.asOptionalFloat64(prefix + "RevenuePerShareTTM"); err != nil {
		return err
	}
	if g.QuarterlyRevenueGrowthYOY, err = reader.asOptionalFloat64(prefix + "QuarterlyRevenueGrowthYOY"); err != nil {
		return err
	}
	if g.GrossProfitTTM, err = reader.asOptionalFloat64(prefix + "GrossProfitTTM"); err != nil {
		return err
	}
	if g.DilutedEpsTTM, err = reader.asOptionalFloat64(prefix + "DilutedEpsTTM"); err != nil {
		return err
	}
	if g.QuarterlyEarningsGrowthYOY, err = reader.asOptionalFloat64(prefix + "QuarterlyEarningsGrowthYOY"); err != nil {
		return err
	}
	return nil

}
func (g *Valuation) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.TrailingPE, err = reader.asOptionalFloat64(prefix + "TrailingPE"); err != nil {
		return err
	}
	if g.ForwardPE, err = reader.asOptionalFloat64(prefix + "ForwardPE"); err != nil {
		return err
	}
	if g.PriceSalesTTM, err = reader.asOptionalFloat64(prefix + "PriceSalesTTM"); err != nil {
		return err
	}
	if g.PriceBookMRQ, err = reader.asOptionalFloat64(prefix + "PriceBookMRQ"); err != nil {
		return err
	}
	if g.EnterpriseValueRevenue, err = reader.asOptionalFloat64(prefix + "EnterpriseValueRevenue"); err != nil {
		return err
	}
	if g.EnterpriseValueEbitda, err = reader.asOptionalFloat64(prefix + "EnterpriseValueEbitda"); err != nil {
		return err
	}
	return nil
}
func (g *Technicals) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.Beta, err = reader.asOptionalFloat64(prefix + "Beta"); err != nil {
		return err
	}
	if g.FiftyTwoWeekHigh, err = reader.asOptionalFloat64(prefix + "52WeekHigh"); err != nil {
		return err
	}
	if g.FiftyTwoWeekLow, err = reader.asOptionalFloat64(prefix + "52WeekLow"); err != nil {
		return err
	}
	if g.FiftyDayMA, err = reader.asOptionalFloat64(prefix + "50DayMA"); err != nil {
		return err
	}
	if g.TwoHundredDayMA, err = reader.asOptionalFloat64(prefix + "200DayMA"); err != nil {
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
	if g.ForwardAnnualDividendRate, err = reader.asOptionalFloat64(prefix + "ForwardAnnualDividendRate"); err != nil {
		return err
	}
	if g.ForwardAnnualDividendYield, err = reader.asOptionalFloat64(prefix + "ForwardAnnualDividendYield"); err != nil {
		return err
	}
	if g.PayoutRatio, err = reader.asOptionalFloat64(prefix + "PayoutRatio"); err != nil {
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
	var err error
	if g.Last0, err = buildEarningsInfo(reader, prefix+"Last_0_"); err != nil {
		return err
	}
	if g.Last1, err = buildEarningsInfo(reader, prefix+"Last_1_"); err != nil {
		return err
	}
	if g.Last2, err = buildEarningsInfo(reader, prefix+"Last_2_"); err != nil {
		return err
	}
	if g.Last3, err = buildEarningsInfo(reader, prefix+"Last_3_"); err != nil {
		return err
	}
	return nil
}
func (g *Financials) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.BalanceSheet, err = buildBalanceSheet(reader, prefix+"Balance_Sheet_"); err != nil {
		return err
	}
	if g.CashFlow, err = buildCashFlow(reader, prefix+"Cash_Flow_"); err != nil {
		return err
	}
	if g.IncomeStatement, err = buildIncomeStatement(reader, prefix+"Income_Statement_"); err != nil {
		return err
	}
	return nil
}

func buildBalanceSheet(reader *csvReaderMap, prefix string) (BalanceSheet, error) {
	var err error
	g := BalanceSheet{}
	if g.CurrencySymbol, err = reader.asString(prefix + "currency_symbol"); err != nil {
		return g, err
	}
	if g.QuarterlyLast0, err = buildBalanceSheetInfo(reader, prefix+"quarterly_last_0_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast1, err = buildBalanceSheetInfo(reader, prefix+"quarterly_last_1_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast2, err = buildBalanceSheetInfo(reader, prefix+"quarterly_last_2_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast3, err = buildBalanceSheetInfo(reader, prefix+"quarterly_last_3_"); err != nil {
		return g, err
	}
	if g.YearlyLast0, err = buildBalanceSheetInfo(reader, prefix+"yearly_last_0_"); err != nil {
		return g, err
	}
	if g.YearlyLast1, err = buildBalanceSheetInfo(reader, prefix+"yearly_last_1_"); err != nil {
		return g, err
	}
	if g.YearlyLast2, err = buildBalanceSheetInfo(reader, prefix+"yearly_last_2_"); err != nil {
		return g, err
	}
	if g.YearlyLast3, err = buildBalanceSheetInfo(reader, prefix+"yearly_last_3_"); err != nil {
		return g, err
	}
	return g, nil
}
func buildBalanceSheetInfo(reader *csvReaderMap, prefix string) (BalanceSheetInfo, error) {
	var err error
	g := BalanceSheetInfo{}
	if g.Date, err = reader.asString(prefix + "date"); err != nil {
		return g, err
	}
	if g.FilingDate, err = reader.asOptionalString(prefix + "filing_date"); err != nil {
		return g, err
	}
	if g.IntangibleAssets, err = reader.asOptionalFloat64(prefix + "intangibleAssets"); err != nil {
		return g, err
	}
	if g.TotalLiab, err = reader.asOptionalFloat64(prefix + "totalLiab"); err != nil {
		return g, err
	}
	if g.TotalStockholderEquity, err = reader.asOptionalFloat64(prefix + "totalStockholderEquity"); err != nil {
		return g, err
	}
	if g.DeferredLongTermLiab, err = reader.asOptionalFloat64(prefix + "deferredLongTermLiab"); err != nil {
		return g, err
	}
	if g.OtherCurrentLiab, err = reader.asOptionalFloat64(prefix + "otherCurrentLiab"); err != nil {
		return g, err
	}
	if g.TotalAssets, err = reader.asOptionalFloat64(prefix + "totalAssets"); err != nil {
		return g, err
	}
	if g.CommonStock, err = reader.asOptionalFloat64(prefix + "commonStock"); err != nil {
		return g, err
	}
	if g.CommonStockSharesOutstanding, err = reader.asOptionalFloat64(prefix + "commonStockSharesOutstanding"); err != nil {
		return g, err
	}
	if g.OtherCurrentAssets, err = reader.asOptionalFloat64(prefix + "otherCurrentAssets"); err != nil {
		return g, err
	}
	if g.RetainedEarnings, err = reader.asOptionalFloat64(prefix + "retainedEarnings"); err != nil {
		return g, err
	}
	if g.OtherLiab, err = reader.asOptionalFloat64(prefix + "otherLiab"); err != nil {
		return g, err
	}
	if g.GoodWill, err = reader.asOptionalFloat64(prefix + "goodWill"); err != nil {
		return g, err
	}
	if g.OtherAssets, err = reader.asOptionalFloat64(prefix + "otherAssets"); err != nil {
		return g, err
	}
	if g.Cash, err = reader.asOptionalFloat64(prefix + "cash"); err != nil {
		return g, err
	}
	if g.TotalCurrentLiabilities, err = reader.asOptionalFloat64(prefix + "totalCurrentLiabilities"); err != nil {
		return g, err
	}
	if g.ShortLongTermDebt, err = reader.asOptionalFloat64(prefix + "shortLongTermDebt"); err != nil {
		return g, err
	}
	if g.OtherStockholderEquity, err = reader.asOptionalFloat64(prefix + "otherStockholderEquity"); err != nil {
		return g, err
	}
	if g.PropertyPlantEquipment, err = reader.asOptionalFloat64(prefix + "propertyPlantEquipment"); err != nil {
		return g, err
	}
	if g.TotalCurrentAssets, err = reader.asOptionalFloat64(prefix + "totalCurrentAssets"); err != nil {
		return g, err
	}
	if g.LongTermInvestments, err = reader.asOptionalFloat64(prefix + "longTermInvestments"); err != nil {
		return g, err
	}
	if g.NetTangibleAssets, err = reader.asOptionalFloat64(prefix + "netTangibleAssets"); err != nil {
		return g, err
	}
	if g.ShortTermInvestments, err = reader.asOptionalFloat64(prefix + "shortTermInvestments"); err != nil {
		return g, err
	}
	if g.NetReceivables, err = reader.asOptionalFloat64(prefix + "netReceivables"); err != nil {
		return g, err
	}
	if g.LongTermDebt, err = reader.asOptionalFloat64(prefix + "longTermDebt"); err != nil {
		return g, err
	}
	if g.Inventory, err = reader.asOptionalFloat64(prefix + "inventory"); err != nil {
		return g, err
	}
	if g.AccountsPayable, err = reader.asOptionalFloat64(prefix + "accountsPayable"); err != nil {
		return g, err
	}
	return g, nil

}
func buildCashFlow(reader *csvReaderMap, prefix string) (CashFlow, error) {
	var err error
	g := CashFlow{}
	if g.CurrencySymbol, err = reader.asString(prefix + "currency_symbol"); err != nil {
		return g, err
	}
	if g.QuarterlyLast0, err = buildCashFlowInfo(reader, prefix+"quarterly_last_0_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast1, err = buildCashFlowInfo(reader, prefix+"quarterly_last_1_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast2, err = buildCashFlowInfo(reader, prefix+"quarterly_last_2_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast3, err = buildCashFlowInfo(reader, prefix+"quarterly_last_3_"); err != nil {
		return g, err
	}
	if g.YearlyLast0, err = buildCashFlowInfo(reader, prefix+"yearly_last_0_"); err != nil {
		return g, err
	}
	if g.YearlyLast1, err = buildCashFlowInfo(reader, prefix+"yearly_last_1_"); err != nil {
		return g, err
	}
	if g.YearlyLast2, err = buildCashFlowInfo(reader, prefix+"yearly_last_2_"); err != nil {
		return g, err
	}
	if g.YearlyLast3, err = buildCashFlowInfo(reader, prefix+"yearly_last_3_"); err != nil {
		return g, err
	}
	return g, nil

}
func buildCashFlowInfo(reader *csvReaderMap, prefix string) (CashFlowInfo, error) {
	var err error
	g := CashFlowInfo{}
	if g.Date, err = reader.asString(prefix + "date"); err != nil {
		return g, err
	}
	if g.FilingDate, err = reader.asOptionalString(prefix + "filing_date"); err != nil {
		return g, err
	}
	if g.Investments, err = reader.asOptionalFloat64(prefix + "investments"); err != nil {
		return g, err
	}
	if g.ChangeToLiabilities, err = reader.asOptionalFloat64(prefix + "changeToLiabilities"); err != nil {
		return g, err
	}
	if g.TotalCashflowsFromInvestingActivities, err = reader.asOptionalFloat64(prefix + "totalCashflowsFromInvestingActivities"); err != nil {
		return g, err
	}
	if g.NetBorrowings, err = reader.asOptionalFloat64(prefix + "netBorrowings"); err != nil {
		return g, err
	}
	if g.TotalCashFromFinancingActivities, err = reader.asOptionalFloat64(prefix + "totalCashFromFinancingActivities"); err != nil {
		return g, err
	}
	if g.ChangeToOperatingActivities, err = reader.asOptionalFloat64(prefix + "changeToOperatingActivities"); err != nil {
		return g, err
	}
	if g.NetIncome, err = reader.asOptionalFloat64(prefix + "netIncome"); err != nil {
		return g, err
	}
	if g.ChangeInCash, err = reader.asOptionalFloat64(prefix + "changeInCash"); err != nil {
		return g, err
	}
	if g.TotalCashFromOperatingActivities, err = reader.asOptionalFloat64(prefix + "totalCashFromOperatingActivities"); err != nil {
		return g, err
	}
	if g.Depreciation, err = reader.asOptionalFloat64(prefix + "depreciation"); err != nil {
		return g, err
	}
	if g.OtherCashflowsFromInvestingActivities, err = reader.asOptionalFloat64(prefix + "otherCashflowsFromInvestingActivities"); err != nil {
		return g, err
	}
	if g.DividendsPaid, err = reader.asOptionalFloat64(prefix + "dividendsPaid"); err != nil {
		return g, err
	}
	if g.ChangeToInventory, err = reader.asOptionalFloat64(prefix + "changeToInventory"); err != nil {
		return g, err
	}
	if g.ChangeToAccountReceivables, err = reader.asOptionalFloat64(prefix + "changeToAccountReceivables"); err != nil {
		return g, err
	}
	if g.SalePurchaseOfStock, err = reader.asOptionalFloat64(prefix + "salePurchaseOfStock"); err != nil {
		return g, err
	}
	if g.OtherCashflowsFromFinancingActivities, err = reader.asOptionalFloat64(prefix + "otherCashflowsFromFinancingActivities"); err != nil {
		return g, err
	}
	if g.ChangeToNetincome, err = reader.asOptionalFloat64(prefix + "changeToNetincome"); err != nil {
		return g, err
	}
	if g.CapitalExpenditures, err = reader.asOptionalFloat64(prefix + "capitalExpenditures"); err != nil {
		return g, err
	}
	return g, nil

}
func buildIncomeStatement(reader *csvReaderMap, prefix string) (IncomeStatement, error) {
	var err error
	g := IncomeStatement{}
	if g.CurrencySymbol, err = reader.asString(prefix + "currency_symbol"); err != nil {
		return g, err
	}
	if g.QuarterlyLast0, err = buildIncomeStatementInfo(reader, prefix+"quarterly_last_0_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast1, err = buildIncomeStatementInfo(reader, prefix+"quarterly_last_1_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast2, err = buildIncomeStatementInfo(reader, prefix+"quarterly_last_2_"); err != nil {
		return g, err
	}
	if g.QuarterlyLast3, err = buildIncomeStatementInfo(reader, prefix+"quarterly_last_3_"); err != nil {
		return g, err
	}
	if g.YearlyLast0, err = buildIncomeStatementInfo(reader, prefix+"yearly_last_0_"); err != nil {
		return g, err
	}
	if g.YearlyLast1, err = buildIncomeStatementInfo(reader, prefix+"yearly_last_1_"); err != nil {
		return g, err
	}
	if g.YearlyLast2, err = buildIncomeStatementInfo(reader, prefix+"yearly_last_2_"); err != nil {
		return g, err
	}
	if g.YearlyLast3, err = buildIncomeStatementInfo(reader, prefix+"yearly_last_3_"); err != nil {
		return g, err
	}
	return g, nil

}
func buildIncomeStatementInfo(reader *csvReaderMap, prefix string) (IncomeStatementInfo, error) {
	var err error
	g := IncomeStatementInfo{}
	if g.Date, err = reader.asString(prefix + "date"); err != nil {
		return g, err
	}
	if g.FilingDate, err = reader.asOptionalString(prefix + "filing_date"); err != nil {
		return g, err
	}
	if g.ResearchDevelopment, err = reader.asOptionalFloat64(prefix + "researchDevelopment"); err != nil {
		return g, err
	}
	if g.EffectOfAccountingCharges, err = reader.asOptionalFloat64(prefix + "effectOfAccountingCharges"); err != nil {
		return g, err
	}
	if g.IncomeBeforeTax, err = reader.asOptionalFloat64(prefix + "incomeBeforeTax"); err != nil {
		return g, err
	}
	if g.MinorityInterest, err = reader.asOptionalFloat64(prefix + "minorityInterest"); err != nil {
		return g, err
	}
	if g.NetIncome, err = reader.asOptionalFloat64(prefix + "netIncome"); err != nil {
		return g, err
	}
	if g.SellingGeneralAdministrative, err = reader.asOptionalFloat64(prefix + "sellingGeneralAdministrative"); err != nil {
		return g, err
	}
	if g.GrossProfit, err = reader.asOptionalFloat64(prefix + "grossProfit"); err != nil {
		return g, err
	}
	if g.Ebit, err = reader.asOptionalFloat64(prefix + "ebit"); err != nil {
		return g, err
	}
	if g.OperatingIncome, err = reader.asOptionalFloat64(prefix + "operatingIncome"); err != nil {
		return g, err
	}
	if g.OtherOperatingExpenses, err = reader.asOptionalFloat64(prefix + "otherOperatingExpenses"); err != nil {
		return g, err
	}
	if g.InterestExpense, err = reader.asOptionalFloat64(prefix + "interestExpense"); err != nil {
		return g, err
	}
	if g.ExtraordinaryItems, err = reader.asOptionalFloat64(prefix + "extraordinaryItems"); err != nil {
		return g, err
	}
	if g.NonRecurring, err = reader.asOptionalFloat64(prefix + "nonRecurring"); err != nil {
		return g, err
	}
	if g.OtherItems, err = reader.asOptionalFloat64(prefix + "otherItems"); err != nil {
		return g, err
	}
	if g.IncomeTaxExpense, err = reader.asOptionalFloat64(prefix + "incomeTaxExpense"); err != nil {
		return g, err
	}
	if g.TotalRevenue, err = reader.asOptionalFloat64(prefix + "totalRevenue"); err != nil {
		return g, err
	}
	if g.TotalOperatingExpenses, err = reader.asOptionalFloat64(prefix + "totalOperatingExpenses"); err != nil {
		return g, err
	}
	if g.CostOfRevenue, err = reader.asOptionalFloat64(prefix + "costOfRevenue"); err != nil {
		return g, err
	}
	if g.TotalOtherIncomeExpenseNet, err = reader.asOptionalFloat64(prefix + "totalOtherIncomeExpenseNet"); err != nil {
		return g, err
	}
	if g.DiscontinuedOperations, err = reader.asOptionalFloat64(prefix + "discontinuedOperations"); err != nil {
		return g, err
	}
	if g.NetIncomeFromContinuingOps, err = reader.asOptionalFloat64(prefix + "netIncomeFromContinuingOps"); err != nil {
		return g, err
	}
	if g.NetIncomeApplicableToCommonShares, err = reader.asOptionalFloat64(prefix + "netIncomeApplicableToCommonShares"); err != nil {
		return g, err
	}
	return g, nil

}
func buildEarningsInfo(reader *csvReaderMap, prefix string) (EarningsInfo, error) {
	var err error
	g := EarningsInfo{}
	if g.Date, err = reader.asString(prefix + "date"); err != nil {
		return g, err
	}
	if g.EpsActual, err = reader.asOptionalFloat64(prefix + "epsActual"); err != nil {
		return g, err
	}
	if g.EpsEstimate, err = reader.asOptionalFloat64(prefix + "epsEstimate"); err != nil {
		return g, err
	}
	if g.EpsDifference, err = reader.asOptionalFloat64(prefix + "epsDifference"); err != nil {
		return g, err
	}
	if g.SurprisePercent, err = reader.asOptionalFloat64(prefix + "surprisePercent"); err != nil {
		return g, err
	}
	return g, nil
}
