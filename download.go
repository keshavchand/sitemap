package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func getXML(site string) ([]byte, error) {
	resp, err := http.Get(site)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getSitesFromSitemap(site string) []string {
	var sites []string
	resp, err := getXML(site)
	if err != nil {
		return nil
	}
	var urlset UrlSet
	xml.Unmarshal(resp, &urlset)
	for _, i := range urlset.UrlLoc {
		sites = append(sites, i)
	}
	return sites
}

func getSites(site string) []string {
	var siteCheck SitemapCheck
	var sites []string
	result, err := getXML(site)
	if err == nil {
		xml.Unmarshal(result, &siteCheck)
	}
	if siteCheck.SitemapName.Local == "sitemap" {
		var sitemap SitemapIndex
		xml.Unmarshal(result, &sitemap)
		for _, i := range sitemap.Loc {
			sites = append(sites, getSitesFromSitemap(i.Location)...)
		}
	} else if siteCheck.SitemapName.Local == "url" {
		sites = append(sites, getSitesFromSitemap(site)...)
	}

	return sites
}
