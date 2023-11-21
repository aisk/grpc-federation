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

	"github.com/cenkalti/backoff/v4"
	"github.com/google/cel-go/cel"
	celtypes "github.com/google/cel-go/common/types"
	grpcfed "github.com/mercari/grpc-federation/grpc/federation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/anypb"

	post "example/post"
	user "example/user"
)

// Federation_GetPostResponseArgument is argument for "federation.GetPostResponse" message.
type Federation_GetPostResponseArgument[T any] struct {
	Id     string
	Post   *Post
	Client T
}

// Federation_PostArgument is argument for "federation.Post" message.
type Federation_PostArgument[T any] struct {
	Id     string
	Post   *post.Post
	User   *User
	Client T
}

// Federation_UserArgument is argument for "federation.User" message.
type Federation_UserArgument[T any] struct {
	Content string
	Id      string
	Title   string
	User    *user.User
	UserId  string
	Client  T
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
	// Post_PostServiceClient create a gRPC Client to be used to call methods in post.PostService.
	Post_PostServiceClient(FederationServiceClientConfig) (post.PostServiceClient, error)
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
	Post_PostServiceClient post.PostServiceClient
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
	FederationService_DependentMethod_Post_PostService_GetPost = "/post.PostService/GetPost"
	FederationService_DependentMethod_User_UserService_GetUser = "/user.UserService/GetUser"
)

// FederationService represents Federation Service.
type FederationService struct {
	*UnimplementedFederationServiceServer
	cfg          FederationServiceConfig
	logger       *slog.Logger
	errorHandler grpcfed.ErrorHandler
	env          *cel.Env
	tracer       trace.Tracer
	client       *FederationServiceDependentClientSet
}

