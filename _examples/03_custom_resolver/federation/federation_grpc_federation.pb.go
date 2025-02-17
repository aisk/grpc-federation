// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
// versions:
//
//	protoc-gen-grpc-federation: dev
//
// source: federation/federation.proto
package federation

import (
	"context"
	"io"
	"log/slog"
	"reflect"
	"runtime/debug"

	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	grpcfedcel "github.com/mercari/grpc-federation/grpc/federation/cel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	post "example/post"
	user "example/user"
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Federation_V2Dev_ForNamelessArgument is argument for "federation.v2dev.ForNameless" message.
type Federation_V2Dev_ForNamelessArgument struct {
	Bar string
}

// Federation_V2Dev_GetPostV2DevResponseArgument is argument for "federation.v2dev.GetPostV2devResponse" message.
type Federation_V2Dev_GetPostV2DevResponseArgument struct {
	Id   string
	Post *PostV2Dev
}

// Federation_V2Dev_PostV2DevArgument is argument for "federation.v2dev.PostV2dev" message.
type Federation_V2Dev_PostV2DevArgument struct {
	Id     string
	Post   *post.Post
	Res    *post.GetPostResponse
	Unused *Unused
	User   *User
	XDef4  *ForNameless
}

// Federation_V2Dev_PostV2Dev_UserArgument is custom resolver's argument for "user" field of "federation.v2dev.PostV2dev" message.
type Federation_V2Dev_PostV2Dev_UserArgument struct {
	*Federation_V2Dev_PostV2DevArgument
}

// Federation_V2Dev_UnusedArgument is argument for "federation.v2dev.Unused" message.
type Federation_V2Dev_UnusedArgument struct {
	Foo string
}

// Federation_V2Dev_UserArgument is argument for "federation.v2dev.User" message.
type Federation_V2Dev_UserArgument struct {
	Content string
	Id      string
	Res     *user.GetUserResponse
	Title   string
	U       *user.User
	UserId  string
}

// Federation_V2Dev_User_NameArgument is custom resolver's argument for "name" field of "federation.v2dev.User" message.
type Federation_V2Dev_User_NameArgument struct {
	*Federation_V2Dev_UserArgument
	Federation_V2Dev_User *User
}

// FederationV2DevServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationV2DevServiceConfig struct {
	// Client provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
	// If this interface is not provided, an error is returned during initialization.
	Client FederationV2DevServiceClientFactory // required
	// Resolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
	// If this interface is not provided, an error is returned during initialization.
	Resolver FederationV2DevServiceResolver // required
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationV2DevServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationV2DevServiceClientFactory interface {
	// Post_PostServiceClient create a gRPC Client to be used to call methods in post.PostService.
	Post_PostServiceClient(FederationV2DevServiceClientConfig) (post.PostServiceClient, error)
	// User_UserServiceClient create a gRPC Client to be used to call methods in user.UserService.
	User_UserServiceClient(FederationV2DevServiceClientConfig) (user.UserServiceClient, error)
}

// FederationV2DevServiceClientConfig helper to create gRPC client.
// Hints for creating a gRPC Client.
type FederationV2DevServiceClientConfig struct {
	// Service FQDN ( `<package-name>.<service-name>` ) of the service on Protocol Buffers.
	Service string
}

// FederationV2DevServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationV2DevServiceDependentClientSet struct {
	Post_PostServiceClient post.PostServiceClient
	User_UserServiceClient user.UserServiceClient
}

// FederationV2DevServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationV2DevServiceResolver interface {
	// Resolve_Federation_V2Dev_ForNameless implements resolver for "federation.v2dev.ForNameless".
	Resolve_Federation_V2Dev_ForNameless(context.Context, *Federation_V2Dev_ForNamelessArgument) (*ForNameless, error)
	// Resolve_Federation_V2Dev_PostV2Dev_User implements resolver for "federation.v2dev.PostV2dev.user".
	Resolve_Federation_V2Dev_PostV2Dev_User(context.Context, *Federation_V2Dev_PostV2Dev_UserArgument) (*User, error)
	// Resolve_Federation_V2Dev_Unused implements resolver for "federation.v2dev.Unused".
	Resolve_Federation_V2Dev_Unused(context.Context, *Federation_V2Dev_UnusedArgument) (*Unused, error)
	// Resolve_Federation_V2Dev_User implements resolver for "federation.v2dev.User".
	Resolve_Federation_V2Dev_User(context.Context, *Federation_V2Dev_UserArgument) (*User, error)
	// Resolve_Federation_V2Dev_User_Name implements resolver for "federation.v2dev.User.name".
	Resolve_Federation_V2Dev_User_Name(context.Context, *Federation_V2Dev_User_NameArgument) (string, error)
}

