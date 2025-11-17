// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/ssestream"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
	"github.com/tidwall/gjson"
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
func NewChatCompletionService(opts ...option.RequestOption) (r *ChatCompletionService) {
	r = &ChatCompletionService{}
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
	Bytes []int64 `json:"bytes,required,nullable"`
	// The log probability of this token, if it is within the top 20 most likely
	// tokens. Otherwise, the value `-9999.0` is used to signify that the token is very
	// unlikely.
	Logprob float64 `json:"logprob,required"`
	// List of the most likely tokens and their log probability, at this token
	// position. In rare cases, there may be fewer than the number of requested
	// `top_logprobs` returned.
	TopLogprobs []TopLogprob                   `json:"top_logprobs,required"`
	JSON        chatCompletionTokenLogprobJSON `json:"-"`
}

// chatCompletionTokenLogprobJSON contains the JSON metadata for the struct
// [ChatCompletionTokenLogprob]
type chatCompletionTokenLogprobJSON struct {
	Token       apijson.Field
	Bytes       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletionTokenLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionTokenLogprobJSON) RawJSON() string {
	return r.raw
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
	Object CompletionObject `json:"object,required"`
	// Information about MCP server failures, if any occurred during the request.
	// Contains details about which servers failed and why, along with recommendations
	// for the user. Only present when MCP server failures occurred.
	MCPServerErrors map[string]interface{} `json:"mcp_server_errors,nullable"`
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
	JSON  completionJSON  `json:"-"`
}

// completionJSON contains the JSON metadata for the struct [Completion]
type completionJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	MCPServerErrors   apijson.Field
	ServiceTier       apijson.Field
	SystemFingerprint apijson.Field
	ToolsExecuted     apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Completion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionJSON) RawJSON() string {
	return r.raw
}

// A chat completion choice.
//
// OpenAI-compatible choice object for non-streaming responses. Part of the
// ChatCompletion response.
type CompletionChoice struct {
	// The index of the choice in the list of choices.
	Index int64 `json:"index,required"`
	// A chat completion message generated by the model.
	Message CompletionChoicesMessage `json:"message,required"`
	// The reason the model stopped generating tokens. This will be `stop` if the model
	// hit a natural stop point or a provided stop sequence, `length` if the maximum
	// number of tokens specified in the request was reached, `content_filter` if
	// content was omitted due to a flag from our content filters, `tool_calls` if the
	// model called a tool, or `function_call` (deprecated) if the model called a
	// function.
	FinishReason CompletionChoicesFinishReason `json:"finish_reason,nullable"`
	// Log probability information for the choice.
	Logprobs CompletionChoicesLogprobs `json:"logprobs,nullable"`
	JSON     completionChoiceJSON      `json:"-"`
}

// completionChoiceJSON contains the JSON metadata for the struct
// [CompletionChoice]
type completionChoiceJSON struct {
	Index        apijson.Field
	Message      apijson.Field
	FinishReason apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoiceJSON) RawJSON() string {
	return r.raw
}

