package model

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidID       = errors.New("asphaltm: invalid identifier")
	ErrNotFound        = errors.New("asphaltm: entity not found")
	ErrConflict        = errors.New("asphaltm: state conflict")
	ErrInterlock       = errors.New("asphaltm: interlock denied")
	ErrMoistureHold    = errors.New("asphaltm: moisture hold active")
	ErrAirflowSetpoint = errors.New("asphaltm: airflow setpoint violation")
	ErrFanFault        = errors.New("asphaltm: fan fault")
	ErrScheduleEmpty   = errors.New("asphaltm: schedule empty")
	ErrGradient        = errors.New("asphaltm: moisture gradient violation")
	ErrBinderDrift   = errors.New("asphaltm: moisture drift exceeded")
	ErrBinderTrip    = errors.New("asphaltm: heat overtemperature")
	ErrHeatHold    = errors.New("asphaltm: gradient hold not satisfied")
	ErrContextCanceled = errors.New("asphaltm: operation canceled")
)

type DomainError struct {
	Op   string
	Code string
	Err  error
}

func (e *DomainError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil {
		return fmt.Sprintf("asphaltm %s [%s]: %v", e.Op, e.Code, e.Err)
	}
	return fmt.Sprintf("asphaltm %s [%s]", e.Op, e.Code)
}

func (e *DomainError) Unwrap() error { return e.Err }

func Wrap(op, code string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Op: op, Code: code, Err: err}
}

func Is(err, target error) bool   { return errors.Is(err, target) }
func As(err error, target any) bool { return errors.As(err, target) }
