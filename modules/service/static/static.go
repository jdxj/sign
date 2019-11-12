package static

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7 = "<h1 align=\"center\">Hello World!</h1>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
	"/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1573556215, 1573556215000000000),
		Data:     []byte(_Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1573556215, 1573556215000000000),
		Data:     nil,
	}}, "")
