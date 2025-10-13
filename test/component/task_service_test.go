package component

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"{{MODULE_PATH}}/internal/domain"
	"{{MODULE_PATH}}/internal/services"
	"{{MODULE_PATH}}/test/mocks"
)

// TaskServiceTestSuite provides isolated testing for TaskService
type TaskServiceTestSuite struct {
	suite.Suite
	service   *services.TaskService
	taskRepo  *mocks.MockTaskRepository
	cacheRepo *mocks.MockCacheRepository
	eventBus  *mocks.MockEventBus
	validator *mocks.MockValidator
	logger    *zap.Logger
}

func (suite *TaskServiceTestSuite) SetupTest() {
	suite.taskRepo = &mocks.MockTaskRepository{}
	suite.cacheRepo = &mocks.MockCacheRepository{}
	suite.eventBus = &mocks.MockEventBus{}
	suite.validator = &mocks.MockValidator{}
	suite.logger = zap.NewNop()

	suite.service = services.NewTaskService(
		suite.taskRepo,
		suite.validator,
		nil, // audit service not needed for these tests
		suite.cacheRepo,
		suite.logger,
		suite.eventBus,
	)
}

func (suite *TaskServiceTestSuite) TearDownTest() {
	suite.taskRepo.AssertExpectations(suite.T())
	suite.cacheRepo.AssertExpectations(suite.T())
	suite.eventBus.AssertExpectations(suite.T())
	suite.validator.AssertExpectations(suite.T())
}

// Test Create Task - Happy Path
func (suite *TaskServiceTestSuite) TestCreateTask_Success() {
	ctx := context.Background()
	userID := uuid.New()

	req := &services.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    domain.PriorityHigh,
		Tags:        []string{"test", "component"},
		Metadata:    map[string]interface{}{"source": "test"},
	}

	expectedTask := &domain.Task{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		Status:      domain.TaskStatusPending,
		Priority:    req.Priority,
		Tags:        req.Tags,
		CreatedBy:   userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Metadata:    req.Metadata,
	}

	// Setup mocks
	suite.validator.On("ValidateStruct", mock.AnythingOfType("*services.CreateTaskRequest")).Return(nil)
	suite.taskRepo.On("Create", ctx, mock.MatchedBy(func(task *domain.Task) bool {
		return task.Title == req.Title &&
			task.Description == req.Description &&
			task.Priority == req.Priority &&
			task.CreatedBy == userID
	})).Return(expectedTask, nil)

	suite.cacheRepo.On("Delete", ctx, mock.MatchedBy(func(key string) bool {
		return key == "tasks:user:"+userID.String()
	})).Return(nil)

	suite.eventBus.On("Publish", ctx, "task.created", mock.AnythingOfType("*events.TaskCreatedEvent")).Return(nil)

	// Execute
	result, err := suite.service.CreateTask(ctx, userID, req)

	// Assert
	require.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), req.Title, result.Title)
	assert.Equal(suite.T(), req.Description, result.Description)
	assert.Equal(suite.T(), req.Priority, result.Priority)
	assert.Equal(suite.T(), userID, result.CreatedBy)
}

// Test Create Task - Validation Error
func (suite *TaskServiceTestSuite) TestCreateTask_ValidationError() {
	ctx := context.Background()
	userID := uuid.New()

	req := &services.CreateTaskRequest{
		Title:       "", // Invalid empty title
		Description: "Test Description",
	}

	validationErr := &services.ValidationError{
		Field:   "title",
		Message: "title is required",
	}

	// Setup mocks
	suite.validator.On("ValidateStruct", req).Return(validationErr)

	// Execute
	result, err := suite.service.CreateTask(ctx, userID, req)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "title is required")
}

// Test Get Task - Cache Hit
func (suite *TaskServiceTestSuite) TestGetTask_CacheHit() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()

	cachedTask := &domain.Task{
		ID:        taskID,
		Title:     "Cached Task",
		CreatedBy: userID,
	}

	// Setup mocks - cache hit, no database call
	suite.cacheRepo.On("Get", ctx, "task:"+taskID.String()).Return(cachedTask, nil)

	// Execute
	result, err := suite.service.GetTask(ctx, taskID, userID)

	// Assert
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), cachedTask, result)
}

// Test Get Task - Cache Miss, Database Hit
func (suite *TaskServiceTestSuite) TestGetTask_CacheMissDbHit() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()

	dbTask := &domain.Task{
		ID:        taskID,
		Title:     "DB Task",
		CreatedBy: userID,
	}

	// Setup mocks - cache miss, database hit, cache update
	suite.cacheRepo.On("Get", ctx, "task:"+taskID.String()).Return((*domain.Task)(nil), services.ErrNotFound)
	suite.taskRepo.On("GetByID", ctx, taskID).Return(dbTask, nil)
	suite.cacheRepo.On("Set", ctx, "task:"+taskID.String(), dbTask, 300).Return(nil)

	// Execute
	result, err := suite.service.GetTask(ctx, taskID, userID)

	// Assert
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), dbTask, result)
}

