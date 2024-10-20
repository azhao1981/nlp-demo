package main

/*
#cgo pkg-config: icu-uc
#include <unicode/ucnv.h>
*/
import "C"
import (
	"fmt"

	"github.com/goodsign/icu"
)

func main() {
	// 定义错误代码变量
	var err C.UErrorCode = C.U_ZERO_ERROR

	// 获取默认的转换器名称
	cname := C.ucnv_getDefaultName()
	name := C.GoString(cname)
	fmt.Println("默认转换器名称：", name)

	// 打开转换器
	converter := C.ucnv_open(nil, &err)
	if err != C.U_ZERO_ERROR {
		fmt.Println("打开转换器时出错：", int(err))
		return
	}
	defer C.ucnv_close(converter)

	// 在这里可以使用 converter 进行进一步操作
	enc, err2 := Detect("สวัสดี พบกันพรุ่งนี้ มีของปีที่แล้ว ไปเมื่อเดือนที่แล้ว บ่ายห้าโมง เขตไห่เตี้ยน กรุงปักกิ่ง 13888886666")
	if err2 != nil {
		fmt.Println("检测编码时出错：", err2)
		return
	}
	fmt.Println("编码：", enc)
}

func Detect(text string) (string, error) {
	detector, err := icu.NewCharsetDetector()
	if err != nil {
		return "", err
	}
	defer detector.Close()
	encMatches, err := detector.GuessCharset([]byte(text))
	if err != nil {
		return "", err
	}
	maxenc := encMatches[0].Charset

	return maxenc, nil
}
