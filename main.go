package main

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/latlong"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	Lat string `json: "lat"`
	Lon string `json: "lon"`
	TZ  string `json: "tz"`
}

type Points []Point

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/timezone/{coordinates}", TimeZoneShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index function returns request path
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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
			fmt.Printf("Error in parsing float, lat: %s", sLatLon[0])
		}

		lon, err := strconv.ParseFloat(sLatLon[1], 64)
		if err != nil {
			fmt.Printf("Error in parsing float, lon: %s", sLatLon[1])
		}
		zone := latlong.LookupZoneName(lat, lon)
		timezones = append(timezones, Point{Lat: sLatLon[0], Lon: sLatLon[1], TZ: zone})
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timezones)
}
