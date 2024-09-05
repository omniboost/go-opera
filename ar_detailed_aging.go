package opera

import "encoding/xml"

// ARAGINGDET was generated 2024-09-05 13:37:58 by https://xml-to-go.github.io/ in Ukraine.
type ARDetailedAging struct {
	XMLName          xml.Name `xml:"ARAGINGDET"`
	LISTGACCOUNTTYPE struct {
		GACCOUNTTYPE struct {
			ACCOUNTT         string `xml:"ACCOUNT_T"`
			ACCOUNTTYPEID    string `xml:"ACCOUNT_TYPE_ID"`
			LISTGACCOUNTNAME struct {
				GACCOUNTNAME []struct {
					ORDERCOL     string `xml:"ORDER_COL"`
					ACCOUNTNO    string `xml:"ACCOUNT_NO"`
					ACCOUNTNAME  string `xml:"ACCOUNT_NAME"`
					ACCOUNTCODE  string `xml:"ACCOUNT_CODE"`
					LISTGCONTACT struct {
						GCONTACT []struct {
							BILLNO     string `xml:"BILL_NO"`
							DORDERCOL  string `xml:"DORDER_COL"`
							ROWNUM     string `xml:"ROWNUM"`
							CONTACT    string `xml:"CONTACT"`
							POSTDATE   string `xml:"POST_DATE"`
							INVOICE    string `xml:"INVOICE"`
							AGE1       string `xml:"AGE1"`
							AGE2       string `xml:"AGE2"`
							AGE3       string `xml:"AGE3"`
							AGE4       string `xml:"AGE4"`
							AGE5       string `xml:"AGE5"`
							AGE6       string `xml:"AGE6"`
							AROPEN     string `xml:"AROPEN"`
							GUESTNAME  string `xml:"GUEST_NAME"`
							ARLEDDEBIT string `xml:"AR_LED_DEBIT"`
							TRXCODE    string `xml:"TRX_CODE"`
							NAMEID     string `xml:"NAME_ID"`
						} `xml:"G_CONTACT"`
					} `xml:"LIST_G_CONTACT"`
					PRINTACCOUNTNO string `xml:"PRINT_ACCOUNT_NO"`
					CS1            string `xml:"CS_1"`
					CS2            string `xml:"CS_2"`
					CS3            string `xml:"CS_3"`
					CS4            string `xml:"CS_4"`
					CS5            string `xml:"CS_5"`
					CS6            string `xml:"CS_6"`
					CS7            string `xml:"CS_7"`
					CS8            string `xml:"CS_8"`
				} `xml:"G_ACCOUNT_NAME"`
			} `xml:"LIST_G_ACCOUNT_NAME"`
		} `xml:"G_ACCOUNT_TYPE"`
	} `xml:"LIST_G_ACCOUNT_TYPE"`
	LISTGHATYPE struct {
		GHATYPE struct {
			HATYPE            string `xml:"HATYPE"`
			LISTGHACCOUNTNAME struct {
				GHACCOUNTNAME struct {
					HORDERCOL         string `xml:"HORDER_COL"`
					HACCNO            string `xml:"HACCNO"`
					HANAME            string `xml:"HANAME"`
					HACODE            string `xml:"HACODE"`
					LISTGACCOUNTCODE1 struct {
						GACCOUNTCODE1 []struct {
							BILLNO1    string `xml:"BILL_NO1"`
							HROWNUM    string `xml:"HROWNUM"`
							HCONTACT   string `xml:"HCONTACT"`
							HPOSTDATE  string `xml:"HPOSTDATE"`
							HINVOICE   string `xml:"HINVOICE"`
							HGUEST     string `xml:"HGUEST"`
							HA1        string `xml:"HA1"`
							HA2        string `xml:"HA2"`
							HA3        string `xml:"HA3"`
							HA4        string `xml:"HA4"`
							HA5        string `xml:"HA5"`
							HA6        string `xml:"HA6"`
							HAOPEN     string `xml:"HAOPEN"`
							HINVAGE    string `xml:"HINVAGE"`
							HINVSTATUS string `xml:"HINVSTATUS"`
							TRXNO      string `xml:"TRX_NO"`
						} `xml:"G_ACCOUNT_CODE1"`
					} `xml:"LIST_G_ACCOUNT_CODE1"`
					CS16 string `xml:"CS_16"`
					CS17 string `xml:"CS_17"`
					CS18 string `xml:"CS_18"`
					CS19 string `xml:"CS_19"`
					CS20 string `xml:"CS_20"`
					CS21 string `xml:"CS_21"`
					CS22 string `xml:"CS_22"`
				} `xml:"G_HACCOUNT_NAME"`
			} `xml:"LIST_G_HACCOUNT_NAME"`
		} `xml:"G_HATYPE"`
	} `xml:"LIST_G_HATYPE"`
	AGE1 string `xml:"AGE_1"`
	AGE2 string `xml:"AGE_2"`
	AGE3 string `xml:"AGE_3"`
	AGE4 string `xml:"AGE_4"`
	AGE5 string `xml:"AGE_5"`
	AGE6 string `xml:"AGE_6"`
	CS9  string `xml:"CS_9"`
	CS10 string `xml:"CS_10"`
	CS11 string `xml:"CS_11"`
	CS12 string `xml:"CS_12"`
	CS13 string `xml:"CS_13"`
	CS14 string `xml:"CS_14"`
	CS15 string `xml:"CS_15"`
	CS23 string `xml:"CS_23"`
	CS24 string `xml:"CS_24"`
	CS25 string `xml:"CS_25"`
	CS26 string `xml:"CS_26"`
	CS27 string `xml:"CS_27"`
	CS28 string `xml:"CS_28"`
	CS29 string `xml:"CS_29"`
	CF2  string `xml:"CF_2"`
	CF3  string `xml:"CF_3"`
	CF4  string `xml:"CF_4"`
	CF5  string `xml:"CF_5"`
	CF6  string `xml:"CF_6"`
	CF7  string `xml:"CF_7"`
	CF8  string `xml:"CF_8"`
	CF9  string `xml:"CF_9"`
	CF10 string `xml:"CF_10"`
	LOGO string `xml:"LOGO"`
} 
