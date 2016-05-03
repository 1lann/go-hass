package hass

import (
	"errors"
	"time"
)

var ErrAPINotRunning = errors.New("hass: API is not running")

// CheckAPI checks whether or not the API is running. It returns an error
// if it is not running.
func (a *Access) CheckAPI() error {
	response := struct {
		Message string `json:"message"`
	}{}
	err := a.httpGet("/api/", &response)
	if err != nil {
		return err
	}

	if response.Message == "" {
		return ErrAPINotRunning
	}

	return nil
}

type Bootstrap struct {
	Config struct {
		Components      []string `json:"components"`
		Latitude        float64  `json:"latitude"`
		Longitude       float64  `json:"longitude"`
		LocationName    string   `json:"location_name"`
		TemperatureUnit string   `json:"temperature_unit"`
		Timezone        string   `json:"time_zone"`
		Version         string   `json:"version"`
	} `json:"config"`
	Events []struct {
		Event         string `json:"event"`
		ListenerCount int    `json:"listener_count"`
	} `json:"events"`
	Services []struct {
		Domain   string `json:"domain"`
		Services map[string]struct {
			Description string      `json:"description"`
			Fields      interface{} `json:"fields"`
		} `json:"services"`
	} `json:"services"`
	States []struct {
		Attributes struct {
			Auto         bool     `json:"auto"`
			EntityID     []string `json:"entity_id"`
			FriendlyName string   `json:"friendly_name"`
			Hidden       bool     `json:"hidden"`
			Order        int      `json:"order"`
		} `json:"attributes"`
		EntityID    string    `json:"entity_id"`
		LastChanged time.Time `json:"last_changed"`
		LastUpdated time.Time `json:"last_updated"`
		State       string    `json:"state"`
	} `json:"states"`
}

// Bootstrap returns the bootstrap information of the system.
func (a *Access) Bootstrap() (Bootstrap, error) {
	var bootstrap Bootstrap
	err := a.httpGet("/api/bootstrap", &bootstrap)
	if err != nil {
		return Bootstrap{}, err
	}

	return bootstrap, nil
}

// FireEvent fires an event.
func (a *Access) FireEvent(eventType string, eventData interface{}) error {
	return a.httpPost("/api/events/"+eventType, eventData)
}

// CallService calls a service with a domain, service, and entity id.
func (a *Access) CallService(domain, service string, entityID string) error {
	serviceData := struct {
		EntityID string `json:"entity_id"`
	}{entityID}

	return a.httpPost("/api/services/"+domain+"/"+service, serviceData)
}
