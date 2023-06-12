// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StackStatus Status of the stack.
//
// swagger:model StackStatus
type StackStatus string

func NewStackStatus(value StackStatus) *StackStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated StackStatus.
func (m StackStatus) Pointer() *StackStatus {
	return &m
}

const (

	// StackStatusREQUESTED captures enum value "REQUESTED"
	StackStatusREQUESTED StackStatus = "REQUESTED"

	// StackStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackStatusCREATEINPROGRESS StackStatus = "CREATE_IN_PROGRESS"

	// StackStatusAVAILABLE captures enum value "AVAILABLE"
	StackStatusAVAILABLE StackStatus = "AVAILABLE"

	// StackStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackStatusUPDATEINPROGRESS StackStatus = "UPDATE_IN_PROGRESS"

	// StackStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackStatusUPDATEREQUESTED StackStatus = "UPDATE_REQUESTED"

	// StackStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackStatusUPDATEFAILED StackStatus = "UPDATE_FAILED"

	// StackStatusBACKUPINPROGRESS captures enum value "BACKUP_IN_PROGRESS"
	StackStatusBACKUPINPROGRESS StackStatus = "BACKUP_IN_PROGRESS"

	// StackStatusBACKUPFAILED captures enum value "BACKUP_FAILED"
	StackStatusBACKUPFAILED StackStatus = "BACKUP_FAILED"

	// StackStatusBACKUPFINISHED captures enum value "BACKUP_FINISHED"
	StackStatusBACKUPFINISHED StackStatus = "BACKUP_FINISHED"

	// StackStatusRESTOREINPROGRESS captures enum value "RESTORE_IN_PROGRESS"
	StackStatusRESTOREINPROGRESS StackStatus = "RESTORE_IN_PROGRESS"

	// StackStatusRESTOREFAILED captures enum value "RESTORE_FAILED"
	StackStatusRESTOREFAILED StackStatus = "RESTORE_FAILED"

	// StackStatusRESTOREFINISHED captures enum value "RESTORE_FINISHED"
	StackStatusRESTOREFINISHED StackStatus = "RESTORE_FINISHED"

	// StackStatusRECOVERYINPROGRESS captures enum value "RECOVERY_IN_PROGRESS"
	StackStatusRECOVERYINPROGRESS StackStatus = "RECOVERY_IN_PROGRESS"

	// StackStatusRECOVERYREQUESTED captures enum value "RECOVERY_REQUESTED"
	StackStatusRECOVERYREQUESTED StackStatus = "RECOVERY_REQUESTED"

	// StackStatusRECOVERYFAILED captures enum value "RECOVERY_FAILED"
	StackStatusRECOVERYFAILED StackStatus = "RECOVERY_FAILED"

	// StackStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackStatusCREATEFAILED StackStatus = "CREATE_FAILED"

	// StackStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackStatusENABLESECURITYFAILED StackStatus = "ENABLE_SECURITY_FAILED"

	// StackStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackStatusPREDELETEINPROGRESS StackStatus = "PRE_DELETE_IN_PROGRESS"

	// StackStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackStatusDELETEINPROGRESS StackStatus = "DELETE_IN_PROGRESS"

	// StackStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackStatusDELETEFAILED StackStatus = "DELETE_FAILED"

	// StackStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	StackStatusDELETEDONPROVIDERSIDE StackStatus = "DELETED_ON_PROVIDER_SIDE"

	// StackStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackStatusDELETECOMPLETED StackStatus = "DELETE_COMPLETED"

	// StackStatusSTOPPED captures enum value "STOPPED"
	StackStatusSTOPPED StackStatus = "STOPPED"

	// StackStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackStatusSTOPREQUESTED StackStatus = "STOP_REQUESTED"

	// StackStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackStatusSTARTREQUESTED StackStatus = "START_REQUESTED"

	// StackStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackStatusSTOPINPROGRESS StackStatus = "STOP_IN_PROGRESS"

	// StackStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackStatusSTARTINPROGRESS StackStatus = "START_IN_PROGRESS"

	// StackStatusSTARTFAILED captures enum value "START_FAILED"
	StackStatusSTARTFAILED StackStatus = "START_FAILED"

	// StackStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackStatusSTOPFAILED StackStatus = "STOP_FAILED"

	// StackStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackStatusWAITFORSYNC StackStatus = "WAIT_FOR_SYNC"

	// StackStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	StackStatusMAINTENANCEMODEENABLED StackStatus = "MAINTENANCE_MODE_ENABLED"

	// StackStatusAMBIGUOUS captures enum value "AMBIGUOUS"
	StackStatusAMBIGUOUS StackStatus = "AMBIGUOUS"

	// StackStatusUNREACHABLE captures enum value "UNREACHABLE"
	StackStatusUNREACHABLE StackStatus = "UNREACHABLE"

	// StackStatusNODEFAILURE captures enum value "NODE_FAILURE"
	StackStatusNODEFAILURE StackStatus = "NODE_FAILURE"

	// StackStatusEXTERNALDATABASECREATIONINPROGRESS captures enum value "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"
	StackStatusEXTERNALDATABASECREATIONINPROGRESS StackStatus = "EXTERNAL_DATABASE_CREATION_IN_PROGRESS"

	// StackStatusEXTERNALDATABASECREATIONFAILED captures enum value "EXTERNAL_DATABASE_CREATION_FAILED"
	StackStatusEXTERNALDATABASECREATIONFAILED StackStatus = "EXTERNAL_DATABASE_CREATION_FAILED"

	// StackStatusEXTERNALDATABASEDELETIONINPROGRESS captures enum value "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"
	StackStatusEXTERNALDATABASEDELETIONINPROGRESS StackStatus = "EXTERNAL_DATABASE_DELETION_IN_PROGRESS"

	// StackStatusEXTERNALDATABASEDELETIONFINISHED captures enum value "EXTERNAL_DATABASE_DELETION_FINISHED"
	StackStatusEXTERNALDATABASEDELETIONFINISHED StackStatus = "EXTERNAL_DATABASE_DELETION_FINISHED"

	// StackStatusEXTERNALDATABASEDELETIONFAILED captures enum value "EXTERNAL_DATABASE_DELETION_FAILED"
	StackStatusEXTERNALDATABASEDELETIONFAILED StackStatus = "EXTERNAL_DATABASE_DELETION_FAILED"

	// StackStatusEXTERNALDATABASESTARTINPROGRESS captures enum value "EXTERNAL_DATABASE_START_IN_PROGRESS"
	StackStatusEXTERNALDATABASESTARTINPROGRESS StackStatus = "EXTERNAL_DATABASE_START_IN_PROGRESS"

	// StackStatusEXTERNALDATABASESTARTFINISHED captures enum value "EXTERNAL_DATABASE_START_FINISHED"
	StackStatusEXTERNALDATABASESTARTFINISHED StackStatus = "EXTERNAL_DATABASE_START_FINISHED"

	// StackStatusEXTERNALDATABASESTARTFAILED captures enum value "EXTERNAL_DATABASE_START_FAILED"
	StackStatusEXTERNALDATABASESTARTFAILED StackStatus = "EXTERNAL_DATABASE_START_FAILED"

	// StackStatusEXTERNALDATABASESTOPINPROGRESS captures enum value "EXTERNAL_DATABASE_STOP_IN_PROGRESS"
	StackStatusEXTERNALDATABASESTOPINPROGRESS StackStatus = "EXTERNAL_DATABASE_STOP_IN_PROGRESS"

	// StackStatusEXTERNALDATABASESTOPFINISHED captures enum value "EXTERNAL_DATABASE_STOP_FINISHED"
	StackStatusEXTERNALDATABASESTOPFINISHED StackStatus = "EXTERNAL_DATABASE_STOP_FINISHED"

	// StackStatusEXTERNALDATABASESTOPFAILED captures enum value "EXTERNAL_DATABASE_STOP_FAILED"
	StackStatusEXTERNALDATABASESTOPFAILED StackStatus = "EXTERNAL_DATABASE_STOP_FAILED"

	// StackStatusLOADBALANCERUPDATEINPROGRESS captures enum value "LOAD_BALANCER_UPDATE_IN_PROGRESS"
	StackStatusLOADBALANCERUPDATEINPROGRESS StackStatus = "LOAD_BALANCER_UPDATE_IN_PROGRESS"

	// StackStatusLOADBALANCERUPDATEFINISHED captures enum value "LOAD_BALANCER_UPDATE_FINISHED"
	StackStatusLOADBALANCERUPDATEFINISHED StackStatus = "LOAD_BALANCER_UPDATE_FINISHED"

	// StackStatusLOADBALANCERUPDATEFAILED captures enum value "LOAD_BALANCER_UPDATE_FAILED"
	StackStatusLOADBALANCERUPDATEFAILED StackStatus = "LOAD_BALANCER_UPDATE_FAILED"
)

