package craft

func GetAtlas() map[string]interface{} {
	craftAtlas:= map[string]interface{}{
		"stone_wall": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 2.0,
			},
			"title": "Stone Wall",
			"description": "Protects from strangers and keeps your animals safe.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
			}, //tools equipped required to craft something
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 5000.0, // ms
			"width": 1.0,
			"height": 2.0,
		},
		"wooden_wall": map[string]interface{}{
			"skill": "lumberjacking",
			"resources": map[string]interface{}{
				"log": 3.0,
			},
			"title": "Wooden Wall",
			"description": "Protects from strangers and keeps your animals safe.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer", "axe",
			}, //tools equipped required to craft something
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 6000.0, // ms
			"width": 0.3,
			"height": 3.0,
		},
		"wooden_shovel": map[string]interface{}{
			"skill": "lumberjacking",
			"resources": map[string]interface{}{
				"log": 2.0,
			},
			"title": "Wooden Shovel",
			"description": "Basic shovel to dig fields for your crops.",
			"inputs": []string{},
			"tools": []string{},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"stone_hammer": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 1.0,
				"log": 1.0,
			},
			"title": "Stone Hammer",
			"description": "Basic hammer used to craft things.",
			"inputs": []string{},
			"tools": []string{},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"stone_spear": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 1.0,
				"log": 1.0,
			},
			"title": "Stone Spear",
			"description": "Basic weapon to defend yourself.",
			"inputs": []string{},
			"tools": []string{
				"axe",
			},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"stone_knife": map[string]interface{}{
			"skill": "stoneworking",
			"resources": map[string]interface{}{
				"stone": 1.0,
				"log": 1.0,
			},
			"title": "Stone Knife",
			"description": "Useful to cut something like cactus.",
			"inputs": []string{},
			"tools": []string{
				"axe",
			},
			"place_in_real_world": false,
			"duration": 5000.0,
		},
		"healing_balm": map[string]interface{}{
			"skill": "herbalism",
			"resources": map[string]interface{}{
				"cactus_slice": 2.0,
			},
			"title": "Healing Balm",
			"description": "Useful to heal small wounds.",
			"inputs": []string{},
			"tools": []string{
				"knife",
			},
			"place_in_real_world": false,
			"duration": 7000.0,
		},
		"fire_dragon_hatchery": map[string]interface{}{
			"skill": "taming",
			"resources": map[string]interface{}{
				"stone": 2.0,
				"log": 2.0,
				"fire_dragon_egg": 1.0,
			},
			"title": "Fire Dragon Hatchery",
			"description": "Want a fire dragon? Hatching time is one minute.",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
				"axe",
			}, //tools equipped required to craft something
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 10000.0, // ms
			"width": 2.0,
			"height": 2.0,
		},
		"claim_obelisk": map[string]interface{}{
			"skill": "householding",
			"resources": map[string]interface{}{
				"stone": 2.0,
				"claim_stone": 1.0,
			},
			"title": "Claim Obelisk",
			"description": "A first step to build your own home",
			"inputs": []string{
				"coordinates",
				"rotation",
			},
			"tools": []string{
				"hammer",
			}, //tools equipped required to craft something
			"place_in_real_world": true, //place item in real world or put into container
			"duration": 20000.0, // ms
			"width": 1.0,
			"height": 1.0,
		},
	}

	return craftAtlas
}