// A chat completion message generated by the model.
type CompletionChoicesMessage struct {
	// The contents of the message.
	Content CompletionChoicesMessageContentUnion `json:"content,required,nullable"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal,required,nullable"`
	// The role of the author of this message.
	Role CompletionChoicesMessageRole `json:"role,required"`
	// Annotations for the message, when applicable, as when using the
	// [web search tool](https://platform.openai.com/docs/guides/tools-web-search?api-mode=chat).
	Annotations []CompletionChoicesMessageAnnotation `json:"annotations"`
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
	Audio CompletionChoicesMessageAudio `json:"audio,nullable"`
	// Deprecated and replaced by `tool_calls`. The name and arguments of a function
	// that should be called, as generated by the model.
	FunctionCall CompletionChoicesMessageFunctionCall `json:"function_call"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []CompletionChoicesMessageToolCall `json:"tool_calls"`
	JSON      completionChoicesMessageJSON       `json:"-"`
}

// completionChoicesMessageJSON contains the JSON metadata for the struct
// [CompletionChoicesMessage]
type completionChoicesMessageJSON struct {
	Content      apijson.Field
	Refusal      apijson.Field
	Role         apijson.Field
	Annotations  apijson.Field
	Audio        apijson.Field
	FunctionCall apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageJSON) RawJSON() string {
	return r.raw
}

// The contents of the message.
//
// Union satisfied by [shared.UnionString] or
// [CompletionChoicesMessageContentArray].
type CompletionChoicesMessageContentUnion interface {
	ImplementsCompletionChoicesMessageContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CompletionChoicesMessageContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CompletionChoicesMessageContentArray{}),
		},
	)
}

type CompletionChoicesMessageContentArray []map[string]interface{}

func (r CompletionChoicesMessageContentArray) ImplementsCompletionChoicesMessageContentUnion() {}

// The role of the author of this message.
type CompletionChoicesMessageRole string

const (
	CompletionChoicesMessageRoleAssistant CompletionChoicesMessageRole = "assistant"
)

func (r CompletionChoicesMessageRole) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageRoleAssistant:
		return true
	}
	return false
}

// A URL citation when using web search.
//
// Fields:
//
// - type (required): Literal['url_citation']
// - url_citation (required): UrlCitation
type CompletionChoicesMessageAnnotation struct {
	// The type of the URL citation. Always `url_citation`.
	Type CompletionChoicesMessageAnnotationsType `json:"type,required"`
	// A URL citation when using web search.
	URLCitation CompletionChoicesMessageAnnotationsURLCitation `json:"url_citation,required"`
	JSON        completionChoicesMessageAnnotationJSON         `json:"-"`
}

// completionChoicesMessageAnnotationJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageAnnotation]
type completionChoicesMessageAnnotationJSON struct {
	Type        apijson.Field
	URLCitation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageAnnotationJSON) RawJSON() string {
	return r.raw
}

// The type of the URL citation. Always `url_citation`.
type CompletionChoicesMessageAnnotationsType string

const (
	CompletionChoicesMessageAnnotationsTypeURLCitation CompletionChoicesMessageAnnotationsType = "url_citation"
)

func (r CompletionChoicesMessageAnnotationsType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageAnnotationsTypeURLCitation:
		return true
	}
	return false
}

// A URL citation when using web search.
type CompletionChoicesMessageAnnotationsURLCitation struct {
	// The index of the last character of the URL citation in the message.
	EndIndex int64 `json:"end_index,required"`
	// The index of the first character of the URL citation in the message.
	StartIndex int64 `json:"start_index,required"`
	// The title of the web resource.
	Title string `json:"title,required"`
	// The URL of the web resource.
	URL  string                                             `json:"url,required"`
	JSON completionChoicesMessageAnnotationsURLCitationJSON `json:"-"`
}

// completionChoicesMessageAnnotationsURLCitationJSON contains the JSON metadata
// for the struct [CompletionChoicesMessageAnnotationsURLCitation]
type completionChoicesMessageAnnotationsURLCitationJSON struct {
	EndIndex    apijson.Field
	StartIndex  apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageAnnotationsURLCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageAnnotationsURLCitationJSON) RawJSON() string {
	return r.raw
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
type CompletionChoicesMessageAudio struct {
	// Unique identifier for this audio response.
	ID string `json:"id,required"`
	// Base64 encoded audio bytes generated by the model, in the format specified in
	// the request.
	Data string `json:"data,required"`
	// The Unix timestamp (in seconds) for when this audio response will no longer be
	// accessible on the server for use in multi-turn conversations.
	ExpiresAt int64 `json:"expires_at,required"`
	// Transcript of the audio generated by the model.
	Transcript string                            `json:"transcript,required"`
	JSON       completionChoicesMessageAudioJSON `json:"-"`
}

// completionChoicesMessageAudioJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageAudio]
type completionChoicesMessageAudioJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	ExpiresAt   apijson.Field
	Transcript  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageAudioJSON) RawJSON() string {
	return r.raw
}

// Deprecated and replaced by `tool_calls`. The name and arguments of a function
// that should be called, as generated by the model.
type CompletionChoicesMessageFunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string                                   `json:"name,required"`
	JSON completionChoicesMessageFunctionCallJSON `json:"-"`
}

// completionChoicesMessageFunctionCallJSON contains the JSON metadata for the
// struct [CompletionChoicesMessageFunctionCall]
type completionChoicesMessageFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageFunctionCallJSON) RawJSON() string {
	return r.raw
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): Function
type CompletionChoicesMessageToolCall struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type CompletionChoicesMessageToolCallsType `json:"type,required"`
	// This field can have the runtime type of
	// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustom].
	Custom interface{} `json:"custom"`
	// This field can have the runtime type of
	// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallFunction].
	Function interface{}                          `json:"function"`
	JSON     completionChoicesMessageToolCallJSON `json:"-"`
	union    CompletionChoicesMessageToolCallsUnion
}

// completionChoicesMessageToolCallJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCall]
type completionChoicesMessageToolCallJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Custom      apijson.Field
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r completionChoicesMessageToolCallJSON) RawJSON() string {
	return r.raw
}

func (r *CompletionChoicesMessageToolCall) UnmarshalJSON(data []byte) (err error) {
	*r = CompletionChoicesMessageToolCall{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CompletionChoicesMessageToolCallsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCall],
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall].
func (r CompletionChoicesMessageToolCall) AsUnion() CompletionChoicesMessageToolCallsUnion {
	return r.union
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): Function
//
// Union satisfied by
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCall] or
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall].
type CompletionChoicesMessageToolCallsUnion interface {
	implementsCompletionChoicesMessageToolCall()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CompletionChoicesMessageToolCallsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompletionChoicesMessageToolCallsChatCompletionMessageToolCall{}),
			DiscriminatorValue: "function",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall{}),
			DiscriminatorValue: "custom",
		},
	)
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): Function
type CompletionChoicesMessageToolCallsChatCompletionMessageToolCall struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The function that the model called.
	Function CompletionChoicesMessageToolCallsChatCompletionMessageToolCallFunction `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallType `json:"type,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageToolCallJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageToolCallJSON contains the
// JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCall]
type completionChoicesMessageToolCallsChatCompletionMessageToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageToolCallJSON) RawJSON() string {
	return r.raw
}

func (r CompletionChoicesMessageToolCallsChatCompletionMessageToolCall) implementsCompletionChoicesMessageToolCall() {
}

// The function that the model called.
type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallFunction struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string                                                                     `json:"name,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageToolCallFunctionJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageToolCallFunctionJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallFunction]
type completionChoicesMessageToolCallsChatCompletionMessageToolCallFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageToolCallFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageToolCallFunctionJSON) RawJSON() string {
	return r.raw
}

// The type of the tool. Currently, only `function` is supported.
type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallType string

const (
	CompletionChoicesMessageToolCallsChatCompletionMessageToolCallTypeFunction CompletionChoicesMessageToolCallsChatCompletionMessageToolCallType = "function"
)

func (r CompletionChoicesMessageToolCallsChatCompletionMessageToolCallType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsChatCompletionMessageToolCallTypeFunction:
		return true
	}
	return false
}

// A call to a custom tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['custom']
// - custom (required): Custom
type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The custom tool that the model called.
	Custom CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustom `json:"custom,required"`
	// The type of the tool. Always `custom`.
	Type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallType `json:"type,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall]
type completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallJSON struct {
	ID          apijson.Field
	Custom      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallJSON) RawJSON() string {
	return r.raw
}

func (r CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCall) implementsCompletionChoicesMessageToolCall() {
}

// The custom tool that the model called.
type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustom struct {
	// The input for the custom tool call generated by the model.
	Input string `json:"input,required"`
	// The name of the custom tool to call.
	Name string                                                                         `json:"name,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustomJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustomJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustom]
type completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustomJSON struct {
	Input       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallCustomJSON) RawJSON() string {
	return r.raw
}

// The type of the tool. Always `custom`.
type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallType string

const (
	CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallTypeCustom CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallType = "custom"
)

func (r CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallTypeCustom:
		return true
	}
	return false
}

// The type of the tool. Currently, only `function` is supported.
type CompletionChoicesMessageToolCallsType string

const (
	CompletionChoicesMessageToolCallsTypeFunction CompletionChoicesMessageToolCallsType = "function"
	CompletionChoicesMessageToolCallsTypeCustom   CompletionChoicesMessageToolCallsType = "custom"
)

func (r CompletionChoicesMessageToolCallsType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsTypeFunction, CompletionChoicesMessageToolCallsTypeCustom:
		return true
	}
	return false
}

