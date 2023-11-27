// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
package federation

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"
	"sync"

	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/singleflight"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

// Federation_GetResponseArgument is argument for "federation.GetResponse" message.
type Federation_GetResponseArgument[T any] struct {
	Id     string
	P      *Post
	Client T
}

// Federation_GetResponse_PostArgument is custom resolver's argument for "post" field of "federation.GetResponse" message.
type Federation_GetResponse_PostArgument[T any] struct {
	*Federation_GetResponseArgument[T]
	Client T
}

// OtherServiceConfig configuration required to initialize the service that use GRPC Federation.
type OtherServiceConfig struct {
	// Resolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
	// If this interface is not provided, an error is returned during initialization.
	Resolver OtherServiceResolver // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// OtherServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type OtherServiceClientFactory interface {
}

// OtherServiceClientConfig information set in `dependencies` of the `grpc.federation.service` option.
// Hints for creating a gRPC Client.
type OtherServiceClientConfig struct {
	// Service returns the name of the service on Protocol Buffers.
	Service string
	// Name is the value set for `name` in `dependencies` of the `grpc.federation.service` option.
	// It must be unique among the services on which the Federation Service depends.
	Name string
}

// OtherServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type OtherServiceDependentClientSet struct {
}

// OtherServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type OtherServiceResolver interface {
	// Resolve_Federation_GetResponse_Post implements resolver for "federation.GetResponse.post".
	Resolve_Federation_GetResponse_Post(context.Context, *Federation_GetResponse_PostArgument[*OtherServiceDependentClientSet]) (*Post, error)
}

// OtherServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type OtherServiceUnimplementedResolver struct{}

// Resolve_Federation_GetResponse_Post resolve "federation.GetResponse.post".
// This method always returns Unimplemented error.
func (OtherServiceUnimplementedResolver) Resolve_Federation_GetResponse_Post(context.Context, *Federation_GetResponse_PostArgument[*OtherServiceDependentClientSet]) (ret *Post, e error) {
	e = grpcstatus.Errorf(grpccodes.Unimplemented, "method Resolve_Federation_GetResponse_Post not implemented")
	return
}

// OtherService represents Federation Service.
type OtherService struct {
	*UnimplementedOtherServiceServer
	cfg          OtherServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	tracer       trace.Tracer
	resolver     OtherServiceResolver
	client       *OtherServiceDependentClientSet
}

// NewOtherService creates OtherService instance by OtherServiceConfig.
func NewOtherService(cfg OtherServiceConfig) (*OtherService, error) {
	if cfg.Resolver == nil {
		return nil, fmt.Errorf("Resolver field in OtherServiceConfig is not set. this field must be set")
	}
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*celtypes.FieldType{
		"grpc.federation.private.GetResponseArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {},
		"grpc.federation.private.UserArgument": {
			"id":   grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
			"name": grpcfed.NewCELFieldType(celtypes.StringType, "Name"),
		},
	})
	env, err := cel.NewCustomEnv(
		cel.StdLib(),
		cel.CustomTypeAdapter(celHelper.TypeAdapter()),
		cel.CustomTypeProvider(celHelper.TypeProvider()),
	)
	if err != nil {
		return nil, err
	}
	return &OtherService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("federation.OtherService"),
		resolver:     cfg.Resolver,
		client:       &OtherServiceDependentClientSet{},
	}, nil
}

// Get implements "federation.OtherService/Get" method.
func (s *OtherService) Get(ctx context.Context, req *GetRequest) (res *GetResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.OtherService/Get")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Federation_GetResponse(ctx, &Federation_GetResponseArgument[*OtherServiceDependentClientSet]{
		Client: s.client,
		Id:     req.Id,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetResponse resolve "federation.GetResponse" message.
func (s *OtherService) resolve_Federation_GetResponse(ctx context.Context, req *Federation_GetResponseArgument[*OtherServiceDependentClientSet]) (*GetResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.GetResponse", slog.Any("message_args", s.logvalue_Federation_GetResponseArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
		valueP  *Post
	)

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "p"
	     message {
	       name: "Post"
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("p", func() (any, error) {
			valueMu.RLock()
			args := &Federation_PostArgument[*OtherServiceDependentClientSet]{
				Client: s.client,
			}
			valueMu.RUnlock()
			return s.resolve_Federation_Post(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*Post)
		valueMu.Lock()
		valueP = value // { name: "p", message: "Post" ... }
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.P = valueP

	// create a message value to be returned.
	ret := &GetResponse{}

	// field binding section.
	{
		// (grpc.federation.field).custom_resolver = true
		var err error
		ret.Post, err = s.resolver.Resolve_Federation_GetResponse_Post(ctx, &Federation_GetResponse_PostArgument[*OtherServiceDependentClientSet]{
			Client:                         s.client,
			Federation_GetResponseArgument: req,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	s.logger.DebugContext(ctx, "resolved federation.GetResponse", slog.Any("federation.GetResponse", s.logvalue_Federation_GetResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *OtherService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument[*OtherServiceDependentClientSet]) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Post")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
		valueU  *User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.PostArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "u"
	     message {
	       name: "User"
	       args: [
	         { name: "id", string: "foo" },
	         { name: "name", string: "bar" }
	       ]
	     }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("u", func() (any, error) {
			valueMu.RLock()
			args := &Federation_UserArgument[*OtherServiceDependentClientSet]{
				Client: s.client,
				Id:     "foo", // { name: "id", string: "foo" }
				Name:   "bar", // { name: "name", string: "bar" }
			}
			valueMu.RUnlock()
			return s.resolve_Federation_User(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*User)
		valueMu.Lock()
		valueU = value // { name: "u", message: "User" ... }
		envOpts = append(envOpts, cel.Variable("u", cel.ObjectType("federation.User")))
		evalValues["u"] = valueU
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.U = valueU

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = "post-id"      // (grpc.federation.field).string = "post-id"
	ret.Title = "title"     // (grpc.federation.field).string = "title"
	ret.Content = "content" // (grpc.federation.field).string = "content"
	// (grpc.federation.field).by = "u"
	{
		value, err := grpcfed.EvalCEL(s.env, "u", envOpts, evalValues, reflect.TypeOf((*User)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.User = value.(*User)
	}

	s.logger.DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *OtherService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument[*OtherServiceDependentClientSet]) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.id"
	{
		value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Id = value.(string)
	}
	// (grpc.federation.field).by = "$.name"
	{
		value, err := grpcfed.EvalCEL(s.env, "$.name", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Name = value.(string)
	}

	s.logger.DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

func (s *OtherService) logvalue_Federation_GetResponse(v *GetResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
	)
}

func (s *OtherService) logvalue_Federation_GetResponseArgument(v *Federation_GetResponseArgument[*OtherServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *OtherService) logvalue_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
	)
}

func (s *OtherService) logvalue_Federation_PostArgument(v *Federation_PostArgument[*OtherServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *OtherService) logvalue_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *OtherService) logvalue_Federation_UserArgument(v *Federation_UserArgument[*OtherServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("name", v.Name),
	)
}
