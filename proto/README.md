# Protocol Buffers

This directory contains the Protocol Buffer definitions for the CloudTV API.

## Structure

```
proto/
├── auth.proto          # Authentication service definitions
├── video.proto         # Video service definitions (future)
└── stream.proto        # Streaming service definitions (future)
```

## Generating Code

### Go (Backend)

```bash
cd backend
make proto
```

Generated files will be placed in `backend/pb/`

### TypeScript (Frontend - Future)

```bash
cd frontend
npm run proto:gen
```

## Adding New Services

1. Create a new `.proto` file in this directory
2. Follow the naming convention: `service_name.proto`
3. Set the package name and go_package option:
   ```protobuf
   syntax = "proto3";
   package service_name;
   option go_package = "github.com/abdulyazidi/cloudtv/backend/pb/service_name";
   ```
4. Run `make proto` in the backend directory

## Style Guide

- Use `snake_case` for field names
- Use `PascalCase` for message and service names
- Always specify `syntax = "proto3"`
- Include clear comments for all services and messages
