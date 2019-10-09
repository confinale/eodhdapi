package eodhdapi

//go:generate go run github.com/mailru/easyjson/easyjson -all fundamentals.go

import (
	"context"
	"fmt"
	"github.com/gitu/eodhdapi/exchanges"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
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
	MarketCapitalization       *decimal.Decimal `json:"MarketCapitalization"`
	MarketCapitalizationMln    string   `json:"MarketCapitalizationMln"`
	EBITDA                     *decimal.Decimal `json:"EBITDA"`
	PERatio                    *decimal.Decimal `json:"PERatio"`
	PEGRatio                   *decimal.Decimal `json:"PEGRatio"`
	WallStreetTargetPrice      *decimal.Decimal `json:"WallStreetTargetPrice"`
	BookValue                  *decimal.Decimal `json:"BookValue"`
	DividendShare              *decimal.Decimal `json:"DividendShare"`
	DividendYield              *decimal.Decimal `json:"DividendYield"`
	EarningsShare              *decimal.Decimal `json:"EarningsShare"`
	EPSEstimateCurrentYear     *decimal.Decimal `json:"EPSEstimateCurrentYear"`
	EPSEstimateNextYear        *decimal.Decimal `json:"EPSEstimateNextYear"`
	EPSEstimateNextQuarter     *decimal.Decimal `json:"EPSEstimateNextQuarter"`
	MostRecentQuarter          string   `json:"MostRecentQuarter"`
	ProfitMargin               *decimal.Decimal `json:"ProfitMargin"`
	OperatingMarginTTM         *decimal.Decimal `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          *decimal.Decimal `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          *decimal.Decimal `json:"ReturnOnEquityTTM"`
	RevenueTTM                 *decimal.Decimal `json:"RevenueTTM"`
	RevenuePerShareTTM         *decimal.Decimal `json:"RevenuePerShareTTM"`
	QuarterlyRevenueGrowthYOY  *decimal.Decimal `json:"QuarterlyRevenueGrowthYOY"`
	GrossProfitTTM             *decimal.Decimal `json:"GrossProfitTTM"`
	DilutedEpsTTM              *decimal.Decimal `json:"DilutedEpsTTM"`
	QuarterlyEarningsGrowthYOY *decimal.Decimal `json:"QuarterlyEarningsGrowthYOY"`
}

type ETFData struct {
	ISIN                    string                `json:"ISIN"`
	CompanyName             string                `json:"Company_Name"`
	CompanyURL              string                `json:"Company_URL"`
	ETFURL                  string                `json:"ETF_URL"`
	Yield                   string                `json:"Yield"`
	DividendPayingFrequency string                `json:"Dividend_Paying_Frequency"`
	InceptionDate           string                `json:"Inception_Date"`
	MaxAnnualMgmtCharge     string                `json:"Max_Annual_Mgmt_Charge"`
	OngoingCharge           string                `json:"Ongoing_Charge"`
	DateOngoingCharge       string                `json:"Date_Ongoing_Charge"`
	NetExpenseRatio         string                `json:"NetExpenseRatio"`
	AnnualHoldingsTurnover  string                `json:"AnnualHoldingsTurnover"`
	TotalAssets             string                `json:"TotalAssets"`
	AverageMktCapMil        string                `json:"Average_Mkt_Cap_Mil"`
	AssetAllocation         map[string]Allocation `json:"Asset_Allocation"`
	WorldRegions            map[string]Weight     `json:"World_Regions"`
	SectorWeights           map[string]Weight     `json:"Sector_Weights"`
	Top10Holdings           map[string]Holding    `json:"Top_10_Holdings"`
	Holdings                map[string]Holding    `json:"Holdings"`
	MorningStar             MorningStar           `json:"MorningStar"`
	//ValuationsGrowth        ValuationsGrowth     `json:"Valuations_Growth"`
	//Performance             Performance          `json:"Performance"`map[string]Weight
}
type MorningStar struct {
	Ratio               int    `json:"Ratio"`
	CategoryBenchmark   string `json:"Category_Benchmark"`
	SustainabilityRatio int    `json:"Sustainability_Ratio"`
}

