package amiibo

const (
	defaultAmiiboName          string = "Toon Link - The Wind Waker"
	defaultAmiiboPage          string = "/amiibo/detail/toon-link…30th-anniversary-series"
	defaultBoxArtURL           string = "/content/dam/noa/en_US/a…nk-WW_30thAnniv_box.png"
	defaultDetailsPath         string = "/content/noa/en_US/amiib…30th-anniversary-series"
	defaultDetailsURL          string = "/amiibo/detail/toon-link…30th-anniversary-series"
	defaultFigureURL           string = "/content/dam/noa/en_US/a…k-WW_30thAnniv_char.png"
	defaultFranchise           string = "The Legend of Zelda"
	defaultGameCode            string = "NVLEAK2A"
	defaultHexCode             string = "#ffdc81"
	defaultIsReleased          bool   = true
	defaultOverviewDescription string = "<p style='margin-top: 0;'>This amiibo figure shows Link™ in his toon-shaded incarnation from The Legend of Zelda™: The Wind Waker™ game. Link holds the Wind Waker itself, a wand that allowed him to control the winds.</p>"
	defaultPresentedBy         string = "noa:publisher/nintendo"
	defaultPrice               string = "24.99"
	defaultReleaseDateMask     string = "12/02/2016"
	defaultSeries              string = "30th anniversary"
	defaultSlug                string = "toon-link-the-wind-waker…30th-anniversary-series"
	defaultType                string = "Figure"
	defaultUnixTimestamp       int64  = 1480636800
	defaultUpc                 string = "045496893064"
)

const (
	defaultDescription  string = "null"
	defaultLastModified int64  = 1554418285473
	defaultPath         string = "/content/noa/en_US/amiibo/detail/wolf-link-amiibo"
	defaultTitle        string = "Wolf Link"
	defaultURL          string = "/amiibo/detail/wolf-link-amiibo"
)

const (
	rawAmiiboDefault string = `{
		"amiiboName": "` + defaultAmiiboName + `",
		"amiiboPage": "` + defaultAmiiboPage + `",
		"boxArtURL": "` + defaultBoxArtURL + `",
		"detailsPath": "` + defaultDetailsPath + `",
		"detailsURL": "` + defaultDetailsURL + `",
		"figureURL": "` + defaultFigureURL + `",
		"franchise": "` + defaultFranchise + `",
		"gameCode": "` + defaultGameCode + `",
		"hexCode": "` + defaultHexCode + `",
		"isReleased": true,
		"overviewDescription": "` + defaultOverviewDescription + `",
		"presentedBy": "` + defaultPresentedBy + `",
		"price": "` + defaultPrice + `",
		"releaseDateMask": "` + defaultReleaseDateMask + `",
		"series": "` + defaultSeries + `",
		"slug": "` + defaultSlug + `",
		"type": "` + defaultType + `",
		"unixTimestamp": 1480636800,
		"upc": "` + defaultUpc + `"}`

	rawItemDefault string = `{
		"description": ` + defaultDescription + `,
		"lastModified": 1554418285473,
		"path": "` + defaultPath + `",
		"title": "` + defaultTitle + `",
		"url": "` + defaultURL + `"}`
)

var (
	rawAmiiboStructDefault = &RawAmiibo{
		AmiiboName:          defaultAmiiboName,
		AmiiboPage:          defaultAmiiboPage,
		BoxArtURL:           defaultBoxArtURL,
		DetailsPath:         defaultDetailsPath,
		DetailsURL:          defaultDetailsURL,
		FigureURL:           defaultFigureURL,
		Franchise:           defaultFranchise,
		GameCode:            defaultGameCode,
		HexCode:             defaultHexCode,
		IsReleased:          defaultIsReleased,
		OverviewDescription: defaultOverviewDescription,
		PresentedBy:         defaultPresentedBy,
		Price:               defaultPrice,
		ReleaseDateMask:     defaultReleaseDateMask,
		Series:              defaultSeries,
		Slug:                defaultSlug,
		Type:                defaultType,
		UnixTimestamp:       defaultUnixTimestamp,
		UPC:                 defaultUpc}

	rawItemStructDefault = &RawItem{
		Description:  defaultDescription,
		LastModified: defaultLastModified,
		Path:         defaultPath,
		Title:        defaultTitle,
		URL:          defaultURL}
)
