// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/ssestream"
	"github.com/dedalus-labs/dedalus-sdk-go/shared/constant"
)

// ChatCompletionService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatCompletionService] method instead.
type ChatCompletionService struct {
	Options []option.RequestOption
}

// NewChatCompletionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChatCompletionService(opts ...option.RequestOption) (r ChatCompletionService) {
	r = ChatCompletionService{}
	r.Options = opts
	return
}

// Create a chat completion.
//
// Unified chat-completions endpoint that works across many model providers.
// Supports optional MCP integration, multi-model routing with agentic handoffs,
// server- or client-executed tools, and both streaming and non-streaming delivery.
//
// Request body:
//
//   - messages: ordered list of chat turns.
//   - model: identifier or a list of identifiers for routing.
//   - tools: optional tool declarations available to the model.
//   - mcp_servers: optional list of MCP server slugs to enable during the run.
//   - stream: boolean to request incremental output.
//   - config: optional generation parameters (e.g., temperature, max_tokens,
//     metadata).
//
// Headers:
//
// - Authorization: bearer key for the calling account.
// - Optional BYOK or provider headers if applicable.
//
// Behavior:
//
//   - If multiple models are supplied, the router may select or hand off across
//     them.
//   - Tools may be invoked on the server or signaled for the client to run.
//   - Streaming responses emit incremental deltas; non-streaming returns a single
//     object.
//   - Usage metrics are computed when available and returned in the response.
//
// Responses:
//
//   - 200 OK: JSON completion object with choices, message content, and usage.
//   - 400 Bad Request: validation error.
//   - 401 Unauthorized: authentication failed.
//   - 402 Payment Required or 429 Too Many Requests: quota, balance, or rate limit
//     issue.
//   - 500 Internal Server Error: unexpected failure.
//
// Billing:
//
// - Token usage metered by the selected model(s).
// - Tool calls and MCP sessions may be billed separately.
// - Streaming is settled after the stream ends via an async task.
//
// Example (non-streaming HTTP): POST /v1/chat/completions Content-Type:
// application/json Authorization: Bearer <key>
//
// { "model": "provider/model-name", "messages": [{"role": "user", "content":
// "Hello"}] }
//
// 200 OK { "id": "cmpl_123", "object": "chat.completion", "choices": [ {"index":
// 0, "message": {"role": "assistant", "content": "Hi there!"}, "finish_reason":
// "stop"} ], "usage": {"prompt_tokens": 3, "completion_tokens": 4, "total_tokens":
// 7} }
//
// Example (streaming over SSE): POST /v1/chat/completions Accept:
// text/event-stream
//
// data: {"id":"cmpl_123","choices":[{"index":0,"delta":{"content":"Hi"}}]} data:
// {"id":"cmpl_123","choices":[{"index":0,"delta":{"content":" there!"}}]} data:
// [DONE]
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *Completion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a chat completion.
//
// Unified chat-completions endpoint that works across many model providers.
// Supports optional MCP integration, multi-model routing with agentic handoffs,
// server- or client-executed tools, and both streaming and non-streaming delivery.
//
// Request body:
//
//   - messages: ordered list of chat turns.
//   - model: identifier or a list of identifiers for routing.
//   - tools: optional tool declarations available to the model.
//   - mcp_servers: optional list of MCP server slugs to enable during the run.
//   - stream: boolean to request incremental output.
//   - config: optional generation parameters (e.g., temperature, max_tokens,
//     metadata).
//
// Headers:
//
// - Authorization: bearer key for the calling account.
// - Optional BYOK or provider headers if applicable.
//
// Behavior:
//
//   - If multiple models are supplied, the router may select or hand off across
//     them.
//   - Tools may be invoked on the server or signaled for the client to run.
//   - Streaming responses emit incremental deltas; non-streaming returns a single
//     object.
//   - Usage metrics are computed when available and returned in the response.
//
// Responses:
//
//   - 200 OK: JSON completion object with choices, message content, and usage.
//   - 400 Bad Request: validation error.
//   - 401 Unauthorized: authentication failed.
//   - 402 Payment Required or 429 Too Many Requests: quota, balance, or rate limit
//     issue.
//   - 500 Internal Server Error: unexpected failure.
//
// Billing:
//
// - Token usage metered by the selected model(s).
// - Tool calls and MCP sessions may be billed separately.
// - Streaming is settled after the stream ends via an async task.
//
// Example (non-streaming HTTP): POST /v1/chat/completions Content-Type:
// application/json Authorization: Bearer <key>
//
// { "model": "provider/model-name", "messages": [{"role": "user", "content":
// "Hello"}] }
//
// 200 OK { "id": "cmpl_123", "object": "chat.completion", "choices": [ {"index":
// 0, "message": {"role": "assistant", "content": "Hi there!"}, "finish_reason":
// "stop"} ], "usage": {"prompt_tokens": 3, "completion_tokens": 4, "total_tokens":
// 7} }
//
// Example (streaming over SSE): POST /v1/chat/completions Accept:
// text/event-stream
//
// data: {"id":"cmpl_123","choices":[{"index":0,"delta":{"content":"Hi"}}]} data:
// {"id":"cmpl_123","choices":[{"index":0,"delta":{"content":" there!"}}]} data:
// [DONE]
func (r *ChatCompletionService) NewStreaming(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[StreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[StreamChunk](ssestream.NewDecoder(raw), err)
}

// Token log probability information.
type ChatCompletionTokenLogprob struct {
	// The token.
	Token string `json:"token,required"`
	// A list of integers representing the UTF-8 bytes representation of the token.
	// Useful in instances where characters are represented by multiple tokens and
	// their byte representations must be combined to generate the correct text
	// representation. Can be `null` if there is no bytes representation for the token.
	Bytes []int64 `json:"bytes,required"`
	// The log probability of this token, if it is within the top 20 most likely
	// tokens. Otherwise, the value `-9999.0` is used to signify that the token is very
	// unlikely.
	Logprob float64 `json:"logprob,required"`
	// List of the most likely tokens and their log probability, at this token
	// position. In rare cases, there may be fewer than the number of requested
	// `top_logprobs` returned.
	TopLogprobs []TopLogprob `json:"top_logprobs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Bytes       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionTokenLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionTokenLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat completion response for Dedalus API.
//
// OpenAI-compatible chat completion response with Dedalus extensions. Maintains
// full compatibility with OpenAI API while providing additional features like
// server-side tool execution tracking and MCP error reporting.
type Completion struct {
	// A unique identifier for the chat completion.
	ID string `json:"id,required"`
	// A list of chat completion choices. Can be more than one if `n` is greater
	// than 1.
	Choices []CompletionChoice `json:"choices,required"`
	// The Unix timestamp (in seconds) of when the chat completion was created.
	Created int64 `json:"created,required"`
	// The model used for the chat completion.
	Model string `json:"model,required"`
	// The object type, which is always `chat.completion`.
	Object constant.ChatCompletion `json:"object,required"`
	// Information about MCP server failures, if any occurred during the request.
	// Contains details about which servers failed and why, along with recommendations
	// for the user. Only present when MCP server failures occurred.
	MCPServerErrors map[string]any `json:"mcp_server_errors,nullable"`
	// Specifies the processing type used for serving the request.
	//
	//   - If set to 'auto', then the request will be processed with the service tier
	//     configured in the Project settings. Unless otherwise configured, the Project
	//     will use 'default'.
	//   - If set to 'default', then the request will be processed with the standard
	//     pricing and performance for the selected model.
	//   - If set to '[flex](https://platform.openai.com/docs/guides/flex-processing)' or
	//     '[priority](https://openai.com/api-priority-processing/)', then the request
	//     will be processed with the corresponding service tier.
	//   - When not set, the default behavior is 'auto'.
	//
	// When the `service_tier` parameter is set, the response body will include the
	// `service_tier` value based on the processing mode actually used to serve the
	// request. This response value may be different from the value set in the
	// parameter.
	//
	// Any of "auto", "default", "flex", "scale", "priority".
	ServiceTier CompletionServiceTier `json:"service_tier,nullable"`
	// This fingerprint represents the backend configuration that the model runs with.
	//
	// Can be used in conjunction with the `seed` request parameter to understand when
	// backend changes have been made that might impact determinism.
	SystemFingerprint string `json:"system_fingerprint"`
	// List of tool names that were executed server-side (e.g., MCP tools). Only
	// present when tools were executed on the server rather than returned for
	// client-side execution.
	ToolsExecuted []string `json:"tools_executed,nullable"`
	// Usage statistics for the completion request.
	Usage CompletionUsage `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Choices           respjson.Field
		Created           respjson.Field
		Model             respjson.Field
		Object            respjson.Field
		MCPServerErrors   respjson.Field
		ServiceTier       respjson.Field
		SystemFingerprint respjson.Field
		ToolsExecuted     respjson.Field
		Usage             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Completion) RawJSON() string { return r.JSON.raw }
func (r *Completion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chat completion choice.
//
// OpenAI-compatible choice object for non-streaming responses. Part of the
// ChatCompletion response.
type CompletionChoice struct {
	// The index of the choice in the list of choices.
	Index int64 `json:"index,required"`
	// A chat completion message generated by the model.
	Message CompletionChoiceMessage `json:"message,required"`
	// The reason the model stopped generating tokens. This will be `stop` if the model
	// hit a natural stop point or a provided stop sequence, `length` if the maximum
	// number of tokens specified in the request was reached, `content_filter` if
	// content was omitted due to a flag from our content filters, `tool_calls` if the
	// model called a tool, or `function_call` (deprecated) if the model called a
	// function.
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason,nullable"`
	// Log probability information for the choice.
	Logprobs CompletionChoiceLogprobs `json:"logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index        respjson.Field
		Message      respjson.Field
		FinishReason respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chat completion message generated by the model.
type CompletionChoiceMessage struct {
	// The contents of the message.
	Content string `json:"content,required"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal,required"`
	// The role of the author of this message.
	Role constant.Assistant `json:"role,required"`
	// Annotations for the message, when applicable, as when using the
	// [web search tool](https://platform.openai.com/docs/guides/tools-web-search?api-mode=chat).
	Annotations []CompletionChoiceMessageAnnotation `json:"annotations"`
	// If the audio output modality is requested, this object contains data
	//
	// about the audio response from the model.
	// [Learn more](https://platform.openai.com/docs/guides/audio).
	//
	// Fields:
	//
	// - id (required): str
	// - expires_at (required): int
	// - data (required): str
	// - transcript (required): str
	Audio CompletionChoiceMessageAudio `json:"audio,nullable"`
	// Deprecated and replaced by `tool_calls`. The name and arguments of a function
	// that should be called, as generated by the model.
	FunctionCall CompletionChoiceMessageFunctionCall `json:"function_call"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []CompletionChoiceMessageToolCallUnion `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		Refusal      respjson.Field
		Role         respjson.Field
		Annotations  respjson.Field
		Audio        respjson.Field
		FunctionCall respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessage) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL citation when using web search.
//
// Fields:
//
// - type (required): Literal['url_citation']
// - url_citation (required): UrlCitation
type CompletionChoiceMessageAnnotation struct {
	// The type of the URL citation. Always `url_citation`.
	Type constant.URLCitation `json:"type,required"`
	// A URL citation when using web search.
	URLCitation CompletionChoiceMessageAnnotationURLCitation `json:"url_citation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URLCitation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageAnnotation) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageAnnotation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL citation when using web search.
type CompletionChoiceMessageAnnotationURLCitation struct {
	// The index of the last character of the URL citation in the message.
	EndIndex int64 `json:"end_index,required"`
	// The index of the first character of the URL citation in the message.
	StartIndex int64 `json:"start_index,required"`
	// The title of the web resource.
	Title string `json:"title,required"`
	// The URL of the web resource.
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndIndex    respjson.Field
		StartIndex  respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageAnnotationURLCitation) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If the audio output modality is requested, this object contains data
//
// about the audio response from the model.
// [Learn more](https://platform.openai.com/docs/guides/audio).
//
// Fields:
//
// - id (required): str
// - expires_at (required): int
// - data (required): str
// - transcript (required): str
type CompletionChoiceMessageAudio struct {
	// Unique identifier for this audio response.
	ID string `json:"id,required"`
	// Base64 encoded audio bytes generated by the model, in the format specified in
	// the request.
	Data string `json:"data,required"`
	// The Unix timestamp (in seconds) for when this audio response will no longer be
	// accessible on the server for use in multi-turn conversations.
	ExpiresAt int64 `json:"expires_at,required"`
	// Transcript of the audio generated by the model.
	Transcript string `json:"transcript,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Data        respjson.Field
		ExpiresAt   respjson.Field
		Transcript  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageAudio) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated and replaced by `tool_calls`. The name and arguments of a function
// that should be called, as generated by the model.
type CompletionChoiceMessageFunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CompletionChoiceMessageToolCallUnion contains all possible properties and values
// from [CompletionChoiceMessageToolCallFunction],
// [CompletionChoiceMessageToolCallCustom].
//
// Use the [CompletionChoiceMessageToolCallUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CompletionChoiceMessageToolCallUnion struct {
	ID string `json:"id"`
	// This field is from variant [CompletionChoiceMessageToolCallFunction].
	Function CompletionChoiceMessageToolCallFunctionFunction `json:"function"`
	// Any of "function", "custom".
	Type string `json:"type"`
	// This field is from variant [CompletionChoiceMessageToolCallCustom].
	Custom CompletionChoiceMessageToolCallCustomCustom `json:"custom"`
	JSON   struct {
		ID       respjson.Field
		Function respjson.Field
		Type     respjson.Field
		Custom   respjson.Field
		raw      string
	} `json:"-"`
}

// anyCompletionChoiceMessageToolCall is implemented by each variant of
// [CompletionChoiceMessageToolCallUnion] to add type safety for the return type of
// [CompletionChoiceMessageToolCallUnion.AsAny]
type anyCompletionChoiceMessageToolCall interface {
	implCompletionChoiceMessageToolCallUnion()
}

func (CompletionChoiceMessageToolCallFunction) implCompletionChoiceMessageToolCallUnion() {}
func (CompletionChoiceMessageToolCallCustom) implCompletionChoiceMessageToolCallUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CompletionChoiceMessageToolCallUnion.AsAny().(type) {
//	case githubcomdedaluslabsdedalussdkgo.CompletionChoiceMessageToolCallFunction:
//	case githubcomdedaluslabsdedalussdkgo.CompletionChoiceMessageToolCallCustom:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CompletionChoiceMessageToolCallUnion) AsAny() anyCompletionChoiceMessageToolCall {
	switch u.Type {
	case "function":
		return u.AsFunction()
	case "custom":
		return u.AsCustom()
	}
	return nil
}

func (u CompletionChoiceMessageToolCallUnion) AsFunction() (v CompletionChoiceMessageToolCallFunction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompletionChoiceMessageToolCallUnion) AsCustom() (v CompletionChoiceMessageToolCallCustom) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CompletionChoiceMessageToolCallUnion) RawJSON() string { return u.JSON.raw }

func (r *CompletionChoiceMessageToolCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): Function
type CompletionChoiceMessageToolCallFunction struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The function that the model called.
	Function CompletionChoiceMessageToolCallFunctionFunction `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type constant.Function `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The function that the model called.
type CompletionChoiceMessageToolCallFunctionFunction struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageToolCallFunctionFunction) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageToolCallFunctionFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A call to a custom tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['custom']
// - custom (required): Custom
type CompletionChoiceMessageToolCallCustom struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The custom tool that the model called.
	Custom CompletionChoiceMessageToolCallCustomCustom `json:"custom,required"`
	// The type of the tool. Always `custom`.
	Type constant.Custom `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Custom      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageToolCallCustom) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageToolCallCustom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The custom tool that the model called.
type CompletionChoiceMessageToolCallCustomCustom struct {
	// The input for the custom tool call generated by the model.
	Input string `json:"input,required"`
	// The name of the custom tool to call.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceMessageToolCallCustomCustom) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceMessageToolCallCustomCustom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log probability information for the choice.
type CompletionChoiceLogprobs struct {
	// A list of message content tokens with log probability information.
	Content []ChatCompletionTokenLogprob `json:"content,nullable"`
	// A list of message refusal tokens with log probability information.
	Refusal []ChatCompletionTokenLogprob `json:"refusal,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *CompletionChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the processing type used for serving the request.
