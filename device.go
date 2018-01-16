package hass

import "strings"

// Device is a generic interface for interacting with devices
type Device interface {
	On() error
	Off() error
	Toggle() error
	EntityID() string
	Domain() string
}

// GetDevice returns a Device object from an State object
func (a *Access) GetDevice(state State) Device {
	dom := strings.TrimSuffix(strings.SplitAfter(state.EntityID, ".")[0], ".")
	switch dom {
	case "light":
		return a.NewLight(state.EntityID)
	case "switch":
		return a.NewSwitch(state.EntityID)
	case "lock":
		return a.NewLock(state.EntityID)
	}
	return nil
}
