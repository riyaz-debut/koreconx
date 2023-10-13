package korecontract

import "time"

// Bond data fields https://github.com/actusfrf/actus-dictionary/blob/master/actus-dictionary-applicability.json
// Principal payment fully at Initial Exchange Date (IED) and repaid at Maturity Date (MD). Fixed and variable rates.
type Bond struct {
	Contract                         string `json:"contract"`
	Calendar                         string `json:"calendar"`
	BusinessDayConvention            string `json:"businessDayConvention"`
	EndOfMonthConvention             string `json:"endOfMonthConvention"`
	ContractType                     string `json:"contractType"`
	StatusDate                       string `json:"statusDate"`
	ContractRole                     string `json:"contractRole"`
	CreatorID                        string `json:"creatorID"`
	ContractID                       string `json:"contractID"`
	MarketObjectCode                 string `json:"marketObjectCode"`
	CounterpartyID                   string `json:"counterpartyID"`
	ContractPerformance              string `json:"contractPerformance"`
	Seniority                        string `json:"seniority"`
	NonPerformingDate                string `json:"nonPerformingDate"`
	PrepaymentPeriod                 string `json:"prepaymentPeriod"`
	GracePeriod                      string `json:"gracePeriod"`
	DelinquencyPeriod                string `json:"delinquencyPeriod"`
	DelinquencyRate                  string `json:"delinquencyRate"`
	CycleAnchorDateOfFee             string `json:"cycleAnchorDateOfFee"`
	CycleOfFee                       string `json:"cycleOfFee"`
	FeeBasis                         string `json:"feeBasis"`
	FeeRate                          string `json:"feeRate"`
	FeeAccrued                       string `json:"feeAccrued"`
	CycleAnchorDateOfInterestPayment string `json:"cycleAnchorDateOfInterestPayment"`
	CycleOfInterestPayment           string `json:"cycleOfInterestPayment"`
	NominalInterestRate              string `json:"nominalInterestRate"`
	DayCountConvention               string `json:"dayCountConvention"`
	AccruedInterest                  string `json:"accruedInterest"`
	CapitalizationEndDate            string `json:"capitalizationEndDate"`
	CyclePointOfInterestPayment      string `json:"cyclePointOfInterestPayment"`
	Currency                         string `json:"currency"`
	ContractDealDate                 string `json:"contractDealDate"`
	InitialExchangeDate              string `json:"initialExchangeDate"`
	PremiumDiscountAtIED             string `json:"premiumDiscountAtIED"`
	MaturityDate                     string `json:"maturityDate"`
	NotionalPrincipal                string `json:"notionalPrincipal"`
	PurchaseDate                     string `json:"purchaseDate"`
	PriceAtPurchaseDate              string `json:"priceAtPurchaseDate"`
	TerminationDate                  string `json:"terminationDate"`
	PriceAtTerminationDate           string `json:"priceAtTerminationDate"`
	CreditLineAmount                 string `json:"creditLineAmount"`
	MarketObjectCodeOfScalingIndex   string `json:"marketObjectCodeOfScalingIndex"`
	ScalingIndexAtStatusDate         string `json:"scalingIndexAtStatusDate"`
	CycleAnchorDateOfScalingIndex    string `json:"cycleAnchorDateOfScalingIndex"`
	CycleOfScalingIndex              string `json:"cycleOfScalingIndex"`
	ScalingEffect                    string `json:"scalingEffect"`
	MarketValueObserved              string `json:"marketValueObserved"`
	OptionExerciseEndDate            string `json:"optionExerciseEndDate"`
	CycleAnchorDateOfOptionality     string `json:"cycleAnchorDateOfOptionality"`
	CycleOfOptionality               string `json:"cycleOfOptionality"`
	PenaltyType                      string `json:"penaltyType"`
	PenaltyRate                      string `json:"penaltyRate"`
	PrepaymentEffect                 string `json:"prepaymentEffect"`
	CycleAnchorDateOfRateReset       string `json:"cycleAnchorDateOfRateReset"`
	CycleOfRateReset                 string `json:"cycleOfRateReset"`
	RateSpread                       string `json:"rateSpread"`
	MarketObjectCodeRateReset        string `json:"marketObjectCodeRateReset"`
	LifeCap                          string `json:"lifeCap"`
	LifeFloor                        string `json:"lifeFloor"`
	PeriodCap                        string `json:"periodCap"`
	PeriodFloor                      string `json:"periodFloor"`
	CyclePointOfRateReset            string `json:"cyclePointOfRateReset"`
	FixingDays                       string `json:"fixingDays"`
	NextResetRate                    string `json:"nextResetRate"`
	RateMultiplier                   string `json:"rateMultiplier"`
	SettlementCurrency               string `json:"settlementCurrency"`
}

