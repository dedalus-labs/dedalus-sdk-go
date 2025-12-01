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

// A URL citation when using web search.
//
// Fields:
//
// - type (required): Literal["url_citation"]
// - url_citation (required): UrlCitation
type Annotation struct {
	// The type of the URL citation. Always `url_citation`.
	Type AnnotationType `json:"type,required"`
	// A URL citation when using web search.
	URLCitation URLCitation    `json:"url_citation,required"`
	JSON        annotationJSON `json:"-"`
}

// annotationJSON contains the JSON metadata for the struct [Annotation]
type annotationJSON struct {
	Type        apijson.Field
	URLCitation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Annotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationJSON) RawJSON() string {
	return r.raw
}

// The type of the URL citation. Always `url_citation`.
type AnnotationType string

const (
	AnnotationTypeURLCitation AnnotationType = "url_citation"
)

func (r AnnotationType) IsKnown() bool {
	switch r {
	case AnnotationTypeURLCitation:
		return true
	}
	return false
}

// Messages sent by the model in response to user messages.
//
// Fields:
//
//   - content (optional): str |
//     Annotated[list[ChatCompletionRequestAssistantMessageContentPart], MinLen(1),
//     ArrayTitle("ChatCompletionRequestAssistantMessageContentArray")] | None
//   - refusal (optional): str | None
//   - role (required): Literal["assistant"]
//   - name (optional): str
//   - audio (optional): Audio | None
//   - tool_calls (optional): ChatCompletionMessageToolCalls
//   - function_call (optional): ChatCompletionRequestAssistantMessageFunctionCall |
//     None
type ChatCompletionAssistantMessageParam struct {
	// The role of the messages author, in this case `assistant`.
	Role param.Field[ChatCompletionAssistantMessageParamRole] `json:"role,required"`
	// Data about a previous audio response from the model.
	// [Learn more](https://platform.openai.com/docs/guides/audio).
	//
	// Fields:
	//
	// - id (required): str
	Audio param.Field[ChatCompletionAudioParam] `json:"audio"`
	// The contents of the assistant message. Required unless `tool_calls` or
	// `function_call` is specified.
	Content param.Field[ChatCompletionAssistantMessageParamContentUnion] `json:"content"`
	// Deprecated and replaced by `tool_calls`. The name and arguments of a function
	// that should be called, as generated by the model.
	//
	// Fields:
	//
	// - arguments (required): str
	// - name (required): str
	FunctionCall param.Field[ChatCompletionAssistantMessageParamFunctionCall] `json:"function_call"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
	// The refusal message by the assistant.
	Refusal param.Field[string] `json:"refusal"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls param.Field[[]ChatCompletionAssistantMessageParamToolCallsUnion] `json:"tool_calls"`
}

func (r ChatCompletionAssistantMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionAssistantMessageParam) implementsChatCompletionNewParamsMessageUnion() {}

// The role of the messages author, in this case `assistant`.
type ChatCompletionAssistantMessageParamRole string

const (
	ChatCompletionAssistantMessageParamRoleAssistant ChatCompletionAssistantMessageParamRole = "assistant"
)

func (r ChatCompletionAssistantMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionAssistantMessageParamRoleAssistant:
		return true
	}
	return false
}

// The contents of the assistant message. Required unless `tool_calls` or
// `function_call` is specified.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArray].
type ChatCompletionAssistantMessageParamContentUnion interface {
	ImplementsChatCompletionAssistantMessageParamContentUnion()
}

type ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArray []ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayUnionItem

func (r ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArray) ImplementsChatCompletionAssistantMessageParamContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal["text"]
// - text (required): str
type ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayItem struct {
	// The type of the content part.
	Type param.Field[ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayType] `json:"type,required"`
	// The refusal message generated by the model.
	Refusal param.Field[string] `json:"refusal"`
	// The text content.
	Text param.Field[string] `json:"text"`
}

func (r ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayItem) implementsChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayUnionItem() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal["text"]
// - text (required): str
//
// Satisfied by [ChatCompletionContentPartTextParam],
// [ChatCompletionContentPartRefusalParam],
// [ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayItem].
type ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayUnionItem interface {
	implementsChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayUnionItem()
}

// The type of the content part.
type ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayType string

const (
	ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayTypeText    ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayType = "text"
	ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayTypeRefusal ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayType = "refusal"
)

func (r ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayType) IsKnown() bool {
	switch r {
	case ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayTypeText, ChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayTypeRefusal:
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
type ChatCompletionAssistantMessageParamFunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments param.Field[string] `json:"arguments,required"`
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionAssistantMessageParamFunctionCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["function"]
// - function (required): ChatCompletionMessageToolCallFunction
type ChatCompletionAssistantMessageParamToolCall struct {
	// The ID of the tool call.
	ID param.Field[string] `json:"id,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type     param.Field[ChatCompletionAssistantMessageParamToolCallsType] `json:"type,required"`
	Custom   param.Field[interface{}]                                      `json:"custom"`
	Function param.Field[interface{}]                                      `json:"function"`
}

func (r ChatCompletionAssistantMessageParamToolCall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionAssistantMessageParamToolCall) implementsChatCompletionAssistantMessageParamToolCallsUnion() {
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["function"]
// - function (required): ChatCompletionMessageToolCallFunction
//
// Satisfied by
// [ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInput],
// [ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInput],
// [ChatCompletionAssistantMessageParamToolCall].
type ChatCompletionAssistantMessageParamToolCallsUnion interface {
	implementsChatCompletionAssistantMessageParamToolCallsUnion()
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["function"]
// - function (required): ChatCompletionMessageToolCallFunction
type ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInput struct {
	// The ID of the tool call.
	ID param.Field[string] `json:"id,required"`
	// The function that the model called.
	Function param.Field[ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputFunction] `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type param.Field[ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputType] `json:"type,required"`
}

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInput) implementsChatCompletionAssistantMessageParamToolCallsUnion() {
}

// The function that the model called.
type ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputFunction struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments param.Field[string] `json:"arguments,required"`
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputType string

const (
	ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputTypeFunction ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputType = "function"
)

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputType) IsKnown() bool {
	switch r {
	case ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageToolCallInputTypeFunction:
		return true
	}
	return false
}

// A call to a custom tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["custom"]
// - custom (required): ChatCompletionMessageCustomToolCallCustom
type ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInput struct {
	// The ID of the tool call.
	ID param.Field[string] `json:"id,required"`
	// The custom tool that the model called.
	Custom param.Field[ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputCustom] `json:"custom,required"`
	// The type of the tool. Always `custom`.
	Type param.Field[ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputType] `json:"type,required"`
}

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInput) implementsChatCompletionAssistantMessageParamToolCallsUnion() {
}

// The custom tool that the model called.
type ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputCustom struct {
	// The input for the custom tool call generated by the model.
	Input param.Field[string] `json:"input,required"`
	// The name of the custom tool to call.
	Name param.Field[string] `json:"name,required"`
}

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the tool. Always `custom`.
type ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputType string

const (
	ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputTypeCustom ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputType = "custom"
)

func (r ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputType) IsKnown() bool {
	switch r {
	case ChatCompletionAssistantMessageParamToolCallsChatCompletionMessageCustomToolCallInputTypeCustom:
		return true
	}
	return false
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionAssistantMessageParamToolCallsType string

const (
	ChatCompletionAssistantMessageParamToolCallsTypeFunction ChatCompletionAssistantMessageParamToolCallsType = "function"
	ChatCompletionAssistantMessageParamToolCallsTypeCustom   ChatCompletionAssistantMessageParamToolCallsType = "custom"
)

func (r ChatCompletionAssistantMessageParamToolCallsType) IsKnown() bool {
	switch r {
	case ChatCompletionAssistantMessageParamToolCallsTypeFunction, ChatCompletionAssistantMessageParamToolCallsTypeCustom:
		return true
	}
	return false
}

// If the audio output modality is requested, this object contains data about the
// audio response from the model.
// [Learn more](https://platform.openai.com/docs/guides/audio).
//
// Fields:
//
// - id (required): str
// - expires_at (required): int
// - data (required): str
// - transcript (required): str
type ChatCompletionAudio struct {
	// Unique identifier for this audio response.
	ID string `json:"id,required"`
	// Base64 encoded audio bytes generated by the model, in the format specified in
	// the request.
	Data string `json:"data,required"`
	// The Unix timestamp (in seconds) for when this audio response will no longer be
	// accessible on the server for use in multi-turn conversations.
	ExpiresAt int64 `json:"expires_at,required"`
	// Transcript of the audio generated by the model.
	Transcript string                  `json:"transcript,required"`
	JSON       chatCompletionAudioJSON `json:"-"`
}

// chatCompletionAudioJSON contains the JSON metadata for the struct
// [ChatCompletionAudio]
type chatCompletionAudioJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	ExpiresAt   apijson.Field
	Transcript  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletionAudio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionAudioJSON) RawJSON() string {
	return r.raw
}

// Data about a previous audio response from the model.
// [Learn more](https://platform.openai.com/docs/guides/audio).
//
// Fields:
//
// - id (required): str
type ChatCompletionAudioParam struct {
	// Unique identifier for a previous audio response from the model.
	ID param.Field[string] `json:"id,required"`
}

