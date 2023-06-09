// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PhotonPersistentDiskVolumeSource Represents a Photon Controller persistent disk resource.
//
// swagger:model PhotonPersistentDiskVolumeSource
type PhotonPersistentDiskVolumeSource struct {

	// fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FSType string `json:"fsType,omitempty"`

	// pdID is the ID that identifies Photon Controller persistent disk
	// Required: true
	PdID *string `json:"pdID"`
}
