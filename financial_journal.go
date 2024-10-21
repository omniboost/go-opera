package opera

import (
	"encoding/xml"
	"time"
)

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
											EXPDATE              string  `xml:"EXP_DATE"`
											RECEIPTNO            string  `xml:"RECEIPT_NO"`
											GUESTFULLNAME        string  `xml:"GUEST_FULL_NAME"`
											TARGETRESORT         string  `xml:"TARGET_RESORT"`
											TRXDESC              string  `xml:"TRX_DESC"`
											MARKETCODE           string  `xml:"MARKET_CODE"`
											BUSINESSFORMATDATE   Date    `xml:"BUSINESS_FORMAT_DATE"`
											BUSINESSTIME         string  `xml:"BUSINESS_TIME"`
											BUSINESSDATE         Date    `xml:"BUSINESS_DATE"`
											REFERENCE            string  `xml:"REFERENCE"`
											TRXNO                string  `xml:"TRX_NO"`
											CASHIERDEBIT         float64 `xml:"CASHIER_DEBIT"`
											CASHIERCREDIT        float64 `xml:"CASHIER_CREDIT"`
											ROOM                 string  `xml:"ROOM"`
											CREDITCARDSUPPLEMENT string  `xml:"CREDIT_CARD_SUPPLEMENT"`
											CURRENCY1            string  `xml:"CURRENCY1"`
											TRXCODE              string  `xml:"TRX_CODE"`
											CASHIERID            string  `xml:"CASHIER_ID"`
											REMARK               string  `xml:"REMARK"`
											INSERTUSER           string  `xml:"INSERT_USER"`
											INSERTDATE           Date    `xml:"INSERT_DATE"`
											CHEQUENUMBER         string  `xml:"CHEQUE_NUMBER"`
											ROOMCLASS            string  `xml:"ROOM_CLASS"`
											CCCODE               string  `xml:"CC_CODE"`
											CASHIERNAME          string  `xml:"CASHIER_NAME"`
											USERNAME             string  `xml:"USER_NAME"`
											DEPNETTAXAMT         float64 `xml:"DEP_NET_TAX_AMT"`
											DEPOSITDEBIT         string  `xml:"DEPOSIT_DEBIT"`
											CASHIDUSERNAME       string  `xml:"CASH_ID_USER_NAME"`
											PRINTCASHIERDEBIT    float64 `xml:"PRINT_CASHIER_DEBIT"`
											PRINTCASHIERCREDIT   float64 `xml:"PRINT_CASHIER_CREDIT"`
										} `xml:"G_TRX_CHAR_DATE"`
									} `xml:"LIST_G_TRX_CHAR_DATE"`
									THIRDDEBIT  float64 `xml:"THIRD_DEBIT"`
									THIRDCREDIT float64 `xml:"THIRD_CREDIT"`
								} `xml:"THIRD"`
							} `xml:"LIST_THIRD"`
							SECONDDEBIT  float64 `xml:"SECOND_DEBIT"`
							SECONDCREDIT float64 `xml:"SECOND_CREDIT"`
						} `xml:"SECOND"`
					} `xml:"LIST_SECOND"`
					FIRSTDEBIT  float64 `xml:"FIRST_DEBIT"`
					FIRSTCREDIT float64 `xml:"FIRST_CREDIT"`
				} `xml:"G_FIRST"`
			} `xml:"LIST_G_FIRST"`
			INTERNALDEBIT  float64 `xml:"INTERNAL_DEBIT"`
			INTERNALCREDIT float64 `xml:"INTERNAL_CREDIT"`
		} `xml:"G_IS_INTERNAL_YN"`
	} `xml:"LIST_G_IS_INTERNAL_YN"`
	RDEBIT  float64 `xml:"R_DEBIT"`
	RCREDIT float64 `xml:"R_CREDIT"`
	LOGO    string  `xml:"LOGO"`
}

type Date struct {
	time.Time
}

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	layout := "02-Jan-06"
	nt, err := time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
		return nil
	}

	layout = "01-02-06"
	nt, err = time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
		return nil
	}

	layout = "2006-01-02"
	nt, err = time.Parse(layout, s)
	if err == nil {
		*t = Date{Time: nt}
		return nil
	}

	return err
}

func (t Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}

type FormatDate struct {
	time.Time
}