func (r ChatCompletionAudioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Learn about [audio inputs](https://platform.openai.com/docs/guides/audio).
//
// Fields:
//
// - type (required): Literal["input_audio"]
// - input_audio (required): ChatCompletionRequestMessageContentPartAudioInputAudio
type ChatCompletionContentPartAudioParam struct {
	// Schema for ChatCompletionRequestMessageContentPartAudioInputAudio.
	//
	// Fields:
	//
	// - data (required): str
	// - format (required): Literal["wav", "mp3"]
	InputAudio param.Field[ChatCompletionContentPartAudioParamInputAudio] `json:"input_audio,required"`
	// The type of the content part. Always `input_audio`.
	Type param.Field[ChatCompletionContentPartAudioParamType] `json:"type,required"`
}

func (r ChatCompletionContentPartAudioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionContentPartAudioParam) implementsChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem() {
}

// Schema for ChatCompletionRequestMessageContentPartAudioInputAudio.
//
// Fields:
//
// - data (required): str
// - format (required): Literal["wav", "mp3"]
type ChatCompletionContentPartAudioParamInputAudio struct {
	// Base64 encoded audio data.
	Data param.Field[string] `json:"data,required"`
	// The format of the encoded audio data. Currently supports "wav" and "mp3".
	Format param.Field[ChatCompletionContentPartAudioParamInputAudioFormat] `json:"format,required"`
}

func (r ChatCompletionContentPartAudioParamInputAudio) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The format of the encoded audio data. Currently supports "wav" and "mp3".
type ChatCompletionContentPartAudioParamInputAudioFormat string

const (
	ChatCompletionContentPartAudioParamInputAudioFormatWav ChatCompletionContentPartAudioParamInputAudioFormat = "wav"
	ChatCompletionContentPartAudioParamInputAudioFormatMP3 ChatCompletionContentPartAudioParamInputAudioFormat = "mp3"
)

func (r ChatCompletionContentPartAudioParamInputAudioFormat) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartAudioParamInputAudioFormatWav, ChatCompletionContentPartAudioParamInputAudioFormatMP3:
		return true
	}
	return false
}

// The type of the content part. Always `input_audio`.
type ChatCompletionContentPartAudioParamType string

const (
	ChatCompletionContentPartAudioParamTypeInputAudio ChatCompletionContentPartAudioParamType = "input_audio"
)

func (r ChatCompletionContentPartAudioParamType) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartAudioParamTypeInputAudio:
		return true
	}
	return false
}

// Learn about [file inputs](https://platform.openai.com/docs/guides/text) for text
// generation.
//
// Fields:
//
// - type (required): Literal["file"]
// - file (required): ChatCompletionRequestMessageContentPartFileFile
type ChatCompletionContentPartFileParam struct {
	// Schema for ChatCompletionRequestMessageContentPartFileFile.
	//
	// Fields:
	//
	// - filename (optional): str
	// - file_data (optional): str
	// - file_id (optional): str
	File param.Field[ChatCompletionContentPartFileParamFile] `json:"file,required"`
	// The type of the content part. Always `file`.
	Type param.Field[ChatCompletionContentPartFileParamType] `json:"type,required"`
}

func (r ChatCompletionContentPartFileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionContentPartFileParam) implementsChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem() {
}

// Schema for ChatCompletionRequestMessageContentPartFileFile.
//
// Fields:
//
// - filename (optional): str
// - file_data (optional): str
// - file_id (optional): str
type ChatCompletionContentPartFileParamFile struct {
	// The base64 encoded file data, used when passing the file to the model as a
	// string.
	FileData param.Field[string] `json:"file_data"`
	// The ID of an uploaded file to use as input.
	FileID param.Field[string] `json:"file_id"`
	// The name of the file, used when passing the file to the model as a string.
	Filename param.Field[string] `json:"filename"`
}

func (r ChatCompletionContentPartFileParamFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the content part. Always `file`.
type ChatCompletionContentPartFileParamType string

const (
	ChatCompletionContentPartFileParamTypeFile ChatCompletionContentPartFileParamType = "file"
)

func (r ChatCompletionContentPartFileParamType) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartFileParamTypeFile:
		return true
	}
	return false
}

// Learn about [image inputs](https://platform.openai.com/docs/guides/vision).
//
// Fields:
//
// - type (required): Literal["image_url"]
// - image_url (required): ChatCompletionRequestMessageContentPartImageImageUrl
type ChatCompletionContentPartImageParam struct {
	// Schema for ChatCompletionRequestMessageContentPartImageImageUrl.
	//
	// Fields:
	//
	// - url (required): AnyUrl
	// - detail (optional): Literal["auto", "low", "high"]
	ImageURL param.Field[ChatCompletionContentPartImageParamImageURL] `json:"image_url,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionContentPartImageParamType] `json:"type,required"`
}

func (r ChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionContentPartImageParam) implementsChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem() {
}

// Schema for ChatCompletionRequestMessageContentPartImageImageUrl.
//
// Fields:
//
// - url (required): AnyUrl
// - detail (optional): Literal["auto", "low", "high"]
type ChatCompletionContentPartImageParamImageURL struct {
	// Either a URL of the image or the base64 encoded image data.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// Specifies the detail level of the image. Learn more in the
	// [Vision guide](https://platform.openai.com/docs/guides/vision#low-or-high-fidelity-image-understanding).
	Detail param.Field[ChatCompletionContentPartImageParamImageURLDetail] `json:"detail"`
}

func (r ChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the detail level of the image. Learn more in the
// [Vision guide](https://platform.openai.com/docs/guides/vision#low-or-high-fidelity-image-understanding).
type ChatCompletionContentPartImageParamImageURLDetail string

const (
	ChatCompletionContentPartImageParamImageURLDetailAuto ChatCompletionContentPartImageParamImageURLDetail = "auto"
	ChatCompletionContentPartImageParamImageURLDetailLow  ChatCompletionContentPartImageParamImageURLDetail = "low"
	ChatCompletionContentPartImageParamImageURLDetailHigh ChatCompletionContentPartImageParamImageURLDetail = "high"
)

func (r ChatCompletionContentPartImageParamImageURLDetail) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartImageParamImageURLDetailAuto, ChatCompletionContentPartImageParamImageURLDetailLow, ChatCompletionContentPartImageParamImageURLDetailHigh:
		return true
	}
	return false
}

// The type of the content part.
type ChatCompletionContentPartImageParamType string

const (
	ChatCompletionContentPartImageParamTypeImageURL ChatCompletionContentPartImageParamType = "image_url"
)

func (r ChatCompletionContentPartImageParamType) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartImageParamTypeImageURL:
		return true
	}
	return false
}

// Schema for ChatCompletionRequestMessageContentPartRefusal.
//
// Fields:
//
// - type (required): Literal["refusal"]
// - refusal (required): str
type ChatCompletionContentPartRefusalParam struct {
	// The refusal message generated by the model.
	Refusal param.Field[string] `json:"refusal,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionContentPartRefusalParamType] `json:"type,required"`
}

func (r ChatCompletionContentPartRefusalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionContentPartRefusalParam) implementsChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayUnionItem() {
}

// The type of the content part.
type ChatCompletionContentPartRefusalParamType string

const (
	ChatCompletionContentPartRefusalParamTypeRefusal ChatCompletionContentPartRefusalParamType = "refusal"
)

func (r ChatCompletionContentPartRefusalParamType) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartRefusalParamTypeRefusal:
		return true
	}
	return false
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal["text"]
// - text (required): str
type ChatCompletionContentPartTextParam struct {
	// The text content.
	Text param.Field[string] `json:"text,required"`
	// The type of the content part.
	Type param.Field[ChatCompletionContentPartTextParamType] `json:"type,required"`
}

func (r ChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionContentPartTextParam) implementsChatCompletionAssistantMessageParamContentChatCompletionRequestAssistantMessageContentArrayUnionItem() {
}

func (r ChatCompletionContentPartTextParam) implementsChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem() {
}

// The type of the content part.
type ChatCompletionContentPartTextParamType string

const (
	ChatCompletionContentPartTextParamTypeText ChatCompletionContentPartTextParamType = "text"
)

func (r ChatCompletionContentPartTextParamType) IsKnown() bool {
	switch r {
	case ChatCompletionContentPartTextParamTypeText:
		return true
	}
	return false
}

