package eodhdapi

//go:generate go run github.com/mailru/easyjson/easyjson -omit_empty -disallow_unknown_fields -all fundamentals.go

import (
	"fmt"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/shopspring/decimal"
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
	MarketCapitalization       *Decimal `json:"MarketCapitalization"`
	MarketCapitalizationMln    string   `json:"MarketCapitalizationMln"`
	EBITDA                     *Decimal `json:"EBITDA"`
	PERatio                    *Decimal `json:"PERatio"`
	PEGRatio                   *Decimal `json:"PEGRatio"`
	WallStreetTargetPrice      *Decimal `json:"WallStreetTargetPrice"`
	BookValue                  *Decimal `json:"BookValue"`
	DividendShare              *Decimal `json:"DividendShare"`
	DividendYield              *Decimal `json:"DividendYield"`
	EarningsShare              *Decimal `json:"EarningsShare"`
	EPSEstimateCurrentYear     *Decimal `json:"EPSEstimateCurrentYear"`
	EPSEstimateNextYear        *Decimal `json:"EPSEstimateNextYear"`
	EPSEstimateNextQuarter     *Decimal `json:"EPSEstimateNextQuarter"`
	MostRecentQuarter          string   `json:"MostRecentQuarter"`
	ProfitMargin               *Decimal `json:"ProfitMargin"`
	OperatingMarginTTM         *Decimal `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          *Decimal `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          *Decimal `json:"ReturnOnEquityTTM"`
	RevenueTTM                 *Decimal `json:"RevenueTTM"`
	RevenuePerShareTTM         *Decimal `json:"RevenuePerShareTTM"`
	QuarterlyRevenueGrowthYOY  *Decimal `json:"QuarterlyRevenueGrowthYOY"`
	GrossProfitTTM             *Decimal `json:"GrossProfitTTM"`
	DilutedEpsTTM              *Decimal `json:"DilutedEpsTTM"`
	QuarterlyEarningsGrowthYOY *Decimal `json:"QuarterlyEarningsGrowthYOY"`
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
	ThreeYVolatility *Decimal `json:"3y_Volatility"`
	ThreeYExpReturn  *Decimal `json:"3y_ExpReturn"`
	ThreeYSharpRatio *Decimal `json:"3y_SharpRatio"`
	ReturnsYTD       *Decimal `json:"Returns_YTD"`
	Returns3Y        *Decimal `json:"Returns_3Y"`
	Returns5Y        *Decimal `json:"Returns_5Y"`
	Returns10Y       *Decimal `json:"Returns_10Y"`
}

type MorningStar struct {
	Ratio               string `json:"Ratio"`
	CategoryBenchmark   string `json:"Category_Benchmark"`
	SustainabilityRatio string `json:"Sustainability_Ratio,string"`
}

type Holding struct {
	Name                  string   `json:"Name"`
	Country               string   `json:"Country"`
	AssetsPercent         *Decimal "json:\"Assets_%\""
	AssetsBackTickPercent *Decimal "json:\"Assets_`%\""
}
type Weight struct {
	Category           string   `json:"Category"`
	EquityPercent      string   `json:"Equity_%"`
	RelativeToCategory *Decimal `json:"Relative_to_Category"`
}

