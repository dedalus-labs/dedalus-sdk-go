# dedalussdk

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#GetRootResponse">GetRootResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#DedalussdkService.GetRoot">GetRoot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#GetRootResponse">GetRootResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ChatCompletionTokenLogprob">ChatCompletionTokenLogprob</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ChatNewResponse">ChatNewResponse</a>

Methods:

- <code title="post /v1/chat">client.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ChatService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ChatNewParams">ChatNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go">dedalussdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dedalus-sdk-go#ChatNewResponse">ChatNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
