// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DfGetTask A download process initiated by dfget or other clients.
//
// swagger:model DfGetTask
type DfGetTask struct {

	// CID means the client ID. It maps to the specific dfget process.
	// When user wishes to download an image/file, user would start a dfget process to do this.
	// This dfget is treated a client and carries a client ID.
	// Thus, multiple dfget processes on the same peer have different CIDs.
	//
	CID string `json:"cID,omitempty"`

	// path is used in one peer A for uploading functionality. When peer B hopes
	// to get piece C from peer A, B must provide a URL for piece C.
	// Then when creating a task in supernode, peer A must provide this URL in request.
	//
	Path string `json:"path,omitempty"`

	// PeerID uniquely identifies a peer, and the cID uniquely identifies a
	// download task belonging to a peer. One peer can initiate multiple download tasks,
	// which means that one peer corresponds to multiple cIDs.
	//
	PeerID string `json:"peerID,omitempty"`

	// The size of pieces which is calculated as per the following strategy
	// 1. If file's total size is less than 200MB, then the piece size is 4MB by default.
	// 2. Otherwise, it equals to the smaller value between totalSize/100MB + 2 MB and 15MB.
	//
	PieceSize int32 `json:"pieceSize,omitempty"`

	// The status of Dfget download process.
	//
	// Enum: [WAITING RUNNING FAILED SUCCESS]
	Status string `json:"status,omitempty"`

	// task Id
	TaskID string `json:"taskId,omitempty"`
}

// Validate validates this df get task
func (m *DfGetTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dfGetTaskTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WAITING","RUNNING","FAILED","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dfGetTaskTypeStatusPropEnum = append(dfGetTaskTypeStatusPropEnum, v)
	}
}

const (

	// DfGetTaskStatusWAITING captures enum value "WAITING"
	DfGetTaskStatusWAITING string = "WAITING"

	// DfGetTaskStatusRUNNING captures enum value "RUNNING"
	DfGetTaskStatusRUNNING string = "RUNNING"

	// DfGetTaskStatusFAILED captures enum value "FAILED"
	DfGetTaskStatusFAILED string = "FAILED"

	// DfGetTaskStatusSUCCESS captures enum value "SUCCESS"
	DfGetTaskStatusSUCCESS string = "SUCCESS"
)

// prop value enum
func (m *DfGetTask) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dfGetTaskTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DfGetTask) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DfGetTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DfGetTask) UnmarshalBinary(b []byte) error {
	var res DfGetTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
