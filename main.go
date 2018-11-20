package main

import (
	"context"
	"log"
	"os"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

// gociadm - scans an oci env and checks for running instances.
// If it determines that those instances are running then it will
// stop them after a certain time. It will not include instances
// that are in the OKE node pool.
func main() {

	if os.Getenv("C") == "" {
		log.Printf("Must declare environement variable $C with compartment-id")
		return
	}
	// every 5 seconds scan for instances in compartment
	scanOciEnv()

	// stop compute instances that are running
	stopComputeInstances()
}

func scanOciEnv() (listOfRunnningInstances []string) {

	// get default configuration file at $HOME/.oci/config
	// TODO change this to another method other than default file and path?
	configProvider := common.DefaultConfigProvider()

	// create compute client with configuration provider
	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		log.Println("error getting configuration file: " + err.Error())
		return
	}

	request := core.ListInstancesRequest{}

	request.CompartmentId = common.String(os.Getenv("C"))

	request.LifecycleState = core.InstanceLifecycleStateRunning

	response, err := client.ListInstances(context.Background(), request)
	if err != nil {
		log.Println("error getting list of instances" + err.Error())
	}

	// scan through items and print their details
	for _, item := range response.Items {
		log.Println("---Found new Running Instance---")

		log.Printf("Display Name: %s\n", *item.DisplayName)
		log.Printf("ID: %s\n", *item.Id)
		// starting of finding oke instances
		if val, ok := item.Metadata["oke-cluster-id"]; ok {
			log.Printf("\t---Found OKE Node Instance---")
			log.Printf("\tFound %s, value: %s\n", "oke-cluster-id", val)
			if val, ok := item.Metadata["oke-pool-id"]; ok {
				log.Printf("\tFound %s, value: %s\n", "oke-pool-id", val)
			}
			log.Printf("\tNot adding instance to list\n")
			continue
		}

		log.Printf("Added to list...")
		listOfRunnningInstances = append(listOfRunnningInstances, *item.Id)
	}

	log.Printf("%s", listOfRunnningInstances)

	return listOfRunnningInstances
}

func stopComputeInstances() {}
