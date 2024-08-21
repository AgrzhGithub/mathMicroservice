package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
)

type Robot struct {
	mass     float64
	velocity float64
}

type RobotInterface interface {
	getKineticEnergy() float64
}

func (r Robot) getKineticEnergy() float64 {
	result := r.mass * math.Pow(r.velocity, 2) * 0.5

	return result
}

func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mass, err := strconv.ParseFloat(vars["mass"], 64)
	if err != nil {
		fmt.Errorf("Error converting mass to float")
	}

	velocity, err := strconv.ParseFloat(vars["velocity"], 64)
	if err != nil {
		fmt.Errorf("Error converting velocity to float")
	}

	bot := Robot{mass: mass, velocity: velocity}
	jsonResponse, jsonErr := json.Marshal(bot.getKineticEnergy())
	if jsonErr != nil {
		fmt.Errorf("Error converting json to bytes")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
