package server

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"go.uber.org/zap"

	"github.com/vertikon/mcp-ultra/internal/config"
	"github.com/vertikon/mcp-ultra/internal/observability"
)

// observabilityUnaryInterceptor adds observability to unary gRPC calls
func observabilityUnaryInterceptor(obsService *observability.Service, logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if obsService == nil {
			return handler(ctx, req)
		}

		startTime := time.Now()
		
		// Extract metadata for tracing
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = obsService.ExtractTraceContext(ctx, md)
		}

		// Start a span for this gRPC call
		ctx, span := obsService.StartSpan(ctx, fmt.Sprintf("grpc.%s", info.FullMethod))
		defer span.End()

		// Add gRPC-specific attributes to span
		span.SetAttributes(map[string]interface{}{
			"grpc.method":      info.FullMethod,
			"grpc.service":     extractServiceName(info.FullMethod),
			"grpc.method_name": extractMethodName(info.FullMethod),
			"grpc.method_type": "unary",
		})

		// Execute the handler
		resp, err := handler(ctx, req)

		// Record metrics
		duration := time.Since(startTime)
		status := codes.OK
		if err != nil {
			status = grpc.Code(err)
		}

		obsService.RecordGRPCMetrics(ctx, info.FullMethod, status.String(), duration)

		// Add status to span
		if err != nil {
			span.RecordError(err)
			span.SetAttributes(map[string]interface{}{
				"grpc.status_code": status,
				"grpc.error":       true,
			})
		} else {
			span.SetAttributes(map[string]interface{}{
				"grpc.status_code": codes.OK,
				"grpc.error":       false,
			})
		}

		// Log the call
		logLevel := zap.InfoLevel
		if err != nil {
			logLevel = zap.ErrorLevel
		}

		logger.Log(logLevel, "gRPC unary call",
			zap.String("method", info.FullMethod),
			zap.Duration("duration", duration),
			zap.String("status", status.String()),
			zap.Error(err),
		)

		return resp, err
	}
}

// observabilityStreamInterceptor adds observability to streaming gRPC calls
func observabilityStreamInterceptor(obsService *observability.Service, logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if obsService == nil {
			return handler(srv, stream)
		}

		startTime := time.Now()
		ctx := stream.Context()

		// Extract metadata for tracing
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = obsService.ExtractTraceContext(ctx, md)
		}

		// Start a span for this gRPC call
		ctx, span := obsService.StartSpan(ctx, fmt.Sprintf("grpc.%s", info.FullMethod))
		defer span.End()

		// Add gRPC-specific attributes to span
		span.SetAttributes(map[string]interface{}{
			"grpc.method":         info.FullMethod,
			"grpc.service":        extractServiceName(info.FullMethod),
			"grpc.method_name":    extractMethodName(info.FullMethod),
			"grpc.method_type":    "stream",
			"grpc.client_stream":  info.IsClientStream,
			"grpc.server_stream":  info.IsServerStream,
		})

		// Wrap the stream to use the new context
		wrappedStream := &wrappedServerStream{
			ServerStream: stream,
			ctx:          ctx,
		}

		// Execute the handler
		err := handler(srv, wrappedStream)

		// Record metrics
		duration := time.Since(startTime)
		status := codes.OK
		if err != nil {
			status = grpc.Code(err)
		}

		obsService.RecordGRPCMetrics(ctx, info.FullMethod, status.String(), duration)

		// Add status to span
		if err != nil {
			span.RecordError(err)
			span.SetAttributes(map[string]interface{}{
				"grpc.status_code": status,
				"grpc.error":       true,
			})
		} else {
			span.SetAttributes(map[string]interface{}{
				"grpc.status_code": codes.OK,
				"grpc.error":       false,
			})
		}

		// Log the call
		logLevel := zap.InfoLevel
		if err != nil {
			logLevel = zap.ErrorLevel
		}

		logger.Log(logLevel, "gRPC stream call",
			zap.String("method", info.FullMethod),
			zap.Duration("duration", duration),
			zap.String("status", status.String()),
			zap.Bool("client_stream", info.IsClientStream),
			zap.Bool("server_stream", info.IsServerStream),
			zap.Error(err),
		)

		return err
	}
}

// authenticationUnaryInterceptor handles authentication for unary calls
func authenticationUnaryInterceptor(config *config.Config, logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Skip authentication for health checks and reflection
		if isPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		// Skip authentication if not configured
		if !config.Security.Auth.Enabled {
			return handler(ctx, req)
		}

		// Extract authorization metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		token := authHeaders[0]
		if token == "" {
			return nil, status.Errorf(codes.Unauthenticated, "empty authorization token")
		}

		// TODO: Implement actual token validation
		// For now, just check if token starts with "Bearer "
		if len(token) < 7 || token[:7] != "Bearer " {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization format")
		}

		// TODO: Validate token and extract user info
		// Add user info to context for downstream services
		ctx = context.WithValue(ctx, "user_id", "authenticated_user")

		logger.Debug("Authentication successful",
			zap.String("method", info.FullMethod),
			zap.String("user_id", "authenticated_user"),
		)

		return handler(ctx, req)
	}
}

