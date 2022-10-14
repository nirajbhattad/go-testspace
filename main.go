package main

import (
	"go-testspace/blockchain"
	goch "go-testspace/go-challenges"
	"go-testspace/kubernetes"
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
	goch.MaskingRedaction()

	// K8's Section
	kubernetes.Interact()

	// BlockChain Section
	blockchain.Initiate()

}
