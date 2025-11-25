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
// Generates a model response for the given conversation and configuration.
// Supports OpenAI-compatible parameters and provider-specific extensions.
//
// Headers:
//
// - Authorization: bearer key for the calling account.
// - Optional BYOK or provider headers if applicable.
//
// Behavior:
//
//   - If multiple models are supplied, the first one is used, and the agent may hand
//     off to another model.
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
// Generates a model response for the given conversation and configuration.
// Supports OpenAI-compatible parameters and provider-specific extensions.
//
// Headers:
//
// - Authorization: bearer key for the calling account.
// - Optional BYOK or provider headers if applicable.
//
// Behavior:
//
//   - If multiple models are supplied, the first one is used, and the agent may hand
//     off to another model.
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
	Content string `json:"content,required,nullable"`
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
	// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustom].
	Custom interface{} `json:"custom"`
	// This field can have the runtime type of
	// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunction].
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
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput],
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput].
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
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput] or
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput].
type CompletionChoicesMessageToolCallsUnion interface {
	implementsCompletionChoicesMessageToolCall()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CompletionChoicesMessageToolCallsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput{}),
			DiscriminatorValue: "function",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput{}),
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
type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The function that the model called.
	Function CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunction `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputType `json:"type,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput]
type completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputJSON) RawJSON() string {
	return r.raw
}

func (r CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutput) implementsCompletionChoicesMessageToolCall() {
}

// The function that the model called.
type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunction struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string                                                                           `json:"name,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunctionJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunctionJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunction]
type completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageToolCallOutputFunctionJSON) RawJSON() string {
	return r.raw
}

// The type of the tool. Currently, only `function` is supported.
type CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputType string

const (
	CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputTypeFunction CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputType = "function"
)

func (r CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsChatCompletionMessageToolCallOutputTypeFunction:
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
type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The custom tool that the model called.
	Custom CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustom `json:"custom,required"`
	// The type of the tool. Always `custom`.
	Type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputType `json:"type,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput]
type completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputJSON struct {
	ID          apijson.Field
	Custom      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputJSON) RawJSON() string {
	return r.raw
}

func (r CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutput) implementsCompletionChoicesMessageToolCall() {
}

// The custom tool that the model called.
type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustom struct {
	// The input for the custom tool call generated by the model.
	Input string `json:"input,required"`
	// The name of the custom tool to call.
	Name string                                                                               `json:"name,required"`
	JSON completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustomJSON `json:"-"`
}

// completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustomJSON
// contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustom]
type completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustomJSON struct {
	Input       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputCustomJSON) RawJSON() string {
	return r.raw
}

// The type of the tool. Always `custom`.
type CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputType string

const (
	CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputTypeCustom CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputType = "custom"
)

