package bot



type AllCases struct {
	Cases int `json:"cases"`
	Deaths int `json:"deaths"`
	Recovered int `json:"recovered"`
}
type Countries struct {
	Country string `json:"country"`
	Cases int `json:"cases"`
	TodayCases int `json:"todayCases"`
	Deaths int `json:"deaths"`
	TodayDeaths int `json:"todayDeaths"`
	Recovered int `json:"recovered"`
	Active int `json:"active"`
	Critical int `json:"critical"`
	CasesPerOneMillion int `json:"casesPerOneMillion"`
}
