package _youtube

type SearchSuggestions struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
	} `json:"responseContext"`
	Contents []struct {
		SearchSuggestionsSectionRenderer struct {
			Contents []struct {
				SearchSuggestionRenderer struct {
					Suggestion struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"suggestion"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						SearchEndpoint      struct {
							Query string `json:"query"`
						} `json:"searchEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					Icon           struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
				} `json:"searchSuggestionRenderer"`
			} `json:"contents"`
		} `json:"searchSuggestionsSectionRenderer"`
	} `json:"contents"`
	TrackingParams string `json:"trackingParams"`
}
