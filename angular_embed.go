package main

import (
	"embed"
	"io/fs"
	"net/http"
)

var assets embed.FS

func GetFrontend() http.FileSystem {
	if subAsset, err := fs.Sub(assets, "frontend"); err == nil {
		return http.FS(subAsset)
	}

	panic("Failed to load assets")
}

var _angularHandler = http.FileServer(GetFrontend())

// AngularHandler loads angular assets
// This version loads angular from files/assets
// Adds indefinite cache to all files
var AngularHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=604800")
		_angularHandler.ServeHTTP(w, r)
	},
)
