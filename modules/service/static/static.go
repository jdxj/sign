package static

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7 = "<h1>Hello World!</h1>"
var _Assetseeecc25dd8b871f37dd7a52a9def696ef0109864 = "<h1>hello world</h1>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html", "index2.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1573117136, 1573117136000000000),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1573117001, 1573117001000000000),
		Data:     []byte(_Assets14fe4559026d4c5b5eb530ee70300c52d99e70d7),
	}, "/index2.html": &assets.File{
		Path:     "/index2.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1573117136, 1573117136000000000),
		Data:     []byte(_Assetseeecc25dd8b871f37dd7a52a9def696ef0109864),
	}}, "")