// authenticationStreamInterceptor handles authentication for stream calls
func authenticationStreamInterceptor(config *config.Config, logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Skip authentication for health checks and reflection
		if isPublicMethod(info.FullMethod) {
			return handler(srv, stream)
		}

		// Skip authentication if not configured
		if !config.Security.Auth.Enabled {
			return handler(srv, stream)
		}

		ctx := stream.Context()

		// Extract authorization metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			return status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		token := authHeaders[0]
		if token == "" {
			return status.Errorf(codes.Unauthenticated, "empty authorization token")
		}

		// TODO: Implement actual token validation
		if len(token) < 7 || token[:7] != "Bearer " {
			return status.Errorf(codes.Unauthenticated, "invalid authorization format")
		}

		// TODO: Validate token and extract user info
		ctx = context.WithValue(ctx, "user_id", "authenticated_user")

		// Wrap the stream to use the new context
		wrappedStream := &wrappedServerStream{
			ServerStream: stream,
			ctx:          ctx,
		}

		logger.Debug("Stream authentication successful",
			zap.String("method", info.FullMethod),
			zap.String("user_id", "authenticated_user"),
		)

		return handler(srv, wrappedStream)
	}
}

// rateLimitingUnaryInterceptor implements rate limiting for unary calls
func rateLimitingUnaryInterceptor(config *config.Config, logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Skip rate limiting if not configured
		if !config.Security.RateLimit.Enabled {
			return handler(ctx, req)
		}

		// Skip rate limiting for health checks
		if isPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		// TODO: Implement actual rate limiting logic
		// For now, just log that rate limiting would be applied
		logger.Debug("Rate limiting check",
			zap.String("method", info.FullMethod),
			zap.Int("limit", config.Security.RateLimit.RequestsPerSecond),
		)

		return handler(ctx, req)
	}
}

// rateLimitingStreamInterceptor implements rate limiting for stream calls
func rateLimitingStreamInterceptor(config *config.Config, logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Skip rate limiting if not configured
		if !config.Security.RateLimit.Enabled {
			return handler(srv, stream)
		}

		// Skip rate limiting for health checks
		if isPublicMethod(info.FullMethod) {
			return handler(srv, stream)
		}

		// TODO: Implement actual rate limiting logic
		logger.Debug("Stream rate limiting check",
			zap.String("method", info.FullMethod),
			zap.Int("limit", config.Security.RateLimit.RequestsPerSecond),
		)

		return handler(srv, stream)
	}
}

// validationUnaryInterceptor validates request messages
func validationUnaryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// TODO: Implement request validation using protobuf validation rules
		// For now, just log that validation would be applied
		logger.Debug("Request validation",
			zap.String("method", info.FullMethod),
			zap.String("request_type", fmt.Sprintf("%T", req)),
		)

		return handler(ctx, req)
	}
}

// Helper functions

// wrappedServerStream wraps grpc.ServerStream to override the context
type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

// extractServiceName extracts the service name from a full method name
func extractServiceName(fullMethod string) string {
	// fullMethod format: "/package.service/method"
	if len(fullMethod) < 2 {
		return ""
	}
	
	// Remove leading slash
	method := fullMethod[1:]
	
	// Find the last slash
	lastSlash := -1
	for i := len(method) - 1; i >= 0; i-- {
		if method[i] == '/' {
			lastSlash = i
			break
		}
	}
	
	if lastSlash == -1 {
		return ""
	}
	
	return method[:lastSlash]
}

// extractMethodName extracts the method name from a full method name
func extractMethodName(fullMethod string) string {
	// fullMethod format: "/package.service/method"
	if len(fullMethod) < 2 {
		return ""
	}
	
	// Remove leading slash
	method := fullMethod[1:]
	
	// Find the last slash
	lastSlash := -1
	for i := len(method) - 1; i >= 0; i-- {
		if method[i] == '/' {
			lastSlash = i
			break
		}
	}
	
	if lastSlash == -1 {
		return method
	}
	
	return method[lastSlash+1:]
}

// isPublicMethod checks if a method should skip authentication
func isPublicMethod(fullMethod string) bool {
	publicMethods := []string{
		"/grpc.health.v1.Health/Check",
		"/grpc.health.v1.Health/Watch",
		"/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo",
		"/system.v1.SystemService/Check", // Allow health checks without auth
	}
	
	for _, public := range publicMethods {
		if fullMethod == public {
			return true
		}
	}
	
	return false
}