// Developer-provided instructions that the model should follow, regardless of
// messages sent by the user. With o1 models and newer, `developer` messages
// replace the previous `system` messages.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText], MinLen(1),
//     ArrayTitle("ChatCompletionRequestDeveloperMessageContentArray")]
//   - role (required): Literal["developer"]
//   - name (optional): str
type ChatCompletionDeveloperMessageParam struct {
	// The contents of the developer message.
	Content param.Field[ChatCompletionDeveloperMessageParamContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `developer`.
	Role param.Field[ChatCompletionDeveloperMessageParamRole] `json:"role,required"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionDeveloperMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionDeveloperMessageParam) implementsChatCompletionNewParamsMessageUnion() {}

// The contents of the developer message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionDeveloperMessageParamContentChatCompletionRequestDeveloperMessageContentArray].
type ChatCompletionDeveloperMessageParamContentUnion interface {
	ImplementsChatCompletionDeveloperMessageParamContentUnion()
}

type ChatCompletionDeveloperMessageParamContentChatCompletionRequestDeveloperMessageContentArray []ChatCompletionContentPartTextParam

func (r ChatCompletionDeveloperMessageParamContentChatCompletionRequestDeveloperMessageContentArray) ImplementsChatCompletionDeveloperMessageParamContentUnion() {
}

// The role of the messages author, in this case `developer`.
type ChatCompletionDeveloperMessageParamRole string

const (
	ChatCompletionDeveloperMessageParamRoleDeveloper ChatCompletionDeveloperMessageParamRole = "developer"
)

func (r ChatCompletionDeveloperMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionDeveloperMessageParamRoleDeveloper:
		return true
	}
	return false
}

// Schema for ChatCompletionRequestFunctionMessage.
//
// Fields:
//
// - role (required): Literal["function"]
// - content (required): str | None
// - name (required): str
type ChatCompletionFunctionMessageParam struct {
	// The contents of the function message.
	Content param.Field[string] `json:"content,required"`
	// The name of the function to call.
	Name param.Field[string] `json:"name,required"`
	// The role of the messages author, in this case `function`.
	Role param.Field[ChatCompletionFunctionMessageParamRole] `json:"role,required"`
}

func (r ChatCompletionFunctionMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionFunctionMessageParam) implementsChatCompletionNewParamsMessageUnion() {}

// The role of the messages author, in this case `function`.
type ChatCompletionFunctionMessageParamRole string

const (
	ChatCompletionFunctionMessageParamRoleFunction ChatCompletionFunctionMessageParamRole = "function"
)

func (r ChatCompletionFunctionMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionFunctionMessageParamRoleFunction:
		return true
	}
	return false
}

// Schema for ChatCompletionFunctions.
//
// Fields:
//
// - description (optional): str
// - name (required): str
// - parameters (optional): FunctionParameters
type ChatCompletionFunctionsParam struct {
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
	Parameters param.Field[shared.FunctionParameters] `json:"parameters"`
}

func (r ChatCompletionFunctionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A chat completion message generated by the model.
//
// Fields:
//
// - content (required): str | None
// - refusal (required): str | None
// - tool_calls (optional): ChatCompletionMessageToolCalls
// - annotations (optional): list[AnnotationsItem]
// - role (required): Literal["assistant"]
// - function_call (optional): FunctionCall
// - audio (optional): Audio | None
type ChatCompletionMessage struct {
	// The contents of the message.
	Content string `json:"content,required,nullable"`
	// The refusal message generated by the model.
	Refusal string `json:"refusal,required,nullable"`
	// The role of the author of this message.
	Role ChatCompletionMessageRole `json:"role,required"`
	// Annotations for the message, when applicable, as when using the
	// [web search tool](https://platform.openai.com/docs/guides/tools-web-search?api-mode=chat).
	Annotations []Annotation `json:"annotations"`
	// If the audio output modality is requested, this object contains data about the
	// audio response from the model.
	// [Learn more](https://platform.openai.com/docs/guides/audio).
	//
	// Fields:
	//
	// - id (required): str
	// - expires_at (required): int
	// - data (required): str
	// - transcript (required): str
	Audio ChatCompletionAudio `json:"audio,nullable"`
	// Deprecated and replaced by `tool_calls`. The name and arguments of a function
	// that should be called, as generated by the model.
	FunctionCall FunctionCall `json:"function_call"`
	// The tool calls generated by the model, such as function calls.
	ToolCalls []ChatCompletionMessageToolCall `json:"tool_calls"`
	JSON      chatCompletionMessageJSON       `json:"-"`
}

// chatCompletionMessageJSON contains the JSON metadata for the struct
// [ChatCompletionMessage]
type chatCompletionMessageJSON struct {
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

func (r *ChatCompletionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionMessageJSON) RawJSON() string {
	return r.raw
}

// The role of the author of this message.
type ChatCompletionMessageRole string

const (
	ChatCompletionMessageRoleAssistant ChatCompletionMessageRole = "assistant"
)

func (r ChatCompletionMessageRole) IsKnown() bool {
	switch r {
	case ChatCompletionMessageRoleAssistant:
		return true
	}
	return false
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["function"]
// - function (required): Function
type ChatCompletionMessageToolCall struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type ChatCompletionMessageToolCallsType `json:"type,required"`
	// The custom tool that the model called.
	Custom Custom `json:"custom"`
	// The function that the model called.
	Function Function                          `json:"function"`
	JSON     chatCompletionMessageToolCallJSON `json:"-"`
	union    ChatCompletionMessageToolCallsUnion
}

// chatCompletionMessageToolCallJSON contains the JSON metadata for the struct
// [ChatCompletionMessageToolCall]
type chatCompletionMessageToolCallJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Custom      apijson.Field
	Function    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r chatCompletionMessageToolCallJSON) RawJSON() string {
	return r.raw
}

func (r *ChatCompletionMessageToolCall) UnmarshalJSON(data []byte) (err error) {
	*r = ChatCompletionMessageToolCall{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ChatCompletionMessageToolCallsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ChatCompletionMessageToolCall],
// [ChatCompletionMessageCustomToolCall].
func (r ChatCompletionMessageToolCall) AsUnion() ChatCompletionMessageToolCallsUnion {
	return r.union
}

// A call to a function tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["function"]
// - function (required): Function
//
// Union satisfied by [ChatCompletionMessageToolCall] or
// [ChatCompletionMessageCustomToolCall].
type ChatCompletionMessageToolCallsUnion interface {
	implementsChatCompletionMessageToolCall()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ChatCompletionMessageToolCallsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ChatCompletionMessageToolCall{}),
			DiscriminatorValue: "function",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ChatCompletionMessageCustomToolCall{}),
			DiscriminatorValue: "custom",
		},
	)
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionMessageToolCallsType string

const (
	ChatCompletionMessageToolCallsTypeFunction ChatCompletionMessageToolCallsType = "function"
	ChatCompletionMessageToolCallsTypeCustom   ChatCompletionMessageToolCallsType = "custom"
)

func (r ChatCompletionMessageToolCallsType) IsKnown() bool {
	switch r {
	case ChatCompletionMessageToolCallsTypeFunction, ChatCompletionMessageToolCallsTypeCustom:
		return true
	}
	return false
}

// A call to a custom tool created by the model.
//
// Fields:
//
// - id (required): str
// - type (required): Literal["custom"]
// - custom (required): Custom
type ChatCompletionMessageCustomToolCall struct {
	// The ID of the tool call.
	ID string `json:"id,required"`
	// The custom tool that the model called.
	Custom Custom `json:"custom,required"`
	// The type of the tool. Always `custom`.
	Type ChatCompletionMessageCustomToolCallType `json:"type,required"`
	JSON chatCompletionMessageCustomToolCallJSON `json:"-"`
}

// chatCompletionMessageCustomToolCallJSON contains the JSON metadata for the
// struct [ChatCompletionMessageCustomToolCall]
type chatCompletionMessageCustomToolCallJSON struct {
	ID          apijson.Field
	Custom      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChatCompletionMessageCustomToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chatCompletionMessageCustomToolCallJSON) RawJSON() string {
	return r.raw
}

func (r ChatCompletionMessageCustomToolCall) implementsChatCompletionMessageToolCall() {}

// The type of the tool. Always `custom`.
type ChatCompletionMessageCustomToolCallType string

const (
	ChatCompletionMessageCustomToolCallTypeCustom ChatCompletionMessageCustomToolCallType = "custom"
)

func (r ChatCompletionMessageCustomToolCallType) IsKnown() bool {
	switch r {
	case ChatCompletionMessageCustomToolCallTypeCustom:
		return true
	}
	return false
}

// Developer-provided instructions that the model should follow, regardless of
// messages sent by the user. With o1 models and newer, use `developer` messages
// for this purpose instead.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestSystemMessageContentPart], MinLen(1),
//     ArrayTitle("ChatCompletionRequestSystemMessageContentArray")]
//   - role (required): Literal["system"]
//   - name (optional): str
type ChatCompletionSystemMessageParam struct {
	// The contents of the system message.
	Content param.Field[ChatCompletionSystemMessageParamContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `system`.
	Role param.Field[ChatCompletionSystemMessageParamRole] `json:"role,required"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionSystemMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionSystemMessageParam) implementsChatCompletionNewParamsMessageUnion() {}

// The contents of the system message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionSystemMessageParamContentChatCompletionRequestSystemMessageContentArray].
type ChatCompletionSystemMessageParamContentUnion interface {
	ImplementsChatCompletionSystemMessageParamContentUnion()
}

type ChatCompletionSystemMessageParamContentChatCompletionRequestSystemMessageContentArray []ChatCompletionContentPartTextParam

func (r ChatCompletionSystemMessageParamContentChatCompletionRequestSystemMessageContentArray) ImplementsChatCompletionSystemMessageParamContentUnion() {
}

// The role of the messages author, in this case `system`.
type ChatCompletionSystemMessageParamRole string

const (
	ChatCompletionSystemMessageParamRoleSystem ChatCompletionSystemMessageParamRole = "system"
)

func (r ChatCompletionSystemMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionSystemMessageParamRoleSystem:
		return true
	}
	return false
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

// Schema for ChatCompletionRequestToolMessage.
//
// Fields:
//
//   - role (required): Literal["tool"]
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestToolMessageContentPart], MinLen(1),
//     ArrayTitle("ChatCompletionRequestToolMessageContentArray")]
//   - tool_call_id (required): str
type ChatCompletionToolMessageParam struct {
	// The contents of the tool message.
	Content param.Field[ChatCompletionToolMessageParamContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `tool`.
	Role param.Field[ChatCompletionToolMessageParamRole] `json:"role,required"`
	// Tool call that this message is responding to.
	ToolCallID param.Field[string] `json:"tool_call_id,required"`
}

func (r ChatCompletionToolMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionToolMessageParam) implementsChatCompletionNewParamsMessageUnion() {}

// The contents of the tool message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionToolMessageParamContentChatCompletionRequestToolMessageContentArray].
type ChatCompletionToolMessageParamContentUnion interface {
	ImplementsChatCompletionToolMessageParamContentUnion()
}

type ChatCompletionToolMessageParamContentChatCompletionRequestToolMessageContentArray []ChatCompletionContentPartTextParam

func (r ChatCompletionToolMessageParamContentChatCompletionRequestToolMessageContentArray) ImplementsChatCompletionToolMessageParamContentUnion() {
}

// The role of the messages author, in this case `tool`.
type ChatCompletionToolMessageParamRole string

const (
	ChatCompletionToolMessageParamRoleTool ChatCompletionToolMessageParamRole = "tool"
)

func (r ChatCompletionToolMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionToolMessageParamRoleTool:
		return true
	}
	return false
}

