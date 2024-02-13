package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\n<body>\n  <p>Can you see this? → {{.Bar}}</p>\n</body>"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\n<body>\n  <p>Hello, {{.Foo}}</p>\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"bar.tmpl", "index.tmpl"}}, map[string]*assets.File{
	"/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1706790617, 1706790617441639445),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1706790756, 1706790756535987981),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1706790627, 1706790627073805850),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1706790633, 1706790633453915731),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}}, "")
