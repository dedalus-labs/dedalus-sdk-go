// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	shimjson "github.com/dedalus-labs/dedalus-sdk-go/internal/encoding/json"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/ssestream"
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

// Create a chat completion using the Agent framework.
//
// This endpoint provides a vendor-agnostic chat completion API that works with
// 100+ LLM providers via the Agent framework. It supports both single and
// multi-model routing, client-side and server-side tool execution, and integration
// with MCP (Model Context Protocol) servers.
//
// Features: - Cross-vendor compatibility (OpenAI, Anthropic, Cohere, etc.) -
// Multi-model routing with intelligent agentic handoffs - Client-side tool
// execution (tools returned as JSON) - Server-side MCP tool execution with
// automatic billing - Streaming and non-streaming responses - Advanced agent
// attributes for routing decisions - Automatic usage tracking and billing
//
// Args: request: Chat completion request with messages, model, and configuration
// http_request: FastAPI request object for accessing headers and state
// background_tasks: FastAPI background tasks for async billing operations user:
// Authenticated user with validated API key and sufficient balance
//
// Returns: ChatCompletion: OpenAI-compatible completion response with usage data
//
// Raises: HTTPException: - 401 if authentication fails or insufficient balance -
// 400 if request validation fails - 500 if internal processing error occurs
//
// Billing: - Token usage billed automatically based on model pricing - MCP tool
// calls billed separately using credits system - Streaming responses billed after
// completion via background task
//
// Example: Basic chat completion: ```python from dedalus_labs import Dedalus
//
//	client = Dedalus(api_key="your-api-key")
//
//	completion = client.chat.completions.create(
//	    model="openai/gpt-5",
//	    messages=[{"role": "user", "content": "Hello, how are you?"}],
//	)
//
//	print(completion.choices[0].message.content)
//	```
//
//	With tools and MCP servers:
//	```python
//	completion = client.chat.completions.create(
//	    model="openai/gpt-5",
//	    messages=[{"role": "user", "content": "Search for recent AI news"}],
//	    tools=[
//	        {
//	            "type": "function",
//	            "function": {
//	                "name": "search_web",
//	                "description": "Search the web for information",
//	            },
//	        }
//	    ],
//	    mcp_servers=["dedalus-labs/brave-search"],
//	)
//	```
//
//	Multi-model routing:
//	```python
//	completion = client.chat.completions.create(
//	    model=[
//	        "openai/gpt-4o-mini",
//	        "openai/gpt-5",
//	        "anthropic/claude-sonnet-4-20250514",
//	    ],
//	    messages=[{"role": "user", "content": "Analyze this complex data"}],
//	    agent_attributes={"complexity": 0.8, "accuracy": 0.9},
//	)
//	```
//
//	Streaming response:
//	```python
//	stream = client.chat.completions.create(
//	    model="openai/gpt-5",
//	    messages=[{"role": "user", "content": "Tell me a story"}],
//	    stream=True,
//	)
//
//	for chunk in stream:
//	    if chunk.choices[0].delta.content:
//	        print(chunk.choices[0].delta.content, end="")
//	```
func (r *ChatCompletionService) NewStreaming(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[StreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[StreamChunk](ssestream.NewDecoder(raw), err)
}

// Token log probability information.
type ChatCompletionTokenLogprob struct {
	// The token
	Token string `json:"token,required"`
	// Log probability of this token
	Logprob float64 `json:"logprob,required"`
	// List of most likely tokens and their log probabilities
	TopLogprobs []TopLogprob `json:"top_logprobs,required"`
	// Bytes representation of the token
	Bytes []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionTokenLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionTokenLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat completion request with 'messages' field (OpenAI standard).
//
// The property Messages is required.
type CompletionRequestMessagesParam struct {
	// Messages to the model. Supports role/content structure and multimodal content
	// arrays.
	Messages []map[string]any `json:"messages,omitzero,required"`
	// Frequency penalty (-2 to 2). Positive values penalize new tokens based on their
	// existing frequency in the text so far, decreasing likelihood of repeated
	// phrases.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// Maximum number of tokens to generate in the completion. Does not include tokens
	// in the input messages.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Maximum number of turns for agent execution before terminating (default: 10).
	// Each turn represents one model inference cycle. Higher values allow more complex
	// reasoning but increase cost and latency.
	MaxTurns param.Opt[int64] `json:"max_turns,omitzero"`
	// Number of completions to generate. Note: only n=1 is currently supported.
	N param.Opt[int64] `json:"n,omitzero"`
	// Presence penalty (-2 to 2). Positive values penalize new tokens based on whether
	// they appear in the text so far, encouraging the model to talk about new topics.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// Sampling temperature (0 to 2). Higher values make output more random, lower
	// values make it more focused and deterministic. 0 = deterministic, 1 = balanced,
	// 2 = very creative.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Nucleus sampling parameter (0 to 1). Alternative to temperature. 0.1 = only top
	// 10% probability mass, 1.0 = consider all tokens.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Unique identifier representing your end-user. Used for monitoring and abuse
	// detection. Should be consistent across requests from the same user.
	User param.Opt[string] `json:"user,omitzero"`
	// Whether to stream back partial message deltas as Server-Sent Events. When true,
	// partial message deltas will be sent as OpenAI-compatible chunks.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Attributes for the agent itself, influencing behavior and model selection.
	// Format: {'attribute': value}, where values are 0.0-1.0. Common attributes:
	// 'complexity', 'accuracy', 'efficiency', 'creativity', 'friendliness'. Higher
	// values indicate stronger preference for that characteristic.
	AgentAttributes map[string]float64 `json:"agent_attributes,omitzero"`
	// Guardrails to apply to the agent for input/output validation and safety checks.
	// Reserved for future use - guardrails configuration format not yet finalized.
	Guardrails []map[string]any `json:"guardrails,omitzero"`
	// Configuration for multi-model handoffs and agent orchestration. Reserved for
	// future use - handoff configuration format not yet finalized.
	HandoffConfig map[string]any `json:"handoff_config,omitzero"`
	// Modify likelihood of specified tokens appearing in the completion. Maps token
	// IDs (as strings) to bias values (-100 to 100). -100 = completely ban token, +100
	// = strongly favor token.
	LogitBias map[string]int64 `json:"logit_bias,omitzero"`
	// MCP (Model Context Protocol) server addresses to make available for server-side
	// tool execution. Can be URLs (e.g., 'https://mcp.example.com') or slugs (e.g.,
	// 'dedalus-labs/brave-search'). MCP tools are executed server-side and billed
	// separately.
	MCPServers []string `json:"mcp_servers,omitzero"`
	// Model(s) to use for completion. Can be a single model ID, a DedalusModel object,
	// or a list for multi-model routing. Single model: 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet-20241022', 'openai/gpt-4o-mini', or a DedalusModel
	// instance. Multi-model routing: ['openai/gpt-4o-mini', 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet'] or list of DedalusModel objects - agent will
	// choose optimal model based on task complexity.
	Model CompletionRequestMessagesModelUnionParam `json:"model,omitzero"`
	// Attributes for individual models used in routing decisions during multi-model
	// execution. Format: {'model_name': {'attribute': value}}, where values are
	// 0.0-1.0. Common attributes: 'intelligence', 'speed', 'cost', 'creativity',
	// 'accuracy'. Used by agent to select optimal model based on task requirements.
	ModelAttributes map[string]map[string]float64 `json:"model_attributes,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The model
	// will stop as soon as it encounters any of these sequences.
	Stop []string `json:"stop,omitzero"`
	// Controls which tool is called by the model. Options: 'auto' (default), 'none',
	// 'required', or specific tool name. Can also be a dict specifying a particular
	// tool.
	ToolChoice CompletionRequestMessagesToolChoiceUnionParam `json:"tool_choice,omitzero"`
	// list of tools available to the model in OpenAI function calling format. Tools
	// are executed client-side and returned as JSON for the application to handle. Use
	// 'mcp_servers' for server-side tool execution.
	Tools []map[string]any `json:"tools,omitzero"`
	paramObj
}

func (r CompletionRequestMessagesParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionRequestMessagesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionRequestMessagesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CompletionRequestMessagesModelUnionParam struct {
	OfModelID      param.Opt[ModelID] `json:",omitzero,inline"`
	OfDedalusModel *DedalusModelParam `json:",omitzero,inline"`
	OfModels       ModelsParam        `json:",omitzero,inline"`
	paramUnion
}

func (u CompletionRequestMessagesModelUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModelID, u.OfDedalusModel, u.OfModels)
}
func (u *CompletionRequestMessagesModelUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CompletionRequestMessagesModelUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModelID) {
		return &u.OfModelID.Value
	} else if !param.IsOmitted(u.OfDedalusModel) {
		return u.OfDedalusModel
	} else if !param.IsOmitted(u.OfModels) {
		return &u.OfModels
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CompletionRequestMessagesToolChoiceUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u CompletionRequestMessagesToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap)
}
func (u *CompletionRequestMessagesToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CompletionRequestMessagesToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

func DedalusModelChoiceParamOfDedalusModel(name string) DedalusModelChoiceUnionParam {
	var variant DedalusModelParam
	variant.Name = name
	return DedalusModelChoiceUnionParam{OfDedalusModel: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DedalusModelChoiceUnionParam struct {
	OfModelID      param.Opt[ModelID] `json:",omitzero,inline"`
	OfDedalusModel *DedalusModelParam `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModelID, u.OfDedalusModel)
}
func (u *DedalusModelChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfModelID) {
		return &u.OfModelID.Value
	} else if !param.IsOmitted(u.OfDedalusModel) {
		return u.OfDedalusModel
	}
	return nil
}