// A function tool that can be used to generate a response.
//
// Fields:
//
// - type (required): Literal["function"]
// - function (required): FunctionObject
type ChatCompletionToolParam struct {
	// Schema for FunctionObject.
	//
	// Fields:
	//
	// - description (optional): str
	// - name (required): str
	// - parameters (optional): FunctionParameters
	// - strict (optional): bool | None
	Function param.Field[shared.FunctionDefinitionParam] `json:"function,required"`
	// The type of the tool. Currently, only `function` is supported.
	Type param.Field[ChatCompletionToolParamType] `json:"type,required"`
}

func (r ChatCompletionToolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionToolParam) implementsChatCompletionNewParamsToolUnion() {}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionToolParamType string

const (
	ChatCompletionToolParamTypeFunction ChatCompletionToolParamType = "function"
)

func (r ChatCompletionToolParamType) IsKnown() bool {
	switch r {
	case ChatCompletionToolParamTypeFunction:
		return true
	}
	return false
}

// Messages sent by an end user, containing prompts or additional context
// information.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestUserMessageContentPart], MinLen(1),
//     ArrayTitle("ChatCompletionRequestUserMessageContentArray")]
//   - role (required): Literal["user"]
//   - name (optional): str
type ChatCompletionUserMessageParam struct {
	// The contents of the user message.
	Content param.Field[ChatCompletionUserMessageParamContentUnion] `json:"content,required"`
	// The role of the messages author, in this case `user`.
	Role param.Field[ChatCompletionUserMessageParamRole] `json:"role,required"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionUserMessageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionUserMessageParam) implementsChatCompletionNewParamsMessageUnion() {}

// The contents of the user message.
//
// Satisfied by [shared.UnionString],
// [ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArray].
type ChatCompletionUserMessageParamContentUnion interface {
	ImplementsChatCompletionUserMessageParamContentUnion()
}

type ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArray []ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem

func (r ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArray) ImplementsChatCompletionUserMessageParamContentUnion() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal["text"]
// - text (required): str
type ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayItem struct {
	// The type of the content part.
	Type       param.Field[ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType] `json:"type,required"`
	File       param.Field[interface{}]                                                                           `json:"file"`
	ImageURL   param.Field[interface{}]                                                                           `json:"image_url"`
	InputAudio param.Field[interface{}]                                                                           `json:"input_audio"`
	// The text content.
	Text param.Field[string] `json:"text"`
}

func (r ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayItem) implementsChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem() {
}

// Learn about
// [text inputs](https://platform.openai.com/docs/guides/text-generation).
//
// Fields:
//
// - type (required): Literal["text"]
// - text (required): str
//
// Satisfied by [ChatCompletionContentPartTextParam],
// [ChatCompletionContentPartImageParam], [ChatCompletionContentPartAudioParam],
// [ChatCompletionContentPartFileParam],
// [ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayItem].
type ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem interface {
	implementsChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayUnionItem()
}

// The type of the content part.
type ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType string

const (
	ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeText       ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType = "text"
	ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeImageURL   ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType = "image_url"
	ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeInputAudio ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType = "input_audio"
	ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeFile       ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType = "file"
)

func (r ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayType) IsKnown() bool {
	switch r {
	case ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeText, ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeImageURL, ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeInputAudio, ChatCompletionUserMessageParamContentChatCompletionRequestUserMessageContentArrayTypeFile:
		return true
	}
	return false
}

// The role of the messages author, in this case `user`.
type ChatCompletionUserMessageParamRole string

const (
	ChatCompletionUserMessageParamRoleUser ChatCompletionUserMessageParamRole = "user"
)

func (r ChatCompletionUserMessageParamRole) IsKnown() bool {
	switch r {
	case ChatCompletionUserMessageParamRoleUser:
		return true
	}
	return false
}

// A chat completion choice.
//
// OpenAI-compatible choice object for non-streaming responses. Part of the
// ChatCompletion response.
type Choice struct {
	// The index of the choice in the list of choices.
	Index int64 `json:"index,required"`
	// A chat completion message generated by the model.
	Message ChatCompletionMessage `json:"message,required"`
	// The reason the model stopped generating tokens. This will be `stop` if the model
	// hit a natural stop point or a provided stop sequence, `length` if the maximum
	// number of tokens specified in the request was reached, `content_filter` if
	// content was omitted due to a flag from our content filters, `tool_calls` if the
	// model called a tool, or `function_call` (deprecated) if the model called a
	// function.
	FinishReason ChoiceFinishReason `json:"finish_reason,nullable"`
	// Log probability information for the choice.
	Logprobs ChoiceLogprobs `json:"logprobs,nullable"`
	JSON     choiceJSON     `json:"-"`
}

// choiceJSON contains the JSON metadata for the struct [Choice]
type choiceJSON struct {
	Index        apijson.Field
	Message      apijson.Field
	FinishReason apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Choice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceJSON) RawJSON() string {
	return r.raw
}

// The reason the model stopped generating tokens. This will be `stop` if the model
// hit a natural stop point or a provided stop sequence, `length` if the maximum
// number of tokens specified in the request was reached, `content_filter` if
// content was omitted due to a flag from our content filters, `tool_calls` if the
// model called a tool, or `function_call` (deprecated) if the model called a
// function.
type ChoiceFinishReason string

const (
	ChoiceFinishReasonStop          ChoiceFinishReason = "stop"
	ChoiceFinishReasonLength        ChoiceFinishReason = "length"
	ChoiceFinishReasonToolCalls     ChoiceFinishReason = "tool_calls"
	ChoiceFinishReasonContentFilter ChoiceFinishReason = "content_filter"
	ChoiceFinishReasonFunctionCall  ChoiceFinishReason = "function_call"
)

func (r ChoiceFinishReason) IsKnown() bool {
	switch r {
	case ChoiceFinishReasonStop, ChoiceFinishReasonLength, ChoiceFinishReasonToolCalls, ChoiceFinishReasonContentFilter, ChoiceFinishReasonFunctionCall:
		return true
	}
	return false
}

type ChoiceDelta struct {
	Content      string                  `json:"content,nullable"`
	FunctionCall ChoiceDeltaFunctionCall `json:"function_call,nullable"`
	Refusal      string                  `json:"refusal,nullable"`
	Role         ChoiceDeltaRole         `json:"role,nullable"`
	ToolCalls    []ChoiceDeltaToolCall   `json:"tool_calls,nullable"`
	ExtraFields  map[string]interface{}  `json:"-,extras"`
	JSON         choiceDeltaJSON         `json:"-"`
}

// choiceDeltaJSON contains the JSON metadata for the struct [ChoiceDelta]
type choiceDeltaJSON struct {
	Content      apijson.Field
	FunctionCall apijson.Field
	Refusal      apijson.Field
	Role         apijson.Field
	ToolCalls    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChoiceDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceDeltaJSON) RawJSON() string {
	return r.raw
}

type ChoiceDeltaRole string

const (
	ChoiceDeltaRoleDeveloper ChoiceDeltaRole = "developer"
	ChoiceDeltaRoleSystem    ChoiceDeltaRole = "system"
	ChoiceDeltaRoleUser      ChoiceDeltaRole = "user"
	ChoiceDeltaRoleAssistant ChoiceDeltaRole = "assistant"
	ChoiceDeltaRoleTool      ChoiceDeltaRole = "tool"
)

func (r ChoiceDeltaRole) IsKnown() bool {
	switch r {
	case ChoiceDeltaRoleDeveloper, ChoiceDeltaRoleSystem, ChoiceDeltaRoleUser, ChoiceDeltaRoleAssistant, ChoiceDeltaRoleTool:
		return true
	}
	return false
}

type ChoiceDeltaFunctionCall struct {
	Arguments   string                      `json:"arguments,nullable"`
	Name        string                      `json:"name,nullable"`
	ExtraFields map[string]interface{}      `json:"-,extras"`
	JSON        choiceDeltaFunctionCallJSON `json:"-"`
}

// choiceDeltaFunctionCallJSON contains the JSON metadata for the struct
// [ChoiceDeltaFunctionCall]
type choiceDeltaFunctionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChoiceDeltaFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceDeltaFunctionCallJSON) RawJSON() string {
	return r.raw
}

type ChoiceDeltaToolCall struct {
	Index       int64                       `json:"index,required"`
	ID          string                      `json:"id,nullable"`
	Function    ChoiceDeltaToolCallFunction `json:"function,nullable"`
	Type        ChoiceDeltaToolCallType     `json:"type,nullable"`
	ExtraFields map[string]interface{}      `json:"-,extras"`
	JSON        choiceDeltaToolCallJSON     `json:"-"`
}

// choiceDeltaToolCallJSON contains the JSON metadata for the struct
// [ChoiceDeltaToolCall]
type choiceDeltaToolCallJSON struct {
	Index       apijson.Field
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChoiceDeltaToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceDeltaToolCallJSON) RawJSON() string {
	return r.raw
}

type ChoiceDeltaToolCallType string

const (
	ChoiceDeltaToolCallTypeFunction ChoiceDeltaToolCallType = "function"
)

func (r ChoiceDeltaToolCallType) IsKnown() bool {
	switch r {
	case ChoiceDeltaToolCallTypeFunction:
		return true
	}
	return false
}

type ChoiceDeltaToolCallFunction struct {
	Arguments   string                          `json:"arguments,nullable"`
	Name        string                          `json:"name,nullable"`
	ExtraFields map[string]interface{}          `json:"-,extras"`
	JSON        choiceDeltaToolCallFunctionJSON `json:"-"`
}

// choiceDeltaToolCallFunctionJSON contains the JSON metadata for the struct
// [ChoiceDeltaToolCallFunction]
type choiceDeltaToolCallFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChoiceDeltaToolCallFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceDeltaToolCallFunctionJSON) RawJSON() string {
	return r.raw
}

// Log probability information for the choice.
type ChoiceLogprobs struct {
	// A list of message content tokens with log probability information.
	Content []ChatCompletionTokenLogprob `json:"content,nullable"`
	// A list of message refusal tokens with log probability information.
	Refusal []ChatCompletionTokenLogprob `json:"refusal,nullable"`
	JSON    choiceLogprobsJSON           `json:"-"`
}

// choiceLogprobsJSON contains the JSON metadata for the struct [ChoiceLogprobs]
type choiceLogprobsJSON struct {
	Content     apijson.Field
	Refusal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChoiceLogprobs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r choiceLogprobsJSON) RawJSON() string {
	return r.raw
}

// A streaming chat completion choice chunk.
//
// OpenAI-compatible choice object for streaming responses. Part of the
// ChatCompletionChunk response in SSE streams.
type ChunkChoice struct {
	// Delta content for streaming responses
	Delta ChoiceDelta `json:"delta,required"`
	// The index of this choice in the list of choices
	Index int64 `json:"index,required"`
	// The reason the model stopped (only in final chunk)
	FinishReason ChunkChoiceFinishReason `json:"finish_reason,nullable"`
	// Log probability information for the choice.
	Logprobs ChoiceLogprobs  `json:"logprobs,nullable"`
	JSON     chunkChoiceJSON `json:"-"`
}

// chunkChoiceJSON contains the JSON metadata for the struct [ChunkChoice]
type chunkChoiceJSON struct {
	Delta        apijson.Field
	Index        apijson.Field
	FinishReason apijson.Field
	Logprobs     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ChunkChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chunkChoiceJSON) RawJSON() string {
	return r.raw
}

// The reason the model stopped (only in final chunk)
type ChunkChoiceFinishReason string

const (
	ChunkChoiceFinishReasonStop          ChunkChoiceFinishReason = "stop"
	ChunkChoiceFinishReasonLength        ChunkChoiceFinishReason = "length"
	ChunkChoiceFinishReasonToolCalls     ChunkChoiceFinishReason = "tool_calls"
	ChunkChoiceFinishReasonContentFilter ChunkChoiceFinishReason = "content_filter"
	ChunkChoiceFinishReasonFunctionCall  ChunkChoiceFinishReason = "function_call"
)

func (r ChunkChoiceFinishReason) IsKnown() bool {
	switch r {
	case ChunkChoiceFinishReasonStop, ChunkChoiceFinishReasonLength, ChunkChoiceFinishReasonToolCalls, ChunkChoiceFinishReasonContentFilter, ChunkChoiceFinishReasonFunctionCall:
		return true
	}
	return false
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
	Choices []Choice `json:"choices,required"`
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

// Breakdown of tokens used in a completion.
//
// Fields:
//
// - accepted_prediction_tokens (optional): int
// - audio_tokens (optional): int
// - reasoning_tokens (optional): int
// - rejected_prediction_tokens (optional): int
type CompletionTokensDetails struct {
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
	RejectedPredictionTokens int64                       `json:"rejected_prediction_tokens"`
	JSON                     completionTokensDetailsJSON `json:"-"`
}

// completionTokensDetailsJSON contains the JSON metadata for the struct
// [CompletionTokensDetails]
type completionTokensDetailsJSON struct {
	AcceptedPredictionTokens apijson.Field
	AudioTokens              apijson.Field
	ReasoningTokens          apijson.Field
	RejectedPredictionTokens apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CompletionTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionTokensDetailsJSON) RawJSON() string {
	return r.raw
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
type CompletionUsage struct {
	// Number of tokens in the generated completion.
	CompletionTokens int64 `json:"completion_tokens,required"`
	// Number of tokens in the prompt.
	PromptTokens int64 `json:"prompt_tokens,required"`
	// Total number of tokens used in the request (prompt + completion).
	TotalTokens int64 `json:"total_tokens,required"`
	// Breakdown of tokens used in a completion.
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
	// Breakdown of tokens used in the prompt.
	PromptTokensDetails PromptTokensDetails `json:"prompt_tokens_details"`
	JSON                completionUsageJSON `json:"-"`
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

// The custom tool that the model called.
//
// Fields:
//
// - name (required): str
// - input (required): str
type Custom struct {
	// The input for the custom tool call generated by the model.
	Input string `json:"input,required"`
	// The name of the custom tool to call.
	Name string     `json:"name,required"`
	JSON customJSON `json:"-"`
}

// customJSON contains the JSON metadata for the struct [Custom]
type customJSON struct {
	Input       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Custom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customJSON) RawJSON() string {
	return r.raw
}

// The function that the model called.
//
// Fields:
//
// - name (required): str
// - arguments (required): str
type Function struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string       `json:"name,required"`
	JSON functionJSON `json:"-"`
}

// functionJSON contains the JSON metadata for the struct [Function]
type functionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Function) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionJSON) RawJSON() string {
	return r.raw
}

// Deprecated and replaced by `tool_calls`. The name and arguments of a function
// that should be called, as generated by the model.
//
// Fields:
//
// - arguments (required): str
// - name (required): str
type FunctionCall struct {
	// The arguments to call the function with, as generated by the model in JSON
	// format. Note that the model does not always generate valid JSON, and may
	// hallucinate parameters not defined by your function schema. Validate the
	// arguments in your code before calling your function.
	Arguments string `json:"arguments,required"`
	// The name of the function to call.
	Name string           `json:"name,required"`
	JSON functionCallJSON `json:"-"`
}

// functionCallJSON contains the JSON metadata for the struct [FunctionCall]
type functionCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r functionCallJSON) RawJSON() string {
	return r.raw
}

// Details about the input tokens billed for this request.
//
// Fields:
//
// - text_tokens (optional): int
// - audio_tokens (optional): int
type InputTokenDetails struct {
	// Number of audio tokens billed for this request.
	AudioTokens int64 `json:"audio_tokens"`
	// Number of text tokens billed for this request.
	TextTokens int64                 `json:"text_tokens"`
	JSON       inputTokenDetailsJSON `json:"-"`
}

// inputTokenDetailsJSON contains the JSON metadata for the struct
// [InputTokenDetails]
type inputTokenDetailsJSON struct {
	AudioTokens apijson.Field
	TextTokens  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InputTokenDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inputTokenDetailsJSON) RawJSON() string {
	return r.raw
}

// Static predicted output content, such as the content of a text file that is
// being regenerated.
//
// Fields:
//
//   - type (required): Literal["content"]
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText], MinLen(1),
//     ArrayTitle("PredictionContentArray")]
type PredictionContentParam struct {
	// The content that should be matched when generating a model response. If
	// generated tokens would match this content, the entire model response can be
	// returned much more quickly.
	Content param.Field[PredictionContentContentUnionParam] `json:"content,required"`
	// The type of the predicted content you want to provide. This type is currently
	// always `content`.
	Type param.Field[PredictionContentType] `json:"type,required"`
}

func (r PredictionContentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The content that should be matched when generating a model response. If
// generated tokens would match this content, the entire model response can be
// returned much more quickly.
//
// Satisfied by [shared.UnionString],
// [PredictionContentContentPredictionContentArrayParam].
type PredictionContentContentUnionParam interface {
	ImplementsPredictionContentContentUnionParam()
}

type PredictionContentContentPredictionContentArrayParam []ChatCompletionContentPartTextParam

func (r PredictionContentContentPredictionContentArrayParam) ImplementsPredictionContentContentUnionParam() {
}

// The type of the predicted content you want to provide. This type is currently
// always `content`.
type PredictionContentType string

const (
	PredictionContentTypeContent PredictionContentType = "content"
)

func (r PredictionContentType) IsKnown() bool {
	switch r {
	case PredictionContentTypeContent:
		return true
	}
	return false
}

// Breakdown of tokens used in the prompt.
//
// Fields:
//
// - audio_tokens (optional): int
// - cached_tokens (optional): int
type PromptTokensDetails struct {
	// Audio input tokens present in the prompt.
	AudioTokens int64 `json:"audio_tokens"`
	// Cached tokens present in the prompt.
	CachedTokens int64                   `json:"cached_tokens"`
	JSON         promptTokensDetailsJSON `json:"-"`
}

// promptTokensDetailsJSON contains the JSON metadata for the struct
// [PromptTokensDetails]
type promptTokensDetailsJSON struct {
	AudioTokens  apijson.Field
	CachedTokens apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PromptTokensDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r promptTokensDetailsJSON) RawJSON() string {
	return r.raw
}

type ReasoningParam struct {
	Effort          param.Field[ReasoningEffort]          `json:"effort"`
	GenerateSummary param.Field[ReasoningGenerateSummary] `json:"generate_summary"`
	Summary         param.Field[ReasoningSummary]         `json:"summary"`
	ExtraFields     map[string]interface{}                `json:"-,extras"`
}

func (r ReasoningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReasoningEffort string

const (
	ReasoningEffortMinimal ReasoningEffort = "minimal"
	ReasoningEffortLow     ReasoningEffort = "low"
	ReasoningEffortMedium  ReasoningEffort = "medium"
	ReasoningEffortHigh    ReasoningEffort = "high"
)

func (r ReasoningEffort) IsKnown() bool {
	switch r {
	case ReasoningEffortMinimal, ReasoningEffortLow, ReasoningEffortMedium, ReasoningEffortHigh:
		return true
	}
	return false
}

type ReasoningGenerateSummary string

const (
	ReasoningGenerateSummaryAuto     ReasoningGenerateSummary = "auto"
	ReasoningGenerateSummaryConcise  ReasoningGenerateSummary = "concise"
	ReasoningGenerateSummaryDetailed ReasoningGenerateSummary = "detailed"
)

func (r ReasoningGenerateSummary) IsKnown() bool {
	switch r {
	case ReasoningGenerateSummaryAuto, ReasoningGenerateSummaryConcise, ReasoningGenerateSummaryDetailed:
		return true
	}
	return false
}

type ReasoningSummary string

const (
	ReasoningSummaryAuto     ReasoningSummary = "auto"
	ReasoningSummaryConcise  ReasoningSummary = "concise"
	ReasoningSummaryDetailed ReasoningSummary = "detailed"
)

func (r ReasoningSummary) IsKnown() bool {
	switch r {
	case ReasoningSummaryAuto, ReasoningSummaryConcise, ReasoningSummaryDetailed:
		return true
	}
	return false
}

// Server-Sent Event streaming format for chat completions
type StreamChunk struct {
	// Unique identifier for the chat completion
	ID string `json:"id,required"`
	// List of completion choice chunks
	Choices []ChunkChoice `json:"choices,required"`
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
	Usage CompletionUsage `json:"usage,nullable"`
	JSON  streamChunkJSON `json:"-"`
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

// Schema for ThinkingConfigDisabled.
//
// Fields:
//
// - type (required): Literal["disabled"]
type ThinkingConfigDisabledParam struct {
	Type param.Field[ThinkingConfigDisabledType] `json:"type,required"`
}

func (r ThinkingConfigDisabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThinkingConfigDisabledParam) implementsChatCompletionNewParamsThinkingUnion() {}

type ThinkingConfigDisabledType string

const (
	ThinkingConfigDisabledTypeDisabled ThinkingConfigDisabledType = "disabled"
)

func (r ThinkingConfigDisabledType) IsKnown() bool {
	switch r {
	case ThinkingConfigDisabledTypeDisabled:
		return true
	}
	return false
}

// Schema for ThinkingConfigEnabled.
//
// Fields:
//
// - budget_tokens (required): int
// - type (required): Literal["enabled"]
type ThinkingConfigEnabledParam struct {
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.claude.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64]                     `json:"budget_tokens,required"`
	Type         param.Field[ThinkingConfigEnabledType] `json:"type,required"`
}

func (r ThinkingConfigEnabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ThinkingConfigEnabledParam) implementsChatCompletionNewParamsThinkingUnion() {}

type ThinkingConfigEnabledType string

const (
	ThinkingConfigEnabledTypeEnabled ThinkingConfigEnabledType = "enabled"
)

func (r ThinkingConfigEnabledType) IsKnown() bool {
	switch r {
	case ThinkingConfigEnabledTypeEnabled:
		return true
	}
	return false
}

// Satisfied by [ToolChoiceString], [shared.UnionString], [ToolChoiceMapParam],
// [ToolChoiceMCPToolChoiceParam].
type ToolChoiceUnionParam interface {
	ImplementsToolChoiceUnionParam()
}

type ToolChoiceString string

const (
	ToolChoiceStringAuto     ToolChoiceString = "auto"
	ToolChoiceStringRequired ToolChoiceString = "required"
	ToolChoiceStringNone     ToolChoiceString = "none"
)

func (r ToolChoiceString) IsKnown() bool {
	switch r {
	case ToolChoiceStringAuto, ToolChoiceStringRequired, ToolChoiceStringNone:
		return true
	}
	return false
}

func (r ToolChoiceString) ImplementsToolChoiceUnionParam() {}

type ToolChoiceMapParam map[string]interface{}

func (r ToolChoiceMapParam) ImplementsToolChoiceUnionParam() {}

type ToolChoiceMCPToolChoiceParam struct {
	Name        param.Field[string] `json:"name,required"`
	ServerLabel param.Field[string] `json:"server_label,required"`
}

func (r ToolChoiceMCPToolChoiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceMCPToolChoiceParam) ImplementsToolChoiceUnionParam() {}

// The model will use any available tools.
//
// Fields:
//
// - disable_parallel_tool_use (optional): bool
// - type (required): Literal["any"]
type ToolChoiceAnyParam struct {
	Type param.Field[ToolChoiceAnyType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r ToolChoiceAnyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceAnyParam) implementsChatCompletionNewParamsToolChoiceUnion() {}

type ToolChoiceAnyType string

const (
	ToolChoiceAnyTypeAny ToolChoiceAnyType = "any"
)

func (r ToolChoiceAnyType) IsKnown() bool {
	switch r {
	case ToolChoiceAnyTypeAny:
		return true
	}
	return false
}

// The model will automatically decide whether to use tools.
//
// Fields:
//
// - disable_parallel_tool_use (optional): bool
// - type (required): Literal["auto"]
type ToolChoiceAutoParam struct {
	Type param.Field[ToolChoiceAutoType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r ToolChoiceAutoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceAutoParam) implementsChatCompletionNewParamsToolChoiceUnion() {}

type ToolChoiceAutoType string

const (
	ToolChoiceAutoTypeAuto ToolChoiceAutoType = "auto"
)

func (r ToolChoiceAutoType) IsKnown() bool {
	switch r {
	case ToolChoiceAutoTypeAuto:
		return true
	}
	return false
}

// The model will not be allowed to use tools.
//
// Fields:
//
// - type (required): Literal["none"]
type ToolChoiceNoneParam struct {
	Type param.Field[ToolChoiceNoneType] `json:"type,required"`
}

func (r ToolChoiceNoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceNoneParam) implementsChatCompletionNewParamsToolChoiceUnion() {}

type ToolChoiceNoneType string

const (
	ToolChoiceNoneTypeNone ToolChoiceNoneType = "none"
)

func (r ToolChoiceNoneType) IsKnown() bool {
	switch r {
	case ToolChoiceNoneTypeNone:
		return true
	}
	return false
}

// The model will use the specified tool with `tool_choice.name`.
//
// Fields:
//
// - disable_parallel_tool_use (optional): bool
// - name (required): str
// - type (required): Literal["tool"]
type ToolChoiceToolParam struct {
	// The name of the tool to use.
	Name param.Field[string]             `json:"name,required"`
	Type param.Field[ToolChoiceToolType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output exactly one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
}

func (r ToolChoiceToolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ToolChoiceToolParam) implementsChatCompletionNewParamsToolChoiceUnion() {}

type ToolChoiceToolType string

const (
	ToolChoiceToolTypeTool ToolChoiceToolType = "tool"
)

func (r ToolChoiceToolType) IsKnown() bool {
	switch r {
	case ToolChoiceToolTypeTool:
		return true
	}
	return false
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

// A URL citation when using web search.
//
// Fields:
//
// - end_index (required): int
// - start_index (required): int
// - url (required): str
// - title (required): str
type URLCitation struct {
	// The index of the last character of the URL citation in the message.
	EndIndex int64 `json:"end_index,required"`
	// The index of the first character of the URL citation in the message.
	StartIndex int64 `json:"start_index,required"`
	// The title of the web resource.
	Title string `json:"title,required"`
	// The URL of the web resource.
	URL  string          `json:"url,required"`
	JSON urlCitationJSON `json:"-"`
}

// urlCitationJSON contains the JSON metadata for the struct [URLCitation]
type urlCitationJSON struct {
	EndIndex    apijson.Field
	StartIndex  apijson.Field
	Title       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLCitation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlCitationJSON) RawJSON() string {
	return r.raw
}

type ChatCompletionNewParams struct {
	// Model identifier. Accepts model ID strings, lists for routing, or DedalusModel
	// objects with per-model settings.
	Model param.Field[ChatCompletionNewParamsModelUnion] `json:"model,required"`
	// Agent attributes. Values in [0.0, 1.0].
	AgentAttributes param.Field[map[string]float64] `json:"agent_attributes"`
	// Parameters for audio output. Required when audio output is requested with `mo...
	Audio param.Field[map[string]interface{}] `json:"audio"`
	// Execute tools server-side. If false, returns raw tool calls for manual handling.
	AutomaticToolExecution param.Field[bool] `json:"automatic_tool_execution"`
	// Optional. The name of the content [cached](https://ai.google.dev/gemini-api/d...
	CachedContent param.Field[string] `json:"cached_content"`
	// If set to `true`, the request returns a `request_id`. You can then get the de...
	Deferred param.Field[bool] `json:"deferred"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on the...
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Wrapper for union variant: function call mode.
	FunctionCall param.Field[string] `json:"function_call"`
	// Deprecated in favor of `tools`. A list of functions the model may generate J...
	Functions param.Field[[]ChatCompletionFunctionsParam] `json:"functions"`
	// Generation parameters wrapper (Google-specific)
	GenerationConfig param.Field[map[string]interface{}] `json:"generation_config"`
	// Content filtering and safety policy configuration.
	Guardrails param.Field[[]map[string]interface{}] `json:"guardrails"`
	// Configuration for multi-model handoffs.
	HandoffConfig param.Field[map[string]interface{}] `json:"handoff_config"`
	// Modify the likelihood of specified tokens appearing in the completion. Accep...
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// Whether to return log probabilities of the output tokens or not. If true, ret...
	Logprobs param.Field[bool] `json:"logprobs"`
	// Maximum tokens in completion (newer parameter name)
	MaxCompletionTokens param.Field[int64] `json:"max_completion_tokens"`
	// Maximum tokens in completion
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Maximum conversation turns.
	MaxTurns param.Field[int64] `json:"max_turns"`
	// MCP server identifiers. Accepts marketplace slugs, URLs, or MCPServerParam
	// objects. MCP tools are executed server-side and billed separately.
	MCPServers param.Field[ChatCompletionNewParamsMCPServersUnion] `json:"mcp_servers"`
	// Conversation history (OpenAI: messages, Google: contents, Responses: input)
	Messages param.Field[[]ChatCompletionNewParamsMessageUnion] `json:"messages"`
	// Set of 16 key-value pairs that can be attached to an object. This can be usef...
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// Output types that you would like the model to generate. Most models are capab...
	Modalities param.Field[[]string] `json:"modalities"`
	// Model attributes for routing. Maps model IDs to attribute dictionaries with
	// values in [0.0, 1.0].
	ModelAttributes param.Field[map[string]map[string]float64] `json:"model_attributes"`
	// How many chat completion choices to generate for each input message. Note tha...
	N param.Field[int64] `json:"n"`
	// Whether to enable parallel tool calls (Anthropic uses inverted polarity)
	ParallelToolCalls param.Field[bool] `json:"parallel_tool_calls"`
	// Static predicted output content, such as the content of a text file that is
	// being regenerated.
	//
	// Fields:
	//
	//   - type (required): Literal["content"]
	//   - content (required): str |
	//     Annotated[list[ChatCompletionRequestMessageContentPartText], MinLen(1),
	//     ArrayTitle("PredictionContentArray")]
	Prediction param.Field[PredictionContentParam] `json:"prediction"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whe...
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Used by OpenAI to cache responses for similar requests to optimize your cache...
	PromptCacheKey param.Field[string] `json:"prompt_cache_key"`
	// The retention policy for the prompt cache. Set to `24h` to enable extended pr...
	PromptCacheRetention param.Field[string] `json:"prompt_cache_retention"`
	// Allows toggling between the reasoning mode and no system prompt. When set to ...
	PromptMode param.Field[ChatCompletionNewParamsPromptMode] `json:"prompt_mode"`
	// Constrains effort on reasoning for [reasoning models](https://platform.openai...
	ReasoningEffort param.Field[string] `json:"reasoning_effort"`
	// An object specifying the format that the model must output. Setting to `{ "...
	ResponseFormat param.Field[ChatCompletionNewParamsResponseFormatUnion] `json:"response_format"`
	// Whether to inject a safety prompt before all conversations.
	SafePrompt param.Field[bool] `json:"safe_prompt"`
	// A stable identifier used to help detect users of your application that may be...
	SafetyIdentifier param.Field[string] `json:"safety_identifier"`
	// Safety/content filtering settings (Google-specific)
	SafetySettings param.Field[[]ChatCompletionNewParamsSafetySetting] `json:"safety_settings"`
	// Set the parameters to be used for searched data. If not set, no data will be ...
	SearchParameters param.Field[map[string]interface{}] `json:"search_parameters"`
	// Random seed for deterministic output
	Seed param.Field[int64] `json:"seed"`
	// Service tier for request processing
	ServiceTier param.Field[string] `json:"service_tier"`
	// Not supported with latest reasoning models `o3` and `o4-mini`. Up to 4 seque...
	Stop param.Field[ChatCompletionNewParamsStopUnion] `json:"stop"`
	// Custom text sequences that will cause the model to stop generating. Our mode...
	StopSequences param.Field[[]string] `json:"stop_sequences"`
	// Whether or not to store the output of this chat completion request for use in...
	Store param.Field[bool] `json:"store"`
	// Options for streaming response. Only set this when you set `stream: true`.
	StreamOptions param.Field[map[string]interface{}] `json:"stream_options"`
	// System instruction/prompt
	SystemInstruction param.Field[ChatCompletionNewParamsSystemInstructionUnion] `json:"system_instruction"`
	// Sampling temperature (0-2 for most providers)
	Temperature param.Field[float64] `json:"temperature"`
	// Extended thinking configuration (Anthropic-specific)
	Thinking param.Field[ChatCompletionNewParamsThinkingUnion] `json:"thinking"`
	// Controls which (if any) tool is called by the model. `none` means the model w...
	ToolChoice param.Field[ChatCompletionNewParamsToolChoiceUnion] `json:"tool_choice"`
	// Tool calling configuration (Google-specific)
	ToolConfig param.Field[map[string]interface{}] `json:"tool_config"`
	// Available tools/functions for the model
	Tools param.Field[[]ChatCompletionNewParamsToolUnion] `json:"tools"`
	// Top-k sampling parameter
	TopK param.Field[int64] `json:"top_k"`
	// An integer between 0 and 20 specifying the number of most likely tokens to re...
	TopLogprobs param.Field[int64] `json:"top_logprobs"`
	// Nucleus sampling threshold
	TopP param.Field[float64] `json:"top_p"`
	// This field is being replaced by `safety_identifier` and `prompt_cache_key`. U...
	User param.Field[string] `json:"user"`
	// Constrains the verbosity of the model's response. Lower values will result in...
	Verbosity param.Field[string] `json:"verbosity"`
	// This tool searches the web for relevant results to use in a response. Learn m...
	WebSearchOptions param.Field[map[string]interface{}] `json:"web_search_options"`
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Model identifier. Accepts model ID strings, lists for routing, or DedalusModel
// objects with per-model settings.
//
// Satisfied by [shared.UnionString], [shared.DedalusModelParam],
// [ChatCompletionNewParamsModelArray].
type ChatCompletionNewParamsModelUnion interface {
	ImplementsChatCompletionNewParamsModelUnion()
}

type ChatCompletionNewParamsModelArray []shared.DedalusModelChoiceUnionParam

func (r ChatCompletionNewParamsModelArray) ImplementsChatCompletionNewParamsModelUnion() {}

// MCP server identifiers. Accepts marketplace slugs, URLs, or MCPServerParam
// objects. MCP tools are executed server-side and billed separately.
//
// Satisfied by [shared.UnionString], [shared.MCPServerParam],
// [shared.MCPServersParam].
type ChatCompletionNewParamsMCPServersUnion interface {
	ImplementsChatCompletionNewParamsMCPServersUnion()
}

// Developer-provided instructions that the model should follow, regardless of
// messages sent by the user. With o1 models and newer, `developer` messages
// replace the previous `system` messages.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText], MinLen(1),
//     ArrayTitle("ChatCompletionRequestDeveloperMessageContentArray")]
//   - role (required): Literal["developer"]
//   - name (optional): str
type ChatCompletionNewParamsMessage struct {
	// The role of the messages author, in this case `developer`.
	Role param.Field[ChatCompletionNewParamsMessagesRole] `json:"role,required"`
	// Data about a previous audio response from the model.
	// [Learn more](https://platform.openai.com/docs/guides/audio).
	//
	// Fields:
	//
	// - id (required): str
	Audio        param.Field[ChatCompletionAudioParam] `json:"audio"`
	Content      param.Field[interface{}]              `json:"content"`
	FunctionCall param.Field[interface{}]              `json:"function_call"`
	// An optional name for the participant. Provides the model information to
	// differentiate between participants of the same role.
	Name param.Field[string] `json:"name"`
	// The refusal message by the assistant.
	Refusal param.Field[string] `json:"refusal"`
	// Tool call that this message is responding to.
	ToolCallID param.Field[string]      `json:"tool_call_id"`
	ToolCalls  param.Field[interface{}] `json:"tool_calls"`
}

func (r ChatCompletionNewParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsMessage) implementsChatCompletionNewParamsMessageUnion() {}

// Developer-provided instructions that the model should follow, regardless of
// messages sent by the user. With o1 models and newer, `developer` messages
// replace the previous `system` messages.
//
// Fields:
//
//   - content (required): str |
//     Annotated[list[ChatCompletionRequestMessageContentPartText], MinLen(1),
//     ArrayTitle("ChatCompletionRequestDeveloperMessageContentArray")]
//   - role (required): Literal["developer"]
//   - name (optional): str
//
// Satisfied by [ChatCompletionDeveloperMessageParam],
// [ChatCompletionSystemMessageParam], [ChatCompletionUserMessageParam],
// [ChatCompletionAssistantMessageParam], [ChatCompletionToolMessageParam],
// [ChatCompletionFunctionMessageParam], [ChatCompletionNewParamsMessage].
type ChatCompletionNewParamsMessageUnion interface {
	implementsChatCompletionNewParamsMessageUnion()
}

// The role of the messages author, in this case `developer`.
type ChatCompletionNewParamsMessagesRole string

const (
	ChatCompletionNewParamsMessagesRoleDeveloper ChatCompletionNewParamsMessagesRole = "developer"
	ChatCompletionNewParamsMessagesRoleSystem    ChatCompletionNewParamsMessagesRole = "system"
	ChatCompletionNewParamsMessagesRoleUser      ChatCompletionNewParamsMessagesRole = "user"
	ChatCompletionNewParamsMessagesRoleAssistant ChatCompletionNewParamsMessagesRole = "assistant"
	ChatCompletionNewParamsMessagesRoleTool      ChatCompletionNewParamsMessagesRole = "tool"
	ChatCompletionNewParamsMessagesRoleFunction  ChatCompletionNewParamsMessagesRole = "function"
)

func (r ChatCompletionNewParamsMessagesRole) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsMessagesRoleDeveloper, ChatCompletionNewParamsMessagesRoleSystem, ChatCompletionNewParamsMessagesRoleUser, ChatCompletionNewParamsMessagesRoleAssistant, ChatCompletionNewParamsMessagesRoleTool, ChatCompletionNewParamsMessagesRoleFunction:
		return true
	}
	return false
}

