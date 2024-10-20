package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-ego/gse"
	"github.com/stretchr/testify/assert"
)

func TestGse(t *testing.T) {
	var seg1 gse.Segmenter
	seg1.DictSep = ","
	err := seg1.LoadDict()
	if err != nil {
		fmt.Println("Load dictionary error: ", err)
	}
	text := "你好世界, Hello world, Helloworld."
	s1 := seg1.Cut(text)
	fmt.Println("seg1 Cut dag: ", s1)

	s2 := seg1.Cut(text, true)
	fmt.Println("seg2 Cut hmm: ", s2)

	hmm := seg1.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	fmt.Println("analyze: ", seg1.Analyze(hmm, text))
}

func TestCut(t *testing.T) {
	var seg1 gse.Segmenter
	seg1.DictSep = ","
	err := seg1.LoadDict()
	if err != nil {
		fmt.Println("Load dictionary error: ", err)
	}
	text := "你好世界, Hello world, Helloworld."
	s1 := seg1.Cut(text)
	fmt.Printf("seg1 Cut dag: %+v, len: %d\n", s1, len(s1))
	s2 := seg1.CutSearch(text)
	fmt.Printf("seg1 Cut search: %+v, len: %d\n", s2, len(s2))
	s3 := seg1.CutAll(text)
	fmt.Printf("seg1 Cut all: %+v, len: %d\n", s3, len(s3))
}

func TestCut2(t *testing.T) {
	var seg1 gse.Segmenter
	seg1.DictSep = ","
	err := seg1.LoadDict()
	if err != nil {
		fmt.Println("Load dictionary error: ", err)
	}
	text := "你好世界, Hello world, Helloworld."
	s1 := seg1.Cut(text, true)
	fmt.Printf("seg1 Cut dag: %+v, len: %d\n", s1, len(s1))
	s2 := seg1.CutSearch(text, true)
	fmt.Printf("seg1 Cut search: %+v, len: %d\n", s2, len(s2))
	s3 := seg1.CutAll(text)
	fmt.Printf("seg1 Cut all: %+v, len: %d\n", s3, len(s3))
}

