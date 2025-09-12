package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"go.uber.org/zap"

	taskv1 "github.com/vertikon/mcp-ultra/api/grpc/gen/task/v1"
	"github.com/vertikon/mcp-ultra/internal/domain"
	"github.com/vertikon/mcp-ultra/internal/services"
)

// TaskServer implements the TaskService gRPC server
type TaskServer struct {
	taskv1.UnimplementedTaskServiceServer
	taskService *services.TaskService
	logger      *zap.Logger
}

// NewTaskServer creates a new TaskServer instance
func NewTaskServer(taskService *services.TaskService, logger *zap.Logger) *TaskServer {
	return &TaskServer{
		taskService: taskService,
		logger:      logger.Named("task-server"),
	}
}

// CreateTask creates a new task
func (s *TaskServer) CreateTask(ctx context.Context, req *taskv1.CreateTaskRequest) (*taskv1.CreateTaskResponse, error) {
	s.logger.Debug("CreateTask called", zap.String("title", req.Task.Title))

	// Convert protobuf task to domain task
	domainTask, err := s.protobufToDomainTask(req.Task)
	if err != nil {
		s.logger.Error("Failed to convert protobuf task to domain task", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid task: %v", err)
	}

	// Create the task
	createdTask, err := s.taskService.CreateTask(ctx, domainTask)
	if err != nil {
		s.logger.Error("Failed to create task", zap.Error(err))
		return nil, s.handleServiceError(err)
	}

	// Convert back to protobuf
	pbTask, err := s.domainToProtobufTask(createdTask)
	if err != nil {
		s.logger.Error("Failed to convert domain task to protobuf", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to serialize task")
	}

	return &taskv1.CreateTaskResponse{
		Task: pbTask,
	}, nil
}

// GetTask retrieves a task by ID
func (s *TaskServer) GetTask(ctx context.Context, req *taskv1.GetTaskRequest) (*taskv1.GetTaskResponse, error) {
	s.logger.Debug("GetTask called", zap.String("id", req.Id))

	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "task ID is required")
	}

	task, err := s.taskService.GetTask(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to get task", zap.String("id", req.Id), zap.Error(err))
		return nil, s.handleServiceError(err)
	}

	pbTask, err := s.domainToProtobufTask(task)
	if err != nil {
		s.logger.Error("Failed to convert domain task to protobuf", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to serialize task")
	}

	// Apply field mask if provided
	if req.FieldMask != nil {
		// In a production implementation, you would apply the field mask here
		// For now, we return the full task
	}

	return &taskv1.GetTaskResponse{
		Task: pbTask,
	}, nil
}

// UpdateTask updates an existing task
func (s *TaskServer) UpdateTask(ctx context.Context, req *taskv1.UpdateTaskRequest) (*taskv1.UpdateTaskResponse, error) {
	s.logger.Debug("UpdateTask called", zap.String("id", req.Task.Id))

	if req.Task.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "task ID is required")
	}

	domainTask, err := s.protobufToDomainTask(req.Task)
	if err != nil {
		s.logger.Error("Failed to convert protobuf task to domain task", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid task: %v", err)
	}

	// Apply update mask if provided
	if req.UpdateMask != nil {
		// In a production implementation, you would apply the update mask here
		// to only update specified fields
	}

	updatedTask, err := s.taskService.UpdateTask(ctx, domainTask)
	if err != nil {
		s.logger.Error("Failed to update task", zap.String("id", req.Task.Id), zap.Error(err))
		return nil, s.handleServiceError(err)
	}

	pbTask, err := s.domainToProtobufTask(updatedTask)
	if err != nil {
		s.logger.Error("Failed to convert domain task to protobuf", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to serialize task")
	}

	return &taskv1.UpdateTaskResponse{
		Task: pbTask,
	}, nil
}

// DeleteTask deletes a task
func (s *TaskServer) DeleteTask(ctx context.Context, req *taskv1.DeleteTaskRequest) (*emptypb.Empty, error) {
	s.logger.Debug("DeleteTask called", zap.String("id", req.Id))

	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "task ID is required")
	}

	err := s.taskService.DeleteTask(ctx, req.Id)
	if err != nil {
		s.logger.Error("Failed to delete task", zap.String("id", req.Id), zap.Error(err))
		return nil, s.handleServiceError(err)
	}

	return &emptypb.Empty{}, nil
}