// Allows toggling between the reasoning mode and no system prompt. When set to ...
type ChatCompletionNewParamsPromptMode string

const (
	ChatCompletionNewParamsPromptModeReasoning ChatCompletionNewParamsPromptMode = "reasoning"
)

func (r ChatCompletionNewParamsPromptMode) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsPromptModeReasoning:
		return true
	}
	return false
}

// An object specifying the format that the model must output. Setting to `{ "...
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

// An object specifying the format that the model must output. Setting to `{ "...
//
// Satisfied by [shared.ResponseFormatTextParam],
// [shared.ResponseFormatJSONSchemaParam], [shared.ResponseFormatJSONObjectParam],
// [ChatCompletionNewParamsResponseFormat].
type ChatCompletionNewParamsResponseFormatUnion interface {
	ImplementsChatCompletionNewParamsResponseFormatUnion()
}

// The type of response format being defined. Always `text`.
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
//   - category (required): HarmCategory
//   - threshold (required): Literal["HARM_BLOCK_THRESHOLD_UNSPECIFIED",
//     "BLOCK_LOW_AND_ABOVE", "BLOCK_MEDIUM_AND_ABOVE", "BLOCK_ONLY_HIGH",
//     "BLOCK_NONE", "OFF"]
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

// Not supported with latest reasoning models `o3` and `o4-mini`. Up to 4 seque...
//
// Satisfied by [ChatCompletionNewParamsStopArray], [shared.UnionString].
type ChatCompletionNewParamsStopUnion interface {
	ImplementsChatCompletionNewParamsStopUnion()
}

