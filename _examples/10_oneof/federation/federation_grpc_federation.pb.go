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
	"golang.org/x/sync/singleflight"

	user "example/user"
)

// Org_Federation_GetResponseArgument is argument for "org.federation.GetResponse" message.
type Org_Federation_GetResponseArgument[T any] struct {
	Sel    *UserSelection
	Client T
}

// Org_Federation_UserArgument is argument for "org.federation.User" message.
type Org_Federation_UserArgument[T any] struct {
	UserId string
	Client T
}

// Org_Federation_UserSelectionArgument is argument for "org.federation.UserSelection" message.
type Org_Federation_UserSelectionArgument[T any] struct {
	Value  string
	Client T
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationServiceClientFactory // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
	// User_UserServiceClient create a gRPC Client to be used to call methods in user.UserService.
	User_UserServiceClient(FederationServiceClientConfig) (user.UserServiceClient, error)
}

// FederationServiceClientConfig information set in `dependencies` of the `grpc.federation.service` option.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service returns the name of the service on Protocol Buffers.
	Service string
	// Name is the value set for `name` in `dependencies` of the `grpc.federation.service` option.
	// It must be unique among the services on which the Federation Service depends.
	Name string
}

// FederationServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependentClientSet struct {
	User_UserServiceClient user.UserServiceClient
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

