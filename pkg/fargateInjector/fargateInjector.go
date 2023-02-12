package fargateInjector

import (
	"encoding/json"
	"fmt"

	yamlUtil "github.com/ghodss/yaml"
	"github.com/mziyabo/fargate-sidecar-injector/m/v2/pkg/shared"
	"github.com/tidwall/gjson"
	v1beta1 "k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var k8sUtil shared.K8sUtil

func init() {
	k8sUtil = shared.K8sUtil{}
}

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

func Mutate(ar v1beta1.AdmissionReview) (*v1beta1.AdmissionResponse, error) {
	req := ar.Request
	var pod *corev1.Pod

	if err := json.Unmarshal(req.Object.Raw, &pod); err != nil {
		panic(err)
	}

	var patchType v1beta1.PatchType = "JSONPatch"

	if err := json.Unmarshal(req.Object.Raw, &pod); err != nil {
		return &v1beta1.AdmissionResponse{
			UID:              ar.Request.UID,
			Allowed:          false,
			Result:           &metav1.Status{Message: err.Error()},
			Patch:            []byte{},
			PatchType:        &patchType,
			AuditAnnotations: map[string]string{},
			Warnings:         []string{},
		}, err
	}

	patch := injectSidecarContainer(pod)

	return &v1beta1.AdmissionResponse{
		UID:     ar.Request.UID,
		Allowed: true,
		Patch:   patch,
		PatchType: func() *v1beta1.PatchType {
			pt := v1beta1.PatchTypeJSONPatch
			return &pt
		}(),
	}, nil
}

// injectSidecarContainer creates a patch that adds sidecar containers to the podspec
func injectSidecarContainer(pod *corev1.Pod) []byte {
	patches := []patchOperation{}
	// TODO: use down-api for webhook data like namespaces and configmaps
	configMapData := k8sUtil.GetConfigMap("default", "fargate-injector-sidecar-config")

	fargateProfile := pod.Labels["eks.amazonaws.com/fargate-profile"]
	if fargateProfile != "" {
		// NOTE: Filter by fargateProfile for now in future it could be using annotations etc.
		containerPatches := injectContainers(configMapData[fargateProfile])
		volumePatches := injectVolumes(configMapData[fargateProfile])

		// TODO: cleanup this code.. yikes
		f := len(containerPatches) + len(volumePatches)
		var _patches []patchOperation = make([]patchOperation, f)

		copy(_patches[:], containerPatches[:])
		copy(_patches[len(containerPatches):], volumePatches[:])

		patches = _patches
	}

	pt, err := json.Marshal(patches)
	if err != nil {
		panic(err)
	}
	return pt
}

// 
func injectContainers(configMap string) []patchOperation {
	path := "containers"
	patches := []patchOperation{}
	patchValue := getPatch(configMap, path)

	containers := []corev1.Container{}
	json.Unmarshal([]byte(patchValue), &containers)
	for _, container := range containers {

		patch := patchOperation{
			Op:    "add",
			Path:  fmt.Sprintf("/spec/%s/-", path),
			Value: container,
		}
		patches = append(patches, patch)
	}

	return patches
}

//
func injectVolumes(configMap string) []patchOperation {
	path := "volumes"
	patches := []patchOperation{}
	patchValue := getPatch(configMap, path)

	volumes := []corev1.Volume{}
	json.Unmarshal([]byte(patchValue), &volumes)
	for _, volume := range volumes {

		patch := patchOperation{
			Op:    "add",
			Path:  fmt.Sprintf("/spec/%s/-", path),
			Value: volume,
		}
		patches = append(patches, patch)
	}

	return patches
}

// getPatch returns path/key from ConfigMap data
func getPatch(configmapData string, path string) string {
	p, err := yamlUtil.YAMLToJSON([]byte(configmapData))
	if err != nil {
		panic(err)
	}

	return gjson.Get(string(p), fmt.Sprintf("spec.%s", path)).String()
}
