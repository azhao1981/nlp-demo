package main

import (
	"embed"
	"fmt"

	"github.com/go-ego/gse"
)

//go:embed testdata/test_en2.txt
var testDict embed.FS

//go:embed testdata/test_en.txt
var testEn embed.FS

var (
	text  = "To be or not to be, that's the question!"
	test1 = "Hiworld, Helloworld!"
)

func main() {
	var seg1 gse.Segmenter
	seg1.DictSep = ","
	err := seg1.LoadDict("testdata/test_en.txt")
	if err != nil {
		fmt.Println("Load dictionary error: ", err)
	}

	s1 := seg1.Cut(text)
	fmt.Println("seg1 Cut: ", s1)
	// seg1 Cut:  [to be   or   not to be ,   that's the question!]

	var seg2 gse.Segmenter
	seg2.AlphaNum = true
	seg2.LoadDict("testdata/test_en_dict3.txt")

	s2 := seg2.Cut(test1)
	fmt.Println("seg2 Cut: ", s2)
	// seg2 Cut:  [hi world ,   hello world !]

	var seg3 gse.Segmenter
	seg3.AlphaNum = true
	seg3.DictSep = ","
	err = seg3.LoadDictEmbed("hello", "world")
	if err != nil {
		fmt.Println("loadDictEmbed error: ", err)
	}
	s3 := seg3.Cut(text + test1)
	fmt.Println("seg3 Cut: ", s3)
	// seg3 Cut:  [to be   or   not to be ,   that's the question! hi world ,   hello world !]

	// example2()
}
