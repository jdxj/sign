package help

import (
	"fmt"
	"sort"
)

func AvailableDomain() string {
	help := "Available domain:"

	keys := make([]int, 0, len(crontab.Domain_name))
	for key := range crontab.Domain_name {
		if key == 0 {
			continue
		}
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		help = fmt.Sprintf("%s\n  %-15s %d",
			help, crontab.Domain_name[int32(key)], key)
	}
	return help
}

func AvailableKind() string {
	help := "Available kind:"

	keys := make([]int, 0, len(crontab.Kind_name))
	for key := range crontab.Kind_name {
		if key == 0 {
			continue
		}
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		help = fmt.Sprintf("%s\n  %-15s %d",
			help, crontab.Kind_name[int32(key)], key)
	}
	return help
}
