package main

import (
	"fmt"
	"log"

	nks "github.com/NetApp/nks-sdk-go/nks"
)

func main() {
	// Set up HTTP client with environment variables for API token and URL
	client, err := nks.NewClientFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	orgID, err := nks.GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get list of configured clusters
	clusters, err := client.GetClusters(orgID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print list of clusters
	for i := 0; i < len(clusters); i++ {
		fmt.Printf("Cluster(%d): %v\n", clusters[i].ID, clusters[i].Name)
	}
	if len(clusters) == 0 {
		fmt.Println("Sorry, no clusters defined yet")
		return
	}
	// Get cluster ID from user to list solutions from
	var clusterID int
	fmt.Printf("Enter cluster ID to list solutions from: ")
	fmt.Scanf("%d", &clusterID)

	// Get list of solutions configured
	solutions, err := client.GetSolutions(orgID, clusterID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// List solutions
	for i := 0; i < len(solutions); i++ {
		fmt.Printf("Solution(%d): %s solution is %s\n", solutions[i].ID, solutions[i].Name, solutions[i].State)
	}
	if len(solutions) == 0 {
		fmt.Printf("Sorry, no solutions found\n")
		return
	}
	// Get solution ID from user to inspect
	var solutionID int
	fmt.Printf("Enter solution ID to inspect: ")
	fmt.Scanf("%d", &solutionID)

	solution, err := client.GetSolution(orgID, clusterID, solutionID)
	if err != nil {
		log.Fatal(err)
	}
	nks.PrettyPrint(solution)
}
