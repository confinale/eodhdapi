package eodhdapi

//go:generate go run github.com/mailru/easyjson/easyjson -omit_empty -disallow_unknown_fields -all fundamentals.go

import (
	"context"
	"fmt"
	"github.com/gitu/eodhdapi/exchanges"
	"github.com/mailru/easyjson/jlexer"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"log"
	"strconv"
	"strings"
	"time"
)

// Fundamentals for a ticker
type Fundamentals struct {
	LastUpdate        time.Time
	Ticker            string
	General           General            `json:"General"`
	Highlights        *Highlights        `json:"Highlights"`
	Valuation         *Valuation         `json:"Valuation"`
	Technicals        *Technicals        `json:"Technicals"`
	SplitsDividends   *SplitsDividends   `json:"SplitsDividends"`
	Earnings          *Earnings          `json:"Earnings"`
	Financials        *Financials        `json:"Financials"`
	ETFData           *ETFData           `json:"ETF_Data"`
	SharesStats       *SharesStats       `json:"SharesStats"`
	OutstandingShares *OutstandingShares `json:"outstandingShares"`
	Components        Components         `json:"Components"`
	MutualFundData    *MutualFundData    `json:"MutualFund_Data"`
}

type General struct {
	Code                  string  `json:"Code"`
	Type                  string  `json:"Type"`
	Name                  string  `json:"Name"`
	Exchange              string  `json:"Exchange"`
	CurrencyCode          string  `json:"CurrencyCode"`
	CurrencyName          string  `json:"CurrencyName"`
	CurrencySymbol        string  `json:"CurrencySymbol"`
	CountryName           string  `json:"CountryName"`
	CountryISO            string  `json:"CountryISO"`
	Sector                string  `json:"Sector"`
	Industry              string  `json:"Industry"`
	Description           string  `json:"Description"`
	ISIN                  *string `json:"ISIN"`
	FullTimeEmployees     *int    `json:"FullTimeEmployees"`
	UpdatedAt             *string `json:"UpdatedAt"`
	Cusip                 *string `json:"CUSIP"`
	LogoURL               *string `json:"LogoURL"`
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
	Category              *string `json:"Category"`
	FundSummary           *string `json:"Fund_Summary"`
	FundFamily            *string `json:"Fund_Family"`
	FundFiscalYearEnd     *string `json:"Fiscal_Year_End"`
}
type Highlights struct {
	MarketCapitalization       *decimal.Decimal `json:"MarketCapitalization"`
	MarketCapitalizationMln    string           `json:"MarketCapitalizationMln"`
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
	MostRecentQuarter          string           `json:"MostRecentQuarter"`
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
	ISIN                    string             `json:"ISIN"`
	CompanyName             string             `json:"Company_Name"`
	CompanyURL              string             `json:"Company_URL"`
	ETFURL                  string             `json:"ETF_URL"`
	Yield                   string             `json:"Yield"`
	DividendPayingFrequency string             `json:"Dividend_Paying_Frequency"`
	InceptionDate           string             `json:"Inception_Date"`
	MaxAnnualMgmtCharge     string             `json:"Max_Annual_Mgmt_Charge"`
	OngoingCharge           string             `json:"Ongoing_Charge"`
	DateOngoingCharge       string             `json:"Date_Ongoing_Charge"`
	NetExpenseRatio         string             `json:"NetExpenseRatio"`
	AnnualHoldingsTurnover  string             `json:"AnnualHoldingsTurnover"`
	TotalAssets             string             `json:"TotalAssets"`
	AverageMktCapMil        string             `json:"Average_Mkt_Cap_Mil"`
	AssetAllocation         ETFAssetAllocation `json:"Asset_Allocation"`
	WorldRegions            Weights            `json:"World_Regions"`
	SectorWeights           Weights            `json:"Sector_Weights"`
	Top10Holdings           Holdings           `json:"Top_10_Holdings"`
	Holdings                Holdings           `json:"Holdings"`
	MorningStar             MorningStar        `json:"MorningStar"`
	//ValuationsGrowth        ValuationsGrowth     `json:"Valuations_Growth"`
	Performance          Performance `json:"Performance"`
	MarketCapitalisation interface{} `json:"Market_Capitalisation"`
	ValuationsGrowth     interface{} `json:"Valuations_Growth"`
}

