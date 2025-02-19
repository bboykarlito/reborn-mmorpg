package moving_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// Check object can move and does not collide with other objects
// TODO: implement accurate check for circles
func (mObj *MovingObject) CanMove(e entity.IEngine, dx float64, dy float64) (float64, float64) {
	obj := mObj.gameObj
	floor := e.Floors()[obj.Floor()]
	// floor edges
	newX := obj.X() + dx
	newY := obj.Y() + dy
	if newX < 0.0 || newX > constants.FloorSize || newY < 0.0 || newY > constants.FloorSize {
		return 0.0, 0.0
	}

	// collisions
	possibleCollidableObjects := floor.RetrieveIntersections(utils.Bounds{
		X:      newX,
		Y:      newY,
		Width:  obj.Width(),
		Height: obj.Height(),
	})
	// Filter collidable objects
	n := 0
	for _, val := range possibleCollidableObjects {
		if collidable, ok := val.(entity.IGameObject).Properties()["collidable"]; ok {
			if collidable.(bool) {
				possibleCollidableObjects[n] = val
				n++
			}
		}
	}
	possibleCollidableObjects = possibleCollidableObjects[:n]

	if len(possibleCollidableObjects) == 0 {
		return dx, dy
	} else {
		return 0.0, 0.0
	}
}
