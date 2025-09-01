package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tomba-io/go/tomba"
	"github.com/tomba-io/go/tomba/models"
)

func main() {
	client := tomba.New("ta_xx", "ts_xxx")
	// Get all search bulks
	// searchBulks, err := client.GetAllSearchBulks(&models.BulkGetParams{
	// 	Page:  1,
	// 	Limit: 10,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(searchBulks)
	// fmt.Printf("Found %d search bulks\n", len(searchBulks.Data))

	// Create a new search bulk
	searchParams := &models.BulkSearchCreateParams{
		BulkCreateParams: models.BulkCreateParams{
			Name:   "My Search Bulk",
			List:   "tomba.io\ndaily.dev",
			Verify: false,
		},
		BulkSearchParams: models.BulkSearchParams{
			Maximum: "10",
			EmailType: models.BulkSearchParamsType{
				Type:         "all",
				PriorityType: "only",
			},
			Department: models.BulkSearchParamsDepartment{
				Name:         []string{"executive"}, // Example department names
				PriorityType: "only",
			},
		},
	}

	result, err := client.CreateSearchBulk(searchParams)
	if err != nil {
		log.Fatal(err)
	}
	if result.Data.ID != nil {
		fmt.Printf("Created bulk with ID: %d\n", *result.Data.ID)
		// Launch the bulk
		_, err = client.LaunchBulk(models.BulkTypeSearch, *result.Data.ID)
		if err != nil {
			log.Fatal(err)
		}

		// Get progress
		progress, err := client.GetBulkProgress(models.BulkTypeSearch, *result.Data.ID)
		if err != nil {
			log.Fatal(err)
		}

		// Wait for the bulk to complete (this is a simple example, you might want to implement a more robust waiting mechanism)
		for progress.Progress < 100 {
			fmt.Printf("Waiting for bulk to complete... Current progress: %d%%\n", progress.Progress)
			// You can add a sleep here to avoid hitting the API too frequently
			time.Sleep(1 * time.Second)
			progress, err = client.GetBulkProgress(models.BulkTypeSearch, *result.Data.ID)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Download results when completed (this would typically be done after checking progress)
		downloadData, err := client.DownloadBulk(models.BulkTypeSearch, *result.Data.ID, &models.BulkDownloadParams{
			Type: "full",
		})
		if err != nil {
			log.Fatal(err)
		}
		// save the downloaded data to a file or process it as needed
		if len(downloadData) != 0 {
			fmt.Printf("Downloaded %d bytes of results\n", len(downloadData))
			// Save full results
			err = client.SaveBulkResults(models.BulkTypeSearch, *result.Data.ID, searchParams.Name+".csv", "full")
			if err != nil {
				log.Printf("Error saving full results: %v", err)
			} else {
				fmt.Printf("✅ Full results saved to: %s\n", searchParams.Name+".csv")
			}
		}
		log.Fatal("No data downloaded, the bulk might not be completed yet. or No result found.")
	}
}
