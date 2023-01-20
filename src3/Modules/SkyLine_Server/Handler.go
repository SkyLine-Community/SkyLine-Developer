package SkyLine_Network

import (
	"fmt"
	"net/http"
)

const (
	ServerStaticFP = "Modules/SkyLine_Server/Static/"
)

var Paths = map[string]string{
	"/":          "Home",
	"/UserBoard": "Board",
	"/Errors":    "Documentation_Of_Errors",
	"/Reference": "SkyLine_Reference",
}

func SkyLine_Network_Handler(writer http.ResponseWriter, requeststream *http.Request) {
	requestpath := requeststream.URL.Path
	var newpath string
	switch requeststream.Method {
	case "GET":
		newpath = ServerStaticFP + Paths[requestpath]
		http.ServeFile(writer, requeststream, newpath)
	case "POST":
		if X := requeststream.ParseForm(); X != nil {
			http.ServeFile(writer, requeststream, newpath)
		}
	default:
		fmt.Fprintf(writer, " Unexpected or Unsupported HTTP method for requests")
	}
}