// FederationV2DevServiceCELPluginWasmConfig type alias for grpcfedcel.WasmConfig.
type FederationV2DevServiceCELPluginWasmConfig = grpcfedcel.WasmConfig

// FederationV2DevServiceCELPluginConfig hints for loading a WebAssembly based plugin.
type FederationV2DevServiceCELPluginConfig struct {
}

// FederationV2DevServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationV2DevServiceUnimplementedResolver struct{}

// Resolve_Federation_V2Dev_ForNameless resolve "federation.v2dev.ForNameless".
// This method always returns Unimplemented error.
func (FederationV2DevServiceUnimplementedResolver) Resolve_Federation_V2Dev_ForNameless(context.Context, *Federation_V2Dev_ForNamelessArgument) (ret *ForNameless, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_V2Dev_ForNameless not implemented")
	return
}

// Resolve_Federation_V2Dev_PostV2Dev_User resolve "federation.v2dev.PostV2dev.user".
// This method always returns Unimplemented error.
func (FederationV2DevServiceUnimplementedResolver) Resolve_Federation_V2Dev_PostV2Dev_User(context.Context, *Federation_V2Dev_PostV2Dev_UserArgument) (ret *User, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_V2Dev_PostV2Dev_User not implemented")
	return
}

// Resolve_Federation_V2Dev_Unused resolve "federation.v2dev.Unused".
// This method always returns Unimplemented error.
func (FederationV2DevServiceUnimplementedResolver) Resolve_Federation_V2Dev_Unused(context.Context, *Federation_V2Dev_UnusedArgument) (ret *Unused, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_V2Dev_Unused not implemented")
	return
}

// Resolve_Federation_V2Dev_User resolve "federation.v2dev.User".
// This method always returns Unimplemented error.
func (FederationV2DevServiceUnimplementedResolver) Resolve_Federation_V2Dev_User(context.Context, *Federation_V2Dev_UserArgument) (ret *User, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_V2Dev_User not implemented")
	return
}

// Resolve_Federation_V2Dev_User_Name resolve "federation.v2dev.User.name".
// This method always returns Unimplemented error.
func (FederationV2DevServiceUnimplementedResolver) Resolve_Federation_V2Dev_User_Name(context.Context, *Federation_V2Dev_User_NameArgument) (ret string, e error) {
	e = grpcfed.GRPCErrorf(grpcfed.UnimplementedCode, "method Resolve_Federation_V2Dev_User_Name not implemented")
	return
}

const (
	FederationV2DevService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
	FederationV2DevService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
)

// FederationV2DevService represents Federation Service.
type FederationV2DevService struct {
	*UnimplementedFederationV2DevServiceServer
	cfg           FederationV2DevServiceConfig
	logger        *slog.Logger
	errorHandler  grpcfed.ErrorHandler
	celCacheMap   *grpcfed.CELCacheMap
	tracer        trace.Tracer
	resolver      FederationV2DevServiceResolver
	celTypeHelper *grpcfed.CELTypeHelper
	envOpts       []grpcfed.CELEnvOption
	celPlugins    []*grpcfedcel.CELPlugin
	client        *FederationV2DevServiceDependentClientSet
}

