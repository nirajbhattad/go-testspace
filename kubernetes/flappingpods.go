package kubernetes

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

type PodDetails struct {
	Name                string                     `json:"name"`
	IP                  string                     `json:"ip"`
	PodPhase            corev1.PodPhase            `json:"podPhase"`
	PodConditionType    string                     `json:"podCondition"`
	PodConditionState   string                     `json:"podConditionState"`
	ClusterName         string                     `json:"clusterName"`
	Namespace           string                     `json:"nameSpace"`
	NodeName            string                     `json:"nodeName"`
	NodeIp              string                     `json:"nodeIp"`
	ContainerNames      []string                   `json:"containerNames"`
	ContainerIp         string                     `json:"containerIp"`
	Label               map[string]string          `json:"label"`
	ContainerStates     []corev1.ContainerStatus   `json:"containerStates"`
	ContainerConditions map[string]string          `json:"containerConditions"`
	ContainerStatuses   map[string]ContainerDetail `json:"containerStatuses"`
}

type ContainerDetail struct {
	Name         string `json:"name"`
	Reason       string `json:"reason"`
	RestartCount int    `json:"restartCount"`
}

func Interact() {

	file, _ := os.UserHomeDir()
	kubeconfig := filepath.Join(
		file, ".kube", "config",
	)
	fmt.Println(file)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	listOptions := metav1.ListOptions{
		Limit: 100,
	}

	podInfo, err := clientset.CoreV1().Pods("default").List(context.TODO(), listOptions)
	if err != nil {
		panic(err.Error())
	}

	var containerNames []string
	var podConditionType string
	var podConditionState string
	flappingPodsData := make([]PodDetails, len(podInfo.Items))

	for _, pod := range podInfo.Items {

		// Filtering On The Basis Of Initial Pod Phase
		if pod.Status.Phase != corev1.PodRunning && pod.Status.Phase != corev1.PodSucceeded {

			containerInit := true
			containerConditions := make(map[string]string)

			for _, v := range pod.Status.Conditions {
				if v.Status == corev1.ConditionFalse {
					containerInit = false
					containerConditions[string(v.Type)] = v.Reason
				}
			}

			// Breaks If The Container Initialization Is Proper
			if containerInit {
				break
			}

			containerStatus := true
			containerStatuses := make(map[string]ContainerDetail)
			for _, status := range pod.Status.ContainerStatuses {
				if status.State.Waiting != nil {
					containerStatus = false
					containerDetail := ContainerDetail{
						Name:         status.Name,
						Reason:       status.State.Waiting.Reason,
						RestartCount: int(status.RestartCount),
					}
					containerStatuses["Waiting"] = containerDetail
				} else if status.State.Terminated != nil {
					containerStatus = false
					containerDetail := ContainerDetail{
						Name:         status.Name,
						Reason:       status.State.Waiting.Reason,
						RestartCount: int(status.RestartCount),
					}
					containerStatuses["Terminated"] = containerDetail
				}
			}

			// Breaks If The Container Status Is Proper
			if containerStatus {
				break
			}

			containerNames = make([]string, len(pod.Spec.Containers))
			for i, container := range pod.Spec.Containers {
				containerNames[i] = container.Name
			}

			podDetails := PodDetails{
				Name:                pod.ObjectMeta.Name,
				IP:                  pod.Status.PodIP,
				PodPhase:            pod.Status.Phase,
				PodConditionType:    podConditionType,
				PodConditionState:   podConditionState,
				ClusterName:         pod.Name,
				Namespace:           pod.Namespace,
				NodeName:            pod.Spec.NodeName,
				NodeIp:              pod.Status.HostIP,
				Label:               pod.Labels,
				ContainerNames:      containerNames,
				ContainerStates:     pod.Status.ContainerStatuses,
				ContainerConditions: containerConditions,
				ContainerStatuses:   containerStatuses,
			}
			flappingPodsData = append(flappingPodsData, podDetails)
		}
	}

	fmt.Println(flappingPodsData)
}