type Performance struct {
	ThreeYVolatility *decimal.Decimal `json:"3y_Volatility"`
	ThreeYExpReturn  *decimal.Decimal `json:"3y_ExpReturn"`
	ThreeYSharpRatio *decimal.Decimal `json:"3y_SharpRatio"`
	ReturnsYTD       *decimal.Decimal `json:"Returns_YTD"`
	Returns3Y        *decimal.Decimal `json:"Returns_3Y"`
	Returns5Y        *decimal.Decimal `json:"Returns_5Y"`
	Returns10Y       *decimal.Decimal `json:"Returns_10Y"`
}

type MorningStar struct {
	Ratio               int    `json:"Ratio,string"`
	CategoryBenchmark   string `json:"Category_Benchmark"`
	SustainabilityRatio int    `json:"Sustainability_Ratio,string"`
}

type Holding struct {
	Name          string           `json:"Name"`
	AssetsPercent *decimal.Decimal `json:"Assets_%"`
}
type Weight struct {
	Category           string           `json:"Category"`
	EquityPercent      string           `json:"Equity_%"`
	RelativeToCategory *decimal.Decimal `json:"Relative_to_Category"`
}

type Allocation struct {
	Category         string           `json:"Category"`
	LongPercent      *decimal.Decimal `json:"Long_%"`
	ShortPercent     *decimal.Decimal `json:"Short_%"`
	NetAssetsPercent *decimal.Decimal `json:"Net_Assets_%"`
}

type Holdings []Holding

func (out *Holdings) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]Holding, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 Holding
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]Holding, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 Holding
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type Weights []Weight

func (out *Weights) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]Weight, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 Weight
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]Weight, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v37 Weight
			(v37).UnmarshalEasyJSON(in)
			v37.Category = key
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type ETFAssetAllocation []Allocation

func (out *ETFAssetAllocation) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]Allocation, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 Allocation
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]Allocation, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v37 Allocation
			(v37).UnmarshalEasyJSON(in)
			v37.Category = key
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type AssetAllocation struct {
	Type            string           `json:"Type"`
	Net             *decimal.Decimal `json:"Net_%"`
	Long            *decimal.Decimal `json:"Long_%"`
	Short           *decimal.Decimal `json:"Short_%"`
	CategoryAverage *decimal.Decimal `json:"Category_Average"`
	Benchmark       *decimal.Decimal `json:"Benchmark"`
}

type AssetAllocations []AssetAllocation

func (out *AssetAllocations) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]AssetAllocation, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 AssetAllocation
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]AssetAllocation, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 AssetAllocation
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
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
	DividendDate               string           `json:"DividendDate"`
	ExDividendDate             string           `json:"ExDividendDate"`
	LastSplitFactor            string           `json:"LastSplitFactor"`
	LastSplitDate              string           `json:"LastSplitDate"`
	NumberDividendsByYear      YearCounts       `json:"NumberDividendsByYear"`
}

type OutstandingShares struct {
	Annual    SharesOutstandings `json:"annual"`
	Quarterly SharesOutstandings `json:"quarterly"`
}

type SharesOutstanding struct {
	Date      string `json:"date"`
	SharesMln string `json:"sharesMln"`
}

type SharesOutstandings []SharesOutstanding

func (out *SharesOutstandings) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]SharesOutstanding, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 SharesOutstanding
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]SharesOutstanding, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 SharesOutstanding
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type YearCount struct {
	Year  int `json:"Year"`
	Count int `json:"Count"`
}

type YearCounts []YearCount

func (out *YearCounts) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]YearCount, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 YearCount
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]YearCount, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 YearCount
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type EarningsInfo struct {
	Date            string           `json:"date"`
	ReportDate      string           `json:"reportDate"`
	EpsActual       *decimal.Decimal `json:"epsActual"`
	EpsEstimate     *decimal.Decimal `json:"epsEstimate"`
	EpsDifference   *decimal.Decimal `json:"epsDifference"`
	SurprisePercent *decimal.Decimal `json:"surprisePercent"`
}

type EarningsInfos []EarningsInfo

