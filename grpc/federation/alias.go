package federation

// Aliases list of types or functions from third-party libraries to minimize the list of imported packages.

import (
	"context"
	"os"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/protoadapt"
)

type (
	ErrorGroup     = errgroup.Group
	ProtoMessage   = protoadapt.MessageV1
	CELTypeDeclare = cel.Type
	CELEnv         = cel.Env
	CELEnvOption   = cel.EnvOption
	CELFieldType   = types.FieldType
	Code           = codes.Code
	RWMutex        = sync.RWMutex
	Status         = status.Status
)

var (
	Getenv                = os.Getenv
	GRPCErrorf            = status.Errorf
	NewGRPCStatus         = status.New
	ErrorGroupWithContext = errgroup.WithContext
	NewCELEnv             = cel.NewCustomEnv
	CELLib                = cel.Lib
	CELDoubleType         = types.DoubleType
	CELIntType            = types.IntType
	CELUintType           = types.UintType
	CELBoolType           = types.BoolType
	CELStringType         = types.StringType
	CELBytesType          = types.BytesType
	CELObjectType         = cel.ObjectType
	CELListType           = cel.ListType
	CELMapType            = cel.MapType
	NewCELListType        = types.NewListType
	NewCELObjectType      = types.NewObjectType
)

const (
	OKCode Code = codes.OK

	// CancelledCode this is not a typo.
	// On proto, it is defined as `CANCELLED`, but in Go's library it is defined as Canceled.
	// The name of the code is automatically generated by (*GRPCError).GoGRPCStatusCode in generator/code_generator.go.
	// So if the name is different, it will not work. Therefore, the names should be the same as the names on the proto side.
	// - Proto: https://github.com/googleapis/googleapis/blob/65277ddce9caa1cfd1a0eb7ab67980fc73d20b50/google/rpc/code.proto#L41
	// - Go: https://github.com/grpc/grpc-go/blob/9952aa83979822b5915c3fcb2bb0f60afe55aa7d/codes/codes.go#L45
	//nolint:misspell
	CancelledCode Code = codes.Canceled

	UnknownCode            Code = codes.Unknown
	InvalidArgumentCode    Code = codes.InvalidArgument
	DeadlineExceededCode   Code = codes.DeadlineExceeded
	NotFoundCode           Code = codes.NotFound
	AlreadyExistsCode      Code = codes.AlreadyExists
	PermissionDeniedCode   Code = codes.PermissionDenied
	ResourceExhaustedCode  Code = codes.ResourceExhausted
	FailedPreconditionCode Code = codes.FailedPrecondition
	AbortedCode            Code = codes.Aborted
	OutOfRangeCode         Code = codes.OutOfRange
	UnimplementedCode      Code = codes.Unimplemented
	InternalCode           Code = codes.Internal
	UnavailableCode        Code = codes.Unavailable
	DataLossCode           Code = codes.DataLoss
	UnauthenticatedCode    Code = codes.Unauthenticated
)

func BackOffWithMaxRetries(b *BackOff, max uint64) *BackOff {
	return &BackOff{
		BackOff: backoff.WithMaxRetries(b, max),
	}
}

func BackOffWithContext(b *BackOff, ctx context.Context) *BackOff {
	return &BackOff{
		BackOff: backoff.WithContext(b, ctx),
	}
}