// Test Get Task - Not Found
func (suite *TaskServiceTestSuite) TestGetTask_NotFound() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()

	// Setup mocks - cache miss, database miss
	suite.cacheRepo.On("Get", ctx, "task:"+taskID.String()).Return((*domain.Task)(nil), services.ErrNotFound)
	suite.taskRepo.On("GetByID", ctx, taskID).Return((*domain.Task)(nil), services.ErrNotFound)

	// Execute
	result, err := suite.service.GetTask(ctx, taskID, userID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), services.ErrNotFound, err)
}

// Test Get Task - Access Denied (Different User)
func (suite *TaskServiceTestSuite) TestGetTask_AccessDenied() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()
	differentUserID := uuid.New()

	task := &domain.Task{
		ID:        taskID,
		Title:     "Other User's Task",
		CreatedBy: differentUserID, // Different user
	}

	// Setup mocks
	suite.cacheRepo.On("Get", ctx, "task:"+taskID.String()).Return((*domain.Task)(nil), services.ErrNotFound)
	suite.taskRepo.On("GetByID", ctx, taskID).Return(task, nil)

	// Execute
	result, err := suite.service.GetTask(ctx, taskID, userID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), services.ErrAccessDenied, err)
}

// Test Update Task - Success
func (suite *TaskServiceTestSuite) TestUpdateTask_Success() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()

	existingTask := &domain.Task{
		ID:        taskID,
		Title:     "Old Title",
		CreatedBy: userID,
		Status:    domain.TaskStatusPending,
	}

	req := &services.UpdateTaskRequest{
		Title:       "New Title",
		Description: "New Description",
		Priority:    domain.PriorityUrgent,
		Tags:        []string{"updated"},
	}

	updatedTask := &domain.Task{
		ID:          taskID,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Tags:        req.Tags,
		CreatedBy:   userID,
		UpdatedAt:   time.Now(),
	}

	// Setup mocks
	suite.validator.On("ValidateStruct", req).Return(nil)
	suite.taskRepo.On("GetByID", ctx, taskID).Return(existingTask, nil)
	suite.taskRepo.On("Update", ctx, mock.MatchedBy(func(task *domain.Task) bool {
		return task.ID == taskID && task.Title == req.Title
	})).Return(updatedTask, nil)

	suite.cacheRepo.On("Delete", ctx, "task:"+taskID.String()).Return(nil)
	suite.cacheRepo.On("Delete", ctx, "tasks:user:"+userID.String()).Return(nil)

	suite.eventBus.On("Publish", ctx, "task.updated", mock.AnythingOfType("*events.TaskUpdatedEvent")).Return(nil)

	// Execute
	result, err := suite.service.UpdateTask(ctx, taskID, userID, req)

	// Assert
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), req.Title, result.Title)
	assert.Equal(suite.T(), req.Description, result.Description)
	assert.Equal(suite.T(), req.Priority, result.Priority)
}

// Test Complete Task - Success
func (suite *TaskServiceTestSuite) TestCompleteTask_Success() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()

	task := &domain.Task{
		ID:        taskID,
		Title:     "Task to Complete",
		CreatedBy: userID,
		Status:    domain.TaskStatusInProgress,
	}

	completedTask := &domain.Task{
		ID:          taskID,
		Title:       task.Title,
		CreatedBy:   userID,
		Status:      domain.TaskStatusCompleted,
		CompletedAt: &time.Time{},
	}

	// Setup mocks
	suite.taskRepo.On("GetByID", ctx, taskID).Return(task, nil)
	suite.taskRepo.On("Update", ctx, mock.MatchedBy(func(t *domain.Task) bool {
		return t.ID == taskID && t.Status == domain.TaskStatusCompleted
	})).Return(completedTask, nil)

	suite.cacheRepo.On("Delete", ctx, "task:"+taskID.String()).Return(nil)
	suite.cacheRepo.On("Delete", ctx, "tasks:user:"+userID.String()).Return(nil)

	suite.eventBus.On("Publish", ctx, "task.completed", mock.AnythingOfType("*events.TaskCompletedEvent")).Return(nil)

	// Execute
	result, err := suite.service.CompleteTask(ctx, taskID, userID)

	// Assert
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), domain.TaskStatusCompleted, result.Status)
	assert.NotNil(suite.T(), result.CompletedAt)
}

