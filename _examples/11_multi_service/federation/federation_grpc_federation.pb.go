// Code generated by protoc-gen-grpc-federation. DO NOT EDIT!
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

	comment "example/comment"
	favorite "example/favorite"
)

var (
	_ = reflect.Invalid // to avoid "imported and not used error"
)

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument struct {
	Id string
	P  *Post
}

// Federation_GetStatusResponseArgument is argument for "federation.GetStatusResponse" message.
type Federation_GetStatusResponseArgument struct {
	U *User
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument struct {
	FavoriteValue favorite.FavoriteType
	Reaction      *Reaction
	U             *User
}

// Federation_UserArgument is argument for "federation.User" message.
type Federation_UserArgument struct {
	Id   string
	Name string
}

// FederationServiceConfig configuration required to initialize the service that use GRPC Federation.
type FederationServiceConfig struct {
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// FederationServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type FederationServiceClientFactory interface {
}

// FederationServiceClientConfig helper to create gRPC client.
// Hints for creating a gRPC Client.
type FederationServiceClientConfig struct {
	// Service FQDN ( `<package-name>.<service-name>` ) of the service on Protocol Buffers.
	Service string
}

// FederationServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type FederationServiceDependentClientSet struct {
}

// FederationServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type FederationServiceResolver interface {
}

// FederationServiceCELPluginWasmConfig type alias for grpcfedcel.WasmConfig.
type FederationServiceCELPluginWasmConfig = grpcfedcel.WasmConfig

// FederationServiceCELPluginConfig hints for loading a WebAssembly based plugin.
type FederationServiceCELPluginConfig struct {
}

// FederationServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type FederationServiceUnimplementedResolver struct{}

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *grpcfed.CELEnv
	tracer       trace.Tracer
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*grpcfed.CELFieldType{
		"grpc.federation.private.GetPostResponseArgument": {
			"id": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {},
		"grpc.federation.private.ReactionArgument": {
			"v": grpcfed.NewCELFieldType(grpcfed.CELIntType, "V"),
		},
		"grpc.federation.private.UserArgument": {
			"id":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
			"name": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Name"),
		},
	})
	envOpts := grpcfed.NewDefaultEnvOptions(celHelper)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("comment.CommentType", comment.CommentType_value, comment.CommentType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("favorite.FavoriteType", favorite.FavoriteType_value, favorite.FavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("favorite.FavoriteType", favorite.FavoriteType_value, favorite.FavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	env, err := grpcfed.NewCELEnv(envOpts...)
	if err != nil {
		return nil, err
	}
	return &FederationService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("federation.FederationService"),
		client:       &FederationServiceDependentClientSet{},
	}, nil
}

// GetPost implements "federation.FederationService/GetPost" method.
func (s *FederationService) GetPost(ctx context.Context, req *GetPostRequest) (res *GetPostResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.FederationService/GetPost")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument{
		Id: req.Id,
	})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetPostResponse resolve "federation.GetPostResponse" message.
