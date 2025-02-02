// Code generated by MockGen. DO NOT EDIT.
// Source: url.go
//
// Generated by this command:
//
//	mockgen -source=url.go -destination=../../../mocks/url.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	domain "url-shortener/internal/core/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockURLCache is a mock of URLCache interface.
type MockURLCache struct {
	ctrl     *gomock.Controller
	recorder *MockURLCacheMockRecorder
}

// MockURLCacheMockRecorder is the mock recorder for MockURLCache.
type MockURLCacheMockRecorder struct {
	mock *MockURLCache
}

// NewMockURLCache creates a new mock instance.
func NewMockURLCache(ctrl *gomock.Controller) *MockURLCache {
	mock := &MockURLCache{ctrl: ctrl}
	mock.recorder = &MockURLCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLCache) EXPECT() *MockURLCacheMockRecorder {
	return m.recorder
}

// Clean mocks base method.
func (m *MockURLCache) Clean(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clean", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clean indicates an expected call of Clean.
func (mr *MockURLCacheMockRecorder) Clean(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockURLCache)(nil).Clean), ctx, key)
}

// Get mocks base method.
func (m *MockURLCache) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockURLCacheMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockURLCache)(nil).Get), ctx, key)
}

// Ping mocks base method.
func (m *MockURLCache) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockURLCacheMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockURLCache)(nil).Ping), ctx)
}

// Set mocks base method.
func (m *MockURLCache) Set(ctx context.Context, key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockURLCacheMockRecorder) Set(ctx, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockURLCache)(nil).Set), ctx, key, value)
}

// Shutdown mocks base method.
func (m *MockURLCache) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockURLCacheMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockURLCache)(nil).Shutdown))
}

// MockURLRepository is a mock of URLRepository interface.
type MockURLRepository struct {
	ctrl     *gomock.Controller
	recorder *MockURLRepositoryMockRecorder
}

// MockURLRepositoryMockRecorder is the mock recorder for MockURLRepository.
type MockURLRepositoryMockRecorder struct {
	mock *MockURLRepository
}

// NewMockURLRepository creates a new mock instance.
func NewMockURLRepository(ctrl *gomock.Controller) *MockURLRepository {
	mock := &MockURLRepository{ctrl: ctrl}
	mock.recorder = &MockURLRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLRepository) EXPECT() *MockURLRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockURLRepository) Find(shortID string) (*domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", shortID)
	ret0, _ := ret[0].(*domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockURLRepositoryMockRecorder) Find(shortID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockURLRepository)(nil).Find), shortID)
}

// Init mocks base method.
func (m *MockURLRepository) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockURLRepositoryMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockURLRepository)(nil).Init))
}

// Save mocks base method.
func (m *MockURLRepository) Save(URL domain.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", URL)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockURLRepositoryMockRecorder) Save(URL any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockURLRepository)(nil).Save), URL)
}

// Update mocks base method.
func (m *MockURLRepository) Update(URL domain.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", URL)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockURLRepositoryMockRecorder) Update(URL any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockURLRepository)(nil).Update), URL)
}

// MockURLShortenerUseCase is a mock of URLShortenerUseCase interface.
type MockURLShortenerUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockURLShortenerUseCaseMockRecorder
}

// MockURLShortenerUseCaseMockRecorder is the mock recorder for MockURLShortenerUseCase.
type MockURLShortenerUseCaseMockRecorder struct {
	mock *MockURLShortenerUseCase
}

// NewMockURLShortenerUseCase creates a new mock instance.
func NewMockURLShortenerUseCase(ctrl *gomock.Controller) *MockURLShortenerUseCase {
	mock := &MockURLShortenerUseCase{ctrl: ctrl}
	mock.recorder = &MockURLShortenerUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLShortenerUseCase) EXPECT() *MockURLShortenerUseCaseMockRecorder {
	return m.recorder
}

// CreateShortURL mocks base method.
func (m *MockURLShortenerUseCase) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortURL", ctx, originalURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShortURL indicates an expected call of CreateShortURL.
func (mr *MockURLShortenerUseCaseMockRecorder) CreateShortURL(ctx, originalURL any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortURL", reflect.TypeOf((*MockURLShortenerUseCase)(nil).CreateShortURL), ctx, originalURL)
}

// DetailURL mocks base method.
func (m *MockURLShortenerUseCase) DetailURL(ctx context.Context, shortID string) (*domain.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetailURL", ctx, shortID)
	ret0, _ := ret[0].(*domain.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetailURL indicates an expected call of DetailURL.
func (mr *MockURLShortenerUseCaseMockRecorder) DetailURL(ctx, shortID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetailURL", reflect.TypeOf((*MockURLShortenerUseCase)(nil).DetailURL), ctx, shortID)
}

// ResolveURL mocks base method.
func (m *MockURLShortenerUseCase) ResolveURL(ctx context.Context, shortID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveURL", ctx, shortID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveURL indicates an expected call of ResolveURL.
func (mr *MockURLShortenerUseCaseMockRecorder) ResolveURL(ctx, shortID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveURL", reflect.TypeOf((*MockURLShortenerUseCase)(nil).ResolveURL), ctx, shortID)
}

// ToggleURLStatus mocks base method.
func (m *MockURLShortenerUseCase) ToggleURLStatus(ctx context.Context, shortID string, enable bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleURLStatus", ctx, shortID, enable)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleURLStatus indicates an expected call of ToggleURLStatus.
func (mr *MockURLShortenerUseCaseMockRecorder) ToggleURLStatus(ctx, shortID, enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleURLStatus", reflect.TypeOf((*MockURLShortenerUseCase)(nil).ToggleURLStatus), ctx, shortID, enable)
}

// UpdateURL mocks base method.
func (m *MockURLShortenerUseCase) UpdateURL(ctx context.Context, shortID, newOriginalURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateURL", ctx, shortID, newOriginalURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateURL indicates an expected call of UpdateURL.
func (mr *MockURLShortenerUseCaseMockRecorder) UpdateURL(ctx, shortID, newOriginalURL any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateURL", reflect.TypeOf((*MockURLShortenerUseCase)(nil).UpdateURL), ctx, shortID, newOriginalURL)
}
