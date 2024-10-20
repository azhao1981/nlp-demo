package main

/*
#cgo pkg-config: icu-i18n
#include <unicode/ubrk.h>
#include <unicode/uclean.h>
#include <stdlib.h>

UBreakIterator* createBreakIterator(const char* locale) {
    UErrorCode status = U_ZERO_ERROR;
    UBreakIterator* bi = ubrk_open(UBRK_WORD, locale, NULL, 0, &status);
    if (U_FAILURE(status)) return NULL;
    return bi;
}

void setText(UBreakIterator* bi, const UChar* text, int32_t length) {
    ubrk_setText(bi, text, length, NULL);
}

int32_t firstBoundary(UBreakIterator* bi) {
    return ubrk_first(bi);
}

int32_t nextBoundary(UBreakIterator* bi) {
    return ubrk_next(bi);
}

void closeBreakIterator(UBreakIterator* bi) {
    ubrk_close(bi);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	locale := C.CString("th_TH")
	defer C.free(unsafe.Pointer(locale))

	bi := C.createBreakIterator(locale)
	if bi == nil {
		fmt.Println("Failed to create BreakIterator")
		return
	}
	defer C.closeBreakIterator(bi)

	// Example Thai text
	thaiText := "สวัสดีครับ นี่คือการทดสอบการตัดคำภาษาไทย"
	utf16 := []uint16{}
	for _, r := range thaiText {
		utf16 = append(utf16, uint16(r))
	}
	ptr := (*C.UChar)(unsafe.Pointer(&utf16[0]))
	C.setText(bi, ptr, C.int32_t(len(utf16)))

	start := C.firstBoundary(bi)
	for {
		end := C.nextBoundary(bi)
		if end == C.UBRK_DONE {
			break
		}
		if end > start {
			word := thaiText[start:end]
			fmt.Println(word)
		}
		start = end
	}
}