func (r CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsChatCompletionMessageCustomToolCallOutputTypeCustom:
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
	// or a list for multi-model routing. Single model: 'openai/gpt-5',
	// 'anthropic/claude-sonnet-4-5-20250929', 'google/gemini-3-pro-preview', or a
	// DedalusModel instance. Multi-model routing: ['openai/gpt-5',
	// 'anthropic/claude-sonnet-4-5-20250929', 'google/gemini-3-pro-preview'] or list
	// of DedalusModel objects - agent will choose optimal model based on task
	// complexity.
	Model param.Field[ChatCompletionNewParamsModelUnion] `json:"model,required"`
	// Attributes for the agent itself, influencing behavior and model selection.
	// Format: {'attribute': value}, where values are 0.0-1.0. Common attributes:
	// 'complexity', 'accuracy', 'efficiency', 'creativity', 'friendliness'. Higher
	// values indicate stronger preference for that characteristic.
	AgentAttributes param.Field[map[string]float64] `json:"agent_attributes"`
	// Parameters for audio output. Required when audio output is requested with
	// `modalities: ["audio"]`.
	// [Learn more](https://platform.openai.com/docs/guides/audio).
	Audio param.Field[map[string]interface{}] `json:"audio"`
	// When False, skip server-side tool execution and return raw OpenAI-style
	// tool_calls in the response.
	AutoExecuteTools param.Field[bool] `json:"auto_execute_tools"`
	// Optional. The name of the content
	// [cached](https://ai.google.dev/gemini-api/docs/caching) to use as context to
	// serve the prediction. Format: `cachedContents/{cachedContent}`
	CachedContent param.Field[string] `json:"cachedContent"`
	// If set to `true`, the request returns a `request_id`. You can then get the
	// deferred response by GET `/v1/chat/deferred-completion/{request_id}`.
	Deferred param.Field[bool] `json:"deferred"`
	// Google SDK control: disable automatic function calling. Agent workflows handle
	// tools manually.
	DisableAutomaticFunctionCalling param.Field[bool] `json:"disable_automatic_function_calling"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing the model's likelihood to
	// repeat the same line verbatim.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Deprecated in favor of `tool_choice`. Controls which (if any) function is called
	// by the model. `none` means the model will not call a function and instead
	// generates a message. `auto` means the model can pick between generating a
	// message or calling a function. Specifying a particular function via
	// `{"name": "my_function"}` forces the model to call that function. `none` is the
	// default when no functions are present. `auto` is the default if functions are
	// present.
	FunctionCall param.Field[ChatCompletionNewParamsFunctionCall] `json:"function_call"`
	// Deprecated in favor of `tools`. A list of functions the model may generate JSON
	// inputs for.
	Functions param.Field[[]ChatCompletionNewParamsFunction] `json:"functions"`
	// Generation parameters wrapper (Google-specific)
	GenerationConfig param.Field[map[string]interface{}] `json:"generation_config"`
	// Guardrails to apply to the agent for input/output validation and safety checks.
	// Reserved for future use - guardrails configuration format not yet finalized.
	Guardrails param.Field[[]map[string]interface{}] `json:"guardrails"`
	// Configuration for multi-model handoffs and agent orchestration. Reserved for
	// future use - handoff configuration format not yet finalized.
	HandoffConfig param.Field[map[string]interface{}] `json:"handoff_config"`
	// Modify the likelihood of specified tokens appearing in the completion. Accepts a
	// JSON object that maps tokens (specified by their token ID in the tokenizer) to
	// an associated bias value from -100 to 100. Mathematically, the bias is added to
	// the logits generated by the model prior to sampling. The exact effect will vary
	// per model, but values between -1 and 1 should decrease or increase likelihood of
	// selection; values like -100 or 100 should result in a ban or exclusive selection
	// of the relevant token.
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// Whether to return log probabilities of the output tokens or not. If true,
	// returns the log probabilities of each output token returned in the `content` of
	// `message`.
	Logprobs param.Field[bool] `json:"logprobs"`
	// An upper bound for the number of tokens that can be generated for a completion,
	// including visible output and reasoning tokens.
	MaxCompletionTokens param.Field[int64] `json:"max_completion_tokens"`
	// Maximum number of tokens the model can generate in the completion. The total
	// token count (input + output) is limited by the model's context window. Setting
	// this prevents unexpectedly long responses and helps control costs. For newer
	// OpenAI models, use max_completion_tokens instead (more precise accounting). For
	// other providers, max_tokens remains the standard parameter name.
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
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard. Keys are strings with a maximum
	// length of 64 characters. Values are strings with a maximum length of 512
	// characters.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Output modalities. Most models generate text by default. Use ['text', 'audio']
	// for audio-capable models.
	Modalities param.Field[[]ChatCompletionNewParamsModality] `json:"modalities"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution. Format: {'model_name': {'attribute': value}}, where values are
	// 0.0-1.0. Common attributes: 'intelligence', 'speed', 'cost', 'creativity',
	// 'accuracy'. Used by agent to select optimal model based on task requirements.
	ModelAttributes param.Field[map[string]map[string]float64] `json:"model_attributes"`
	// How many chat completion choices to generate for each input message. Note that
	// you will be charged based on the number of generated tokens across all of the
	// choices. Keep `n` as `1` to minimize costs.
	N param.Field[int64] `json:"n"`
	// Whether to enable parallel tool calls (Anthropic uses inverted polarity)
	ParallelToolCalls param.Field[bool] `json:"parallel_tool_calls"`
	// Static predicted output content, such as the content of a text file that is
	// being regenerated.
	Prediction param.Field[ChatCompletionNewParamsPrediction] `json:"prediction"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Used by OpenAI to cache responses for similar requests to optimize your cache
	// hit rates. Replaces the `user` field.
	// [Learn more](https://platform.openai.com/docs/guides/prompt-caching).
	PromptCacheKey param.Field[string] `json:"prompt_cache_key"`
	// The retention policy for the prompt cache. Set to `24h` to enable extended
	// prompt caching, which keeps cached prefixes active for longer, up to a maximum
	// of 24 hours.
	// [Learn more](https://platform.openai.com/docs/guides/prompt-caching#prompt-cache-retention).
	PromptCacheRetention param.Field[ChatCompletionNewParamsPromptCacheRetention] `json:"prompt_cache_retention"`
	// Allows toggling between the reasoning mode and no system prompt. When set to
	// `reasoning` the system prompt for reasoning models will be used.
	PromptMode param.Field[map[string]interface{}] `json:"prompt_mode"`
	// Constrains effort on reasoning for
	// [reasoning models](https://platform.openai.com/docs/guides/reasoning). Currently
	// supported values are `none`, `minimal`, `low`, `medium`, and `high`. Reducing
	// reasoning effort can result in faster responses and fewer tokens used on
	// reasoning in a response. - `gpt-5.1` defaults to `none`, which does not perform
	// reasoning. The supported reasoning values for `gpt-5.1` are `none`, `low`,
	// `medium`, and `high`. Tool calls are supported for all reasoning values in
	// gpt-5.1. - All models before `gpt-5.1` default to `medium` reasoning effort, and
	// do not support `none`. - The `gpt-5-pro` model defaults to (and only supports)
	// `high` reasoning effort.
	ReasoningEffort param.Field[ChatCompletionNewParamsReasoningEffort] `json:"reasoning_effort"`
	// An object specifying the format that the model must output. Setting to
	// `{ "type": "json_schema", "json_schema": {...} }` enables Structured Outputs
	// which ensures the model will match your supplied JSON schema. Learn more in the
	// [Structured Outputs guide](https://platform.openai.com/docs/guides/structured-outputs).
	// Setting to `{ "type": "json_object" }` enables the older JSON mode, which
	// ensures the message the model generates is valid JSON. Using `json_schema` is
	// preferred for models that support it.
	ResponseFormat param.Field[ChatCompletionNewParamsResponseFormatUnion] `json:"response_format"`
	// Whether to inject a safety prompt before all conversations.
	SafePrompt param.Field[bool] `json:"safe_prompt"`
	// A stable identifier used to help detect users of your application that may be
	// violating OpenAI's usage policies. The IDs should be a string that uniquely
	// identifies each user. We recommend hashing their username or email address, in
	// order to avoid sending us any identifying information.
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#safety-identifiers).
	SafetyIdentifier param.Field[string] `json:"safety_identifier"`
	// Safety/content filtering settings (Google-specific)
	SafetySettings param.Field[[]ChatCompletionNewParamsSafetySetting] `json:"safety_settings"`
	// Set the parameters to be used for searched data. If not set, no data will be
	// acquired by the model.
	SearchParameters param.Field[map[string]interface{}] `json:"search_parameters"`
	// Random seed for deterministic output
	Seed param.Field[int64] `json:"seed"`
	// Service tier for request processing
	ServiceTier param.Field[ChatCompletionNewParamsServiceTier] `json:"service_tier"`
	// Not supported with latest reasoning models 'o3' and 'o4-mini'. Up to 4 sequences
	// where the API will stop generating further tokens; the returned text will not
	// contain the stop sequence.
	Stop param.Field[ChatCompletionNewParamsStopUnion] `json:"stop"`
	// Whether or not to store the output of this chat completion request for use in
	// our [model distillation](https://platform.openai.com/docs/guides/distillation)
	// or [evals](https://platform.openai.com/docs/guides/evals) products. Supports
	// text and image inputs. Note: image inputs over 8MB will be dropped.
	Store param.Field[bool] `json:"store"`
	// Options for streaming response. Only set this when you set `stream: true`.
	StreamOptions param.Field[map[string]interface{}] `json:"stream_options"`
	// System-level instructions defining the assistant's behavior, role, and
	// constraints. Sets the context and personality for the entire conversation.
	// Different from user/assistant messages as it provides meta-instructions about
	// how to respond rather than conversation content. OpenAI: Provided as system role
	// message in messages array. Google: Top-level systemInstruction field (adapter
	// extracts from messages). Anthropic: Top-level system parameter (adapter extracts
	// from messages). Accepts both string and structured object formats depending on
	// provider capabilities.
	SystemInstruction param.Field[ChatCompletionNewParamsSystemInstructionUnion] `json:"system_instruction"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will
	// make the output more random, while lower values like 0.2 will make it more
	// focused and deterministic. We generally recommend altering this or top_p but not
	// both.
	Temperature param.Field[float64] `json:"temperature"`
	// Extended thinking configuration (Anthropic-specific)
	Thinking param.Field[ChatCompletionNewParamsThinkingUnion] `json:"thinking"`
	// Controls which (if any) tool is called by the model. `none` means the model will
	// not call any tool and instead generates a message. `auto` means the model can
	// pick between generating a message or calling one or more tools. `required` means
	// the model must call one or more tools. Specifying a particular tool via
	// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
	// call that tool. `none` is the default when no tools are present. `auto` is the
	// default if tools are present.
	ToolChoice param.Field[ChatCompletionNewParamsToolChoiceUnion] `json:"tool_choice"`
	// Tool calling configuration (Google-specific)
	ToolConfig param.Field[map[string]interface{}] `json:"tool_config"`
	// A list of tools the model may call. You can provide either custom tools or
	// function tools. All providers support tools. Adapters handle translation to
	// provider-specific formats.
	Tools param.Field[[]ChatCompletionNewParamsTool] `json:"tools"`
	// Top-k sampling parameter limiting token selection to k most likely candidates.
	// Only considers the top k highest probability tokens at each generation step,
	// setting all other tokens' probabilities to zero. Reduces tail probability mass.
	// Helps prevent selection of highly unlikely tokens, improving output coherence.
	// Supported by Google and Anthropic; not available in OpenAI's API.
	TopK param.Field[int64] `json:"top_k"`
	// An integer between 0 and 20 specifying the number of most likely tokens to
	// return at each token position, each with an associated log probability.
	// `logprobs` must be set to `true` if this parameter is used.
	TopLogprobs param.Field[int64] `json:"top_logprobs"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered. We
	// generally recommend altering this or temperature but not both.
	TopP param.Field[float64] `json:"top_p"`
	// This field is being replaced by `safety_identifier` and `prompt_cache_key`. Use
	// `prompt_cache_key` instead to maintain caching optimizations. A stable
	// identifier for your end-users. Used to boost cache hit rates by better bucketing
	// similar requests and to help OpenAI detect and prevent abuse.
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#safety-identifiers).
	User param.Field[string] `json:"user"`
	// Constrains the verbosity of the model's response. Lower values will result in
	// more concise responses, while higher values will result in more verbose
	// responses. Currently supported values are `low`, `medium`, and `high`.
	Verbosity param.Field[ChatCompletionNewParamsVerbosity] `json:"verbosity"`
	// This tool searches the web for relevant results to use in a response. Learn more
	// about the
	// [web search tool](https://platform.openai.com/docs/guides/tools-web-search?api-mode=chat).
	WebSearchOptions param.Field[map[string]interface{}] `json:"web_search_options"`
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Model(s) to use for completion. Can be a single model ID, a DedalusModel object,
// or a list for multi-model routing. Single model: 'openai/gpt-5',
// 'anthropic/claude-sonnet-4-5-20250929', 'google/gemini-3-pro-preview', or a
// DedalusModel instance. Multi-model routing: ['openai/gpt-5',
// 'anthropic/claude-sonnet-4-5-20250929', 'google/gemini-3-pro-preview'] or list
// of DedalusModel objects - agent will choose optimal model based on task
// complexity.
//
// Satisfied by [shared.UnionString], [shared.DedalusModelParam],
// [ChatCompletionNewParamsModelArray].
type ChatCompletionNewParamsModelUnion interface {
	ImplementsChatCompletionNewParamsModelUnion()
}

type ChatCompletionNewParamsModelArray []shared.DedalusModelChoiceUnionParam

func (r ChatCompletionNewParamsModelArray) ImplementsChatCompletionNewParamsModelUnion() {}

// Deprecated in favor of `tool_choice`. Controls which (if any) function is called
// by the model. `none` means the model will not call a function and instead
// generates a message. `auto` means the model can pick between generating a
// message or calling a function. Specifying a particular function via
// `{"name": "my_function"}` forces the model to call that function. `none` is the
// default when no functions are present. `auto` is the default if functions are
// present.
type ChatCompletionNewParamsFunctionCall string

const (
	ChatCompletionNewParamsFunctionCallAuto ChatCompletionNewParamsFunctionCall = "auto"
	ChatCompletionNewParamsFunctionCallNone ChatCompletionNewParamsFunctionCall = "none"
)

func (r ChatCompletionNewParamsFunctionCall) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsFunctionCallAuto, ChatCompletionNewParamsFunctionCallNone:
		return true
	}
	return false
}

// Fields:
//
// - description (optional): str
// - name (required): str
// - parameters (optional): FunctionParameters
type ChatCompletionNewParamsFunction struct {
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain
	// underscores and dashes, with a maximum length of 64.
	Name param.Field[string] `json:"name,required"`
	// A description of what the function does, used by the model to choose when and
	// how to call the function.
	Description param.Field[string] `json:"description"`
	// The parameters the functions accepts, described as a JSON Schema object. See the
	// [guide](https://platform.openai.com/docs/guides/function-calling) for examples,
	// and the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for
	// documentation about the format.
	//
	// Omitting `parameters` defines a function with an empty parameter list.
	Parameters param.Field[map[string]interface{}] `json:"parameters"`
}

func (r ChatCompletionNewParamsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
// Satisfied by [ChatCompletionNewParamsMessagesMessages], [shared.UnionString].
type ChatCompletionNewParamsMessagesUnion interface {
	ImplementsChatCompletionNewParamsMessagesUnion()
}

type ChatCompletionNewParamsMessagesMessages []ChatCompletionNewParamsMessagesMessageUnion

func (r ChatCompletionNewParamsMessagesMessages) ImplementsChatCompletionNewParamsMessagesUnion() {}

// Developer-provided instructions that the model should follow, regardless of
//
// messages sent by the user. With o1 models and newer, `developer` messages
// replace the previous `system` messages.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText],
//     Field(json_schema_extra={"title": "Content3"}), MinLen(1)]
//   - role (required): Literal['developer']
//   - name (optional): str
type ChatCompletionNewParamsMessagesMessage struct {
	// The role of the messages author, in this case `developer`.
	Role         param.Field[ChatCompletionNewParamsMessagesMessagesRole] `json:"role,required"`
	Audio        param.Field[interface{}]                                 `json:"audio"`
	Content      param.Field[interface{}]                                 `json:"content"`
	FunctionCall param.Field[interface{}]                                 `json:"function_call"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
	// The refusal message by the assistant.
	Refusal param.Field[string] `json:"refusal"`
	// Tool call that this message is responding to.
	ToolCallID param.Field[string]      `json:"tool_call_id"`
	ToolCalls  param.Field[interface{}] `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// Developer-provided instructions that the model should follow, regardless of
//
// messages sent by the user. With o1 models and newer, `developer` messages
// replace the previous `system` messages.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText],
//     Field(json_schema_extra={"title": "Content3"}), MinLen(1)]
//   - role (required): Literal['developer']
//   - name (optional): str
//
// Satisfied by
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessage],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessage],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessage],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessage],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessage],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessage],
// [ChatCompletionNewParamsMessagesMessage].
type ChatCompletionNewParamsMessagesMessageUnion interface {
	implementsChatCompletionNewParamsMessagesMessageUnion()
}

// Developer-provided instructions that the model should follow, regardless of
//
// messages sent by the user. With o1 models and newer, `developer` messages
// replace the previous `system` messages.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText],
//     Field(json_schema_extra={"title": "Content3"}), MinLen(1)]
//   - role (required): Literal['developer']
//   - name (optional): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessage struct {
	// The contents of the developer message.
	Content param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `developer`.
	Role param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageRole] `json:"role,required"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// The contents of the developer message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentUnion()
}

