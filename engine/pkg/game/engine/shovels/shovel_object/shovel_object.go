package shovel_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/pickable_object"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/mixins/equipable_object"
)

type ShovelObject struct {
	pickable_object.PickableObject
	equipable_object.EquipableObject
	entity.GameObject
}

func NewShovelObject(gameObj entity.IGameObject) *ShovelObject {
	shovel := &ShovelObject{pickable_object.PickableObject{}, equipable_object.EquipableObject{}, *gameObj.(*entity.GameObject)}
	shovel.InitPickableObject(shovel)
	shovel.InitEquipableObject(shovel)
	return shovel
}