type Allocation struct {
	Category         string   `json:"Category"`
	LongPercent      *Decimal `json:"Long_%"`
	ShortPercent     *Decimal `json:"Short_%"`
	NetAssetsPercent *Decimal `json:"Net_Assets_%"`
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
	Type            string   `json:"Type"`
	Net             *Decimal `json:"Net_%"`
	Long            *Decimal `json:"Long_%"`
	Short           *Decimal `json:"Short_%"`
	CategoryAverage *Decimal `json:"Category_Average"`
	Benchmark       *Decimal `json:"Benchmark"`
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
	TrailingPE             *Decimal `json:"TrailingPE"`
	ForwardPE              *Decimal `json:"ForwardPE"`
	PriceSalesTTM          *Decimal `json:"PriceSalesTTM"`
	PriceBookMRQ           *Decimal `json:"PriceBookMRQ"`
	EnterpriseValueRevenue *Decimal `json:"EnterpriseValueRevenue"`
	EnterpriseValueEbitda  *Decimal `json:"EnterpriseValueEbitda"`
}
type Technicals struct {
	Beta                  *Decimal `json:"Beta"`
	FiftyTwoWeekHigh      *Decimal `json:"52WeekHigh"`
	FiftyTwoWeekLow       *Decimal `json:"52WeekLow"`
	FiftyDayMA            *Decimal `json:"50DayMA"`
	TwoHundredDayMA       *Decimal `json:"200DayMA"`
	SharesShort           *Decimal `json:"SharesShort"`
	SharesShortPriorMonth *Decimal `json:"SharesShortPriorMonth"`
	ShortRatio            *Decimal `json:"ShortRatio"`
	ShortPercent          *Decimal `json:"ShortPercent"`
}
type SplitsDividends struct {
	ForwardAnnualDividendRate  *Decimal   `json:"ForwardAnnualDividendRate"`
	ForwardAnnualDividendYield *Decimal   `json:"ForwardAnnualDividendYield"`
	PayoutRatio                *Decimal   `json:"PayoutRatio"`
	DividendDate               string     `json:"DividendDate"`
	ExDividendDate             string     `json:"ExDividendDate"`
	LastSplitFactor            string     `json:"LastSplitFactor"`
	LastSplitDate              string     `json:"LastSplitDate"`
	NumberDividendsByYear      YearCounts `json:"NumberDividendsByYear"`
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
	Date            string   `json:"date"`
	ReportDate      string   `json:"reportDate"`
	EpsActual       *Decimal `json:"epsActual"`
	EpsEstimate     *Decimal `json:"epsEstimate"`
	EpsDifference   *Decimal `json:"epsDifference"`
	SurprisePercent *Decimal `json:"surprisePercent"`
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
	Date                             string   `json:"date"`
	Period                           string   `json:"period"`
	Growth                           *Decimal `json:"growth"`
	EarningsEstimateAvg              *Decimal `json:"earningsEstimateAvg"`
	EarningsEstimateLow              *Decimal `json:"earningsEstimateLow"`
	EarningsEstimateHigh             *Decimal `json:"earningsEstimateHigh"`
	EarningsEstimateYearAgoEps       *Decimal `json:"earningsEstimateYearAgoEps"`
	EarningsEstimateNumberOfAnalysts *Decimal `json:"earningsEstimateNumberOfAnalysts"`
	EarningsEstimateGrowth           *Decimal `json:"earningsEstimateGrowth"`
	RevenueEstimateAvg               *Decimal `json:"revenueEstimateAvg"`
	RevenueEstimateLow               *Decimal `json:"revenueEstimateLow"`
	RevenueEstimateHigh              *Decimal `json:"revenueEstimateHigh"`
	RevenueEstimateYearAgoEps        *Decimal `json:"revenueEstimateYearAgoEps"`
	RevenueEstimateNumberOfAnalysts  *Decimal `json:"revenueEstimateNumberOfAnalysts"`
	RevenueEstimateGrowth            *Decimal `json:"revenueEstimateGrowth"`
	EpsTrendCurrent                  *Decimal `json:"epsTrendCurrent"`
	EpsTrend7DaysAgo                 *Decimal `json:"epsTrend7daysAgo"`
	EpsTrend30DaysAgo                *Decimal `json:"epsTrend30daysAgo"`
	EpsTrend60DaysAgo                *Decimal `json:"epsTrend60daysAgo"`
	EpsTrend90DaysAgo                *Decimal `json:"epsTrend90daysAgo"`
	EpsRevisionsUpLast7Days          *Decimal `json:"epsRevisionsUpLast7days"`
	EpsRevisionsUpLast30Days         *Decimal `json:"epsRevisionsUpLast30days"`
	EpsRevisionsDownLast30Days       *Decimal `json:"epsRevisionsDownLast30days"`
	EpsRevisionsDownLast90Days       *Decimal `json:"epsRevisionsDownLast90days"`
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

	IntangibleAssets                                 *Decimal `json:"intangibleAssets"`
	TotalLiab                                        *Decimal `json:"totalLiab"`
	TotalStockholderEquity                           *Decimal `json:"totalStockholderEquity"`
	DeferredLongTermLiab                             *Decimal `json:"deferredLongTermLiab"`
	OtherCurrentLiab                                 *Decimal `json:"otherCurrentLiab"`
	TotalAssets                                      *Decimal `json:"totalAssets"`
	CommonStock                                      *Decimal `json:"commonStock"`
	OtherCurrentAssets                               *Decimal `json:"otherCurrentAssets"`
	RetainedEarnings                                 *Decimal `json:"retainedEarnings"`
	OtherLiab                                        *Decimal `json:"otherLiab"`
	GoodWill                                         *Decimal `json:"goodWill"`
	OtherAssets                                      *Decimal `json:"otherAssets"`
	Cash                                             *Decimal `json:"cash"`
	TotalCurrentLiabilities                          *Decimal `json:"totalCurrentLiabilities"`
	ShortLongTermDebt                                *Decimal `json:"shortLongTermDebt"`
	OtherStockholderEquity                           *Decimal `json:"otherStockholderEquity"`
	PropertyPlantEquipment                           *Decimal `json:"propertyPlantEquipment"`
	TotalCurrentAssets                               *Decimal `json:"totalCurrentAssets"`
	LongTermInvestments                              *Decimal `json:"longTermInvestments"`
	NetTangibleAssets                                *Decimal `json:"netTangibleAssets"`
	ShortTermInvestments                             *Decimal `json:"shortTermInvestments"`
	NetReceivables                                   *Decimal `json:"netReceivables"`
	LongTermDebt                                     *Decimal `json:"longTermDebt"`
	Inventory                                        *Decimal `json:"inventory"`
	AccountsPayable                                  *Decimal `json:"accountsPayable"`
	TotalPermanentEquity                             *Decimal `json:"totalPermanentEquity"`
	NoncontrollingInterestInConsolidatedEntity       *Decimal `json:"noncontrollingInterestInConsolidatedEntity"`
	TemporaryEquityRedeemableNoncontrollingInterests *Decimal `json:"temporaryEquityRedeemableNoncontrollingInterests"`
	AccumulatedOtherComprehensiveIncome              *Decimal `json:"accumulatedOtherComprehensiveIncome"`
	AdditionalPaidInCapital                          *Decimal `json:"additionalPaidInCapital"`
	CommonStockTotalEquity                           *Decimal `json:"commonStockTotalEquity"`
	PreferredStockTotalEquity                        *Decimal `json:"preferredStockTotalEquity"`
	RetainedEarningsTotalEquity                      *Decimal `json:"retainedEarningsTotalEquity"`
	TreasuryStock                                    *Decimal `json:"treasuryStock"`
	AccumulatedAmortization                          *Decimal `json:"accumulatedAmortization"`
	NonCurrrentAssetsOther                           *Decimal `json:"nonCurrrentAssetsOther"`
	DeferredLongTermAssetCharges                     *Decimal `json:"deferredLongTermAssetCharges"`
	NonCurrentAssetsTotal                            *Decimal `json:"nonCurrentAssetsTotal"`
	ShortTermDebt                                    *Decimal `json:"shortTermDebt"`
	CapitalLeaseObligations                          *Decimal `json:"capitalLeaseObligations"`
	LongTermDebtTotal                                *Decimal `json:"longTermDebtTotal"`
	NonCurrentLiabilitiesOther                       *Decimal `json:"nonCurrentLiabilitiesOther"`
	NonCurrentLiabilitiesTotal                       *Decimal `json:"nonCurrentLiabilitiesTotal"`
	NegativeGoodwill                                 *Decimal `json:"negativeGoodwill"`
	Warrants                                         *Decimal `json:"warrants"`
	PreferredStockRedeemable                         *Decimal `json:"preferredStockRedeemable"`
	CapitalSurpluse                                  *Decimal `json:"capitalSurpluse"`
	LiabilitiesAndStockholdersEquity                 *Decimal `json:"liabilitiesAndStockholdersEquity"`
	CashAndShortTermInvestments                      *Decimal `json:"cashAndShortTermInvestments"`
	PropertyPlantAndEquipmentGross                   *Decimal `json:"propertyPlantAndEquipmentGross"`
	AccumulatedDepreciation                          *Decimal `json:"accumulatedDepreciation"`
	CommonStockSharesOutstanding                     *Decimal `json:"commonStockSharesOutstanding"`
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
	Date                                  string   `json:"date"`
	FilingDate                            *string  `json:"filing_date"`
	Investments                           *Decimal `json:"investments"`
	ChangeToLiabilities                   *Decimal `json:"changeToLiabilities"`
	TotalCashflowsFromInvestingActivities *Decimal `json:"totalCashflowsFromInvestingActivities"`
	NetBorrowings                         *Decimal `json:"netBorrowings"`
	TotalCashFromFinancingActivities      *Decimal `json:"totalCashFromFinancingActivities"`
	ChangeToOperatingActivities           *Decimal `json:"changeToOperatingActivities"`
	NetIncome                             *Decimal `json:"netIncome"`
	ChangeInCash                          *Decimal `json:"changeInCash"`
	TotalCashFromOperatingActivities      *Decimal `json:"totalCashFromOperatingActivities"`
	Depreciation                          *Decimal `json:"depreciation"`
	OtherCashflowsFromInvestingActivities *Decimal `json:"otherCashflowsFromInvestingActivities"`
	DividendsPaid                         *Decimal `json:"dividendsPaid"`
	ChangeToInventory                     *Decimal `json:"changeToInventory"`
	ChangeToAccountReceivables            *Decimal `json:"changeToAccountReceivables"`
	SalePurchaseOfStock                   *Decimal `json:"salePurchaseOfStock"`
	OtherCashflowsFromFinancingActivities *Decimal `json:"otherCashflowsFromFinancingActivities"`
	ChangeToNetincome                     *Decimal `json:"changeToNetincome"`
	CapitalExpenditures                   *Decimal `json:"capitalExpenditures"`
	ChangeReceivables                     *Decimal `json:"changeReceivables"`
	CashFlowsOtherOperating               *Decimal `json:"cashFlowsOtherOperating"`
	ExchangeRateChanges                   *Decimal `json:"exchangeRateChanges"`
	CashAndCashEquivalentsChanges         *Decimal `json:"cashAndCashEquivalentsChanges"`
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
	Date                              string   `json:"date"`
	FilingDate                        *string  `json:"filing_date"`
	ResearchDevelopment               *Decimal `json:"researchDevelopment"`
	EffectOfAccountingCharges         *Decimal `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   *Decimal `json:"incomeBeforeTax"`
	MinorityInterest                  *Decimal `json:"minorityInterest"`
	NetIncome                         *Decimal `json:"netIncome"`
	SellingGeneralAdministrative      *Decimal `json:"sellingGeneralAdministrative"`
	GrossProfit                       *Decimal `json:"grossProfit"`
	Ebit                              *Decimal `json:"ebit"`
	NonOperatingIncomeNetOther        *Decimal `json:"nonOperatingIncomeNetOther"`
	OperatingIncome                   *Decimal `json:"operatingIncome"`
	OtherOperatingExpenses            *Decimal `json:"otherOperatingExpenses"`
	InterestExpense                   *Decimal `json:"interestExpense"`
	ExtraordinaryItems                *Decimal `json:"extraordinaryItems"`
	NonRecurring                      *Decimal `json:"nonRecurring"`
	OtherItems                        *Decimal `json:"otherItems"`
	IncomeTaxExpense                  *Decimal `json:"incomeTaxExpense"`
	TotalRevenue                      *Decimal `json:"totalRevenue"`
	TotalOperatingExpenses            *Decimal `json:"totalOperatingExpenses"`
	CostOfRevenue                     *Decimal `json:"costOfRevenue"`
	TotalOtherIncomeExpenseNet        *Decimal `json:"totalOtherIncomeExpenseNet"`
	DiscontinuedOperations            *Decimal `json:"discontinuedOperations"`
	NetIncomeFromContinuingOps        *Decimal `json:"netIncomeFromContinuingOps"`
	NetIncomeApplicableToCommonShares *Decimal `json:"netIncomeApplicableToCommonShares"`
	PreferredStockAndOtherAdjustments *Decimal `json:"preferredStockAndOtherAdjustments"`
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
	SharesOutstanding       *Decimal `json:"SharesOutstanding"`
	SharesFloat             *Decimal `json:"SharesFloat"`
	PercentInsiders         *Decimal `json:"PercentInsiders"`
	PercentInstitutions     *Decimal `json:"PercentInstitutions"`
	SharesShort             *Decimal `json:"SharesShort"`
	SharesShortPriorMonth   *Decimal `json:"SharesShortPriorMonth"`
	ShortRatio              *Decimal `json:"ShortRatio"`
	ShortPercentOutstanding *Decimal `json:"ShortPercentOutstanding"`
	ShortPercentFloat       *Decimal `json:"ShortPercentFloat"`
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
	MorningStarRating     Rating                `json:"Morning_Star_Rating"`
	MorningStarRiskRating Rating                `json:"Morning_Star_Risk_Rating"`
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
	Category        string   `json:"Category"`
	Size            string   `json:"Size"`
	CategoryAverage *Decimal `json:"Category_Average"`
	Benchmark       *Decimal `json:"Benchmark"`
	Portfolio       *Decimal `json:"Portfolio_%"`
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
	Name            string   `json:"Name"`
	CategoryAverage *Decimal `json:"Category_Average"`
	Benchmark       *Decimal `json:"Benchmark"`
	StockPortfolio  *Decimal `json:"Stock_Portfolio"`
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

type Rating string

func (r *Rating) UnmarshalEasyJSON(in *jlexer.Lexer) {
	raw := in.Raw()
	str, err := unquoteIfQuoted(raw)
	if err != nil {
		in.AddError(fmt.Errorf("error decoding string '%s': %s", raw, err))
	}
	*r = Rating(str)
}

func (r Rating) MarshalEasyJSON(w *jwriter.Writer) {
	w.String(string(r))
}

type Decimal decimal.Decimal

func (dec *Decimal) UnmarshalEasyJSON(in *jlexer.Lexer) {

	decimalBytes := in.Raw()

	str, err := unquoteIfQuoted(decimalBytes)
	if err != nil {
		in.AddError(fmt.Errorf("error decoding string '%s': %s", decimalBytes, err))
	}

	if str == "null" || str == "-" {
		return
	}

	d, err := decimal.NewFromString(str)

	if err != nil {
		if strings.Count(str, ",") == 1 {
			d, err = decimal.NewFromString(strings.Replace(str, ",", ".", 1))
		}
	}

	if err != nil {
		in.AddError(err)
		return
	}

	*dec = Decimal(d)
}

func (dec Decimal) MarshalEasyJSON(w *jwriter.Writer) {
	w.String(dec.D().String())
}

// D gets the casted decimal.Decimal
func (dec *Decimal) D() *decimal.Decimal {
	if dec == nil {
		return nil
	}

	d := decimal.Decimal(*dec)
	return &d
}

func unquoteIfQuoted(value interface{}) (string, error) {
	var bytes []byte

	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return "", fmt.Errorf("Could not convert value '%+v' to byte array of type '%T'",
			value, value)
	}

	// If the amount is quoted, strip the quotes
	if len(bytes) > 2 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}
	return string(bytes), nil
}
