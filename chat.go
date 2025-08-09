// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/param"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/ssestream"
)

// ChatService contains methods and other services that help with interacting with
// the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r *ChatService) {
	r = &ChatService{}
	r.Options = opts
	return
}

// Create a chat completion using the Agent framework.
//
// This endpoint provides a vendor-agnostic chat completion API that works with
// 100+ LLM providers via the Agent framework. It supports both single and
// multi-model routing, client-side and server-side tool execution, and integration
// with MCP (Model Context Protocol) servers.
func (r *ChatService) NewStreaming(ctx context.Context, body ChatNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[StreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	path := "v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[StreamChunk](ssestream.NewDecoder(raw), err)
}

type Completion struct {
	ID                string                `json:"id,required"`
	Choices           []CompletionChoice    `json:"choices,required"`
	Created           int64                 `json:"created,required"`
	Model             string                `json:"model,required"`
	Object            CompletionObject      `json:"object,required"`
	ServiceTier       CompletionServiceTier `json:"service_tier,nullable"`
	SystemFingerprint string                `json:"system_fingerprint,nullable"`
	Usage             CompletionUsage       `json:"usage,nullable"`
	JSON              completionJSON        `json:"-"`
}

// completionJSON contains the JSON metadata for the struct [Completion]
type completionJSON struct {
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

func (r *Completion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionJSON) RawJSON() string {
	return r.raw
}

type CompletionChoice struct {
	FinishReason CompletionChoicesFinishReason `json:"finish_reason,required"`
	Index        int64                         `json:"index,required"`
	Message      CompletionChoicesMessage      `json:"message,required"`
	Logprobs     CompletionChoicesLogprobs     `json:"logprobs,nullable"`
	JSON         completionChoiceJSON          `json:"-"`
}

// completionChoiceJSON contains the JSON metadata for the struct
// [CompletionChoice]
type completionChoiceJSON struct {
	FinishReason apijson.Field
	Index        apijson.Field
	Message      apijson.Field
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

type CompletionChoicesMessage struct {
	Role         CompletionChoicesMessageRole         `json:"role,required"`
	Annotations  []CompletionChoicesMessageAnnotation `json:"annotations,nullable"`
	Audio        CompletionChoicesMessageAudio        `json:"audio,nullable"`
	Content      string                               `json:"content,nullable"`
	FunctionCall CompletionChoicesMessageFunctionCall `json:"function_call,nullable"`
	Refusal      string                               `json:"refusal,nullable"`
	ToolCalls    []CompletionChoicesMessageToolCall   `json:"tool_calls,nullable"`
	JSON         completionChoicesMessageJSON         `json:"-"`
}

// completionChoicesMessageJSON contains the JSON metadata for the struct
// [CompletionChoicesMessage]
type completionChoicesMessageJSON struct {
	Role         apijson.Field
	Annotations  apijson.Field
	Audio        apijson.Field
	Content      apijson.Field
	FunctionCall apijson.Field
	Refusal      apijson.Field
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

type CompletionChoicesMessageAnnotation struct {
	Type        CompletionChoicesMessageAnnotationsType        `json:"type,required"`
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

type CompletionChoicesMessageAnnotationsURLCitation struct {
	EndIndex   int64                                              `json:"end_index,required"`
	StartIndex int64                                              `json:"start_index,required"`
	Title      string                                             `json:"title,required"`
	URL        string                                             `json:"url,required"`
	JSON       completionChoicesMessageAnnotationsURLCitationJSON `json:"-"`
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

type CompletionChoicesMessageAudio struct {
	ID         string                            `json:"id,required"`
	Data       string                            `json:"data,required"`
	ExpiresAt  int64                             `json:"expires_at,required"`
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

type CompletionChoicesMessageFunctionCall struct {
	Arguments string                                   `json:"arguments,required"`
	Name      string                                   `json:"name,required"`
	JSON      completionChoicesMessageFunctionCallJSON `json:"-"`
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

type CompletionChoicesMessageToolCall struct {
	ID       string                                    `json:"id,required"`
	Function CompletionChoicesMessageToolCallsFunction `json:"function,required"`
	Type     CompletionChoicesMessageToolCallsType     `json:"type,required"`
	JSON     completionChoicesMessageToolCallJSON      `json:"-"`
}

// completionChoicesMessageToolCallJSON contains the JSON metadata for the struct
// [CompletionChoicesMessageToolCall]
type completionChoicesMessageToolCallJSON struct {
	ID          apijson.Field
	Function    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageToolCallsFunction struct {
	Arguments string                                        `json:"arguments,required"`
	Name      string                                        `json:"name,required"`
	JSON      completionChoicesMessageToolCallsFunctionJSON `json:"-"`
}

// completionChoicesMessageToolCallsFunctionJSON contains the JSON metadata for the
// struct [CompletionChoicesMessageToolCallsFunction]
type completionChoicesMessageToolCallsFunctionJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesMessageToolCallsFunction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesMessageToolCallsFunctionJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesMessageToolCallsType string

const (
	CompletionChoicesMessageToolCallsTypeFunction CompletionChoicesMessageToolCallsType = "function"
)

func (r CompletionChoicesMessageToolCallsType) IsKnown() bool {
	switch r {
	case CompletionChoicesMessageToolCallsTypeFunction:
		return true
	}
	return false
}

type CompletionChoicesLogprobs struct {
	Content []CompletionChoicesLogprobsContent `json:"content,nullable"`
	Refusal []CompletionChoicesLogprobsRefusal `json:"refusal,nullable"`
	JSON    completionChoicesLogprobsJSON      `json:"-"`
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

type CompletionChoicesLogprobsContent struct {
	Token       string                                       `json:"token,required"`
	Logprob     float64                                      `json:"logprob,required"`
	TopLogprobs []CompletionChoicesLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                      `json:"bytes,nullable"`
	JSON        completionChoicesLogprobsContentJSON         `json:"-"`
}

// completionChoicesLogprobsContentJSON contains the JSON metadata for the struct
// [CompletionChoicesLogprobsContent]
type completionChoicesLogprobsContentJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsContentJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsContentTopLogprob struct {
	Token   string                                         `json:"token,required"`
	Logprob float64                                        `json:"logprob,required"`
	Bytes   []int64                                        `json:"bytes,nullable"`
	JSON    completionChoicesLogprobsContentTopLogprobJSON `json:"-"`
}

// completionChoicesLogprobsContentTopLogprobJSON contains the JSON metadata for
// the struct [CompletionChoicesLogprobsContentTopLogprob]
type completionChoicesLogprobsContentTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsContentTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsContentTopLogprobJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsRefusal struct {
	Token       string                                       `json:"token,required"`
	Logprob     float64                                      `json:"logprob,required"`
	TopLogprobs []CompletionChoicesLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                      `json:"bytes,nullable"`
	JSON        completionChoicesLogprobsRefusalJSON         `json:"-"`
}

// completionChoicesLogprobsRefusalJSON contains the JSON metadata for the struct
// [CompletionChoicesLogprobsRefusal]
type completionChoicesLogprobsRefusalJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsRefusal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsRefusalJSON) RawJSON() string {
	return r.raw
}

type CompletionChoicesLogprobsRefusalTopLogprob struct {
	Token   string                                         `json:"token,required"`
	Logprob float64                                        `json:"logprob,required"`
	Bytes   []int64                                        `json:"bytes,nullable"`
	JSON    completionChoicesLogprobsRefusalTopLogprobJSON `json:"-"`
}

// completionChoicesLogprobsRefusalTopLogprobJSON contains the JSON metadata for
// the struct [CompletionChoicesLogprobsRefusalTopLogprob]
type completionChoicesLogprobsRefusalTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompletionChoicesLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r completionChoicesLogprobsRefusalTopLogprobJSON) RawJSON() string {
	return r.raw
}

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

type CompletionUsage struct {
	CompletionTokens        int64                                  `json:"completion_tokens,required"`
	PromptTokens            int64                                  `json:"prompt_tokens,required"`
	TotalTokens             int64                                  `json:"total_tokens,required"`
	CompletionTokensDetails CompletionUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	PromptTokensDetails     CompletionUsagePromptTokensDetails     `json:"prompt_tokens_details,nullable"`
	JSON                    completionUsageJSON                    `json:"-"`
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

type CompletionUsageCompletionTokensDetails struct {
	AcceptedPredictionTokens int64                                      `json:"accepted_prediction_tokens,nullable"`
	AudioTokens              int64                                      `json:"audio_tokens,nullable"`
	ReasoningTokens          int64                                      `json:"reasoning_tokens,nullable"`
	RejectedPredictionTokens int64                                      `json:"rejected_prediction_tokens,nullable"`
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

type CompletionUsagePromptTokensDetails struct {
	AudioTokens  int64                                  `json:"audio_tokens,nullable"`
	CachedTokens int64                                  `json:"cached_tokens,nullable"`
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

// Request model for chat completions.
type CompletionRequestParam struct {
	// Attributes for the agent itself, influencing behavior and model selection.
	AgentAttributes param.Field[map[string]float64] `json:"agent_attributes"`
	// Frequency penalty (-2 to 2). Positive values penalize new tokens based on their
	// existing frequency in the text so far.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Input to the model - can be messages, images, or other modalities. Supports
	// OpenAI chat format with role/content structure. For multimodal inputs, content
	// can include text, images, or other media types.
	Input param.Field[[]interface{}] `json:"input"`
	// Modify likelihood of specified tokens appearing in the completion.
	LogitBias param.Field[map[string]int64] `json:"logit_bias"`
	// Maximum number of tokens to generate in the completion.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Maximum number of turns for agent execution before terminating (default: 10).
	MaxTurns param.Field[int64] `json:"max_turns"`
	// MCP (Model Context Protocol) server addresses to make available for server-side
	// tool execution.
	McpServers param.Field[[]string] `json:"mcp_servers"`
	// Model(s) to use for completion. Can be a single model ID, a Model object, or a
	// list for multi-model routing.
	Model param.Field[CompletionRequestModelUnionParam] `json:"model"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution.
	ModelAttributes param.Field[map[string]map[string]float64] `json:"model_attributes"`
	// Number of completions to generate. Note: only n=1 is currently supported.
	N param.Field[int64] `json:"n"`
	// Presence penalty (-2 to 2). Positive values penalize new tokens based on whether
	// they appear in the text so far.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Up to 4 sequences where the API will stop generating further tokens.
	Stop param.Field[[]string] `json:"stop"`
	// Whether to stream back partial message deltas as Server-Sent Events.
	Stream param.Field[bool] `json:"stream"`
	// Sampling temperature (0 to 2). Higher values make output more random, lower
	// values make it more focused and deterministic.
	Temperature param.Field[float64] `json:"temperature"`
	// Controls which tool is called by the model.
	ToolChoice param.Field[interface{}] `json:"tool_choice"`
	// List of tools available to the model in OpenAI function calling format.
	Tools param.Field[[]interface{}] `json:"tools"`
	// Nucleus sampling parameter (0 to 1). Alternative to temperature.
	TopP param.Field[float64] `json:"top_p"`
	// Unique identifier representing your end-user.
	User param.Field[string] `json:"user"`
}

func (r CompletionRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Model(s) to use for completion. Can be a single model ID, a Model object, or a
// list for multi-model routing.
//
// Satisfied by [shared.UnionString], [CompletionRequestModelArrayParam],
// [CompletionRequestModelModelParam], [CompletionRequestModelArrayParam].
type CompletionRequestModelUnionParam interface {
	ImplementsCompletionRequestModelUnionParam()
}

type CompletionRequestModelArrayParam []string

func (r CompletionRequestModelArrayParam) ImplementsCompletionRequestModelUnionParam() {}

type CompletionRequestModelModelParam struct {
	// Model identifier (e.g., 'gpt-4', 'claude-3-5-sonnet')
	Name param.Field[string] `json:"name,required"`
	// Model attributes as scores between 0-1. Used for multi-model routing decisions.
	Attributes param.Field[map[string]float64] `json:"attributes"`
	// Model generation settings including temperature, max_tokens, and other
	// parameters.
	Settings param.Field[CompletionRequestModelModelSettingsParam] `json:"settings"`
}

func (r CompletionRequestModelModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CompletionRequestModelModelParam) ImplementsCompletionRequestModelUnionParam() {}

// Model generation settings including temperature, max_tokens, and other
// parameters.
type CompletionRequestModelModelSettingsParam struct {
	FrequencyPenalty  param.Field[float64]           `json:"frequency_penalty"`
	IncludeUsage      param.Field[bool]              `json:"include_usage"`
	InputAudioFormat  param.Field[string]            `json:"input_audio_format"`
	MaxTokens         param.Field[int64]             `json:"max_tokens"`
	Metadata          param.Field[map[string]string] `json:"metadata"`
	Modalities        param.Field[[]string]          `json:"modalities"`
	OutputAudioFormat param.Field[string]            `json:"output_audio_format"`
	ParallelToolCalls param.Field[bool]              `json:"parallel_tool_calls"`
	PresencePenalty   param.Field[float64]           `json:"presence_penalty"`
	Store             param.Field[bool]              `json:"store"`
	Temperature       param.Field[float64]           `json:"temperature"`
	TopP              param.Field[float64]           `json:"top_p"`
	Voice             param.Field[string]            `json:"voice"`
}

func (r CompletionRequestModelModelSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamChunk struct {
	ID                string                 `json:"id,required"`
	Choices           []StreamChunkChoice    `json:"choices,required"`
	Created           int64                  `json:"created,required"`
	Model             string                 `json:"model,required"`
	Object            StreamChunkObject      `json:"object,required"`
	ServiceTier       StreamChunkServiceTier `json:"service_tier,nullable"`
	SystemFingerprint string                 `json:"system_fingerprint,nullable"`
	Usage             StreamChunkUsage       `json:"usage,nullable"`
	JSON              streamChunkJSON        `json:"-"`
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

type StreamChunkChoice struct {
	Delta        StreamChunkChoicesDelta        `json:"delta,required"`
	Index        int64                          `json:"index,required"`
	FinishReason StreamChunkChoicesFinishReason `json:"finish_reason,nullable"`
	Logprobs     StreamChunkChoicesLogprobs     `json:"logprobs,nullable"`
	JSON         streamChunkChoiceJSON          `json:"-"`
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

type StreamChunkChoicesDelta struct {
	Content      string                              `json:"content,nullable"`
	FunctionCall StreamChunkChoicesDeltaFunctionCall `json:"function_call,nullable"`
	Refusal      string                              `json:"refusal,nullable"`
	Role         StreamChunkChoicesDeltaRole         `json:"role,nullable"`
	ToolCalls    []StreamChunkChoicesDeltaToolCall   `json:"tool_calls,nullable"`
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
	Arguments string                                  `json:"arguments,nullable"`
	Name      string                                  `json:"name,nullable"`
	JSON      streamChunkChoicesDeltaFunctionCallJSON `json:"-"`
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
	Index    int64                                    `json:"index,required"`
	ID       string                                   `json:"id,nullable"`
	Function StreamChunkChoicesDeltaToolCallsFunction `json:"function,nullable"`
	Type     StreamChunkChoicesDeltaToolCallsType     `json:"type,nullable"`
	JSON     streamChunkChoicesDeltaToolCallJSON      `json:"-"`
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
	Arguments string                                       `json:"arguments,nullable"`
	Name      string                                       `json:"name,nullable"`
	JSON      streamChunkChoicesDeltaToolCallsFunctionJSON `json:"-"`
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

type StreamChunkChoicesLogprobs struct {
	Content []StreamChunkChoicesLogprobsContent `json:"content,nullable"`
	Refusal []StreamChunkChoicesLogprobsRefusal `json:"refusal,nullable"`
	JSON    streamChunkChoicesLogprobsJSON      `json:"-"`
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

type StreamChunkChoicesLogprobsContent struct {
	Token       string                                        `json:"token,required"`
	Logprob     float64                                       `json:"logprob,required"`
	TopLogprobs []StreamChunkChoicesLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                       `json:"bytes,nullable"`
	JSON        streamChunkChoicesLogprobsContentJSON         `json:"-"`
}

// streamChunkChoicesLogprobsContentJSON contains the JSON metadata for the struct
// [StreamChunkChoicesLogprobsContent]
type streamChunkChoicesLogprobsContentJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsContentJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsContentTopLogprob struct {
	Token   string                                          `json:"token,required"`
	Logprob float64                                         `json:"logprob,required"`
	Bytes   []int64                                         `json:"bytes,nullable"`
	JSON    streamChunkChoicesLogprobsContentTopLogprobJSON `json:"-"`
}

// streamChunkChoicesLogprobsContentTopLogprobJSON contains the JSON metadata for
// the struct [StreamChunkChoicesLogprobsContentTopLogprob]
type streamChunkChoicesLogprobsContentTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsContentTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsContentTopLogprobJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsRefusal struct {
	Token       string                                        `json:"token,required"`
	Logprob     float64                                       `json:"logprob,required"`
	TopLogprobs []StreamChunkChoicesLogprobsRefusalTopLogprob `json:"top_logprobs,required"`
	Bytes       []int64                                       `json:"bytes,nullable"`
	JSON        streamChunkChoicesLogprobsRefusalJSON         `json:"-"`
}

// streamChunkChoicesLogprobsRefusalJSON contains the JSON metadata for the struct
// [StreamChunkChoicesLogprobsRefusal]
type streamChunkChoicesLogprobsRefusalJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	TopLogprobs apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsRefusal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsRefusalJSON) RawJSON() string {
	return r.raw
}

type StreamChunkChoicesLogprobsRefusalTopLogprob struct {
	Token   string                                          `json:"token,required"`
	Logprob float64                                         `json:"logprob,required"`
	Bytes   []int64                                         `json:"bytes,nullable"`
	JSON    streamChunkChoicesLogprobsRefusalTopLogprobJSON `json:"-"`
}

// streamChunkChoicesLogprobsRefusalTopLogprobJSON contains the JSON metadata for
// the struct [StreamChunkChoicesLogprobsRefusalTopLogprob]
type streamChunkChoicesLogprobsRefusalTopLogprobJSON struct {
	Token       apijson.Field
	Logprob     apijson.Field
	Bytes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamChunkChoicesLogprobsRefusalTopLogprob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamChunkChoicesLogprobsRefusalTopLogprobJSON) RawJSON() string {
	return r.raw
}

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

type StreamChunkUsage struct {
	CompletionTokens        int64                                   `json:"completion_tokens,required"`
	PromptTokens            int64                                   `json:"prompt_tokens,required"`
	TotalTokens             int64                                   `json:"total_tokens,required"`
	CompletionTokensDetails StreamChunkUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	PromptTokensDetails     StreamChunkUsagePromptTokensDetails     `json:"prompt_tokens_details,nullable"`
	JSON                    streamChunkUsageJSON                    `json:"-"`
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

type StreamChunkUsageCompletionTokensDetails struct {
	AcceptedPredictionTokens int64                                       `json:"accepted_prediction_tokens,nullable"`
	AudioTokens              int64                                       `json:"audio_tokens,nullable"`
	ReasoningTokens          int64                                       `json:"reasoning_tokens,nullable"`
	RejectedPredictionTokens int64                                       `json:"rejected_prediction_tokens,nullable"`
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

type StreamChunkUsagePromptTokensDetails struct {
	AudioTokens  int64                                   `json:"audio_tokens,nullable"`
	CachedTokens int64                                   `json:"cached_tokens,nullable"`
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

type ChatNewParams struct {
	// Request model for chat completions.
	CompletionRequest CompletionRequestParam `json:"completion_request,required"`
}

func (r ChatNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CompletionRequest)
}