// for schema
var stackStatusEnum []interface{}

func init() {
	var res []StackStatus
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","BACKUP_IN_PROGRESS","BACKUP_FAILED","BACKUP_FINISHED","RESTORE_IN_PROGRESS","RESTORE_FAILED","RESTORE_FINISHED","RECOVERY_IN_PROGRESS","RECOVERY_REQUESTED","RECOVERY_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETED_ON_PROVIDER_SIDE","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","AMBIGUOUS","UNREACHABLE","NODE_FAILURE","EXTERNAL_DATABASE_CREATION_IN_PROGRESS","EXTERNAL_DATABASE_CREATION_FAILED","EXTERNAL_DATABASE_DELETION_IN_PROGRESS","EXTERNAL_DATABASE_DELETION_FINISHED","EXTERNAL_DATABASE_DELETION_FAILED","EXTERNAL_DATABASE_START_IN_PROGRESS","EXTERNAL_DATABASE_START_FINISHED","EXTERNAL_DATABASE_START_FAILED","EXTERNAL_DATABASE_STOP_IN_PROGRESS","EXTERNAL_DATABASE_STOP_FINISHED","EXTERNAL_DATABASE_STOP_FAILED","LOAD_BALANCER_UPDATE_IN_PROGRESS","LOAD_BALANCER_UPDATE_FINISHED","LOAD_BALANCER_UPDATE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackStatusEnum = append(stackStatusEnum, v)
	}
}

func (m StackStatus) validateStackStatusEnum(path, location string, value StackStatus) error {
	if err := validate.EnumCase(path, location, value, stackStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this stack status
func (m StackStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStackStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this stack status based on context it is used
func (m StackStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}