const (
	FederationService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationServiceClientConfig{
		Service: "user.UserService",
		Name:    "",
	})
	if err != nil {
		return nil, err
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
		"grpc.federation.private.GetResponseArgument": map[string]*celtypes.FieldType{},
		"grpc.federation.private.UserArgument": map[string]*celtypes.FieldType{
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
		},
		"grpc.federation.private.UserSelectionArgument": map[string]*celtypes.FieldType{
			"value": grpcfed.NewCELFieldType(celtypes.StringType, "Value"),
		},
		"org.federation.UserSelection": map[string]*celtypes.FieldType{
			"user": grpcfed.NewOneofSelectorFieldType(
				celtypes.NewObjectType("org.federation.User"), "User",
				[]reflect.Type{reflect.TypeOf((*UserSelection_UserA)(nil)), reflect.TypeOf((*UserSelection_UserB)(nil)), reflect.TypeOf((*UserSelection_UserC)(nil))},
				[]string{"GetUserA", "GetUserB", "GetUserC"},
				reflect.Zero(reflect.TypeOf((*User)(nil))),
			),
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
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		client: &FederationServiceDependentClientSet{
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

// Get implements "org.federation.FederationService/Get" method.
func (s *FederationService) Get(ctx context.Context, req *GetRequest) (res *GetResponse, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Org_Federation_GetResponse(ctx, &Org_Federation_GetResponseArgument[*FederationServiceDependentClientSet]{
		Client: s.client,
	})
	if err != nil {
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Org_Federation_GetResponse resolve "org.federation.GetResponse" message.
func (s *FederationService) resolve_Org_Federation_GetResponse(ctx context.Context, req *Org_Federation_GetResponseArgument[*FederationServiceDependentClientSet]) (*GetResponse, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.GetResponse", slog.Any("message_args", s.logvalue_Org_Federation_GetResponseArgument(req)))
	var (
		sg       singleflight.Group
		valueMu  sync.RWMutex
		valueSel *UserSelection
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.GetResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "sel"
	     message: "UserSelection"
	     args { name: "value", string: "foo" }
	   }
	*/
	resUserSelectionIface, err, _ := sg.Do("sel_org.federation.UserSelection", func() (any, error) {
		valueMu.RLock()
		args := &Org_Federation_UserSelectionArgument[*FederationServiceDependentClientSet]{
			Client: s.client,
			Value:  "foo", // { name: "value", string: "foo" }
		}
		valueMu.RUnlock()
		return s.resolve_Org_Federation_UserSelection(ctx, args)
	})
	if err != nil {
		return nil, err
	}
	resUserSelection := resUserSelectionIface.(*UserSelection)
	valueMu.Lock()
	valueSel = resUserSelection // { name: "sel", message: "UserSelection" ... }
	envOpts = append(envOpts, cel.Variable("sel", cel.ObjectType("org.federation.UserSelection")))
	evalValues["sel"] = valueSel
	valueMu.Unlock()

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Sel = valueSel

	// create a message value to be returned.
	ret := &GetResponse{}

	// field binding section.
	// (grpc.federation.field).by = "sel.user"
	{
		_value, err := grpcfed.EvalCEL(s.env, "sel.user", envOpts, evalValues, reflect.TypeOf((*User)(nil)))
		if err != nil {
			return nil, err
		}
		ret.User = _value.(*User)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.GetResponse", slog.Any("org.federation.GetResponse", s.logvalue_Org_Federation_GetResponse(ret)))
	return ret, nil
}

// resolve_Org_Federation_User resolve "org.federation.User" message.
func (s *FederationService) resolve_Org_Federation_User(ctx context.Context, req *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.User", slog.Any("message_args", s.logvalue_Org_Federation_UserArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "user.UserService/GetUser"
	     request { field: "id", by: "$.user_id" }
	   }
	*/
	if _, err, _ := sg.Do("user.UserService/GetUser", func() (any, error) {
		valueMu.RLock()
		args := &user.GetUserRequest{}
		// { field: "id", by: "$.user_id" }
		{
			_value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
			if err != nil {
				return nil, err
			}
			args.Id = _value.(string)
		}
		valueMu.RUnlock()
		return s.client.User_UserServiceClient.GetUser(ctx, args)
	}); err != nil {
		if err := s.errorHandler(ctx, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
			return nil, err
		}
	}

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.user_id"
	{
		_value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
		if err != nil {
			return nil, err
		}
		ret.Id = _value.(string)
	}

	s.logger.DebugContext(ctx, "resolved org.federation.User", slog.Any("org.federation.User", s.logvalue_Org_Federation_User(ret)))
	return ret, nil
}

// resolve_Org_Federation_UserSelection resolve "org.federation.UserSelection" message.
func (s *FederationService) resolve_Org_Federation_UserSelection(ctx context.Context, req *Org_Federation_UserSelectionArgument[*FederationServiceDependentClientSet]) (*UserSelection, error) {
	s.logger.DebugContext(ctx, "resolve  org.federation.UserSelection", slog.Any("message_args", s.logvalue_Org_Federation_UserSelectionArgument(req)))
	var (
		sg      singleflight.Group
		valueMu sync.RWMutex
		valueUa *User
		valueUb *User
		valueUc *User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserSelectionArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// create a message value to be returned.
	ret := &UserSelection{}

	// field binding section.

	oneof_UserA, err := grpcfed.EvalCEL(s.env, "false", envOpts, evalValues, reflect.TypeOf(true))
	if err != nil {
		return nil, err
	}
	oneof_UserB, err := grpcfed.EvalCEL(s.env, "true", envOpts, evalValues, reflect.TypeOf(true))
	if err != nil {
		return nil, err
	}
	switch {
	case oneof_UserA.(bool):

		// This section's codes are generated by the following proto definition.
		/*
		   {
		     name: "ua"
		     message: "User"
		     args { name: "user_id", string: "a" }
		   }
		*/
		resUserIface, err, _ := sg.Do("ua_org.federation.User", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
				UserId: "a", // { name: "user_id", string: "a" }
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_User(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		resUser := resUserIface.(*User)
		valueMu.Lock()
		valueUa = resUser // { name: "ua", message: "User" ... }
		envOpts = append(envOpts, cel.Variable("ua", cel.ObjectType("org.federation.User")))
		evalValues["ua"] = valueUa
		valueMu.Unlock()
		_value, err := grpcfed.EvalCEL(s.env, "ua", envOpts, evalValues, nil)
		if err != nil {
			return nil, err
		}
		ret.User = &UserSelection_UserA{UserA: _value.(*User)}
	case oneof_UserB.(bool):

		// This section's codes are generated by the following proto definition.
		/*
		   {
		     name: "ub"
		     message: "User"
		     args { name: "user_id", string: "b" }
		   }
		*/
		resUserIface, err, _ := sg.Do("ub_org.federation.User", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
				UserId: "b", // { name: "user_id", string: "b" }
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_User(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		resUser := resUserIface.(*User)
		valueMu.Lock()
		valueUb = resUser // { name: "ub", message: "User" ... }
		envOpts = append(envOpts, cel.Variable("ub", cel.ObjectType("org.federation.User")))
		evalValues["ub"] = valueUb
		valueMu.Unlock()
		_value, err := grpcfed.EvalCEL(s.env, "ub", envOpts, evalValues, nil)
		if err != nil {
			return nil, err
		}
		ret.User = &UserSelection_UserB{UserB: _value.(*User)}
	default:

		// This section's codes are generated by the following proto definition.
		/*
		   {
		     name: "uc"
		     message: "User"
		     args { name: "user_id", by: "$.value" }
		   }
		*/
		resUserIface, err, _ := sg.Do("uc_org.federation.User", func() (any, error) {
			valueMu.RLock()
			args := &Org_Federation_UserArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "user_id", by: "$.value" }
			{
				_value, err := grpcfed.EvalCEL(s.env, "$.value", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					return nil, err
				}
				args.UserId = _value.(string)
			}
			valueMu.RUnlock()
			return s.resolve_Org_Federation_User(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		resUser := resUserIface.(*User)
		valueMu.Lock()
		valueUc = resUser // { name: "uc", message: "User" ... }
		envOpts = append(envOpts, cel.Variable("uc", cel.ObjectType("org.federation.User")))
		evalValues["uc"] = valueUc
		valueMu.Unlock()
		_value, err := grpcfed.EvalCEL(s.env, "uc", envOpts, evalValues, nil)
		if err != nil {
			return nil, err
		}
		ret.User = &UserSelection_UserC{UserC: _value.(*User)}
	}

	s.logger.DebugContext(ctx, "resolved org.federation.UserSelection", slog.Any("org.federation.UserSelection", s.logvalue_Org_Federation_UserSelection(ret)))
	return ret, nil
}

func (s *FederationService) logvalue_Org_Federation_GetResponse(v *GetResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("user", s.logvalue_Org_Federation_User(v.GetUser())),
	)
}

func (s *FederationService) logvalue_Org_Federation_GetResponseArgument(v *Org_Federation_GetResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Org_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserArgument(v *Org_Federation_UserArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserSelection(v *UserSelection) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("user_a", s.logvalue_Org_Federation_User(v.GetUserA())),
		slog.Any("user_b", s.logvalue_Org_Federation_User(v.GetUserB())),
		slog.Any("user_c", s.logvalue_Org_Federation_User(v.GetUserC())),
	)
}

func (s *FederationService) logvalue_Org_Federation_UserSelectionArgument(v *Org_Federation_UserSelectionArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("value", v.Value),
	)
}