// NewFederationService creates FederationService instance by FederationServiceConfig.
func NewFederationService(cfg FederationServiceConfig) (*FederationService, error) {
	if cfg.Client == nil {
		return nil, fmt.Errorf("Client field in FederationServiceConfig is not set. this field must be set")
	}
	Post_PostServiceClient, err := cfg.Client.Post_PostServiceClient(FederationServiceClientConfig{
		Service: "post.PostService",
		Name:    "post_service",
	})
	if err != nil {
		return nil, err
	}
	User_UserServiceClient, err := cfg.Client.User_UserServiceClient(FederationServiceClientConfig{
		Service: "user.UserService",
		Name:    "user_service",
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
		"grpc.federation.private.GetPostResponseArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.PostArgument": {
			"id": grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
		},
		"grpc.federation.private.UserArgument": {
			"id":      grpcfed.NewCELFieldType(celtypes.StringType, "Id"),
			"title":   grpcfed.NewCELFieldType(celtypes.StringType, "Title"),
			"content": grpcfed.NewCELFieldType(celtypes.StringType, "Content"),
			"user_id": grpcfed.NewCELFieldType(celtypes.StringType, "UserId"),
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
		tracer:       otel.Tracer("federation.FederationService"),
		client: &FederationServiceDependentClientSet{
			Post_PostServiceClient: Post_PostServiceClient,
			User_UserServiceClient: User_UserServiceClient,
		},
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
	res, err := s.resolve_Federation_GetPostResponse(ctx, &Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]{
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

// resolve_Federation_GetPostResponse resolve "federation.GetPostResponse" message.
func (s *FederationService) resolve_Federation_GetPostResponse(ctx context.Context, req *Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) (*GetPostResponse, error) {
	ctx, span := s.tracer.Start(ctx, "federation.GetPostResponse")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.GetPostResponse", slog.Any("message_args", s.logvalue_Federation_GetPostResponseArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *Post
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.GetPostResponseArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "post"
	     message: "Post"
	     args { name: "id", by: "$.id" }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("post_federation.Post", func() (any, error) {
			valueMu.RLock()
			args := &Federation_PostArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { name: "id", by: "$.id" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Id = value.(string)
			}
			valueMu.RUnlock()
			return s.resolve_Federation_Post(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*Post)
		valueMu.Lock()
		valuePost = value // { name: "post", message: "Post" ... }
		envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("federation.Post")))
		evalValues["post"] = valuePost
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost

	// create a message value to be returned.
	ret := &GetPostResponse{}

	// field binding section.
	// (grpc.federation.field).by = "post"
	{
		value, err := grpcfed.EvalCEL(s.env, "post", envOpts, evalValues, reflect.TypeOf((*Post)(nil)))
		if err != nil {
			grpcfed.RecordErrorToSpan(ctx, err)
			return nil, err
		}
		ret.Post = value.(*Post)
	}
	ret.Str = "hello" // (grpc.federation.field).string = "hello"

	s.logger.DebugContext(ctx, "resolved federation.GetPostResponse", slog.Any("federation.GetPostResponse", s.logvalue_Federation_GetPostResponse(ret)))
	return ret, nil
}

// resolve_Federation_Post resolve "federation.Post" message.
func (s *FederationService) resolve_Federation_Post(ctx context.Context, req *Federation_PostArgument[*FederationServiceDependentClientSet]) (*Post, error) {
	ctx, span := s.tracer.Start(ctx, "federation.Post")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.Post", slog.Any("message_args", s.logvalue_Federation_PostArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valuePost *post.Post
		valueUser *User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.PostArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "post.PostService/GetPost"
	     request { field: "id", by: "$.id" }
	     response { name: "post", field: "post", autobind: true }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("post.PostService/GetPost", func() (any, error) {
			valueMu.RLock()
			args := &post.GetPostRequest{}
			// { field: "id", by: "$.id" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.id", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Id = value.(string)
			}
			valueMu.RUnlock()
			return grpcfed.WithTimeout[post.GetPostResponse](ctx, "post.PostService/GetPost", 10000000000 /* 10s */, func(ctx context.Context) (*post.GetPostResponse, error) {
				var b backoff.BackOff = backoff.NewConstantBackOff(2000000000 /* 2s */)
				b = backoff.WithMaxRetries(b, 3)
				b = backoff.WithContext(b, ctx)
				return grpcfed.WithRetry[post.GetPostResponse](b, func() (*post.GetPostResponse, error) {
					return s.client.Post_PostServiceClient.GetPost(ctx, args)
				})
			})
		})
		if err != nil {
			if err := s.errorHandler(ctx, FederationService_DependentMethod_Post_PostService_GetPost, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
		}
		value := valueIface.(*post.GetPostResponse)
		valueMu.Lock()
		valuePost = value.GetPost() // { name: "post", field: "post", autobind: true }
		envOpts = append(envOpts, cel.Variable("post", cel.ObjectType("post.Post")))
		evalValues["post"] = valuePost
		valueMu.Unlock()
	}

	// This section's codes are generated by the following proto definition.
	/*
	   {
	     name: "user"
	     message: "User"
	     args { inline: "post" }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("user_federation.User", func() (any, error) {
			valueMu.RLock()
			args := &Federation_UserArgument[*FederationServiceDependentClientSet]{
				Client: s.client,
			}
			// { inline: "post" }
			{
				value, err := grpcfed.EvalCEL(s.env, "post", envOpts, evalValues, reflect.TypeOf((*post.Post)(nil)))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				inlineValue := value.(*post.Post)
				args.Id = inlineValue.GetId()
				args.Title = inlineValue.GetTitle()
				args.Content = inlineValue.GetContent()
				args.UserId = inlineValue.GetUserId()
			}
			valueMu.RUnlock()
			return s.resolve_Federation_User(ctx, args)
		})
		if err != nil {
			return nil, err
		}
		value := valueIface.(*User)
		valueMu.Lock()
		valueUser = value // { name: "user", message: "User" ... }
		envOpts = append(envOpts, cel.Variable("user", cel.ObjectType("federation.User")))
		evalValues["user"] = valueUser
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.Post = valuePost
	req.User = valueUser

	// create a message value to be returned.
	ret := &Post{}

	// field binding section.
	ret.Id = valuePost.GetId()           // { name: "post", autobind: true }
	ret.Title = valuePost.GetTitle()     // { name: "post", autobind: true }
	ret.Content = valuePost.GetContent() // { name: "post", autobind: true }
	// (grpc.federation.field).by = "user"
	{
		value, err := grpcfed.EvalCEL(s.env, "user", envOpts, evalValues, reflect.TypeOf((*User)(nil)))
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
func (s *FederationService) resolve_Federation_User(ctx context.Context, req *Federation_UserArgument[*FederationServiceDependentClientSet]) (*User, error) {
	ctx, span := s.tracer.Start(ctx, "federation.User")
	defer span.End()

	s.logger.DebugContext(ctx, "resolve federation.User", slog.Any("message_args", s.logvalue_Federation_UserArgument(req)))
	var (
		sg        singleflight.Group
		valueMu   sync.RWMutex
		valueUser *user.User
	)
	envOpts := []cel.EnvOption{cel.Variable(grpcfed.MessageArgumentVariableName, cel.ObjectType("grpc.federation.private.UserArgument"))}
	evalValues := map[string]any{grpcfed.MessageArgumentVariableName: req}

	// This section's codes are generated by the following proto definition.
	/*
	   resolver: {
	     method: "user.UserService/GetUser"
	     request { field: "id", by: "$.user_id" }
	     response { name: "user", field: "user", autobind: true }
	   }
	*/
	{
		valueIface, err, _ := sg.Do("user.UserService/GetUser", func() (any, error) {
			valueMu.RLock()
			args := &user.GetUserRequest{}
			// { field: "id", by: "$.user_id" }
			{
				value, err := grpcfed.EvalCEL(s.env, "$.user_id", envOpts, evalValues, reflect.TypeOf(""))
				if err != nil {
					grpcfed.RecordErrorToSpan(ctx, err)
					return nil, err
				}
				args.Id = value.(string)
			}
			valueMu.RUnlock()
			return grpcfed.WithTimeout[user.GetUserResponse](ctx, "user.UserService/GetUser", 20000000000 /* 20s */, func(ctx context.Context) (*user.GetUserResponse, error) {
				eb := backoff.NewExponentialBackOff()
				eb.InitialInterval = 1000000000 /* 1s */
				eb.RandomizationFactor = 0.7
				eb.Multiplier = 1.7
				eb.MaxInterval = 30000000000    /* 30s */
				eb.MaxElapsedTime = 20000000000 /* 20s */

				var b backoff.BackOff = eb
				b = backoff.WithMaxRetries(b, 3)
				b = backoff.WithContext(b, ctx)
				return grpcfed.WithRetry[user.GetUserResponse](b, func() (*user.GetUserResponse, error) {
					return s.client.User_UserServiceClient.GetUser(ctx, args)
				})
			})
		})
		if err != nil {
			if err := s.errorHandler(ctx, FederationService_DependentMethod_User_UserService_GetUser, err); err != nil {
				grpcfed.RecordErrorToSpan(ctx, err)
				return nil, err
			}
		}
		value := valueIface.(*user.GetUserResponse)
		valueMu.Lock()
		valueUser = value.GetUser() // { name: "user", field: "user", autobind: true }
		envOpts = append(envOpts, cel.Variable("user", cel.ObjectType("user.User")))
		evalValues["user"] = valueUser
		valueMu.Unlock()
	}

	// assign named parameters to message arguments to pass to the custom resolver.
	req.User = valueUser

	// create a message value to be returned.
	ret := &User{}

	// field binding section.
	ret.Id = valueUser.GetId()                                                                // { name: "user", autobind: true }
	ret.Name = valueUser.GetName()                                                            // { name: "user", autobind: true }
	ret.Items = s.cast_repeated_User_Item__to__repeated_Federation_Item(valueUser.GetItems()) // { name: "user", autobind: true }
	ret.Profile = valueUser.GetProfile()                                                      // { name: "user", autobind: true }

	switch {
	case s.cast_User_User_AttrA___to__Federation_User_AttrA_(valueUser.GetAttrA()) != nil:

		ret.Attr = s.cast_User_User_AttrA___to__Federation_User_AttrA_(valueUser.GetAttrA())
	case s.cast_User_User_B__to__Federation_User_B(valueUser.GetB()) != nil:

		ret.Attr = s.cast_User_User_B__to__Federation_User_B(valueUser.GetB())
	}

	s.logger.DebugContext(ctx, "resolved federation.User", slog.Any("federation.User", s.logvalue_Federation_User(ret)))
	return ret, nil
}

// cast_User_Item_ItemType__to__Federation_Item_ItemType cast from "user.Item.ItemType" to "federation.Item.ItemType".
func (s *FederationService) cast_User_Item_ItemType__to__Federation_Item_ItemType(from user.Item_ItemType) Item_ItemType {
	switch from {
	case user.Item_ITEM_TYPE_UNSPECIFIED:
		return Item_ITEM_TYPE_UNSPECIFIED
	case user.Item_ITEM_TYPE_1:
		return Item_ITEM_TYPE_1
	case user.Item_ITEM_TYPE_2:
		return Item_ITEM_TYPE_2
	case user.Item_ITEM_TYPE_3:
		return Item_ITEM_TYPE_3
	default:
		return 0
	}
}

// cast_User_Item_Location_AddrA___to__Federation_Item_Location_AddrA_ cast from "user.Item.Location.addr_a" to "federation.Item.Location.addr_a".
func (s *FederationService) cast_User_Item_Location_AddrA___to__Federation_Item_Location_AddrA_(from *user.Item_Location_AddrA) *Item_Location_AddrA_ {
	if from == nil {
		return nil
	}
	return &Item_Location_AddrA_{
		AddrA: s.cast_User_Item_Location_AddrA__to__Federation_Item_Location_AddrA(from),
	}
}

// cast_User_Item_Location_AddrA__to__Federation_Item_Location_AddrA cast from "user.Item.Location.AddrA" to "federation.Item.Location.AddrA".
func (s *FederationService) cast_User_Item_Location_AddrA__to__Federation_Item_Location_AddrA(from *user.Item_Location_AddrA) *Item_Location_AddrA {
	if from == nil {
		return nil
	}

	return &Item_Location_AddrA{
		Foo: from.GetFoo(),
	}
}

// cast_User_Item_Location_AddrB__to__Federation_Item_Location_AddrB cast from "user.Item.Location.AddrB" to "federation.Item.Location.AddrB".
func (s *FederationService) cast_User_Item_Location_AddrB__to__Federation_Item_Location_AddrB(from *user.Item_Location_AddrB) *Item_Location_AddrB {
	if from == nil {
		return nil
	}

	return &Item_Location_AddrB{
		Bar: from.GetBar(),
	}
}

// cast_User_Item_Location_B__to__Federation_Item_Location_B cast from "user.Item.Location.b" to "federation.Item.Location.b".
func (s *FederationService) cast_User_Item_Location_B__to__Federation_Item_Location_B(from *user.Item_Location_AddrB) *Item_Location_B {
	if from == nil {
		return nil
	}
	return &Item_Location_B{
		B: s.cast_User_Item_Location_AddrB__to__Federation_Item_Location_AddrB(from),
	}
}

// cast_User_Item_Location__to__Federation_Item_Location cast from "user.Item.Location" to "federation.Item.Location".
func (s *FederationService) cast_User_Item_Location__to__Federation_Item_Location(from *user.Item_Location) *Item_Location {
	if from == nil {
		return nil
	}

	ret := &Item_Location{
		Addr1: from.GetAddr1(),
		Addr2: from.GetAddr2(),
	}
	switch {

	case from.GetAddrA() != nil:
		ret.Addr3 = s.cast_User_Item_Location_AddrA___to__Federation_Item_Location_AddrA_(from.GetAddrA())
	case from.GetB() != nil:
		ret.Addr3 = s.cast_User_Item_Location_B__to__Federation_Item_Location_B(from.GetB())
	}
	return ret
}

// cast_User_Item__to__Federation_Item cast from "user.Item" to "federation.Item".
func (s *FederationService) cast_User_Item__to__Federation_Item(from *user.Item) *Item {
	if from == nil {
		return nil
	}

	return &Item{
		Name:     from.GetName(),
		Type:     s.cast_User_Item_ItemType__to__Federation_Item_ItemType(from.GetType()),
		Value:    from.GetValue(),
		Location: s.cast_User_Item_Location__to__Federation_Item_Location(from.GetLocation()),
	}
}

// cast_User_User_AttrA___to__Federation_User_AttrA_ cast from "user.User.attr_a" to "federation.User.attr_a".
func (s *FederationService) cast_User_User_AttrA___to__Federation_User_AttrA_(from *user.User_AttrA) *User_AttrA_ {
	if from == nil {
		return nil
	}
	return &User_AttrA_{
		AttrA: s.cast_User_User_AttrA__to__Federation_User_AttrA(from),
	}
}

// cast_User_User_AttrA__to__Federation_User_AttrA cast from "user.User.AttrA" to "federation.User.AttrA".
func (s *FederationService) cast_User_User_AttrA__to__Federation_User_AttrA(from *user.User_AttrA) *User_AttrA {
	if from == nil {
		return nil
	}

	return &User_AttrA{
		Foo: from.GetFoo(),
	}
}

// cast_User_User_AttrB__to__Federation_User_AttrB cast from "user.User.AttrB" to "federation.User.AttrB".
func (s *FederationService) cast_User_User_AttrB__to__Federation_User_AttrB(from *user.User_AttrB) *User_AttrB {
	if from == nil {
		return nil
	}

	return &User_AttrB{
		Bar: from.GetBar(),
	}
}

// cast_User_User_B__to__Federation_User_B cast from "user.User.b" to "federation.User.b".
func (s *FederationService) cast_User_User_B__to__Federation_User_B(from *user.User_AttrB) *User_B {
	if from == nil {
		return nil
	}
	return &User_B{
		B: s.cast_User_User_AttrB__to__Federation_User_AttrB(from),
	}
}

// cast_repeated_User_Item__to__repeated_Federation_Item cast from "repeated user.Item" to "repeated federation.Item".
func (s *FederationService) cast_repeated_User_Item__to__repeated_Federation_Item(from []*user.Item) []*Item {
	ret := make([]*Item, 0, len(from))
	for _, v := range from {
		ret = append(ret, s.cast_User_Item__to__Federation_Item(v))
	}
	return ret
}

func (s *FederationService) logvalue_Federation_GetPostResponse(v *GetPostResponse) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Any("post", s.logvalue_Federation_Post(v.GetPost())),
		slog.String("str", v.GetStr()),
	)
}

func (s *FederationService) logvalue_Federation_GetPostResponseArgument(v *Federation_GetPostResponseArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_Item(v *Item) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("name", v.GetName()),
		slog.String("type", s.logvalue_Federation_Item_ItemType(v.GetType()).String()),
		slog.Int64("value", v.GetValue()),
		slog.Any("location", s.logvalue_Federation_Item_Location(v.GetLocation())),
	)
}

func (s *FederationService) logvalue_Federation_Item_ItemType(v Item_ItemType) slog.Value {
	switch v {
	case Item_ITEM_TYPE_UNSPECIFIED:
		return slog.StringValue("ITEM_TYPE_UNSPECIFIED")
	case Item_ITEM_TYPE_1:
		return slog.StringValue("ITEM_TYPE_1")
	case Item_ITEM_TYPE_2:
		return slog.StringValue("ITEM_TYPE_2")
	case Item_ITEM_TYPE_3:
		return slog.StringValue("ITEM_TYPE_3")
	}
	return slog.StringValue("")
}

func (s *FederationService) logvalue_Federation_Item_Location(v *Item_Location) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("addr1", v.GetAddr1()),
		slog.String("addr2", v.GetAddr2()),
		slog.Any("addr_a", s.logvalue_Federation_Item_Location_AddrA(v.GetAddrA())),
		slog.Any("b", s.logvalue_Federation_Item_Location_AddrB(v.GetB())),
	)
}

