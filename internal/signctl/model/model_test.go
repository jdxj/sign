package model

import (
	"fmt"
	"testing"
)

func TestGetTasksRsp_String(t *testing.T) {
	gtr := &GetTasksRsp{
		Count:    6,
		PageID:   1,
		PageSize: 3,
		List: []*Task{
			{
				TaskID:   1,
				Describe: "2",
				Kind:     "3",
				Spec:     "4",
				Param:    []byte("5"),
			},
			{
				TaskID:   6,
				Describe: "7",
				Kind:     "8",
				Spec:     "9",
				Param:    []byte("10"),
			},
			{
				TaskID:   11,
				Describe: "12",
				Kind:     "13",
				Spec:     "14",
				Param:    []byte("15"),
			},
		},
	}
	fmt.Printf("%s\n", gtr)
}