// The reason the model stopped generating tokens. This will be `stop` if the model
// hit a natural stop point or a provided stop sequence, `length` if the maximum
// number of tokens specified in the request was reached, `content_filter` if
// content was omitted due to a flag from our content filters, `tool_calls` if the
// model called a tool, or `function_call` (deprecated) if the model called a
// function.
type CompletionChoicesFinishReason string

const (
	CompletionChoicesFinishReasonStop          CompletionChoicesFinishReason = "stop"
	CompletionChoicesFinishReasonLength        CompletionChoicesFinishReason = "length"
	CompletionChoicesFinishReasonToolCalls     CompletionChoicesFinishReason = "tool_calls"
	CompletionChoicesFinishReasonContentFilter CompletionChoicesFinishReason = "content_filter"
	CompletionChoicesFinishReasonFunctionCall  CompletionChoicesFinishReason = "function_call"
)

func (r CompletionChoicesFinishReason) IsKnown() bool {
	switch r {
	case CompletionChoicesFinishReasonStop, CompletionChoicesFinishReasonLength, CompletionChoicesFinishReasonToolCalls, CompletionChoicesFinishReasonContentFilter, CompletionChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

// Log probability information for the choice.
type CompletionChoicesLogprobs struct {
	// A list of message content tokens with log probability information.
	Content []ChatCompletionTokenLogprob `json:"content,nullable"`
	// A list of message refusal tokens with log probability information.
	Refusal []ChatCompletionTokenLogprob  `json:"refusal,nullable"`
	JSON    completionChoicesLogprobsJSON `json:"-"`
}

// completionChoicesLogprobsJSON contains the JSON metadata for the struct
// [CompletionChoicesLogprobs]
type completionChoicesLogprobsJSON struct {
	Content     apijson.Field
	Refusal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsJSON) RawJSON() string {
	return r.raw
}

// The object type, which is always `chat.completion`.
type CompletionObject string

const (
	CompletionObjectChatCompletion CompletionObject = "chat.completion"
)

func (r CompletionObject) IsKnown() bool {
	switch r {
	case CompletionObjectChatCompletion:
		return true
	}
	return false
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

func (r CompletionServiceTier) IsKnown() bool {
	switch r {
	case CompletionServiceTierAuto, CompletionServiceTierDefault, CompletionServiceTierFlex, CompletionServiceTierScale, CompletionServiceTierPriority:
		return true
	}
	return false
}

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
	JSON                completionUsageJSON                `json:"-"`
}

// completionUsageJSON contains the JSON metadata for the struct [CompletionUsage]
type completionUsageJSON struct {
	CompletionTokens        apijson.Field
	PromptTokens            apijson.Field
	TotalTokens             apijson.Field
	CompletionTokensDetails apijson.Field
	PromptTokensDetails     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CompletionUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionUsageJSON) RawJSON() string {
	return r.raw
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
	RejectedPredictionTokens int64                                      `json:"rejected_prediction_tokens"`
	JSON                     completionUsageCompletionTokensDetailsJSON `json:"-"`
}

// completionUsageCompletionTokensDetailsJSON contains the JSON metadata for the
// struct [CompletionUsageCompletionTokensDetails]
type completionUsageCompletionTokensDetailsJSON struct {
	AcceptedPredictionTokens apijson.Field
	AudioTokens              apijson.Field
	ReasoningTokens          apijson.Field
	RejectedPredictionTokens apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CompletionUsageCompletionTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionUsageCompletionTokensDetailsJSON) RawJSON() string {
	return r.raw
}

// Breakdown of tokens used in the prompt.
type CompletionUsagePromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens int64 `json:"audio_tokens"`
	// Cached tokens present in the prompt.
	CachedTokens int64                                  `json:"cached_tokens"`
	JSON         completionUsagePromptTokensDetailsJSON `json:"-"`
}

// completionUsagePromptTokensDetailsJSON contains the JSON metadata for the struct
// [CompletionUsagePromptTokensDetails]
type completionUsagePromptTokensDetailsJSON struct {
	AudioTokens  apijson.Field
	CachedTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompletionUsagePromptTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionUsagePromptTokensDetailsJSON) RawJSON() string {
	return r.raw
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
	Object StreamChunkObject `json:"object"`
	// Service tier used for processing the request
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
	JSON  streamChunkJSON  `json:"-"`
}

// streamChunkJSON contains the JSON metadata for the struct [StreamChunk]
type streamChunkJSON struct {
	ID                apijson.Field
	Choices           apijson.Field
	Created           apijson.Field
	Model             apijson.Field
	Object            apijson.Field
	ServiceTier       apijson.Field
	SystemFingerprint apijson.Field
	Usage             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *StreamChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkJSON) RawJSON() string {
	return r.raw
}

// A streaming chat completion choice chunk.
//
// OpenAI-compatible choice object for streaming responses. Part of the
// ChatCompletionChunk response in SSE streams.
type StreamChunkChoice struct {
	// Delta content for streaming responses
	Delta StreamChunkChoicesDelta `json:"delta,required"`
	// The index of this choice in the list of choices
	Index int64 `json:"index,required"`
	// The reason the model stopped (only in final chunk)
	FinishReason StreamChunkChoicesFinishReason `json:"finish_reason,nullable"`
	// Log probability information for the choice.
	Logprobs StreamChunkChoicesLogprobs `json:"logprobs,nullable"`
	JSON     streamChunkChoiceJSON      `json:"-"`
}

// streamChunkChoiceJSON contains the JSON metadata for the struct
// [StreamChunkChoice]
type streamChunkChoiceJSON struct {
	Delta        apijson.Field
	Index        apijson.Field
	FinishReason apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StreamChunkChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoiceJSON) RawJSON() string {
	return r.raw
}

// Delta content for streaming responses
type StreamChunkChoicesDelta struct {
	Content      string                              `json:"content,nullable"`
	FunctionCall StreamChunkChoicesDeltaFunctionCall `json:"function_call,nullable"`
	Refusal      string                              `json:"refusal,nullable"`
	Role         StreamChunkChoicesDeltaRole         `json:"role,nullable"`
	ToolCalls    []StreamChunkChoicesDeltaToolCall   `json:"tool_calls,nullable"`
	ExtraFields  map[string]interface{}              `json:"-,extras"`
	JSON         streamChunkChoicesDeltaJSON         `json:"-"`
}

// streamChunkChoicesDeltaJSON contains the JSON metadata for the struct
// [StreamChunkChoicesDelta]
type streamChunkChoicesDeltaJSON struct {
	Content      apijson.Field
	FunctionCall apijson.Field
	Refusal      apijson.Field
	Role         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StreamChunkChoicesDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaFunctionCall struct {
	Arguments   string                                  `json:"arguments,nullable"`
	Name        string                                  `json:"name,nullable"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
	JSON        streamChunkChoicesDeltaFunctionCallJSON `json:"-"`
}

// streamChunkChoicesDeltaFunctionCallJSON contains the JSON metadata for the
// struct [StreamChunkChoicesDeltaFunctionCall]
type streamChunkChoicesDeltaFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesDeltaFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaFunctionCallJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaRole string

const (
	StreamChunkChoicesDeltaRoleDeveloper StreamChunkChoicesDeltaRole = "developer"
	StreamChunkChoicesDeltaRoleSystem    StreamChunkChoicesDeltaRole = "system"
	StreamChunkChoicesDeltaRoleUser      StreamChunkChoicesDeltaRole = "user"
	StreamChunkChoicesDeltaRoleAssistant StreamChunkChoicesDeltaRole = "assistant"
	StreamChunkChoicesDeltaRoleTool      StreamChunkChoicesDeltaRole = "tool"
)

func (r StreamChunkChoicesDeltaRole) IsKnown() bool {
	switch r {
	case StreamChunkChoicesDeltaRoleDeveloper, StreamChunkChoicesDeltaRoleSystem, StreamChunkChoicesDeltaRoleUser, StreamChunkChoicesDeltaRoleAssistant, StreamChunkChoicesDeltaRoleTool:
		return true
	}
	return false
}

type StreamChunkChoicesDeltaToolCall struct {
	Index       int64                                    `json:"index,required"`
	ID          string                                   `json:"id,nullable"`
	Function    StreamChunkChoicesDeltaToolCallsFunction `json:"function,nullable"`
	Type        StreamChunkChoicesDeltaToolCallsType     `json:"type,nullable"`
	ExtraFields map[string]interface{}                   `json:"-,extras"`
	JSON        streamChunkChoicesDeltaToolCallJSON      `json:"-"`
}

// streamChunkChoicesDeltaToolCallJSON contains the JSON metadata for the struct
// [StreamChunkChoicesDeltaToolCall]
type streamChunkChoicesDeltaToolCallJSON struct {
	Index       apijson.Field
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesDeltaToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaToolCallJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaToolCallsFunction struct {
	Arguments   string                                       `json:"arguments,nullable"`
	Name        string                                       `json:"name,nullable"`
	ExtraFields map[string]interface{}                       `json:"-,extras"`
	JSON        streamChunkChoicesDeltaToolCallsFunctionJSON `json:"-"`
}

// streamChunkChoicesDeltaToolCallsFunctionJSON contains the JSON metadata for the
// struct [StreamChunkChoicesDeltaToolCallsFunction]
type streamChunkChoicesDeltaToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesDeltaToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesDeltaToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesDeltaToolCallsType string

const (
	StreamChunkChoicesDeltaToolCallsTypeFunction StreamChunkChoicesDeltaToolCallsType = "function"
)

func (r StreamChunkChoicesDeltaToolCallsType) IsKnown() bool {
	switch r {
	case StreamChunkChoicesDeltaToolCallsTypeFunction:
		return true
	}
	return false
}

// The reason the model stopped (only in final chunk)
type StreamChunkChoicesFinishReason string

const (
	StreamChunkChoicesFinishReasonStop          StreamChunkChoicesFinishReason = "stop"
	StreamChunkChoicesFinishReasonLength        StreamChunkChoicesFinishReason = "length"
	StreamChunkChoicesFinishReasonToolCalls     StreamChunkChoicesFinishReason = "tool_calls"
	StreamChunkChoicesFinishReasonContentFilter StreamChunkChoicesFinishReason = "content_filter"
	StreamChunkChoicesFinishReasonFunctionCall  StreamChunkChoicesFinishReason = "function_call"
)

func (r StreamChunkChoicesFinishReason) IsKnown() bool {
	switch r {
	case StreamChunkChoicesFinishReasonStop, StreamChunkChoicesFinishReasonLength, StreamChunkChoicesFinishReasonToolCalls, StreamChunkChoicesFinishReasonContentFilter, StreamChunkChoicesFinishReasonFunctionCall:
		return true
	}
	return false
}

// Log probability information for the choice.
type StreamChunkChoicesLogprobs struct {
	// A list of message content tokens with log probability information.
	Content []ChatCompletionTokenLogprob `json:"content,nullable"`
	// A list of message refusal tokens with log probability information.
	Refusal []ChatCompletionTokenLogprob   `json:"refusal,nullable"`
	JSON    streamChunkChoicesLogprobsJSON `json:"-"`
}

// streamChunkChoicesLogprobsJSON contains the JSON metadata for the struct
// [StreamChunkChoicesLogprobs]
type streamChunkChoicesLogprobsJSON struct {
	Content     apijson.Field
	Refusal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsJSON) RawJSON() string {
	return r.raw
}

// Object type, always 'chat.completion.chunk'
type StreamChunkObject string

const (
	StreamChunkObjectChatCompletionChunk StreamChunkObject = "chat.completion.chunk"
)

func (r StreamChunkObject) IsKnown() bool {
	switch r {
	case StreamChunkObjectChatCompletionChunk:
		return true
	}
	return false
}

// Service tier used for processing the request
type StreamChunkServiceTier string

const (
	StreamChunkServiceTierAuto     StreamChunkServiceTier = "auto"
	StreamChunkServiceTierDefault  StreamChunkServiceTier = "default"
	StreamChunkServiceTierFlex     StreamChunkServiceTier = "flex"
	StreamChunkServiceTierScale    StreamChunkServiceTier = "scale"
	StreamChunkServiceTierPriority StreamChunkServiceTier = "priority"
)

func (r StreamChunkServiceTier) IsKnown() bool {
	switch r {
	case StreamChunkServiceTierAuto, StreamChunkServiceTierDefault, StreamChunkServiceTierFlex, StreamChunkServiceTierScale, StreamChunkServiceTierPriority:
		return true
	}
	return false
}

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
	JSON                streamChunkUsageJSON                `json:"-"`
}

// streamChunkUsageJSON contains the JSON metadata for the struct
// [StreamChunkUsage]
type streamChunkUsageJSON struct {
	CompletionTokens        apijson.Field
	PromptTokens            apijson.Field
	TotalTokens             apijson.Field
	CompletionTokensDetails apijson.Field
	PromptTokensDetails     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *StreamChunkUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkUsageJSON) RawJSON() string {
	return r.raw
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
	RejectedPredictionTokens int64                                       `json:"rejected_prediction_tokens"`
	JSON                     streamChunkUsageCompletionTokensDetailsJSON `json:"-"`
}

// streamChunkUsageCompletionTokensDetailsJSON contains the JSON metadata for the
// struct [StreamChunkUsageCompletionTokensDetails]
type streamChunkUsageCompletionTokensDetailsJSON struct {
	AcceptedPredictionTokens apijson.Field
	AudioTokens              apijson.Field
	ReasoningTokens          apijson.Field
	RejectedPredictionTokens apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *StreamChunkUsageCompletionTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkUsageCompletionTokensDetailsJSON) RawJSON() string {
	return r.raw
}

// Breakdown of tokens used in the prompt.
type StreamChunkUsagePromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens int64 `json:"audio_tokens"`
	// Cached tokens present in the prompt.
	CachedTokens int64                                   `json:"cached_tokens"`
	JSON         streamChunkUsagePromptTokensDetailsJSON `json:"-"`
}

// streamChunkUsagePromptTokensDetailsJSON contains the JSON metadata for the
// struct [StreamChunkUsagePromptTokensDetails]
type streamChunkUsagePromptTokensDetailsJSON struct {
	AudioTokens  apijson.Field
	CachedTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StreamChunkUsagePromptTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkUsagePromptTokensDetailsJSON) RawJSON() string {
	return r.raw
}

// Token and its log probability.
type TopLogprob struct {
	// The token.
	Token string `json:"token,required"`
	// A list of integers representing the UTF-8 bytes representation of the token.
	// Useful in instances where characters are represented by multiple tokens and
	// their byte representations must be combined to generate the correct text
	// representation. Can be `null` if there is no bytes representation for the token.
	Bytes []int64 `json:"bytes,required,nullable"`
	// The log probability of this token, if it is within the top 20 most likely
	// tokens. Otherwise, the value `-9999.0` is used to signify that the token is very
	// unlikely.
	Logprob float64        `json:"logprob,required"`
	JSON    topLogprobJSON `json:"-"`
}

// topLogprobJSON contains the JSON metadata for the struct [TopLogprob]
type topLogprobJSON struct {
	Token       apijson.Field
	Bytes       apijson.Field
	Logprob     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r topLogprobJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionNewParams struct {
	// Model(s) to use for completion. Can be a single model ID, a DedalusModel object,
	// or a list for multi-model routing. Single model: 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet-20241022', 'openai/gpt-4o-mini', or a DedalusModel
	// instance. Multi-model routing: ['openai/gpt-4o-mini', 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet'] or list of DedalusModel objects - agent will
	// choose optimal model based on task complexity.
	Model param.Field[ChatCompletionNewParamsModelUnion] `json:"model,required"`
	// Attributes for the agent itself, influencing behavior and model selection.
	// Format: {'attribute': value}, where values are 0.0-1.0. Common attributes:
	// 'complexity', 'accuracy', 'efficiency', 'creativity', 'friendliness'. Higher
	// values indicate stronger preference for that characteristic.
	AgentAttributes param.Field[map[string]float64] `json:"agent_attributes"`
	// Parameters for audio output. Required when requesting audio responses (for
	// example, modalities including 'audio').
	Audio param.Field[map[string]interface{}] `json:"audio"`
	// When False, skip server-side tool execution and return raw OpenAI-style
	// tool_calls in the response.
	AutoExecuteTools param.Field[bool] `json:"auto_execute_tools"`
	// xAI-specific parameter. If set to true, the request returns a request_id for
	// async completion retrieval via GET /v1/chat/deferred-completion/{request_id}.
	Deferred param.Field[bool] `json:"deferred"`
	// Google-only flag to disable the SDK's automatic function execution. When true,
	// the model returns function calls for the client to execute manually.
	DisableAutomaticFunctionCalling param.Field[bool] `json:"disable_automatic_function_calling"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing the model's likelihood to
	// repeat the same line verbatim.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Deprecated in favor of 'tool_choice'. Controls which function is called by the
	// model (none, auto, or specific name).
	FunctionCall param.Field[ChatCompletionNewParamsFunctionCallUnion] `json:"function_call"`
	// Deprecated in favor of 'tools'. Legacy list of function definitions the model
	// may generate JSON inputs for.
	Functions param.Field[[]map[string]interface{}] `json:"functions"`
	// Google generationConfig object. Merged with auto-generated config. Use for
	// Google-specific params (candidateCount, responseMimeType, etc.).
	GenerationConfig param.Field[map[string]interface{}] `json:"generation_config"`
	// Guardrails to apply to the agent for input/output validation and safety checks.
	// Reserved for future use - guardrails configuration format not yet finalized.
	Guardrails param.Field[[]map[string]interface{}] `json:"guardrails"`
	// Configuration for multi-model handoffs and agent orchestration. Reserved for
	// future use - handoff configuration format not yet finalized.
	HandoffConfig param.Field[map[string]interface{}] `json:"handoff_config"`
	// Convenience alias for Responses-style `input`. Used when `messages` is omitted
	// to provide the user prompt directly.
	Input param.Field[ChatCompletionNewParamsInputUnion] `json:"input"`
	// Convenience alias for Responses-style `instructions`. Takes precedence over
	// `system` and over system-role messages when provided.
	Instructions param.Field[ChatCompletionNewParamsInstructionsUnion] `json:"instructions"`
	// Modify the likelihood of specified tokens appearing in the completion. Accepts a
	// JSON object mapping token IDs (as strings) to bias values from -100 to 100. The
	// bias is added to the logits before sampling; values between -1 and 1 nudge
	// selection probability, while values like -100 or 100 effectively ban or require
	// a token.
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// Whether to return log probabilities of the output tokens. If true, returns the
	// log probabilities for each token in the response content.
	Logprobs param.Field[bool] `json:"logprobs"`
	// An upper bound for the number of tokens that can be generated for a completion,
	// including visible output and reasoning tokens.
	MaxCompletionTokens param.Field[int64] `json:"max_completion_tokens"`
	// The maximum number of tokens that can be generated in the chat completion. This
	// value can be used to control costs for text generated via API. This value is now
	// deprecated in favor of 'max_completion_tokens' and is not compatible with
	// o-series models.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Maximum number of turns for agent execution before terminating (default: 10).
	// Each turn represents one model inference cycle. Higher values allow more complex
	// reasoning but increase cost and latency.
	MaxTurns param.Field[int64] `json:"max_turns"`
	// MCP (Model Context Protocol) server addresses to make available for server-side
	// tool execution. Entries can be URLs (e.g., 'https://mcp.example.com'), slugs
	// (e.g., 'dedalus-labs/brave-search'), or structured objects specifying
	// slug/version/url. MCP tools are executed server-side and billed separately.
	MCPServers param.Field[ChatCompletionNewParamsMCPServersUnion] `json:"mcp_servers"`
	// Conversation history. Accepts either a list of message objects or a string,
	// which is treated as a single user message. Optional if `input` or `instructions`
	// is provided.
	Messages param.Field[ChatCompletionNewParamsMessagesUnion] `json:"messages"`
	// Set of up to 16 key-value string pairs that can be attached to the request for
	// structured metadata.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Output types you would like the model to generate. Most models default to
	// ['text']; some support ['text', 'audio'].
	Modalities param.Field[[]string] `json:"modalities"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution. Format: {'model_name': {'attribute': value}}, where values are
	// 0.0-1.0. Common attributes: 'intelligence', 'speed', 'cost', 'creativity',
	// 'accuracy'. Used by agent to select optimal model based on task requirements.
	ModelAttributes param.Field[map[string]map[string]float64] `json:"model_attributes"`
	// How many chat completion choices to generate for each input message. Keep 'n' as
	// 1 to minimize costs.
	N param.Field[int64] `json:"n"`
	// Whether to enable parallel function calling during tool use.
	ParallelToolCalls param.Field[bool] `json:"parallel_tool_calls"`
	// Configuration for predicted outputs. Improves response times when you already
	// know large portions of the response content.
	Prediction param.Field[map[string]interface{}] `json:"prediction"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Used by OpenAI to cache responses for similar requests and optimize cache hit
	// rates. Replaces the legacy 'user' field for caching.
	PromptCacheKey param.Field[string] `json:"prompt_cache_key"`
	// Constrains effort on reasoning for supported reasoning models. Higher values use
	// more compute, potentially improving reasoning quality at the cost of latency and
	// tokens.
	ReasoningEffort param.Field[ChatCompletionNewParamsReasoningEffort] `json:"reasoning_effort"`
	// An object specifying the format that the model must output. Use {'type':
	// 'json_schema', 'json_schema': {...}} for structured outputs or {'type':
	// 'json_object'} for the legacy JSON mode. Currently only OpenAI-prefixed models
	// honor this field; Anthropic and Google requests will return an
	// invalid_request_error if it is supplied.
	ResponseFormat param.Field[ChatCompletionNewParamsResponseFormatUnion] `json:"response_format"`
	// Stable identifier used to help detect users who might violate OpenAI usage
	// policies. Consider hashing end-user identifiers before sending.
	SafetyIdentifier param.Field[string] `json:"safety_identifier"`
	// Google safety settings (harm categories and thresholds).
	SafetySettings param.Field[[]map[string]interface{}] `json:"safety_settings"`
	// xAI-specific parameter for configuring web search data acquisition. If not set,
	// no data will be acquired by the model.
	SearchParameters param.Field[map[string]interface{}] `json:"search_parameters"`
	// If specified, system will make a best effort to sample deterministically.
	// Determinism is not guaranteed for the same seed across different models or API
	// versions.
	Seed param.Field[int64] `json:"seed"`
	// Specifies the processing tier used for the request. 'auto' uses project
	// defaults, while 'default' forces standard pricing and performance.
	ServiceTier param.Field[ChatCompletionNewParamsServiceTier] `json:"service_tier"`
	// Not supported with latest reasoning models 'o3' and 'o4-mini'.
	//
	//	Up to 4 sequences where the API will stop generating further tokens; the returned text will not contain the stop sequence.
	Stop param.Field[[]string] `json:"stop"`
	// Whether to store the output of this chat completion request for OpenAI model
	// distillation or eval products. Image inputs over 8MB are dropped if storage is
	// enabled.
	Store param.Field[bool] `json:"store"`
	// Options for streaming responses. Only set when 'stream' is true (supports
	// 'include_usage' and 'include_obfuscation').
	StreamOptions param.Field[map[string]interface{}] `json:"stream_options"`
	// System prompt/instructions. Anthropic: pass-through. Google: converted to
	// systemInstruction. OpenAI: extracted from messages.
	System param.Field[ChatCompletionNewParamsSystemUnion] `json:"system"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 make
	// the output more random, while lower values like 0.2 make it more focused and
	// deterministic. We generally recommend altering this or 'top_p' but not both.
	Temperature param.Field[float64] `json:"temperature"`
	// Extended thinking configuration (Anthropic only). Enables thinking blocks
	// showing reasoning process. Requires min 1,024 token budget.
	Thinking param.Field[ChatCompletionNewParamsThinkingUnion] `json:"thinking"`
	// Controls which (if any) tool is called by the model. 'none' stops tool calling,
	// 'auto' lets the model decide, and 'required' forces at least one tool
	// invocation. Specific tool payloads force that tool.
	ToolChoice param.Field[ChatCompletionNewParamsToolChoiceUnion] `json:"tool_choice"`
	// Google tool configuration (function calling mode, etc.).
	ToolConfig param.Field[map[string]interface{}] `json:"tool_config"`
	// A list of tools the model may call. Supports OpenAI function tools and custom
	// tools; use 'mcp_servers' for Dedalus-managed server-side tools.
	Tools param.Field[[]map[string]interface{}] `json:"tools"`
	// Top-k sampling. Anthropic: pass-through. Google: injected into
	// generationConfig.topK.
	TopK param.Field[int64] `json:"top_k"`
	// An integer between 0 and 20 specifying how many of the most likely tokens to
	// return at each position, with log probabilities. Requires 'logprobs' to be true.
	TopLogprobs param.Field[int64] `json:"top_logprobs"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered. We
	// generally recommend altering this or 'temperature' but not both.
	TopP param.Field[float64] `json:"top_p"`
	// Stable identifier for your end-users. Helps OpenAI detect and prevent abuse and
	// may boost cache hit rates. This field is being replaced by 'safety_identifier'
	// and 'prompt_cache_key'.
	User param.Field[string] `json:"user"`
	// Constrains the verbosity of the model's response. Lower values produce concise
	// answers, higher values allow more detail.
	Verbosity param.Field[ChatCompletionNewParamsVerbosity] `json:"verbosity"`
	// Configuration for OpenAI's web search tool. Learn more at
	// https://platform.openai.com/docs/guides/tools-web-search?api-mode=chat.
	WebSearchOptions param.Field[map[string]interface{}] `json:"web_search_options"`
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Model(s) to use for completion. Can be a single model ID, a DedalusModel object,
// or a list for multi-model routing. Single model: 'openai/gpt-4',
// 'anthropic/claude-3-5-sonnet-20241022', 'openai/gpt-4o-mini', or a DedalusModel
// instance. Multi-model routing: ['openai/gpt-4o-mini', 'openai/gpt-4',
// 'anthropic/claude-3-5-sonnet'] or list of DedalusModel objects - agent will
// choose optimal model based on task complexity.
//
// Satisfied by [shared.UnionString], [shared.DedalusModelParam],
// [ChatCompletionNewParamsModelModels].
type ChatCompletionNewParamsModelUnion interface {
	ImplementsChatCompletionNewParamsModelUnion()
}

type ChatCompletionNewParamsModelModels []shared.DedalusModelChoiceUnionParam

func (r ChatCompletionNewParamsModelModels) ImplementsChatCompletionNewParamsModelUnion() {}

// Deprecated in favor of 'tool_choice'. Controls which function is called by the
// model (none, auto, or specific name).
//
// Satisfied by [shared.UnionString], [ChatCompletionNewParamsFunctionCallMap].
type ChatCompletionNewParamsFunctionCallUnion interface {
	ImplementsChatCompletionNewParamsFunctionCallUnion()
}

type ChatCompletionNewParamsFunctionCallMap map[string]interface{}

func (r ChatCompletionNewParamsFunctionCallMap) ImplementsChatCompletionNewParamsFunctionCallUnion() {
}

// Convenience alias for Responses-style `input`. Used when `messages` is omitted
// to provide the user prompt directly.
//
// Satisfied by [ChatCompletionNewParamsInputArray], [shared.UnionString].
type ChatCompletionNewParamsInputUnion interface {
	ImplementsChatCompletionNewParamsInputUnion()
}

type ChatCompletionNewParamsInputArray []map[string]interface{}

func (r ChatCompletionNewParamsInputArray) ImplementsChatCompletionNewParamsInputUnion() {}

// Convenience alias for Responses-style `instructions`. Takes precedence over
// `system` and over system-role messages when provided.
//
// Satisfied by [shared.UnionString], [ChatCompletionNewParamsInstructionsArray].
type ChatCompletionNewParamsInstructionsUnion interface {
	ImplementsChatCompletionNewParamsInstructionsUnion()
}

type ChatCompletionNewParamsInstructionsArray []map[string]interface{}

func (r ChatCompletionNewParamsInstructionsArray) ImplementsChatCompletionNewParamsInstructionsUnion() {
}

// MCP (Model Context Protocol) server addresses to make available for server-side
// tool execution. Entries can be URLs (e.g., 'https://mcp.example.com'), slugs
// (e.g., 'dedalus-labs/brave-search'), or structured objects specifying
// slug/version/url. MCP tools are executed server-side and billed separately.
//
// Satisfied by [shared.UnionString], [ChatCompletionNewParamsMCPServersArray].
type ChatCompletionNewParamsMCPServersUnion interface {
	ImplementsChatCompletionNewParamsMCPServersUnion()
}

type ChatCompletionNewParamsMCPServersArray []string

func (r ChatCompletionNewParamsMCPServersArray) ImplementsChatCompletionNewParamsMCPServersUnion() {}

// Conversation history. Accepts either a list of message objects or a string,
// which is treated as a single user message. Optional if `input` or `instructions`
// is provided.
//
// Satisfied by [ChatCompletionNewParamsMessagesArray], [shared.UnionString].
type ChatCompletionNewParamsMessagesUnion interface {
	ImplementsChatCompletionNewParamsMessagesUnion()
}

type ChatCompletionNewParamsMessagesArray []map[string]interface{}

func (r ChatCompletionNewParamsMessagesArray) ImplementsChatCompletionNewParamsMessagesUnion() {}

// Constrains effort on reasoning for supported reasoning models. Higher values use
// more compute, potentially improving reasoning quality at the cost of latency and
// tokens.
type ChatCompletionNewParamsReasoningEffort string

const (
	ChatCompletionNewParamsReasoningEffortLow    ChatCompletionNewParamsReasoningEffort = "low"
	ChatCompletionNewParamsReasoningEffortMedium ChatCompletionNewParamsReasoningEffort = "medium"
	ChatCompletionNewParamsReasoningEffortHigh   ChatCompletionNewParamsReasoningEffort = "high"
)

func (r ChatCompletionNewParamsReasoningEffort) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsReasoningEffortLow, ChatCompletionNewParamsReasoningEffortMedium, ChatCompletionNewParamsReasoningEffortHigh:
		return true
	}
	return false
}

// An object specifying the format that the model must output. Use {'type':
// 'json_schema', 'json_schema': {...}} for structured outputs or {'type':
// 'json_object'} for the legacy JSON mode. Currently only OpenAI-prefixed models
// honor this field; Anthropic and Google requests will return an
// invalid_request_error if it is supplied.
type ChatCompletionNewParamsResponseFormat struct {
	// The type of response format being defined. Always `text`.
	Type       param.Field[ChatCompletionNewParamsResponseFormatType] `json:"type,required"`
	JSONSchema param.Field[interface{}]                               `json:"json_schema"`
}

func (r ChatCompletionNewParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsResponseFormat) ImplementsChatCompletionNewParamsResponseFormatUnion() {
}

// An object specifying the format that the model must output. Use {'type':
// 'json_schema', 'json_schema': {...}} for structured outputs or {'type':
// 'json_object'} for the legacy JSON mode. Currently only OpenAI-prefixed models
// honor this field; Anthropic and Google requests will return an
// invalid_request_error if it is supplied.
//
// Satisfied by [shared.ResponseFormatTextParam],
// [shared.ResponseFormatJSONObjectParam], [shared.ResponseFormatJSONSchemaParam],
// [ChatCompletionNewParamsResponseFormat].
type ChatCompletionNewParamsResponseFormatUnion interface {
	ImplementsChatCompletionNewParamsResponseFormatUnion()
}

// The type of response format being defined. Always `text`.
type ChatCompletionNewParamsResponseFormatType string

const (
	ChatCompletionNewParamsResponseFormatTypeText       ChatCompletionNewParamsResponseFormatType = "text"
	ChatCompletionNewParamsResponseFormatTypeJSONObject ChatCompletionNewParamsResponseFormatType = "json_object"
	ChatCompletionNewParamsResponseFormatTypeJSONSchema ChatCompletionNewParamsResponseFormatType = "json_schema"
)

func (r ChatCompletionNewParamsResponseFormatType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsResponseFormatTypeText, ChatCompletionNewParamsResponseFormatTypeJSONObject, ChatCompletionNewParamsResponseFormatTypeJSONSchema:
		return true
	}
	return false
}

