package main

// Point struct holds point and timezone information
type Point struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	TZ  string `json:"tz"`
}

// Points array holds an array of points to return as json
type Points []Point