// NewFederationV2DevService creates FederationV2DevService instance by FederationV2DevServiceConfig.
func NewFederationV2DevService(cfg FederationV2DevServiceConfig) (*FederationV2DevService, error) {
	if cfg.Client == nil {
		return nil, grpcfed.ErrClientConfig
	}
	if cfg.Resolver == nil {
		return nil, grpcfed.ErrResolverConfig
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationV2DevServiceClientConfig{
		Service: "post.PostService",
	})
	if err != nil {
		return nil, err
	}
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationV2DevServiceClientConfig{
		Service: "user.UserService",
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
	celTypeHelperFieldMap := grpcfed.CELTypeHelperFieldMap{
		"grpc.federation.private.ForNamelessArgument": {
			"bar": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Bar"),
		},
		"grpc.federation.private.GetPostV2devResponseArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.PostV2devArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.UnusedArgument": {
			"foo": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Foo"),
		},
		"grpc.federation.private.UserArgument": {
			"id":      grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
			"title":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Title"),
			"content": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "UserId"),
		},
	}
	celTypeHelper := grpcfed.NewCELTypeHelper(celTypeHelperFieldMap)
	var envOpts []grpcfed.CELEnvOption
	envOpts = append(envOpts, grpcfed.NewDefaultEnvOptions(celTypeHelper)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.v2dev.PostV2devType", PostV2DevType_value, PostV2DevType_name)...)
	return &FederationV2DevService{
		cfg:           cfg,
		logger:        logger,
		errorHandler:  errorHandler,
		envOpts:       envOpts,
		celTypeHelper: celTypeHelper,
		celCacheMap:   grpcfed.NewCELCacheMap(),
		tracer:        otel.Tracer("federation.v2dev.FederationV2devService"),
		resolver:      cfg.Resolver,
		client: &FederationV2DevServiceDependentClientSet{
			Post_PostServiceClient: Post_PostServiceClient,
			User_UserServiceClient: User_UserServiceClient,
		},
	}, nil
}

// GetPostV2Dev implements "federation.v2dev.FederationV2devService/GetPostV2dev" method.
func (s *FederationV2DevService) GetPostV2Dev(ctx context.Context, req *GetPostV2DevRequest) (res *GetPostV2DevResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.v2dev.FederationV2devService/GetPostV2dev")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	ctx = grpcfed.WithCELCacheMap(ctx, s.celCacheMap)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, e)
		}
	}()
	res, err := s.resolve_Federation_V2Dev_GetPostV2DevResponse(ctx, &Federation_V2Dev_GetPostV2DevResponseArgument{
		Id: req.GetId(),
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_V2Dev_ForNameless resolve "federation.v2dev.ForNameless" message.
func (s *FederationV2DevService) resolve_Federation_V2Dev_ForNameless(ctx context.Context, req *Federation_V2Dev_ForNamelessArgument) (*ForNameless, error) {
	ctx, span := s.tracer.Start(ctx, "federation.v2dev.ForNameless")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.v2dev.ForNameless", slog.Any("message_args", s.logvalue_Federation_V2Dev_ForNamelessArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx)) // create a new reference to logger.
	ret, err := s.resolver.Resolve_Federation_V2Dev_ForNameless(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.v2dev.ForNameless", slog.Any("federation.v2dev.ForNameless", s.logvalue_Federation_V2Dev_ForNameless(ret)))
	return ret, nil
}

// resolve_Federation_V2Dev_GetPostV2DevResponse resolve "federation.v2dev.GetPostV2devResponse" message.
func (s *FederationV2DevService) resolve_Federation_V2Dev_GetPostV2DevResponse(ctx context.Context, req *Federation_V2Dev_GetPostV2DevResponseArgument) (*GetPostV2DevResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.v2dev.GetPostV2devResponse")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.v2dev.GetPostV2devResponse", slog.Any("message_args", s.logvalue_Federation_V2Dev_GetPostV2DevResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			post *PostV2Dev
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.GetPostV2devResponseArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "post"
	     message {
	       name: "PostV2dev"
	       args { name: "id", by: "$.id" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*PostV2Dev, *localValueType]{
		Name: `post`,
		Type: grpcfed.CELObjectType("federation.v2dev.PostV2dev"),
		Setter: func(value *localValueType, v *PostV2Dev) error {
			value.vars.post = v
			return nil
		},
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Federation_V2Dev_PostV2DevArgument{}
			// { name: "id", by: "$.id" }
			if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
				Value:             value,
				Expr:              `$.id`,
				UseContextLibrary: false,
				CacheIndex:        1,
				Setter: func(v string) error {
					args.Id = v
					return nil
				},
			}); err != nil {
				return nil, err
			}
			return s.resolve_Federation_V2Dev_PostV2Dev(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = value.vars.post

	// create a message value to be returned.
	ret := &GetPostV2DevResponse{}

	// field binding section.
	// (grpc.federation.field).by = "post"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*PostV2Dev]{
		Value:             value,
		Expr:              `post`,
		UseContextLibrary: false,
		CacheIndex:        2,
		Setter: func(v *PostV2Dev) error {
			ret.Post = v
			return nil
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "federation.v2dev.PostV2devType.POST_V2_DEV_TYPE"
	if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[PostV2DevType]{
		Value:             value,
		Expr:              `federation.v2dev.PostV2devType.POST_V2_DEV_TYPE`,
		UseContextLibrary: false,
		CacheIndex:        3,
		Setter: func(v PostV2DevType) error {
			ret.Type = v
			return nil
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.v2dev.GetPostV2devResponse", slog.Any("federation.v2dev.GetPostV2devResponse", s.logvalue_Federation_V2Dev_GetPostV2DevResponse(ret)))
	return ret, nil
}

// resolve_Federation_V2Dev_PostV2Dev resolve "federation.v2dev.PostV2dev" message.
func (s *FederationV2DevService) resolve_Federation_V2Dev_PostV2Dev(ctx context.Context, req *Federation_V2Dev_PostV2DevArgument) (*PostV2Dev, error) {
	ctx, span := s.tracer.Start(ctx, "federation.v2dev.PostV2dev")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.v2dev.PostV2dev", slog.Any("message_args", s.logvalue_Federation_V2Dev_PostV2DevArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			_def4  *ForNameless
			post   *post.Post
			res    *post.GetPostResponse
			unused *Unused
			user   *User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.PostV2devArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()
	// A tree view of message dependencies is shown below.
	/*
	               _def4 ─┐
	              unused ─┤
	   res ─┐             │
	        post ─┐       │
	                user ─┤
	*/
	eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "_def4"
		     message {
		       name: "ForNameless"
		       args { name: "bar", by: "'bar'" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*ForNameless, *localValueType]{
			Name: `_def4`,
			Type: grpcfed.CELObjectType("federation.v2dev.ForNameless"),
			Setter: func(value *localValueType, v *ForNameless) error {
				value.vars._def4 = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_V2Dev_ForNamelessArgument{}
				// { name: "bar", by: "'bar'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'bar'`,
					UseContextLibrary: false,
					CacheIndex:        4,
					Setter: func(v string) error {
						args.Bar = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Federation_V2Dev_ForNameless(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "unused"
		     message {
		       name: "Unused"
		       args { name: "foo", by: "'foo'" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Unused, *localValueType]{
			Name: `unused`,
			Type: grpcfed.CELObjectType("federation.v2dev.Unused"),
			Setter: func(value *localValueType, v *Unused) error {
				value.vars.unused = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_V2Dev_UnusedArgument{}
				// { name: "foo", by: "'foo'" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `'foo'`,
					UseContextLibrary: false,
					CacheIndex:        5,
					Setter: func(v string) error {
						args.Foo = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Federation_V2Dev_Unused(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "res"
		     call {
		       method: "post.PostService/GetPost"
		       request { field: "id", by: "$.id" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*post.GetPostResponse, *localValueType]{
			Name: `res`,
			Type: grpcfed.CELObjectType("post.GetPostResponse"),
			Setter: func(value *localValueType, v *post.GetPostResponse) error {
				value.vars.res = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &post.GetPostRequest{}
				// { field: "id", by: "$.id" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
					Value:             value,
					Expr:              `$.id`,
					UseContextLibrary: false,
					CacheIndex:        6,
					Setter: func(v string) error {
						args.Id = v
						return nil
					},
				}); err != nil {
					return nil, err
				}
				grpcfed.Logger(ctx).DebugContext(ctx, "call post.PostService/GetPost", slog.Any("post.GetPostRequest", s.logvalue_Post_GetPostRequest(args)))
				return s.client.Post_PostServiceClient.GetPost(ctx, args)
			},
		}); err != nil {
			if err := s.errorHandler(ctx1, FederationV2DevService_DependentMethod_Post_PostService_GetPost, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx1, err)
				return nil, err
			}
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "post"
		     autobind: true
		     by: "res.post"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*post.Post, *localValueType]{
			Name: `post`,
			Type: grpcfed.CELObjectType("post.Post"),
			Setter: func(value *localValueType, v *post.Post) error {
				value.vars.post = v
				return nil
			},
			By:                  `res.post`,
			ByUseContextLibrary: false,
			ByCacheIndex:        7,
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "user"
		     message {
		       name: "User"
		       args { inline: "post" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*User, *localValueType]{
			Name: `user`,
			Type: grpcfed.CELObjectType("federation.v2dev.User"),
			Setter: func(value *localValueType, v *User) error {
				value.vars.user = v
				return nil
			},
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_V2Dev_UserArgument{}
				// { inline: "post" }
				if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[*post.Post]{
					Value:             value,
					Expr:              `post`,
					UseContextLibrary: false,
					CacheIndex:        8,
					Setter: func(v *post.Post) error {
						args.Id = v.GetId()
						args.Title = v.GetTitle()
						args.Content = v.GetContent()
						args.UserId = v.GetUserId()
						return nil
					},
				}); err != nil {
					return nil, err
				}
				return s.resolve_Federation_V2Dev_User(ctx, args)
			},
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}
		return nil, nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = value.vars.post
	req.Res = value.vars.res
	req.Unused = value.vars.unused
	req.User = value.vars.user
	req.XDef4 = value.vars._def4

	// create a message value to be returned.
	ret := &PostV2Dev{}

	// field binding section.
	ret.Id = value.vars.post.GetId()           // { name: "post", autobind: true }
	ret.Title = value.vars.post.GetTitle()     // { name: "post", autobind: true }
	ret.Content = value.vars.post.GetContent() // { name: "post", autobind: true }
	{
		// (grpc.federation.field).custom_resolver = true
		ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx)) // create a new reference to logger.
		var err error
		ret.User, err = s.resolver.Resolve_Federation_V2Dev_PostV2Dev_User(ctx, &Federation_V2Dev_PostV2Dev_UserArgument{
			Federation_V2Dev_PostV2DevArgument: req,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.v2dev.PostV2dev", slog.Any("federation.v2dev.PostV2dev", s.logvalue_Federation_V2Dev_PostV2Dev(ret)))
	return ret, nil
}

// resolve_Federation_V2Dev_Unused resolve "federation.v2dev.Unused" message.
func (s *FederationV2DevService) resolve_Federation_V2Dev_Unused(ctx context.Context, req *Federation_V2Dev_UnusedArgument) (*Unused, error) {
	ctx, span := s.tracer.Start(ctx, "federation.v2dev.Unused")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.v2dev.Unused", slog.Any("message_args", s.logvalue_Federation_V2Dev_UnusedArgument(req)))

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx)) // create a new reference to logger.
	ret, err := s.resolver.Resolve_Federation_V2Dev_Unused(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.v2dev.Unused", slog.Any("federation.v2dev.Unused", s.logvalue_Federation_V2Dev_Unused(ret)))
	return ret, nil
}

// resolve_Federation_V2Dev_User resolve "federation.v2dev.User" message.
func (s *FederationV2DevService) resolve_Federation_V2Dev_User(ctx context.Context, req *Federation_V2Dev_UserArgument) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.v2dev.User")
	defer span.End()

	grpcfed.Logger(ctx).DebugContext(ctx, "resolve federation.v2dev.User", slog.Any("message_args", s.logvalue_Federation_V2Dev_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			res *user.GetUserResponse
			u   *user.User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(ctx, s.celTypeHelper, s.envOpts, s.celPlugins, "grpc.federation.private.UserArgument", req)}
	defer func() {
		if err := value.Close(ctx); err != nil {
			grpcfed.Logger(ctx).ErrorContext(ctx, err.Error())
		}
	}()

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "res"
	     call {
	       method: "user.UserService/GetUser"
	       request { field: "id", by: "$.user_id" }
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.GetUserResponse, *localValueType]{
		Name: `res`,
		Type: grpcfed.CELObjectType("user.GetUserResponse"),
		Setter: func(value *localValueType, v *user.GetUserResponse) error {
			value.vars.res = v
			return nil
		},
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &user.GetUserRequest{}
			// { field: "id", by: "$.user_id" }
			if err := grpcfed.SetCELValue(ctx, &grpcfed.SetCELValueParam[string]{
				Value:             value,
				Expr:              `$.user_id`,
				UseContextLibrary: false,
				CacheIndex:        9,
				Setter: func(v string) error {
					args.Id = v
					return nil
				},
			}); err != nil {
				return nil, err
			}
			grpcfed.Logger(ctx).DebugContext(ctx, "call user.UserService/GetUser", slog.Any("user.GetUserRequest", s.logvalue_User_GetUserRequest(args)))
			return s.client.User_UserServiceClient.GetUser(ctx, args)
		},
	}); err != nil {
		if err := s.errorHandler(ctx, FederationV2DevService_DependentMethod_User_UserService_GetUser, err); err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "u"
	     by: "res.user"
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*user.User, *localValueType]{
		Name: `u`,
		Type: grpcfed.CELObjectType("user.User"),
		Setter: func(value *localValueType, v *user.User) error {
			value.vars.u = v
			return nil
		},
		By:                  `res.user`,
		ByUseContextLibrary: false,
		ByCacheIndex:        10,
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Res = value.vars.res
	req.U = value.vars.u

	// create a message value to be returned.
	// `custom_resolver = true` in "grpc.federation.message" option.
	ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx)) // create a new reference to logger.
	ret, err := s.resolver.Resolve_Federation_V2Dev_User(ctx, req)
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// field binding section.
	{
		// (grpc.federation.field).custom_resolver = true
		ctx = grpcfed.WithLogger(ctx, grpcfed.Logger(ctx)) // create a new reference to logger.
		var err error
		ret.Name, err = s.resolver.Resolve_Federation_V2Dev_User_Name(ctx, &Federation_V2Dev_User_NameArgument{
			Federation_V2Dev_UserArgument: req,
			Federation_V2Dev_User:         ret,
		})
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
	}

	grpcfed.Logger(ctx).DebugContext(ctx, "resolved federation.v2dev.User", slog.Any("federation.v2dev.User", s.logvalue_Federation_V2Dev_User(ret)))
	return ret, nil
}

// cast_int64__to__Federation_V2Dev_PostV2DevType cast from "int64" to "federation.v2dev.PostV2devType".
func (s *FederationV2DevService) cast_int64__to__Federation_V2Dev_PostV2DevType(from int64) (PostV2DevType, error) {
	return PostV2DevType(from), nil
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_ForNameless(v *ForNameless) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("bar", v.GetBar()),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_ForNamelessArgument(v *Federation_V2Dev_ForNamelessArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("bar", v.Bar),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_GetPostV2DevResponse(v *GetPostV2DevResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_V2Dev_PostV2Dev(v.GetPost())),
		slog.String("type", s.logvalue_Federation_V2Dev_PostV2DevType(v.GetType()).String()),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_GetPostV2DevResponseArgument(v *Federation_V2Dev_GetPostV2DevResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_PostV2Dev(v *PostV2Dev) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Federation_V2Dev_User(v.GetUser())),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_PostV2DevArgument(v *Federation_V2Dev_PostV2DevArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_PostV2DevType(v PostV2DevType) slog.Value {
	switch v {
	case PostV2DevType_POST_V2_DEV_TYPE:
		return slog.StringValue("POST_V2_DEV_TYPE")
	}
	return slog.StringValue("")
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_Unused(v *Unused) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("foo", v.GetFoo()),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_UnusedArgument(v *Federation_V2Dev_UnusedArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("foo", v.Foo),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *FederationV2DevService) logvalue_Federation_V2Dev_UserArgument(v *Federation_V2Dev_UserArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("title", v.Title),
		slog.String("content", v.Content),
		slog.String("user_id", v.UserId),
	)
}

func (s *FederationV2DevService) logvalue_Post_GetPostRequest(v *post.GetPostRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
	)
}

func (s *FederationV2DevService) logvalue_Post_GetPostsRequest(v *post.GetPostsRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.GetIds()),
	)
}

func (s *FederationV2DevService) logvalue_User_GetUserRequest(v *user.GetUserRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
	)
}

func (s *FederationV2DevService) logvalue_User_GetUsersRequest(v *user.GetUsersRequest) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("ids", v.GetIds()),
	)
}