// Specifies the processing tier used for the request. 'auto' uses project
// defaults, while 'default' forces standard pricing and performance.
type ChatCompletionNewParamsServiceTier string

const (
	ChatCompletionNewParamsServiceTierAuto    ChatCompletionNewParamsServiceTier = "auto"
	ChatCompletionNewParamsServiceTierDefault ChatCompletionNewParamsServiceTier = "default"
)

func (r ChatCompletionNewParamsServiceTier) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsServiceTierAuto, ChatCompletionNewParamsServiceTierDefault:
		return true
	}
	return false
}

// System prompt/instructions. Anthropic: pass-through. Google: converted to
// systemInstruction. OpenAI: extracted from messages.
//
// Satisfied by [shared.UnionString], [ChatCompletionNewParamsSystemArray].
type ChatCompletionNewParamsSystemUnion interface {
	ImplementsChatCompletionNewParamsSystemUnion()
}

type ChatCompletionNewParamsSystemArray []map[string]interface{}

func (r ChatCompletionNewParamsSystemArray) ImplementsChatCompletionNewParamsSystemUnion() {}

// Extended thinking configuration (Anthropic only). Enables thinking blocks
// showing reasoning process. Requires min 1,024 token budget.
type ChatCompletionNewParamsThinking struct {
	Type param.Field[ChatCompletionNewParamsThinkingType] `json:"type,required"`
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64] `json:"budget_tokens"`
}

