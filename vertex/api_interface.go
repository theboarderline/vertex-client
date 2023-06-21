package vertexai

type Request struct {
	Instances  []Instance  `json:"instances"`
	Parameters []Parameter `json:"parameters"`
}

type Instance struct {
	Content string `json:"content"`
	Prefix  string `json:"prefix"`
}

type Parameter struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens int     `json:"maxOutputTokens"`
	TopP            float64 `json:"topP"`
	TopK            int     `json:"topK"`
}

type Response struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`

		Details []struct {
			Type     string `json:"@type"`
			Reason   string `json:"reason"`
			Domain   string `json:"domain"`
			Metadata struct {
				Permission string `json:"permission"`
				Resource   string `json:"resource"`
			} `json:"metadata"`
		} `json:"details"`
	} `json:"error"`

	Predictions []struct {
		Content          string `json:"content"`
		SafetyAttributes struct {
			Blocked    bool      `json:"blocked"`
			Categories []string  `json:"categories"`
			Scores     []float64 `json:"scores"`
		} `json:"safetyAttributes"`
		CitationMetadata struct {
			Citations []interface{} `json:"citations"`
		} `json:"citationMetadata"`
	} `json:"predictions"`
}
