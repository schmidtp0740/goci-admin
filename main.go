package main

import (
	"log"
	"os"
	"time"

	"github.com/schmidtp0740/goci-admin/pkg/admin"
	"github.com/schmidtp0740/goci-admin/pkg/scan"
)

// gociadm - scans an oci env and checks for running instances.
// If it determines that those instances are running then it will
// stop them after a certain time. It will not include instances
// that are in the OKE node pool.
func main() {

	if os.Getenv("C") == "" {
		log.Fatalf("Must declare environement variable $C with compartment-id")
		return
	}

	// run 4ever to scan for running instances then stop them
	for {

		// every 5 seconds scan for instances in compartment
		// change this timer to some time after 6 pm CST
		log.Printf("Waiting 5 seconds\n")
		time.Sleep(5 * time.Second)

		listOfRunningInstances := scan.OciEnv()

		// stop compute instances that are running
		admin.StopOCIComputeInstances(listOfRunningInstances)

	}

}