func (out *EarningsInfos) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]EarningsInfo, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 EarningsInfo
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]EarningsInfo, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 EarningsInfo
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type EarningsEstimateInfo struct {
	Date                             string           `json:"date"`
	Period                           string           `json:"period"`
	Growth                           *decimal.Decimal `json:"growth"`
	EarningsEstimateAvg              *decimal.Decimal `json:"earningsEstimateAvg"`
	EarningsEstimateLow              *decimal.Decimal `json:"earningsEstimateLow"`
	EarningsEstimateHigh             *decimal.Decimal `json:"earningsEstimateHigh"`
	EarningsEstimateYearAgoEps       *decimal.Decimal `json:"earningsEstimateYearAgoEps"`
	EarningsEstimateNumberOfAnalysts *decimal.Decimal `json:"earningsEstimateNumberOfAnalysts"`
	EarningsEstimateGrowth           *decimal.Decimal `json:"earningsEstimateGrowth"`
	RevenueEstimateAvg               *decimal.Decimal `json:"revenueEstimateAvg"`
	RevenueEstimateLow               *decimal.Decimal `json:"revenueEstimateLow"`
	RevenueEstimateHigh              *decimal.Decimal `json:"revenueEstimateHigh"`
	RevenueEstimateYearAgoEps        *decimal.Decimal `json:"revenueEstimateYearAgoEps"`
	RevenueEstimateNumberOfAnalysts  *decimal.Decimal `json:"revenueEstimateNumberOfAnalysts"`
	RevenueEstimateGrowth            *decimal.Decimal `json:"revenueEstimateGrowth"`
	EpsTrendCurrent                  *decimal.Decimal `json:"epsTrendCurrent"`
	EpsTrend7DaysAgo                 *decimal.Decimal `json:"epsTrend7daysAgo"`
	EpsTrend30DaysAgo                *decimal.Decimal `json:"epsTrend30daysAgo"`
	EpsTrend60DaysAgo                *decimal.Decimal `json:"epsTrend60daysAgo"`
	EpsTrend90DaysAgo                *decimal.Decimal `json:"epsTrend90daysAgo"`
	EpsRevisionsUpLast7Days          *decimal.Decimal `json:"epsRevisionsUpLast7days"`
	EpsRevisionsUpLast30Days         *decimal.Decimal `json:"epsRevisionsUpLast30days"`
	EpsRevisionsDownLast30Days       *decimal.Decimal `json:"epsRevisionsDownLast30days"`
	EpsRevisionsDownLast90Days       *decimal.Decimal `json:"epsRevisionsDownLast90days"`
}

type EarningsEstimateInfos []EarningsEstimateInfo

func (out *EarningsEstimateInfos) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]EarningsEstimateInfo, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 EarningsEstimateInfo
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]EarningsEstimateInfo, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 EarningsEstimateInfo
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type Earnings struct {
	History EarningsInfos         `json:"History"`
	Trend   EarningsEstimateInfos `json:"Trend"`
	Annual  EarningsInfos         `json:"Annual"`
}