type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3 []ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Item

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Item struct {
	// The text content.
	Text param.Field[string] `json:"text,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Type] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Item) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Type string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3TypeText ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Type = "text"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3Type) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentContent3TypeText:
		return true
	}
	return false
}

// The role of the messages author, in this case `developer`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageRole string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageRoleDeveloper ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageRole = "developer"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageRoleDeveloper:
		return true
	}
	return false
}

// Developer-provided instructions that the model should follow, regardless of
//
// messages sent by the user. With o1 models and newer, use `developer` messages
// for this purpose instead.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestSystemMessageContentPart],
//     Field(json_schema_extra={"title": "Content4"}), MinLen(1)]
//   - role (required): Literal['system']
//   - name (optional): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessage struct {
	// The contents of the system message.
	Content param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `system`.
	Role param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageRole] `json:"role,required"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// The contents of the system message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentUnion()
}

type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4 []ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Item

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Item struct {
	// The text content.
	Text param.Field[string] `json:"text,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Type] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Item) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Type string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4TypeText ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Type = "text"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4Type) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentContent4TypeText:
		return true
	}
	return false
}

// The role of the messages author, in this case `system`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageRole string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageRoleSystem ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageRole = "system"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageRoleSystem:
		return true
	}
	return false
}

// Messages sent by an end user, containing prompts or additional context
//
// information.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestUserMessageContentPart],
//     Field(json_schema_extra={"title": "Content5"}), MinLen(1)]
//   - role (required): Literal['user']
//   - name (optional): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessage struct {
	// The contents of the user message.
	Content param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `user`.
	Role param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageRole] `json:"role,required"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// The contents of the user message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentUnion()
}

