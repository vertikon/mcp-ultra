package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vertikon/mcp-ultra/internal/domain"
	"github.com/vertikon/mcp-ultra/internal/services"
	"go.uber.org/zap"
)

// MockHealthService for testing
type MockHealthService struct {
	mock.Mock
}

func (m *MockHealthService) Check(ctx context.Context) map[string]services.HealthStatus {
	args := m.Called(ctx)
	return args.Get(0).(map[string]services.HealthStatus)
}

func (m *MockHealthService) IsReady(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

func (m *MockHealthService) IsLive(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

func (m *MockHealthService) RegisterChecker(name string, checker services.HealthChecker) {
	m.Called(name, checker)
}

// MockTaskService for testing
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(ctx context.Context, req domain.CreateTaskRequest) (*domain.Task, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskService) GetTask(ctx context.Context, taskID string) (*domain.Task, error) {
	args := m.Called(ctx, taskID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(ctx context.Context, taskID string, req domain.UpdateTaskRequest) (*domain.Task, error) {
	args := m.Called(ctx, taskID, req)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskService) DeleteTask(ctx context.Context, taskID string) error {
	args := m.Called(ctx, taskID)
	return args.Error(0)
}

func (m *MockTaskService) ListTasks(ctx context.Context, filters domain.TaskFilters) (*domain.TaskList, error) {
	args := m.Called(ctx, filters)
	return args.Get(0).(*domain.TaskList), args.Error(1)
}

func TestNewRouter(t *testing.T) {
	logger := zap.NewNop()
	mockHealthService := &MockHealthService{}
	mockTaskService := &MockTaskService{}

	router := NewRouter(logger, mockHealthService, mockTaskService)

	assert.NotNil(t, router)
}

func TestRouter_HealthEndpoints(t *testing.T) {
	logger := zap.NewNop()
	mockHealthService := &MockHealthService{}
	mockTaskService := &MockTaskService{}

	tests := []struct {
		name           string
		endpoint       string
		setupMock      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "health endpoint returns healthy status",
			endpoint: "/health",
			setupMock: func() {
				mockHealthService.On("Check", mock.Anything).Return(map[string]services.HealthStatus{
					"database": {Status: "healthy", Message: "Connected"},
					"redis":    {Status: "healthy", Message: "Connected"},
				})
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:     "ready endpoint returns true",
			endpoint: "/ready",
			setupMock: func() {
				mockHealthService.On("IsReady", mock.Anything).Return(true)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"ready":true}`,
		},
		{
			name:     "ready endpoint returns false",
			endpoint: "/ready",
			setupMock: func() {
				mockHealthService.On("IsReady", mock.Anything).Return(false)
			},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   `{"ready":false}`,
		},
		{
			name:     "live endpoint returns true",
			endpoint: "/live",
			setupMock: func() {
				mockHealthService.On("IsLive", mock.Anything).Return(true)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"live":true}`,
		},
		{
			name:     "live endpoint returns false",
			endpoint: "/live",
			setupMock: func() {
				mockHealthService.On("IsLive", mock.Anything).Return(false)
			},
			expectedStatus: http.StatusServiceUnavailable,
			expectedBody:   `{"live":false}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset mocks
			mockHealthService.ExpectedCalls = nil
			
			tt.setupMock()

			router := NewRouter(logger, mockHealthService, mockTaskService)
			req := httptest.NewRequest(http.MethodGet, tt.endpoint, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}

			mockHealthService.AssertExpectations(t)
		})
	}
}

func TestRouter_TaskEndpoints(t *testing.T) {
	logger := zap.NewNop()
	mockHealthService := &MockHealthService{}
	mockTaskService := &MockTaskService{}

	router := NewRouter(logger, mockHealthService, mockTaskService)

	t.Run("POST /tasks - create task", func(t *testing.T) {
		taskRequest := domain.CreateTaskRequest{
			Title:       "Test Task",
			Description: "Test Description",
			Priority:    "high",
			Category:    "development",
		}

		expectedTask := &domain.Task{
			ID:          "task-123",
			Title:       "Test Task",
			Description: "Test Description",
			Priority:    "high",
			Category:    "development",
			Status:      "pending",
		}

		mockTaskService.On("CreateTask", mock.Anything, taskRequest).Return(expectedTask, nil)

		body, _ := json.Marshal(taskRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		
		var response domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.ID, response.ID)
		assert.Equal(t, expectedTask.Title, response.Title)

		mockTaskService.AssertExpectations(t)
	})

	t.Run("GET /tasks/:id - get task", func(t *testing.T) {
		taskID := "task-123"
		expectedTask := &domain.Task{
			ID:          taskID,
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "pending",
		}

		mockTaskService.On("GetTask", mock.Anything, taskID).Return(expectedTask, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks/"+taskID, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.ID, response.ID)

		mockTaskService.AssertExpectations(t)
	})

	t.Run("PUT /tasks/:id - update task", func(t *testing.T) {
		taskID := "task-123"
		updateRequest := domain.UpdateTaskRequest{
			Title:  "Updated Task",
			Status: "completed",
		}

		expectedTask := &domain.Task{
			ID:     taskID,
			Title:  "Updated Task",
			Status: "completed",
		}

		mockTaskService.On("UpdateTask", mock.Anything, taskID, updateRequest).Return(expectedTask, nil)

		body, _ := json.Marshal(updateRequest)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/tasks/"+taskID, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response domain.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.Title, response.Title)

		mockTaskService.AssertExpectations(t)
	})

	t.Run("DELETE /tasks/:id - delete task", func(t *testing.T) {
		taskID := "task-123"

		mockTaskService.On("DeleteTask", mock.Anything, taskID).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/api/v1/tasks/"+taskID, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)

		mockTaskService.AssertExpectations(t)
	})

	t.Run("GET /tasks - list tasks", func(t *testing.T) {
		filters := domain.TaskFilters{
			Status:   "pending",
			Priority: "high",
		}

		taskList := &domain.TaskList{
			Tasks: []domain.Task{
				{ID: "task-1", Title: "Task 1", Status: "pending", Priority: "high"},
				{ID: "task-2", Title: "Task 2", Status: "pending", Priority: "high"},
			},
			Total: 2,
			Page:  1,
			Size:  10,
		}

		mockTaskService.On("ListTasks", mock.Anything, mock.MatchedBy(func(f domain.TaskFilters) bool {
			return f.Status == "pending" && f.Priority == "high"
		})).Return(taskList, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks?status=pending&priority=high", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response domain.TaskList
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(response.Tasks))
		assert.Equal(t, 2, response.Total)

		mockTaskService.AssertExpectations(t)
	})
}

func TestRouter_Middleware(t *testing.T) {
	logger := zap.NewNop()
	mockHealthService := &MockHealthService{}
	mockTaskService := &MockTaskService{}

	router := NewRouter(logger, mockHealthService, mockTaskService)

	t.Run("CORS headers are set", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/api/v1/tasks", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		req.Header.Set("Access-Control-Request-Method", "POST")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Origin"), "*")
		assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "POST")
	})

	t.Run("Content-Type header is set for JSON responses", func(t *testing.T) {
		mockHealthService.On("IsReady", mock.Anything).Return(true)

		req := httptest.NewRequest(http.MethodGet, "/ready", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		mockHealthService.AssertExpectations(t)
	})
}

func TestRouter_ErrorHandling(t *testing.T) {
	logger := zap.NewNop()
	mockHealthService := &MockHealthService{}
	mockTaskService := &MockTaskService{}

	router := NewRouter(logger, mockHealthService, mockTaskService)

	t.Run("404 for non-existent endpoint", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/non-existent", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("405 for unsupported method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/health", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})

	t.Run("400 for invalid JSON in POST request", func(t *testing.T) {
		invalidJSON := `{"title": "Test", "invalid": json}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}