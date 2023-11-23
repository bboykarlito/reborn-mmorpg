package entity

import (
	"encoding/json"
	"math"

	"github.com/gavr-games/reborn-mmorpg/pkg/utils"
)

const (
	MaxDistance = 0.1
)

type IGameObject interface {
	X() float64
	SetX(x float64) 
	Y() float64
	SetY(y float64)
	Width() float64
	SetWidth(width float64)
	Height() float64
	SetHeight(height float64)
	Id() string
	SetId(id string)
	Type() string
	SetType(t string)
	Floor() int
	SetFloor(floor int)
	Rotation() float64
	SetRotation(rotation float64)
	CurrentAction() *DelayedAction
	SetCurrentAction( currentAction *DelayedAction)
	Properties() map[string]interface{}
	SetProperties(properties map[string]interface{})
	Effects() map[string]interface{}
	SetEffects(effects map[string]interface{})
	HitBox() utils.Bounds
	IsPoint() bool
	Intersects(b utils.Bounds) bool
	Clone() *GameObject
	GetDistance(b IGameObject) float64
	IsCloseTo(b IGameObject) bool
	Rotate(rotation float64)
}

type GameObject struct {
	// params for quadtree
	x      float64
	y      float64
	width  float64
	height float64

	// game params
	id string
	objType string
	floor int // -1 for does not belong to any floor
	currentAction *DelayedAction
	rotation float64 // from 0 to math.Pi * 2
	properties map[string]interface{}
	effects map[string]interface{}
}

func (obj *GameObject) X() float64 {
	return obj.x
}

func (obj *GameObject) SetX(x float64) {
	obj.x = x
	obj.properties["x"] = x
}

func (obj *GameObject) Y() float64 {
	return obj.y
}

func (obj *GameObject) SetY(y float64) {
	obj.y = y
	obj.properties["y"] = y
}

func (obj *GameObject) Width() float64 {
	return obj.width
}

func (obj *GameObject) SetWidth(width float64) {
	obj.width = width
	obj.properties["width"] = width
}

func (obj *GameObject) Height() float64 {
	return obj.height
}

func (obj *GameObject) SetHeight(height float64) {
	obj.height = height
	obj.properties["height"] = height
}

func (obj *GameObject) Id() string {
	return obj.id
}

func (obj *GameObject) SetId(id string) {
	obj.id = id
	obj.properties["id"] = id
}

func (obj *GameObject) Type() string {
	return obj.objType
}

func (obj *GameObject) SetType(t string) {
	obj.objType = t
	obj.properties["type"] = t
}

func (obj *GameObject) Floor() int {
	return obj.floor
}

func (obj *GameObject) SetFloor(floor int) {
	obj.floor = floor
}

func (obj *GameObject) Rotation() float64 {
	return obj.rotation
}

func (obj *GameObject) SetRotation(rotation float64) {
	obj.rotation = rotation
}

func (obj *GameObject) CurrentAction() *DelayedAction {
	return obj.currentAction
}

func (obj *GameObject) SetCurrentAction(currentAction *DelayedAction) {
	obj.currentAction = currentAction
}

func (obj *GameObject) Properties() map[string]interface{} {
	return obj.properties
}

func (obj *GameObject) SetProperties(properties map[string]interface{}) {
	obj.properties = properties
}

func (obj *GameObject) Effects() map[string]interface{} {
	return obj.effects
}

func (obj *GameObject) SetEffects(effects map[string]interface{}) {
	obj.effects = effects
}

func (obj *GameObject) UnmarshalJSON(b []byte) error {
	var tmp struct {
		X float64
		Y float64
		Width float64
		Height float64
		Id string
		Type string
		Floor int 
		CurrentAction *DelayedAction
		Rotation float64
		Properties map[string]interface{}
		Effects map[string]interface{}
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
			return err
	}
	obj.x = tmp.X
	obj.y = tmp.Y
	obj.width = tmp.Width
	obj.height = tmp.Height
	obj.id = tmp.Id
	obj.objType = tmp.Type
	obj.floor = tmp.Floor
	obj.currentAction = tmp.CurrentAction
	obj.rotation = tmp.Rotation
	obj.properties = tmp.Properties
	obj.effects = tmp.Effects
	return nil 
}

