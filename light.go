package hass

// Light describes a Light class
type Light struct {
	id     string
	state  string
	access *Access
}

// NewLight creates a new Light instance
func (a *Access) NewLight(id string) (light *Light) {
	return &Light{
		id:     id,
		access: a,
	}
}

// On turns on a light
func (l *Light) On() (err error) {
	return l.access.CallService("light", "turn_on", l.id)
}

// Off turns off a light
func (l *Light) Off() (err error) {
	return l.access.CallService("light", "turn_off", l.id)
}

// Toggle toggles a switch
func (l *Light) Toggle() (err error) {
	return l.access.CallService("light", "toggle", l.id)
}

// EntityID returns the id of the device object
func (l *Light) EntityID() string {
	return l.id
}

// Domain returns the Home Assistant domain for the device
func (l *Light) Domain() string {
	return "light"
}