type BalanceSheetInfo struct {
	Date       string  `json:"date"`
	FilingDate *string `json:"filing_date"`

	IntangibleAssets                                 *decimal.Decimal `json:"intangibleAssets"`
	TotalLiab                                        *decimal.Decimal `json:"totalLiab"`
	TotalStockholderEquity                           *decimal.Decimal `json:"totalStockholderEquity"`
	DeferredLongTermLiab                             *decimal.Decimal `json:"deferredLongTermLiab"`
	OtherCurrentLiab                                 *decimal.Decimal `json:"otherCurrentLiab"`
	TotalAssets                                      *decimal.Decimal `json:"totalAssets"`
	CommonStock                                      *decimal.Decimal `json:"commonStock"`
	OtherCurrentAssets                               *decimal.Decimal `json:"otherCurrentAssets"`
	RetainedEarnings                                 *decimal.Decimal `json:"retainedEarnings"`
	OtherLiab                                        *decimal.Decimal `json:"otherLiab"`
	GoodWill                                         *decimal.Decimal `json:"goodWill"`
	OtherAssets                                      *decimal.Decimal `json:"otherAssets"`
	Cash                                             *decimal.Decimal `json:"cash"`
	TotalCurrentLiabilities                          *decimal.Decimal `json:"totalCurrentLiabilities"`
	ShortLongTermDebt                                *decimal.Decimal `json:"shortLongTermDebt"`
	OtherStockholderEquity                           *decimal.Decimal `json:"otherStockholderEquity"`
	PropertyPlantEquipment                           *decimal.Decimal `json:"propertyPlantEquipment"`
	TotalCurrentAssets                               *decimal.Decimal `json:"totalCurrentAssets"`
	LongTermInvestments                              *decimal.Decimal `json:"longTermInvestments"`
	NetTangibleAssets                                *decimal.Decimal `json:"netTangibleAssets"`
	ShortTermInvestments                             *decimal.Decimal `json:"shortTermInvestments"`
	NetReceivables                                   *decimal.Decimal `json:"netReceivables"`
	LongTermDebt                                     *decimal.Decimal `json:"longTermDebt"`
	Inventory                                        *decimal.Decimal `json:"inventory"`
	AccountsPayable                                  *decimal.Decimal `json:"accountsPayable"`
	TotalPermanentEquity                             *decimal.Decimal `json:"totalPermanentEquity"`
	NoncontrollingInterestInConsolidatedEntity       *decimal.Decimal `json:"noncontrollingInterestInConsolidatedEntity"`
	TemporaryEquityRedeemableNoncontrollingInterests *decimal.Decimal `json:"temporaryEquityRedeemableNoncontrollingInterests"`
	AccumulatedOtherComprehensiveIncome              *decimal.Decimal `json:"accumulatedOtherComprehensiveIncome"`
	AdditionalPaidInCapital                          *decimal.Decimal `json:"additionalPaidInCapital"`
	CommonStockTotalEquity                           *decimal.Decimal `json:"commonStockTotalEquity"`
	PreferredStockTotalEquity                        *decimal.Decimal `json:"preferredStockTotalEquity"`
	RetainedEarningsTotalEquity                      *decimal.Decimal `json:"retainedEarningsTotalEquity"`
	TreasuryStock                                    *decimal.Decimal `json:"treasuryStock"`
	AccumulatedAmortization                          *decimal.Decimal `json:"accumulatedAmortization"`
	NonCurrrentAssetsOther                           *decimal.Decimal `json:"nonCurrrentAssetsOther"`
	DeferredLongTermAssetCharges                     *decimal.Decimal `json:"deferredLongTermAssetCharges"`
	NonCurrentAssetsTotal                            *decimal.Decimal `json:"nonCurrentAssetsTotal"`
	ShortTermDebt                                    *decimal.Decimal `json:"shortTermDebt"`
	CapitalLeaseObligations                          *decimal.Decimal `json:"capitalLeaseObligations"`
	LongTermDebtTotal                                *decimal.Decimal `json:"longTermDebtTotal"`
	NonCurrentLiabilitiesOther                       *decimal.Decimal `json:"nonCurrentLiabilitiesOther"`
	NonCurrentLiabilitiesTotal                       *decimal.Decimal `json:"nonCurrentLiabilitiesTotal"`
	NegativeGoodwill                                 *decimal.Decimal `json:"negativeGoodwill"`
	Warrants                                         *decimal.Decimal `json:"warrants"`
	PreferredStockRedeemable                         *decimal.Decimal `json:"preferredStockRedeemable"`
	CapitalSurpluse                                  *decimal.Decimal `json:"capitalSurpluse"`
	LiabilitiesAndStockholdersEquity                 *decimal.Decimal `json:"liabilitiesAndStockholdersEquity"`
	CashAndShortTermInvestments                      *decimal.Decimal `json:"cashAndShortTermInvestments"`
	PropertyPlantAndEquipmentGross                   *decimal.Decimal `json:"propertyPlantAndEquipmentGross"`
	AccumulatedDepreciation                          *decimal.Decimal `json:"accumulatedDepreciation"`
	CommonStockSharesOutstanding                     *decimal.Decimal `json:"commonStockSharesOutstanding"`
}

type BalanceSheetInfos []BalanceSheetInfo

func (out *BalanceSheetInfos) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]BalanceSheetInfo, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 BalanceSheetInfo
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]BalanceSheetInfo, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 BalanceSheetInfo
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type CashFlowInfo struct {
	Date                                  string           `json:"date"`
	FilingDate                            *string          `json:"filing_date"`
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
	ChangeReceivables                     *decimal.Decimal `json:"changeReceivables"`
	CashFlowsOtherOperating               *decimal.Decimal `json:"cashFlowsOtherOperating"`
	ExchangeRateChanges                   *decimal.Decimal `json:"exchangeRateChanges"`
	CashAndCashEquivalentsChanges         *decimal.Decimal `json:"cashAndCashEquivalentsChanges"`
}

type CashFlowInfos []CashFlowInfo

