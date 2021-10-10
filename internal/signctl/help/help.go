package help

import (
	"fmt"
	"sort"

	"github.com/jdxj/sign/internal/proto/crontab"
)

func AvailableDomain() string {
	help := "Available domain:"

	var keys []int
	for key := range crontab.Domain_name {
		if key == 0 {
			continue
		}
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		help = fmt.Sprintf("%s\n  %s\t%d",
			help, crontab.Domain_name[int32(key)], key)
	}
	return help
}

func AvailableKind() string {
	help := "Available kind:"

	var keys []int
	for key := range crontab.Kind_name {
		if key == 0 {
			continue
		}
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		help = fmt.Sprintf("%s\n  %s\t%d",
			help, crontab.Kind_name[int32(key)], key)
	}
	return help
}
