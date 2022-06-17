package web

import "embed"

//go:embed templates/*.gohtml
var Content embed.FS

//go:embed static/*
var Static embed.FS