// CrowdSafe data fields
type CrowdSafe struct {
	InvestorID        string    `json:"investor_id"`
	GrantorID         string    `json:"grantor_id"`
	DateOfTransaction time.Time `json:"date_of_transaction"`
	CertificateNumber string    `json:"certificate_number"`
	TotalInvestment   int       `json:"total_investment"`
	Term              int       `json:"term"`
	ValuationCap      int       `json:"valuation_cap"`
	DiscountRate      int       `json:"discount_rate"`
	Payment           int       `json:"payment"`
	ExpiryDate        time.Time `json:"expiry_date"`
}

// Debenture data fields
type Debenture struct {
}

// Debt data fields
type Debt struct {
	Bonds           Bond           `json:"bonds"`
	Debentures      Debenture      `json:"debentures"`
	Loans           Loan           `json:"loans"`
	PromissoryNotes PromissoryNote `json:"promissory_notes"`
}

// Derivative data fields
type Derivative struct {
	Options  Option  `json:"options"`
	Warrants Warrant `json:"warrants"`
}

// RightOfDragAlong data fields
type RightOfDragAlong struct {
	Exists                           bool      `json:"exists"`
	SaleExpectedDate                 time.Time `json:"sale_expected_date"`
	SaleNotification                 bool      `json:"sale_notification"`
	SaleNotificationDate             time.Time `json:"sale_notification_date"`
	SaleNotificationResponseDuration int       `json:"sale_notification_response_duration"`
	AmountCalculationMethod          int       `json:"amount_calculation_method"`
	AmountCalculationMethodType      int       `json:"amount_calculation_method_type"`
	DurationBasis                    int       `json:"duration_basis"`
	NotificationresponseDeadline     time.Time `json:"notification_response_deadline"`
	PostResponseSaleDuration         int       `json:"post_response_sale_duration"`
	PostResponseSaleDeadline         time.Time `json:"post_response_sale_deadline"`
}

// Equity data fields
type Equity struct {
	Shares Share `json:"shares"`
	Units  Unit  `json:"units"`
}

// Share data fields
type Share struct {
	ShareHolderAgreementID  string `json:"shareholder_agreement_id"`
	SubscriptionAgreementID string `json:"subscription_agreement_id"`
}

// Unit data fields
type Unit struct {
	ShareHolderAgreementID  string `json:"shareholder_agreement_id"`
	SubscriptionAgreementID string `json:"subscription_agreement_id"`
}

// Loan data fields
type Loan struct {
	InvestorID        string    `json:"investor_id"`
	LenderID          string    `json:"lender_id"`
	DateOfTransaction time.Time `json:"date_of_transaction"`
	CertificateNumber string    `json:"certificate_number"`
	TotalLoan         int       `json:"total_loan"`
	Term              int       `json:"term"`
	Interest          int       `json:"interest"`
	Payment           int       `json:"payment"`
	ExpiryDate        time.Time `json:"expiry_date"`
}

// CallMoney Lonas that are rolled over as long as they are not called. Once called it has to be paid back after the stipulated notice period
type CallMoney struct {
	Contract                         string `json:"contract"`
	Calendar                         string `json:"calendar"`
	BusinessDayConvention            string `json:"businessDayConvention"`
	EndOfMonthConvention             string `json:"endOfMonthConvention"`
	ContractType                     string `json:"contractType"`
	StatusDate                       string `json:"statusDate"`
	ContractRole                     string `json:"contractRole"`
	CreatorID                        string `json:"creatorID"`
	ContractID                       string `json:"contractID"`
	CounterpartyID                   string `json:"counterpartyID"`
	ContractPerformance              string `json:"contractPerformance"`
	Seniority                        string `json:"seniority"`
	NonPerformingDate                string `json:"nonPerformingDate"`
	PrepaymentPeriod                 string `json:"prepaymentPeriod"`
	GracePeriod                      string `json:"gracePeriod"`
	DelinquencyPeriod                string `json:"delinquencyPeriod"`
	DelinquencyRate                  string `json:"delinquencyRate"`
	CycleAnchorDateOfFee             string `json:"cycleAnchorDateOfFee"`
	CycleOfFee                       string `json:"cycleOfFee"`
	FeeBasis                         string `json:"feeBasis"`
	FeeRate                          string `json:"feeRate"`
	FeeAccrued                       string `json:"feeAccrued"`
	CycleAnchorDateOfInterestPayment string `json:"cycleAnchorDateOfInterestPayment"`
	CycleOfInterestPayment           string `json:"cycleOfInterestPayment"`
	NominalInterestRate              string `json:"nominalInterestRate"`
	DayCountConvention               string `json:"dayCountConvention"`
	AccruedInterest                  string `json:"accruedInterest"`
	Currency                         string `json:"currency"`
	ContractDealDate                 string `json:"contractDealDate"`
	InitialExchangeDate              string `json:"initialExchangeDate"`
	MaturityDate                     string `json:"maturityDate"`
	NotionalPrincipal                string `json:"notionalPrincipal"`
	XDayNotice                       string `json:"xDayNotice"`
	CycleAnchorDateOfRateReset       string `json:"cycleAnchorDateOfRateReset"`
	CycleOfRateReset                 string `json:"cycleOfRateReset"`
	RateSpread                       string `json:"rateSpread"`
	MarketObjectCodeRateReset        string `json:"marketObjectCodeRateReset"`
	FixingDays                       string `json:"fixingDays"`
	NextResetRate                    string `json:"nextResetRate"`
	RateMultiplier                   string `json:"rateMultiplier"`
	SettlementCurrency               string `json:"settlementCurrency"`
}

