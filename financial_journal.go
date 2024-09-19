package opera

import "encoding/xml"

// FINJRNLBYTRANS was generated 2024-09-05 13:59:27 by https://xml-to-go.github.io/ in Ukraine.
type FinancialJournalByTransaction struct {
	XMLName           xml.Name `xml:"FINJRNLBYTRANS"`
	LISTGISINTERNALYN struct {
		GISINTERNALYN struct {
			ISINTERNALYN string `xml:"IS_INTERNAL_YN"`
			LISTGFIRST   struct {
				GFIRST []struct {
					FIRST      string `xml:"FIRST"`
					LISTSECOND struct {
						SECOND struct {
							SECOND    string `xml:"SECOND"`
							LISTTHIRD struct {
								THIRD struct {
									THIRD            string `xml:"THIRD"`
									LISTGTRXCHARDATE struct {
										GTRXCHARDATE []struct {
											EXPDATE              string `xml:"EXP_DATE"`
											RECEIPTNO            string `xml:"RECEIPT_NO"`
											GUESTFULLNAME        string `xml:"GUEST_FULL_NAME"`
											TARGETRESORT         string `xml:"TARGET_RESORT"`
											TRXDESC              string `xml:"TRX_DESC"`
											MARKETCODE           string `xml:"MARKET_CODE"`
											BUSINESSFORMATDATE   string `xml:"BUSINESS_FORMAT_DATE"`
											BUSINESSTIME         string `xml:"BUSINESS_TIME"`
											BUSINESSDATE         string `xml:"BUSINESS_DATE"`
											REFERENCE            string `xml:"REFERENCE"`
											TRXNO                string `xml:"TRX_NO"`
											CASHIERDEBIT         string `xml:"CASHIER_DEBIT"`
											CASHIERCREDIT        string `xml:"CASHIER_CREDIT"`
											ROOM                 string `xml:"ROOM"`
											CREDITCARDSUPPLEMENT string `xml:"CREDIT_CARD_SUPPLEMENT"`
											CURRENCY1            string `xml:"CURRENCY1"`
											TRXCODE              string `xml:"TRX_CODE"`
											CASHIERID            string `xml:"CASHIER_ID"`
											REMARK               string `xml:"REMARK"`
											INSERTUSER           string `xml:"INSERT_USER"`
											INSERTDATE           string `xml:"INSERT_DATE"`
											CHEQUENUMBER         string `xml:"CHEQUE_NUMBER"`
											ROOMCLASS            string `xml:"ROOM_CLASS"`
											CCCODE               string `xml:"CC_CODE"`
											CASHIERNAME          string `xml:"CASHIER_NAME"`
											USERNAME             string `xml:"USER_NAME"`
											DEPNETTAXAMT         string `xml:"DEP_NET_TAX_AMT"`
											DEPOSITDEBIT         string `xml:"DEPOSIT_DEBIT"`
											CASHIDUSERNAME       string `xml:"CASH_ID_USER_NAME"`
											PRINTCASHIERDEBIT    string `xml:"PRINT_CASHIER_DEBIT"`
											PRINTCASHIERCREDIT   string `xml:"PRINT_CASHIER_CREDIT"`
										} `xml:"G_TRX_CHAR_DATE"`
									} `xml:"LIST_G_TRX_CHAR_DATE"`
									THIRDDEBIT  string `xml:"THIRD_DEBIT"`
									THIRDCREDIT string `xml:"THIRD_CREDIT"`
								} `xml:"THIRD"`
							} `xml:"LIST_THIRD"`
							SECONDDEBIT  string `xml:"SECOND_DEBIT"`
							SECONDCREDIT string `xml:"SECOND_CREDIT"`
						} `xml:"SECOND"`
					} `xml:"LIST_SECOND"`
					FIRSTDEBIT  string `xml:"FIRST_DEBIT"`
					FIRSTCREDIT string `xml:"FIRST_CREDIT"`
				} `xml:"G_FIRST"`
			} `xml:"LIST_G_FIRST"`
			INTERNALDEBIT  string `xml:"INTERNAL_DEBIT"`
			INTERNALCREDIT string `xml:"INTERNAL_CREDIT"`
		} `xml:"G_IS_INTERNAL_YN"`
	} `xml:"LIST_G_IS_INTERNAL_YN"`
	RDEBIT  string `xml:"R_DEBIT"`
	RCREDIT string `xml:"R_CREDIT"`
	LOGO    string `xml:"LOGO"`
}
