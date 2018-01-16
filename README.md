# go-hass
Incomplete API for interfacing with Home Assistant in Go.

## Example use
```go
package main

import (
	"fmt"

	"github.com/pawal/go-hass"
)

func main() {
	h := hass.NewAccess("http://localhost:8123", "")
	err := h.CheckAPI()
	if err != nil {
		panic(err)
	}
	fmt.Println("API ok")

    // Get a filtered list of devices
	list, err := h.FilterStates("switch", "lock", "light")
	if err != nil {
		panic(err)
	}
    // Print that list of devices
	for d := range list {
		fmt.Printf("%s (%s): %s\n", list[d].EntityID,
			list[d].Attributes.FriendlyName,
			list[d].State)
	}

    // Get the state of a device
	s, err := h.GetState("group.kitchen")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Group kitchen state: %s\n", s.State)

    // Create and interact with a device object
    dev := h.GetDevice(s)
    	fmt.Println("DEV: " + dev.EntityID())
	dev.On()
}
```

## License
MIT.
