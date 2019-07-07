package models

import "fmt"

// "context"

// Ship - model
type Ship struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Faction - model
type Faction struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Ships []string `json:"ships"`
}

var xwing = &Ship{"1", "X-wing"}
var ywing = &Ship{"1", "Y-wing"}
var awing = &Ship{"1", "A-wing"}

var falcon = &Ship{"4", "Millenium Falcon"}
var homeOne = &Ship{"5", "Home One"}
var tIEFighter = &Ship{"6", "TIE Fighter"}
var tIEInterceptor = &Ship{"7", "TIE Interceptor"}
var executor = &Ship{"8", "Executor"}

var rebels = &Faction{
	"1",
	"Alliance to Restore the Republic",
	[]string{"1", "2", "3", "4", "5"},
}

var empire = &Faction{
	"2",
	"Galactic Empire",
	[]string{"6", "7", "8"},
}

var factions = map[string]*Faction{
	"1": rebels,
	"2": empire,
}

var ships = map[string]*Ship{
	"1": xwing,
	"2": ywing,
	"3": awing,
	"4": falcon,
	"5": homeOne,
	"6": tIEFighter,
	"7": tIEInterceptor,
	"8": executor,
}

var nextShip = 9

// CreateShip - new ship
func CreateShip(shipName string, factionID string) *Ship {
	nextShip = nextShip + 1
	newShip := &Ship{
		fmt.Sprintf("%v", nextShip),
		shipName,
	}
	ships[newShip.ID] = newShip

	faction := GetFaction(factionID)
	if faction != nil {
		faction.Ships = append(faction.Ships, newShip.ID)
	}
	return newShip
}

// GetShip - return a ship
func GetShip(id string) *Ship {
	if ship, ok := ships[id]; ok {
		return ship
	}
	return nil
}

// GetFaction - return a ship
func GetFaction(id string) *Faction {
	if faction, ok := factions[id]; ok {
		return faction
	}
	return nil
}

// GetRebels - return a ship
func GetRebels() *Faction {
	return rebels
}

// GetEmpire - return a ship
func GetEmpire() *Faction {
	return empire
}
