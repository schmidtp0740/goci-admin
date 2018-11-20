package admin

import (
	"context"
	"log"
	"sync"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

// StopComputeInstances
func StopOCIComputeInstances(listOfRunnningInstances []string) {
	configProvider := common.DefaultConfigProvider()

	client, err := core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		log.Fatalf("Error getting compute client with configuration provider: %s", err.Error())
	}

	request := core.InstanceActionRequest{}

	var wg sync.WaitGroup

	for _, instanceID := range listOfRunnningInstances {
		request.Action = core.InstanceActionActionStop

		wg.Add(1)
		go func(instanceID string, client core.ComputeClient, request core.InstanceActionRequest) {
			defer wg.Done()

			request.InstanceId = common.String(instanceID)

			ctx := context.Background()

			response, err := client.InstanceAction(ctx, request)
			if err != nil {
				log.Fatalf("Error unable to stop instance: %s", err.Error())
			}

			log.Printf("---Updated Instance---")
			log.Printf("Instance ID: %s", *response.Id)
			log.Printf("Display-Name: %s", *response.DisplayName)
			log.Printf("Lifecycle status: %s\n", response.LifecycleState)
		}(instanceID, client, request)

	}

	wg.Wait()

}
