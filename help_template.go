package main

const HelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
   
USAGE:
   {{.HelpName}} [options] servername
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
	 {{end}}
`
