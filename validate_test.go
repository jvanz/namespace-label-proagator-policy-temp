package main

import (
	"encoding/json"
	"fmt"

	appsv1 "github.com/kubewarden/k8s-objects/api/apps/v1"
	batchv1 "github.com/kubewarden/k8s-objects/api/batch/v1"
	corev1 "github.com/kubewarden/k8s-objects/api/core/v1"
	metav1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	capabilities "github.com/kubewarden/policy-sdk-go/pkg/capabilities"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	kubewarden_testing "github.com/kubewarden/policy-sdk-go/testing"
	"github.com/mailru/easyjson"
	"testing"
)

const SHOULD_ACCEPT = true
const SHOULD_REJECT = false
const SHOULD_MUTATE = true
const NO_MUTATION = false

func buildValidationRequest(propagatedLabels []string, resource easyjson.Marshaler, kind string) ([]byte, error) {
	settings := Settings{PropagatedLabels: propagatedLabels}
	payload, err := kubewarden_testing.BuildValidationRequest(resource, &settings)

	if err != nil {
		return nil, err
	}
	payload, err = updateValidationRequestKind(payload, kind)
	if err != nil {
		return nil, err
	}
	return payload, nil

}
func buildWapcClient(namespaceLabels map[string]string) error {
	clientResponse := corev1.Namespace{
		Metadata: &metav1.ObjectMeta{
			Labels: namespaceLabels,
		},
	}
	var err error
	wapcClient, err = capabilities.NewSuccessfulMockWapcClient(clientResponse)
	if err != nil {
		return err
	}
	return nil

}

func basicResposeValidation(responsePayload []byte, accepted, should_mutate bool) (*kubewarden_protocol.ValidationResponse, error) {
	var response kubewarden_protocol.ValidationResponse
	if err := easyjson.Unmarshal(responsePayload, &response); err != nil {
		return nil, fmt.Errorf("Unexpected error: %+v", err)
	}

	if response.Accepted != accepted {
		return nil, fmt.Errorf("Unexpected rejection: msg %s - code %d", *response.Message, *response.Code)
	}

	if response.MutatedObject == nil && should_mutate {
		return nil, fmt.Errorf("Missing mutated resource")
	}
	return &response, nil

}

func validateLabels(resourceLabels, expectedLabels map[string]string) error {
	for expectedLabel, expectedValue := range expectedLabels {
		if resourceValue, ok := resourceLabels[expectedLabel]; ok {
			if resourceValue != expectedValue {
				return fmt.Errorf("Resource label \"%s\" expected value:  \"%s\". Found \"%s\"", expectedLabel, expectedValue, resourceValue)
			}
		} else {
			return fmt.Errorf("Mutated resource missing label \"%s\"", expectedLabel)
		}

	}

	if len(resourceLabels) != len(expectedLabels) {
		return fmt.Errorf("Mutated resource contains %d labels. But the expected is %d", len(resourceLabels), len(expectedLabels))
	}
	return nil
}

func updateValidationRequestKind(payload []byte, kind string) ([]byte, error) {
	validationRequest := kubewarden_protocol.ValidationRequest{}
	err := easyjson.Unmarshal(payload, &validationRequest)
	if err != nil {
		return nil, err
	}
	validationRequest.Request.Kind.Kind = kind
	return easyjson.Marshal(validationRequest)
}

func TestPodWithNoLabels(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{
		"testing": "foo",
	}

	resource := corev1.Pod{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestPodLabelsShouldNotMutateWithItHasTheExpectedValue(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}

	resource := corev1.Pod{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels: map[string]string{
				"testing":  "foo",
				"testing2": "zzz",
			},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	_, err = basicResposeValidation(responsePayload, SHOULD_ACCEPT, NO_MUTATION)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}
}

func TestPodLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := corev1.Pod{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestDeploymentLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := appsv1.Deployment{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestReplicaSetLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := appsv1.ReplicaSet{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestStatefulSetLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := appsv1.StatefulSet{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestDaemonSetLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := appsv1.DaemonSet{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestReplicationControllerLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := corev1.ReplicationController{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}

func TestCronJobLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := batchv1.CronJob{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}
func TestJobLabelsShouldOverwrittenLabelsOnlyDefinedInSettings(t *testing.T) {
	propagatedLabels := []string{"testing"}
	namespaceLabels := map[string]string{
		"testing":  "foo",
		"testing2": "zpto",
	}
	expectedLabels := map[string]string{"testing": "foo", "testing2": "zzz"}

	resource := batchv1.Job{
		Metadata: &metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
			Labels:    map[string]string{"testing": "bar", "testing2": "zzz"},
		},
	}

	payload, err := buildValidationRequest(propagatedLabels, resource, DEPLOYMENT_KIND)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	err = buildWapcClient(namespaceLabels)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	response, err := basicResposeValidation(responsePayload, SHOULD_ACCEPT, SHOULD_MUTATE)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	mutatedResourceJSON, err := json.Marshal(response.MutatedObject.(map[string]interface{}))
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := easyjson.Unmarshal(mutatedResourceJSON, &resource); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if err := validateLabels(resource.Metadata.Labels, expectedLabels); err != nil {
		t.Error(err.Error())
	}
}
