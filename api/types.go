package api

import "time"

 type NasaResponse struct {
	Num675 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num9 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"9"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"675"`
	Num676 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num2 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"2"`
			Num5 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"5"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"676"`
	Num677 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num2 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"2"`
			Num5 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"5"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"677"`
	Num678 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num2 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"2"`
			Num3 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"3"`
			Num5 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"5"`
			Num6 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"6"`
			Num7 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"7"`
			Num8 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"8"`
			Num9 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"9"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"678"`
	Num679 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num2 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"2"`
			Num3 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"3"`
			Num5 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"5"`
			Num6 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"6"`
			Num7 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"7"`
			Num8 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"8"`
			Num9 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"9"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"679"`
	Num680 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num5 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"5"`
			Num6 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"6"`
			Num7 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"7"`
			Num8 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"8"`
			Num9 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"9"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"680"`
	Num681 struct {
		AT struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"AT"`
		FirstUTC time.Time `json:"First_UTC"`
		HWS      struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"HWS"`
		LastUTC        time.Time `json:"Last_UTC"`
		MonthOrdinal   int       `json:"Month_ordinal"`
		NorthernSeason string    `json:"Northern_season"`
		PRE            struct {
			Av float64 `json:"av"`
			Ct int     `json:"ct"`
			Mn float64 `json:"mn"`
			Mx float64 `json:"mx"`
		} `json:"PRE"`
		Season         string `json:"Season"`
		SouthernSeason string `json:"Southern_season"`
		WD             struct {
			Num0 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"0"`
			Num1 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"1"`
			Num2 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"2"`
			Num3 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"3"`
			Num5 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"5"`
			Num6 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"6"`
			Num7 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"7"`
			Num8 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"8"`
			Num9 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"9"`
			Num10 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"10"`
			Num11 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"11"`
			Num12 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"12"`
			Num13 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"13"`
			Num14 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"14"`
			Num15 struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"15"`
			MostCommon struct {
				CompassDegrees float64 `json:"compass_degrees"`
				CompassPoint   string  `json:"compass_point"`
				CompassRight   float64 `json:"compass_right"`
				CompassUp      float64 `json:"compass_up"`
				Ct             int     `json:"ct"`
			} `json:"most_common"`
		} `json:"WD"`
	} `json:"681"`
	SolKeys        []string `json:"sol_keys"`
	ValidityChecks struct {
		Num674 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"674"`
		Num675 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"675"`
		Num676 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"676"`
		Num677 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"677"`
		Num678 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"678"`
		Num679 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"679"`
		Num680 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"680"`
		Num681 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"681"`
		Num682 struct {
			AT struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"AT"`
			HWS struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"HWS"`
			PRE struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"PRE"`
			WD struct {
				SolHoursWithData []int `json:"sol_hours_with_data"`
				Valid            bool  `json:"valid"`
			} `json:"WD"`
		} `json:"682"`
		SolHoursRequired int      `json:"sol_hours_required"`
		SolsChecked      []string `json:"sols_checked"`
	} `json:"validity_checks"`
}