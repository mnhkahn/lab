package models

type CyeamDoodle struct {
	Title  string `json:"title"`
	Doodle string `json:"doodle"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string        `xml:"title"`
	Items []ChannelItem `xml:"item"`
}

type ChannelItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}