type ModelID = string

type ModelsParam []DedalusModelChoiceUnionParam

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
	// Usage statistics (only in final chunk with stream_options.include_usage=true)
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
	// Log probabilities for the content tokens
	Content []ChatCompletionTokenLogprob `json:"content,nullable"`
	// Log probabilities for refusal tokens, if any
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

// Usage statistics (only in final chunk with stream_options.include_usage=true)
type StreamChunkUsage struct {
	CompletionTokens        int64                                   `json:"completion_tokens,required"`
	PromptTokens            int64                                   `json:"prompt_tokens,required"`
	TotalTokens             int64                                   `json:"total_tokens,required"`
	CompletionTokensDetails StreamChunkUsageCompletionTokensDetails `json:"completion_tokens_details,nullable"`
	PromptTokensDetails     StreamChunkUsagePromptTokensDetails     `json:"prompt_tokens_details,nullable"`
	ExtraFields             map[string]any                          `json:",extras"`
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

type StreamChunkUsageCompletionTokensDetails struct {
	AcceptedPredictionTokens int64          `json:"accepted_prediction_tokens,nullable"`
	AudioTokens              int64          `json:"audio_tokens,nullable"`
	ReasoningTokens          int64          `json:"reasoning_tokens,nullable"`
	RejectedPredictionTokens int64          `json:"rejected_prediction_tokens,nullable"`
	ExtraFields              map[string]any `json:",extras"`
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

type StreamChunkUsagePromptTokensDetails struct {
	AudioTokens  int64          `json:"audio_tokens,nullable"`
	CachedTokens int64          `json:"cached_tokens,nullable"`
	ExtraFields  map[string]any `json:",extras"`
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
	// The token
	Token string `json:"token,required"`
	// Log probability of this token
	Logprob float64 `json:"logprob,required"`
	// Bytes representation of the token
	Bytes []int64 `json:"bytes,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
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
	// Chat completion request with 'messages' field (OpenAI standard).
	CompletionRequestMessages CompletionRequestMessagesParam
	paramObj
}

func (r ChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CompletionRequestMessages)
}
func (r *ChatCompletionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CompletionRequestMessages)
}