// ExoticLinearAmortizer Exotic version of LAM. However step ups with respect to (i) Principal, (ii) Interest rates are possible. Highly flexible to match totally irregular principal payments. Principal can also be paid out in steps.
type ExoticLinearAmortizer struct {
	Contract                                  string `json:"contract"`
	Calendar                                  string `json:"calendar"`
	BusinessDayConvention                     string `json:"businessDayConvention"`
	EndOfMonthConvention                      string `json:"endOfMonthConvention"`
	ContractType                              string `json:"contractType"`
	StatusDate                                string `json:"statusDate"`
	ContractRole                              string `json:"contractRole"`
	CreatorID                                 string `json:"creatorID"`
	ContractID                                string `json:"contractID"`
	MarketObjectCode                          string `json:"marketObjectCode"`
	CounterpartyID                            string `json:"counterpartyID"`
	ContractPerformance                       string `json:"contractPerformance"`
	Seniority                                 string `json:"seniority"`
	NonPerformingDate                         string `json:"nonPerformingDate"`
	PrepaymentPeriod                          string `json:"prepaymentPeriod"`
	GracePeriod                               string `json:"gracePeriod"`
	DelinquencyPeriod                         string `json:"delinquencyPeriod"`
	DelinquencyRate                           string `json:"delinquencyRate"`
	CycleAnchorDateOfFee                      string `json:"cycleAnchorDateOfFee"`
	CycleOfFee                                string `json:"cycleOfFee"`
	FeeBasis                                  string `json:"feeBasis"`
	FeeRate                                   string `json:"feeRate"`
	FeeAccrued                                string `json:"feeAccrued"`
	ArrayCycleAnchorDateOfInterestPayment     string `json:"arrayCycleAnchorDateOfInterestPayment"`
	ArrayCycleOfInterestPayment               string `json:"arrayCycleOfInterestPayment"`
	NominalInterestRate                       string `json:"nominalInterestRate"`
	DayCountConvention                        string `json:"dayCountConvention"`
	AccruedInterest                           string `json:"accruedInterest"`
	CapitalizationEndDate                     string `json:"capitalizationEndDate"`
	CycleAnchorDateOfInterestCalculationBase  string `json:"cycleAnchorDateOfInterestCalculationBase"`
	CycleOfInterestCalculationBase            string `json:"cycleOfInterestCalculationBase"`
	InterestCalculationBase                   string `json:"interestCalculationBase"`
	InterestCalculationBaseAmount             string `json:"interestCalculationBaseAmount"`
	CyclePointOfInterestPayment               string `json:"cyclePointOfInterestPayment"`
	Currency                                  string `json:"currency"`
	ContractDealDate                          string `json:"contractDealDate"`
	InitialExchangeDate                       string `json:"initialExchangeDate"`
	PremiumDiscountAtIED                      string `json:"premiumDiscountAtIED"`
	MaturityDate                              string `json:"maturityDate"`
	NotionalPrincipal                         string `json:"notionalPrincipal"`
	ArrayCycleAnchorDateOfPrincipalRedemption string `json:"arrayCycleAnchorDateOfPrincipalRedemption"`
	ArrayCycleOfPrincipalRedemption           string `json:"arrayCycleOfPrincipalRedemption"`
	ArrayNextPrincipalRedemptionPayment       string `json:"arrayNextPrincipalRedemptionPayment"`
	ArrayIncreaseDecrease                     string `json:"arrayIncreaseDecrease"`
	PurchaseDate                              string `json:"purchaseDate"`
	PriceAtPurchaseDate                       string `json:"priceAtPurchaseDate"`
	TerminationDate                           string `json:"terminationDate"`
	PriceAtTerminationDate                    string `json:"priceAtTerminationDate"`
	MarketObjectCodeOfScalingIndex            string `json:"marketObjectCodeOfScalingIndex"`
	ScalingIndexAtStatusDate                  string `json:"scalingIndexAtStatusDate"`
	CycleAnchorDateOfScalingIndex             string `json:"cycleAnchorDateOfScalingIndex"`
	CycleOfScalingIndex                       string `json:"cycleOfScalingIndex"`
	ScalingEffect                             string `json:"scalingEffect"`
	MarketValueObserved                       string `json:"marketValueObserved"`
	OptionExerciseEndDate                     string `json:"optionExerciseEndDate"`
	CycleAnchorDateOfOptionality              string `json:"cycleAnchorDateOfOptionality"`
	CycleOfOptionality                        string `json:"cycleOfOptionality"`
	PenaltyType                               string `json:"penaltyType"`
	PenaltyRate                               string `json:"penaltyRate"`
	PrepaymentEffect                          string `json:"prepaymentEffect"`
	ArrayCycleAnchorDateOfRateReset           string `json:"arrayCycleAnchorDateOfRateReset"`
	ArrayCycleOfRateReset                     string `json:"arrayCycleOfRateReset"`
	ArrayRate                                 string `json:"arrayRate"`
	ArrayFixedVariable                        string `json:"arrayFixedVariable"`
	MarketObjectCodeRateReset                 string `json:"marketObjectCodeRateReset"`
	LifeCap                                   string `json:"lifeCap"`
	LifeFloor                                 string `json:"lifeFloor"`
	PeriodCap                                 string `json:"periodCap"`
	PeriodFloor                               string `json:"periodFloor"`
	CyclePointOfRateReset                     string `json:"cyclePointOfRateReset"`
	FixingDays                                string `json:"fixingDays"`
	NextResetRate                             string `json:"nextResetRate"`
	RateMultiplier                            string `json:"rateMultiplier"`
	SettlementCurrency                        string `json:"settlementCurrency"`
}

