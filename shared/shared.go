// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
//
// The property Model is required.
type DedalusModelParam struct {
	// Model identifier with provider prefix (e.g., 'openai/gpt-5',
	// 'anthropic/claude-3-5-sonnet').
	Model string `json:"model,required"`
	// Optional default generation settings (e.g., temperature, max_tokens) applied
	// when this model is selected.
	Settings DedalusModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r DedalusModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
type DedalusModelSettingsParam struct {
	FrequencyPenalty                param.Opt[float64]                 `json:"frequency_penalty,omitzero"`
	IncludeUsage                    param.Opt[bool]                    `json:"include_usage,omitzero"`
	InputAudioFormat                param.Opt[string]                  `json:"input_audio_format,omitzero"`
	Logprobs                        param.Opt[bool]                    `json:"logprobs,omitzero"`
	MaxCompletionTokens             param.Opt[int64]                   `json:"max_completion_tokens,omitzero"`
	MaxTokens                       param.Opt[int64]                   `json:"max_tokens,omitzero"`
	N                               param.Opt[int64]                   `json:"n,omitzero"`
	OutputAudioFormat               param.Opt[string]                  `json:"output_audio_format,omitzero"`
	ParallelToolCalls               param.Opt[bool]                    `json:"parallel_tool_calls,omitzero"`
	PresencePenalty                 param.Opt[float64]                 `json:"presence_penalty,omitzero"`
	PromptCacheKey                  param.Opt[string]                  `json:"prompt_cache_key,omitzero"`
	ReasoningEffort                 param.Opt[string]                  `json:"reasoning_effort,omitzero"`
	SafetyIdentifier                param.Opt[string]                  `json:"safety_identifier,omitzero"`
	Seed                            param.Opt[int64]                   `json:"seed,omitzero"`
	ServiceTier                     param.Opt[string]                  `json:"service_tier,omitzero"`
	Store                           param.Opt[bool]                    `json:"store,omitzero"`
	Stream                          param.Opt[bool]                    `json:"stream,omitzero"`
	Temperature                     param.Opt[float64]                 `json:"temperature,omitzero"`
	Timeout                         param.Opt[float64]                 `json:"timeout,omitzero"`
	TopK                            param.Opt[int64]                   `json:"top_k,omitzero"`
	TopLogprobs                     param.Opt[int64]                   `json:"top_logprobs,omitzero"`
	TopP                            param.Opt[float64]                 `json:"top_p,omitzero"`
	User                            param.Opt[string]                  `json:"user,omitzero"`
	Verbosity                       param.Opt[string]                  `json:"verbosity,omitzero"`
	Voice                           param.Opt[string]                  `json:"voice,omitzero"`
	DisableAutomaticFunctionCalling param.Opt[bool]                    `json:"disable_automatic_function_calling,omitzero"`
	UseResponses                    param.Opt[bool]                    `json:"use_responses,omitzero"`
	Audio                           map[string]any                     `json:"audio,omitzero"`
	ExtraArgs                       map[string]any                     `json:"extra_args,omitzero"`
	ExtraHeaders                    map[string]string                  `json:"extra_headers,omitzero"`
	ExtraQuery                      map[string]any                     `json:"extra_query,omitzero"`
	GenerationConfig                map[string]any                     `json:"generation_config,omitzero"`
	InputAudioTranscription         map[string]any                     `json:"input_audio_transcription,omitzero"`
	LogitBias                       map[string]int64                   `json:"logit_bias,omitzero"`
	Metadata                        map[string]string                  `json:"metadata,omitzero"`
	Modalities                      []string                           `json:"modalities,omitzero"`
	Prediction                      map[string]any                     `json:"prediction,omitzero"`
	Reasoning                       DedalusModelSettingsReasoningParam `json:"reasoning,omitzero"`
	ResponseFormat                  map[string]any                     `json:"response_format,omitzero"`
	// Any of "code_interpreter_call.outputs", "computer_call_output.output.image_url",
	// "file_search_call.results", "message.input_image.image_url",
	// "message.output_text.logprobs", "reasoning.encrypted_content".
	ResponseInclude   []string                                 `json:"response_include,omitzero"`
	SafetySettings    []map[string]any                         `json:"safety_settings,omitzero"`
	Stop              DedalusModelSettingsStopUnionParam       `json:"stop,omitzero"`
	StreamOptions     map[string]any                           `json:"stream_options,omitzero"`
	SystemInstruction map[string]any                           `json:"system_instruction,omitzero"`
	Thinking          map[string]any                           `json:"thinking,omitzero"`
	ToolChoice        DedalusModelSettingsToolChoiceUnionParam `json:"tool_choice,omitzero"`
	ToolConfig        map[string]any                           `json:"tool_config,omitzero"`
	// Any of "auto", "disabled".
	Truncation       string         `json:"truncation,omitzero"`
	TurnDetection    map[string]any `json:"turn_detection,omitzero"`
	WebSearchOptions map[string]any `json:"web_search_options,omitzero"`
	Attributes       map[string]any `json:"attributes,omitzero"`
	StructuredOutput any            `json:"structured_output,omitzero"`
	paramObj
}

func (r DedalusModelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DedalusModelSettingsParam](
		"truncation", "auto", "disabled",
	)
}