type ChatCompletionNewParamsStopArray []string

func (r ChatCompletionNewParamsStopArray) ImplementsChatCompletionNewParamsStopUnion() {}

// System instruction/prompt
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
	Type param.Field[ChatCompletionNewParamsThinkingType] `json:"type,required"`
	// Determines how many tokens Claude can use for its internal reasoning process.
	// Larger budgets can enable more thorough analysis for complex problems, improving
	// response quality.
	//
	// Must be ≥1024 and less than `max_tokens`.
	//
	// See
	// [extended thinking](https://docs.claude.com/en/docs/build-with-claude/extended-thinking)
	// for details.
	BudgetTokens param.Field[int64] `json:"budget_tokens"`
}

func (r ChatCompletionNewParamsThinking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsThinking) implementsChatCompletionNewParamsThinkingUnion() {}

// Extended thinking configuration (Anthropic-specific)
//
// Satisfied by [ThinkingConfigEnabledParam], [ThinkingConfigDisabledParam],
// [ChatCompletionNewParamsThinking].
type ChatCompletionNewParamsThinkingUnion interface {
	implementsChatCompletionNewParamsThinkingUnion()
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

// Controls which (if any) tool is called by the model. `none` means the model w...
type ChatCompletionNewParamsToolChoice struct {
	Type param.Field[ChatCompletionNewParamsToolChoiceType] `json:"type,required"`
	// Whether to disable parallel tool use.
	//
	// Defaults to `false`. If set to `true`, the model will output at most one tool
	// use.
	DisableParallelToolUse param.Field[bool] `json:"disable_parallel_tool_use"`
	// The name of the tool to use.
	Name param.Field[string] `json:"name"`
}

func (r ChatCompletionNewParamsToolChoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolChoice) implementsChatCompletionNewParamsToolChoiceUnion() {}

// Controls which (if any) tool is called by the model. `none` means the model w...
//
// Satisfied by [ToolChoiceAutoParam], [ToolChoiceAnyParam], [ToolChoiceToolParam],
// [ToolChoiceNoneParam], [ChatCompletionNewParamsToolChoice].
type ChatCompletionNewParamsToolChoiceUnion interface {
	implementsChatCompletionNewParamsToolChoiceUnion()
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
// - type (required): Literal["function"]
// - function (required): FunctionObject
type ChatCompletionNewParamsTool struct {
	// The type of the tool. Currently, only `function` is supported.
	Type   param.Field[ChatCompletionNewParamsToolsType] `json:"type,required"`
	Custom param.Field[interface{}]                      `json:"custom"`
	// Schema for FunctionObject.
	//
	// Fields:
	//
	// - description (optional): str
	// - name (required): str
	// - parameters (optional): FunctionParameters
	// - strict (optional): bool | None
	Function param.Field[shared.FunctionDefinitionParam] `json:"function"`
}

func (r ChatCompletionNewParamsTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsTool) implementsChatCompletionNewParamsToolUnion() {}

// A function tool that can be used to generate a response.
//
// Fields:
//
// - type (required): Literal["function"]
// - function (required): FunctionObject
//
// Satisfied by [ChatCompletionToolParam],
// [ChatCompletionNewParamsToolsCustomToolChatCompletions],
// [ChatCompletionNewParamsTool].
type ChatCompletionNewParamsToolUnion interface {
	implementsChatCompletionNewParamsToolUnion()
}

// A custom tool that processes input using a specified format.
//
// Fields:
//
// - type (required): Literal["custom"]
// - custom (required): CustomToolProperties
type ChatCompletionNewParamsToolsCustomToolChatCompletions struct {
	// Properties of the custom tool.
	Custom param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustom] `json:"custom,required"`
	// The type of the custom tool. Always `custom`.
	Type param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsType] `json:"type,required"`
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletions) implementsChatCompletionNewParamsToolUnion() {
}

