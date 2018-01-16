package hass

import "errors"

// Lock describes a Lock class
type Lock struct {
	id     string
	state  string
	access *Access
}

// NewLock creates a new Lock instance
func (a *Access) NewLock(id string) (lock *Lock) {
	return &Lock{
		id:     id,
		access: a,
	}
}

// On locks a lock
func (l *Lock) On() (err error) {
	return l.access.CallService("lock", "lock", l.id)
}

// Off unlocks a lock
func (l *Lock) Off() (err error) {
	return l.access.CallService("lock", "unlock", l.id)
}

// Toggle is supposed to toggle the lock, but is not implemented
func (l *Lock) Toggle() (err error) {
	return errors.New("The command 'toggle' is not implemented for locks")
}

// EntityID returns the id of the device object
func (l *Lock) EntityID() string {
	return l.id
}

// Domain returns the Home Assistant domain for the device
func (l *Lock) Domain() string {
	return "lock"
}