func (s *FederationService) resolve_Federation_GetPostResponse(ctx context.Context, req *Federation_GetPostResponseArgument) (*GetPostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetPostResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.GetPostResponse", slog.Any("message_args", s.logvalue_Federation_GetPostResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			p *Post
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.GetPostResponseArgument", req)}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "p"
	     message {
	       name: "Post"
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*Post, *localValueType]{
		Name:   "p",
		Type:   grpcfed.CELObjectType("federation.Post"),
		Setter: func(value *localValueType, v *Post) { value.vars.p = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Federation_PostArgument{}
			return s.resolve_Federation_Post(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.P = value.vars.p

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "p"
	if err := grpcfed.SetCELValue(ctx, value, "p", func(v *Post) { ret.Post = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.GetPostResponse", slog.Any("federation.GetPostResponse", s.logvalue_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *FederationService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Post")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			favorite_value favorite.FavoriteType
			reaction       *Reaction
			u              *User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.PostArgument", req)}
	// A tree view of message dependencies is shown below.
	/*
	   favorite_value ─┐
	                   reaction ─┐
	                          u ─┤
	*/
	eg, ctx1 := grpcfed.ErrorGroupWithContext(ctx)

	grpcfed.GoWithRecover(eg, func() (any, error) {

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "favorite_value"
		     by: "favorite.FavoriteType.value('TYPE1')"
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[favorite.FavoriteType, *localValueType]{
			Name:   "favorite_value",
			Type:   grpcfed.CELIntType,
			Setter: func(value *localValueType, v favorite.FavoriteType) { value.vars.favorite_value = v },
			By:     "favorite.FavoriteType.value('TYPE1')",
		}); err != nil {
			grpcfed.RecordErrorToSpan(ctx1, err)
			return nil, err
		}

		// This section's codes are generated by the following proto definition.
		/*
		   def {
		     name: "reaction"
		     message {
		       name: "Reaction"
		       args { name: "v", by: "favorite_value" }
		     }
		   }
		*/
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*Reaction, *localValueType]{
			Name:   "reaction",
			Type:   grpcfed.CELObjectType("federation.Reaction"),
			Setter: func(value *localValueType, v *Reaction) { value.vars.reaction = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_ReactionArgument{}
				// { name: "v", by: "favorite_value" }
				if err := grpcfed.SetCELValue(ctx, value, "favorite_value", func(v favorite.FavoriteType) {
					args.V = v
				}); err != nil {
					return nil, err
				}
				return s.resolve_Federation_Reaction(ctx, args)
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
		if err := grpcfed.EvalDef(ctx1, value, grpcfed.Def[*User, *localValueType]{
			Name:   "u",
			Type:   grpcfed.CELObjectType("federation.User"),
			Setter: func(value *localValueType, v *User) { value.vars.u = v },
			Message: func(ctx context.Context, value *localValueType) (any, error) {
				args := &Federation_UserArgument{
					Id:   "foo", // { name: "id", string: "foo" }
					Name: "bar", // { name: "name", string: "bar" }
				}
				return s.resolve_Federation_User(ctx, args)
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
	req.FavoriteValue = value.vars.favorite_value
	req.Reaction = value.vars.reaction
	req.U = value.vars.u

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = "post-id"      // (grpc.federation.field).string = "post-id"
	ret.Title = "title"     // (grpc.federation.field).string = "title"
	ret.Content = "content" // (grpc.federation.field).string = "content"
	// (grpc.federation.field).by = "u"
	if err := grpcfed.SetCELValue(ctx, value, "u", func(v *User) { ret.User = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "reaction"
	if err := grpcfed.SetCELValue(ctx, value, "reaction", func(v *Reaction) { ret.Reaction = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "favorite_value"
	if err := grpcfed.SetCELValue(ctx, value, "favorite_value", func(v favorite.FavoriteType) {
		ret.FavoriteValue = s.cast_Favorite_FavoriteType__to__Federation_MyFavoriteType(v)
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.Post", slog.Any("federation.Post", s.logvalue_Federation_Post(ret)))
	return ret, nil
}

// resolve_Federation_Reaction resolve "federation.Reaction" message.
func (s *FederationService) resolve_Federation_Reaction(ctx context.Context, req *Federation_ReactionArgument) (*Reaction, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Reaction")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.Reaction", slog.Any("message_args", s.logvalue_Federation_ReactionArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.ReactionArgument", req)}

	// create a message value to be returned.
	ret := &Reaction{}

	// field binding section.
	// (grpc.federation.field).by = "favorite.FavoriteType.TYPE1"
	if err := grpcfed.SetCELValue(ctx, value, "favorite.FavoriteType.TYPE1", func(v favorite.FavoriteType) { ret.FavoriteType = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.Reaction", slog.Any("federation.Reaction", s.logvalue_Federation_Reaction(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *FederationService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.UserArgument", req)}

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.id"
	if err := grpcfed.SetCELValue(ctx, value, "$.id", func(v string) { ret.Id = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "$.name"
	if err := grpcfed.SetCELValue(ctx, value, "$.name", func(v string) { ret.Name = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

// cast_Favorite_FavoriteType__to__Federation_MyFavoriteType cast from "favorite.FavoriteType" to "federation.MyFavoriteType".
func (s *FederationService) cast_Favorite_FavoriteType__to__Federation_MyFavoriteType(from favorite.FavoriteType) MyFavoriteType {
	switch from {
	case favorite.FavoriteType_UNKNOWN:
		return MyFavoriteType_UNKNOWN
	case favorite.FavoriteType_TYPE1:
		return MyFavoriteType_TYPE1
	default:
		return 0
	}
}

// cast_int64__to__Favorite_FavoriteType cast from "int64" to "favorite.FavoriteType".
func (s *FederationService) cast_int64__to__Favorite_FavoriteType(from int64) favorite.FavoriteType {
	return favorite.FavoriteType(from)
}

func (s *FederationService) logvalue_Favorite_FavoriteType(v favorite.FavoriteType) slog.Value {
	switch v {
	case favorite.FavoriteType_UNKNOWN:
		return slog.StringValue("UNKNOWN")
	case favorite.FavoriteType_TYPE1:
		return slog.StringValue("TYPE1")
	case favorite.FavoriteType_TYPE2:
		return slog.StringValue("TYPE2")
	}
	return slog.StringValue("")
}

func (s *FederationService) logvalue_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponseArgument(v *Federation_GetPostResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_MyFavoriteType(v MyFavoriteType) slog.Value {
	switch v {
	case MyFavoriteType_UNKNOWN:
		return slog.StringValue("UNKNOWN")
	case MyFavoriteType_TYPE1:
		return slog.StringValue("TYPE1")
	}
	return slog.StringValue("")
}

func (s *FederationService) logvalue_Federation_Post(v *Post) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("title", v.GetTitle()),
		slog.String("content", v.GetContent()),
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
		slog.Any("reaction", s.logvalue_Federation_Reaction(v.GetReaction())),
		slog.String("favorite_value", s.logvalue_Federation_MyFavoriteType(v.GetFavoriteValue()).String()),
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *FederationService) logvalue_Federation_Reaction(v *Reaction) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("favorite_type", s.logvalue_Favorite_FavoriteType(v.GetFavoriteType()).String()),
	)
}

func (s *FederationService) logvalue_Federation_ReactionArgument(v *Federation_ReactionArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("v", s.logvalue_Favorite_FavoriteType(v.V).String()),
	)
}

func (s *FederationService) logvalue_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *FederationService) logvalue_Federation_UserArgument(v *Federation_UserArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("name", v.Name),
	)
}

// DebugServiceConfig configuration required to initialize the service that use GRPC Federation.
type DebugServiceConfig struct {
	// ErrorHandler Federation Service often needs to convert errors received from downstream services.
	// If an error occurs during method execution in the Federation Service, this error handler is called and the returned error is treated as a final error.
	ErrorHandler grpcfed.ErrorHandler
	// Logger sets the logger used to output Debug/Info/Error information.
	Logger *slog.Logger
}

// DebugServiceClientFactory provides a factory that creates the gRPC Client needed to invoke methods of the gRPC Service on which the Federation Service depends.
type DebugServiceClientFactory interface {
}

// DebugServiceClientConfig helper to create gRPC client.
// Hints for creating a gRPC Client.
type DebugServiceClientConfig struct {
	// Service FQDN ( `<package-name>.<service-name>` ) of the service on Protocol Buffers.
	Service string
}

// DebugServiceDependentClientSet has a gRPC client for all services on which the federation service depends.
// This is provided as an argument when implementing the custom resolver.
type DebugServiceDependentClientSet struct {
}

// DebugServiceResolver provides an interface to directly implement message resolver and field resolver not defined in Protocol Buffers.
type DebugServiceResolver interface {
}

// DebugServiceCELPluginWasmConfig type alias for grpcfedcel.WasmConfig.
type DebugServiceCELPluginWasmConfig = grpcfedcel.WasmConfig

// DebugServiceCELPluginConfig hints for loading a WebAssembly based plugin.
type DebugServiceCELPluginConfig struct {
}

// DebugServiceUnimplementedResolver a structure implemented to satisfy the Resolver interface.
// An Unimplemented error is always returned.
// This is intended for use when there are many Resolver interfaces that do not need to be implemented,
// by embedding them in a resolver structure that you have created.
type DebugServiceUnimplementedResolver struct{}

// DebugService represents Federation Service.
type DebugService struct {
	*UnimplementedDebugServiceServer
	cfg          DebugServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *grpcfed.CELEnv
	tracer       trace.Tracer
	client       *DebugServiceDependentClientSet
}

// NewDebugService creates DebugService instance by DebugServiceConfig.
func NewDebugService(cfg DebugServiceConfig) (*DebugService, error) {
	logger := cfg.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	errorHandler := cfg.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(ctx context.Context, methodName string, err error) error { return err }
	}
	celHelper := grpcfed.NewCELTypeHelper(map[string]map[string]*grpcfed.CELFieldType{
		"grpc.federation.private.GetStatusResponseArgument": {},
		"grpc.federation.private.UserArgument": {
			"id":   grpcfed.NewCELFieldType(grpcfed.CELStringType, "Id"),
			"name": grpcfed.NewCELFieldType(grpcfed.CELStringType, "Name"),
		},
	})
	envOpts := grpcfed.NewDefaultEnvOptions(celHelper)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("comment.CommentType", comment.CommentType_value, comment.CommentType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("favorite.FavoriteType", favorite.FavoriteType_value, favorite.FavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("favorite.FavoriteType", favorite.FavoriteType_value, favorite.FavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	envOpts = append(envOpts, grpcfed.EnumAccessorOptions("federation.MyFavoriteType", MyFavoriteType_value, MyFavoriteType_name)...)
	env, err := grpcfed.NewCELEnv(envOpts...)
	if err != nil {
		return nil, err
	}
	return &DebugService{
		cfg:          cfg,
		logger:       logger,
		errorHandler: errorHandler,
		env:          env,
		tracer:       otel.Tracer("federation.DebugService"),
		client:       &DebugServiceDependentClientSet{},
	}, nil
}

// GetStatus implements "federation.DebugService/GetStatus" method.
func (s *DebugService) GetStatus(ctx context.Context, req *GetStatusRequest) (res *GetStatusResponse, e error) {
	ctx, span := s.tracer.Start(ctx, "federation.DebugService/GetStatus")
	defer span.End()

	ctx = grpcfed.WithLogger(ctx, s.logger)
	defer func() {
		if r := recover(); r != nil {
			e = grpcfed.RecoverError(r, debug.Stack())
			grpcfed.OutputErrorLog(ctx, s.logger, e)
		}
	}()
	res, err := s.resolve_Federation_GetStatusResponse(ctx, &Federation_GetStatusResponseArgument{})
	if err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		grpcfed.OutputErrorLog(ctx, s.logger, err)
		return nil, err
	}
	return res, nil
}

// resolve_Federation_GetStatusResponse resolve "federation.GetStatusResponse" message.
func (s *DebugService) resolve_Federation_GetStatusResponse(ctx context.Context, req *Federation_GetStatusResponseArgument) (*GetStatusResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetStatusResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.GetStatusResponse", slog.Any("message_args", s.logvalue_Federation_GetStatusResponseArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
			u *User
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.GetStatusResponseArgument", req)}

	// This section's codes are generated by the following proto definition.
	/*
	   def {
	     name: "u"
	     message {
	       name: "User"
	       args: [
	         { name: "id", string: "xxxx" },
	         { name: "name", string: "yyyy" }
	       ]
	     }
	   }
	*/
	if err := grpcfed.EvalDef(ctx, value, grpcfed.Def[*User, *localValueType]{
		Name:   "u",
		Type:   grpcfed.CELObjectType("federation.User"),
		Setter: func(value *localValueType, v *User) { value.vars.u = v },
		Message: func(ctx context.Context, value *localValueType) (any, error) {
			args := &Federation_UserArgument{
				Id:   "xxxx", // { name: "id", string: "xxxx" }
				Name: "yyyy", // { name: "name", string: "yyyy" }
			}
			return s.resolve_Federation_User(ctx, args)
		},
	}); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.U = value.vars.u

	// create a message value to be returned.
	ret := &GetStatusResponse{}

	// field binding section.
	// (grpc.federation.field).by = "u"
	if err := grpcfed.SetCELValue(ctx, value, "u", func(v *User) { ret.User = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.GetStatusResponse", slog.Any("federation.GetStatusResponse", s.logvalue_Federation_GetStatusResponse(ret)))
	return ret, nil
}

// resolve_Federation_User resolve "federation.User" message.
func (s *DebugService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	type localValueType struct {
		*grpcfed.LocalValue
		vars struct {
		}
	}
	value := &localValueType{LocalValue: grpcfed.NewLocalValue(s.env, "grpc.federation.private.UserArgument", req)}

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	// (grpc.federation.field).by = "$.id"
	if err := grpcfed.SetCELValue(ctx, value, "$.id", func(v string) { ret.Id = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}
	// (grpc.federation.field).by = "$.name"
	if err := grpcfed.SetCELValue(ctx, value, "$.name", func(v string) { ret.Name = v }); err != nil {
		grpcfed.RecordErrorToSpan(ctx, err)
		return nil, err
	}

	s.logger.DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

func (s *DebugService) logvalue_Federation_GetStatusResponse(v *GetStatusResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("user", s.logvalue_Federation_User(v.GetUser())),
	)
}

func (s *DebugService) logvalue_Federation_GetStatusResponseArgument(v *Federation_GetStatusResponseArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue()
}

func (s *DebugService) logvalue_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
	)
}

func (s *DebugService) logvalue_Federation_UserArgument(v *Federation_UserArgument) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
		slog.String("name", v.Name),
	)
}