// LinearAmortizer Principal payment fully at IED. Principal repaid periodically in constant amounts till MD. Interest gets reduced accordingly. If variable rate, only interest payment is recalculated. Fixed and variable rates.
type LinearAmortizer struct {
	Contract                                 string `json:"contract"`
	Calendar                                 string `json:"calendar"`
	BusinessDayConvention                    string `json:"businessDayConvention"`
	EndOfMonthConvention                     string `json:"endOfMonthConvention"`
	ContractType                             string `json:"contractType"`
	StatusDate                               string `json:"statusDate"`
	ContractRole                             string `json:"contractRole"`
	CreatorID                                string `json:"creatorID"`
	ContractID                               string `json:"contractID"`
	MarketObjectCode                         string `json:"marketObjectCode"`
	CounterpartyID                           string `json:"counterpartyID"`
	ContractPerformance                      string `json:"contractPerformance"`
	Seniority                                string `json:"seniority"`
	NonPerformingDate                        string `json:"nonPerformingDate"`
	PrepaymentPeriod                         string `json:"prepaymentPeriod"`
	GracePeriod                              string `json:"gracePeriod"`
	DelinquencyPeriod                        string `json:"delinquencyPeriod"`
	DelinquencyRate                          string `json:"delinquencyRate"`
	CycleAnchorDateOfFee                     string `json:"cycleAnchorDateOfFee"`
	CycleOfFee                               string `json:"cycleOfFee"`
	FeeBasis                                 string `json:"feeBasis"`
	FeeRate                                  string `json:"feeRate"`
	FeeAccrued                               string `json:"feeAccrued"`
	CycleAnchorDateOfInterestPayment         string `json:"cycleAnchorDateOfInterestPayment"`
	CycleOfInterestPayment                   string `json:"cycleOfInterestPayment"`
	NominalInterestRate                      string `json:"nominalInterestRate"`
	DayCountConvention                       string `json:"dayCountConvention"`
	AccruedInterest                          string `json:"accruedInterest"`
	CapitalizationEndDate                    string `json:"capitalizationEndDate"`
	CycleAnchorDateOfInterestCalculationBase string `json:"cycleAnchorDateOfInterestCalculationBase"`
	CycleOfInterestCalculationBase           string `json:"cycleOfInterestCalculationBase"`
	InterestCalculationBase                  string `json:"interestCalculationBase"`
	InterestCalculationBaseAmount            string `json:"interestCalculationBaseAmount"`
	CyclePointOfInterestPayment              string `json:"cyclePointOfInterestPayment"`
	Currency                                 string `json:"currency"`
	ContractDealDate                         string `json:"contractDealDate"`
	InitialExchangeDate                      string `json:"initialExchangeDate"`
	PremiumDiscountAtIED                     string `json:"premiumDiscountAtIED"`
	MaturityDate                             string `json:"maturityDate"`
	NotionalPrincipal                        string `json:"notionalPrincipal"`
	CycleAnchorDateOfPrincipalRedemption     string `json:"cycleAnchorDateOfPrincipalRedemption"`
	CycleOfPrincipalRedemption               string `json:"cycleOfPrincipalRedemption"`
	NextPrincipalRedemptionPayment           string `json:"nextPrincipalRedemptionPayment"`
	PurchaseDate                             string `json:"purchaseDate"`
	PriceAtPurchaseDate                      string `json:"priceAtPurchaseDate"`
	TerminationDate                          string `json:"terminationDate"`
	PriceAtTerminationDate                   string `json:"priceAtTerminationDate"`
	CreditLineAmount                         string `json:"creditLineAmount"`
	MarketObjectCodeOfScalingIndex           string `json:"marketObjectCodeOfScalingIndex"`
	ScalingIndexAtStatusDate                 string `json:"scalingIndexAtStatusDate"`
	CycleAnchorDateOfScalingIndex            string `json:"cycleAnchorDateOfScalingIndex"`
	CycleOfScalingIndex                      string `json:"cycleOfScalingIndex"`
	ScalingEffect                            string `json:"scalingEffect"`
	MarketValueObserved                      string `json:"marketValueObserved"`
	OptionExerciseEndDate                    string `json:"optionExerciseEndDate"`
	CycleAnchorDateOfOptionality             string `json:"cycleAnchorDateOfOptionality"`
	CycleOfOptionality                       string `json:"cycleOfOptionality"`
	PenaltyType                              string `json:"penaltyType"`
	PenaltyRate                              string `json:"penaltyRate"`
	PrepaymentEffect                         string `json:"prepaymentEffect"`
	CycleAnchorDateOfRateReset               string `json:"cycleAnchorDateOfRateReset"`
	CycleOfRateReset                         string `json:"cycleOfRateReset"`
	RateSpread                               string `json:"rateSpread"`
	MarketObjectCodeRateReset                string `json:"marketObjectCodeRateReset"`
	LifeCap                                  string `json:"lifeCap"`
	LifeFloor                                string `json:"lifeFloor"`
	PeriodCap                                string `json:"periodCap"`
	PeriodFloor                              string `json:"periodFloor"`
	CyclePointOfRateReset                    string `json:"cyclePointOfRateReset"`
	FixingDays                               string `json:"fixingDays"`
	NextResetRate                            string `json:"nextResetRate"`
	RateMultiplier                           string `json:"rateMultiplier"`
	SettlementCurrency                       string `json:"settlementCurrency"`
}

