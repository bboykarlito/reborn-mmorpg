package delayed_actions

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/engine/game_objects/serializers"
)

func UpdateAll(e entity.IEngine, tickDelta int64) {
	for _, gameObj := range e.GameObjects() {
		delayedAction := gameObj.CurrentAction()
		if gameObj != nil && delayedAction != nil {
			// Moving to coords has higher priority, then action. For example: first move to coords, then build a wall there.
			if gameObj.MoveToCoords() == nil {
				if delayedAction.Status == entity.DelayedActionReady {
					delayedAction.Status = entity.DelayedActionStarted
					e.SendResponseToVisionAreas(gameObj, "start_delayed_action", map[string]interface{}{
						"object": serializers.GetInfo(e.GameObjects(), gameObj),
						"duration": delayedAction.TimeLeft,
						"action": delayedAction.FuncName,
					})
				}
				delayedAction.TimeLeft = delayedAction.TimeLeft - float64(tickDelta)
				if (delayedAction.TimeLeft <= 0.0) {
					Finish(e, gameObj)
				}
			}
		}
	}
}