// ListTasks lists tasks with filtering and pagination
func (s *TaskServer) ListTasks(ctx context.Context, req *taskv1.ListTasksRequest) (*taskv1.ListTasksResponse, error) {
	s.logger.Debug("ListTasks called", zap.Int32("page_size", req.PageSize))

	// Set default page size if not provided
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	// Convert filters
	filter := s.convertTaskFilter(req.Filter)

	// Get tasks from service
	result, err := s.taskService.ListTasks(ctx, services.ListTasksOptions{
		Filter:    filter,
		PageSize:  int(pageSize),
		PageToken: req.PageToken,
	})
	if err != nil {
		s.logger.Error("Failed to list tasks", zap.Error(err))
		return nil, s.handleServiceError(err)
	}

	// Convert tasks to protobuf
	pbTasks := make([]*taskv1.Task, len(result.Tasks))
	for i, task := range result.Tasks {
		pbTask, err := s.domainToProtobufTask(&task)
		if err != nil {
			s.logger.Error("Failed to convert domain task to protobuf", zap.Error(err))
			continue // Skip this task rather than failing the entire request
		}
		pbTasks[i] = pbTask
	}

	return &taskv1.ListTasksResponse{
		Tasks:         pbTasks,
		NextPageToken: result.NextPageToken,
		TotalCount:    int32(result.TotalCount),
	}, nil
}

// BatchCreateTasks creates multiple tasks in a single request
func (s *TaskServer) BatchCreateTasks(ctx context.Context, req *taskv1.BatchCreateTasksRequest) (*taskv1.BatchCreateTasksResponse, error) {
	s.logger.Debug("BatchCreateTasks called", zap.Int("count", len(req.Requests)))

	responses := make([]*taskv1.CreateTaskResponse, 0, len(req.Requests))
	errors := make([]*taskv1.BatchError, 0)

	for i, createReq := range req.Requests {
		resp, err := s.CreateTask(ctx, createReq)
		if err != nil {
			errors = append(errors, &taskv1.BatchError{
				Index:   int32(i),
				Code:    status.Code(err).String(),
				Message: err.Error(),
			})
			continue
		}
		responses = append(responses, resp)
	}

	return &taskv1.BatchCreateTasksResponse{
		Responses: responses,
		Errors:    errors,
	}, nil
}

// BatchUpdateTasks updates multiple tasks in a single request
func (s *TaskServer) BatchUpdateTasks(ctx context.Context, req *taskv1.BatchUpdateTasksRequest) (*taskv1.BatchUpdateTasksResponse, error) {
	s.logger.Debug("BatchUpdateTasks called", zap.Int("count", len(req.Requests)))

	responses := make([]*taskv1.UpdateTaskResponse, 0, len(req.Requests))
	errors := make([]*taskv1.BatchError, 0)

	for i, updateReq := range req.Requests {
		resp, err := s.UpdateTask(ctx, updateReq)
		if err != nil {
			errors = append(errors, &taskv1.BatchError{
				Index:   int32(i),
				Code:    status.Code(err).String(),
				Message: err.Error(),
			})
			continue
		}
		responses = append(responses, resp)
	}

	return &taskv1.BatchUpdateTasksResponse{
		Responses: responses,
		Errors:    errors,
	}, nil
}

// BatchDeleteTasks deletes multiple tasks in a single request
func (s *TaskServer) BatchDeleteTasks(ctx context.Context, req *taskv1.BatchDeleteTasksRequest) (*emptypb.Empty, error) {
	s.logger.Debug("BatchDeleteTasks called", zap.Int("count", len(req.Ids)))

	for _, id := range req.Ids {
		if err := s.taskService.DeleteTask(ctx, id); err != nil {
			s.logger.Error("Failed to delete task in batch", zap.String("id", id), zap.Error(err))
			// In a production implementation, you might want to collect errors
			// and return them as part of the response
		}
	}

	return &emptypb.Empty{}, nil
}