type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5 []ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Item struct {
	// The type of the content part.
	Type       param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type] `json:"type,required"`
	File       param.Field[interface{}]                                                                                `json:"file"`
	ImageURL   param.Field[interface{}]                                                                                `json:"image_url"`
	InputAudio param.Field[interface{}]                                                                                `json:"input_audio"`
	// The text content.
	Text param.Field[string] `json:"text"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Item) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Item) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
//
// Satisfied by
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartText],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImage],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudio],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFile],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Item].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion interface {
	implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion()
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartText struct {
	// The text content.
	Text param.Field[string] `json:"text,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartTextType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartText) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion() {
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartTextType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartTextTypeText ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartTextType = "text"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartTextType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartTextTypeText:
		return true
	}
	return false
}

// Learn about [image inputs](https://platform.openai.com/docs/guides/vision).
//
// Fields:
//
// - type (required): Literal['image_url']
// - image_url (required): ImageUrl
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImage struct {
	// Fields:
	//
	// - url (required): AnyUrl
	// - detail (optional): Literal['auto', 'low', 'high']
	ImageURL param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURL] `json:"image_url,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImage) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion() {
}

// Fields:
//
// - url (required): AnyUrl
// - detail (optional): Literal['auto', 'low', 'high']
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURL struct {
	// Either a URL of the image or the base64 encoded image data.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Specifies the detail level of the image. Learn more in the
	// [Vision guide](https://platform.openai.com/docs/guides/vision#low-or-high-fidelity-image-understanding).
	Detail param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetail] `json:"detail"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the detail level of the image. Learn more in the
// [Vision guide](https://platform.openai.com/docs/guides/vision#low-or-high-fidelity-image-understanding).
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetail string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetailAuto ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetail = "auto"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetailLow  ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetail = "low"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetailHigh ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetail = "high"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetail) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetailAuto, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetailLow, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageImageURLDetailHigh:
		return true
	}
	return false
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageTypeImageURL ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageType = "image_url"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartImageTypeImageURL:
		return true
	}
	return false
}

// Learn about [audio inputs](https://platform.openai.com/docs/guides/audio).
//
// Fields:
//
// - type (required): Literal['input_audio']
// - input_audio (required): InputAudio82c696db
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudio struct {
	// Fields:
	//
	// - data (required): str
	// - format (required): Literal['wav', 'mp3']
	InputAudio param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudio] `json:"input_audio,required"`
	// The type of the content part. Always `input_audio`.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudio) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion() {
}

// Fields:
//
// - data (required): str
// - format (required): Literal['wav', 'mp3']
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudio struct {
	// Base64 encoded audio data.
	Data param.Field[string] `json:"data,required"`
	// The format of the encoded audio data. Currently supports "wav" and "mp3".
	Format param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormat] `json:"format,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the encoded audio data. Currently supports "wav" and "mp3".
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormat string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormatWav ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormat = "wav"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormatMP3 ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormat = "mp3"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormat) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormatWav, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioInputAudioFormatMP3:
		return true
	}
	return false
}

// The type of the content part. Always `input_audio`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioTypeInputAudio ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioType = "input_audio"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartAudioTypeInputAudio:
		return true
	}
	return false
}

// Learn about [file inputs](https://platform.openai.com/docs/guides/text) for text
// generation.
//
// Fields:
//
// - type (required): Literal['file']
// - file (required): File
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFile struct {
	// Fields:
	//
	// - filename (optional): str
	// - file_data (optional): str
	// - file_id (optional): str
	File param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileFile] `json:"file,required"`
	// The type of the content part. Always `file`.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFile) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ItemUnion() {
}

// Fields:
//
// - filename (optional): str
// - file_data (optional): str
// - file_id (optional): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileFile struct {
	// The base64 encoded file data, used when passing the file to the model as a
	// string.
	FileData param.Field[string] `json:"file_data"`
	// The ID of an uploaded file to use as input.
	FileID param.Field[string] `json:"file_id"`
	// The name of the file, used when passing the file to the model as a string.
	Filename param.Field[string] `json:"filename"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the content part. Always `file`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileTypeFile ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileType = "file"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5ChatCompletionRequestMessageContentPartFileTypeFile:
		return true
	}
	return false
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeText       ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type = "text"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeImageURL   ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type = "image_url"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeInputAudio ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type = "input_audio"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeFile       ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type = "file"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5Type) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeText, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeImageURL, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeInputAudio, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentContent5TypeFile:
		return true
	}
	return false
}

// The role of the messages author, in this case `user`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageRole string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageRoleUser ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageRole = "user"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageRoleUser:
		return true
	}
	return false
}

// Messages sent by the model in response to user messages.
//
// Fields:
//
//   - content (optional): str |
//     Annotated[list[ChatCompletionRequestAssistantMessageContentPart],
//     Field(json_schema_extra={"title": "Content6"}), MinLen(1)] | None
//   - refusal (optional): str | None
//   - role (required): Literal['assistant']
//   - name (optional): str
//   - audio (optional): Audio815cb4c9 | None
//   - tool_calls (optional): ChatCompletionMessageToolCalls
//   - function_call (optional): FunctionCall | None
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessage struct {
	// The role of the messages author, in this case `assistant`.
	Role param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageRole] `json:"role,required"`
	// Data about a previous audio response from the model.
	//
	// [Learn more](https://platform.openai.com/docs/guides/audio).
	//
	// Fields:
	//
	// - id (required): str
	Audio param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageAudio] `json:"audio"`
	// The contents of the assistant message. Required unless `tool_calls` or
	// `function_call` is specified.
	Content param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentUnion] `json:"content"`
	// Deprecated and replaced by `tool_calls`. The name and arguments of a function
	// that should be called, as generated by the model.
	//
	// Fields:
	//
	// - arguments (required): str
	// - name (required): str
	FunctionCall param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageFunctionCall] `json:"function_call"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
	// The refusal message by the assistant.
	Refusal param.Field[string] `json:"refusal"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls param.Field[[]ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallUnion] `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// The role of the messages author, in this case `assistant`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageRole string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageRoleAssistant ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageRole = "assistant"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageRoleAssistant:
		return true
	}
	return false
}

// Data about a previous audio response from the model.
//
// [Learn more](https://platform.openai.com/docs/guides/audio).
//
// Fields:
//
// - id (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageAudio struct {
	// Unique identifier for a previous audio response from the model.
	ID param.Field[string] `json:"id,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The contents of the assistant message. Required unless `tool_calls` or
// `function_call` is specified.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentUnion()
}

type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6 []ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ItemUnion

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Item struct {
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Type] `json:"type,required"`
	// The refusal message generated by the model.
	Refusal param.Field[string] `json:"refusal"`
	// The text content.
	Text param.Field[string] `json:"text"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Item) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Item) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ItemUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
//
// Satisfied by
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartText],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusal],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Item].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ItemUnion interface {
	implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ItemUnion()
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartText struct {
	// The text content.
	Text param.Field[string] `json:"text,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartTextType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartText) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ItemUnion() {
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartTextType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartTextTypeText ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartTextType = "text"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartTextType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartTextTypeText:
		return true
	}
	return false
}

