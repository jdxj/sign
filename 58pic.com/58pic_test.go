package pic

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	Start()
}

func TestBAE(t *testing.T) {
	start, end := beginAndEnd()
	fmt.Println(start)
	fmt.Println(end)
}

func TestTimeUnix(t2 *testing.T) {
	t := unixTimeMill()
	fmt.Println(t)
	fmt.Println(postUrl())
}
