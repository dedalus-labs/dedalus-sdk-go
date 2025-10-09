// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"context"
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
// This endpoint provides a vendor-agnostic chat completions API that works with
// thousands of LLMs. It supports MCP integration, multi-model routing with
// intelligent agentic handoffs, client-side and server-side tool execution, and
// streaming and non-streaming responses.
//
// Args: request: Chat completion request with messages, model, and configuration.
// http_request: FastAPI request object for accessing headers and state.
// background_tasks: FastAPI background tasks for async billing operations. user:
// Authenticated user with validated API key and sufficient balance.
//
// Returns: ChatCompletion: OpenAI-compatible completion response with usage data.
//
// Raises: HTTPException: - 401 if authentication fails or insufficient balance. -
// 400 if request validation fails. - 500 if internal processing error occurs.
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
func (r *ChatCompletionService) New(ctx context.Context, body ChatCompletionNewParams, opts ...option.RequestOption) (res *StreamChunk, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a chat completion.
//
// This endpoint provides a vendor-agnostic chat completions API that works with
// thousands of LLMs. It supports MCP integration, multi-model routing with
// intelligent agentic handoffs, client-side and server-side tool execution, and
// streaming and non-streaming responses.
//
// Args: request: Chat completion request with messages, model, and configuration.
// http_request: FastAPI request object for accessing headers and state.
// background_tasks: FastAPI background tasks for async billing operations. user:
// Authenticated user with validated API key and sufficient balance.
//
// Returns: ChatCompletion: OpenAI-compatible completion response with usage data.
//
// Raises: HTTPException: - 401 if authentication fails or insufficient balance. -
// 400 if request validation fails. - 500 if internal processing error occurs.
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
	// Log probability of this token.
	Logprob float64 `json:"logprob,required"`
	// List of the most likely tokens and their log probability information.
	TopLogprobs []TopLogprob `json:"top_logprobs,required"`
	// Bytes representation of the token.
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

func DedalusModelChoiceParamOfDedalusModel(model string) DedalusModelChoiceUnionParam {
	var variant DedalusModelParam
	variant.Model = model
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
	// The token.
	Token string `json:"token,required"`
	// Log probability of this token.
	Logprob float64 `json:"logprob,required"`
	// Bytes representation of the token.
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
	// A list of messages comprising the conversation so far. Depending on the model
	// you use, different message types (modalities) are supported, like text, images,
	// and audio.
	Messages []map[string]any `json:"messages,omitzero,required"`
	// Model(s) to use for completion. Can be a single model ID, a DedalusModel object,
	// or a list for multi-model routing. Single model: 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet-20241022', 'openai/gpt-4o-mini', or a DedalusModel
	// instance. Multi-model routing: ['openai/gpt-4o-mini', 'openai/gpt-4',
	// 'anthropic/claude-3-5-sonnet'] or list of DedalusModel objects - agent will
	// choose optimal model based on task complexity.
	Model ChatCompletionNewParamsModelUnion `json:"model,omitzero,required"`
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
	// 'json_object'} for the legacy JSON mode.
	ResponseFormat map[string]any `json:"response_format,omitzero"`
	// Google safety settings (harm categories and thresholds).
	SafetySettings []map[string]any `json:"safety_settings,omitzero"`
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
type ChatCompletionNewParamsModelUnion struct {
	OfModelID      param.Opt[ModelID] `json:",omitzero,inline"`
	OfDedalusModel *DedalusModelParam `json:",omitzero,inline"`
	OfModels       ModelsParam        `json:",omitzero,inline"`
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
type ChatCompletionNewParamsMCPServersUnion struct {
	OfMCPServers    []ChatCompletionNewParamsMCPServersMCPServerUnion `json:",omitzero,inline"`
	OfMCPServerSlug param.Opt[string]                                 `json:",omitzero,inline"`
	OfMCPServerSpec *ChatCompletionNewParamsMCPServersMCPServerSpec   `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMCPServersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMCPServers, u.OfMCPServerSlug, u.OfMCPServerSpec)
}
func (u *ChatCompletionNewParamsMCPServersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMCPServersUnion) asAny() any {
	if !param.IsOmitted(u.OfMCPServers) {
		return &u.OfMCPServers
	} else if !param.IsOmitted(u.OfMCPServerSlug) {
		return &u.OfMCPServerSlug.Value
	} else if !param.IsOmitted(u.OfMCPServerSpec) {
		return u.OfMCPServerSpec
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionNewParamsMCPServersMCPServerUnion struct {
	OfMCPServerSlug param.Opt[string]                                        `json:",omitzero,inline"`
	OfMCPServerSpec *ChatCompletionNewParamsMCPServersMCPServerMCPServerSpec `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionNewParamsMCPServersMCPServerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMCPServerSlug, u.OfMCPServerSpec)
}
func (u *ChatCompletionNewParamsMCPServersMCPServerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionNewParamsMCPServersMCPServerUnion) asAny() any {
	if !param.IsOmitted(u.OfMCPServerSlug) {
		return &u.OfMCPServerSlug.Value
	} else if !param.IsOmitted(u.OfMCPServerSpec) {
		return u.OfMCPServerSpec
	}
	return nil
}

// Structured representation of an MCP server reference.
type ChatCompletionNewParamsMCPServersMCPServerMCPServerSpec struct {
	// Slug identifying an MCP server (e.g., 'dedalus-labs/brave-search').
	Slug param.Opt[string] `json:"slug,omitzero"`
	// Explicit MCP server URL.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Optional explicit version to target when using a slug.
	Version param.Opt[string] `json:"version,omitzero"`
	// Optional metadata associated with the MCP server entry.
	Metadata    map[string]any `json:"metadata,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ChatCompletionNewParamsMCPServersMCPServerMCPServerSpec) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMCPServersMCPServerMCPServerSpec
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ChatCompletionNewParamsMCPServersMCPServerMCPServerSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured representation of an MCP server reference.
type ChatCompletionNewParamsMCPServersMCPServerSpec struct {
	// Slug identifying an MCP server (e.g., 'dedalus-labs/brave-search').
	Slug param.Opt[string] `json:"slug,omitzero"`
	// Explicit MCP server URL.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Optional explicit version to target when using a slug.
	Version param.Opt[string] `json:"version,omitzero"`
	// Optional metadata associated with the MCP server entry.
	Metadata    map[string]any `json:"metadata,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ChatCompletionNewParamsMCPServersMCPServerSpec) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionNewParamsMCPServersMCPServerSpec
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ChatCompletionNewParamsMCPServersMCPServerSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
