package opera

import (
	"encoding/xml"
	"time"
)

// FINJRNLBYTRANS was generated 2024-09-05 13:59:27 by https://xml-to-go.github.io/ in Ukraine.
type FinancialJournalByTransaction struct {
	Finjrnlbytrans    xml.Name `xml:"FINJRNLBYTRANS"`
	ListGIsInternalYn struct {
		GIsInternalYn struct {
			IsInternalYn string `xml:"IS_INTERNAL_YN"`
			ListGFirst   struct {
				GFirst []struct {
					First      string `xml:"FIRST"`
					ListSecond struct {
						Second struct {
							Second    string `xml:"SECOND"`
							ListThird struct {
								Third struct {
									Third            string `xml:"THIRD"`
									ListGTrxCharDate struct {
										GTrxCharDate []struct {
											ExpDate              string  `xml:"EXP_DATE"`
											ReceiptNo            string  `xml:"RECEIPT_NO"`
											GuestFullName        string  `xml:"GUEST_FULL_NAME"`
											TargetResort         string  `xml:"TARGET_RESORT"`
											TrxDesc              string  `xml:"TRX_DESC"`
											MarketCode           string  `xml:"MARKET_CODE"`
											BusinessFormatDate   Date    `xml:"BUSINESS_FORMAT_DATE"`
											BusinessTime         string  `xml:"BUSINESS_TIME"`
											BusinessDate         Date    `xml:"BUSINESS_DATE"`
											Reference            string  `xml:"REFERENCE"`
											TrxNo                string  `xml:"TRX_NO"`
											CashierDebit         float64 `xml:"CASHIER_DEBIT"`
											CashierCredit        float64 `xml:"CASHIER_CREDIT"`
											Room                 string  `xml:"ROOM"`
											CreditCardSupplement string  `xml:"CREDIT_CARD_SUPPLEMENT"`
											Currency1            string  `xml:"CURRENCY1"`
											TrxCode              string  `xml:"TRX_CODE"`
											CashierId            string  `xml:"CASHIER_ID"`
											Remark               string  `xml:"REMARK"`
											InsertUser           string  `xml:"INSERT_USER"`
											InsertDate           Date    `xml:"INSERT_DATE"`
											ChequeNumber         string  `xml:"CHEQUE_NUMBER"`
											RoomClass            string  `xml:"ROOM_CLASS"`
											CcCode               string  `xml:"CC_CODE"`
											CashierName          string  `xml:"CASHIER_NAME"`
											UserName             string  `xml:"USER_NAME"`
											DepNetTaxAmt         float64 `xml:"DEP_NET_TAX_AMT"`
											DepositDebit         string  `xml:"DEPOSIT_DEBIT"`
											CashIdUserName       string  `xml:"CASH_ID_USER_NAME"`
											PrintCashierDebit    float64 `xml:"PRINT_CASHIER_DEBIT"`
											PrintCashierCredit   float64 `xml:"PRINT_CASHIER_CREDIT"`
										} `xml:"G_TRX_CHAR_DATE"`
									} `xml:"LIST_G_TRX_CHAR_DATE"`
									ThirdDebit  float64 `xml:"THIRD_DEBIT"`
									ThirdCredit float64 `xml:"THIRD_CREDIT"`
								} `xml:"THIRD"`
							} `xml:"LIST_THIRD"`
							SecondDebit  float64 `xml:"SECOND_DEBIT"`
							SecondCredit float64 `xml:"SECOND_CREDIT"`
						} `xml:"SECOND"`
					} `xml:"LIST_SECOND"`
					FirstDebit  float64 `xml:"FIRST_DEBIT"`
					FirstCredit float64 `xml:"FIRST_CREDIT"`
				} `xml:"G_FIRST"`
			} `xml:"LIST_G_FIRST"`
			InternalDebit  float64 `xml:"INTERNAL_DEBIT"`
			InternalCredit float64 `xml:"INTERNAL_CREDIT"`
		} `xml:"G_IS_INTERNAL_YN"`
	} `xml:"LIST_G_IS_INTERNAL_YN"`
	RDebit  float64 `xml:"R_DEBIT"`
	RCredit float64 `xml:"R_CREDIT"`
	Logo    string  `xml:"LOGO"`
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
