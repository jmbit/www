package hugo

import (
	"embed"
)

//go:embed public/*
var PublicFS embed.FS