// Fields:
//
// - type (required): Literal['refusal']
// - refusal (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusal struct {
	// The refusal message generated by the model.
	Refusal param.Field[string] `json:"refusal,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusalType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusal) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusal) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ItemUnion() {
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusalType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusalTypeRefusal ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusalType = "refusal"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusalType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6ChatCompletionRequestMessageContentPartRefusalTypeRefusal:
		return true
	}
	return false
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Type string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6TypeText    ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Type = "text"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6TypeRefusal ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Type = "refusal"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6Type) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6TypeText, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentContent6TypeRefusal:
		return true
	}
	return false
}

// Deprecated and replaced by `tool_calls`. The name and arguments of a function
// that should be called, as generated by the model.
//
// Fields:
//
// - arguments (required): str
// - name (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageFunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments param.Field[string] `json:"arguments,required"`
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): FunctionD877ee33
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCall struct {
	// The ID of the tool call.
	ID param.Field[string] `json:"id,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type     param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsType] `json:"type,required"`
	Custom   param.Field[interface{}]                                                                               `json:"custom"`
	Function param.Field[interface{}]                                                                               `json:"function"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCall) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallUnion() {
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): FunctionD877ee33
//
// Satisfied by
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInput],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInput],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCall].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallUnion interface {
	implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallUnion()
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal['function']
// - function (required): FunctionD877ee33
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInput struct {
	// The ID of the tool call.
	ID param.Field[string] `json:"id,required"`
	// The function that the model called.
	Function param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputFunction] `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInput) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallUnion() {
}

// The function that the model called.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputFunction struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments param.Field[string] `json:"arguments,required"`
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputTypeFunction ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputType = "function"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageToolCallInputTypeFunction:
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
// - custom (required): Custom314518a6
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInput struct {
	// The ID of the tool call.
	ID param.Field[string] `json:"id,required"`
	// The custom tool that the model called.
	Custom param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputCustom] `json:"custom,required"`
	// The type of the tool. Always `custom`.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputType] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInput) implementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallUnion() {
}

// The custom tool that the model called.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputCustom struct {
	// The input for the custom tool call generated by the model.
	Input param.Field[string] `json:"input,required"`
	// The name of the custom tool to call.
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the tool. Always `custom`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputTypeCustom ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputType = "custom"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsChatCompletionMessageCustomToolCallInputTypeCustom:
		return true
	}
	return false
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsType string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsTypeFunction ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsType = "function"
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsTypeCustom   ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsType = "custom"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsTypeFunction, ChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageToolCallsTypeCustom:
		return true
	}
	return false
}

// Fields:
//
//   - role (required): Literal['tool']
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestToolMessageContentPart],
//     Field(json_schema_extra={"title": "Content7"}), MinLen(1)]
//   - tool_call_id (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessage struct {
	// The contents of the tool message.
	Content param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `tool`.
	Role param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageRole] `json:"role,required"`
	// Tool call that this message is responding to.
	ToolCallID param.Field[string] `json:"tool_call_id,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// The contents of the tool message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7].
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentUnion interface {
	ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentUnion()
}

type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7 []ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Item

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal['text']
// - text (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Item struct {
	// The text content.
	Text param.Field[string] `json:"text,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Type] `json:"type,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Item) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the content part.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Type string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7TypeText ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Type = "text"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7Type) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentContent7TypeText:
		return true
	}
	return false
}

// The role of the messages author, in this case `tool`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageRole string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageRoleTool ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageRole = "tool"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageRoleTool:
		return true
	}
	return false
}

// Fields:
//
// - role (required): Literal['function']
// - content (required): str | None
// - name (required): str
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessage struct {
	// The contents of the function message.
	Content param.Field[string] `json:"content,required"`
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
	// The role of the messages author, in this case `function`.
	Role param.Field[ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessageRole] `json:"role,required"`
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessage) implementsChatCompletionNewParamsMessagesMessageUnion() {
}

// The role of the messages author, in this case `function`.
type ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessageRole string

const (
	ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessageRoleFunction ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessageRole = "function"
)

func (r ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesChatCompletionRequestFunctionMessageRoleFunction:
		return true
	}
	return false
}

// The role of the messages author, in this case `developer`.
type ChatCompletionNewParamsMessagesMessagesRole string

const (
	ChatCompletionNewParamsMessagesMessagesRoleDeveloper ChatCompletionNewParamsMessagesMessagesRole = "developer"
	ChatCompletionNewParamsMessagesMessagesRoleSystem    ChatCompletionNewParamsMessagesMessagesRole = "system"
	ChatCompletionNewParamsMessagesMessagesRoleUser      ChatCompletionNewParamsMessagesMessagesRole = "user"
	ChatCompletionNewParamsMessagesMessagesRoleAssistant ChatCompletionNewParamsMessagesMessagesRole = "assistant"
	ChatCompletionNewParamsMessagesMessagesRoleTool      ChatCompletionNewParamsMessagesMessagesRole = "tool"
	ChatCompletionNewParamsMessagesMessagesRoleFunction  ChatCompletionNewParamsMessagesMessagesRole = "function"
)

func (r ChatCompletionNewParamsMessagesMessagesRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesMessagesRoleDeveloper, ChatCompletionNewParamsMessagesMessagesRoleSystem, ChatCompletionNewParamsMessagesMessagesRoleUser, ChatCompletionNewParamsMessagesMessagesRoleAssistant, ChatCompletionNewParamsMessagesMessagesRoleTool, ChatCompletionNewParamsMessagesMessagesRoleFunction:
		return true
	}
	return false
}

type ChatCompletionNewParamsModality string

const (
	ChatCompletionNewParamsModalityText  ChatCompletionNewParamsModality = "text"
	ChatCompletionNewParamsModalityAudio ChatCompletionNewParamsModality = "audio"
)

func (r ChatCompletionNewParamsModality) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsModalityText, ChatCompletionNewParamsModalityAudio:
		return true
	}
	return false
}

// Static predicted output content, such as the content of a text file that is
// being regenerated.
type ChatCompletionNewParamsPrediction struct {
	Content param.Field[map[string]interface{}]                `json:"content,required"`
	Type    param.Field[ChatCompletionNewParamsPredictionType] `json:"type"`
}

func (r ChatCompletionNewParamsPrediction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChatCompletionNewParamsPredictionType string

const (
	ChatCompletionNewParamsPredictionTypeContent ChatCompletionNewParamsPredictionType = "content"
)

func (r ChatCompletionNewParamsPredictionType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsPredictionTypeContent:
		return true
	}
	return false
}

// The retention policy for the prompt cache. Set to `24h` to enable extended
// prompt caching, which keeps cached prefixes active for longer, up to a maximum
// of 24 hours.
// [Learn more](https://platform.openai.com/docs/guides/prompt-caching#prompt-cache-retention).
type ChatCompletionNewParamsPromptCacheRetention string

const (
	ChatCompletionNewParamsPromptCacheRetention24h      ChatCompletionNewParamsPromptCacheRetention = "24h"
	ChatCompletionNewParamsPromptCacheRetentionInMemory ChatCompletionNewParamsPromptCacheRetention = "in-memory"
)

