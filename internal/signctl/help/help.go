package help

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/jdxj/sign/internal/proto/task"
)

func AvailableKind() string {
	help := "Available kind and params:\n"
	help = fmt.Sprintf("%s  %-18s\t%s\n", help, "Kind Name", "Kind ID")
	help = fmt.Sprintf("%s    Task Param\n", help)
	help = fmt.Sprintf("%s  %-18s\t%s\n", help, "---------", "-------")

	kinds := make([]string, 0, len(task.Kind_value))
	for k := range task.Kind_value {
		if k == task.Kind_UNKNOWN_KIND.String() ||
			k == task.Kind_MOCK.String() ||
			k == task.Kind_HPI_SIGN_IN.String() {
			continue
		}
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)

	for _, k := range kinds {
		help = fmt.Sprintf("%s  %-18s\t%7d\n", help, k, task.Kind_value[k])
		paramList := getParamList(k)
		if paramList != "" {
			help = fmt.Sprintf("%s   %s\n", help, getParamList(k))
		}
	}
	return help
}

func getParamList(kind string) string {
	msg := task.GetParamByKind(kind)
	if msg == nil {
		return ""
	}

	var paramList string
	msgRt := reflect.TypeOf(msg).Elem()
	for i := 0; i < msgRt.NumField(); i++ {
		key := msgRt.Field(i).Tag.Get("json")
		if key == "" {
			continue
		}
		key = strings.Split(key, ",")[0]
		keyKind := msgRt.Field(i).Type.Kind()
		paramList = fmt.Sprintf("%s %s(%s)", paramList, key, keyKind)
	}
	return paramList
}
