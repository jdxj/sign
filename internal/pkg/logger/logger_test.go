package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(t *testing.M) {
	Init("./tgf", "test_pod")
	os.Exit(t.Run())
}

func TestDebugf(t *testing.T) {
	Debugf("abc: %s", "haha")
	Infof("def: %s", "123")
	Errorf("ghi: %s", "456")
	err := logger.Sync()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestFilePathJoin(t *testing.T) {
	p1 := []string{"", "/", "", "/", "", "/c"}
	p2 := []string{"", "", "a", "b", " ", "/d"}
	for i := range p1 {
		fmt.Printf("p1: [%s], p2: [%s], res: %s\n",
			p1[i], p2[i], filepath.Join(p1[i], p2[i]))
	}

}
func TestDir(t *testing.T) {
	fmt.Printf("base: %s\n", filepath.Base("/abc/def/name.123"))
	fmt.Printf("dir: %s\n", filepath.Dir("abc/def/name.123"))
}