type Holding struct {
	Name          string  `json:"Name"`
	AssetsPercent *decimal.Decimal `json:"Assets_%"`
}
type Weight struct {
	EquityPercent      string  `json:"Equity_%"`
	RelativeToCategory *decimal.Decimal `json:"Relative_to_Category"`
}

type Allocation struct {
	LongPercent      *decimal.Decimal `json:"Long_%"`
	ShortPercent     *decimal.Decimal `json:"Short_%"`
	NetAssetsPercent *decimal.Decimal `json:"Net_Assets_%"`
}

type Valuation struct {
	TrailingPE             *decimal.Decimal `json:"TrailingPE"`
	ForwardPE              *decimal.Decimal `json:"ForwardPE"`
	PriceSalesTTM          *decimal.Decimal `json:"PriceSalesTTM"`
	PriceBookMRQ           *decimal.Decimal `json:"PriceBookMRQ"`
	EnterpriseValueRevenue *decimal.Decimal `json:"EnterpriseValueRevenue"`
	EnterpriseValueEbitda  *decimal.Decimal `json:"EnterpriseValueEbitda"`
}
type Technicals struct {
	Beta                  *decimal.Decimal `json:"Beta"`
	FiftyTwoWeekHigh      *decimal.Decimal `json:"52WeekHigh"`
	FiftyTwoWeekLow       *decimal.Decimal `json:"52WeekLow"`
	FiftyDayMA            *decimal.Decimal `json:"50DayMA"`
	TwoHundredDayMA       *decimal.Decimal `json:"200DayMA"`
	SharesShort           *decimal.Decimal `json:"SharesShort"`
	SharesShortPriorMonth *decimal.Decimal `json:"SharesShortPriorMonth"`
	ShortRatio            *decimal.Decimal `json:"ShortRatio"`
	ShortPercent          *decimal.Decimal `json:"ShortPercent"`
}
type SplitsDividends struct {
	ForwardAnnualDividendRate  *decimal.Decimal `json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield *decimal.Decimal `json:"ForwardAnnualDividendYield"`
	PayoutRatio                *decimal.Decimal `json:"PayoutRatio"`
	DividendDate               string   `json:"DividendDate"`
	ExDividendDate             string   `json:"ExDividendDate"`
	LastSplitFactor            string   `json:"LastSplitFactor"`
	LastSplitDate              string   `json:"LastSplitDate"`
}
type EarningsInfo struct {
	Date            string   `json:"date"`
	EpsActual       *decimal.Decimal `json:"epsActual"`
	EpsEstimate     *decimal.Decimal `json:"epsEstimate"`
	EpsDifference   *decimal.Decimal `json:"epsDifference"`
	SurprisePercent *decimal.Decimal `json:"surprisePercent"`
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
	IntangibleAssets             *decimal.Decimal `json:"intangibleAssets"`
	TotalLiab                    *decimal.Decimal `json:"totalLiab"`
	TotalStockholderEquity       *decimal.Decimal `json:"totalStockholderEquity"`
	DeferredLongTermLiab         *decimal.Decimal `json:"deferredLongTermLiab"`
	OtherCurrentLiab             *decimal.Decimal `json:"otherCurrentLiab"`
	TotalAssets                  *decimal.Decimal `json:"totalAssets"`
	CommonStock                  *decimal.Decimal `json:"commonStock"`
	CommonStockSharesOutstanding *decimal.Decimal `json:"commonStockSharesOutStanding"`
	OtherCurrentAssets           *decimal.Decimal `json:"otherCurrentAssets"`
	RetainedEarnings             *decimal.Decimal `json:"retainedEarnings"`
	OtherLiab                    *decimal.Decimal `json:"otherLiab"`
	GoodWill                     *decimal.Decimal `json:"goodWill"`
	OtherAssets                  *decimal.Decimal `json:"otherAssets"`
	Cash                         *decimal.Decimal `json:"cash"`
	TotalCurrentLiabilities      *decimal.Decimal `json:"totalCurrentLiabilities"`
	ShortLongTermDebt            *decimal.Decimal `json:"shortLongTermDebt"`
	OtherStockholderEquity       *decimal.Decimal `json:"otherStockholderEquity"`
	PropertyPlantEquipment       *decimal.Decimal `json:"propertyPlantEquipment"`
	TotalCurrentAssets           *decimal.Decimal `json:"totalCurrentAssets"`
	LongTermInvestments          *decimal.Decimal `json:"longTermInvestments"`
	NetTangibleAssets            *decimal.Decimal `json:"netTangibleAssets"`
	ShortTermInvestments         *decimal.Decimal `json:"shortTermInvestments"`
	NetReceivables               *decimal.Decimal `json:"netReceivables"`
	LongTermDebt                 *decimal.Decimal `json:"longTermDebt"`
	Inventory                    *decimal.Decimal `json:"inventory"`
	AccountsPayable              *decimal.Decimal `json:"accountsPayable"`
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
	Investments                           *decimal.Decimal `json:"investments"`
	ChangeToLiabilities                   *decimal.Decimal `json:"changeToLiabilities"`
	TotalCashflowsFromInvestingActivities *decimal.Decimal `json:"totalCashflowsFromInvestingActivities"`
	NetBorrowings                         *decimal.Decimal `json:"netBorrowings"`
	TotalCashFromFinancingActivities      *decimal.Decimal `json:"totalCashFromFinancingActivities"`
	ChangeToOperatingActivities           *decimal.Decimal `json:"changeToOperatingActivities"`
	NetIncome                             *decimal.Decimal `json:"netIncome"`
	ChangeInCash                          *decimal.Decimal `json:"changeInCash"`
	TotalCashFromOperatingActivities      *decimal.Decimal `json:"totalCashFromOperatingActivities"`
	Depreciation                          *decimal.Decimal `json:"depreciation"`
	OtherCashflowsFromInvestingActivities *decimal.Decimal `json:"otherCashflowsFromInvestingActivities"`
	DividendsPaid                         *decimal.Decimal `json:"dividendsPaid"`
	ChangeToInventory                     *decimal.Decimal `json:"changeToInventory"`
	ChangeToAccountReceivables            *decimal.Decimal `json:"changeToAccountReceivables"`
	SalePurchaseOfStock                   *decimal.Decimal `json:"salePurchaseOfStock"`
	OtherCashflowsFromFinancingActivities *decimal.Decimal `json:"otherCashflowsFromFinancingActivities"`
	ChangeToNetincome                     *decimal.Decimal `json:"changeToNetincome"`
	CapitalExpenditures                   *decimal.Decimal `json:"capitalExpenditures"`
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
	ResearchDevelopment               *decimal.Decimal `json:"researchDevelopment"`
	EffectOfAccountingCharges         *decimal.Decimal `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   *decimal.Decimal `json:"incomeBeforeTax"`
	MinorityInterest                  *decimal.Decimal `json:"minorityInterest"`
	NetIncome                         *decimal.Decimal `json:"netIncome"`
	SellingGeneralAdministrative      *decimal.Decimal `json:"sellingGeneralAdministrative"`
	GrossProfit                       *decimal.Decimal `json:"grossProfit"`
	Ebit                              *decimal.Decimal `json:"ebit"`
	OperatingIncome                   *decimal.Decimal `json:"operatingIncome"`
	OtherOperatingExpenses            *decimal.Decimal `json:"otherOperatingExpenses"`
	InterestExpense                   *decimal.Decimal `json:"interestExpense"`
	ExtraordinaryItems                *decimal.Decimal `json:"extraordinaryItems"`
	NonRecurring                      *decimal.Decimal `json:"nonRecurring"`
	OtherItems                        *decimal.Decimal `json:"otherItems"`
	IncomeTaxExpense                  *decimal.Decimal `json:"incomeTaxExpense"`
	TotalRevenue                      *decimal.Decimal `json:"totalRevenue"`
	TotalOperatingExpenses            *decimal.Decimal `json:"totalOperatingExpenses"`
	CostOfRevenue                     *decimal.Decimal `json:"costOfRevenue"`
	TotalOtherIncomeExpenseNet        *decimal.Decimal `json:"totalOtherIncomeExpenseNet"`
	DiscontinuedOperations            *decimal.Decimal `json:"discontinuedOperations"`
	NetIncomeFromContinuingOps        *decimal.Decimal `json:"netIncomeFromContinuingOps"`
	NetIncomeApplicableToCommonShares *decimal.Decimal `json:"netIncomeApplicableToCommonShares"`
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
	if g.MarketCapitalization, err = reader.asOptionalDecimal(prefix + "MarketCapitalization"); err != nil {
		return err
	}
	if g.MarketCapitalizationMln, err = reader.asString(prefix + "MarketCapitalizationMln"); err != nil {
		return err
	}
	if g.EBITDA, err = reader.asOptionalDecimal(prefix + "EBITDA"); err != nil {
		return err
	}
	if g.PERatio, err = reader.asOptionalDecimal(prefix + "PERatio"); err != nil {
		return err
	}
	if g.PEGRatio, err = reader.asOptionalDecimal(prefix + "PEGRatio"); err != nil {
		return err
	}
	if g.WallStreetTargetPrice, err = reader.asOptionalDecimal(prefix + "WallStreetTargetPrice"); err != nil {
		return err
	}
	if g.BookValue, err = reader.asOptionalDecimal(prefix + "BookValue"); err != nil {
		return err
	}
	if g.DividendShare, err = reader.asOptionalDecimal(prefix + "DividendShare"); err != nil {
		return err
	}
	if g.DividendYield, err = reader.asOptionalDecimal(prefix + "DividendYield"); err != nil {
		return err
	}
	if g.EarningsShare, err = reader.asOptionalDecimal(prefix + "EarningsShare"); err != nil {
		return err
	}
	if g.EPSEstimateCurrentYear, err = reader.asOptionalDecimal(prefix + "EPSEstimateCurrentYear"); err != nil {
		return err
	}
	if g.EPSEstimateNextYear, err = reader.asOptionalDecimal(prefix + "EPSEstimateNextYear"); err != nil {
		return err
	}
	if g.EPSEstimateNextQuarter, err = reader.asOptionalDecimal(prefix + "EPSEstimateNextQuarter"); err != nil {
		return err
	}
	if g.MostRecentQuarter, err = reader.asString(prefix + "MostRecentQuarter"); err != nil {
		return err
	}
	if g.ProfitMargin, err = reader.asOptionalDecimal(prefix + "ProfitMargin"); err != nil {
		return err
	}
	if g.OperatingMarginTTM, err = reader.asOptionalDecimal(prefix + "OperatingMarginTTM"); err != nil {
		return err
	}
	if g.ReturnOnAssetsTTM, err = reader.asOptionalDecimal(prefix + "ReturnOnAssetsTTM"); err != nil {
		return err
	}
	if g.ReturnOnEquityTTM, err = reader.asOptionalDecimal(prefix + "ReturnOnEquityTTM"); err != nil {
		return err
	}
	if g.RevenueTTM, err = reader.asOptionalDecimal(prefix + "RevenueTTM"); err != nil {
		return err
	}
	if g.RevenuePerShareTTM, err = reader.asOptionalDecimal(prefix + "RevenuePerShareTTM"); err != nil {
		return err
	}
	if g.QuarterlyRevenueGrowthYOY, err = reader.asOptionalDecimal(prefix + "QuarterlyRevenueGrowthYOY"); err != nil {
		return err
	}
	if g.GrossProfitTTM, err = reader.asOptionalDecimal(prefix + "GrossProfitTTM"); err != nil {
		return err
	}
	if g.DilutedEpsTTM, err = reader.asOptionalDecimal(prefix + "DilutedEpsTTM"); err != nil {
		return err
	}
	if g.QuarterlyEarningsGrowthYOY, err = reader.asOptionalDecimal(prefix + "QuarterlyEarningsGrowthYOY"); err != nil {
		return err
	}
	return nil

}
func (g *Valuation) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.TrailingPE, err = reader.asOptionalDecimal(prefix + "TrailingPE"); err != nil {
		return err
	}
	if g.ForwardPE, err = reader.asOptionalDecimal(prefix + "ForwardPE"); err != nil {
		return err
	}
	if g.PriceSalesTTM, err = reader.asOptionalDecimal(prefix + "PriceSalesTTM"); err != nil {
		return err
	}
	if g.PriceBookMRQ, err = reader.asOptionalDecimal(prefix + "PriceBookMRQ"); err != nil {
		return err
	}
	if g.EnterpriseValueRevenue, err = reader.asOptionalDecimal(prefix + "EnterpriseValueRevenue"); err != nil {
		return err
	}
	if g.EnterpriseValueEbitda, err = reader.asOptionalDecimal(prefix + "EnterpriseValueEbitda"); err != nil {
		return err
	}
	return nil
}
func (g *Technicals) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.Beta, err = reader.asOptionalDecimal(prefix + "Beta"); err != nil {
		return err
	}
	if g.FiftyTwoWeekHigh, err = reader.asOptionalDecimal(prefix + "52WeekHigh"); err != nil {
		return err
	}
	if g.FiftyTwoWeekLow, err = reader.asOptionalDecimal(prefix + "52WeekLow"); err != nil {
		return err
	}
	if g.FiftyDayMA, err = reader.asOptionalDecimal(prefix + "50DayMA"); err != nil {
		return err
	}
	if g.TwoHundredDayMA, err = reader.asOptionalDecimal(prefix + "200DayMA"); err != nil {
		return err
	}
	if g.SharesShort, err = reader.asOptionalDecimal(prefix + "SharesShort"); err != nil {
		return err
	}
	if g.SharesShortPriorMonth, err = reader.asOptionalDecimal(prefix + "SharesShortPriorMonth"); err != nil {
		return err
	}
	if g.ShortRatio, err = reader.asOptionalDecimal(prefix + "ShortRatio"); err != nil {
		return err
	}
	if g.ShortPercent, err = reader.asOptionalDecimal(prefix + "ShortPercent"); err != nil {
		return err
	}
	return nil
}
func (g *SplitsDividends) fill(reader *csvReaderMap, prefix string) error {
	var err error
	if g.ForwardAnnualDividendRate, err = reader.asOptionalDecimal(prefix + "ForwardAnnualDividendRate"); err != nil {
		return err
	}
	if g.ForwardAnnualDividendYield, err = reader.asOptionalDecimal(prefix + "ForwardAnnualDividendYield"); err != nil {
		return err
	}
	if g.PayoutRatio, err = reader.asOptionalDecimal(prefix + "PayoutRatio"); err != nil {
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
	if g.IntangibleAssets, err = reader.asOptionalDecimal(prefix + "intangibleAssets"); err != nil {
		return g, err
	}
	if g.TotalLiab, err = reader.asOptionalDecimal(prefix + "totalLiab"); err != nil {
		return g, err
	}
	if g.TotalStockholderEquity, err = reader.asOptionalDecimal(prefix + "totalStockholderEquity"); err != nil {
		return g, err
	}
	if g.DeferredLongTermLiab, err = reader.asOptionalDecimal(prefix + "deferredLongTermLiab"); err != nil {
		return g, err
	}
	if g.OtherCurrentLiab, err = reader.asOptionalDecimal(prefix + "otherCurrentLiab"); err != nil {
		return g, err
	}
	if g.TotalAssets, err = reader.asOptionalDecimal(prefix + "totalAssets"); err != nil {
		return g, err
	}
	if g.CommonStock, err = reader.asOptionalDecimal(prefix + "commonStock"); err != nil {
		return g, err
	}
	if g.CommonStockSharesOutstanding, err = reader.asOptionalDecimal(prefix + "commonStockSharesOutstanding"); err != nil {
		return g, err
	}
	if g.OtherCurrentAssets, err = reader.asOptionalDecimal(prefix + "otherCurrentAssets"); err != nil {
		return g, err
	}
	if g.RetainedEarnings, err = reader.asOptionalDecimal(prefix + "retainedEarnings"); err != nil {
		return g, err
	}
	if g.OtherLiab, err = reader.asOptionalDecimal(prefix + "otherLiab"); err != nil {
		return g, err
	}
	if g.GoodWill, err = reader.asOptionalDecimal(prefix + "goodWill"); err != nil {
		return g, err
	}
	if g.OtherAssets, err = reader.asOptionalDecimal(prefix + "otherAssets"); err != nil {
		return g, err
	}
	if g.Cash, err = reader.asOptionalDecimal(prefix + "cash"); err != nil {
		return g, err
	}
	if g.TotalCurrentLiabilities, err = reader.asOptionalDecimal(prefix + "totalCurrentLiabilities"); err != nil {
		return g, err
	}
	if g.ShortLongTermDebt, err = reader.asOptionalDecimal(prefix + "shortLongTermDebt"); err != nil {
		return g, err
	}
	if g.OtherStockholderEquity, err = reader.asOptionalDecimal(prefix + "otherStockholderEquity"); err != nil {
		return g, err
	}
	if g.PropertyPlantEquipment, err = reader.asOptionalDecimal(prefix + "propertyPlantEquipment"); err != nil {
		return g, err
	}
	if g.TotalCurrentAssets, err = reader.asOptionalDecimal(prefix + "totalCurrentAssets"); err != nil {
		return g, err
	}
	if g.LongTermInvestments, err = reader.asOptionalDecimal(prefix + "longTermInvestments"); err != nil {
		return g, err
	}
	if g.NetTangibleAssets, err = reader.asOptionalDecimal(prefix + "netTangibleAssets"); err != nil {
		return g, err
	}
	if g.ShortTermInvestments, err = reader.asOptionalDecimal(prefix + "shortTermInvestments"); err != nil {
		return g, err
	}
	if g.NetReceivables, err = reader.asOptionalDecimal(prefix + "netReceivables"); err != nil {
		return g, err
	}
	if g.LongTermDebt, err = reader.asOptionalDecimal(prefix + "longTermDebt"); err != nil {
		return g, err
	}
	if g.Inventory, err = reader.asOptionalDecimal(prefix + "inventory"); err != nil {
		return g, err
	}
	if g.AccountsPayable, err = reader.asOptionalDecimal(prefix + "accountsPayable"); err != nil {
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
	if g.Investments, err = reader.asOptionalDecimal(prefix + "investments"); err != nil {
		return g, err
	}
	if g.ChangeToLiabilities, err = reader.asOptionalDecimal(prefix + "changeToLiabilities"); err != nil {
		return g, err
	}
	if g.TotalCashflowsFromInvestingActivities, err = reader.asOptionalDecimal(prefix + "totalCashflowsFromInvestingActivities"); err != nil {
		return g, err
	}
	if g.NetBorrowings, err = reader.asOptionalDecimal(prefix + "netBorrowings"); err != nil {
		return g, err
	}
	if g.TotalCashFromFinancingActivities, err = reader.asOptionalDecimal(prefix + "totalCashFromFinancingActivities"); err != nil {
		return g, err
	}
	if g.ChangeToOperatingActivities, err = reader.asOptionalDecimal(prefix + "changeToOperatingActivities"); err != nil {
		return g, err
	}
	if g.NetIncome, err = reader.asOptionalDecimal(prefix + "netIncome"); err != nil {
		return g, err
	}
	if g.ChangeInCash, err = reader.asOptionalDecimal(prefix + "changeInCash"); err != nil {
		return g, err
	}
	if g.TotalCashFromOperatingActivities, err = reader.asOptionalDecimal(prefix + "totalCashFromOperatingActivities"); err != nil {
		return g, err
	}
	if g.Depreciation, err = reader.asOptionalDecimal(prefix + "depreciation"); err != nil {
		return g, err
	}
	if g.OtherCashflowsFromInvestingActivities, err = reader.asOptionalDecimal(prefix + "otherCashflowsFromInvestingActivities"); err != nil {
		return g, err
	}
	if g.DividendsPaid, err = reader.asOptionalDecimal(prefix + "dividendsPaid"); err != nil {
		return g, err
	}
	if g.ChangeToInventory, err = reader.asOptionalDecimal(prefix + "changeToInventory"); err != nil {
		return g, err
	}
	if g.ChangeToAccountReceivables, err = reader.asOptionalDecimal(prefix + "changeToAccountReceivables"); err != nil {
		return g, err
	}
	if g.SalePurchaseOfStock, err = reader.asOptionalDecimal(prefix + "salePurchaseOfStock"); err != nil {
		return g, err
	}
	if g.OtherCashflowsFromFinancingActivities, err = reader.asOptionalDecimal(prefix + "otherCashflowsFromFinancingActivities"); err != nil {
		return g, err
	}
	if g.ChangeToNetincome, err = reader.asOptionalDecimal(prefix + "changeToNetincome"); err != nil {
		return g, err
	}
	if g.CapitalExpenditures, err = reader.asOptionalDecimal(prefix + "capitalExpenditures"); err != nil {
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
	if g.ResearchDevelopment, err = reader.asOptionalDecimal(prefix + "researchDevelopment"); err != nil {
		return g, err
	}
	if g.EffectOfAccountingCharges, err = reader.asOptionalDecimal(prefix + "effectOfAccountingCharges"); err != nil {
		return g, err
	}
	if g.IncomeBeforeTax, err = reader.asOptionalDecimal(prefix + "incomeBeforeTax"); err != nil {
		return g, err
	}
	if g.MinorityInterest, err = reader.asOptionalDecimal(prefix + "minorityInterest"); err != nil {
		return g, err
	}
	if g.NetIncome, err = reader.asOptionalDecimal(prefix + "netIncome"); err != nil {
		return g, err
	}
	if g.SellingGeneralAdministrative, err = reader.asOptionalDecimal(prefix + "sellingGeneralAdministrative"); err != nil {
		return g, err
	}
	if g.GrossProfit, err = reader.asOptionalDecimal(prefix + "grossProfit"); err != nil {
		return g, err
	}
	if g.Ebit, err = reader.asOptionalDecimal(prefix + "ebit"); err != nil {
		return g, err
	}
	if g.OperatingIncome, err = reader.asOptionalDecimal(prefix + "operatingIncome"); err != nil {
		return g, err
	}
	if g.OtherOperatingExpenses, err = reader.asOptionalDecimal(prefix + "otherOperatingExpenses"); err != nil {
		return g, err
	}
	if g.InterestExpense, err = reader.asOptionalDecimal(prefix + "interestExpense"); err != nil {
		return g, err
	}
	if g.ExtraordinaryItems, err = reader.asOptionalDecimal(prefix + "extraordinaryItems"); err != nil {
		return g, err
	}
	if g.NonRecurring, err = reader.asOptionalDecimal(prefix + "nonRecurring"); err != nil {
		return g, err
	}
	if g.OtherItems, err = reader.asOptionalDecimal(prefix + "otherItems"); err != nil {
		return g, err
	}
	if g.IncomeTaxExpense, err = reader.asOptionalDecimal(prefix + "incomeTaxExpense"); err != nil {
		return g, err
	}
	if g.TotalRevenue, err = reader.asOptionalDecimal(prefix + "totalRevenue"); err != nil {
		return g, err
	}
	if g.TotalOperatingExpenses, err = reader.asOptionalDecimal(prefix + "totalOperatingExpenses"); err != nil {
		return g, err
	}
	if g.CostOfRevenue, err = reader.asOptionalDecimal(prefix + "costOfRevenue"); err != nil {
		return g, err
	}
	if g.TotalOtherIncomeExpenseNet, err = reader.asOptionalDecimal(prefix + "totalOtherIncomeExpenseNet"); err != nil {
		return g, err
	}
	if g.DiscontinuedOperations, err = reader.asOptionalDecimal(prefix + "discontinuedOperations"); err != nil {
		return g, err
	}
	if g.NetIncomeFromContinuingOps, err = reader.asOptionalDecimal(prefix + "netIncomeFromContinuingOps"); err != nil {
		return g, err
	}
	if g.NetIncomeApplicableToCommonShares, err = reader.asOptionalDecimal(prefix + "netIncomeApplicableToCommonShares"); err != nil {
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
	if g.EpsActual, err = reader.asOptionalDecimal(prefix + "epsActual"); err != nil {
		return g, err
	}
	if g.EpsEstimate, err = reader.asOptionalDecimal(prefix + "epsEstimate"); err != nil {
		return g, err
	}
	if g.EpsDifference, err = reader.asOptionalDecimal(prefix + "epsDifference"); err != nil {
		return g, err
	}
	if g.SurprisePercent, err = reader.asOptionalDecimal(prefix + "surprisePercent"); err != nil {
		return g, err
	}
	return g, nil
}
