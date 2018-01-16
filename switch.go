package hass

// Switch describes a Switch class
type Switch struct {
	id     string
	state  string
	access *Access
}

// NewSwitch creates a new Switch instance
func (a *Access) NewSwitch(id string) (s *Switch) {
	// d := strings.SplitAfter(id, ".")[0]
	return &Switch{
		id:     id,
		access: a,
	}
}

// On turns on a switch
func (s *Switch) On() (err error) {
	return s.access.CallService("switch", "turn_on", s.id)
}

// Off turns off a switch
func (s *Switch) Off() (err error) {
	return s.access.CallService("switch", "turn_off", s.id)
}

// Toggle toggles a switch
func (s *Switch) Toggle() (err error) {
	return s.access.CallService("switch", "toggle", s.id)
}

// EntityID returns the id of the device object
func (s *Switch) EntityID() string {
	return s.id
}

// Domain returns the Home Assistant domain for the device
func (s *Switch) Domain() string {
	return "switch"
}
