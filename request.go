package hass

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Access struct {
	host     string
	password string
}

// NewAccess returns a new *Access to be used to interface with the
// Home Assistant system.
func NewAccess(host, password string) *Access {
	return &Access{host, password}
}

func (a *Access) httpGet(path string, v interface{}) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", a.host+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-ha-access", a.password)

	success := false
	for i := 0; i < 3; i++ {
		func() {
			var resp *http.Response
			resp, err = client.Do(req)
			if err != nil {
				return
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				err = errors.New("hass: status not OK: " + resp.Status)
				return
			}

			dec := json.NewDecoder(resp.Body)
			err = dec.Decode(v)
			success = true
		}()

		if success {
			break
		}
	}

	return err
}

func (a *Access) httpPost(path string, v interface{}) error {
	var req *http.Request

	if v != nil {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}

		req, err = http.NewRequest("POST", a.host+path, bytes.NewReader(data))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		var err error
		req, err = http.NewRequest("POST", a.host+path, nil)
		if err != nil {
			return err
		}
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req.Header.Set("x-ha-access", a.password)

	var err error
	success := false
	for i := 0; i < 3; i++ {
		func() {
			var resp *http.Response
			resp, err = client.Do(req)
			if err != nil {
				return
			}

			defer resp.Body.Close()

			success = true
		}()

		if success {
			break
		}
	}

	return err
}
