# Root

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#RootGetResponse">RootGetResponse</a>

Methods:

- <code title="get /">client.Root.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#RootService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#RootGetResponse">RootGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelInfo">ModelInfo</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelsResponse">ModelsResponse</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelInfo">ModelInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelsResponse">ModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CompletionRequestParam">CompletionRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Completion">Completion</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#StreamChunk">StreamChunk</a>

Methods:

- <code title="post /v1/chat">client.Chat.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatNewParams">ChatNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Completion">Completion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
