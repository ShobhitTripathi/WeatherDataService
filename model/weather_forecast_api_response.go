package model

type WeatherAPIResponse struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
	Forecast struct {
		ForecastDay []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
	Alerts struct {
		Alert []Alert `json:"alert"`
	} `json:"alerts"`
}

type Alert struct {
	Headline    string      `json:"headline"`
	Msgtype     interface{} `json:"msgtype"`
	Severity    interface{} `json:"severity"`
	Urgency     interface{} `json:"urgency"`
	Areas       interface{} `json:"areas"`
	Category    string      `json:"category"`
	Certainty   interface{} `json:"certainty"`
	Event       string      `json:"event"`
	Note        interface{} `json:"note"`
	Effective   string      `json:"effective"`
	Expires     string      `json:"expires"`
	Desc        string      `json:"desc"`
	Instruction string      `json:"instruction"`
}

type ForecastDay struct {
	Date      string `json:"date"`
	DateEpoch int    `json:"date_epoch"`
	Day       Day    `json:"day"`
	//Astro     Astro  `json:"astro"`
	Hour []Hour `json:"hour"`
}

type Day struct {
	MaxTempC          float64   `json:"maxtemp_c"`
	MaxTempF          float64   `json:"maxtemp_f"`
	MinTempC          float64   `json:"mintemp_c"`
	MinTempF          float64   `json:"mintemp_f"`
	AvgTempC          float64   `json:"avgtemp_c"`
	AvgTempF          float64   `json:"avgtemp_f"`
	MaxWindMph        float64   `json:"maxwind_mph"`
	MaxWindKph        float64   `json:"maxwind_kph"`
	TotalPrecipMm     float64   `json:"totalprecip_mm"`
	TotalPrecipIn     float64   `json:"totalprecip_in"`
	AvgVisKm          float64   `json:"avgvis_km"`
	AvgVisMiles       float64   `json:"avgvis_miles"`
	AvgHumidity       int       `json:"avghumidity"`
	DailyWillItRain   int       `json:"daily_will_it_rain"`
	DailyChanceOfRain int       `json:"daily_chance_of_rain"`
	DailyWillItSnow   int       `json:"daily_will_it_snow"`
	DailyChanceOfSnow int       `json:"daily_chance_of_snow"`
	Condition         Condition `json:"condition"`
	UV                float64   `json:"uv"`
}

type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination string `json:"moon_illumination"`
	IsSunUp          int    `json:"is_sun_up"`
	IsMoonUp         int    `json:"is_moon_up"`
}

type Hour struct {
	TimeEpoch    int       `json:"time_epoch"`
	Time         string    `json:"time"`
	TempC        float64   `json:"temp_c"`
	TempF        float64   `json:"temp_f"`
	IsDay        int       `json:"is_day"`
	Condition    Condition `json:"condition"`
	WindMph      float64   `json:"wind_mph"`
	WindKph      float64   `json:"wind_kph"`
	WindDegree   int       `json:"wind_degree"`
	WindDir      string    `json:"wind_dir"`
	PressureMb   float64   `json:"pressure_mb"`
	PressureIn   float64   `json:"pressure_in"`
	PrecipMm     float64   `json:"precip_mm"`
	PrecipIn     float64   `json:"precip_in"`
	Humidity     int       `json:"humidity"`
	Cloud        int       `json:"cloud"`
	FeelslikeC   float64   `json:"feelslike_c"`
	FeelslikeF   float64   `json:"feelslike_f"`
	WindchillC   float64   `json:"windchill_c"`
	WindchillF   float64   `json:"windchill_f"`
	HeatindexC   float64   `json:"heatindex_c"`
	HeatindexF   float64   `json:"heatindex_f"`
	DewpointC    float64   `json:"dewpoint_c"`
	DewpointF    float64   `json:"dewpoint_f"`
	WillItRain   int       `json:"will_it_rain"`
	ChanceOfRain int       `json:"chance_of_rain"`
	WillItSnow   int       `json:"will_it_snow"`
	ChanceOfSnow int       `json:"chance_of_snow"`
	VisKm        float64   `json:"vis_km"`
	VisMiles     float64   `json:"vis_miles"`
	GustMph      float64   `json:"gust_mph"`
	GustKph      float64   `json:"gust_kph"`
	UV           float64   `json:"uv"`
}
