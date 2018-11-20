package main

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
