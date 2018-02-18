package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bradfitz/latlong"
	"github.com/gorilla/mux"
)

// Health function returns request path
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

// TimeZoneShow returns timezone for a given set of points
func TimeZoneShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coordinates := vars["coordinates"]
	points := strings.Split(coordinates, ";")
	var timezones Points
	for _, p := range points {
		sLatLon := strings.Split(p, ",")
		lat, err := strconv.ParseFloat(sLatLon[0], 64)
		if err != nil {
			fmt.Println("Error in parsing float, lat:", sLatLon[0])
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Error in parsing float, lat\n"))
			return
		}

		lon, err := strconv.ParseFloat(sLatLon[1], 64)
		if err != nil {
			fmt.Println("Error in parsing float, lon:", sLatLon[1])
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Error in parsing float, lon\n"))
			return
		}
		zone := latlong.LookupZoneName(lat, lon)
		timezones = append(timezones, Point{Lat: sLatLon[0], Lon: sLatLon[1], TZ: zone})
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timezones)
}
