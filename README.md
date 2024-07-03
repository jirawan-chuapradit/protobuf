First step: create a file .proto
  - Define message
  - Define a service
    - with RPC method
Second step: generate the Go code that will be used (with protoc)
  - To encode/decode messages with protocol Buffers
  - To handle incoming gRPC requests

- Each time you change your proto file then you need to regenerate the code
- Do not touch the generated code (your changes will be removed at next generation)
  
