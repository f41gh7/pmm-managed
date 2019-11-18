// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/postgresql.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddPostgreSQLRequest) Validate() error {
	if this.AddNode != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AddNode); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AddNode", err)
		}
	}
	if this.ServiceName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ServiceName", fmt.Errorf(`value '%v' must not be an empty string`, this.ServiceName))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	if this.PmmAgentId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PmmAgentId", fmt.Errorf(`value '%v' must not be an empty string`, this.PmmAgentId))
	}
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddPostgreSQLResponse) Validate() error {
	if this.Service != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Service); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Service", err)
		}
	}
	if this.PostgresExporter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PostgresExporter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PostgresExporter", err)
		}
	}
	if this.QanPostgresqlPgstatementsAgent != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.QanPostgresqlPgstatementsAgent); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("QanPostgresqlPgstatementsAgent", err)
		}
	}
	return nil
}