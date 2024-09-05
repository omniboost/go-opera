package opera

import "encoding/xml"

// TRIALBALANCE was generated 2024-09-05 14:32:20 by https://xml-to-go.github.io/ in Ukraine.
type TrialBalance struct {
	XMLName      xml.Name `xml:"TRIAL_BALANCE"`
	LISTGTRXTYPE struct {
		GTRXTYPE []struct {
			TRXTYPESORT        string `xml:"TRX_TYPE_SORT"`
			TRXTYPE            string `xml:"TRX_TYPE"`
			TRXTYPEDESCRIPTION string `xml:"TRX_TYPE_DESCRIPTION"`
			ISTBTRXTYPE01      string `xml:"IS_TB_TRX_TYPE_01"`
			LISTGTRXCODE       struct {
				GTRXCODE []struct {
					TRXCODESORT          string `xml:"TRX_CODE_SORT"`
					TRXCODE              string `xml:"TRX_CODE"`
					DESCRIPTION          string `xml:"DESCRIPTION"`
					TBAMOUNT             string `xml:"TB_AMOUNT"`
					TRXDATE              string `xml:"TRX_DATE"`
					NETAMOUNT            string `xml:"NET_AMOUNT"`
					GUESTLEDDEBIT        string `xml:"GUEST_LED_DEBIT"`
					GUESTLEDCREDIT       string `xml:"GUEST_LED_CREDIT"`
					ARLEDDEBIT           string `xml:"AR_LED_DEBIT"`
					ARLEDCREDIT          string `xml:"AR_LED_CREDIT"`
					DEPLEDDEBIT          string `xml:"DEP_LED_DEBIT"`
					DEPLEDCREDIT         string `xml:"DEP_LED_CREDIT"`
					PACKAGELEDDEBIT      string `xml:"PACKAGE_LED_DEBIT"`
					PACKAGELEDCREDIT     string `xml:"PACKAGE_LED_CREDIT"`
					INHDEBIT             string `xml:"INH_DEBIT"`
					INHCREDIT            string `xml:"INH_CREDIT"`
					NONREVENUEAMT        string `xml:"NON_REVENUE_AMT"`
					TODAYSACCRUALS       string `xml:"TODAYS_ACCRUALS"`
					DEPOSITATCHECKIN     string `xml:"DEPOSIT_AT_CHECKIN"`
					PACKAGELEDTAX        string `xml:"PACKAGE_LED_TAX"`
					DEPTAXLEDDEBIT       string `xml:"DEP_TAX_LED_DEBIT"`
					OWNERLEDDEBIT        string `xml:"OWNER_LED_DEBIT"`
					OWNERLEDCREDIT       string `xml:"OWNER_LED_CREDIT"`
					DEPLEDDEBITPLCZ      string `xml:"DEP_LED_DEBIT_PL_CZ"`
					LISTGCURRENCYDETAILS string `xml:"LIST_G_CURRENCY_DETAILS"`
					CSCURRENCYCOUNT      string `xml:"CS_CURRENCY_COUNT"`
				} `xml:"G_TRX_CODE"`
			} `xml:"LIST_G_TRX_CODE"`
			CSTBAMOUNTTYPE string `xml:"CS_TB_AMOUNT_TYPE"`
		} `xml:"G_TRX_TYPE"`
	} `xml:"LIST_G_TRX_TYPE"`
	LISTGCHKBALANCE struct {
		GCHKBALANCE []struct {
			CHKBALCODE  string `xml:"CHK_BAL_CODE"`
			CHKBALVALUE string `xml:"CHK_BAL_VALUE"`
		} `xml:"G_CHK_BALANCE"`
	} `xml:"LIST_G_CHK_BALANCE"`
	CSTBAMOUNTREP           string `xml:"CS_TB_AMOUNT_REP"`
	CSGUESTLEDDEBITREP      string `xml:"CS_GUEST_LED_DEBIT_REP"`
	CSGUESTLEDCREDITREP     string `xml:"CS_GUEST_LED_CREDIT_REP"`
	CSARLEDDEBITREP         string `xml:"CS_AR_LED_DEBIT_REP"`
	CSARLEDCREDITREP        string `xml:"CS_AR_LED_CREDIT_REP"`
	CSDEPOSITLEDDEBITREP    string `xml:"CS_DEPOSIT_LED_DEBIT_REP"`
	CSDEPOSITLEDCREDITREP   string `xml:"CS_DEPOSIT_LED_CREDIT_REP"`
	CSPACKAGELEDDEBITREP    string `xml:"CS_PACKAGE_LED_DEBIT_REP"`
	CSPACKAGELEDCREDITREP   string `xml:"CS_PACKAGE_LED_CREDIT_REP"`
	CFGUESTLEDREP           string `xml:"CF_GUEST_LED_REP"`
	CFARLEDREP              string `xml:"CF_AR_LED_REP"`
	CFDEPOSITLEDREP         string `xml:"CF_DEPOSIT_LED_REP"`
	CFPACKAGELEDREP         string `xml:"CF_PACKAGE_LED_REP"`
	CPGUESTLEDDEBITYEST     string `xml:"CP_GUEST_LED_DEBIT_YEST"`
	CPGUESTLEDCREDITYEST    string `xml:"CP_GUEST_LED_CREDIT_YEST"`
	CPARLEDDEBITYEST        string `xml:"CP_AR_LED_DEBIT_YEST"`
	CPARLEDCREDITYEST       string `xml:"CP_AR_LED_CREDIT_YEST"`
	CPDEPOSITLEDDEBITYEST   string `xml:"CP_DEPOSIT_LED_DEBIT_YEST"`
	CPDEPOSITLEDCREDITYEST  string `xml:"CP_DEPOSIT_LED_CREDIT_YEST"`
	CPPACKAGELEDDEBITYEST   string `xml:"CP_PACKAGE_LED_DEBIT_YEST"`
	CPPACKAGELEDCREDITYEST  string `xml:"CP_PACKAGE_LED_CREDIT_YEST"`
	CPTBBALANCEYEST         string `xml:"CP_TB_BALANCE_YEST"`
	CFGUESTLEDYEST          string `xml:"CF_GUEST_LED_YEST"`
	CFCALCPREVIOUSBALANCE   string `xml:"CF_CALC_PREVIOUS_BALANCE"`
	CFARLEDYEST             string `xml:"CF_AR_LED_YEST"`
	CFDEPOSITLEDYEST        string `xml:"CF_DEPOSIT_LED_YEST"`
	CFPACKAGELEDYEST        string `xml:"CF_PACKAGE_LED_YEST"`
	CFTBBALANCETODAY        string `xml:"CF_TB_BALANCE_TODAY"`
	CPINHDEBITYEST          string `xml:"CP_INH_DEBIT_YEST"`
	CPINHCREDITYEST         string `xml:"CP_INH_CREDIT_YEST"`
	CFINHYEST               string `xml:"CF_INH_YEST"`
	CSINHDEBITREP           string `xml:"CS_INH_DEBIT_REP"`
	CSINHCREDITREP          string `xml:"CS_INH_CREDIT_REP"`
	CFINHREP                string `xml:"CF_INH_REP"`
	CFFUTUREACCRUALS        string `xml:"CF_FUTURE_ACCRUALS"`
	CFPACKAGELEDTODAY       string `xml:"CF_PACKAGE_LED_TODAY"`
	CPPACKAGEACCRUALS       string `xml:"CP_PACKAGE_ACCRUALS"`
	CFTBMSG                 string `xml:"CF_TBMSG"`
	CSACCRUALSTODAY         string `xml:"CS_ACCRUALS_TODAY"`
	CFBALANCETODAY          string `xml:"CF_BALANCE_TODAY"`
	CSDEPOSITATCHECKIN      string `xml:"CS_DEPOSIT_AT_CHECKIN"`
	CFBALCARRIEDFORWARD     string `xml:"CF_BAL_CARRIED_FORWARD"`
	CFDEPOSITACTIVITYREP    string `xml:"CF_DEPOSIT_ACTIVITY_REP"`
	CFHOTELLEDGERTODAY      string `xml:"CF_HOTEL_LEDGER_TODAY"`
	CPPACKAGELEDTAXYEST     string `xml:"CP_PACKAGE_LED_TAX_YEST"`
	CFLOGO                  string `xml:"CF_LOGO"`
	CSDEPOSITTAXLEDDEBITREP string `xml:"CS_DEPOSIT_TAX_LED_DEBIT_REP"`
	CSOWNERLEDCREDITREP     string `xml:"CS_OWNER_LED_CREDIT_REP"`
	CSOWNERLEDDEBITREP      string `xml:"CS_OWNER_LED_DEBIT_REP"`
	CFOWNERLEDREP           string `xml:"CF_OWNER_LED_REP"`
	CPOWNERLEDDEBITYEST     string `xml:"CP_OWNER_LED_DEBIT_YEST"`
	CPOWNERLEDCREDITYEST    string `xml:"CP_OWNER_LED_CREDIT_YEST"`
	CFOWNERLEDYEST          string `xml:"CF_OWNER_LED_YEST"`
	CSDEPOSITLEDACTIVITYREP string `xml:"CS_DEPOSIT_LED_ACTIVITY_REP"`
	CPDEPLEDDEBITPLCZ       string `xml:"CP_DEP_LED_DEBIT_PL_CZ"`
} 