type DedalusModelSettingsReasoningParam struct {
	// Any of "minimal", "low", "medium", "high".
	Effort string `json:"effort,omitzero"`
	// Any of "auto", "concise", "detailed".
	GenerateSummary string `json:"generate_summary,omitzero"`
	// Any of "auto", "concise", "detailed".
	Summary     string         `json:"summary,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r DedalusModelSettingsReasoningParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelSettingsReasoningParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DedalusModelSettingsReasoningParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DedalusModelSettingsReasoningParam](
		"effort", "minimal", "low", "medium", "high",
	)
	apijson.RegisterFieldValidator[DedalusModelSettingsReasoningParam](
		"generate_summary", "auto", "concise", "detailed",
	)
	apijson.RegisterFieldValidator[DedalusModelSettingsReasoningParam](
		"summary", "auto", "concise", "detailed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DedalusModelSettingsStopUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelSettingsStopUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *DedalusModelSettingsStopUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelSettingsStopUnionParam) asAny() any {
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
type DedalusModelSettingsToolChoiceUnionParam struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfDedalusModelSettingsToolChoiceString)
	OfDedalusModelSettingsToolChoiceString param.Opt[string]                                 `json:",omitzero,inline"`
	OfString                               param.Opt[string]                                 `json:",omitzero,inline"`
	OfMCPToolChoice                        *DedalusModelSettingsToolChoiceMCPToolChoiceParam `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelSettingsToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDedalusModelSettingsToolChoiceString, u.OfString, u.OfMCPToolChoice)
}
func (u *DedalusModelSettingsToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelSettingsToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDedalusModelSettingsToolChoiceString) {
		return &u.OfDedalusModelSettingsToolChoiceString
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMCPToolChoice) {
		return u.OfMCPToolChoice
	}
	return nil
}

type DedalusModelSettingsToolChoiceString string

const (
	DedalusModelSettingsToolChoiceStringAuto     DedalusModelSettingsToolChoiceString = "auto"
	DedalusModelSettingsToolChoiceStringRequired DedalusModelSettingsToolChoiceString = "required"
	DedalusModelSettingsToolChoiceStringNone     DedalusModelSettingsToolChoiceString = "none"
)

// The properties Name, ServerLabel are required.
type DedalusModelSettingsToolChoiceMCPToolChoiceParam struct {
	Name        string `json:"name,required"`
	ServerLabel string `json:"server_label,required"`
	paramObj
}