func (out *CashFlowInfos) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]CashFlowInfo, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 CashFlowInfo
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]CashFlowInfo, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 CashFlowInfo
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type IncomeStatementInfo struct {
	Date                              string           `json:"date"`
	FilingDate                        *string          `json:"filing_date"`
	ResearchDevelopment               *decimal.Decimal `json:"researchDevelopment"`
	EffectOfAccountingCharges         *decimal.Decimal `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   *decimal.Decimal `json:"incomeBeforeTax"`
	MinorityInterest                  *decimal.Decimal `json:"minorityInterest"`
	NetIncome                         *decimal.Decimal `json:"netIncome"`
	SellingGeneralAdministrative      *decimal.Decimal `json:"sellingGeneralAdministrative"`
	GrossProfit                       *decimal.Decimal `json:"grossProfit"`
	Ebit                              *decimal.Decimal `json:"ebit"`
	NonOperatingIncomeNetOther        *decimal.Decimal `json:"nonOperatingIncomeNetOther"`
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
	PreferredStockAndOtherAdjustments *decimal.Decimal `json:"preferredStockAndOtherAdjustments"`
}

type IncomeStatementInfos []IncomeStatementInfo

func (out *IncomeStatementInfos) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]IncomeStatementInfo, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 IncomeStatementInfo
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]IncomeStatementInfo, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 IncomeStatementInfo
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type Financials struct {
	BalanceSheet    BalanceSheet    `json:"Balance_Sheet"`
	CashFlow        CashFlow        `json:"Cash_Flow"`
	IncomeStatement IncomeStatement `json:"Income_Statement"`
}

type IncomeStatement struct {
	CurrencySymbol string               `json:"currency_symbol"`
	Quarterly      IncomeStatementInfos `json:"quarterly"`
	Yearly         IncomeStatementInfos `json:"yearly"`
}
type BalanceSheet struct {
	CurrencySymbol string            `json:"currency_symbol"`
	Quarterly      BalanceSheetInfos `json:"quarterly"`
	Yearly         BalanceSheetInfos `json:"yearly"`
}
type CashFlow struct {
	CurrencySymbol string        `json:"currency_symbol"`
	Quarterly      CashFlowInfos `json:"quarterly"`
	Yearly         CashFlowInfos `json:"yearly"`
}
type SharesStats struct {
	SharesOutstanding       *decimal.Decimal `json:"SharesOutstanding"`
	SharesFloat             *decimal.Decimal `json:"SharesFloat"`
	PercentInsiders         *decimal.Decimal `json:"PercentInsiders"`
	PercentInstitutions     *decimal.Decimal `json:"PercentInstitutions"`
	SharesShort             *decimal.Decimal `json:"SharesShort"`
	SharesShortPriorMonth   *decimal.Decimal `json:"SharesShortPriorMonth"`
	ShortRatio              *decimal.Decimal `json:"ShortRatio"`
	ShortPercentOutstanding *decimal.Decimal `json:"ShortPercentOutstanding"`
	ShortPercentFloat       *decimal.Decimal `json:"ShortPercentFloat"`
}

type Component struct {
	Code     string `json:"Code"`
	Exchange string `json:"Exchange"`
	Name     string `json:"Name"`
	Sector   string `json:"Sector"`
	Industry string `json:"Industry"`
}

type Components []Component

func (out *Components) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]Component, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 Component
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]Component, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 Component
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type MutualFundData struct {
	Nav                   string                `json:"Nav"`
	PrevClosePrice        string                `json:"Prev_Close_Price"`
	UpdateDate            string                `json:"Update_Date"`
	PortfolioNetAssets    string                `json:"Portfolio_Net_Assets"`
	ShareClassNetAssets   string                `json:"Share_Class_Net_Assets"`
	MorningStarRating     int                   `json:"Morning_Star_Rating"`
	MorningStarRiskRating int                   `json:"Morning_Star_Risk_Rating"`
	MorningStarCategory   string                `json:"Morning_Star_Category"`
	InceptonDate          string                `json:"Incepton_Date"`
	Currency              string                `json:"Currency"`
	Domicile              string                `json:"Domicile"`
	Yield                 string                `json:"Yield"`
	YieldYTD              string                `json:"Yield_YTD"`
	Yield1YearYTD         string                `json:"Yield_1Year_YTD"`
	Yield3YearYTD         string                `json:"Yield_3Year_YTD"`
	Yield5YearYTD         string                `json:"Yield_5Year_YTD"`
	ExpenseRatio          string                `json:"Expense_Ratio"`
	ExpenseRatioDate      string                `json:"Expense_Ratio_Date"`
	AssetAllocation       AssetAllocations      `json:"Asset_Allocation"`
	ValueGrowth           ValueGrowths          `json:"Value_Growth"`
	TopHoldings           TopHoldings           `json:"Top_Holdings"`
	MarketCapitalization  MarketCapitalizations `json:"Market_Capitalization"`
	SectorWeights         SectorWeightsGroup    `json:"Sector_Weights"`
	WorldRegions          RegionWeights         `json:"World_Regions"`
}

type SectorWeightsGroup struct {
	Cyclical  SectorWeights `json:"Cyclical"`
	Defensive SectorWeights `json:"Defensive"`
	Sensitive SectorWeights `json:"Sensitive"`
}

type SectorWeights []SectorWeight

func (out *SectorWeights) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]SectorWeight, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 SectorWeight
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]SectorWeight, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 SectorWeight
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type SectorWeight struct {
	Type            string `json:"Type"`
	CategoryAverage string `json:"Category_Average"`
	Amount          string `json:"Amount_%"`
	Benchmark       string `json:"Benchmark"`
}

type RegionWeight struct {
	Category        string `json:"Category"`
	Name            string `json:"Name"`
	CategoryAverage string `json:"Category_Average"`
	Stocks          string `json:"Stocks_%"`
	Benchmark       string `json:"Benchmark"`
}

type RegionWeights []RegionWeight

func (out *RegionWeights) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]RegionWeight, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 RegionWeight
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]RegionWeight, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()

			in.Delim('{')
			if in.IsDelim('}') {
				continue
			}

			for !in.IsDelim('}') {
				in.Skip()
				in.WantColon()

				var v37 RegionWeight
				(v37).UnmarshalEasyJSON(in)
				v37.Category = key
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim('}')

			in.WantComma()
		}
		in.Delim('}')
	}
}

type MarketCapitalization struct {
	Category        string           `json:"Category"`
	Size            string           `json:"Size"`
	CategoryAverage *decimal.Decimal `json:"Category_Average"`
	Benchmark       *decimal.Decimal `json:"Benchmark"`
	Portfolio       *decimal.Decimal `json:"Portfolio_%"`
}

type MarketCapitalizations []MarketCapitalization

func (out *MarketCapitalizations) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]MarketCapitalization, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 MarketCapitalization
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]MarketCapitalization, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v37 MarketCapitalization
			(v37).UnmarshalEasyJSON(in)
			v37.Category = key
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type TopHolding struct {
	Name   string  `json:"Name"`
	Owned  *string `json:"Owned"`
	Change *string `json:"Change"`
	Weight *string `json:"Weight"`
}

type TopHoldings []TopHolding

func (out *TopHoldings) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]TopHolding, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 TopHolding
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]TopHolding, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 TopHolding
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
}

type ValueGrowth struct {
	Name            string           `json:"Name"`
	CategoryAverage *decimal.Decimal `json:"Category_Average"`
	Benchmark       *decimal.Decimal `json:"Benchmark"`
	StockPortfolio  *decimal.Decimal `json:"Stock_Portfolio"`
}

type ValueGrowths []ValueGrowth

func (out *ValueGrowths) UnmarshalEasyJSON(in *jlexer.Lexer) {
	if in.IsNull() {
		in.Skip()
	} else {
		if in.IsDelim('[') {
			in.Delim('[')
			if !in.IsDelim(']') {
				*out = make([]ValueGrowth, 0)
			} else {
				*out = nil
			}
			for !in.IsDelim(']') {
				var v37 ValueGrowth
				(v37).UnmarshalEasyJSON(in)
				*out = append(*out, v37)
				in.WantComma()
			}
			in.Delim(']')
			return
		}

		in.Delim('{')
		if !in.IsDelim('}') {
			*out = make([]ValueGrowth, 0)
		} else {
			*out = nil
		}
		for !in.IsDelim('}') {
			in.Skip()
			in.WantColon()
			var v37 ValueGrowth
			(v37).UnmarshalEasyJSON(in)
			*out = append(*out, v37)
			in.WantComma()
		}
		in.Delim('}')
	}
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
			offset += newElements
		}
	}

	return nil
}

func buildFundamental(reader *csvReaderMap, exchange *exchanges.Exchange) (Fundamentals, error) {
	var err error
	f := Fundamentals{
		LastUpdate: time.Now(),
		General:    General{},
	}
	err = f.General.fill(reader, "General_")
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
