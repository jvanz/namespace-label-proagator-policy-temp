package main

import (
	// "errors"
	"fmt"

	appsv1 "github.com/kubewarden/k8s-objects/api/apps/v1"
	batchv1 "github.com/kubewarden/k8s-objects/api/batch/v1"
	corev1 "github.com/kubewarden/k8s-objects/api/core/v1"
	apimachineryv1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubernetes "github.com/kubewarden/policy-sdk-go/pkg/capabilities/kubernetes"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	"github.com/mailru/easyjson"
)

const DEPLOYMENT_KIND = "deployment"
const REPLICASET_KIND = "replicaset"
const STATEFULSET_KIND = "statefulset"
const DAEMONSET_KIND = "daemonset"
const REPLICATIONCONTROLLER_KIND = "replicationController"
const CRONJOB_KIND = "cronjob"
const JOB_KIND = "job"
const POD_KIND = "Pod"

func getNamespace(validationRequest kubewarden_protocol.ValidationRequest) (*corev1.Namespace, error) {

	metadata, err := getResourceMetadata(validationRequest)
	if err != nil {
		return nil, err
	}
	host := getWapcHost()

	resourceRequest := kubernetes.GetResourceRequest{
		APIVersion: "v1",
		Kind:       "Namespace",
		Name:       metadata.Namespace,
	}

	responseBytes, err := kubernetes.GetResource(&host, resourceRequest)
	if err != nil {
		return nil, fmt.Errorf("Cannot get namespace data")
	}
	namespace := &corev1.Namespace{}
	if err := easyjson.Unmarshal(responseBytes, namespace); err != nil {
		return nil, fmt.Errorf("Cannot parse namespace data")
	}
	return namespace, nil
}

func validateResourceLabels(namespaceLabels map[string]string, request kubewarden_protocol.ValidationRequest, settings Settings) ([]byte, error) {
	newLabels := make(map[string]string)
	for _, label := range settings.PropagatedLabels {
		if value, namespace_has_label := namespaceLabels[label]; namespace_has_label {
			newLabels[label] = value
		}
	}
	return updateResourceLabels(request, newLabels)

}
func getResourceMetadata(object kubewarden_protocol.ValidationRequest) (*apimachineryv1.ObjectMeta, error) {
	switch object.Request.Kind.Kind {
	case DEPLOYMENT_KIND:
		deployment := appsv1.Deployment{}
		if err := easyjson.Unmarshal(object.Request.Object, &deployment); err != nil {
			return nil, err
		}
		return deployment.Metadata, nil
	case REPLICASET_KIND:
		replicaset := appsv1.ReplicaSet{}
		if err := easyjson.Unmarshal(object.Request.Object, &replicaset); err != nil {
			return nil, err
		}
		return replicaset.Metadata, nil
	case STATEFULSET_KIND:
		statefulset := appsv1.StatefulSet{}
		if err := easyjson.Unmarshal(object.Request.Object, &statefulset); err != nil {
			return nil, err
		}
		return statefulset.Metadata, nil
	case DAEMONSET_KIND:
		daemonset := appsv1.DaemonSet{}
		if err := easyjson.Unmarshal(object.Request.Object, &daemonset); err != nil {
			return nil, err
		}
		return daemonset.Metadata, nil
	case REPLICATIONCONTROLLER_KIND:
		replicationController := corev1.ReplicationController{}
		if err := easyjson.Unmarshal(object.Request.Object, &replicationController); err != nil {
			return nil, err
		}
		return replicationController.Metadata, nil
	case CRONJOB_KIND:
		cronjob := batchv1.CronJob{}
		if err := easyjson.Unmarshal(object.Request.Object, &cronjob); err != nil {
			return nil, err
		}
		return cronjob.Metadata, nil
	case JOB_KIND:
		job := batchv1.Job{}
		if err := easyjson.Unmarshal(object.Request.Object, &job); err != nil {
			return nil, err
		}
		return job.Metadata, nil
	case POD_KIND:
		pod := corev1.Pod{}
		if err := easyjson.Unmarshal(object.Request.Object, &pod); err != nil {
			return nil, err
		}
		return pod.Metadata, nil
	default:
		return nil, fmt.Errorf("object should be one of these kinds: Deployment, ReplicaSet, StatefulSet, DaemonSet, ReplicationController, Job, CronJob, Pod. Found %s", object.Request.Kind.Kind)
	}
}