func (s *FederationService) logvalue_Federation_Item_Location_AddrA(v *Item_Location_AddrA) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("foo", v.GetFoo()),
	)
}

func (s *FederationService) logvalue_Federation_Item_Location_AddrB(v *Item_Location_AddrB) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Int64("bar", v.GetBar()),
	)
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
	)
}

func (s *FederationService) logvalue_Federation_PostArgument(v *Federation_PostArgument[*FederationServiceDependentClientSet]) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.Id),
	)
}

func (s *FederationService) logvalue_Federation_User(v *User) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("id", v.GetId()),
		slog.String("name", v.GetName()),
		slog.Any("items", s.logvalue_repeated_Federation_Item(v.GetItems())),
		slog.Any("profile", s.logvalue_Federation_User_ProfileEntry(v.GetProfile())),
		slog.Any("attr_a", s.logvalue_Federation_User_AttrA(v.GetAttrA())),
		slog.Any("b", s.logvalue_Federation_User_AttrB(v.GetB())),
	)
}

func (s *FederationService) logvalue_Federation_UserArgument(v *Federation_UserArgument[*FederationServiceDependentClientSet]) slog.Value {
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

func (s *FederationService) logvalue_Federation_User_AttrA(v *User_AttrA) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("foo", v.GetFoo()),
	)
}

func (s *FederationService) logvalue_Federation_User_AttrB(v *User_AttrB) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.Bool("bar", v.GetBar()),
	)
}

func (s *FederationService) logvalue_Federation_User_ProfileEntry(v map[string]*anypb.Any) slog.Value {
	attrs := make([]slog.Attr, 0, len(v))
	for key, value := range v {
		attrs = append(attrs, slog.Attr{
			Key:   fmt.Sprint(key),
			Value: s.logvalue_Google_Protobuf_Any(value),
		})
	}
	return slog.GroupValue(attrs...)
}

func (s *FederationService) logvalue_Google_Protobuf_Any(v *anypb.Any) slog.Value {
	if v == nil {
		return slog.GroupValue()
	}
	return slog.GroupValue(
		slog.String("type_url", v.GetTypeUrl()),
		slog.String("value", string(v.GetValue())),
	)
}

func (s *FederationService) logvalue_repeated_Federation_Item(v []*Item) slog.Value {
	attrs := make([]slog.Attr, 0, len(v))
	for idx, vv := range v {
		attrs = append(attrs, slog.Attr{
			Key:   fmt.Sprint(idx),
			Value: s.logvalue_Federation_Item(vv),
		})
	}
	return slog.GroupValue(attrs...)
}
