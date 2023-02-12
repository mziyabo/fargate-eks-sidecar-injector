package shared

import (
	"context"

	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//
type K8sUtil struct{}

var (
	client *kubernetes.Clientset
)

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
}

// GetConfigMap returns ConfigMap data
func (k K8sUtil) GetConfigMap(namespace, configMapName string) map[string]string {
	configMap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	return configMap.Data
}
