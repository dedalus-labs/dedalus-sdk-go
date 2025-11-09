// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
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
	Deferred                        param.Opt[bool]                    `json:"deferred,omitzero"`
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
	// Any of "file_search_call.results", "web_search_call.results",
	// "web_search_call.action.sources", "message.input_image.image_url",
	// "computer_call_output.output.image_url", "code_interpreter_call.outputs",
	// "reasoning.encrypted_content", "message.output_text.logprobs".
	ResponseInclude   []string                                 `json:"response_include,omitzero"`
	SafetySettings    []map[string]any                         `json:"safety_settings,omitzero"`
	SearchParameters  map[string]any                           `json:"search_parameters,omitzero"`
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
	OfAnyMap                               map[string]any                                    `json:",omitzero,inline"`
	OfMCPToolChoice                        *DedalusModelSettingsToolChoiceMCPToolChoiceParam `json:",omitzero,inline"`
	paramUnion
}

func (u DedalusModelSettingsToolChoiceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDedalusModelSettingsToolChoiceString, u.OfString, u.OfAnyMap, u.OfMCPToolChoice)
}
func (u *DedalusModelSettingsToolChoiceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DedalusModelSettingsToolChoiceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDedalusModelSettingsToolChoiceString) {
		return &u.OfDedalusModelSettingsToolChoiceString
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
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

func DedalusModelChoiceParamOfDedalusModel(model string) DedalusModelChoiceUnionParam {
	var variant DedalusModelParam
	variant.Model = model
	return DedalusModelChoiceUnionParam{OfDedalusModel: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DedalusModelChoiceUnionParam struct {
	OfModelID      param.Opt[githubcomdedaluslabsdedalussdkgo.ModelID] `json:",omitzero,inline"`
	OfDedalusModel *DedalusModelParam                                  `json:",omitzero,inline"`
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
