// +build !windows

package diskinfo

import (
	"github.com/getsentry/sentry-go"
	"golang.org/x/sys/unix"
)

// calculates disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := unix.Statfs_t{}
	err := unix.Statfs(path, &fs)
	if err != nil {
		sentry.CaptureException(err)
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