func (r ChatCompletionNewParamsThinking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsThinking) implementsChatCompletionNewParamsThinkingUnion() {}

// Extended thinking configuration (Anthropic only). Enables thinking blocks
// showing reasoning process. Requires min 1,024 token budget.
//
// Satisfied by [ChatCompletionNewParamsThinkingThinkingConfigDisabled],
// [ChatCompletionNewParamsThinkingThinkingConfigEnabled],
// [ChatCompletionNewParamsThinking].
type ChatCompletionNewParamsThinkingUnion interface {
	implementsChatCompletionNewParamsThinkingUnion()
}

// Fields:
//
// - type (required): Literal['disabled']
type ChatCompletionNewParamsThinkingThinkingConfigDisabled struct {
	Type param.Field[ChatCompletionNewParamsThinkingThinkingConfigDisabledType] `json:"type,required"`
}

func (r ChatCompletionNewParamsThinkingThinkingConfigDisabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsThinkingThinkingConfigDisabled) implementsChatCompletionNewParamsThinkingUnion() {
}

type ChatCompletionNewParamsThinkingThinkingConfigDisabledType string

const (
	ChatCompletionNewParamsThinkingThinkingConfigDisabledTypeDisabled ChatCompletionNewParamsThinkingThinkingConfigDisabledType = "disabled"
)

