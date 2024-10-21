package opera

import "encoding/xml"

// TRIALBALANCE was generated 2024-09-05 14:32:20 by https://xml-to-go.github.io/ in Ukraine.
type TrialBalance struct {
	XMLName      xml.Name `xml:"TRIAL_BALANCE"`
	ListGTrxType struct {
		GTrxType []struct {
			TrxTypeSort        string `xml:"TRX_TYPE_SORT"`
			TrxType            string `xml:"TRX_TYPE"`
			TrxTypeDescription string `xml:"TRX_TYPE_DESCRIPTION"`
			IsTbTrxType01      string `xml:"IS_TB_TRX_TYPE_01"`
			ListGTrxCode       struct {
				GTrxCode []struct {
					TrxCodeSort          string  `xml:"TRX_CODE_SORT"`
					TrxCode              string  `xml:"TRX_CODE"`
					Description          string  `xml:"DESCRIPTION"`
					TbAmount             float64 `xml:"TB_AMOUNT"`
					TrxDate              Date    `xml:"TRX_DATE"`
					NetAmount            float64 `xml:"NET_AMOUNT"`
					GuestLedDebit        float64 `xml:"GUEST_LED_DEBIT"`
					GuestLedCredit       float64 `xml:"GUEST_LED_CREDIT"`
					ArLedDebit           float64 `xml:"AR_LED_DEBIT"`
					ArLedCredit          float64 `xml:"AR_LED_CREDIT"`
					DepLedDebit          float64 `xml:"DEP_LED_DEBIT"`
					DepLedCredit         float64 `xml:"DEP_LED_CREDIT"`
					PackageLedDebit      float64 `xml:"PACKAGE_LED_DEBIT"`
					PackageLedCredit     float64 `xml:"PACKAGE_LED_CREDIT"`
					InhDebit             float64 `xml:"INH_DEBIT"`
					InhCredit            float64 `xml:"INH_CREDIT"`
					NonRevenueAmt        float64 `xml:"NON_REVENUE_AMT"`
					TodaysAccruals       float64 `xml:"TODAYS_ACCRUALS"`
					DepositAtCheckin     float64 `xml:"DEPOSIT_AT_CHECKIN"`
					PackageLedTax        float64 `xml:"PACKAGE_LED_TAX"`
					DepTaxLedDebit       float64 `xml:"DEP_TAX_LED_DEBIT"`
					OwnerLedDebit        float64 `xml:"OWNER_LED_DEBIT"`
					OwnerLedCredit       float64 `xml:"OWNER_LED_CREDIT"`
					DepLedDebitPlCz      float64 `xml:"DEP_LED_DEBIT_PL_CZ"`
					ListGCurrencyDetails string  `xml:"LIST_G_CURRENCY_DETAILS"`
					CsCurrencyCount      float64 `xml:"CS_CURRENCY_COUNT"`
				} `xml:"G_TRX_CODE"`
			} `xml:"LIST_G_TRX_CODE"`
			CsTbAmountType float64 `xml:"CS_TB_AMOUNT_TYPE"`
		} `xml:"G_TRX_TYPE"`
	} `xml:"LIST_G_TRX_TYPE"`
	ListGChkBalance struct {
		GChkBalance []struct {
			ChkBalCode  string `xml:"CHK_BAL_CODE"`
			ChkBalValue string `xml:"CHK_BAL_VALUE"`
		} `xml:"G_CHK_BALANCE"`
	} `xml:"LIST_G_CHK_BALANCE"`
	CsTbAmountRep           float64 `xml:"CS_TB_AMOUNT_REP"`
	CsGuestLedDebitRep      float64 `xml:"CS_GUEST_LED_DEBIT_REP"`
	CsGuestLedCreditRep     float64 `xml:"CS_GUEST_LED_CREDIT_REP"`
	CsArLedDebitRep         float64 `xml:"CS_AR_LED_DEBIT_REP"`
	CsArLedCreditRep        float64 `xml:"CS_AR_LED_CREDIT_REP"`
	CsDepositLedDebitRep    float64 `xml:"CS_DEPOSIT_LED_DEBIT_REP"`
	CsDepositLedCreditRep   float64 `xml:"CS_DEPOSIT_LED_CREDIT_REP"`
	CsPackageLedDebitRep    float64 `xml:"CS_PACKAGE_LED_DEBIT_REP"`
	CsPackageLedCreditRep   float64 `xml:"CS_PACKAGE_LED_CREDIT_REP"`
	CfGuestLedRep           float64 `xml:"CF_GUEST_LED_REP"`
	CfArLedRep              float64 `xml:"CF_AR_LED_REP"`
	CfDepositLedRep         float64 `xml:"CF_DEPOSIT_LED_REP"`
	CfPackageLedRep         float64 `xml:"CF_PACKAGE_LED_REP"`
	CpGuestLedDebitYest     float64 `xml:"CP_GUEST_LED_DEBIT_YEST"`
	CpGuestLedCreditYest    float64 `xml:"CP_GUEST_LED_CREDIT_YEST"`
	CpArLedDebitYest        float64 `xml:"CP_AR_LED_DEBIT_YEST"`
	CpArLedCreditYest       float64 `xml:"CP_AR_LED_CREDIT_YEST"`
	CpDepositLedDebitYest   float64 `xml:"CP_DEPOSIT_LED_DEBIT_YEST"`
	CpDepositLedCreditYest  float64 `xml:"CP_DEPOSIT_LED_CREDIT_YEST"`
	CpPackageLedDebitYest   float64 `xml:"CP_PACKAGE_LED_DEBIT_YEST"`
	CpPackageLedCreditYest  float64 `xml:"CP_PACKAGE_LED_CREDIT_YEST"`
	CpTbBalanceYest         float64 `xml:"CP_TB_BALANCE_YEST"`
	CfGuestLedYest          float64 `xml:"CF_GUEST_LED_YEST"`
	CfCalcPreviousBalance   float64 `xml:"CF_CALC_PREVIOUS_BALANCE"`
	CfArLedYest             float64 `xml:"CF_AR_LED_YEST"`
	CfDepositLedYest        float64 `xml:"CF_DEPOSIT_LED_YEST"`
	CfPackageLedYest        float64 `xml:"CF_PACKAGE_LED_YEST"`
	CfTbBalanceToday        float64 `xml:"CF_TB_BALANCE_TODAY"`
	CpInhDebitYest          float64 `xml:"CP_INH_DEBIT_YEST"`
	CpInhCreditYest         float64 `xml:"CP_INH_CREDIT_YEST"`
	CfInhYest               float64 `xml:"CF_INH_YEST"`
	CsInhDebitRep           float64 `xml:"CS_INH_DEBIT_REP"`
	CsInhCreditRep          float64 `xml:"CS_INH_CREDIT_REP"`
	CfInhRep                float64 `xml:"CF_INH_REP"`
	CfFutureAccruals        float64 `xml:"CF_FUTURE_ACCRUALS"`
	CfPackageLedToday       float64 `xml:"CF_PACKAGE_LED_TODAY"`
	CpPackageAccruals       float64 `xml:"CP_PACKAGE_ACCRUALS"`
	CfTbmsg                 string  `xml:"CF_TBMSG"`
	CsAccrualsToday         float64 `xml:"CS_ACCRUALS_TODAY"`
	CfBalanceToday          float64 `xml:"CF_BALANCE_TODAY"`
	CsDepositAtCheckin      float64 `xml:"CS_DEPOSIT_AT_CHECKIN"`
	CfBalCarriedForward     float64 `xml:"CF_BAL_CARRIED_FORWARD"`
	CfDepositActivityRep    float64 `xml:"CF_DEPOSIT_ACTIVITY_REP"`
	CfHotelLedgerToday      float64 `xml:"CF_HOTEL_LEDGER_TODAY"`
	CpPackageLedTaxYest     float64 `xml:"CP_PACKAGE_LED_TAX_YEST"`
	CfLogo                  float64 `xml:"CF_LOGO"`
	CsDepositTaxLedDebitRep float64 `xml:"CS_DEPOSIT_TAX_LED_DEBIT_REP"`
	CsOwnerLedCreditRep     float64 `xml:"CS_OWNER_LED_CREDIT_REP"`
	CsOwnerLedDebitRep      float64 `xml:"CS_OWNER_LED_DEBIT_REP"`
	CfOwnerLedRep           float64 `xml:"CF_OWNER_LED_REP"`
	CpOwnerLedDebitYest     float64 `xml:"CP_OWNER_LED_DEBIT_YEST"`
	CpOwnerLedCreditYest    float64 `xml:"CP_OWNER_LED_CREDIT_YEST"`
	CfOwnerLedYest          float64 `xml:"CF_OWNER_LED_YEST"`
	CsDepositLedActivityRep float64 `xml:"CS_DEPOSIT_LED_ACTIVITY_REP"`
	CpDepLedDebitPlCz       float64 `xml:"CP_DEP_LED_DEBIT_PL_CZ"`
}
