package _youtube

type NextResponse struct {
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
		SingleColumnMusicWatchNextResultsRenderer struct {
			TabbedRenderer struct {
				WatchNextTabbedResultsRenderer struct {
					Tabs []struct {
						TabRenderer struct {
							Title   string `json:"title"`
							Content struct {
								MusicQueueRenderer struct {
									Content struct {
										PlaylistPanelRenderer struct {
											Title    string `json:"title"`
											Contents []struct {
												PlaylistPanelVideoRenderer PlaylistPanelVideoRenderer `json:"playlistPanelVideoRenderer"`
											} `json:"contents"`
											PlaylistId    string `json:"playlistId"`
											IsInfinite    bool   `json:"isInfinite"`
											Continuations []struct {
												NextRadioContinuationData struct {
													Continuation        string `json:"continuation"`
													ClickTrackingParams string `json:"clickTrackingParams"`
												} `json:"nextRadioContinuationData"`
											} `json:"continuations"`
											ShortBylineText struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														BrowseEndpoint      struct {
															BrowseId         string `json:"browseId"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"shortBylineText"`
											LongBylineText struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														BrowseEndpoint      struct {
															BrowseId         string `json:"browseId"`
															CanonicalBaseUrl string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"longBylineText"`
											TrackingParams string `json:"trackingParams"`
											TitleText      struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														BrowseEndpoint      struct {
															BrowseId string `json:"browseId"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"titleText"`
										} `json:"playlistPanelRenderer"`
									} `json:"content"`
									Hack   bool `json:"hack"`
									Header struct {
										MusicQueueHeaderRenderer struct {
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"title"`
											Subtitle struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"subtitle"`
											Buttons []struct {
												ChipCloudChipRenderer struct {
													Style struct {
														StyleType string `json:"styleType"`
													} `json:"style"`
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														ModalEndpoint       struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"title"`
																	Content struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				Runs []struct {
																					Text string `json:"text"`
																				} `json:"runs"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				SignInEndpoint      struct {
																					Hack bool `json:"hack"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
														SaveQueueToPlaylistCommand struct {
														} `json:"saveQueueToPlaylistCommand"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
													Icon           struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
													IsSelected bool   `json:"isSelected"`
													UniqueId   string `json:"uniqueId"`
												} `json:"chipCloudChipRenderer"`
											} `json:"buttons"`
											TrackingParams string `json:"trackingParams"`
										} `json:"musicQueueHeaderRenderer"`
									} `json:"header"`
								} `json:"musicQueueRenderer"`
							} `json:"content,omitempty"`
							TrackingParams string `json:"trackingParams"`
							Endpoint       struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								BrowseEndpoint      struct {
									BrowseId                              string `json:"browseId"`
									BrowseEndpointContextSupportedConfigs struct {
										BrowseEndpointContextMusicConfig struct {
											PageType string `json:"pageType"`
										} `json:"browseEndpointContextMusicConfig"`
									} `json:"browseEndpointContextSupportedConfigs"`
								} `json:"browseEndpoint"`
							} `json:"endpoint,omitempty"`
							Unselectable bool `json:"unselectable,omitempty"`
						} `json:"tabRenderer"`
					} `json:"tabs"`
				} `json:"watchNextTabbedResultsRenderer"`
			} `json:"tabbedRenderer"`
		} `json:"singleColumnMusicWatchNextResultsRenderer"`
	} `json:"contents"`
	CurrentVideoEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		WatchEndpoint       struct {
			VideoId            string `json:"videoId"`
			PlaylistId         string `json:"playlistId"`
			Index              int    `json:"index"`
			PlaylistSetVideoId string `json:"playlistSetVideoId"`
			LoggingContext     struct {
				VssLoggingContext struct {
					SerializedContextData string `json:"serializedContextData"`
				} `json:"vssLoggingContext"`
			} `json:"loggingContext"`
		} `json:"watchEndpoint"`
	} `json:"currentVideoEndpoint"`
	TrackingParams string `json:"trackingParams"`
	PlayerOverlays struct {
		PlayerOverlayRenderer struct {
			Actions []struct {
				LikeButtonRenderer struct {
					Target struct {
						VideoId string `json:"videoId"`
					} `json:"target"`
					LikeStatus                string `json:"likeStatus"`
					TrackingParams            string `json:"trackingParams"`
					LikesAllowed              bool   `json:"likesAllowed"`
					DislikeNavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"dislikeNavigationEndpoint"`
					LikeCommand struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"likeCommand"`
				} `json:"likeButtonRenderer"`
			} `json:"actions"`
			BrowserMediaSession struct {
				BrowserMediaSessionRenderer struct {
					Album struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"album"`
					ThumbnailDetails struct {
						Thumbnails []struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"thumbnailDetails"`
				} `json:"browserMediaSessionRenderer"`
			} `json:"browserMediaSession"`
		} `json:"playerOverlayRenderer"`
	} `json:"playerOverlays"`
	VideoReporting struct {
		ReportFormModalRenderer struct {
			OptionsSupportedRenderers struct {
				OptionsRenderer struct {
					Items []struct {
						OptionSelectableItemRenderer struct {
							Text struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
							SubmitEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								FlagEndpoint        struct {
									FlagAction string `json:"flagAction"`
								} `json:"flagEndpoint"`
							} `json:"submitEndpoint"`
						} `json:"optionSelectableItemRenderer"`
					} `json:"items"`
					TrackingParams string `json:"trackingParams"`
				} `json:"optionsRenderer"`
			} `json:"optionsSupportedRenderers"`
			TrackingParams string `json:"trackingParams"`
			Title          struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"title"`
			SubmitButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"submitButton"`
			CancelButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"cancelButton"`
			Footer struct {
				Runs []struct {
					Text               string `json:"text"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						UrlEndpoint         struct {
							Url string `json:"url"`
						} `json:"urlEndpoint"`
					} `json:"navigationEndpoint,omitempty"`
				} `json:"runs"`
			} `json:"footer"`
		} `json:"reportFormModalRenderer"`
	} `json:"videoReporting"`
	QueueContextParams string `json:"queueContextParams"`
}

type PlaylistPanelVideoRenderer struct {
	Title struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"title"`
	LongBylineText struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				BrowseEndpoint      struct {
					BrowseId                              string `json:"browseId"`
					BrowseEndpointContextSupportedConfigs struct {
						BrowseEndpointContextMusicConfig struct {
							PageType string `json:"pageType"`
						} `json:"browseEndpointContextMusicConfig"`
					} `json:"browseEndpointContextSupportedConfigs"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint,omitempty"`
		} `json:"runs"`
	} `json:"longBylineText"`
	Thumbnail struct {
		Thumbnails []struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"thumbnails"`
	} `json:"thumbnail"`
	LengthText struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
		Accessibility struct {
			AccessibilityData struct {
				Label string `json:"label"`
			} `json:"accessibilityData"`
		} `json:"accessibility"`
	} `json:"lengthText"`
	Selected           bool `json:"selected"`
	NavigationEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		WatchEndpoint       struct {
			VideoId            string `json:"videoId"`
			PlaylistId         string `json:"playlistId"`
			Index              int    `json:"index"`
			Params             string `json:"params"`
			PlayerParams       string `json:"playerParams"`
			PlaylistSetVideoId string `json:"playlistSetVideoId"`
			LoggingContext     struct {
				VssLoggingContext struct {
					SerializedContextData string `json:"serializedContextData"`
				} `json:"vssLoggingContext"`
			} `json:"loggingContext"`
			WatchEndpointMusicSupportedConfigs struct {
				WatchEndpointMusicConfig struct {
					HasPersistentPlaylistPanel bool   `json:"hasPersistentPlaylistPanel"`
					MusicVideoType             string `json:"musicVideoType"`
				} `json:"watchEndpointMusicConfig"`
			} `json:"watchEndpointMusicSupportedConfigs"`
		} `json:"watchEndpoint"`
	} `json:"navigationEndpoint"`
	VideoId         string `json:"videoId"`
	ShortBylineText struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"shortBylineText"`
	TrackingParams string `json:"trackingParams"`
	Menu           struct {
		MenuRenderer struct {
			Items []struct {
				MenuNavigationItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId        string `json:"videoId"`
							PlaylistId     string `json:"playlistId"`
							Params         string `json:"params"`
							LoggingContext struct {
								VssLoggingContext struct {
									SerializedContextData string `json:"serializedContextData"`
								} `json:"vssLoggingContext"`
							} `json:"loggingContext"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint,omitempty"`
						ModalEndpoint struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint,omitempty"`
						BrowseEndpoint struct {
							BrowseId                              string `json:"browseId"`
							BrowseEndpointContextSupportedConfigs struct {
								BrowseEndpointContextMusicConfig struct {
									PageType string `json:"pageType"`
								} `json:"browseEndpointContextMusicConfig"`
							} `json:"browseEndpointContextSupportedConfigs"`
						} `json:"browseEndpoint,omitempty"`
						ShareEntityEndpoint struct {
							SerializedShareEntity string `json:"serializedShareEntity"`
							SharePanelType        string `json:"sharePanelType"`
						} `json:"shareEntityEndpoint,omitempty"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuNavigationItemRenderer,omitempty"`
				MenuServiceItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						QueueAddEndpoint    struct {
							QueueTarget struct {
								VideoId      string `json:"videoId"`
								OnEmptyQueue struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									WatchEndpoint       struct {
										VideoId string `json:"videoId"`
									} `json:"watchEndpoint"`
								} `json:"onEmptyQueue"`
							} `json:"queueTarget"`
							QueueInsertPosition string `json:"queueInsertPosition"`
							Commands            []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								AddToToastAction    struct {
									Item struct {
										NotificationTextRenderer struct {
											SuccessResponseText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"successResponseText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"notificationTextRenderer"`
									} `json:"item"`
								} `json:"addToToastAction"`
							} `json:"commands"`
						} `json:"queueAddEndpoint,omitempty"`
						RemoveFromQueueEndpoint struct {
							VideoId  string `json:"videoId"`
							Commands []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								AddToToastAction    struct {
									Item struct {
										NotificationTextRenderer struct {
											SuccessResponseText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"successResponseText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"notificationTextRenderer"`
									} `json:"item"`
								} `json:"addToToastAction"`
							} `json:"commands"`
						} `json:"removeFromQueueEndpoint,omitempty"`
						ModalEndpoint struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint,omitempty"`
					} `json:"serviceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuServiceItemRenderer,omitempty"`
				ToggleMenuServiceItemRenderer struct {
					DefaultText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"defaultText"`
					DefaultIcon struct {
						IconType string `json:"iconType"`
					} `json:"defaultIcon"`
					DefaultServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"defaultServiceEndpoint"`
					ToggledText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"toggledText"`
					ToggledIcon struct {
						IconType string `json:"iconType"`
					} `json:"toggledIcon"`
					TrackingParams string `json:"trackingParams"`
				} `json:"toggleMenuServiceItemRenderer,omitempty"`
			} `json:"items"`
			TrackingParams string `json:"trackingParams"`
			Accessibility  struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"menuRenderer"`
	} `json:"menu"`
	PlaylistSetVideoId string `json:"playlistSetVideoId"`
	CanReorder         bool   `json:"canReorder"`
	Badges             []struct {
		MusicInlineBadgeRenderer struct {
			TrackingParams string `json:"trackingParams"`
			Icon           struct {
				IconType string `json:"iconType"`
			} `json:"icon"`
			AccessibilityData struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibilityData"`
		} `json:"musicInlineBadgeRenderer"`
	} `json:"badges,omitempty"`
}
