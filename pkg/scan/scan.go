package scan

import (
	"context"
	"log"
	"os"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

// OciEnv ...
func OciEnv() (listOfRunnningInstances []string) {

	// get default configuration file at $HOME/.oci/config
	// TODO change this to another method other than default file and path?
	configProvider := common.DefaultConfigProvider()

	// create compute client with configuration provider
	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		log.Fatalf("error getting configuration file: %s", err.Error())
		return
	}

	request := core.ListInstancesRequest{}

	request.CompartmentId = common.String(os.Getenv("C"))

	request.LifecycleState = core.InstanceLifecycleStateRunning

	response, err := client.ListInstances(context.Background(), request)
	if err != nil {
		log.Fatalf("error getting list of instances: %s", err.Error())
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

	log.Printf("Updated list of running instance")
	for key, instanceID := range listOfRunnningInstances {
		log.Printf("list[%d]: %s\n", key, instanceID)
	}

	return listOfRunnningInstances
}
