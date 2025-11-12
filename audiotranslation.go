// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/dedalus-labs/dedalus-sdk-go/internal/apiform"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/apijson"
	"github.com/dedalus-labs/dedalus-sdk-go/internal/requestconfig"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/param"
	"github.com/dedalus-labs/dedalus-sdk-go/packages/respjson"
)

// AudioTranslationService contains methods and other services that help with
// interacting with the Dedalus API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioTranslationService] method instead.
type AudioTranslationService struct {
	Options []option.RequestOption
}

// NewAudioTranslationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAudioTranslationService(opts ...option.RequestOption) (r AudioTranslationService) {
	r = AudioTranslationService{}
	r.Options = opts
	return
}

// Translate audio into English.
//
// Translates audio files in any supported language to English text using OpenAI's
// Whisper model. Supports the same audio formats as transcription. Maximum file
// size is 25 MB.
//
// Args: file: Audio file to translate (required) model: Model ID to use (e.g.,
// "openai/whisper-1") prompt: Optional text to guide the model's style
// response_format: Format of the output (json, text, srt, verbose_json, vtt)
// temperature: Sampling temperature between 0 and 1
//
// Returns: Translation object with the English translation
func (r *AudioTranslationService) New(ctx context.Context, body AudioTranslationNewParams, opts ...option.RequestOption) (res *AudioTranslationNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/audio/translations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// AudioTranslationNewResponseUnion contains all possible properties and values
// from [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON],
// [AudioTranslationNewResponseCreateTranslationResponseJSON].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AudioTranslationNewResponseUnion struct {
	// This field is from variant
	// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON].
	Duration float64 `json:"duration"`
	// This field is from variant
	// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON].
	Language string `json:"language"`
	Text     string `json:"text"`
	// This field is from variant
	// [AudioTranslationNewResponseCreateTranslationResponseVerboseJSON].
	Segments []AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment `json:"segments"`
	JSON     struct {
		Duration respjson.Field
		Language respjson.Field
		Text     respjson.Field
		Segments respjson.Field
		raw      string
	} `json:"-"`
}

func (u AudioTranslationNewResponseUnion) AsCreateTranslationResponseVerboseJSON() (v AudioTranslationNewResponseCreateTranslationResponseVerboseJSON) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AudioTranslationNewResponseUnion) AsCreateTranslationResponseJSON() (v AudioTranslationNewResponseCreateTranslationResponseJSON) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AudioTranslationNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AudioTranslationNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - language (required): str
// - duration (required): float
// - text (required): str
// - segments (optional): list[TranscriptionSegment]
type AudioTranslationNewResponseCreateTranslationResponseVerboseJSON struct {
	// The duration of the input audio.
	Duration float64 `json:"duration,required"`
	// The language of the output translation (always `english`).
	Language string `json:"language,required"`
	// The translated text.
	Text string `json:"text,required"`
	// Segments of the translated text and their corresponding details.
	Segments []AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment `json:"segments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Language    respjson.Field
		Text        respjson.Field
		Segments    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseCreateTranslationResponseVerboseJSON) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranslationNewResponseCreateTranslationResponseVerboseJSON) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - id (required): int
// - seek (required): int
// - start (required): float
// - end (required): float
// - text (required): str
// - tokens (required): list[int]
// - temperature (required): float
// - avg_logprob (required): float
// - compression_ratio (required): float
// - no_speech_prob (required): float
type AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment struct {
	// Unique identifier of the segment.
	ID int64 `json:"id,required"`
	// Average logprob of the segment. If the value is lower than -1, consider the
	// logprobs failed.
	AvgLogprob float64 `json:"avg_logprob,required"`
	// Compression ratio of the segment. If the value is greater than 2.4, consider the
	// compression failed.
	CompressionRatio float64 `json:"compression_ratio,required"`
	// End time of the segment in seconds.
	End float64 `json:"end,required"`
	// Probability of no speech in the segment. If the value is higher than 1.0 and the
	// `avg_logprob` is below -1, consider this segment silent.
	NoSpeechProb float64 `json:"no_speech_prob,required"`
	// Seek offset of the segment.
	Seek int64 `json:"seek,required"`
	// Start time of the segment in seconds.
	Start float64 `json:"start,required"`
	// Temperature parameter used for generating the segment.
	Temperature float64 `json:"temperature,required"`
	// Text content of the segment.
	Text string `json:"text,required"`
	// Array of token IDs for the text content.
	Tokens []int64 `json:"tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvgLogprob       respjson.Field
		CompressionRatio respjson.Field
		End              respjson.Field
		NoSpeechProb     respjson.Field
		Seek             respjson.Field
		Start            respjson.Field
		Temperature      respjson.Field
		Text             respjson.Field
		Tokens           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment) RawJSON() string {
	return r.JSON.raw
}
func (r *AudioTranslationNewResponseCreateTranslationResponseVerboseJSONSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fields:
//
// - text (required): str
type AudioTranslationNewResponseCreateTranslationResponseJSON struct {
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponseCreateTranslationResponseJSON) RawJSON() string { return r.JSON.raw }
func (r *AudioTranslationNewResponseCreateTranslationResponseJSON) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioTranslationNewParams struct {
	File           io.Reader          `json:"file,omitzero,required" format:"binary"`
	Model          string             `json:"model,required"`
	Prompt         param.Opt[string]  `json:"prompt,omitzero"`
	ResponseFormat param.Opt[string]  `json:"response_format,omitzero"`
	Temperature    param.Opt[float64] `json:"temperature,omitzero"`
	paramObj
}

func (r AudioTranslationNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
