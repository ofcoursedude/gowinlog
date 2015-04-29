package winlog

/*
#cgo CPPFLAGS: -I C:/mingw-w64/x86_64-4.9.2-posix-seh-rt_v4-rev2/mingw64/x86_64-w64-mingw32/include
#cgo CFLAGS: -I C:/mingw-w64/x86_64-4.9.2-posix-seh-rt_v4-rev2/mingw64/x86_64-w64-mingw32/include
#cgo LDFLAGS: -l wevtapi -L C:/mingw-w64/x86_64-4.9.2-posix-seh-rt_v4-rev2/mingw64/x86_64-w64-mingw32/lib
#include "bookmark.h"
*/
import "C"
import (
	"unsafe"
)

type BookmarkHandle uint64

func CreateBookmark() (BookmarkHandle, error) {
	bookmark := BookmarkHandle(C.CreateBookmark())
	if bookmark == 0 {
		return 0, GetLastError()
	}
	return bookmark, nil
}

func CreateBookmarkFromXml(xmlString string) (BookmarkHandle, error) {
	cString := C.CString(xmlString)
	bookmark := C.CreateBookmarkFromXML(cString)
	C.free(unsafe.Pointer(cString))
	if bookmark == 0 {
		return 0, GetLastError()
	}
	return BookmarkHandle(bookmark), nil
}

func UpdateBookmark(bookmarkHandle BookmarkHandle, eventHandle EventHandle) error {
	if C.UpdateBookmark(C.ULONGLONG(bookmarkHandle), C.ULONGLONG(eventHandle)) == 0 {
		return GetLastError()
	}
	return nil
}

func RenderBookmark(bookmarkHandle BookmarkHandle) (string, error) {
	cString := C.RenderBookmark(C.ULONGLONG(bookmarkHandle))
	if cString == nil {
		return "", GetLastError()
	}
	bookmarkXml := C.GoString(cString)
	C.free(unsafe.Pointer(cString))
	return bookmarkXml, nil
}
