# Shared Params Types

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#CredentialsBindingSpecParam">CredentialsBindingSpecParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#DedalusModelParam">DedalusModelParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#DedalusModelChoiceUnionParam">DedalusModelChoiceUnionParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#FunctionDefinitionParam">FunctionDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#FunctionParameters">FunctionParameters</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#MCPServerInputUnionParam">MCPServerInputUnionParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#MCPServerSpecParam">MCPServerSpecParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#MCPServersParam">MCPServersParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#ResponseFormatJSONObjectParam">ResponseFormatJSONObjectParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#ResponseFormatJSONSchemaParam">ResponseFormatJSONSchemaParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go/shared#ResponseFormatTextParam">ResponseFormatTextParam</a>

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

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionNewResponse">AudioTranscriptionNewResponse</a>

Methods:

- <code title="post /v1/audio/transcriptions">client.Audio.Transcriptions.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionNewParams">AudioTranscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranscriptionNewResponse">AudioTranscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Translations

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationNewResponse">AudioTranslationNewResponse</a>

Methods:

- <code title="post /v1/audio/translations">client.Audio.Translations.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationNewParams">AudioTranslationNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#AudioTranslationNewResponse">AudioTranslationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionAssistantMessageParam">ChatCompletionAssistantMessageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionAudioParam">ChatCompletionAudioParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionContentPartAudioParam">ChatCompletionContentPartAudioParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionContentPartFileParam">ChatCompletionContentPartFileParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionContentPartImageParam">ChatCompletionContentPartImageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionContentPartRefusalParam">ChatCompletionContentPartRefusalParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionContentPartTextParam">ChatCompletionContentPartTextParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionDeveloperMessageParam">ChatCompletionDeveloperMessageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionFunctionMessageParam">ChatCompletionFunctionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionFunctionsParam">ChatCompletionFunctionsParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionSystemMessageParam">ChatCompletionSystemMessageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionToolMessageParam">ChatCompletionToolMessageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionToolParam">ChatCompletionToolParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionUserMessageParam">ChatCompletionUserMessageParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChoiceDeltaToolCallFunctionParam">ChoiceDeltaToolCallFunctionParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#PredictionContentParam">PredictionContentParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ReasoningParam">ReasoningParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ThinkingConfigDisabledParam">ThinkingConfigDisabledParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ThinkingConfigEnabledParam">ThinkingConfigEnabledParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ToolChoiceUnionParam">ToolChoiceUnionParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ToolChoiceAnyParam">ToolChoiceAnyParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ToolChoiceAutoParam">ToolChoiceAutoParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ToolChoiceNoneParam">ToolChoiceNoneParam</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ToolChoiceToolParam">ToolChoiceToolParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Annotation">Annotation</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletion">ChatCompletion</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionAudio">ChatCompletionAudio</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionChunk">ChatCompletionChunk</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionMessage">ChatCompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionMessageCustomToolCall">ChatCompletionMessageCustomToolCall</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionMessageToolCall">ChatCompletionMessageToolCall</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionTokenLogprob">ChatCompletionTokenLogprob</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Choice">Choice</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChoiceDelta">ChoiceDelta</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChoiceDeltaToolCall">ChoiceDeltaToolCall</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChoiceLogprobs">ChoiceLogprobs</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CompletionTokensDetails">CompletionTokensDetails</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#CompletionUsage">CompletionUsage</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Custom">Custom</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#Function">Function</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#FunctionCall">FunctionCall</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#InputTokenDetails">InputTokenDetails</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#PromptTokensDetails">PromptTokensDetails</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#TopLogprob">TopLogprob</a>
- <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#URLCitation">URLCitation</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go">githubcomdedaluslabsdedalussdkgo</a>.<a href="https://pkg.go.dev/github.com/dedalus-labs/dedalus-sdk-go#ChatCompletion">ChatCompletion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
