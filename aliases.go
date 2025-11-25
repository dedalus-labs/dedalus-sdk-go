// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apierror"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
)

type Error = apierror.Error

// Structured model selection entry used in request payloads.
//
// Supports OpenAI-style semantics (string model id) while enabling optional
// per-model default settings for Dedalus multi-model routing.
//
// This is an alias to an internal type.
type DedalusModelParam = shared.DedalusModelParam

// Optional default generation settings (e.g., temperature, max_tokens) applied
// when this model is selected.
//
// This is an alias to an internal type.
type DedalusModelSettingsParam = shared.DedalusModelSettingsParam

// This is an alias to an internal type.
type DedalusModelSettingsReasoningParam = shared.DedalusModelSettingsReasoningParam

// This is an alias to an internal type.
type DedalusModelSettingsReasoningEffort = shared.DedalusModelSettingsReasoningEffort

// This is an alias to an internal value.
const DedalusModelSettingsReasoningEffortMinimal = shared.DedalusModelSettingsReasoningEffortMinimal

// This is an alias to an internal value.
const DedalusModelSettingsReasoningEffortLow = shared.DedalusModelSettingsReasoningEffortLow

// This is an alias to an internal value.
const DedalusModelSettingsReasoningEffortMedium = shared.DedalusModelSettingsReasoningEffortMedium

// This is an alias to an internal value.
const DedalusModelSettingsReasoningEffortHigh = shared.DedalusModelSettingsReasoningEffortHigh

// This is an alias to an internal type.
type DedalusModelSettingsReasoningGenerateSummary = shared.DedalusModelSettingsReasoningGenerateSummary

// This is an alias to an internal value.
const DedalusModelSettingsReasoningGenerateSummaryAuto = shared.DedalusModelSettingsReasoningGenerateSummaryAuto

// This is an alias to an internal value.
const DedalusModelSettingsReasoningGenerateSummaryConcise = shared.DedalusModelSettingsReasoningGenerateSummaryConcise

// This is an alias to an internal value.
const DedalusModelSettingsReasoningGenerateSummaryDetailed = shared.DedalusModelSettingsReasoningGenerateSummaryDetailed

// This is an alias to an internal type.
type DedalusModelSettingsReasoningSummary = shared.DedalusModelSettingsReasoningSummary

// This is an alias to an internal value.
const DedalusModelSettingsReasoningSummaryAuto = shared.DedalusModelSettingsReasoningSummaryAuto

// This is an alias to an internal value.
const DedalusModelSettingsReasoningSummaryConcise = shared.DedalusModelSettingsReasoningSummaryConcise

// This is an alias to an internal value.
const DedalusModelSettingsReasoningSummaryDetailed = shared.DedalusModelSettingsReasoningSummaryDetailed

// This is an alias to an internal type.
type DedalusModelSettingsResponseInclude = shared.DedalusModelSettingsResponseInclude

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeFileSearchCallResults = shared.DedalusModelSettingsResponseIncludeFileSearchCallResults

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeWebSearchCallResults = shared.DedalusModelSettingsResponseIncludeWebSearchCallResults

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeWebSearchCallActionSources = shared.DedalusModelSettingsResponseIncludeWebSearchCallActionSources

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeMessageInputImageImageURL = shared.DedalusModelSettingsResponseIncludeMessageInputImageImageURL

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeComputerCallOutputOutputImageURL = shared.DedalusModelSettingsResponseIncludeComputerCallOutputOutputImageURL

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeCodeInterpreterCallOutputs = shared.DedalusModelSettingsResponseIncludeCodeInterpreterCallOutputs

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeReasoningEncryptedContent = shared.DedalusModelSettingsResponseIncludeReasoningEncryptedContent

// This is an alias to an internal value.
const DedalusModelSettingsResponseIncludeMessageOutputTextLogprobs = shared.DedalusModelSettingsResponseIncludeMessageOutputTextLogprobs

// This is an alias to an internal type.
type DedalusModelSettingsStopUnionParam = shared.DedalusModelSettingsStopUnionParam

// This is an alias to an internal type.
type DedalusModelSettingsStopArrayParam = shared.DedalusModelSettingsStopArrayParam

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceUnionParam = shared.DedalusModelSettingsToolChoiceUnionParam

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceString = shared.DedalusModelSettingsToolChoiceString

// This is an alias to an internal value.
const DedalusModelSettingsToolChoiceStringAuto = shared.DedalusModelSettingsToolChoiceStringAuto

// This is an alias to an internal value.
const DedalusModelSettingsToolChoiceStringRequired = shared.DedalusModelSettingsToolChoiceStringRequired

// This is an alias to an internal value.
const DedalusModelSettingsToolChoiceStringNone = shared.DedalusModelSettingsToolChoiceStringNone

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceMapParam = shared.DedalusModelSettingsToolChoiceMapParam

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceMCPToolChoiceParam = shared.DedalusModelSettingsToolChoiceMCPToolChoiceParam

// This is an alias to an internal type.
type DedalusModelSettingsTruncation = shared.DedalusModelSettingsTruncation

// This is an alias to an internal value.
const DedalusModelSettingsTruncationAuto = shared.DedalusModelSettingsTruncationAuto

// This is an alias to an internal value.
const DedalusModelSettingsTruncationDisabled = shared.DedalusModelSettingsTruncationDisabled

// Dedalus model choice - either a string ID or DedalusModel configuration object.
//
// This is an alias to an internal type.
type DedalusModelChoiceUnionParam = shared.DedalusModelChoiceUnionParam

// JSON object response format. An older method of generating JSON responses. Using
// `json_schema` is recommended for models that support it. Note that the model
// will not generate JSON without a system or user message instructing it to do so.
//
// This is an alias to an internal type.
type ResponseFormatJSONObjectParam = shared.ResponseFormatJSONObjectParam

// This is an alias to an internal type.
type ResponseFormatJSONObjectType = shared.ResponseFormatJSONObjectType

// This is an alias to an internal value.
const ResponseFormatJSONObjectTypeJSONObject = shared.ResponseFormatJSONObjectTypeJSONObject

// JSON Schema response format. Used to generate structured JSON responses. Learn
// more about
// [Structured Outputs](https://platform.openai.com/docs/guides/structured-outputs).
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaParam = shared.ResponseFormatJSONSchemaParam

// Structured Outputs configuration options, including a JSON Schema.
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaJSONSchemaParam = shared.ResponseFormatJSONSchemaJSONSchemaParam

// This is an alias to an internal type.
type ResponseFormatJSONSchemaType = shared.ResponseFormatJSONSchemaType

// This is an alias to an internal value.
const ResponseFormatJSONSchemaTypeJSONSchema = shared.ResponseFormatJSONSchemaTypeJSONSchema

// Default response format. Used to generate text responses.
//
// This is an alias to an internal type.
type ResponseFormatTextParam = shared.ResponseFormatTextParam

// This is an alias to an internal type.
type ResponseFormatTextType = shared.ResponseFormatTextType

// This is an alias to an internal value.
const ResponseFormatTextTypeText = shared.ResponseFormatTextTypeText
