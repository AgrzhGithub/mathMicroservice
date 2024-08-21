package pkg

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Point struct {
	X float64
	Y float64
}

type Coordinates2d interface {
	Distance() float64
}

func Distance(A, B Point) float64 {
	result := math.Pow(B.X-A.X, 2) + math.Pow(B.Y-A.Y, 2)
	result = math.Sqrt(result)
	return result
}

func DotsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	x1, err := strconv.ParseFloat(vars["x1"], 64)
	if err != nil {
		panic(err)
	}
	x2, err := strconv.ParseFloat(vars["x2"], 64)
	if err != nil {
		panic(err)
	}
	y1, err := strconv.ParseFloat(vars["y1"], 64)
	if err != nil {
		panic(err)
	}
	y2, err := strconv.ParseFloat(vars["y2"], 64)
	if err != nil {
		panic(err)
	}

	point1 := Point{X: x1, Y: y1}
	point2 := Point{X: x2, Y: y2}

	distance := Distance(point1, point2)

	jsonResponse, _ := json.Marshal(distance)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("User-Role")
		if val == "Admin228" {
			log.Println("Accepted")
		}
		next(w, r)
	}
}