func (r ChatCompletionNewParamsPromptCacheRetention) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsPromptCacheRetention24h, ChatCompletionNewParamsPromptCacheRetentionInMemory:
		return true
	}
	return false
}

// Constrains effort on reasoning for
// [reasoning models](https://platform.openai.com/docs/guides/reasoning). Currently
// supported values are `none`, `minimal`, `low`, `medium`, and `high`. Reducing
// reasoning effort can result in faster responses and fewer tokens used on
// reasoning in a response. - `gpt-5.1` defaults to `none`, which does not perform
// reasoning. The supported reasoning values for `gpt-5.1` are `none`, `low`,
// `medium`, and `high`. Tool calls are supported for all reasoning values in
// gpt-5.1. - All models before `gpt-5.1` default to `medium` reasoning effort, and
// do not support `none`. - The `gpt-5-pro` model defaults to (and only supports)
// `high` reasoning effort.
type ChatCompletionNewParamsReasoningEffort string

const (
	ChatCompletionNewParamsReasoningEffortHigh    ChatCompletionNewParamsReasoningEffort = "high"
	ChatCompletionNewParamsReasoningEffortLow     ChatCompletionNewParamsReasoningEffort = "low"
	ChatCompletionNewParamsReasoningEffortMedium  ChatCompletionNewParamsReasoningEffort = "medium"
	ChatCompletionNewParamsReasoningEffortMinimal ChatCompletionNewParamsReasoningEffort = "minimal"
	ChatCompletionNewParamsReasoningEffortNone    ChatCompletionNewParamsReasoningEffort = "none"
)

func (r ChatCompletionNewParamsReasoningEffort) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsReasoningEffortHigh, ChatCompletionNewParamsReasoningEffortLow, ChatCompletionNewParamsReasoningEffortMedium, ChatCompletionNewParamsReasoningEffortMinimal, ChatCompletionNewParamsReasoningEffortNone:
		return true
	}
	return false
}

// An object specifying the format that the model must output. Setting to
// `{ "type": "json_schema", "json_schema": {...} }` enables Structured Outputs
// which ensures the model will match your supplied JSON schema. Learn more in the
// [Structured Outputs guide](https://platform.openai.com/docs/guides/structured-outputs).
// Setting to `{ "type": "json_object" }` enables the older JSON mode, which
// ensures the message the model generates is valid JSON. Using `json_schema` is
// preferred for models that support it.
type ChatCompletionNewParamsResponseFormat struct {
	JSONSchema param.Field[interface{}]                               `json:"json_schema"`
	Type       param.Field[ChatCompletionNewParamsResponseFormatType] `json:"type"`
}

func (r ChatCompletionNewParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsResponseFormat) ImplementsChatCompletionNewParamsResponseFormatUnion() {
}

// An object specifying the format that the model must output. Setting to
// `{ "type": "json_schema", "json_schema": {...} }` enables Structured Outputs
// which ensures the model will match your supplied JSON schema. Learn more in the
// [Structured Outputs guide](https://platform.openai.com/docs/guides/structured-outputs).
// Setting to `{ "type": "json_object" }` enables the older JSON mode, which
// ensures the message the model generates is valid JSON. Using `json_schema` is
// preferred for models that support it.
//
// Satisfied by [shared.ResponseFormatTextParam],
// [shared.ResponseFormatJSONSchemaParam], [shared.ResponseFormatJSONObjectParam],
// [ChatCompletionNewParamsResponseFormat].
type ChatCompletionNewParamsResponseFormatUnion interface {
	ImplementsChatCompletionNewParamsResponseFormatUnion()
}

type ChatCompletionNewParamsResponseFormatType string

const (
	ChatCompletionNewParamsResponseFormatTypeText       ChatCompletionNewParamsResponseFormatType = "text"
	ChatCompletionNewParamsResponseFormatTypeJSONSchema ChatCompletionNewParamsResponseFormatType = "json_schema"
	ChatCompletionNewParamsResponseFormatTypeJSONObject ChatCompletionNewParamsResponseFormatType = "json_object"
)

func (r ChatCompletionNewParamsResponseFormatType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsResponseFormatTypeText, ChatCompletionNewParamsResponseFormatTypeJSONSchema, ChatCompletionNewParamsResponseFormatTypeJSONObject:
		return true
	}
	return false
}

// Safety setting, affecting the safety-blocking behavior.
//
// Passing a safety setting for a category changes the allowed probability that
// content is blocked.
//
// Fields:
//
//   - category (required): Literal['HARM_CATEGORY_UNSPECIFIED',
//     'HARM_CATEGORY_DEROGATORY', 'HARM_CATEGORY_TOXICITY',
//     'HARM_CATEGORY_VIOLENCE', 'HARM_CATEGORY_SEXUAL', 'HARM_CATEGORY_MEDICAL',
//     'HARM_CATEGORY_DANGEROUS', 'HARM_CATEGORY_HARASSMENT',
//     'HARM_CATEGORY_HATE_SPEECH', 'HARM_CATEGORY_SEXUALLY_EXPLICIT',
//     'HARM_CATEGORY_DANGEROUS_CONTENT', 'HARM_CATEGORY_CIVIC_INTEGRITY']
//   - threshold (required): Literal['HARM_BLOCK_THRESHOLD_UNSPECIFIED',
//     'BLOCK_LOW_AND_ABOVE', 'BLOCK_MEDIUM_AND_ABOVE', 'BLOCK_ONLY_HIGH',
//     'BLOCK_NONE', 'OFF']
type ChatCompletionNewParamsSafetySetting struct {
	// Required. The category for this setting.
	Category param.Field[ChatCompletionNewParamsSafetySettingsCategory] `json:"category,required"`
	// Required. Controls the probability threshold at which harm is blocked.
	Threshold param.Field[ChatCompletionNewParamsSafetySettingsThreshold] `json:"threshold,required"`
}

func (r ChatCompletionNewParamsSafetySetting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required. The category for this setting.
type ChatCompletionNewParamsSafetySettingsCategory string

const (
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryUnspecified      ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_UNSPECIFIED"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryDerogatory       ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_DEROGATORY"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryToxicity         ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_TOXICITY"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryViolence         ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_VIOLENCE"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategorySexual           ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_SEXUAL"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryMedical          ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_MEDICAL"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryDangerous        ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_DANGEROUS"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryHarassment       ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_HARASSMENT"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryHateSpeech       ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_HATE_SPEECH"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategorySexuallyExplicit ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_SEXUALLY_EXPLICIT"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryDangerousContent ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_DANGEROUS_CONTENT"
	ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryCivicIntegrity   ChatCompletionNewParamsSafetySettingsCategory = "HARM_CATEGORY_CIVIC_INTEGRITY"
)

func (r ChatCompletionNewParamsSafetySettingsCategory) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryUnspecified, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryDerogatory, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryToxicity, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryViolence, ChatCompletionNewParamsSafetySettingsCategoryHarmCategorySexual, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryMedical, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryDangerous, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryHarassment, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryHateSpeech, ChatCompletionNewParamsSafetySettingsCategoryHarmCategorySexuallyExplicit, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryDangerousContent, ChatCompletionNewParamsSafetySettingsCategoryHarmCategoryCivicIntegrity:
		return true
	}
	return false
}

