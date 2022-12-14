package variables

import (
	"runtime"
	"sync"
)

var (
	WriteWg          = sync.WaitGroup{}
	FileUriSeparator = "/"
)

func init() {
	osType := runtime.GOOS
	if osType == "windows" {
		FileUriSeparator = "\\"
	}
}