// Properties of the custom tool.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustom struct {
	// The name of the custom tool, used to identify it in tool calls.
	Name param.Field[string] `json:"name,required"`
	// Optional description of the custom tool, used to provide more context.
	Description param.Field[string] `json:"description"`
	// The input format for the custom tool. Default is unconstrained text.
	Format param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatUnion] `json:"format"`
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The input format for the custom tool. Default is unconstrained text.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormat struct {
	// Unconstrained text format. Always `text`.
	Type    param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatType] `json:"type,required"`
	Grammar param.Field[interface{}]                                                           `json:"grammar"`
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormat) implementsChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatUnion() {
}

// The input format for the custom tool. Default is unconstrained text.
//
// Satisfied by
// [ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormat],
// [ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormat],
// [ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormat].
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatUnion interface {
	implementsChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatUnion()
}

// Unconstrained free-form text.
//
// Fields:
//
// - type (required): Literal["text"]
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormat struct {
	// Unconstrained text format. Always `text`.
	Type param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormatType] `json:"type,required"`
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormat) implementsChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatUnion() {
}

// Unconstrained text format. Always `text`.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormatType string

const (
	ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormatTypeText ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormatType = "text"
)

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormatType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTextFormatTypeText:
		return true
	}
	return false
}

