package main

import (
	"log"
	"os"
)

func main() {
	sitemap_location := "https://www.google.com/sitemap.xml"
	if len(os.Args) >= 2 {
		sitemap_location = os.Args[1]
	}

	//fmt.Println(getSites(sitemap_location))
	for _, i := range getSites(sitemap_location) {
		log.Println(i)
	}
}
