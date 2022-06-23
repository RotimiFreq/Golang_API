package main

import (
	"fmt"
	"html"
	"net/http"
	"soundapidata"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("confiming")
	var test = soundapidata.Mixer{}

	test.Name = "quantum"
	test.ModelNo = 19945
	test.YearProduce = 2019
	test.Price = 8000
	test.StillProducing = true
	test.ChannelNo = 48
	test.OutputNo = 32
	test.UsbConn = true
	test.WifiConn = true
	test.DanteSupport = true
	test.About = "a wonderful system"

	var test2 = soundapidata.Mixer{}

	test2.Name = "qu-32"
	test2.ModelNo = 19945
	test2.YearProduce = 2014
	test2.Price = 2300
	test2.StillProducing = true
	test2.ChannelNo = 32
	test2.OutputNo = 20
	test2.UsbConn = true
	test2.WifiConn = true
	test2.DanteSupport = false
	test2.About = "user freiendly"

	var test3 = soundapidata.Mixer{}

	test3.Name = "X-Air"
	test3.ModelNo = 193245
	test3.YearProduce = 2015
	test3.Price = 3920
	test3.StillProducing = true
	test3.ChannelNo = 32
	test3.OutputNo = 24
	test3.UsbConn = true
	test3.WifiConn = true
	test3.DanteSupport = false
	test3.About = "very good for live production"

	var Mixerslice []soundapidata.Mixer

	Mixerslice = append(Mixerslice, test)
	Mixerslice = append(Mixerslice, test2)
	Mixerslice = append(Mixerslice, test3)

	// fmt.Print(len(Mixerslice))
	// fmt.Print(Mixerslice[]//)

	//http package has methods for dealing with requests

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// path of the http request

		urlPathElements := strings.Split(r.URL.Path, "/")

		// If request is GET with correct syntax

		if urlPathElements[1] == "mixer" {

			// // converting a string integer to integer using Atoi function
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))

			x := len(Mixerslice)

			if number == 0 || number > x {

				// If resource is not in the list, send Not Found status

				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				dataRet := Mixerslice[number]
				fmt.Fprintf(w, "%q", html.EscapeString(dataRet)
			}

		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})
	// Create a server and run it on 8000 port
	s := &http.Server{
		Addr: ":8000",
		// ReadTimeout: 10 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