func (r DedalusModelSettingsToolChoiceMCPToolChoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow DedalusModelSettingsToolChoiceMCPToolChoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DedalusModelSettingsToolChoiceMCPToolChoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified model metadata across all providers.
//
// Combines provider-specific schemas into a single, consistent format. Fields that
// aren't available from a provider are set to None.
type Model struct {
	// Unique model identifier with provider prefix (e.g., 'openai/gpt-4')
	ID string `json:"id,required"`
	// When the model was released (RFC 3339)
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Provider that hosts this model
	//
	// Any of "openai", "anthropic", "google", "xai", "mistral", "groq".
	Provider ModelProvider `json:"provider,required"`
	// Normalized model capabilities across all providers.
	Capabilities ModelCapabilities `json:"capabilities,nullable"`
	// Provider-declared default parameters for model generation.
	Defaults ModelDefaults `json:"defaults,nullable"`
	// Model description
	Description string `json:"description,nullable"`
	// Human-readable model name
	DisplayName string `json:"display_name,nullable"`
	// Provider-specific generation method names (None = not declared)
	ProviderDeclaredGenerationMethods []string `json:"provider_declared_generation_methods,nullable"`
	// Raw provider-specific metadata
	ProviderInfo map[string]any `json:"provider_info,nullable"`
	// Model version identifier
	Version string `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                respjson.Field
		CreatedAt                         respjson.Field
		Provider                          respjson.Field
		Capabilities                      respjson.Field
		Defaults                          respjson.Field
		Description                       respjson.Field
		DisplayName                       respjson.Field
		ProviderDeclaredGenerationMethods respjson.Field
		ProviderInfo                      respjson.Field
		Version                           respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider that hosts this model
type ModelProvider string

const (
	ModelProviderOpenAI    ModelProvider = "openai"
	ModelProviderAnthropic ModelProvider = "anthropic"
	ModelProviderGoogle    ModelProvider = "google"
	ModelProviderXai       ModelProvider = "xai"
	ModelProviderMistral   ModelProvider = "mistral"
	ModelProviderGroq      ModelProvider = "groq"
)

// Normalized model capabilities across all providers.
type ModelCapabilities struct {
	// Supports audio processing
	Audio bool `json:"audio,nullable"`
	// Supports image generation
	ImageGeneration bool `json:"image_generation,nullable"`
	// Maximum input tokens
	InputTokenLimit int64 `json:"input_token_limit,nullable"`
	// Maximum output tokens
	OutputTokenLimit int64 `json:"output_token_limit,nullable"`
	// Supports streaming responses
	Streaming bool `json:"streaming,nullable"`
	// Supports structured JSON output
	StructuredOutput bool `json:"structured_output,nullable"`
	// Supports text generation
	Text bool `json:"text,nullable"`
	// Supports extended thinking/reasoning
	Thinking bool `json:"thinking,nullable"`
	// Supports function/tool calling
	Tools bool `json:"tools,nullable"`
	// Supports image understanding
	Vision bool `json:"vision,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio            respjson.Field
		ImageGeneration  respjson.Field
		InputTokenLimit  respjson.Field
		OutputTokenLimit respjson.Field
		Streaming        respjson.Field
		StructuredOutput respjson.Field
		Text             respjson.Field
		Thinking         respjson.Field
		Tools            respjson.Field
		Vision           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelCapabilities) RawJSON() string { return r.JSON.raw }
func (r *ModelCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider-declared default parameters for model generation.
type ModelDefaults struct {
	// Default maximum output tokens
	MaxOutputTokens int64 `json:"max_output_tokens,nullable"`
	// Default temperature setting
	Temperature float64 `json:"temperature,nullable"`
	// Default top_k setting
	TopK int64 `json:"top_k,nullable"`
	// Default top_p setting
	TopP float64 `json:"top_p,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxOutputTokens respjson.Field
		Temperature     respjson.Field
		TopK            respjson.Field
		TopP            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelDefaults) RawJSON() string { return r.JSON.raw }
func (r *ModelDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