func (r ChatCompletionNewParamsThinkingThinkingConfigDisabledType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsThinkingThinkingConfigDisabledTypeDisabled:
		return true
	}
	return false
}

// Fields:
//
// - budget_tokens (required): int
// - type (required): Literal['enabled']
type ChatCompletionNewParamsThinkingThinkingConfigEnabled struct {
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.anthropic.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64]                                                    `json:"budget_tokens,required"`
	Type         param.Field[ChatCompletionNewParamsThinkingThinkingConfigEnabledType] `json:"type,required"`
}

func (r ChatCompletionNewParamsThinkingThinkingConfigEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsThinkingThinkingConfigEnabled) implementsChatCompletionNewParamsThinkingUnion() {
}

type ChatCompletionNewParamsThinkingThinkingConfigEnabledType string

const (
	ChatCompletionNewParamsThinkingThinkingConfigEnabledTypeEnabled ChatCompletionNewParamsThinkingThinkingConfigEnabledType = "enabled"
)

func (r ChatCompletionNewParamsThinkingThinkingConfigEnabledType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsThinkingThinkingConfigEnabledTypeEnabled:
		return true
	}
	return false
}

type ChatCompletionNewParamsThinkingType string

const (
	ChatCompletionNewParamsThinkingTypeDisabled ChatCompletionNewParamsThinkingType = "disabled"
	ChatCompletionNewParamsThinkingTypeEnabled  ChatCompletionNewParamsThinkingType = "enabled"
)