// StreamTasks provides real-time task updates
func (s *TaskServer) StreamTasks(req *taskv1.StreamTasksRequest, stream taskv1.TaskService_StreamTasksServer) error {
	s.logger.Debug("StreamTasks called")

	// In a production implementation, you would:
	// 1. Subscribe to task events from the event bus
	// 2. Filter events based on the request filter
	// 3. Stream matching events to the client
	// 4. Handle client disconnections and cleanup

	// For now, we'll simulate with a ticker
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			// This is a mock event - in production, this would come from the event bus
			event := &taskv1.TaskEvent{
				Type:      taskv1.TaskEventType_TASK_EVENT_TYPE_UPDATED,
				Timestamp: timestamppb.Now(),
				UserId:    "system",
			}

			if err := stream.Send(&taskv1.StreamTasksResponse{
				Event: event,
			}); err != nil {
				s.logger.Error("Failed to send stream event", zap.Error(err))
				return err
			}
		}
	}
}

// GetTaskAnalytics provides task analytics and reporting
func (s *TaskServer) GetTaskAnalytics(ctx context.Context, req *taskv1.GetTaskAnalyticsRequest) (*taskv1.GetTaskAnalyticsResponse, error) {
	s.logger.Debug("GetTaskAnalytics called")

	// In a production implementation, you would:
	// 1. Query the database for analytics data
	// 2. Apply the requested filters and date range
	// 3. Calculate the requested metrics
	// 4. Return the results

	// For now, return mock data
	return &taskv1.GetTaskAnalyticsResponse{
		Results: []*taskv1.AnalyticsResult{
			{
				Metric: taskv1.AnalyticsMetric_ANALYTICS_METRIC_TASK_COUNT,
				Values: map[string]float64{
					"total": 150,
					"today": 5,
				},
			},
		},
	}, nil
}

// Helper methods

func (s *TaskServer) protobufToDomainTask(pbTask *taskv1.Task) (*domain.Task, error) {
	if pbTask == nil {
		return nil, fmt.Errorf("task cannot be nil")
	}

	task := &domain.Task{
		ID:          pbTask.Id,
		Title:       pbTask.Title,
		Description: pbTask.Description,
		Status:      s.convertStatus(pbTask.Status),
		Priority:    s.convertPriority(pbTask.Priority),
		Category:    pbTask.Category,
		Tags:        pbTask.Tags,
		Metadata:    pbTask.Metadata,
	}

	// Convert timestamps
	if pbTask.CreatedAt != nil {
		task.CreatedAt = pbTask.CreatedAt.AsTime()
	}
	if pbTask.UpdatedAt != nil {
		task.UpdatedAt = pbTask.UpdatedAt.AsTime()
	}
	if pbTask.DueDate != nil {
		dueDate := pbTask.DueDate.AsTime()
		task.DueDate = &dueDate
	}

	// Convert assignee
	if pbTask.Assignee != nil {
		task.AssigneeID = pbTask.Assignee.UserId
		task.AssigneeEmail = pbTask.Assignee.UserEmail
	}

	return task, nil
}

func (s *TaskServer) domainToProtobufTask(task *domain.Task) (*taskv1.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task cannot be nil")
	}

	pbTask := &taskv1.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      s.convertStatusToPb(task.Status),
		Priority:    s.convertPriorityToPb(task.Priority),
		Category:    task.Category,
		Tags:        task.Tags,
		Metadata:    task.Metadata,
		CreatedAt:   timestamppb.New(task.CreatedAt),
		UpdatedAt:   timestamppb.New(task.UpdatedAt),
	}

	// Convert due date
	if task.DueDate != nil {
		pbTask.DueDate = timestamppb.New(*task.DueDate)
	}

	// Convert assignee
	if task.AssigneeID != "" {
		pbTask.Assignee = &taskv1.TaskAssignee{
			UserId:    task.AssigneeID,
			UserEmail: task.AssigneeEmail,
		}
	}

	return pbTask, nil
}

