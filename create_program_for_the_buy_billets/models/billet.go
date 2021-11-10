package models

import "math/rand"

var DistanceToMars uint32 = 62100000

var speeds [3]uint32 = [3]uint32{
	16000,
	23000,
	30000,
}

type SpaceLine string
type TripType string

type Rocket struct {
	name  string
	speed uint32
}

func (rocket *Rocket) newRocket(name string) (*Rocket, error) {
	rocket.name = name
	rnd := rand.Intn(3)
	rocket.speed = speeds[rnd]
	return rocket, nil
}

type Billet struct {
	spaceLineName SpaceLine
	tripTypeName  TripType
	rocket        Rocket
	days          uint32
	price         uint32
}

func NewBillet(spaceLine SpaceLine, tripType TripType, rocket Rocket) (Billet, error) {
	billet := Billet{}
	billet.spaceLineName = spaceLine
	billet.tripTypeName = tripType
	billet.rocket = rocket
	var speed uint32 = billet.rocket.speed
	billet.days = DistanceToMars / speed
	if speed == speeds[0] {
		billet.price = 36000000
	} else if speed == speeds[1] {
		billet.price = 43000000
	} else if speed == speeds[2] {
		billet.price = 50000000
	}
	return billet, nil
}

func (billet *Billet) GetSpaceLineName() SpaceLine {
	return billet.spaceLineName
}
func (billet *Billet) GetTripType() TripType {
	return billet.tripTypeName
}
func (billet *Billet) GetDays() uint32 {
	return billet.days
}
func (billet *Billet) GetPrice() uint32 {
	return billet.price
}

// datas
var Spacelines []SpaceLine = []SpaceLine{
	"Virgin Galactic",
	"SpaceX",
	"Space Adventures",
}
var TripTypes []TripType = []TripType{
	"Round-trip",
	"One-way",
}

var Rockets []Rocket = []Rocket{
	Rocket{name: "Спутник", speed: speeds[2]},
	Rocket{name: "Спутник-3", speed: speeds[1]},
	Rocket{name: "Thor Able", speed: speeds[0]},
	Rocket{name: "Восток", speed: speeds[2]},
	Rocket{name: "Titan II", speed: speeds[1]},
}
