package mob_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (mob *MobObject) MeleeHit(targetObj entity.IGameObject) bool {
	// Check Cooldown
	// here we cast everything to float64, because go restores from json everything as float64
	lastHitAt, hitted := mob.Properties()["last_hit_at"]
	if hitted {
		if float64(utils.MakeTimestamp()) - lastHitAt.(float64) >= mob.Properties()["cooldown"].(float64) {
			mob.Properties()["last_hit_at"] = float64(utils.MakeTimestamp())
		} else {
			return false
		}
	} else {
		mob.Properties()["last_hit_at"] = float64(utils.MakeTimestamp())
	}

	// check collision with target
	if !mob.CanHit(mob, targetObj) {
		return false
	}

	// Send hit attempt to client
	mob.Engine.SendResponseToVisionAreas(mob, "melee_hit_attempt", map[string]interface{}{
		"object": mob,
		"weapon": mob, // mob has all required weapon attributes itself to act like weapon
	})

	// deduct health and update object
	targetObj.Properties()["health"] = targetObj.Properties()["health"].(float64) - mob.Properties()["damage"].(float64)
	if targetObj.Properties()["health"].(float64) <= 0.0 {
		targetObj.Properties()["health"] = 0.0
	}
	// Trigger mob to aggro
	if targetObj.Properties()["type"].(string) == "mob" {
		mob.Engine.Mobs()[targetObj.Id()].Attack(mob.Id())
	}
	mob.Engine.SendGameObjectUpdate(targetObj, "update_object")

	// die if health < 0
	if targetObj.Properties()["health"].(float64) == 0.0 {
		mob.StopAttacking()
		if targetObj.Properties()["type"].(string) == "mob" {
			mob.Engine.Mobs()[targetObj.Id()].Die()
		} else {
			// for characters
			targetObj.(entity.ICharacterObject).Reborn(mob.Engine)
		}
	}

	return true
}