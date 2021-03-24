package main

import "encoding/xml"

type Loc struct {
	XMLName  xml.Name `xml:"loc"`
	Location string   `xml:",chardata"`
}
type LastMod struct {
	XMLName xml.Name `xml:"lastmod"`
	LastMod string   `xml:",chardata"`
}
type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     Loc      `xml:"loc"`
	Lastmod LastMod  `xml:"lastmod"`
}
type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	//Url     []Url    `xml:"url"`
	UrlLoc []string `xml:"url>loc"`
}

type SitemapIndex struct {
	XMLName xml.Name `xml:"sitemapindex"`
	Loc     []Loc    `xml:"sitemap>loc"`
}

type SitemapCheck struct {
	SitemapName xml.Name `xml:",any"`
}