// NegativeAmortizer Similar as ANN. However when resetting rate, total amount (interest plus principal) stay constant. MD shifts. Only variable rates
type NegativeAmortizer struct {
	Contract                                 string `json:"contract"`
	Calendar                                 string `json:"calendar"`
	BusinessDayConvention                    string `json:"businessDayConvention"`
	EndOfMonthConvention                     string `json:"endOfMonthConvention"`
	ContractType                             string `json:"contractType"`
	StatusDate                               string `json:"statusDate"`
	ContractRole                             string `json:"contractRole"`
	CreatorID                                string `json:"creatorID"`
	ContractID                               string `json:"contractID"`
	MarketObjectCode                         string `json:"marketObjectCode"`
	CounterpartyID                           string `json:"counterpartyID"`
	ContractPerformance                      string `json:"contractPerformance"`
	Seniority                                string `json:"seniority"`
	NonPerformingDate                        string `json:"nonPerformingDate"`
	PrepaymentPeriod                         string `json:"prepaymentPeriod"`
	GracePeriod                              string `json:"gracePeriod"`
	DelinquencyPeriod                        string `json:"delinquencyPeriod"`
	DelinquencyRate                          string `json:"delinquencyRate"`
	CycleAnchorDateOfFee                     string `json:"cycleAnchorDateOfFee"`
	CycleOfFee                               string `json:"cycleOfFee"`
	FeeBasis                                 string `json:"feeBasis"`
	FeeRate                                  string `json:"feeRate"`
	FeeAccrued                               string `json:"feeAccrued"`
	CycleAnchorDateOfInterestPayment         string `json:"cycleAnchorDateOfInterestPayment"`
	CycleOfInterestPayment                   string `json:"cycleOfInterestPayment"`
	NominalInterestRate                      string `json:"nominalInterestRate"`
	DayCountConvention                       string `json:"dayCountConvention"`
	AccruedInterest                          string `json:"accruedInterest"`
	CapitalizationEndDate                    string `json:"capitalizationEndDate"`
	CycleAnchorDateOfInterestCalculationBase string `json:"cycleAnchorDateOfInterestCalculationBase"`
	CycleOfInterestCalculationBase           string `json:"cycleOfInterestCalculationBase"`
	InterestCalculationBase                  string `json:"interestCalculationBase"`
	InterestCalculationBaseAmount            string `json:"interestCalculationBaseAmount"`
	Currency                                 string `json:"currency"`
	ContractDealDate                         string `json:"contractDealDate"`
	InitialExchangeDate                      string `json:"initialExchangeDate"`
	PremiumDiscountAtIED                     string `json:"premiumDiscountAtIED"`
	MaturityDate                             string `json:"maturityDate"`
	NotionalPrincipal                        string `json:"notionalPrincipal"`
	CycleAnchorDateOfPrincipalRedemption     string `json:"cycleAnchorDateOfPrincipalRedemption"`
	CycleOfPrincipalRedemption               string `json:"cycleOfPrincipalRedemption"`
	NextPrincipalRedemptionPayment           string `json:"nextPrincipalRedemptionPayment"`
	PurchaseDate                             string `json:"purchaseDate"`
	PriceAtPurchaseDate                      string `json:"priceAtPurchaseDate"`
	TerminationDate                          string `json:"terminationDate"`
	PriceAtTerminationDate                   string `json:"priceAtTerminationDate"`
	CreditLineAmount                         string `json:"creditLineAmount"`
	MarketObjectCodeOfScalingIndex           string `json:"marketObjectCodeOfScalingIndex"`
	ScalingIndexAtStatusDate                 string `json:"scalingIndexAtStatusDate"`
	CycleAnchorDateOfScalingIndex            string `json:"cycleAnchorDateOfScalingIndex"`
	CycleOfScalingIndex                      string `json:"cycleOfScalingIndex"`
	ScalingEffect                            string `json:"scalingEffect"`
	MarketValueObserved                      string `json:"marketValueObserved"`
	OptionExerciseEndDate                    string `json:"optionExerciseEndDate"`
	CycleAnchorDateOfOptionality             string `json:"cycleAnchorDateOfOptionality"`
	CycleOfOptionality                       string `json:"cycleOfOptionality"`
	PenaltyType                              string `json:"penaltyType"`
	PenaltyRate                              string `json:"penaltyRate"`
	PrepaymentEffect                         string `json:"prepaymentEffect"`
	CycleAnchorDateOfRateReset               string `json:"cycleAnchorDateOfRateReset"`
	CycleOfRateReset                         string `json:"cycleOfRateReset"`
	RateSpread                               string `json:"rateSpread"`
	MarketObjectCodeRateReset                string `json:"marketObjectCodeRateReset"`
	LifeCap                                  string `json:"lifeCap"`
	LifeFloor                                string `json:"lifeFloor"`
	PeriodCap                                string `json:"periodCap"`
	PeriodFloor                              string `json:"periodFloor"`
	FixingDays                               string `json:"fixingDays"`
	NextResetRate                            string `json:"nextResetRate"`
	RateMultiplier                           string `json:"rateMultiplier"`
	SettlementCurrency                       string `json:"settlementCurrency"`
}

