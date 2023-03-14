package model

type HostMeta struct {
	XMLName string  `xml:"XRD" json:"-"`
	Link    XMLLink `xml:"Link" json:"link"`
	Xmlns   string  `xml:"xmlns,attr" json:"-"`
}

type XMLLink struct {
	Rel      string `xml:"rel,attr" json:"rel"`
	Type     string `xml:"type,attr" json:"type"`
	Template string `xml:"template,attr" json:"template"`
}
