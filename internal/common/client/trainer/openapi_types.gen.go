// Package trainer provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package trainer

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Date defines model for Date.
type Date struct {
	Date         openapi_types.Date `json:"date"`
	HasFreeHours bool               `json:"hasFreeHours"`
	Hours        []Hour             `json:"hours"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// Hour defines model for Hour.
type Hour struct {
	Available            bool      `json:"available"`
	HasTrainingScheduled bool      `json:"hasTrainingScheduled"`
	Hour                 time.Time `json:"hour"`
}

// HourUpdate defines model for HourUpdate.
type HourUpdate struct {
	Hours []time.Time `json:"hours"`
}

// GetTrainerAvailableHoursParams defines parameters for GetTrainerAvailableHours.
type GetTrainerAvailableHoursParams struct {
	DateFrom time.Time `form:"dateFrom" json:"dateFrom"`
	DateTo   time.Time `form:"dateTo" json:"dateTo"`
}

// MakeHourAvailableJSONRequestBody defines body for MakeHourAvailable for application/json ContentType.
type MakeHourAvailableJSONRequestBody = HourUpdate

// MakeHourUnavailableJSONRequestBody defines body for MakeHourUnavailable for application/json ContentType.
type MakeHourUnavailableJSONRequestBody = HourUpdate