// Option data fields
type Option struct {
	InvestorID         string    `json:"investor_id"`
	GrantorID          string    `json:"grantor_id"`
	DateOfTransaction  time.Time `json:"date_of_transaction"`
	CertificateNumber  string    `json:"certificate_number"`
	NumberOfSecurities int       `json:"number_of_securities"`
	Term               int       `json:"term"`
	PricePerSecurity   float64   `json:"price_per_security"`
	VestingPeriod      int       `json:"vesting_period"`
	ExpiryDate         time.Time `json:"expiry_date"`
}

// Options data fields https://github.com/actusfrf/actus-dictionary/blob/master/actus-dictionary-applicability.json
// Calculates straight option pay-off for any basic CT as underlying (PAM, ANN etc.) but also SWAPS, FXOUT, STK and COM. Single, periodic and continuous strike is supported.
type Options struct {
	Contract                     string `json:"contract"`
	Calendar                     string `json:"calendar"`
	BusinessDayConvention        string `json:"businessDayConvention"`
	EndOfMonthConvention         string `json:"endOfMonthConvention"`
	ContractType                 string `json:"contractType"`
	StatusDate                   string `json:"statusDate"`
	ContractRole                 string `json:"contractRole"`
	CreatorID                    string `json:"creatorID"`
	ContractID                   string `json:"contractID"`
	MarketObjectCode             string `json:"marketObjectCode"`
	ContractStructure            string `json:"contractStructure"`
	CounterpartyID               string `json:"counterpartyID"`
	ContractPerformance          string `json:"contractPerformance"`
	Seniority                    string `json:"seniority"`
	NonPerformingDate            string `json:"nonPerformingDate"`
	GracePeriod                  string `json:"gracePeriod"`
	DelinquencyPeriod            string `json:"delinquencyPeriod"`
	DelinquencyRate              string `json:"delinquencyRate"`
	Currency                     string `json:"currency"`
	ContractDealDate             string `json:"contractDealDate"`
	MaturityDate                 string `json:"maturityDate"`
	PurchaseDate                 string `json:"purchaseDate"`
	PriceAtPurchaseDate          string `json:"priceAtPurchaseDate"`
	TerminationDate              string `json:"terminationDate"`
	PriceAtTerminationDate       string `json:"priceAtTerminationDate"`
	MarketValueObserved          string `json:"marketValueObserved"`
	OptionExecutionType          string `json:"optionExecutionType"`
	OptionExerciseEndDate        string `json:"optionExerciseEndDate"`
	OptionStrike1                string `json:"optionStrike1"`
	OptionStrike2                string `json:"optionStrike2"`
	OptionType                   string `json:"optionType"`
	CycleAnchorDateOfOptionality string `json:"cycleAnchorDateOfOptionality"`
	CycleOfOptionality           string `json:"cycleOfOptionality"`
	ExerciseDate                 string `json:"exerciseDate"`
	ExerciseAmount               string `json:"exerciseAmount"`
	SettlementDays               string `json:"settlementDays"`
	DeliverySettlement           string `json:"deliverySettlement"`
	SettlementCurrency           string `json:"settlementCurrency"`
}

