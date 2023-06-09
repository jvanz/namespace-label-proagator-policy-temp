// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PodSchedulingGate PodSchedulingGate is associated to a Pod to guard its scheduling.
//
// swagger:model PodSchedulingGate
type PodSchedulingGate struct {

	// Name of the scheduling gate. Each scheduling gate must have a unique name field.
	// Required: true
	Name *string `json:"name"`
}
