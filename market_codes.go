package opera

import "encoding/xml"

// MarketCodes was generated 2025-04-28 15:47:22 by https://xml-to-go.github.io/ in Ukraine.
type MarketCodes struct {
	XMLName   xml.Name `xml:"STAT_DMY_SEG"`
	ListGGrp1 struct {
		GGrp1 []struct {
			GrpSortOrder           string `xml:"GRP_SORT_ORDER"`
			Grp1CodeDesc           string `xml:"GRP1_CODE_DESC"`
			ParentMarketCode       string `xml:"PARENT_MARKET_CODE"`
			MarketGroupDescription string `xml:"MARKET_GROUP_DESCRIPTION"`
			ListGGrp2              struct {
				GGrp2 []struct {
					Grp2Code       string `xml:"GRP2_CODE"`
					SubGrpCodeDesc string `xml:"SUB_GRP_CODE_DESC"`
					Description    string `xml:"DESCRIPTION"`
					GraphXCode     string `xml:"GRAPH_X_CODE"`
					RoomsDay       string `xml:"ROOMS_DAY"`
					RoomRevDay     string `xml:"ROOM_REV_DAY"`
					RoomsMtd       string `xml:"ROOMS_MTD"`
					GuestMtd       string `xml:"GUEST_MTD"`
					RoomRevMtd     string `xml:"ROOM_REV_MTD"`
					RoomsYtd       string `xml:"ROOMS_YTD"`
					GuestYtd       string `xml:"GUEST_YTD"`
					RoomRevYtd     string `xml:"ROOM_REV_YTD"`
					GuestDay       string `xml:"GUEST_DAY"`
					AdrDay         string `xml:"ADR_DAY"`
					AdrMtd         string `xml:"ADR_MTD"`
					AdrYtd         string `xml:"ADR_YTD"`
					PerOccDay      string `xml:"PER_OCC_DAY"`
					PerOccMtd      string `xml:"PER_OCC_MTD"`
					PerOccYtd      string `xml:"PER_OCC_YTD"`
				} `xml:"G_GRP2"`
			} `xml:"LIST_G_GRP2"`
			PmDayRoom     string `xml:"PM_DAY_ROOM"`
			PmDayPerson   string `xml:"PM_DAY_PERSON"`
			PmMonthRoom   string `xml:"PM_MONTH_ROOM"`
			PmMonthPerson string `xml:"PM_MONTH_PERSON"`
			PmYearRoom    string `xml:"PM_YEAR_ROOM"`
			PmYearPrs     string `xml:"PM_YEAR_PRS"`
			PmDayCrev     string `xml:"PM_DAY_CREV"`
			PmMonthCrev   string `xml:"PM_MONTH_CREV"`
			PmYearCrev    string `xml:"PM_YEAR_CREV"`
			PmYearArr     string `xml:"PM_YEAR_ARR"`
			PmMonthArr    string `xml:"PM_MONTH_ARR"`
			PmDayArr      string `xml:"PM_DAY_ARR"`
			PmDayPerOcc   string `xml:"PM_DAY_PER_OCC"`
			PmMtdPerOcc   string `xml:"PM_MTD_PER_OCC"`
			PmYtdPerOcc   string `xml:"PM_YTD_PER_OCC"`
		} `xml:"G_GRP1"`
	} `xml:"LIST_G_GRP1"`
	SDayRooms       string `xml:"S_DAY_ROOMS"`
	SDayPersons     string `xml:"S_DAY_PERSONS"`
	SMonthRms       string `xml:"S_MONTH_RMS"`
	SMonthPrs       string `xml:"S_MONTH_PRS"`
	SYearRms        string `xml:"S_YEAR_RMS"`
	SYearPrs        string `xml:"S_YEAR_PRS"`
	SDayArr         string `xml:"S_DAY_ARR"`
	SMonthArr       string `xml:"S_MONTH_ARR"`
	SYearArr        string `xml:"S_YEAR_ARR"`
	Logo            string `xml:"LOGO"`
	SDayCrev        string `xml:"S_DAY_CREV"`
	SMonthCrev      string `xml:"S_MONTH_CREV"`
	SYearCrev       string `xml:"S_YEAR_CREV"`
	SYtdPerOcc      string `xml:"S_YTD_PER_OCC"`
	SMtdPerOcc      string `xml:"S_MTD_PER_OCC"`
	SDayPerOcc      string `xml:"S_DAY_PER_OCC"`
	CfPeriodText    string `xml:"CF_PERIOD_TEXT"`
	CfLegRoomsMtd   string `xml:"CF_LEG_ROOMS_MTD"`
	CfLegRevenueMtd string `xml:"CF_LEG_REVENUE_MTD"`
}