func getResourceLabels(object kubewarden_protocol.ValidationRequest) (map[string]string, error) {
	metadata, err := getResourceMetadata(object)
	if err != nil {
		return nil, err
	}
	return metadata.Labels, nil
}

func updateLabels(labels *map[string]string, newLabels map[string]string) bool {
	has_mutation := false
	for label, newValue := range newLabels {
		if oldValue, has_label := (*labels)[label]; !has_label || oldValue != newValue {
			(*labels)[label] = newValue
			has_mutation = has_mutation || true
		}
	}
	return has_mutation
}

func updateResourceLabels(object kubewarden_protocol.ValidationRequest, newLabels map[string]string) ([]byte, error) {
	switch object.Request.Kind.Kind {
	case DEPLOYMENT_KIND:
		deployment := appsv1.Deployment{}
		if err := easyjson.Unmarshal(object.Request.Object, &deployment); err != nil {
			return nil, err
		}
		if deployment.Metadata.Labels == nil {
			deployment.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&deployment.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(deployment)
		}
		return kubewarden.AcceptRequest()
	case REPLICASET_KIND:
		replicaset := appsv1.ReplicaSet{}
		if err := easyjson.Unmarshal(object.Request.Object, &replicaset); err != nil {
			return nil, err
		}
		if replicaset.Metadata.Labels == nil {
			replicaset.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&replicaset.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(replicaset)
		}
		return kubewarden.AcceptRequest()
	case STATEFULSET_KIND:
		statefulset := appsv1.StatefulSet{}
		if err := easyjson.Unmarshal(object.Request.Object, &statefulset); err != nil {
			return nil, err
		}
		if statefulset.Metadata.Labels == nil {
			statefulset.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&statefulset.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(statefulset)
		}
		return kubewarden.AcceptRequest()
	case DAEMONSET_KIND:
		daemonset := appsv1.DaemonSet{}
		if err := easyjson.Unmarshal(object.Request.Object, &daemonset); err != nil {
			return nil, err
		}
		if daemonset.Metadata.Labels == nil {
			daemonset.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&daemonset.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(daemonset)
		}
		return kubewarden.AcceptRequest()
	case REPLICATIONCONTROLLER_KIND:
		replicationController := corev1.ReplicationController{}
		if err := easyjson.Unmarshal(object.Request.Object, &replicationController); err != nil {
			return nil, err
		}
		if replicationController.Metadata.Labels == nil {
			replicationController.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&replicationController.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(replicationController)
		}
		return kubewarden.AcceptRequest()
	case CRONJOB_KIND:
		cronjob := batchv1.CronJob{}
		if err := easyjson.Unmarshal(object.Request.Object, &cronjob); err != nil {
			return nil, err
		}
		if cronjob.Metadata.Labels == nil {
			cronjob.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&cronjob.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(cronjob)
		}
		return kubewarden.AcceptRequest()
	case JOB_KIND:
		job := batchv1.Job{}
		if err := easyjson.Unmarshal(object.Request.Object, &job); err != nil {
			return nil, err
		}
		if job.Metadata.Labels == nil {
			job.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&job.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(job)
		}
		return kubewarden.AcceptRequest()
	case POD_KIND:
		pod := corev1.Pod{}
		if err := easyjson.Unmarshal(object.Request.Object, &pod); err != nil {
			return nil, err
		}
		if pod.Metadata.Labels == nil {
			pod.Metadata.Labels = make(map[string]string)
		}
		if updateLabels(&pod.Metadata.Labels, newLabels) {
			return kubewarden.MutateRequest(pod)
		}
		return kubewarden.AcceptRequest()
	default:
		return nil, fmt.Errorf("object should be one of these kinds: Deployment, ReplicaSet, StatefulSet, DaemonSet, ReplicationController, Job, CronJob, Pod. Found %s", object.Request.Kind.Kind)
	}
}

func validate(payload []byte) ([]byte, error) {
	// Create a ValidationRequest instance from the incoming payload
	validationRequest := kubewarden_protocol.ValidationRequest{}
	err := easyjson.Unmarshal(payload, &validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}
	fmt.Print(validationRequest.Request.Kind)

	// Create a Settings instance from the ValidationRequest object
	settings, err := NewSettingsFromValidationReq(&validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	namespace, err := getNamespace(validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(kubewarden.Message(err.Error()), kubewarden.Code(400))
	}

	return validateResourceLabels(namespace.Metadata.Labels, validationRequest, settings)
}