// CapFloor data fields
// Interest rate option expressed in a maximum or minimum interest rate.
type CapFloor struct {
	Contract               string `json:"contract"`
	ContractType           string `json:"contractType"`
	StatusDate             string `json:"statusDate"`
	ContractRole           string `json:"contractRole"`
	CreatorID              string `json:"creatorID"`
	ContractID             string `json:"contractID"`
	MarketObjectCode       string `json:"marketObjectCode"`
	ContractStructure      string `json:"contractStructure"`
	CounterpartyID         string `json:"counterpartyID"`
	ContractPerformance    string `json:"contractPerformance"`
	Seniority              string `json:"seniority"`
	NonPerformingDate      string `json:"nonPerformingDate"`
	GracePeriod            string `json:"gracePeriod"`
	DelinquencyPeriod      string `json:"delinquencyPeriod"`
	DelinquencyRate        string `json:"delinquencyRate"`
	Currency               string `json:"currency"`
	ContractDealDate       string `json:"contractDealDate"`
	PurchaseDate           string `json:"purchaseDate"`
	PriceAtPurchaseDate    string `json:"priceAtPurchaseDate"`
	TerminationDate        string `json:"terminationDate"`
	PriceAtTerminationDate string `json:"priceAtTerminationDate"`
	MarketValueObserved    string `json:"marketValueObserved"`
	LifeCap                string `json:"lifeCap"`
	LifeFloor              string `json:"lifeFloor"`
	SettlementCurrency     string `json:"settlementCurrency"`
}

// PromissoryNote data fields
type PromissoryNote struct {
	InvestorID          string            `json:"investor_id"`
	NoteMakerID         string            `json:"note_maker_id"`
	DateOfTransaction   time.Time         `json:"date_of_transaction"`
	CertificateNumber   string            `json:"certificate_number"`
	TotalPromissoryNote int               `json:"total_promissory_note"`
	Term                int               `json:"term"`
	Interest            int               `json:"interest"`
	PaymentSchedule     []PaymentSchedule `json:"payment_schedule"`
	ExpiryDate          time.Time         `json:"expiry_date"`
	Convert             int               `json:"convert"`
	ValuationCap        int               `json:"valuation_cap"`
	DiscountRate        int               `json:"discount_rate"`
	NumberOfShares      int               `json:"number_of_shares"`
	PricePerShare       int               `json:"price_per_share"`
	ClassOfShares       string            `json:"class_of_shares"`
	ConversionEvent     string            `json:"conversion_event"`
}

// PaymentSchedule data fields
type PaymentSchedule struct {
	InstallmentNumber string    `json:"installment_number"`
	PaymentDate       time.Time `json:"payment_date"`
	PrincipalAmount   int       `json:"principal_amount"`
	InterestAmount    int       `json:"interest_amount"`
}

