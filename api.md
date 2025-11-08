# Shared Params Types

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#DedalusModelParam">DedalusModelParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#DedalusModelChoiceUnionParam">DedalusModelChoiceUnionParam</a>

# Root

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#RootGetResponse">RootGetResponse</a>

Methods:

- <code title="get /">client.Root.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#RootService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#RootGetResponse">RootGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# \_Private

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ListModelsResponse">ListModelsResponse</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Model">Model</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ListModelsResponse">ListModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Params Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CreateEmbeddingRequestParam">CreateEmbeddingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CreateEmbeddingResponse">CreateEmbeddingResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CreateEmbeddingResponse">CreateEmbeddingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Audio

## Speech

Methods:

- <code title="post /v1/audio/speech">client.Audio.Speech.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioSpeechService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioSpeechNewParams">AudioSpeechNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionNewResponseUnion">AudioTranscriptionNewResponseUnion</a>

Methods:

- <code title="post /v1/audio/transcriptions">client.Audio.Transcriptions.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionNewParams">AudioTranscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionNewResponseUnion">AudioTranscriptionNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Translations

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationNewResponseUnion">AudioTranslationNewResponseUnion</a>

Methods:

- <code title="post /v1/audio/translations">client.Audio.Translations.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationNewParams">AudioTranslationNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationNewResponseUnion">AudioTranslationNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Params Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CreateImageRequestParam">CreateImageRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Image">Image</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImagesResponse">ImagesResponse</a>

Methods:

- <code title="post /v1/images/variations">client.Images.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImageService.NewVariation">NewVariation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImageNewVariationParams">ImageNewVariationParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImagesResponse">ImagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/images/edits">client.Images.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImageService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImageEditParams">ImageEditParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImagesResponse">ImagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/images/generations">client.Images.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImageService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImageGenerateParams">ImageGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ImagesResponse">ImagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

## Completions

Params Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelID">ModelID</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ModelsParam">ModelsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionTokenLogprob">ChatCompletionTokenLogprob</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Completion">Completion</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#StreamChunk">StreamChunk</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#TopLogprob">TopLogprob</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Completion">Completion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
