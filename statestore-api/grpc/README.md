# Schemas from .proto files

Generated using: https://github.com/solo-io/protoc-gen-openapi/ (similar to https://github.com/kollalabs/protoc-gen-openapi) which is a protoc plugin that export OpenAPI format. 

To generate this file I needed to bundle all .proto files together with it's dependencies. 

I started with: `dapr/proto/components/v1/state.proto` which depends on `dapr/proto/components/v1/common.proto` and `google/protobuf/any.proto`.