func (r ChatCompletionNewParamsThinkingType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsThinkingTypeDisabled, ChatCompletionNewParamsThinkingTypeEnabled:
		return true
	}
	return false
}

// Controls which (if any) tool is called by the model. 'none' stops tool calling,
// 'auto' lets the model decide, and 'required' forces at least one tool
// invocation. Specific tool payloads force that tool.
//
// Satisfied by [shared.UnionString], [ChatCompletionNewParamsToolChoiceMap].
type ChatCompletionNewParamsToolChoiceUnion interface {
	ImplementsChatCompletionNewParamsToolChoiceUnion()
}

type ChatCompletionNewParamsToolChoiceMap map[string]interface{}

func (r ChatCompletionNewParamsToolChoiceMap) ImplementsChatCompletionNewParamsToolChoiceUnion() {}

// Constrains the verbosity of the model's response. Lower values produce concise
// answers, higher values allow more detail.
type ChatCompletionNewParamsVerbosity string

const (
	ChatCompletionNewParamsVerbosityLow    ChatCompletionNewParamsVerbosity = "low"
	ChatCompletionNewParamsVerbosityMedium ChatCompletionNewParamsVerbosity = "medium"
	ChatCompletionNewParamsVerbosityHigh   ChatCompletionNewParamsVerbosity = "high"
)

func (r ChatCompletionNewParamsVerbosity) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsVerbosityLow, ChatCompletionNewParamsVerbosityMedium, ChatCompletionNewParamsVerbosityHigh:
		return true
	}
	return false
}