// Required. Controls the probability threshold at which harm is blocked.
type ChatCompletionNewParamsSafetySettingsThreshold string

const (
	ChatCompletionNewParamsSafetySettingsThresholdHarmBlockThresholdUnspecified ChatCompletionNewParamsSafetySettingsThreshold = "HARM_BLOCK_THRESHOLD_UNSPECIFIED"
	ChatCompletionNewParamsSafetySettingsThresholdBlockLowAndAbove              ChatCompletionNewParamsSafetySettingsThreshold = "BLOCK_LOW_AND_ABOVE"
	ChatCompletionNewParamsSafetySettingsThresholdBlockMediumAndAbove           ChatCompletionNewParamsSafetySettingsThreshold = "BLOCK_MEDIUM_AND_ABOVE"
	ChatCompletionNewParamsSafetySettingsThresholdBlockOnlyHigh                 ChatCompletionNewParamsSafetySettingsThreshold = "BLOCK_ONLY_HIGH"
	ChatCompletionNewParamsSafetySettingsThresholdBlockNone                     ChatCompletionNewParamsSafetySettingsThreshold = "BLOCK_NONE"
	ChatCompletionNewParamsSafetySettingsThresholdOff                           ChatCompletionNewParamsSafetySettingsThreshold = "OFF"
)

func (r ChatCompletionNewParamsSafetySettingsThreshold) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsSafetySettingsThresholdHarmBlockThresholdUnspecified, ChatCompletionNewParamsSafetySettingsThresholdBlockLowAndAbove, ChatCompletionNewParamsSafetySettingsThresholdBlockMediumAndAbove, ChatCompletionNewParamsSafetySettingsThresholdBlockOnlyHigh, ChatCompletionNewParamsSafetySettingsThresholdBlockNone, ChatCompletionNewParamsSafetySettingsThresholdOff:
		return true
	}
	return false
}

// Service tier for request processing
type ChatCompletionNewParamsServiceTier string

const (
	ChatCompletionNewParamsServiceTierAuto         ChatCompletionNewParamsServiceTier = "auto"
	ChatCompletionNewParamsServiceTierDefault      ChatCompletionNewParamsServiceTier = "default"
	ChatCompletionNewParamsServiceTierFlex         ChatCompletionNewParamsServiceTier = "flex"
	ChatCompletionNewParamsServiceTierPriority     ChatCompletionNewParamsServiceTier = "priority"
	ChatCompletionNewParamsServiceTierScale        ChatCompletionNewParamsServiceTier = "scale"
	ChatCompletionNewParamsServiceTierStandardOnly ChatCompletionNewParamsServiceTier = "standard_only"
)

func (r ChatCompletionNewParamsServiceTier) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsServiceTierAuto, ChatCompletionNewParamsServiceTierDefault, ChatCompletionNewParamsServiceTierFlex, ChatCompletionNewParamsServiceTierPriority, ChatCompletionNewParamsServiceTierScale, ChatCompletionNewParamsServiceTierStandardOnly:
		return true
	}
	return false
}

// Not supported with latest reasoning models 'o3' and 'o4-mini'. Up to 4 sequences
// where the API will stop generating further tokens; the returned text will not
// contain the stop sequence.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionNewParamsStopStopconfiguration].
type ChatCompletionNewParamsStopUnion interface {
	ImplementsChatCompletionNewParamsStopUnion()
}

type ChatCompletionNewParamsStopStopconfiguration []string

func (r ChatCompletionNewParamsStopStopconfiguration) ImplementsChatCompletionNewParamsStopUnion() {}

// System-level instructions defining the assistant's behavior, role, and
// constraints. Sets the context and personality for the entire conversation.
// Different from user/assistant messages as it provides meta-instructions about
// how to respond rather than conversation content. OpenAI: Provided as system role
// message in messages array. Google: Top-level systemInstruction field (adapter
// extracts from messages). Anthropic: Top-level system parameter (adapter extracts
// from messages). Accepts both string and structured object formats depending on
// provider capabilities.
//
// Satisfied by [ChatCompletionNewParamsSystemInstructionMap],
// [shared.UnionString].
type ChatCompletionNewParamsSystemInstructionUnion interface {
	ImplementsChatCompletionNewParamsSystemInstructionUnion()
}

type ChatCompletionNewParamsSystemInstructionMap map[string]interface{}

func (r ChatCompletionNewParamsSystemInstructionMap) ImplementsChatCompletionNewParamsSystemInstructionUnion() {
}

// Extended thinking configuration (Anthropic-specific)
type ChatCompletionNewParamsThinking struct {
	BudgetTokens param.Field[int64]                               `json:"budget_tokens"`
	Type         param.Field[ChatCompletionNewParamsThinkingType] `json:"type"`
}

func (r ChatCompletionNewParamsThinking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsThinking) implementsChatCompletionNewParamsThinkingUnion() {}

// Extended thinking configuration (Anthropic-specific)
//
// Satisfied by [ChatCompletionNewParamsThinkingThinkingConfigEnabled],
// [ChatCompletionNewParamsThinkingThinkingConfigDisabled],
// [ChatCompletionNewParamsThinking].
type ChatCompletionNewParamsThinkingUnion interface {
	implementsChatCompletionNewParamsThinkingUnion()
}

type ChatCompletionNewParamsThinkingThinkingConfigEnabled struct {
	BudgetTokens param.Field[int64]                                                    `json:"budget_tokens,required"`
	Type         param.Field[ChatCompletionNewParamsThinkingThinkingConfigEnabledType] `json:"type"`
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

type ChatCompletionNewParamsThinkingThinkingConfigDisabled struct {
	Type param.Field[ChatCompletionNewParamsThinkingThinkingConfigDisabledType] `json:"type"`
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

type ChatCompletionNewParamsThinkingType string

const (
	ChatCompletionNewParamsThinkingTypeEnabled  ChatCompletionNewParamsThinkingType = "enabled"
	ChatCompletionNewParamsThinkingTypeDisabled ChatCompletionNewParamsThinkingType = "disabled"
)

func (r ChatCompletionNewParamsThinkingType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsThinkingTypeEnabled, ChatCompletionNewParamsThinkingTypeDisabled:
		return true
	}
	return false
}

// Controls which (if any) tool is called by the model. `none` means the model will
// not call any tool and instead generates a message. `auto` means the model can
// pick between generating a message or calling one or more tools. `required` means
// the model must call one or more tools. Specifying a particular tool via
// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
// call that tool. `none` is the default when no tools are present. `auto` is the
// default if tools are present.
type ChatCompletionNewParamsToolChoice struct {
	DisableParallelToolUse param.Field[bool]                                  `json:"disable_parallel_tool_use"`
	Name                   param.Field[string]                                `json:"name"`
	Type                   param.Field[ChatCompletionNewParamsToolChoiceType] `json:"type"`
}

func (r ChatCompletionNewParamsToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolChoice) implementsChatCompletionNewParamsToolChoiceUnion() {}

// Controls which (if any) tool is called by the model. `none` means the model will
// not call any tool and instead generates a message. `auto` means the model can
// pick between generating a message or calling one or more tools. `required` means
// the model must call one or more tools. Specifying a particular tool via
// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
// call that tool. `none` is the default when no tools are present. `auto` is the
// default if tools are present.
//
// Satisfied by [ChatCompletionNewParamsToolChoiceToolChoiceAuto],
// [ChatCompletionNewParamsToolChoiceToolChoiceAny],
// [ChatCompletionNewParamsToolChoiceToolChoiceTool],
// [ChatCompletionNewParamsToolChoiceToolChoiceNone],
// [ChatCompletionNewParamsToolChoice].
type ChatCompletionNewParamsToolChoiceUnion interface {
	implementsChatCompletionNewParamsToolChoiceUnion()
}

// The model will automatically decide whether to use tools.
type ChatCompletionNewParamsToolChoiceToolChoiceAuto struct {
	DisableParallelToolUse param.Field[bool]                                                `json:"disable_parallel_tool_use"`
	Type                   param.Field[ChatCompletionNewParamsToolChoiceToolChoiceAutoType] `json:"type"`
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceAuto) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceAuto) implementsChatCompletionNewParamsToolChoiceUnion() {
}

