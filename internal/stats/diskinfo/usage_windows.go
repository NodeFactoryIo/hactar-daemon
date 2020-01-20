package diskinfo

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

// calculates disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	ds := &DiskStatus{}

	av := new(int)

	_, _, _ = c.Call(
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&ds.Free)),
		uintptr(unsafe.Pointer(&ds.All)),
		uintptr(unsafe.Pointer(&av)))

	ds.Used = ds.All - ds.Free

	return *ds
}
