package opera

import "encoding/xml"

// ManagerReport was generated 2025-04-28 10:00:43 by https://xml-to-go.github.io/ in Ukraine.
type ManagerReport struct {
	XMLName              xml.Name `xml:"MANAGER_REPORT"`
	ListMasterValueOrder struct {
		GMasterValueOrder struct {
			MasterValueOrder string `xml:"MASTER_VALUE_ORDER"`
			MasterValue      string `xml:"MASTER_VALUE"`
			Resort           string `xml:"RESORT"`
			ListGLastYear01  struct {
				GLastYear01 struct {
					LastYear01 string `xml:"LAST_YEAR_01"`
					ListGCross struct {
						GCross struct {
							ListGMasterValue struct {
								GMasterValue []struct {
									SubGrp1Order       string `xml:"SUB_GRP_1_ORDER"`
									SubGrp1            string `xml:"SUB_GRP_1"`
									Description        string `xml:"DESCRIPTION"`
									AmountFormatType   string `xml:"AMOUNT_FORMAT_TYPE"`
									PrintLineAfterYN   string `xml:"PRINT_LINE_AFTER_YN"`
									ListGHeading1Order struct {
										GHeading1Order []struct {
											Heading1Order  string `xml:"HEADING_1_ORDER"`
											Heading1       string `xml:"HEADING_1"`
											Heading2       string `xml:"HEADING_2"`
											ListGSumAmount struct {
												GSumAmount struct {
													SumAmount       string `xml:"SUM_AMOUNT"`
													FormattedAmount string `xml:"FORMATTED_AMOUNT"`
												} `xml:"G_SUM_AMOUNT"`
											} `xml:"LIST_G_SUM_AMOUNT"`
										} `xml:"G_HEADING_1_ORDER"`
									} `xml:"LIST_G_HEADING_1_ORDER"`
								} `xml:"G_MASTER_VALUE"`
							} `xml:"LIST_G_MASTER_VALUE"`
						} `xml:"G_CROSS"`
					} `xml:"LIST_G_CROSS"`
				} `xml:"G_LAST_YEAR_01"`
			} `xml:"LIST_G_LAST_YEAR_01"`
			ListGForecast            string `xml:"LIST_G_FORECAST"`
			CSHeadingCountMaster     string `xml:"CS_HEADING_COUNT_MASTER"`
			CSFSArrRoomsMaster       string `xml:"CS_FS_ARR_ROOMS_MASTER"`
			CSFSDepRoomsMaster       string `xml:"CS_FS_DEP_ROOMS_MASTER"`
			CSFSNoRoomsMaster        string `xml:"CS_FS_NO_ROOMS_MASTER"`
			CSFSGuestsMaster         string `xml:"CS_FS_GUESTS_MASTER"`
			CSFSTotalRevenueMaster   string `xml:"CS_FS_TOTAL_REVENUE_MASTER"`
			CSFSRoomRevenueMaster    string `xml:"CS_FS_ROOM_REVENUE_MASTER"`
			CSFSInventoryRoomsMaster string `xml:"CS_FS_INVENTORY_ROOMS_MASTER"`
			CFFSPERCOCCRoomsMaster   string `xml:"CF_FS_PERC_OCC_ROOMS_MASTER"`
			CFFSAVGRoomRateMaster    string `xml:"CF_FS_AVG_ROOM_RATE_MASTER"`
			CSFSOORoomsMaster        string `xml:"CS_FS_OO_ROOMS_MASTER"`
		} `xml:"G_MASTER_VALUE_ORDER"`
	} `xml:"LIST_G_MASTER_VALUE_ORDER"`
	CFLogo            string `xml:"CF_LOGO"`
	CSHeadingCountRep string `xml:"CS_HEADING_COUNT_REP"`
	CPPageNumDummy    string `xml:"CP_PAGE_NUM_DUMMY"`
}