// RightOfFirstRefusal struct
type RightOfFirstRefusal struct {
	Exists                               bool `json:"exists"`                                   //does an FRR exist in this KoreContract?
	AllocationBasis                      int  `json:"allocation_basis"`                         //1: do not normalize by excluding selling Holder 2: normalize by excluding selling Holder
	TransfereeIdentity                   bool `json:"transferee_identity"`                      //does the selling shareholder have to disclose the identity of the proposed transferee? Default: yes
	SecondaryMarketOption                bool `json:"secondary_market_option"`                  //can the selling shareholder have the option to sell refused shares in the secondary market?
	TransferApprovalbyCompany            bool `json:"transfer_approvalby_company"`              //must the Company approve the transfer of refused shares to the transferee? Default: yes
	SecondaryMarketSaleApprovalbyCompany bool `json:"secondary_market_sale_approvalby_company"` //must the Company approve the sale of refused shares in the secondary market? Default: yes
	ResponseDuration                     int  `json:"response_duration"`                        //number of days within which the company or other shareholders are required to respond with an accept or reject the decision of each entity (company, individual, other holding org such as an institutional shareholder, etc.) is to be recorded on the world state
	ResponseDurationBasis                int  `json:"response_duration_basis"`                  //1: business days 2: calendar days
	PaymentInstrument                    int  `json:"payment_instrument"`                       //1: cash 2: non-cash consideration 3: FMV-equivalent of non-cash consideration
	PostResponseTransferDuration         int  `json:"post_response_transfer_duration"`          //number of days within which the selling shareholder is required to complete the transfer of all shares (or put it for sale on a secondary market) that were refused by the company or its other shareholders
	NewHolderBinding                     bool `json:"new_holder_binding"`                       //must the new holder (transferee or buyer) be obligated to be bound by the terms and conditions of the holding or shareholders’ agreement? Default: yes
	FailedTransferNotification           bool `json:"failed_transfer_notification"`             //must the transferring or selling shareholder provide notice to Company in case the refused shares will be withdrawn from the offering (either through transfer, sale, or for any other reason)? Default: yes
	SubsequentNotifications              bool `json:"subsequent_notifications"`                 //after the first FRR transaction, must any subsequent transfer or sale transactions by the same shareholder go through the same FRR process? Default: yes

	//Below are flags or data fields that deal with exclusions to the FRR provision:
	ExclusionFamily                      bool `json:"exclusion_family"`                        //Is FRR necessary if transferring to a family member or to a trust for family?
	ExclusionAffiliate                   bool `json:"exclusion_affiliate"`                     //Is FRR necessary if transferring to holder’s affiliate?
	ExclusionEquityHolderofHolder        bool `json:"exclusion_equity_holderof_holder"`        //Is FRR necessary if transferring to an equity holder of the Holder?
	ExclusionPartner                     bool `json:"exclusion_partner"`                       //Is FRR necessary if transferring to a partner of the Holder?
	ExclusionTransactionMethod           int  `json:"exclusion_transaction_method"`            //What was the reason or method of this allowed transfer (where FRR is waived)? Examples are 1: sale, 2: gift, 3: intestate, 4: inclusion in Will, 5: “like” exchange
	ExclusionNotificationDurationBasis   int  `json:"exclusion_notification_duration_basis"`   //1: business days 2: calendar days
	ExclusionNotificationMinimumDuration int  `json:"exclusion_notification_minimum_duration"` //Transferring Holder should notify Company of intended FRR-excluded transfer no less than this minimum number of days in advance
	ExclusionNotificationMaximumDuration int  `json:"exclusion_notification_maximum_duration"` //Transferring Holder should notify Company of intended FRR-excluded transfer no more than this maximum number of days in advance

	//General restrictions to transfers must be considered during the FRR process
	TransferRestrictionCompetitor                    bool      `json:"transfer_restriction_competitor"`                      //Are transfers to competitors allowed or not?
	TransferRestrictionCompetitorWaiver              int       `json:"transfer_restriction_competitor_waiver"`               //1: Never waived 2: Waived for defined scenarios
	TransferRestrictionCompetitorWaiverScenarios     []int     `json:"transfer_restriction_competitor_waiver_scenarios"`     //If waived, multiple values can be selected: 1. Waived in case of merger 2: Waived in case of exit, dissolution 3: Waived in case of bankruptcy process 4: Waived in case of regulatory action (what types?) 5: Waived through shareholder or Board approval
	TransferRestrictionCancellationReason            int       `json:"transfer_restriction_cancellation_reason"`             //Transfer restrictions may be cancelled upon various conditions. For example, 1: Termination of Shareholders’ Agreement, 2: Dissolution or Bankruptcy of the Company, 3: Consummation of Public Offering, 4: A pre-determined date as set forth in the shareholders’ agreement.
	TransferRestrictionCancellationDeterminationDate time.Time `json:"transfer_restriction_cancellation_determination_date"` //Date on which the determiniation was made to cancel transfer restrictions
	TransferRestrictionCancellationEffectiveDate     time.Time `json:"transfer_restriction_cancellation_effective_date"`     //Date on which the cancellation of transfer restrictions takes effect
}

// RightOfTagAlong data fields
type RightOfTagAlong struct {
	Exists                           bool      `json:"exists"`
	SaleExpectedDate                 time.Time `json:"sale_expected_date"`
	SaleNotification                 bool      `json:"sale_notification"`
	SaleNotificationDate             time.Time `json:"sale_notification_date"`
	SaleNotificationResponseDuration int       `json:"sale_notification_response_duration"`
	AmountCalculationMethod          int       `json:"amount_calculation_method"`
	AmountCalculationMethodType      int       `json:"amount_calculation_method_type"`
	DurationBasis                    int       `json:"duration_basis"`
	NotificationresponseDeadline     time.Time `json:"notification_response_deadline"`
	PostResponseSaleDuration         int       `json:"post_response_sale_duration"`
	PostResponseSaleDeadline         time.Time `json:"post_response_sale_deadline"`
}

// Warrant data fields
type Warrant struct {
	InvestorID         string    `json:"investor_id"`
	GrantorID          string    `json:"grantor_id"`
	DateOfTransaction  time.Time `json:"date_of_transaction"`
	CertificateNumber  string    `json:"certificate_number"`
	NumberOfSecurities int       `json:"number_of_securities"`
	PricePerSecurity   float64   `json:"price_per_security"`
	ExpiryDate         time.Time `json:"expiry_date"`
}
