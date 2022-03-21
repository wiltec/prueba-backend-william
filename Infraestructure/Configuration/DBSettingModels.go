package Configuration

import "encoding/xml"

type DBSettingsModel struct {
	XMLName  xml.Name  `xml:"DBSettings"`
	DataBase []DBModel `xml:"DB"`
}

type DBModel struct {
	XMLName  xml.Name `xml:"DB"`
	Name     string   `xml:"name,attr"`
	Engine   string   `xml:"Engine"`
	Server   string   `xml:"Server"`
	Port     string   `xml:"Port"`
	User     string   `xml:"User"`
	Password string   `xml:"Password"`
	Database string   `xml:"Database"`
	SslMode  string   `xml:"SslMode"`
}
