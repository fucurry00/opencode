package models

const (
	ProviderOpenRouter ModelProvider = "openrouter"

	OpenRouterDeepSeekV3Free        ModelID = "openrouter.deepseek-v3-free"
	OpenRouterDeepSeekR1Free        ModelID = "openrouter.deepseek-r1-free"
	OpenRouterDeepSeekR1Qwen38BFree ModelID = "openrouter.deepseek-r1-qwen3-free"
	OpenRouterLlama4ScoutFree       ModelID = "openrouter.llama4-scout-free"
	OpenRouterLlama4MaverickFree    ModelID = "openrouter.llama4-maverick-free"
)

var OpenRouterModels = map[ModelID]Model{
	OpenRouterDeepSeekR1Free: {
		ID:                 OpenRouterDeepSeekR1Free,
		Name:               "OpenRouter – DeepSeek R1 Free",
		Provider:           ProviderOpenRouter,
		APIModel:           "deepseek/deepseek-r1-0528:free",
		CostPer1MIn:        0,
		CostPer1MInCached:  0,
		CostPer1MOut:       0,
		CostPer1MOutCached: 0,
		ContextWindow:      163_840,
		DefaultMaxTokens:   10000,
	},

	OpenRouterDeepSeekV3Free: {
		ID:                 OpenRouterDeepSeekV3Free,
		Name:               "OpenRouter – DeepSeek V3 Free",
		Provider:           ProviderOpenRouter,
		APIModel:           "deepseek/deepseek-chat-v3-0324:free",
		CostPer1MIn:        0,
		CostPer1MInCached:  0,
		CostPer1MOut:       0,
		CostPer1MOutCached: 0,
		ContextWindow:      163_840,
		DefaultMaxTokens:   10000,
	},
	
	OpenRouterDeepSeekR1Qwen38BFree: {
		ID:                 OpenRouterDeepSeekR1Qwen38BFree,
		Name:               "OpenRouter – DeepSeek R1 Qwen3 8B Free",
		Provider:           ProviderOpenRouter,
		APIModel:           "deepseek/deepseek-r1-0528-qwen3-8b:free",
		CostPer1MIn:        0,
		CostPer1MInCached:  0,
		CostPer1MOut:       0,
		CostPer1MOutCached: 0,
		ContextWindow:      163_840,
		DefaultMaxTokens:   10000,
	},
	
	OpenRouterLlama4ScoutFree: {
		ID:                 OpenRouterLlama4ScoutFree,
		Name:               "OpenRouter – Llama4 Scout Free",
		Provider:           ProviderOpenRouter,
		APIModel:           "meta-llama/llama-4-scout:free",
		CostPer1MIn:        0,
		CostPer1MInCached:  0,
		CostPer1MOut:       0,
		CostPer1MOutCached: 0,
		ContextWindow:      163_840,
		DefaultMaxTokens:   10000,
	},

	OpenRouterLlama4MaverickFree: {
		ID:                 OpenRouterLlama4MaverickFree,
		Name:               "OpenRouter – Llama4 Maverick Free",
		Provider:           ProviderOpenRouter,
		APIModel:           "meta-llama/llama-4-maverick:free",
		CostPer1MIn:        0,
		CostPer1MInCached:  0,
		CostPer1MOut:       0,
		CostPer1MOutCached: 0,
		ContextWindow:      163_840,
		DefaultMaxTokens:   10000,
	},
}
