package shared

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sUtil struct{}

var client *kubernetes.Clientset

func init() {
	// Create a new Kubernetes client.
	config, err := rest.InClusterConfig()
	if err != nil {
		// handle error
	}
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		// handle error
	}
}

// FetchConfigMap returns configmap data
func (k K8sUtil) FetchConfigMap(namespace, configMapName string) map[string]string {
	configMap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		// handle error
	}

	return configMap.Data
}

// 
func (k K8sUtil) InjectSidecarContainers(pod *corev1.Pod, data map[string]string) {

}
