package types

//나이스에서 가져오는 급식 정보 구조체입니다.
type MealServiceDietInfoType struct {
	MealServiceDietInfo []struct {
		Head []struct {
			Result struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT"`
		} `json:"head,omitempty"`
		Row []struct {
			AtptOfcdcScCode string `json:"ATPT_OFCDC_SC_CODE"`
			AtptOfcdcScNm   string `json:"ATPT_OFCDC_SC_NM"`
			SdSchulCode     string `json:"SD_SCHUL_CODE"`
			SchulNm         string `json:"SCHUL_NM"`
			MmealScCode     string `json:"MMEAL_SC_CODE"`
			MmealScNm       string `json:"MMEAL_SC_NM"`
			MlsvYmd         string `json:"MLSV_YMD"`
			MlsvFgr         string `json:"MLSV_FGR"`
			DdishNm         string `json:"DDISH_NM"`
			OrplcInfo       string `json:"ORPLC_INFO"`
			CalInfo         string `json:"CAL_INFO"`
			NtrInfo         string `json:"NTR_INFO"`
			MlsvFromYmd     string `json:"MLSV_FROM_YMD"`
			MlsvToYmd       string `json:"MLSV_TO_YMD"`
		} `json:"row,omitempty"`
	} `json:"mealServiceDietInfo"`
}

//나이스에서 가져오는 시간표 정보 구조체입니다.
type HisTimetableType struct {
	HisTimetable []struct {
		Head []struct {
			ListTotalCount int `json:"list_total_count,omitempty"`
			Result         struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT,omitempty"`
		} `json:"head,omitempty"`
		Row []struct {
			AtptOfcdcScCode string `json:"ATPT_OFCDC_SC_CODE"`
			AtptOfcdcScNm   string `json:"ATPT_OFCDC_SC_NM"`
			SdSchulCode     string `json:"SD_SCHUL_CODE"`
			SchulNm         string `json:"SCHUL_NM"`
			Ay              string `json:"AY"`
			Sem             string `json:"SEM"`
			AllTiYmd        string `json:"ALL_TI_YMD"`
			DghtCrseScNm    string `json:"DGHT_CRSE_SC_NM"`
			OrdScNm         string `json:"ORD_SC_NM"`
			DddepNm         string `json:"DDDEP_NM"`
			Grade           string `json:"GRADE"`
			ClrmNm          string `json:"CLRM_NM"`
			ClassNm         string `json:"CLASS_NM"`
			Perio           string `json:"PERIO"`
			ItrtCntnt       string `json:"ITRT_CNTNT"`
			LoadDtm         string `json:"LOAD_DTM"`
		} `json:"row,omitempty"`
	} `json:"hisTimetable"`
}