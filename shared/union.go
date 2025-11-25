// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsDedalusModelSettingsStopUnionParam()        {}
func (UnionString) ImplementsDedalusModelSettingsToolChoiceUnionParam()  {}
func (UnionString) ImplementsDedalusModelChoiceUnionParam()              {}
func (UnionString) ImplementsCreateEmbeddingRequestInputUnionParam()     {}
func (UnionString) ImplementsCreateEmbeddingResponseDataEmbeddingUnion() {}
func (UnionString) ImplementsChatCompletionNewParamsModelUnion()         {}
func (UnionString) ImplementsChatCompletionNewParamsMCPServersUnion()    {}
func (UnionString) ImplementsChatCompletionNewParamsMessagesUnion()      {}
func (UnionString) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestDeveloperMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestSystemMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestUserMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestAssistantMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsMessagesMessagesChatCompletionRequestToolMessageContentUnion() {
}
func (UnionString) ImplementsChatCompletionNewParamsStopUnion()              {}
func (UnionString) ImplementsChatCompletionNewParamsSystemInstructionUnion() {}
