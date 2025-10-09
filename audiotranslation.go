// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdedaluslabsdedalussdkgo

import (
	"bytes"
	"context"
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

// Translate audio to English text.
//
// OpenAI only endpoint.
func (r *AudioTranslationService) New(ctx context.Context, body AudioTranslationNewParams, opts ...option.RequestOption) (res *AudioTranslationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/audio/translations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from translation endpoint.
//
// For response_format='json' or 'text', only 'text' is returned. For
// response_format='verbose_json', additional fields are included.
type AudioTranslationNewResponse struct {
	// The translated text (in English)
	Text string `json:"text,required"`
	// The duration of the input audio in seconds
	Duration float64 `json:"duration,nullable"`
	// The language of the output translation (always 'english')
	Language string `json:"language,nullable"`
	// Segments of the translated text and their corresponding details
	Segments []AudioTranslationNewResponseSegment `json:"segments,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Duration    respjson.Field
		Language    respjson.Field
		Segments    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranslationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioTranslationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Segment-level details for transcription.
type AudioTranslationNewResponseSegment struct {
	// Unique identifier of the segment
	ID int64 `json:"id,required"`
	// Average log probability of the segment
	AvgLogprob float64 `json:"avg_logprob,required"`
	// Compression ratio of the segment. If greater than 2.4, consider the compression
	// failed
	CompressionRatio float64 `json:"compression_ratio,required"`
	// End time of the segment in seconds
	End float64 `json:"end,required"`
	// Probability of no speech in the segment. If higher than 1.0 and avg_logprob is
	// below -1, consider this segment silent
	NoSpeechProb float64 `json:"no_speech_prob,required"`
	// Seek offset of the segment
	Seek int64 `json:"seek,required"`
	// Start time of the segment in seconds
	Start float64 `json:"start,required"`
	// Temperature parameter used for generating this segment
	Temperature float64 `json:"temperature,required"`
	// Text content of the segment
	Text string `json:"text,required"`
	// Array of token IDs for the segment
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
func (r AudioTranslationNewResponseSegment) RawJSON() string { return r.JSON.raw }
func (r *AudioTranslationNewResponseSegment) UnmarshalJSON(data []byte) error {
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
