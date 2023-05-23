// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CinderPersistentVolumeSource Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
//
// swagger:model CinderPersistentVolumeSource
type CinderPersistentVolumeSource struct {

	// fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	FSType string `json:"fsType,omitempty"`

	// readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	ReadOnly bool `json:"readOnly,omitempty"`

	// secretRef is Optional: points to a secret object containing parameters used to connect to OpenStack.
	SecretRef *SecretReference `json:"secretRef,omitempty"`

	// volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// Required: true
	VolumeID *string `json:"volumeID"`
}