func TestCutThai(t *testing.T) {
	time.Sleep(11 * time.Microsecond)
	var seg1 gse.Segmenter
	seg1.DictSep = ","
	err := seg1.LoadDict("/home/weiz/projects/demo/nlp-demo/gse-demo/tests/words_th_modified.txt")
	if err != nil {
		fmt.Println("Load dictionary error: ", err)
	}
	// https://www.thairath.co.th/news/local/2819456
	// สำหรับหลายคนแล้ว => สำหรับ หลาย คน แล้ว
	// text := "สำหรับหลายคนแล้ว เมืองฮ่องกง เป็นหนึ่งในสถานที่ท่องเที่ยวที่น่าประทับใจ ทั้งในเรื่องของอาหารการกินที่หลากหลายและแปลกใหม่"
	// text = text + "ศิลปะสมัยใหม่จากศิลปินระดับโลก แหล่งช้อปปิ้งเสื้อผ้าแฟชั่นและเทคโนโลยีที่ทันสมัย รวมไปถึงที่พึ่งทางใจและส่งเสริมโชคลาภ"
	// text = text + "จึงไม่แปลกที่หลายคนยกให้ฮ่องกงเป็นเมืองที่ถ้ามีโอกาสก็จะกลับมาเยี่ยมเยียนทุกครั้ง"
	// 你好,明天见,去年的有，上个月走了，下午5点钟，北京市海淀区, 13888886666, 一只三个，都是王先生的，还要等上半天
	text := "สวัสดี พบกันพรุ่งนี้ มีของปีที่แล้ว ไปเมื่อเดือนที่แล้ว บ่ายห้าโมง เขตไห่เตี้ยน กรุงปักกิ่ง 13888886666 หนึ่งตัวสามอัน ทั้งหมดเป็นของคุณหวัง ต้องรออีกครึ่งวัน"
	// text = strings.ReplaceAll(text, " ", "")
	s1 := seg1.Cut(text, true)
	fmt.Printf("seg1 Cut dag: %+v, len: %d\n", s1, len(s1))
	for _, s := range s1 {
		fmt.Printf("%s|", s)
	}
	fmt.Println()
	s2 := seg1.CutSearch(text, true)
	fmt.Printf("seg1 Cut search: %+v, len: %d\n", s2, len(s2))
	s3 := seg1.CutAll(text)
	fmt.Printf("seg1 Cut all: %+v, len: %d\n", s3, len(s3))
}
func bytesToString(bytes []gse.Text) (output string) {
	for _, b := range bytes {
		output += (string(b) + "/")
	}
	return
}
func TestCutThai2(t *testing.T) {
	time.Sleep(12 * time.Microsecond)
	var seg gse.Segmenter
	// seg.DictSep = ","
	seg, err := gse.NewEmbed("สำหรับ 2000000 n")
	assert.NoError(t, err)

	// seg.LoadDictEmbed()
	seg.LoadStopEmbed()
	// https://www.thairath.co.th/news/local/2819456
	// สำหรับหลายคนแล้ว => สำหรับ หลาย คน แล้ว
	text := "สำหรับหลายคนแล้ว เมืองฮ่องกง เป็นหนึ่งในสถานที่ท่องเที่ยวที่น่าประทับใจ ทั้งในเรื่องของอาหารการกินที่หลากหลายและแปลกใหม่"
	text = text + "ศิลปะสมัยใหม่จากศิลปินระดับโลก แหล่งช้อปปิ้งเสื้อผ้าแฟชั่นและเทคโนโลยีที่ทันสมัย รวมไปถึงที่พึ่งทางใจและส่งเสริมโชคลาภ"
	text = text + "จึงไม่แปลกที่หลายคนยกให้ฮ่องกงเป็นเมืองที่ถ้ามีโอกาสก็จะกลับมาเยี่ยมเยียนทุกครั้ง"
	// text := "สวัสดี พบกันพรุ่งนี้ มีของปีที่แล้ว ไปเมื่อเดือนที่แล้ว บ่ายห้าโมง เขตไห่เตี้ยน กรุงปักกิ่ง 13888886666 หนึ่งตัวสามอัน ทั้งหมดเป็นของคุณหวัง ต้องรออีกครึ่งวัน"
	// text = strings.ReplaceAll(text, " ", "")
	// s1 := seg.Cut(text, true)
	s1 := seg.SplitTextToWords([]byte(text))
	fmt.Printf("seg1 Cut dag: %+v, len: %d\n", bytesToString(s1), len(s1))
	s11 := seg.Cut(text, true)
	fmt.Printf("seg1 Cut hmm: %+v, len: %d\n", s11, len(s11))
	s2 := seg.CutSearch(text, true)
	fmt.Printf("seg1 Cut search: %+v, len: %d\n", s2, len(s2))
	s3 := seg.CutAll(text)
	fmt.Printf("seg1 Cut all: %+v, len: %d\n", s3, len(s3))
}

func TestCutThai3(t *testing.T) {
	time.Sleep(11 * time.Microsecond)
	var seg1 gse.Segmenter
	seg1.DictSep = ","
	err := seg1.LoadDict("/home/weiz/projects/demo/nlp-demo/gse-demo/tests/words_th_modified.txt")
	if err != nil {
		fmt.Println("Load dictionary error: ", err)
	}
	// 你好,明天见,去年的有，上个月走了，下午5点钟，北京市海淀区, 13888886666, 一只三个，都是王先生的，还要等上半天
	text := "สวัสดี พบกันพรุ่งนี้ มีของปีที่แล้ว ไปเมื่อเดือนที่แล้ว บ่ายห้าโมง เขตไห่เตี้ยน กรุงปักกิ่ง 13888886666 หนึ่งตัวสามอัน ทั้งหมดเป็นของคุณหวัง ต้องรออีกครึ่งวัน"
	s1 := seg1.Cut(text, true)
	fmt.Printf("seg1 Cut dag: %+v, len: %d\n", s1, len(s1))
	for _, s := range s1 {
		fmt.Printf("%s|", s)
	}
	fmt.Println()
	seg1.LoadDictStr("สามอัน, 30, n")
	s2 := seg1.Cut(text, true)
	fmt.Printf("seg1 Cut dag: %+v, len: %d\n", s2, len(s2))
	for _, s := range s2 {
		fmt.Printf("%s|", s)
	}
	fmt.Println()
}
