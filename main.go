package main

import (
	blockchain "go-testspace/blockchain"
	gobcs "go-testspace/go-basics"
	goch "go-testspace/go-challenges"
	k8s "go-testspace/kubernetes"
)

func main() {

	// Go Challenges Section
	goch.ImplementTypeInterface()
	goch.OverrideInterfaceFunction()
	goch.ImplementOverrideSortFunction()
	goch.FindUnique()
	goch.StackImplementation()
	goch.QueueImplementation()
	goch.MinMax()
	goch.CheckPermutations()
	goch.SingleLinkedList()
	goch.PlayWithSyncMap()
	goch.ClosestMatch()

	// // Basics Section

	gobcs.PlayWithPointers()
	gobcs.PlayWithSlices()
	gobcs.PlayWithMaps()

	// // K8's Section
	k8s.Interact()

	// // BlockChain Section
	blockchain.Initiate()

}
