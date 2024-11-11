# ml-server

# cmd/server
- main.go -> entry point

# internal/api 
- handlers.go -> http handlers for job submission, status, monitoring
- routes.go -> api routes

The api layer accepts job submissions, reports the job status, handles result downloads and manages authentication

# internal/gpu
- monitor.go -> gpu status monitoring
- manager.go -> gpu resource allocation

The GPU management layer monitors GPU status, allocates GPU memory, prevents overallocation and tracks GPU health

# internal/queue
- queue.go -> job queue implementation
- job.go -> job type definition

The Job Queue maintains job order, tracks job status, handles job priority and manages the job metadata.

# internal/runner
- executor.go -> code execution and process management
- sandbox.go -> training env set up

The Job Runner sets up the training enivornment, executes the training code, and monitors training process while handling handling failures.

# internal/storage
- results.go -> training results 
- files.go -> file system ops

The storage saves the training results, manages temporary files and handles cleanup/

# configs/
- config.yaml -> server config

The yaml file configs the server at a defined ip

go.mod