type ChatCompletionNewParamsToolChoiceToolChoiceAutoType string

const (
	ChatCompletionNewParamsToolChoiceToolChoiceAutoTypeAuto ChatCompletionNewParamsToolChoiceToolChoiceAutoType = "auto"
)

func (r ChatCompletionNewParamsToolChoiceToolChoiceAutoType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolChoiceToolChoiceAutoTypeAuto:
		return true
	}
	return false
}

// The model will use any available tools.
type ChatCompletionNewParamsToolChoiceToolChoiceAny struct {
	DisableParallelToolUse param.Field[bool]                                               `json:"disable_parallel_tool_use"`
	Type                   param.Field[ChatCompletionNewParamsToolChoiceToolChoiceAnyType] `json:"type"`
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceAny) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceAny) implementsChatCompletionNewParamsToolChoiceUnion() {
}

type ChatCompletionNewParamsToolChoiceToolChoiceAnyType string

const (
	ChatCompletionNewParamsToolChoiceToolChoiceAnyTypeAny ChatCompletionNewParamsToolChoiceToolChoiceAnyType = "any"
)

func (r ChatCompletionNewParamsToolChoiceToolChoiceAnyType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolChoiceToolChoiceAnyTypeAny:
		return true
	}
	return false
}

// The model will use the specified tool with `tool_choice.name`.
type ChatCompletionNewParamsToolChoiceToolChoiceTool struct {
	Name                   param.Field[string]                                              `json:"name,required"`
	DisableParallelToolUse param.Field[bool]                                                `json:"disable_parallel_tool_use"`
	Type                   param.Field[ChatCompletionNewParamsToolChoiceToolChoiceToolType] `json:"type"`
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceTool) implementsChatCompletionNewParamsToolChoiceUnion() {
}

type ChatCompletionNewParamsToolChoiceToolChoiceToolType string

const (
	ChatCompletionNewParamsToolChoiceToolChoiceToolTypeTool ChatCompletionNewParamsToolChoiceToolChoiceToolType = "tool"
)

func (r ChatCompletionNewParamsToolChoiceToolChoiceToolType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolChoiceToolChoiceToolTypeTool:
		return true
	}
	return false
}

// The model will not be allowed to use tools.
type ChatCompletionNewParamsToolChoiceToolChoiceNone struct {
	Type param.Field[ChatCompletionNewParamsToolChoiceToolChoiceNoneType] `json:"type"`
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceNone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolChoiceToolChoiceNone) implementsChatCompletionNewParamsToolChoiceUnion() {
}

type ChatCompletionNewParamsToolChoiceToolChoiceNoneType string

const (
	ChatCompletionNewParamsToolChoiceToolChoiceNoneTypeNone ChatCompletionNewParamsToolChoiceToolChoiceNoneType = "none"
)

func (r ChatCompletionNewParamsToolChoiceToolChoiceNoneType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolChoiceToolChoiceNoneTypeNone:
		return true
	}
	return false
}

type ChatCompletionNewParamsToolChoiceType string

const (
	ChatCompletionNewParamsToolChoiceTypeAuto ChatCompletionNewParamsToolChoiceType = "auto"
	ChatCompletionNewParamsToolChoiceTypeAny  ChatCompletionNewParamsToolChoiceType = "any"
	ChatCompletionNewParamsToolChoiceTypeTool ChatCompletionNewParamsToolChoiceType = "tool"
	ChatCompletionNewParamsToolChoiceTypeNone ChatCompletionNewParamsToolChoiceType = "none"
)

func (r ChatCompletionNewParamsToolChoiceType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolChoiceTypeAuto, ChatCompletionNewParamsToolChoiceTypeAny, ChatCompletionNewParamsToolChoiceTypeTool, ChatCompletionNewParamsToolChoiceTypeNone:
		return true
	}
	return false
}

// A function tool that can be used to generate a response.
//
// Fields:
//
// - type (required): Literal['function']
// - function (required): FunctionObject
type ChatCompletionNewParamsTool struct {
	// Fields:
	//
	// - description (optional): str
	// - name (required): str
	// - parameters (optional): FunctionParameters
	// - strict (optional): bool | None
	Function param.Field[ChatCompletionNewParamsToolsFunction] `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type param.Field[ChatCompletionNewParamsToolsType] `json:"type,required"`
}

func (r ChatCompletionNewParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Fields:
//
// - description (optional): str
// - name (required): str
// - parameters (optional): FunctionParameters
// - strict (optional): bool | None
type ChatCompletionNewParamsToolsFunction struct {
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain
	// underscores and dashes, with a maximum length of 64.
	Name param.Field[string] `json:"name,required"`
	// A description of what the function does, used by the model to choose when and
	// how to call the function.
	Description param.Field[string] `json:"description"`
	// The parameters the functions accepts, described as a JSON Schema object. See the
	// [guide](https://platform.openai.com/docs/guides/function-calling) for examples,
	// and the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for
	// documentation about the format.
	//
	// Omitting `parameters` defines a function with an empty parameter list.
	Parameters param.Field[map[string]interface{}] `json:"parameters"`
	// Whether to enable strict schema adherence when generating the function call. If
	// set to true, the model will follow the exact schema defined in the `parameters`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. Learn
	// more about Structured Outputs in the
	// [function calling guide](https://platform.openai.com/docs/guides/function-calling).
	Strict param.Field[bool] `json:"strict"`
}

func (r ChatCompletionNewParamsToolsFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionNewParamsToolsType string

const (
	ChatCompletionNewParamsToolsTypeFunction ChatCompletionNewParamsToolsType = "function"
)

func (r ChatCompletionNewParamsToolsType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsTypeFunction:
		return true
	}
	return false
}

// Constrains the verbosity of the model's response. Lower values will result in
// more concise responses, while higher values will result in more verbose
// responses. Currently supported values are `low`, `medium`, and `high`.
type ChatCompletionNewParamsVerbosity string

const (
	ChatCompletionNewParamsVerbosityHigh   ChatCompletionNewParamsVerbosity = "high"
	ChatCompletionNewParamsVerbosityLow    ChatCompletionNewParamsVerbosity = "low"
	ChatCompletionNewParamsVerbosityMedium ChatCompletionNewParamsVerbosity = "medium"
)

func (r ChatCompletionNewParamsVerbosity) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsVerbosityHigh, ChatCompletionNewParamsVerbosityLow, ChatCompletionNewParamsVerbosityMedium:
		return true
	}
	return false
}
