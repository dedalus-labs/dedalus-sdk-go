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

// Unified model metadata across all providers.
//
// Combines provider-specific schemas into a single, consistent format. Fields that
// aren't available from a provider are set to None.
//
// This is an alias to an internal type.
type Model = shared.Model

// Provider that hosts this model
//
// This is an alias to an internal type.
type ModelProvider = shared.ModelProvider

// Equals "openai"
const ModelProviderOpenAI = shared.ModelProviderOpenAI

// Equals "anthropic"
const ModelProviderAnthropic = shared.ModelProviderAnthropic

// Equals "google"
const ModelProviderGoogle = shared.ModelProviderGoogle

// Equals "xai"
const ModelProviderXai = shared.ModelProviderXai

// Equals "mistral"
const ModelProviderMistral = shared.ModelProviderMistral

// Equals "groq"
const ModelProviderGroq = shared.ModelProviderGroq

// Normalized model capabilities across all providers.
//
// This is an alias to an internal type.
type ModelCapabilities = shared.ModelCapabilities

// Provider-declared default parameters for model generation.
//
// This is an alias to an internal type.
type ModelDefaults = shared.ModelDefaults
