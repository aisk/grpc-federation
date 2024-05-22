package testutil

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/mercari/grpc-federation/resolver"
)

func ResolverCmpOpts() []cmp.Option {
	return []cmp.Option{
		cmpopts.IgnoreUnexported(resolver.VariableDefinition{}),
		cmpopts.IgnoreFields(resolver.File{}, "Messages", "Services", "Enums", "Desc", "CELPlugins", "ImportFiles"),
		cmpopts.IgnoreFields(resolver.Service{}, "CELPlugins"),
		cmpopts.IgnoreFields(resolver.Package{}, "Files"),
		cmpopts.IgnoreFields(resolver.Method{}, "Service"),
		cmpopts.IgnoreFields(resolver.Message{}, "File", "ParentMessage"),
		cmpopts.IgnoreFields(resolver.Enum{}, "File", "Message.Rule"),
		cmpopts.IgnoreFields(resolver.EnumValue{}, "Enum"),
		cmpopts.IgnoreFields(resolver.EnumRule{}, "Aliases"),
		cmpopts.IgnoreFields(resolver.MessageRule{}, "Aliases"),
		cmpopts.IgnoreFields(resolver.MessageDependencyGraph{}, "Roots"),
		cmpopts.IgnoreFields(resolver.SequentialVariableDefinitionGroup{}, "End.Idx", "End.If", "End.AutoBind", "End.Used", "End.Expr"),
		cmpopts.IgnoreFields(resolver.ConcurrentVariableDefinitionGroup{}, "End.Idx", "End.If", "End.AutoBind", "End.Used", "End.Expr"),
		cmpopts.IgnoreFields(resolver.MessageDependencyGraphNode{}, "BaseMessage", "VariableDefinition", "Parent", "ParentMap", "Children", "ChildrenMap"),
		cmpopts.IgnoreFields(resolver.AllMessageDependencyGraph{}),
		cmpopts.IgnoreFields(resolver.AllMessageDependencyGraphNode{}, "Parent", "Children", "Message.Rule"),
		cmpopts.IgnoreFields(resolver.GRPCError{}, "DefSet.Graph", "DefSet.Groups"),
		cmpopts.IgnoreFields(resolver.GRPCErrorDetail{}, "DefSet.Graph", "DefSet.Groups", "Messages.Graph", "Messages.Groups"),
		cmpopts.IgnoreFields(resolver.AutoBindField{}, "VariableDefinition"),
		cmpopts.IgnoreFields(resolver.Type{}, "Message.Rule", "Enum.Rule", "OneofField"),
		cmpopts.IgnoreFields(resolver.Oneof{}, "Message"),
		cmpopts.IgnoreFields(resolver.Field{}, "Message", "Oneof.Message", "Oneof.Fields"),
		cmpopts.IgnoreFields(resolver.Value{}, "CEL", "Const"),
		cmpopts.IgnoreFields(resolver.CELValue{}, "CheckedExpr"),
		cmpopts.IgnoreFields(resolver.MessageExpr{}, "Message.Rule", "Message.Fields"),
	}
}
