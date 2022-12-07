package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	// once the device been allocated we expect the env variable to be there
	// but on first time, it might not, hence this check is useful.
	reg, err := regexp.Compile("EXAMPLECOMDEVICE.*_DEV.*=[0-9]+")
	if err != nil {
		log.Fatalf("%v", err)
	}
	var deviceFound bool
	envs := os.Environ()
	for _, env := range envs {
		if reg.MatchString(env) {
			log.Printf("device found: %q", env)
			deviceFound = true
		}
	}

	if !deviceFound {
		log.Fatalf("did not find a device that is matching the regex %q", reg.String())
	}

	// TODO make it configurable
	fi, err := os.Stat("/host-var/run/test-file")
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("file found: %v", fi.Name())
	for {
	}
}
