package main

// gociadm - scans an oci env and checks for running instances.
// If it determines that those instances are running then it will
// stop them after a certain time. It will not include instances
// that are in the OKE node pool.
func main() {
	// every 5 seconds scan for instances in compartment
	scanOciEnv()

	// determine if compute instances are equal to OKE node-pool instances
	removeOkeInstancesFromList()

	// stop compute instances that are running
	stopComputeInstances()
}

func scanOciEnv() {}

func removeOkeInstancesFromList() {}

func stopComputeInstances() {}
