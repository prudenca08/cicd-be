// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	patientses "finalproject/features/patientses"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AllPatientses provides a mock function with given fields:
func (_m *Repository) AllPatientses() ([]patientses.Domain, error) {
	ret := _m.Called()

	var r0 []patientses.Domain
	if rf, ok := ret.Get(0).(func() []patientses.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patientses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: pssID, domain
func (_m *Repository) Create(pssID int, domain *patientses.Domain) (patientses.Domain, error) {
	ret := _m.Called(pssID, domain)

	var r0 patientses.Domain
	if rf, ok := ret.Get(0).(func(int, *patientses.Domain) patientses.Domain); ok {
		r0 = rf(pssID, domain)
	} else {
		r0 = ret.Get(0).(patientses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *patientses.Domain) error); ok {
		r1 = rf(pssID, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: pssID, id
func (_m *Repository) Delete(pssID int, id int) (string, error) {
	ret := _m.Called(pssID, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(pssID, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(pssID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatientsesByID provides a mock function with given fields: id
func (_m *Repository) PatientsesByID(id int) (patientses.Domain, error) {
	ret := _m.Called(id)

	var r0 patientses.Domain
	if rf, ok := ret.Get(0).(func(int) patientses.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(patientses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: admID, pssID, domain
func (_m *Repository) Update(admID int, pssID int, domain *patientses.Domain) (patientses.Domain, error) {
	ret := _m.Called(admID, pssID, domain)

	var r0 patientses.Domain
	if rf, ok := ret.Get(0).(func(int, int, *patientses.Domain) patientses.Domain); ok {
		r0 = rf(admID, pssID, domain)
	} else {
		r0 = ret.Get(0).(patientses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, *patientses.Domain) error); ok {
		r1 = rf(admID, pssID, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