// A grammar defined by the user.
//
// Fields:
//
// - type (required): Literal["grammar"]
// - grammar (required): GrammarFormatGrammarFormat
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormat struct {
	// Your chosen grammar.
	Grammar param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammar] `json:"grammar,required"`
	// Grammar format. Always `grammar`.
	Type param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatType] `json:"type,required"`
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormat) implementsChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatUnion() {
}

// Your chosen grammar.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammar struct {
	// The grammar definition.
	Definition param.Field[string] `json:"definition,required"`
	// The syntax of the grammar definition. One of `lark` or `regex`.
	Syntax param.Field[ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntax] `json:"syntax,required"`
}

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The syntax of the grammar definition. One of `lark` or `regex`.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntax string

const (
	ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntaxLark  ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntax = "lark"
	ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntaxRegex ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntax = "regex"
)

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntax) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntaxLark, ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatGrammarSyntaxRegex:
		return true
	}
	return false
}

// Grammar format. Always `grammar`.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatType string

const (
	ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatTypeGrammar ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatType = "grammar"
)

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatGrammarFormatTypeGrammar:
		return true
	}
	return false
}

// Unconstrained text format. Always `text`.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatType string

const (
	ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTypeText    ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatType = "text"
	ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTypeGrammar ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatType = "grammar"
)

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTypeText, ChatCompletionNewParamsToolsCustomToolChatCompletionsCustomFormatTypeGrammar:
		return true
	}
	return false
}

// The type of the custom tool. Always `custom`.
type ChatCompletionNewParamsToolsCustomToolChatCompletionsType string

const (
	ChatCompletionNewParamsToolsCustomToolChatCompletionsTypeCustom ChatCompletionNewParamsToolsCustomToolChatCompletionsType = "custom"
)

func (r ChatCompletionNewParamsToolsCustomToolChatCompletionsType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsCustomToolChatCompletionsTypeCustom:
		return true
	}
	return false
}

// The type of the tool. Currently, only `function` is supported.
type ChatCompletionNewParamsToolsType string

const (
	ChatCompletionNewParamsToolsTypeFunction ChatCompletionNewParamsToolsType = "function"
	ChatCompletionNewParamsToolsTypeCustom   ChatCompletionNewParamsToolsType = "custom"
)

func (r ChatCompletionNewParamsToolsType) IsKnown() bool {
	switch r {
	case ChatCompletionNewParamsToolsTypeFunction, ChatCompletionNewParamsToolsTypeCustom:
		return true
	}
	return false
}