func (s *TaskServer) convertStatus(pbStatus taskv1.TaskStatus) string {
	switch pbStatus {
	case taskv1.TaskStatus_TASK_STATUS_DRAFT:
		return "draft"
	case taskv1.TaskStatus_TASK_STATUS_PENDING:
		return "pending"
	case taskv1.TaskStatus_TASK_STATUS_IN_PROGRESS:
		return "in_progress"
	case taskv1.TaskStatus_TASK_STATUS_ON_HOLD:
		return "on_hold"
	case taskv1.TaskStatus_TASK_STATUS_COMPLETED:
		return "completed"
	case taskv1.TaskStatus_TASK_STATUS_CANCELLED:
		return "cancelled"
	case taskv1.TaskStatus_TASK_STATUS_ARCHIVED:
		return "archived"
	default:
		return "pending"
	}
}

func (s *TaskServer) convertStatusToPb(status string) taskv1.TaskStatus {
	switch status {
	case "draft":
		return taskv1.TaskStatus_TASK_STATUS_DRAFT
	case "pending":
		return taskv1.TaskStatus_TASK_STATUS_PENDING
	case "in_progress":
		return taskv1.TaskStatus_TASK_STATUS_IN_PROGRESS
	case "on_hold":
		return taskv1.TaskStatus_TASK_STATUS_ON_HOLD
	case "completed":
		return taskv1.TaskStatus_TASK_STATUS_COMPLETED
	case "cancelled":
		return taskv1.TaskStatus_TASK_STATUS_CANCELLED
	case "archived":
		return taskv1.TaskStatus_TASK_STATUS_ARCHIVED
	default:
		return taskv1.TaskStatus_TASK_STATUS_PENDING
	}
}

func (s *TaskServer) convertPriority(pbPriority taskv1.TaskPriority) string {
	switch pbPriority {
	case taskv1.TaskPriority_TASK_PRIORITY_LOW:
		return "low"
	case taskv1.TaskPriority_TASK_PRIORITY_MEDIUM:
		return "medium"
	case taskv1.TaskPriority_TASK_PRIORITY_HIGH:
		return "high"
	case taskv1.TaskPriority_TASK_PRIORITY_URGENT:
		return "urgent"
	case taskv1.TaskPriority_TASK_PRIORITY_CRITICAL:
		return "critical"
	default:
		return "medium"
	}
}

func (s *TaskServer) convertPriorityToPb(priority string) taskv1.TaskPriority {
	switch priority {
	case "low":
		return taskv1.TaskPriority_TASK_PRIORITY_LOW
	case "medium":
		return taskv1.TaskPriority_TASK_PRIORITY_MEDIUM
	case "high":
		return taskv1.TaskPriority_TASK_PRIORITY_HIGH
	case "urgent":
		return taskv1.TaskPriority_TASK_PRIORITY_URGENT
	case "critical":
		return taskv1.TaskPriority_TASK_PRIORITY_CRITICAL
	default:
		return taskv1.TaskPriority_TASK_PRIORITY_MEDIUM
	}
}

func (s *TaskServer) convertTaskFilter(pbFilter *taskv1.TaskFilter) services.TaskFilter {
	if pbFilter == nil {
		return services.TaskFilter{}
	}

	filter := services.TaskFilter{
		Categories: pbFilter.Categories,
		Tags:       pbFilter.Tags,
		Query:      pbFilter.SearchQuery,
	}

	// Convert status filters
	for _, status := range pbFilter.Status {
		filter.Statuses = append(filter.Statuses, s.convertStatus(status))
	}

	// Convert priority filters
	for _, priority := range pbFilter.Priority {
		filter.Priorities = append(filter.Priorities, s.convertPriority(priority))
	}

	return filter
}

func (s *TaskServer) handleServiceError(err error) error {
	// Convert domain/service errors to gRPC errors
	switch {
	case err == services.ErrTaskNotFound:
		return status.Errorf(codes.NotFound, "task not found")
	case err == services.ErrTaskAlreadyExists:
		return status.Errorf(codes.AlreadyExists, "task already exists")
	case err == services.ErrInvalidTaskData:
		return status.Errorf(codes.InvalidArgument, "invalid task data")
	case err == services.ErrUnauthorized:
		return status.Errorf(codes.PermissionDenied, "unauthorized")
	default:
		s.logger.Error("Unhandled service error", zap.Error(err))
		return status.Errorf(codes.Internal, "internal server error")
	}
}