// Test Complete Task - Invalid Status Transition
func (suite *TaskServiceTestSuite) TestCompleteTask_InvalidStatusTransition() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()

	task := &domain.Task{
		ID:        taskID,
		Title:     "Already Completed Task",
		CreatedBy: userID,
		Status:    domain.TaskStatusCompleted, // Already completed
	}

	// Setup mocks
	suite.taskRepo.On("GetByID", ctx, taskID).Return(task, nil)

	// Execute
	result, err := suite.service.CompleteTask(ctx, taskID, userID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "invalid status transition")
}

// Test List Tasks - With Pagination and Filters
func (suite *TaskServiceTestSuite) TestListTasks_WithFilters() {
	ctx := context.Background()
	userID := uuid.New()

	req := &services.ListTasksRequest{
		UserID: userID,
		Page:   1,
		Limit:  10,
		Status: domain.TaskStatusPending,
		Tags:   []string{"important"},
	}

	expectedTasks := []*domain.Task{
		{
			ID:        uuid.New(),
			Title:     "Important Task 1",
			CreatedBy: userID,
			Status:    domain.TaskStatusPending,
			Tags:      []string{"important", "urgent"},
		},
		{
			ID:        uuid.New(),
			Title:     "Important Task 2",
			CreatedBy: userID,
			Status:    domain.TaskStatusPending,
			Tags:      []string{"important"},
		},
	}

	totalCount := int64(2)

	// Setup mocks
	cacheKey := "tasks:user:" + userID.String() + ":page:1:limit:10:status:pending:tags:important"
	suite.cacheRepo.On("Get", ctx, cacheKey).Return(nil, services.ErrNotFound)

	suite.taskRepo.On("List", ctx, mock.MatchedBy(func(filter *domain.TaskFilter) bool {
		return filter.UserID == userID &&
			filter.Status == domain.TaskStatusPending &&
			len(filter.Tags) == 1 && filter.Tags[0] == "important"
	})).Return(expectedTasks, totalCount, nil)

	suite.cacheRepo.On("Set", ctx, cacheKey, mock.Anything, 60).Return(nil)

	// Execute
	result, total, err := suite.service.ListTasks(ctx, req)

	// Assert
	require.NoError(suite.T(), err)
	assert.Len(suite.T(), result, 2)
	assert.Equal(suite.T(), totalCount, total)
	assert.Equal(suite.T(), expectedTasks[0].Title, result[0].Title)
	assert.Equal(suite.T(), expectedTasks[1].Title, result[1].Title)
}

// Test Delete Task - Success (Admin User)
func (suite *TaskServiceTestSuite) TestDeleteTask_AdminSuccess() {
	ctx := context.Background()
	taskID := uuid.New()
	adminUserID := uuid.New()
	taskOwnerID := uuid.New()

	task := &domain.Task{
		ID:        taskID,
		Title:     "Task to Delete",
		CreatedBy: taskOwnerID,
	}

	// Mock admin context
	ctx = context.WithValue(ctx, "user_role", "admin")

	// Setup mocks
	suite.taskRepo.On("GetByID", ctx, taskID).Return(task, nil)
	suite.taskRepo.On("Delete", ctx, taskID).Return(nil)

	suite.cacheRepo.On("Delete", ctx, "task:"+taskID.String()).Return(nil)
	suite.cacheRepo.On("Delete", ctx, "tasks:user:"+taskOwnerID.String()).Return(nil)

	suite.eventBus.On("Publish", ctx, "task.deleted", mock.AnythingOfType("*events.TaskDeletedEvent")).Return(nil)

	// Execute
	err := suite.service.DeleteTask(ctx, taskID, adminUserID)

	// Assert
	require.NoError(suite.T(), err)
}

// Test Delete Task - Access Denied (Non-admin User)
func (suite *TaskServiceTestSuite) TestDeleteTask_AccessDenied() {
	ctx := context.Background()
	taskID := uuid.New()
	userID := uuid.New()
	taskOwnerID := uuid.New()

	task := &domain.Task{
		ID:        taskID,
		Title:     "Someone Else's Task",
		CreatedBy: taskOwnerID,
	}

	// Setup mocks
	suite.taskRepo.On("GetByID", ctx, taskID).Return(task, nil)

	// Execute
	err := suite.service.DeleteTask(ctx, taskID, userID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), services.ErrAccessDenied, err)
}

// Test Concurrent Operations
func (suite *TaskServiceTestSuite) TestConcurrentOperations() {
	// This test would be more complex and test concurrent access patterns
	// For brevity, we'll skip the full implementation but it would test:
	// - Concurrent task updates
	// - Cache consistency under load
	// - Event ordering
	suite.T().Skip("Concurrent operations test - implement if needed")
}

// Run the test suite
func TestTaskServiceSuite(t *testing.T) {
	suite.Run(t, new(TaskServiceTestSuite))
}
