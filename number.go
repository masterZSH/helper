package helper

import "fmt"

// ByteSize byte size B,KB...
type ByteSize int

var ByteSizeStr []string = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const (
	// B 1B => 8bit
	B ByteSize = iota

	// KB 1KB => 1024B
	KB

	// MB 1MB => 1024KB
	MB

	// GB 1GB => 1024MB
	GB

	// TB 1TB => 1024GB
	TB

	// PB 1PB => 1024TB
	PB

	// EB 1EB => 1024PB
	EB
)

// 1<<10  1024
const unit = 1 << 10

// ByteCount converts bytes number to human readable byte string
// 1024 => 1KB
// 1100000 => 1.1MB
func ByteCount(b int, size ByteSize) string {
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div := unit
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		size++
	}
	size++
	return fmt.Sprintf("%.1f %s", float64(b)/float64(div), ByteSizeStr[size])
}