func (obj *GameObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		X float64
		Y float64
		Width float64
		Height float64
		Id string
		Type string
		Floor int 
		CurrentAction *DelayedAction
		Rotation float64
		Properties map[string]interface{}
		Effects map[string]interface{}
	}{
		X: obj.x,
		Y: obj.y,
		Width: obj.width,
		Height: obj.height,
		Id: obj.id,
		Type: obj.objType,
		Floor: obj.floor,
		CurrentAction: obj.currentAction,
		Rotation: obj.rotation,
		Properties: obj.properties,
		Effects: obj.effects,
	})
}

func (obj *GameObject) HitBox() utils.Bounds {
	return utils.Bounds{
		X: obj.X(),
		Y: obj.Y(),
		Width: obj.Width(),
		Height: obj.Height(),
	}
}

//IsPoint - Checks if a bounds object is a point or not (has no width or height)
func (obj *GameObject) IsPoint() bool {
	if obj.Width() == 0 && obj.Height() == 0 {
		return true
	}
	return false
}

// Intersects - Checks if a Bounds object intersects with another Bounds
func (a GameObject) Intersects(b utils.Bounds) bool {
	aMaxX := a.X() + a.Width()
	aMaxY := a.Y() + a.Height()
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	// a is left of b
	if aMaxX < b.X {
		return false
	}

	// a is right of b
	if a.X() > bMaxX {
		return false
	}

	// a is above b
	if aMaxY < b.Y {
		return false
	}

	// a is below b
	if a.Y() > bMaxY {
		return false
	}

	// The two overlap
	return true
}

func (obj *GameObject) Clone() *GameObject {
	clone := &GameObject{
		x: obj.X(),
		y: obj.Y(),
		width: obj.Width(),
		height: obj.Height(),
		id: obj.Id(),
		objType: obj.Type(),
		floor: obj.Floor(),
		rotation: obj.Rotation(),
		properties: make(map[string]interface{}),
		effects: make(map[string]interface{}),
	}
	clone.SetProperties(utils.CopyMap(obj.Properties()))
	clone.SetEffects(utils.CopyMap(obj.Effects()))
	return clone
}

// Get approximate distance between objects. Assuming all of them are rectangles
func (a GameObject) GetDistance(b IGameObject) float64 {
	aXCenter := a.X() + a.Width() / 2
	aYCenter := a.Y() + a.Height() / 2
	
	bXCenter := b.X() + b.Width() / 2
	bYCenter := b.Y() + b.Height() / 2

	xDistance := math.Abs(aXCenter - bXCenter) - (a.Width() / 2 + b.Width() / 2)
	if xDistance < 0 {
		xDistance = 0.0
	}

	yDistance := math.Abs(aYCenter - bYCenter) - (a.Height() / 2 + b.Height() / 2)
	if yDistance < 0 {
		yDistance = 0.0
	}

	return math.Sqrt(math.Pow(xDistance, 2.0) + math.Pow(yDistance, 2.0))
}

// Determines if 2 objects are close enough to each other
func (a GameObject) IsCloseTo(b IGameObject) bool {
	if (a.Floor() != b.Floor()) {
		return false
	}
	return a.GetDistance(b) < MaxDistance
}

// Rotates Game object. Possible rotations 0 and 1 (0 and 90 dergrees)
func (obj *GameObject) Rotate(rotation float64) {
	if obj.Rotation() != rotation {
		if rotation == 0 {
			obj.SetRotation(rotation)
		} else {
			obj.SetRotation(math.Pi / 2)
		}
		width := obj.Width()
		obj.SetWidth(obj.Height())
		obj.SetHeight(width)
	}
}
