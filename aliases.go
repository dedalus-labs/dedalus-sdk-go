// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apierror"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

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
type DedalusModelSettingsStopUnionParam = shared.DedalusModelSettingsStopUnionParam

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceUnionParam = shared.DedalusModelSettingsToolChoiceUnionParam

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceString = shared.DedalusModelSettingsToolChoiceString

// This is an alias to an internal type.
type DedalusModelSettingsToolChoiceMCPToolChoiceParam = shared.DedalusModelSettingsToolChoiceMCPToolChoiceParam

// Dedalus model choice - either a string ID or DedalusModel configuration object.
//
// This is an alias to an internal type.
type DedalusModelChoiceUnionParam = shared.DedalusModelChoiceUnionParam

// JSON object response format. An older method of generating JSON responses.
//
// Using `json_schema` is recommended for models that support it. Note that the
// model will not generate JSON without a system or user message instructing it to
// do so.
//
// Fields:
//
// - type (required): Literal['json_object']
//
// This is an alias to an internal type.
type ResponseFormatJSONObjectParam = shared.ResponseFormatJSONObjectParam

// JSON Schema response format. Used to generate structured JSON responses.
//
// Learn more about
// [Structured Outputs](https://platform.openai.com/docs/guides/structured-outputs).
//
// Fields:
//
// - type (required): Literal['json_schema']
// - json_schema (required): JSONSchema
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaParam = shared.ResponseFormatJSONSchemaParam

// Structured Outputs configuration options, including a JSON Schema.
//
// This is an alias to an internal type.
type ResponseFormatJSONSchemaJSONSchemaParam = shared.ResponseFormatJSONSchemaJSONSchemaParam

// Default response format. Used to generate text responses.
//
// Fields:
//
// - type (required): Literal['text']
//
// This is an alias to an internal type.
type ResponseFormatTextParam = shared.ResponseFormatTextParam
