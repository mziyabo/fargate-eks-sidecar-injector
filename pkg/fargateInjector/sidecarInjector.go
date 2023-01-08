package injector

import (
	"encoding/json"

	"github.com/mziyabo/fargate-sidecar-injector/m/v2/pkg/shared"
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var awsUtil shared.AWSUtil
var k8s shared.K8sUtil

func init() {
	awsUtil = shared.AWSUtil{}
	k8s = shared.K8sUtil{}
}

func Mutate(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	req := ar.Request
	var pod corev1.Pod

	var patchType v1beta1.PatchType = "JSONPatch"

	if err := json.Unmarshal(req.Object.Raw, &pod); err != nil {
		return &v1beta1.AdmissionResponse{
			UID:              "",
			Allowed:          false,
			Result:           &metav1.Status{Message: err.Error()},
			Patch:            []byte{},
			PatchType:        &patchType,
			AuditAnnotations: map[string]string{},
			Warnings:         []string{},
		}
	}

	// TODO: URGENT: Modify the pod object as needed here
	mutatedPod := injectSidecarContainer(pod)
	patch, _ := json.Marshal(mutatedPod)

	return &v1beta1.AdmissionResponse{
		Allowed: true,
		Patch:   patch,
		PatchType: func() *v1beta1.PatchType {
			pt := v1beta1.PatchTypeJSONPatch
			return &pt
		}(),
	}
}

// TODO: Implement
// injectSidecarContainer injects sidecar containers into the pod
func injectSidecarContainer(pod corev1.Pod) corev1.Pod {
	if awsUtil.IsFargatePod(pod.Namespace, pod.Labels) {
		data := k8s.FetchConfigMap(pod.Namespace, "REPLACE_CONFIGMAP_NAME")
		k8s.InjectSidecarContainers(&pod, data)
	}

	return pod
}

// IsFargatePod checks if pod is scheduled on fargate
func IsFargatePod(namespace string, labels map[string]string) {
	panic("unimplemented")
}
