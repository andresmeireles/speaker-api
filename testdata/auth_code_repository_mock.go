package testdata

import (
	"github.com/andresmeireles/speaker/internal/codesender"
	"github.com/stretchr/testify/mock"
)

type AuthCodeRepositoryMock struct {
	mock.Mock
}

func (r *AuthCodeRepositoryMock) Add(authCode codesender.AuthCode) error {
	return r.Called(authCode).Error(0)
}

func (r *AuthCodeRepositoryMock) GetById(authCodeId int) (codesender.AuthCode, error) {
	args := r.Called(authCodeId)
	return args.Get(0).(codesender.AuthCode), args.Error(1)
}

func (r *AuthCodeRepositoryMock) GetAll() ([]codesender.AuthCode, error) {
	args := r.Called()
	return args.Get(0).([]codesender.AuthCode), args.Error(1)
}

func (r *AuthCodeRepositoryMock) GetByCode(code string) (codesender.AuthCode, error) {
	args := r.Called(code)
	return args.Get(0).(codesender.AuthCode), args.Error(1)
}

func (r *AuthCodeRepositoryMock) Update(authCode codesender.AuthCode) error {
	return r.Called(authCode).Error(0)
}

func (r *AuthCodeRepositoryMock) Delete(authCode codesender.AuthCode) error {
	return r.Called(authCode).Error(0)
}
