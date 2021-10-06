package evo

import (
	"fmt"
	"time"
)

const (
	buildURL = "https://raw.githubusercontent.com/Evolution-X-Devices/official_devices/master/builds/%s.json"
)

const (
	gb = 1 << 30
)

type buildInfo struct {
	Filename string `json:"filename"`
	Datetime int64  `json:"datetime"`
	Size     int64  `json:"size"`
	URL      string `json:"url"`
}

func (bi *buildInfo) String() string {
	format := `Filename: %s
UpdateTime: %s
Size: %.3fGB
DownloadURL: %s`
	return fmt.Sprintf(format, bi.Filename, time.Unix(bi.Datetime, 0).Format(time.RFC3339),
		float64(bi.Size)/gb, bi.URL)
}
