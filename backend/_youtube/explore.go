package _youtube

type Explore struct {
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
	Contents struct {
		SingleColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []struct {
								GridRenderer struct {
									Items []struct {
										MusicNavigationButtonRenderer MusicNavigationButtonRenderer `json:"musicNavigationButtonRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									ItemSize       string `json:"itemSize"`
									Header         struct {
										GridHeaderRenderer struct {
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"title"`
										} `json:"gridHeaderRenderer"`
									} `json:"header"`
								} `json:"gridRenderer"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"singleColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		MusicHeaderRenderer struct {
			Title struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"title"`
			TrackingParams string `json:"trackingParams"`
		} `json:"musicHeaderRenderer"`
	} `json:"header"`
	TrackingParams     string `json:"trackingParams"`
	MaxAgeStoreSeconds int    `json:"maxAgeStoreSeconds"`
}

type MusicNavigationButtonRenderer struct {
	ButtonText struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"buttonText"`
	Solid struct {
		LeftStripeColor int64 `json:"leftStripeColor"`
	} `json:"solid"`
	ClickCommand struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		BrowseEndpoint      struct {
			BrowseID string `json:"browseId"`
			Params   string `json:"params"`
		} `json:"browseEndpoint"`
	} `json:"clickCommand"`
	TrackingParams string `json:"trackingParams"`
}
