package pic

import (
	"fmt"
	"github.com/gocolly/redisstorage"
	"net/url"
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

func TestRedisStorageBackend(t *testing.T) {
	storage := redisstorage.Storage{
		Address:  "redis.aaronkir.xyz:6379",
		Password: "",
		DB:       0,
		Prefix:   "58pic.com",
	}
	err := storage.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.Clear()
	if err != nil {
		fmt.Println(err)
		return
	}

	cookies := storage.Cookies(&url.URL{})
	fmt.Println("cookies:", cookies)
}

func TestParseUrl(t *testing.T) {
	cookieUrl, err := url.Parse("https://www.58pic.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cookieUrl)
}
