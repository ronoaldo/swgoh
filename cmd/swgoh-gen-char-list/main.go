package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	o = flag.String("o", "-", "Output file; - for stdout")
)

func main() {
	flag.Parse()

	var out io.Writer
	out = os.Stdout
	if *o != "-" {
		fd, err := os.Create(*o)
		if err != nil {
			panic(err)
		}
		defer fd.Close()
		out = fd
	}
	genCharList(out)
}

func genCharList(out io.Writer) {
	doc, err := goquery.NewDocument("https://swgoh.gg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(out, `package swgohgg

import "strings"

func CharSlug(charName string) string {
	switch strings.ToLower(charName) {
`)
	doc.Find("a.character").Each(func(i int, el *goquery.Selection) {
		name := strings.ToLower(el.Find("h5").First().Text())
		href := el.AttrOr("href", "?")
		slug := strings.Split(href, "/")[2]
		fmt.Fprintf(out, "\tcase `%s`:\n\t\treturn \"%s\"\n", name, slug)
	})
	fmt.Fprintf(out, "\tdefault:\n\t\treturn strings.ToLower(charName)\n\t}\n}\n")
	if *o != "-" {
		log.Printf("File %s generated.", *o)
	}
}
