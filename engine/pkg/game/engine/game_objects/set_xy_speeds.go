package game_objects

import (
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Order is simportant here for finding direction by angle in mobs/mob.go
var PossibleDirections = [...]string {"move_east", "move_north_east", "move_north", "move_north_west", "move_west", "move_south_west", "move_south", "move_south_east"}

// Set x and y speeds depending on the direction
func SetXYSpeeds(obj entity.IGameObject, direction string) {
	speed := obj.Properties()["speed"].(float64)
	axisSpeed := math.Sqrt(speed * speed / 2)

	validDirection := false
	for _, dir := range PossibleDirections {
		if dir == direction {
			validDirection = true
			break
		}
	}

	if !validDirection {
		//TODO: log error
		return
	}

	switch direction {
		case "move_north":
			obj.Properties()["speed_x"] = 0.0
			obj.Properties()["speed_y"] = speed
			// this is required to determine where character/mob looks, when stopped
			// it is important to later hit the target with weapon
			obj.SetRotation(math.Pi / 2)
		case "move_south":
			obj.Properties()["speed_x"] = 0.0
			obj.Properties()["speed_y"] = -speed
			obj.SetRotation(math.Pi * 3 / 2)
		case "move_east":
			obj.Properties()["speed_x"] = speed
			obj.Properties()["speed_y"] = 0.0
			obj.SetRotation(0)
		case "move_west":
			obj.Properties()["speed_x"] = -speed
			obj.Properties()["speed_y"] = 0.0
			obj.SetRotation(math.Pi)
		case "move_north_east":
			obj.Properties()["speed_x"] = axisSpeed
			obj.Properties()["speed_y"] = axisSpeed
			obj.SetRotation(math.Pi / 4)
		case "move_north_west":
			obj.Properties()["speed_x"] = -axisSpeed
			obj.Properties()["speed_y"] = axisSpeed
			obj.SetRotation(math.Pi * 3 / 4)
		case "move_south_east":
			obj.Properties()["speed_x"] = axisSpeed
			obj.Properties()["speed_y"] = -axisSpeed
			obj.SetRotation(math.Pi * 7 / 4)
		case "move_south_west":
			obj.Properties()["speed_x"] = -axisSpeed
			obj.Properties()["speed_y"] = -axisSpeed
			obj.SetRotation(math.Pi * 5 / 4)
	}
}

// Set rotation depending on the direction
func SetRotation(obj entity.IGameObject, direction string) {
	obj.SetRotation(GetRotation(direction))
}

func GetRotation(direction string) float64 {
	validDirection := false
	for _, dir := range PossibleDirections {
		if dir == direction {
			validDirection = true
			break
		}
	}

	if !validDirection {
		//TODO: log error
		return 0.0
	}

	switch direction {
		case "move_north":
			return math.Pi / 2
		case "move_south":
			return math.Pi * 3 / 2
		case "move_east":
			return 0.0
		case "move_west":
			return math.Pi
		case "move_north_east":
			return math.Pi / 4
		case "move_north_west":
			return math.Pi * 3 / 4
		case "move_south_east":
			return math.Pi * 7 / 4
		case "move_south_west":
			return math.Pi * 5 / 4
	}

	return 0.0
}