//
//   - If set to 'auto', then the request will be processed with the service tier
//     configured in the Project settings. Unless otherwise configured, the Project
//     will use 'default'.
//   - If set to 'default', then the request will be processed with the standard
//     pricing and performance for the selected model.
//   - If set to '[flex](https://platform.openai.com/docs/guides/flex-processing)' or
//     '[priority](https://openai.com/api-priority-processing/)', then the request
//     will be processed with the corresponding service tier.
//   - When not set, the default behavior is 'auto'.
//
// When the `service_tier` parameter is set, the response body will include the
// `service_tier` value based on the processing mode actually used to serve the
// request. This response value may be different from the value set in the
// parameter.
type CompletionServiceTier string

const (
	CompletionServiceTierAuto     CompletionServiceTier = "auto"
	CompletionServiceTierDefault  CompletionServiceTier = "default"
	CompletionServiceTierFlex     CompletionServiceTier = "flex"
	CompletionServiceTierScale    CompletionServiceTier = "scale"
	CompletionServiceTierPriority CompletionServiceTier = "priority"
)

// Usage statistics for the completion request.
type CompletionUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
	// Breakdown of tokens used in a completion.
	CompletionTokensDetails CompletionUsageCompletionTokensDetails `json:"completion_tokens_details"`
	// Breakdown of tokens used in the prompt.
	PromptTokensDetails CompletionUsagePromptTokensDetails `json:"prompt_tokens_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		PromptTokens            respjson.Field
		TotalTokens             respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokensDetails     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionUsage) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of tokens used in a completion.
type CompletionUsageCompletionTokensDetails struct {
	// When using Predicted Outputs, the number of tokens in the prediction that
	// appeared in the completion.
	AcceptedPredictionTokens int64 `json:"accepted_prediction_tokens"`
	// Audio input tokens generated by the model.
	AudioTokens int64 `json:"audio_tokens"`
	// Tokens generated by the model for reasoning.
	ReasoningTokens int64 `json:"reasoning_tokens"`
	// When using Predicted Outputs, the number of tokens in the prediction that did
	// not appear in the completion. However, like reasoning tokens, these tokens are
	// still counted in the total completion tokens for purposes of billing, output,
	// and context window limits.
	RejectedPredictionTokens int64 `json:"rejected_prediction_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedPredictionTokens respjson.Field
		AudioTokens              respjson.Field
		ReasoningTokens          respjson.Field
		RejectedPredictionTokens respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of tokens used in the prompt.
type CompletionUsagePromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens int64 `json:"audio_tokens"`
	// Cached tokens present in the prompt.
	CachedTokens int64 `json:"cached_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioTokens  respjson.Field
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *CompletionUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelID = string

type ModelsParam []ModelUnionParam

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelUnionParam struct {
	OfModelID      param.Opt[ModelID]      `json:",omitzero,inline"`
	OfDedalusModel *ModelDedalusModelParam `json:",omitzero,inline"`
	paramUnion
}

func (u ModelUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModelID, u.OfDedalusModel)
}
func (u *ModelUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModelID) {
		return &u.OfModelID.Value
	} else if !param.IsOmitted(u.OfDedalusModel) {
		return u.OfDedalusModel
	}
	return nil
}

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
//
// The property Model is required.
type ModelDedalusModelParam struct {
	// Model identifier with provider prefix (e.g., 'openai/gpt-5',
	// 'anthropic/claude-3-5-sonnet').
	Model string `json:"model,required"`
	// Optional default generation settings (e.g., temperature, max_tokens) applied
	// when this model is selected.
	Settings ModelDedalusModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r ModelDedalusModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelDedalusModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelDedalusModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
type ModelDedalusModelSettingsParam struct {
	Deferred                        param.Opt[bool]                         `json:"deferred,omitzero"`
	FrequencyPenalty                param.Opt[float64]                      `json:"frequency_penalty,omitzero"`
	IncludeUsage                    param.Opt[bool]                         `json:"include_usage,omitzero"`
	InputAudioFormat                param.Opt[string]                       `json:"input_audio_format,omitzero"`
	Logprobs                        param.Opt[bool]                         `json:"logprobs,omitzero"`
	MaxCompletionTokens             param.Opt[int64]                        `json:"max_completion_tokens,omitzero"`
	MaxTokens                       param.Opt[int64]                        `json:"max_tokens,omitzero"`
	N                               param.Opt[int64]                        `json:"n,omitzero"`
	OutputAudioFormat               param.Opt[string]                       `json:"output_audio_format,omitzero"`
	ParallelToolCalls               param.Opt[bool]                         `json:"parallel_tool_calls,omitzero"`
	PresencePenalty                 param.Opt[float64]                      `json:"presence_penalty,omitzero"`
	PromptCacheKey                  param.Opt[string]                       `json:"prompt_cache_key,omitzero"`
	ReasoningEffort                 param.Opt[string]                       `json:"reasoning_effort,omitzero"`
	SafetyIdentifier                param.Opt[string]                       `json:"safety_identifier,omitzero"`
	Seed                            param.Opt[int64]                        `json:"seed,omitzero"`
	ServiceTier                     param.Opt[string]                       `json:"service_tier,omitzero"`
	Store                           param.Opt[bool]                         `json:"store,omitzero"`
	Stream                          param.Opt[bool]                         `json:"stream,omitzero"`
	Temperature                     param.Opt[float64]                      `json:"temperature,omitzero"`
	Timeout                         param.Opt[float64]                      `json:"timeout,omitzero"`
	TopK                            param.Opt[int64]                        `json:"top_k,omitzero"`
	TopLogprobs                     param.Opt[int64]                        `json:"top_logprobs,omitzero"`
	TopP                            param.Opt[float64]                      `json:"top_p,omitzero"`
	User                            param.Opt[string]                       `json:"user,omitzero"`
	Verbosity                       param.Opt[string]                       `json:"verbosity,omitzero"`
	Voice                           param.Opt[string]                       `json:"voice,omitzero"`
	DisableAutomaticFunctionCalling param.Opt[bool]                         `json:"disable_automatic_function_calling,omitzero"`
	UseResponses                    param.Opt[bool]                         `json:"use_responses,omitzero"`
	Audio                           map[string]any                          `json:"audio,omitzero"`
	ExtraArgs                       map[string]any                          `json:"extra_args,omitzero"`
	ExtraHeaders                    map[string]string                       `json:"extra_headers,omitzero"`
	ExtraQuery                      map[string]any                          `json:"extra_query,omitzero"`
	GenerationConfig                map[string]any                          `json:"generation_config,omitzero"`
	InputAudioTranscription         map[string]any                          `json:"input_audio_transcription,omitzero"`
	LogitBias                       map[string]int64                        `json:"logit_bias,omitzero"`
	Metadata                        map[string]string                       `json:"metadata,omitzero"`
	Modalities                      []string                                `json:"modalities,omitzero"`
	Prediction                      map[string]any                          `json:"prediction,omitzero"`
	Reasoning                       ModelDedalusModelSettingsReasoningParam `json:"reasoning,omitzero"`
	ResponseFormat                  map[string]any                          `json:"response_format,omitzero"`
	// Any of "code_interpreter_call.outputs", "computer_call_output.output.image_url",
	// "file_search_call.results", "message.input_image.image_url",
	// "message.output_text.logprobs", "reasoning.encrypted_content".
	ResponseInclude   []string                                      `json:"response_include,omitzero"`
	SafetySettings    []map[string]any                              `json:"safety_settings,omitzero"`
	SearchParameters  map[string]any                                `json:"search_parameters,omitzero"`
	Stop              ModelDedalusModelSettingsStopUnionParam       `json:"stop,omitzero"`
	StreamOptions     map[string]any                                `json:"stream_options,omitzero"`
	SystemInstruction map[string]any                                `json:"system_instruction,omitzero"`
	Thinking          map[string]any                                `json:"thinking,omitzero"`
	ToolChoice        ModelDedalusModelSettingsToolChoiceUnionParam `json:"tool_choice,omitzero"`
	ToolConfig        map[string]any                                `json:"tool_config,omitzero"`
	// Any of "auto", "disabled".
	Truncation       string         `json:"truncation,omitzero"`
	TurnDetection    map[string]any `json:"turn_detection,omitzero"`
	WebSearchOptions map[string]any `json:"web_search_options,omitzero"`
	Attributes       map[string]any `json:"attributes,omitzero"`
	StructuredOutput any            `json:"structured_output,omitzero"`
	paramObj
}

func (r ModelDedalusModelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelDedalusModelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelDedalusModelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ModelDedalusModelSettingsParam](
		"truncation", "auto", "disabled",
	)
}

type ModelDedalusModelSettingsReasoningParam struct {
	// Any of "minimal", "low", "medium", "high".
	Effort string `json:"effort,omitzero"`
	// Any of "auto", "concise", "detailed".
	GenerateSummary string `json:"generate_summary,omitzero"`
	// Any of "auto", "concise", "detailed".
	Summary     string         `json:"summary,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ModelDedalusModelSettingsReasoningParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelDedalusModelSettingsReasoningParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ModelDedalusModelSettingsReasoningParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ModelDedalusModelSettingsReasoningParam](
		"effort", "minimal", "low", "medium", "high",
	)
	apijson.RegisterFieldValidator[ModelDedalusModelSettingsReasoningParam](
		"generate_summary", "auto", "concise", "detailed",
	)
	apijson.RegisterFieldValidator[ModelDedalusModelSettingsReasoningParam](
		"summary", "auto", "concise", "detailed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelDedalusModelSettingsStopUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ModelDedalusModelSettingsStopUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ModelDedalusModelSettingsStopUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelDedalusModelSettingsStopUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ModelDedalusModelSettingsToolChoiceUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfModelDedalusModelSettingsToolChoiceString)
	OfModelDedalusModelSettingsToolChoiceString param.Opt[string]                                      `json:",omitzero,inline"`
	OfString                                    param.Opt[string]                                      `json:",omitzero,inline"`
	OfAnyMap                                    map[string]any                                         `json:",omitzero,inline"`
	OfMCPToolChoice                             *ModelDedalusModelSettingsToolChoiceMCPToolChoiceParam `json:",omitzero,inline"`
	paramUnion
}

func (u ModelDedalusModelSettingsToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModelDedalusModelSettingsToolChoiceString, u.OfString, u.OfAnyMap, u.OfMCPToolChoice)
}
func (u *ModelDedalusModelSettingsToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ModelDedalusModelSettingsToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModelDedalusModelSettingsToolChoiceString) {
		return &u.OfModelDedalusModelSettingsToolChoiceString
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfMCPToolChoice) {
		return u.OfMCPToolChoice
	}
	return nil
}

type ModelDedalusModelSettingsToolChoiceString string

const (
	ModelDedalusModelSettingsToolChoiceStringAuto     ModelDedalusModelSettingsToolChoiceString = "auto"
	ModelDedalusModelSettingsToolChoiceStringRequired ModelDedalusModelSettingsToolChoiceString = "required"
	ModelDedalusModelSettingsToolChoiceStringNone     ModelDedalusModelSettingsToolChoiceString = "none"
)

// The properties Name, ServerLabel are required.
type ModelDedalusModelSettingsToolChoiceMCPToolChoiceParam struct {
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	paramObj
}

