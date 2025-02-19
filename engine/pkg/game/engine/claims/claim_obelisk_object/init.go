package claim_obelisk_object

import (
	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/constants"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/storage"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

func (claimObelisk *ClaimObeliskObject) Init(e entity.IEngine) bool {
	charGameObj := e.GameObjects()[claimObelisk.Properties()["crafted_by_character_id"].(string)]
	if charGameObj == nil {
		return false
	}

	// Create claim area
	additionalProps := make(map[string]interface{})
	additionalProps["claim_obelisk_id"] = claimObelisk.Id()
	claimArea := e.CreateGameObject("claim/claim_area", claimObelisk.X() - constants.ClaimArea / 2, claimObelisk.Y() - constants.ClaimArea / 2, 0.0, claimObelisk.Floor(), additionalProps)

	claimObelisk.Properties()["claim_area_id"] = claimArea.Id()

	// Init rent
	claimObelisk.Properties()["payed_until"] = float64(utils.MakeTimestamp()) + constants.ClaimRentDuration

	delayedAction := &entity.DelayedAction{
		FuncName: "ExpireClaim",
		Params: map[string]interface{}{
			"claim_obelisk_id": claimObelisk.Id(),
		},
		TimeLeft: constants.ClaimRentDuration,
		Status: entity.DelayedActionReady,
	}
	claimObelisk.SetCurrentAction(delayedAction)

	// Set claim obelisk id for character
	charGameObj.Properties()["claim_obelisk_id"] = claimObelisk.Id()

	storage.GetClient().Updates <- claimObelisk.Clone()
	storage.GetClient().Updates <- charGameObj.Clone()

	e.SendGameObjectUpdate(claimArea, "add_object")

	return true
}
