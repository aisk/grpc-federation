package generator_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"

	"github.com/mercari/grpc-federation/grpc/federation/generator"
	"github.com/mercari/grpc-federation/internal/testutil"
	"github.com/mercari/grpc-federation/resolver"
)

func TestRoundTrip(t *testing.T) {
	tests := []string{
		"simple_aggregation",
		"minimum",
		"create_post",
		"custom_resolver",
		"async",
		"alias",
		"autobind",
		"const_value",
		"multi_user",
		"resolver_overlaps",
		"oneof",
		"validation",
		"map",
		"condition",
	}
	for _, test := range tests {
		files := testutil.Compile(t, filepath.Join(testutil.RepoRoot(), "testdata", fmt.Sprintf("%s.proto", test)))
		r := resolver.New(files)
		result, err := r.Resolve()
		if err != nil {
			t.Fatal(err)
		}
		file := generator.ResolverFilesToCodeGeneratorRequest(result.Files)
		req, err := proto.Marshal(file)
		if err != nil {
			t.Fatal(err)
		}
		decoded, err := generator.ToCodeGeneratorRequest(req)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(
			decoded.GRPCFederationFiles,
			result.Files,
			testutil.ResolverCmpOpts()...,
		); diff != "" {
			t.Errorf("(-got, +want)\n%s", diff)
		}
	}
}