func (r ModelDedalusModelSettingsToolChoiceMCPToolChoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelDedalusModelSettingsToolChoiceMCPToolChoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelDedalusModelSettingsToolChoiceMCPToolChoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-Sent Event streaming format for chat completions
type StreamChunk struct {
	// Unique identifier for the chat completion
	ID string `json:"id,required"`
	// List of completion choice chunks
	Choices []StreamChunkChoice `json:"choices,required"`
	// Unix timestamp when the chunk was created
	Created int64 `json:"created,required"`
	// ID of the model used for the completion
	Model string `json:"model,required"`
	// Object type, always 'chat.completion.chunk'
	//
	// Any of "chat.completion.chunk".
	Object StreamChunkObject `json:"object"`
	// Service tier used for processing the request
	//
	// Any of "auto", "default", "flex", "scale", "priority".
	ServiceTier StreamChunkServiceTier `json:"service_tier,nullable"`
	// System fingerprint representing backend configuration
	SystemFingerprint string `json:"system_fingerprint,nullable"`
	// Usage statistics for the completion request.
	//
	// Fields:
	//
	// - completion_tokens (required): int
	// - prompt_tokens (required): int
	// - total_tokens (required): int
	// - completion_tokens_details (optional): CompletionTokensDetails
	// - prompt_tokens_details (optional): PromptTokensDetails
	Usage StreamChunkUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Choices           respjson.Field
		Created           respjson.Field
		Model             respjson.Field
		Object            respjson.Field
		ServiceTier       respjson.Field
		SystemFingerprint respjson.Field
		Usage             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunk) RawJSON() string { return r.JSON.raw }
func (r *StreamChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A streaming chat completion choice chunk.
//
// OpenAI-compatible choice object for streaming responses. Part of the
// ChatCompletionChunk response in SSE streams.
type StreamChunkChoice struct {
	// Delta content for streaming responses
	Delta StreamChunkChoiceDelta `json:"delta,required"`
	// The index of this choice in the list of choices
	Index int64 `json:"index,required"`
	// The reason the model stopped (only in final chunk)
	//
	// Any of "stop", "length", "tool_calls", "content_filter", "function_call".
	FinishReason string `json:"finish_reason,nullable"`
	// Log probability information for the choice.
	Logprobs StreamChunkChoiceLogprobs `json:"logprobs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta        respjson.Field
		Index        respjson.Field
		FinishReason respjson.Field
		Logprobs     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkChoice) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delta content for streaming responses
type StreamChunkChoiceDelta struct {
	Content      string                             `json:"content,nullable"`
	FunctionCall StreamChunkChoiceDeltaFunctionCall `json:"function_call,nullable"`
	Refusal      string                             `json:"refusal,nullable"`
	// Any of "developer", "system", "user", "assistant", "tool".
	Role        string                           `json:"role,nullable"`
	ToolCalls   []StreamChunkChoiceDeltaToolCall `json:"tool_calls,nullable"`
	ExtraFields map[string]any                   `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		FunctionCall respjson.Field
		Refusal      respjson.Field
		Role         respjson.Field
		ToolCalls    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkChoiceDelta) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkChoiceDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamChunkChoiceDeltaFunctionCall struct {
	Arguments   string         `json:"arguments,nullable"`
	Name        string         `json:"name,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkChoiceDeltaFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkChoiceDeltaFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamChunkChoiceDeltaToolCall struct {
	Index    int64                                  `json:"index,required"`
	ID       string                                 `json:"id,nullable"`
	Function StreamChunkChoiceDeltaToolCallFunction `json:"function,nullable"`
	// Any of "function".
	Type        string         `json:"type,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index       respjson.Field
		ID          respjson.Field
		Function    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkChoiceDeltaToolCall) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkChoiceDeltaToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamChunkChoiceDeltaToolCallFunction struct {
	Arguments   string         `json:"arguments,nullable"`
	Name        string         `json:"name,nullable"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkChoiceDeltaToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkChoiceDeltaToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log probability information for the choice.
type StreamChunkChoiceLogprobs struct {
	// A list of message content tokens with log probability information.
	Content []ChatCompletionTokenLogprob `json:"content,nullable"`
	// A list of message refusal tokens with log probability information.
	Refusal []ChatCompletionTokenLogprob `json:"refusal,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Refusal     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object type, always 'chat.completion.chunk'
type StreamChunkObject string

const (
	StreamChunkObjectChatCompletionChunk StreamChunkObject = "chat.completion.chunk"
)

// Service tier used for processing the request
type StreamChunkServiceTier string

const (
	StreamChunkServiceTierAuto     StreamChunkServiceTier = "auto"
	StreamChunkServiceTierDefault  StreamChunkServiceTier = "default"
	StreamChunkServiceTierFlex     StreamChunkServiceTier = "flex"
	StreamChunkServiceTierScale    StreamChunkServiceTier = "scale"
	StreamChunkServiceTierPriority StreamChunkServiceTier = "priority"
)

// Usage statistics for the completion request.
//
// Fields:
//
// - completion_tokens (required): int
// - prompt_tokens (required): int
// - total_tokens (required): int
// - completion_tokens_details (optional): CompletionTokensDetails
// - prompt_tokens_details (optional): PromptTokensDetails
type StreamChunkUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
	// Breakdown of tokens used in a completion.
	CompletionTokensDetails StreamChunkUsageCompletionTokensDetails `json:"completion_tokens_details"`
	// Breakdown of tokens used in the prompt.
	PromptTokensDetails StreamChunkUsagePromptTokensDetails `json:"prompt_tokens_details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		PromptTokens            respjson.Field
		TotalTokens             respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokensDetails     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkUsage) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of tokens used in a completion.
type StreamChunkUsageCompletionTokensDetails struct {
	// When using Predicted Outputs, the number of tokens in the prediction that
	// appeared in the completion.
	AcceptedPredictionTokens int64 `json:"accepted_prediction_tokens"`
	// Audio input tokens generated by the model.
	AudioTokens int64 `json:"audio_tokens"`
	// Tokens generated by the model for reasoning.
	ReasoningTokens int64 `json:"reasoning_tokens"`
	// When using Predicted Outputs, the number of tokens in the prediction that did
	// not appear in the completion. However, like reasoning tokens, these tokens are
	// still counted in the total completion tokens for purposes of billing, output,
	// and context window limits.
	RejectedPredictionTokens int64 `json:"rejected_prediction_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedPredictionTokens respjson.Field
		AudioTokens              respjson.Field
		ReasoningTokens          respjson.Field
		RejectedPredictionTokens respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkUsageCompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkUsageCompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of tokens used in the prompt.
type StreamChunkUsagePromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens int64 `json:"audio_tokens"`
	// Cached tokens present in the prompt.
	CachedTokens int64 `json:"cached_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioTokens  respjson.Field
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamChunkUsagePromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *StreamChunkUsagePromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token and its log probability.
type TopLogprob struct {
	// The token.
	Token string `json:"token,required"`
	// A list of integers representing the UTF-8 bytes representation of the token.
	// Useful in instances where characters are represented by multiple tokens and
	// their byte representations must be combined to generate the correct text
	// representation. Can be `null` if there is no bytes representation for the token.
	Bytes []int64 `json:"bytes,required"`
	// The log probability of this token, if it is within the top 20 most likely
	// tokens. Otherwise, the value `-9999.0` is used to signify that the token is very
	// unlikely.
	Logprob float64 `json:"logprob,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Bytes       respjson.Field
		Logprob     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopLogprob) RawJSON() string { return r.JSON.raw }
func (r *TopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionNewParams struct {
	// Conversation history. Accepts either a list of message objects or a string,
	// which is treated as a single user message.
	Messages ChatCompletionNewParamsMessagesUnion `json:"messages,omitzero,required"`
	// Model(s) to use for completion. Can be a single model ID, a DedalusModel object,
	// or a list for multi-model routing. Single model: 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet-20241022', 'openai/gpt-4o-mini', or a DedalusModel
	// instance. Multi-model routing: ['openai/gpt-4o-mini', 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet'] or list of DedalusModel objects - agent will
	// choose optimal model based on task complexity.
	Model ChatCompletionNewParamsModelUnion `json:"model,omitzero,required"`
	// xAI-specific parameter. If set to true, the request returns a request_id for
	// async completion retrieval via GET /v1/chat/deferred-completion/{request_id}.
	Deferred param.Opt[bool] `json:"deferred,omitzero"`
	// Google-only flag to disable the SDK's automatic function execution. When true,
	// the model returns function calls for the client to execute manually.
	DisableAutomaticFunctionCalling param.Opt[bool] `json:"disable_automatic_function_calling,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing the model's likelihood to
	// repeat the same line verbatim.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// Whether to return log probabilities of the output tokens. If true, returns the
	// log probabilities for each token in the response content.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// An upper bound for the number of tokens that can be generated for a completion,
	// including visible output and reasoning tokens.
	MaxCompletionTokens param.Opt[int64] `json:"max_completion_tokens,omitzero"`
	// The maximum number of tokens that can be generated in the chat completion. This
	// value can be used to control costs for text generated via API. This value is now
	// deprecated in favor of 'max_completion_tokens' and is not compatible with
	// o-series models.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Maximum number of turns for agent execution before terminating (default: 10).
	// Each turn represents one model inference cycle. Higher values allow more complex
	// reasoning but increase cost and latency.
	MaxTurns param.Opt[int64] `json:"max_turns,omitzero"`
	// How many chat completion choices to generate for each input message. Keep 'n' as
	// 1 to minimize costs.
	N param.Opt[int64] `json:"n,omitzero"`
	// Whether to enable parallel function calling during tool use.
	ParallelToolCalls param.Opt[bool] `json:"parallel_tool_calls,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// Used by OpenAI to cache responses for similar requests and optimize cache hit
	// rates. Replaces the legacy 'user' field for caching.
	PromptCacheKey param.Opt[string] `json:"prompt_cache_key,omitzero"`
	// Stable identifier used to help detect users who might violate OpenAI usage
	// policies. Consider hashing end-user identifiers before sending.
	SafetyIdentifier param.Opt[string] `json:"safety_identifier,omitzero"`
	// If specified, system will make a best effort to sample deterministically.
	// Determinism is not guaranteed for the same seed across different models or API
	// versions.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Whether to store the output of this chat completion request for OpenAI model
	// distillation or eval products. Image inputs over 8MB are dropped if storage is
	// enabled.
	Store param.Opt[bool] `json:"store,omitzero"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 make
	// the output more random, while lower values like 0.2 make it more focused and
	// deterministic. We generally recommend altering this or 'top_p' but not both.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Top-k sampling. Anthropic: pass-through. Google: injected into
	// generationConfig.topK.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// An integer between 0 and 20 specifying how many of the most likely tokens to
	// return at each position, with log probabilities. Requires 'logprobs' to be true.
	TopLogprobs param.Opt[int64] `json:"top_logprobs,omitzero"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered. We
	// generally recommend altering this or 'temperature' but not both.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Stable identifier for your end-users. Helps OpenAI detect and prevent abuse and
	// may boost cache hit rates. This field is being replaced by 'safety_identifier'
	// and 'prompt_cache_key'.
	User param.Opt[string] `json:"user,omitzero"`
	// When False, skip server-side tool execution and return raw OpenAI-style
	// tool_calls in the response.
	AutoExecuteTools param.Opt[bool] `json:"auto_execute_tools,omitzero"`
	// Attributes for the agent itself, influencing behavior and model selection.
	// Format: {'attribute': value}, where values are 0.0-1.0. Common attributes:
	// 'complexity', 'accuracy', 'efficiency', 'creativity', 'friendliness'. Higher
	// values indicate stronger preference for that characteristic.
	AgentAttributes map[string]float64 `json:"agent_attributes,omitzero"`
	// Parameters for audio output. Required when requesting audio responses (for
	// example, modalities including 'audio').
	Audio map[string]any `json:"audio,omitzero"`
	// Deprecated in favor of 'tool_choice'. Controls which function is called by the
	// model (none, auto, or specific name).
	FunctionCall ChatCompletionNewParamsFunctionCallUnion `json:"function_call,omitzero"`
	// Deprecated in favor of 'tools'. Legacy list of function definitions the model
	// may generate JSON inputs for.
	Functions []map[string]any `json:"functions,omitzero"`
	// Google generationConfig object. Merged with auto-generated config. Use for
	// Google-specific params (candidateCount, responseMimeType, etc.).
	GenerationConfig map[string]any `json:"generation_config,omitzero"`
	// Guardrails to apply to the agent for input/output validation and safety checks.
	// Reserved for future use - guardrails configuration format not yet finalized.
	Guardrails []map[string]any `json:"guardrails,omitzero"`
	// Configuration for multi-model handoffs and agent orchestration. Reserved for
	// future use - handoff configuration format not yet finalized.
	HandoffConfig map[string]any `json:"handoff_config,omitzero"`
	// Convenience alias for Responses-style `input`. Used when `messages` is omitted
	// to provide the user prompt directly.
	Input ChatCompletionNewParamsInputUnion `json:"input,omitzero"`
	// Convenience alias for Responses-style `instructions`. Takes precedence over
	// `system` and over system-role messages when provided.
	Instructions ChatCompletionNewParamsInstructionsUnion `json:"instructions,omitzero"`
	// Modify the likelihood of specified tokens appearing in the completion. Accepts a
	// JSON object mapping token IDs (as strings) to bias values from -100 to 100. The
	// bias is added to the logits before sampling; values between -1 and 1 nudge
	// selection probability, while values like -100 or 100 effectively ban or require
	// a token.
	LogitBias map[string]int64 `json:"logit_bias,omitzero"`
	// MCP (Model Context Protocol) server addresses to make available for server-side
	// tool execution. Entries can be URLs (e.g., 'https://mcp.example.com'), slugs
	// (e.g., 'dedalus-labs/brave-search'), or structured objects specifying
	// slug/version/url. MCP tools are executed server-side and billed separately.
	MCPServers ChatCompletionNewParamsMCPServersUnion `json:"mcp_servers,omitzero"`
	// Set of up to 16 key-value string pairs that can be attached to the request for
	// structured metadata.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Output types you would like the model to generate. Most models default to
	// ['text']; some support ['text', 'audio'].
	Modalities []string `json:"modalities,omitzero"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution. Format: {'model_name': {'attribute': value}}, where values are
	// 0.0-1.0. Common attributes: 'intelligence', 'speed', 'cost', 'creativity',
	// 'accuracy'. Used by agent to select optimal model based on task requirements.
	ModelAttributes map[string]map[string]float64 `json:"model_attributes,omitzero"`
	// Configuration for predicted outputs. Improves response times when you already
	// know large portions of the response content.
	Prediction map[string]any `json:"prediction,omitzero"`
	// Constrains effort on reasoning for supported reasoning models. Higher values use
	// more compute, potentially improving reasoning quality at the cost of latency and
	// tokens.
	//
	// Any of "low", "medium", "high".
	ReasoningEffort ChatCompletionNewParamsReasoningEffort `json:"reasoning_effort,omitzero"`
	// An object specifying the format that the model must output. Use {'type':
	// 'json_schema', 'json_schema': {...}} for structured outputs or {'type':
	// 'json_object'} for the legacy JSON mode. Currently only OpenAI-prefixed models
	// honour this field; Anthropic and Google requests will return an
	// invalid_request_error if it is supplied.
	ResponseFormat map[string]any `json:"response_format,omitzero"`
	// Google safety settings (harm categories and thresholds).
	SafetySettings []map[string]any `json:"safety_settings,omitzero"`
	// xAI-specific parameter for configuring web search data acquisition. If not set,
	// no data will be acquired by the model.
	SearchParameters map[string]any `json:"search_parameters,omitzero"`
	// Specifies the processing tier used for the request. 'auto' uses project
	// defaults, while 'default' forces standard pricing and performance.
	//
	// Any of "auto", "default".
	ServiceTier ChatCompletionNewParamsServiceTier `json:"service_tier,omitzero"`
	// Not supported with latest reasoning models 'o3' and 'o4-mini'.
	//
	//	Up to 4 sequences where the API will stop generating further tokens; the returned text will not contain the stop sequence.
	Stop []string `json:"stop,omitzero"`
	// Options for streaming responses. Only set when 'stream' is true (supports
	// 'include_usage' and 'include_obfuscation').
	StreamOptions map[string]any `json:"stream_options,omitzero"`
	// System prompt/instructions. Anthropic: pass-through. Google: converted to
	// systemInstruction. OpenAI: extracted from messages.
	System ChatCompletionNewParamsSystemUnion `json:"system,omitzero"`
	// Extended thinking configuration (Anthropic only). Enables thinking blocks
	// showing reasoning process. Requires min 1,024 token budget.
	Thinking ChatCompletionNewParamsThinkingUnion `json:"thinking,omitzero"`
	// Controls which (if any) tool is called by the model. 'none' stops tool calling,
	// 'auto' lets the model decide, and 'required' forces at least one tool
	// invocation. Specific tool payloads force that tool.
	ToolChoice ChatCompletionNewParamsToolChoiceUnion `json:"tool_choice,omitzero"`
	// Google tool configuration (function calling mode, etc.).
	ToolConfig map[string]any `json:"tool_config,omitzero"`
	// A list of tools the model may call. Supports OpenAI function tools and custom
	// tools; use 'mcp_servers' for Dedalus-managed server-side tools.
	Tools []map[string]any `json:"tools,omitzero"`
	// Constrains the verbosity of the model's response. Lower values produce concise
	// answers, higher values allow more detail.
	//
	// Any of "low", "medium", "high".
	Verbosity ChatCompletionNewParamsVerbosity `json:"verbosity,omitzero"`
	// Configuration for OpenAI's web search tool. Learn more at
	// https://platform.openai.com/docs/guides/tools-web-search?api-mode=chat.
	WebSearchOptions map[string]any `json:"web_search_options,omitzero"`
	paramObj
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMessagesUnion struct {
	OfMapOfAnyMap []map[string]any  `json:",omitzero,inline"`
	OfString      param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMessagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfString)
}
func (u *ChatCompletionNewParamsMessagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMessagesUnion) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsModelUnion struct {
	OfModelID      param.Opt[ModelID]                        `json:",omitzero,inline"`
	OfDedalusModel *ChatCompletionNewParamsModelDedalusModel `json:",omitzero,inline"`
	OfModels       ModelsParam                               `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsModelUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModelID, u.OfDedalusModel, u.OfModels)
}
func (u *ChatCompletionNewParamsModelUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsModelUnion) asAny() any {
	if !param.IsOmitted(u.OfModelID) {
		return &u.OfModelID.Value
	} else if !param.IsOmitted(u.OfDedalusModel) {
		return u.OfDedalusModel
	} else if !param.IsOmitted(u.OfModels) {
		return &u.OfModels
	}
	return nil
}

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
//
// The property Model is required.
type ChatCompletionNewParamsModelDedalusModel struct {
	// Model identifier with provider prefix (e.g., 'openai/gpt-5',
	// 'anthropic/claude-3-5-sonnet').
	Model string `json:"model,required"`
	// Optional default generation settings (e.g., temperature, max_tokens) applied
	// when this model is selected.
	Settings ChatCompletionNewParamsModelDedalusModelSettings `json:"settings,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsModelDedalusModel) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsModelDedalusModel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsModelDedalusModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
type ChatCompletionNewParamsModelDedalusModelSettings struct {
	Deferred                        param.Opt[bool]                                           `json:"deferred,omitzero"`
	FrequencyPenalty                param.Opt[float64]                                        `json:"frequency_penalty,omitzero"`
	IncludeUsage                    param.Opt[bool]                                           `json:"include_usage,omitzero"`
	InputAudioFormat                param.Opt[string]                                         `json:"input_audio_format,omitzero"`
	Logprobs                        param.Opt[bool]                                           `json:"logprobs,omitzero"`
	MaxCompletionTokens             param.Opt[int64]                                          `json:"max_completion_tokens,omitzero"`
	MaxTokens                       param.Opt[int64]                                          `json:"max_tokens,omitzero"`
	N                               param.Opt[int64]                                          `json:"n,omitzero"`
	OutputAudioFormat               param.Opt[string]                                         `json:"output_audio_format,omitzero"`
	ParallelToolCalls               param.Opt[bool]                                           `json:"parallel_tool_calls,omitzero"`
	PresencePenalty                 param.Opt[float64]                                        `json:"presence_penalty,omitzero"`
	PromptCacheKey                  param.Opt[string]                                         `json:"prompt_cache_key,omitzero"`
	ReasoningEffort                 param.Opt[string]                                         `json:"reasoning_effort,omitzero"`
	SafetyIdentifier                param.Opt[string]                                         `json:"safety_identifier,omitzero"`
	Seed                            param.Opt[int64]                                          `json:"seed,omitzero"`
	ServiceTier                     param.Opt[string]                                         `json:"service_tier,omitzero"`
	Store                           param.Opt[bool]                                           `json:"store,omitzero"`
	Stream                          param.Opt[bool]                                           `json:"stream,omitzero"`
	Temperature                     param.Opt[float64]                                        `json:"temperature,omitzero"`
	Timeout                         param.Opt[float64]                                        `json:"timeout,omitzero"`
	TopK                            param.Opt[int64]                                          `json:"top_k,omitzero"`
	TopLogprobs                     param.Opt[int64]                                          `json:"top_logprobs,omitzero"`
	TopP                            param.Opt[float64]                                        `json:"top_p,omitzero"`
	User                            param.Opt[string]                                         `json:"user,omitzero"`
	Verbosity                       param.Opt[string]                                         `json:"verbosity,omitzero"`
	Voice                           param.Opt[string]                                         `json:"voice,omitzero"`
	DisableAutomaticFunctionCalling param.Opt[bool]                                           `json:"disable_automatic_function_calling,omitzero"`
	UseResponses                    param.Opt[bool]                                           `json:"use_responses,omitzero"`
	Audio                           map[string]any                                            `json:"audio,omitzero"`
	ExtraArgs                       map[string]any                                            `json:"extra_args,omitzero"`
	ExtraHeaders                    map[string]string                                         `json:"extra_headers,omitzero"`
	ExtraQuery                      map[string]any                                            `json:"extra_query,omitzero"`
	GenerationConfig                map[string]any                                            `json:"generation_config,omitzero"`
	InputAudioTranscription         map[string]any                                            `json:"input_audio_transcription,omitzero"`
	LogitBias                       map[string]int64                                          `json:"logit_bias,omitzero"`
	Metadata                        map[string]string                                         `json:"metadata,omitzero"`
	Modalities                      []string                                                  `json:"modalities,omitzero"`
	Prediction                      map[string]any                                            `json:"prediction,omitzero"`
	Reasoning                       ChatCompletionNewParamsModelDedalusModelSettingsReasoning `json:"reasoning,omitzero"`
	ResponseFormat                  map[string]any                                            `json:"response_format,omitzero"`
	// Any of "code_interpreter_call.outputs", "computer_call_output.output.image_url",
	// "file_search_call.results", "message.input_image.image_url",
	// "message.output_text.logprobs", "reasoning.encrypted_content".
	ResponseInclude   []string                                                        `json:"response_include,omitzero"`
	SafetySettings    []map[string]any                                                `json:"safety_settings,omitzero"`
	SearchParameters  map[string]any                                                  `json:"search_parameters,omitzero"`
	Stop              ChatCompletionNewParamsModelDedalusModelSettingsStopUnion       `json:"stop,omitzero"`
	StreamOptions     map[string]any                                                  `json:"stream_options,omitzero"`
	SystemInstruction map[string]any                                                  `json:"system_instruction,omitzero"`
	Thinking          map[string]any                                                  `json:"thinking,omitzero"`
	ToolChoice        ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceUnion `json:"tool_choice,omitzero"`
	ToolConfig        map[string]any                                                  `json:"tool_config,omitzero"`
	// Any of "auto", "disabled".
	Truncation       string         `json:"truncation,omitzero"`
	TurnDetection    map[string]any `json:"turn_detection,omitzero"`
	WebSearchOptions map[string]any `json:"web_search_options,omitzero"`
	Attributes       map[string]any `json:"attributes,omitzero"`
	StructuredOutput any            `json:"structured_output,omitzero"`
	paramObj
}

func (r ChatCompletionNewParamsModelDedalusModelSettings) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsModelDedalusModelSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsModelDedalusModelSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsModelDedalusModelSettings](
		"truncation", "auto", "disabled",
	)
}

type ChatCompletionNewParamsModelDedalusModelSettingsReasoning struct {
	// Any of "minimal", "low", "medium", "high".
	Effort string `json:"effort,omitzero"`
	// Any of "auto", "concise", "detailed".
	GenerateSummary string `json:"generate_summary,omitzero"`
	// Any of "auto", "concise", "detailed".
	Summary     string         `json:"summary,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ChatCompletionNewParamsModelDedalusModelSettingsReasoning) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsModelDedalusModelSettingsReasoning
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ChatCompletionNewParamsModelDedalusModelSettingsReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionNewParamsModelDedalusModelSettingsReasoning](
		"effort", "minimal", "low", "medium", "high",
	)
	apijson.RegisterFieldValidator[ChatCompletionNewParamsModelDedalusModelSettingsReasoning](
		"generate_summary", "auto", "concise", "detailed",
	)
	apijson.RegisterFieldValidator[ChatCompletionNewParamsModelDedalusModelSettingsReasoning](
		"summary", "auto", "concise", "detailed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsModelDedalusModelSettingsStopUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsModelDedalusModelSettingsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ChatCompletionNewParamsModelDedalusModelSettingsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsModelDedalusModelSettingsStopUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfChatCompletionNewsModelDedalusModelSettingsToolChoiceString)
	OfChatCompletionNewsModelDedalusModelSettingsToolChoiceString param.Opt[string]                                                        `json:",omitzero,inline"`
	OfString                                                      param.Opt[string]                                                        `json:",omitzero,inline"`
	OfAnyMap                                                      map[string]any                                                           `json:",omitzero,inline"`
	OfMCPToolChoice                                               *ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceMCPToolChoice `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChatCompletionNewsModelDedalusModelSettingsToolChoiceString, u.OfString, u.OfAnyMap, u.OfMCPToolChoice)
}
func (u *ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfChatCompletionNewsModelDedalusModelSettingsToolChoiceString) {
		return &u.OfChatCompletionNewsModelDedalusModelSettingsToolChoiceString
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfMCPToolChoice) {
		return u.OfMCPToolChoice
	}
	return nil
}

type ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceString string

const (
	ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceStringAuto     ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceString = "auto"
	ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceStringRequired ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceString = "required"
	ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceStringNone     ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceString = "none"
)

// The properties Name, ServerLabel are required.
type ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceMCPToolChoice struct {
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	paramObj
}

func (r ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceMCPToolChoice) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceMCPToolChoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsModelDedalusModelSettingsToolChoiceMCPToolChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsFunctionCallUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsFunctionCallUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *ChatCompletionNewParamsFunctionCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsFunctionCallUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsInputUnion struct {
	OfMapOfAnyMap []map[string]any  `json:",omitzero,inline"`
	OfString      param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfString)
}
func (u *ChatCompletionNewParamsInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsInputUnion) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsInstructionsUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfMapOfAnyMap []map[string]any  `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsInstructionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMapOfAnyMap)
}
func (u *ChatCompletionNewParamsInstructionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsInstructionsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMCPServersUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMCPServersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ChatCompletionNewParamsMCPServersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMCPServersUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Constrains effort on reasoning for supported reasoning models. Higher values use
// more compute, potentially improving reasoning quality at the cost of latency and
// tokens.
type ChatCompletionNewParamsReasoningEffort string

const (
	ChatCompletionNewParamsReasoningEffortLow    ChatCompletionNewParamsReasoningEffort = "low"
	ChatCompletionNewParamsReasoningEffortMedium ChatCompletionNewParamsReasoningEffort = "medium"
	ChatCompletionNewParamsReasoningEffortHigh   ChatCompletionNewParamsReasoningEffort = "high"
)

// Specifies the processing tier used for the request. 'auto' uses project
// defaults, while 'default' forces standard pricing and performance.
type ChatCompletionNewParamsServiceTier string

const (
	ChatCompletionNewParamsServiceTierAuto    ChatCompletionNewParamsServiceTier = "auto"
	ChatCompletionNewParamsServiceTierDefault ChatCompletionNewParamsServiceTier = "default"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsSystemUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfMapOfAnyMap []map[string]any  `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsSystemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMapOfAnyMap)
}
func (u *ChatCompletionNewParamsSystemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsSystemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsThinkingUnion struct {
	OfDisabled *ChatCompletionNewParamsThinkingDisabled `json:",omitzero,inline"`
	OfEnabled  *ChatCompletionNewParamsThinkingEnabled  `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsThinkingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDisabled, u.OfEnabled)
}
func (u *ChatCompletionNewParamsThinkingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsThinkingUnion) asAny() any {
	if !param.IsOmitted(u.OfDisabled) {
		return u.OfDisabled
	} else if !param.IsOmitted(u.OfEnabled) {
		return u.OfEnabled
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsThinkingUnion) GetBudgetTokens() *int64 {
	if vt := u.OfEnabled; vt != nil {
		return &vt.BudgetTokens
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionNewParamsThinkingUnion) GetType() *string {
	if vt := u.OfDisabled; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEnabled; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionNewParamsThinkingUnion](
		"type",
		apijson.Discriminator[ChatCompletionNewParamsThinkingDisabled]("disabled"),
		apijson.Discriminator[ChatCompletionNewParamsThinkingEnabled]("enabled"),
	)
}

func NewChatCompletionNewParamsThinkingDisabled() ChatCompletionNewParamsThinkingDisabled {
	return ChatCompletionNewParamsThinkingDisabled{
		Type: "disabled",
	}
}

// Fields:
//
// - type (required): Literal['disabled']
//
// This struct has a constant value, construct it with
// [NewChatCompletionNewParamsThinkingDisabled].
type ChatCompletionNewParamsThinkingDisabled struct {
	Type constant.Disabled `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsThinkingDisabled) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsThinkingDisabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsThinkingDisabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - budget_tokens (required): int
// - type (required): Literal['enabled']
//
// The properties BudgetTokens, Type are required.
type ChatCompletionNewParamsThinkingEnabled struct {
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens int64 `json:"budget_tokens,required"`
	// This field can be elided, and will marshal its zero value as "enabled".
	Type constant.Enabled `json:"type,required"`
	paramObj
}

func (r ChatCompletionNewParamsThinkingEnabled) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsThinkingEnabled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionNewParamsThinkingEnabled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsToolChoiceUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *ChatCompletionNewParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Constrains the verbosity of the model's response. Lower values produce concise
// answers, higher values allow more detail.
type ChatCompletionNewParamsVerbosity string

const (
	ChatCompletionNewParamsVerbosityLow    ChatCompletionNewParamsVerbosity = "low"
	ChatCompletionNewParamsVerbosityMedium ChatCompletionNewParamsVerbosity = "medium"
	ChatCompletionNewParamsVerbosityHigh   ChatCompletionNewParamsVerbosity = "high"
)
