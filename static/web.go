package static

import "embed"

//go:embed index.html assets/* favicon.ico
var Static embed.FS
