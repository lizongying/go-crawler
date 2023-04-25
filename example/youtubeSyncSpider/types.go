package main

type RespSearch struct {
	Contents struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents []struct {
								SearchPyvRenderer struct {
									Ads []struct {
										AdSlotRenderer struct {
											AdSlotMetadata struct {
												SlotID               string `json:"slotId"`
												SlotType             string `json:"slotType"`
												SlotPhysicalPosition int    `json:"slotPhysicalPosition"`
											} `json:"adSlotMetadata"`
											FulfillmentContent struct {
												FulfilledLayout struct {
													InFeedAdLayoutRenderer struct {
														AdLayoutMetadata struct {
															LayoutID            string `json:"layoutId"`
															LayoutType          string `json:"layoutType"`
															AdLayoutLoggingData struct {
																SerializedAdServingDataEntry string `json:"serializedAdServingDataEntry"`
															} `json:"adLayoutLoggingData"`
														} `json:"adLayoutMetadata"`
														RenderingContent struct {
															PromotedVideoRenderer struct {
																VideoID   string `json:"videoId"`
																Thumbnail struct {
																	Thumbnails []struct {
																		URL string `json:"url"`
																	} `json:"thumbnails"`
																} `json:"thumbnail"`
																Title struct {
																	SimpleText string `json:"simpleText"`
																} `json:"title"`
																Description struct {
																	Runs []struct {
																		Text               string `json:"text"`
																		NavigationEndpoint struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			LoggingUrls         []struct {
																				BaseURL string `json:"baseUrl"`
																			} `json:"loggingUrls"`
																			CommandMetadata struct {
																				WebCommandMetadata struct {
																					URL         string `json:"url"`
																					WebPageType string `json:"webPageType"`
																					RootVe      int    `json:"rootVe"`
																				} `json:"webCommandMetadata"`
																			} `json:"commandMetadata"`
																			URLEndpoint struct {
																				URL    string `json:"url"`
																				Target string `json:"target"`
																			} `json:"urlEndpoint"`
																		} `json:"navigationEndpoint"`
																	} `json:"runs"`
																} `json:"description"`
																LongBylineText struct {
																	Runs []struct {
																		Text               string `json:"text"`
																		NavigationEndpoint struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			CommandMetadata     struct {
																				WebCommandMetadata struct {
																					URL         string `json:"url"`
																					WebPageType string `json:"webPageType"`
																					RootVe      int    `json:"rootVe"`
																					APIURL      string `json:"apiUrl"`
																				} `json:"webCommandMetadata"`
																			} `json:"commandMetadata"`
																			BrowseEndpoint struct {
																				BrowseID         string `json:"browseId"`
																				CanonicalBaseURL string `json:"canonicalBaseUrl"`
																			} `json:"browseEndpoint"`
																		} `json:"navigationEndpoint"`
																	} `json:"runs"`
																} `json:"longBylineText"`
																ShortBylineText struct {
																	Runs []struct {
																		Text               string `json:"text"`
																		NavigationEndpoint struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			CommandMetadata     struct {
																				WebCommandMetadata struct {
																					URL         string `json:"url"`
																					WebPageType string `json:"webPageType"`
																					RootVe      int    `json:"rootVe"`
																					APIURL      string `json:"apiUrl"`
																				} `json:"webCommandMetadata"`
																			} `json:"commandMetadata"`
																			BrowseEndpoint struct {
																				BrowseID         string `json:"browseId"`
																				CanonicalBaseURL string `json:"canonicalBaseUrl"`
																			} `json:"browseEndpoint"`
																		} `json:"navigationEndpoint"`
																	} `json:"runs"`
																} `json:"shortBylineText"`
																LengthText struct {
																	Accessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"accessibility"`
																	SimpleText string `json:"simpleText"`
																} `json:"lengthText"`
																NavigationEndpoint struct {
																	ClickTrackingParams string `json:"clickTrackingParams"`
																	CommandMetadata     struct {
																		WebCommandMetadata struct {
																			URL         string `json:"url"`
																			WebPageType string `json:"webPageType"`
																			RootVe      int    `json:"rootVe"`
																		} `json:"webCommandMetadata"`
																	} `json:"commandMetadata"`
																	URLEndpoint struct {
																		URL string `json:"url"`
																	} `json:"urlEndpoint"`
																} `json:"navigationEndpoint"`
																CtaRenderer struct {
																	ButtonRenderer struct {
																		Style string `json:"style"`
																		Text  struct {
																			SimpleText string `json:"simpleText"`
																		} `json:"text"`
																		Icon struct {
																			IconType string `json:"iconType"`
																		} `json:"icon"`
																		TrackingParams string `json:"trackingParams"`
																		Command        struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			CommandMetadata     struct {
																				WebCommandMetadata struct {
																					URL         string `json:"url"`
																					WebPageType string `json:"webPageType"`
																					RootVe      int    `json:"rootVe"`
																				} `json:"webCommandMetadata"`
																			} `json:"commandMetadata"`
																			URLEndpoint struct {
																				URL    string `json:"url"`
																				Target string `json:"target"`
																			} `json:"urlEndpoint"`
																		} `json:"command"`
																		IconPosition string `json:"iconPosition"`
																	} `json:"buttonRenderer"`
																} `json:"ctaRenderer"`
																ImpressionUrls    []string `json:"impressionUrls"`
																ClickTrackingUrls []string `json:"clickTrackingUrls"`
																TrackingParams    string   `json:"trackingParams"`
																Menu              struct {
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
																					OpenPopupAction     struct {
																						Popup struct {
																							AboutThisAdRenderer struct {
																								URL struct {
																									PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
																								} `json:"url"`
																								TrackingParams string `json:"trackingParams"`
																							} `json:"aboutThisAdRenderer"`
																						} `json:"popup"`
																						PopupType string `json:"popupType"`
																					} `json:"openPopupAction"`
																				} `json:"navigationEndpoint"`
																				TrackingParams string `json:"trackingParams"`
																			} `json:"menuNavigationItemRenderer"`
																		} `json:"items"`
																		TrackingParams string `json:"trackingParams"`
																	} `json:"menuRenderer"`
																} `json:"menu"`
																ThumbnailOverlays []struct {
																	ThumbnailOverlayTimeStatusRenderer struct {
																		Text struct {
																			Accessibility struct {
																				AccessibilityData struct {
																					Label string `json:"label"`
																				} `json:"accessibilityData"`
																			} `json:"accessibility"`
																			SimpleText string `json:"simpleText"`
																		} `json:"text"`
																		Style string `json:"style"`
																	} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
																	ThumbnailOverlayToggleButtonRenderer struct {
																		IsToggled     bool `json:"isToggled"`
																		UntoggledIcon struct {
																			IconType string `json:"iconType"`
																		} `json:"untoggledIcon"`
																		ToggledIcon struct {
																			IconType string `json:"iconType"`
																		} `json:"toggledIcon"`
																		UntoggledTooltip         string `json:"untoggledTooltip"`
																		ToggledTooltip           string `json:"toggledTooltip"`
																		UntoggledServiceEndpoint struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			CommandMetadata     struct {
																				WebCommandMetadata struct {
																					SendPost bool   `json:"sendPost"`
																					APIURL   string `json:"apiUrl"`
																				} `json:"webCommandMetadata"`
																			} `json:"commandMetadata"`
																			PlaylistEditEndpoint struct {
																				PlaylistID string `json:"playlistId"`
																				Actions    []struct {
																					AddedVideoID string `json:"addedVideoId"`
																					Action       string `json:"action"`
																				} `json:"actions"`
																			} `json:"playlistEditEndpoint"`
																		} `json:"untoggledServiceEndpoint"`
																		ToggledServiceEndpoint struct {
																			ClickTrackingParams string `json:"clickTrackingParams"`
																			CommandMetadata     struct {
																				WebCommandMetadata struct {
																					SendPost bool   `json:"sendPost"`
																					APIURL   string `json:"apiUrl"`
																				} `json:"webCommandMetadata"`
																			} `json:"commandMetadata"`
																			PlaylistEditEndpoint struct {
																				PlaylistID string `json:"playlistId"`
																				Actions    []struct {
																					Action         string `json:"action"`
																					RemovedVideoID string `json:"removedVideoId"`
																				} `json:"actions"`
																			} `json:"playlistEditEndpoint"`
																		} `json:"toggledServiceEndpoint"`
																		UntoggledAccessibility struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"untoggledAccessibility"`
																		ToggledAccessibility struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"toggledAccessibility"`
																		TrackingParams string `json:"trackingParams"`
																	} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
																} `json:"thumbnailOverlays"`
																RichThumbnail struct {
																	MovingThumbnailRenderer struct {
																		MovingThumbnailDetails struct {
																			Thumbnails []struct {
																				URL    string `json:"url"`
																				Width  int    `json:"width"`
																				Height int    `json:"height"`
																			} `json:"thumbnails"`
																			LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
																		} `json:"movingThumbnailDetails"`
																		EnableHoveredLogging bool `json:"enableHoveredLogging"`
																		EnableOverlay        bool `json:"enableOverlay"`
																	} `json:"movingThumbnailRenderer"`
																} `json:"richThumbnail"`
																ActiveView struct {
																	ViewableCommands []struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		LoggingUrls         []struct {
																			BaseURL string `json:"baseUrl"`
																		} `json:"loggingUrls"`
																		PingingEndpoint struct {
																			Hack bool `json:"hack"`
																		} `json:"pingingEndpoint"`
																	} `json:"viewableCommands"`
																	EndOfSessionCommands []struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		LoggingUrls         []struct {
																			BaseURL string `json:"baseUrl"`
																		} `json:"loggingUrls"`
																		PingingEndpoint struct {
																			Hack bool `json:"hack"`
																		} `json:"pingingEndpoint"`
																	} `json:"endOfSessionCommands"`
																	RegexURIMacroValidator struct {
																		EmptyMap bool `json:"emptyMap"`
																	} `json:"regexUriMacroValidator"`
																} `json:"activeView"`
																AdPlaybackContextParams string `json:"adPlaybackContextParams"`
																AdBadge                 struct {
																	MetadataBadgeRenderer struct {
																		Style          string `json:"style"`
																		Label          string `json:"label"`
																		TrackingParams string `json:"trackingParams"`
																	} `json:"metadataBadgeRenderer"`
																} `json:"adBadge"`
															} `json:"promotedVideoRenderer"`
														} `json:"renderingContent"`
													} `json:"inFeedAdLayoutRenderer"`
												} `json:"fulfilledLayout"`
											} `json:"fulfillmentContent"`
											EnablePacfLoggingWeb bool `json:"enablePacfLoggingWeb"`
										} `json:"adSlotRenderer"`
									} `json:"ads"`
									TrackingParams string `json:"trackingParams"`
								} `json:"searchPyvRenderer,omitempty"`
								VideoRenderer struct {
									VideoID   string `json:"videoId"`
									Thumbnail struct {
										Thumbnails []struct {
											URL    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
									} `json:"title"`
									LongBylineText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"longBylineText"`
									PublishedTimeText struct {
										SimpleText string `json:"simpleText"`
									} `json:"publishedTimeText"`
									LengthText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"lengthText"`
									ViewCountText struct {
										SimpleText string `json:"simpleText"`
									} `json:"viewCountText"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoID                            string `json:"videoId"`
											Params                             string `json:"params"`
											WatchEndpointSupportedOnesieConfig struct {
												HTML5PlaybackOnesieConfig struct {
													CommonConfig struct {
														URL string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint"`
									OwnerBadges []struct {
										MetadataBadgeRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Style             string `json:"style"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"metadataBadgeRenderer"`
									} `json:"ownerBadges"`
									OwnerText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"ownerText"`
									ShortBylineText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"shortBylineText"`
									TrackingParams     string `json:"trackingParams"`
									ShowActionMenu     bool   `json:"showActionMenu"`
									ShortViewCountText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"shortViewCountText"`
									Menu struct {
										MenuRenderer struct {
											Items []struct {
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
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool `json:"sendPost"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														SignalServiceEndpoint struct {
															Signal  string `json:"signal"`
															Actions []struct {
																ClickTrackingParams  string `json:"clickTrackingParams"`
																AddToPlaylistCommand struct {
																	OpenMiniplayer      bool   `json:"openMiniplayer"`
																	VideoID             string `json:"videoId"`
																	ListType            string `json:"listType"`
																	OnCreateListCommand struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				APIURL   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		CreatePlaylistServiceEndpoint struct {
																			VideoIds []string `json:"videoIds"`
																			Params   string   `json:"params"`
																		} `json:"createPlaylistServiceEndpoint"`
																	} `json:"onCreateListCommand"`
																	VideoIds []string `json:"videoIds"`
																} `json:"addToPlaylistCommand"`
															} `json:"actions"`
														} `json:"signalServiceEndpoint"`
													} `json:"serviceEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"menuServiceItemRenderer"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"menuRenderer"`
									} `json:"menu"`
									ChannelThumbnailSupportedRenderers struct {
										ChannelThumbnailWithLinkRenderer struct {
											Thumbnail struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
											Accessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"channelThumbnailWithLinkRenderer"`
									} `json:"channelThumbnailSupportedRenderers"`
									ThumbnailOverlays []struct {
										ThumbnailOverlayTimeStatusRenderer struct {
											Text struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											Style string `json:"style"`
										} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
										ThumbnailOverlayToggleButtonRenderer struct {
											IsToggled     bool `json:"isToggled"`
											UntoggledIcon struct {
												IconType string `json:"iconType"`
											} `json:"untoggledIcon"`
											ToggledIcon struct {
												IconType string `json:"iconType"`
											} `json:"toggledIcon"`
											UntoggledTooltip         string `json:"untoggledTooltip"`
											ToggledTooltip           string `json:"toggledTooltip"`
											UntoggledServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														APIURL   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												PlaylistEditEndpoint struct {
													PlaylistID string `json:"playlistId"`
													Actions    []struct {
														AddedVideoID string `json:"addedVideoId"`
														Action       string `json:"action"`
													} `json:"actions"`
												} `json:"playlistEditEndpoint"`
											} `json:"untoggledServiceEndpoint"`
											ToggledServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														APIURL   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												PlaylistEditEndpoint struct {
													PlaylistID string `json:"playlistId"`
													Actions    []struct {
														Action         string `json:"action"`
														RemovedVideoID string `json:"removedVideoId"`
													} `json:"actions"`
												} `json:"playlistEditEndpoint"`
											} `json:"toggledServiceEndpoint"`
											UntoggledAccessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"untoggledAccessibility"`
											ToggledAccessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"toggledAccessibility"`
											TrackingParams string `json:"trackingParams"`
										} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
										ThumbnailOverlayNowPlayingRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
										} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
										ThumbnailOverlayLoadingPreviewRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
										} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
									} `json:"thumbnailOverlays"`
									RichThumbnail struct {
										MovingThumbnailRenderer struct {
											MovingThumbnailDetails struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
												LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
											} `json:"movingThumbnailDetails"`
											EnableHoveredLogging bool `json:"enableHoveredLogging"`
											EnableOverlay        bool `json:"enableOverlay"`
										} `json:"movingThumbnailRenderer"`
									} `json:"richThumbnail"`
									DetailedMetadataSnippets []struct {
										SnippetText struct {
											Runs []struct {
												Text string `json:"text"`
												Bold bool   `json:"bold,omitempty"`
											} `json:"runs"`
										} `json:"snippetText"`
										SnippetHoverText struct {
											Runs []struct {
												Text string `json:"text"`
											} `json:"runs"`
										} `json:"snippetHoverText"`
										MaxOneLine bool `json:"maxOneLine"`
									} `json:"detailedMetadataSnippets"`
									InlinePlaybackEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoID              string `json:"videoId"`
											PlayerParams         string `json:"playerParams"`
											PlayerExtraURLParams []struct {
												Key   string `json:"key"`
												Value string `json:"value"`
											} `json:"playerExtraUrlParams"`
											WatchEndpointSupportedOnesieConfig struct {
												HTML5PlaybackOnesieConfig struct {
													CommonConfig struct {
														URL string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"inlinePlaybackEndpoint"`
									SearchVideoResultEntityKey string `json:"searchVideoResultEntityKey"`
								} `json:"videoRenderer,omitempty"`
								ShelfRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Content struct {
										VerticalListRenderer struct {
											Items []struct {
												VideoRenderer struct {
													VideoID   string `json:"videoId"`
													Thumbnail struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
													} `json:"title"`
													LongBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		APIURL      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseID         string `json:"browseId"`
																	CanonicalBaseURL string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"longBylineText"`
													PublishedTimeText struct {
														SimpleText string `json:"simpleText"`
													} `json:"publishedTimeText"`
													LengthText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"lengthText"`
													ViewCountText struct {
														SimpleText string `json:"simpleText"`
													} `json:"viewCountText"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoID                            string `json:"videoId"`
															WatchEndpointSupportedOnesieConfig struct {
																HTML5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		URL string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													OwnerText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		APIURL      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseID         string `json:"browseId"`
																	CanonicalBaseURL string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"ownerText"`
													ShortBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		APIURL      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseID         string `json:"browseId"`
																	CanonicalBaseURL string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"shortBylineText"`
													TrackingParams     string `json:"trackingParams"`
													ShowActionMenu     bool   `json:"showActionMenu"`
													ShortViewCountText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"shortViewCountText"`
													Menu struct {
														MenuRenderer struct {
															Items []struct {
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
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool `json:"sendPost"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignalServiceEndpoint struct {
																			Signal  string `json:"signal"`
																			Actions []struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				AddToPlaylistCommand struct {
																					OpenMiniplayer      bool   `json:"openMiniplayer"`
																					VideoID             string `json:"videoId"`
																					ListType            string `json:"listType"`
																					OnCreateListCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost bool   `json:"sendPost"`
																								APIURL   string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						CreatePlaylistServiceEndpoint struct {
																							VideoIds []string `json:"videoIds"`
																							Params   string   `json:"params"`
																						} `json:"createPlaylistServiceEndpoint"`
																					} `json:"onCreateListCommand"`
																					VideoIds []string `json:"videoIds"`
																				} `json:"addToPlaylistCommand"`
																			} `json:"actions"`
																		} `json:"signalServiceEndpoint"`
																	} `json:"serviceEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"menuServiceItemRenderer"`
															} `json:"items"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"menuRenderer"`
													} `json:"menu"`
													ChannelThumbnailSupportedRenderers struct {
														ChannelThumbnailWithLinkRenderer struct {
															Thumbnail struct {
																Thumbnails []struct {
																	URL    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
															} `json:"thumbnail"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		APIURL      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseID         string `json:"browseId"`
																	CanonicalBaseURL string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
															Accessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"channelThumbnailWithLinkRenderer"`
													} `json:"channelThumbnailSupportedRenderers"`
													ThumbnailOverlays []struct {
														ThumbnailOverlayTimeStatusRenderer struct {
															Text struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Style string `json:"style"`
														} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
														ThumbnailOverlayToggleButtonRenderer struct {
															IsToggled     bool `json:"isToggled"`
															UntoggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"untoggledIcon"`
															ToggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"toggledIcon"`
															UntoggledTooltip         string `json:"untoggledTooltip"`
															ToggledTooltip           string `json:"toggledTooltip"`
															UntoggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																PlaylistEditEndpoint struct {
																	PlaylistID string `json:"playlistId"`
																	Actions    []struct {
																		AddedVideoID string `json:"addedVideoId"`
																		Action       string `json:"action"`
																	} `json:"actions"`
																} `json:"playlistEditEndpoint"`
															} `json:"untoggledServiceEndpoint"`
															ToggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																PlaylistEditEndpoint struct {
																	PlaylistID string `json:"playlistId"`
																	Actions    []struct {
																		Action         string `json:"action"`
																		RemovedVideoID string `json:"removedVideoId"`
																	} `json:"actions"`
																} `json:"playlistEditEndpoint"`
															} `json:"toggledServiceEndpoint"`
															UntoggledAccessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"untoggledAccessibility"`
															ToggledAccessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"toggledAccessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
														ThumbnailOverlayNowPlayingRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
														} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
														ThumbnailOverlayLoadingPreviewRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
														} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
													} `json:"thumbnailOverlays"`
													RichThumbnail struct {
														MovingThumbnailRenderer struct {
															MovingThumbnailDetails struct {
																Thumbnails []struct {
																	URL    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
																LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
															} `json:"movingThumbnailDetails"`
															EnableHoveredLogging bool `json:"enableHoveredLogging"`
															EnableOverlay        bool `json:"enableOverlay"`
														} `json:"movingThumbnailRenderer"`
													} `json:"richThumbnail"`
													DetailedMetadataSnippets []struct {
														SnippetText struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"snippetText"`
														SnippetHoverText struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"snippetHoverText"`
														MaxOneLine bool `json:"maxOneLine"`
													} `json:"detailedMetadataSnippets"`
													InlinePlaybackEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoID              string `json:"videoId"`
															PlayerParams         string `json:"playerParams"`
															PlayerExtraURLParams []struct {
																Key   string `json:"key"`
																Value string `json:"value"`
															} `json:"playerExtraUrlParams"`
															WatchEndpointSupportedOnesieConfig struct {
																HTML5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		URL string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"inlinePlaybackEndpoint"`
													SearchVideoResultEntityKey string `json:"searchVideoResultEntityKey"`
												} `json:"videoRenderer"`
											} `json:"items"`
											CollapsedItemCount       int `json:"collapsedItemCount"`
											CollapsedStateButtonText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"collapsedStateButtonText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"verticalListRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
								} `json:"shelfRenderer,omitempty"`
								ReelShelfRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Button struct {
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
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														UserFeedbackEndpoint struct {
															AdditionalDatas []struct {
																UserFeedbackEndpointProductSpecificValueData struct {
																	Key   string `json:"key"`
																	Value string `json:"value"`
																} `json:"userFeedbackEndpointProductSpecificValueData"`
															} `json:"additionalDatas"`
														} `json:"userFeedbackEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuNavigationItemRenderer"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"menuRenderer"`
									} `json:"button"`
									Items []struct {
										ReelItemRenderer struct {
											VideoID  string `json:"videoId"`
											Headline struct {
												SimpleText string `json:"simpleText"`
											} `json:"headline"`
											Thumbnail struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
												IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
											} `json:"thumbnail"`
											ViewCountText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"viewCountText"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ReelWatchEndpoint struct {
													VideoID      string `json:"videoId"`
													PlayerParams string `json:"playerParams"`
													Thumbnail    struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
														IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
													} `json:"thumbnail"`
													Overlay struct {
														ReelPlayerOverlayRenderer struct {
															Style                     string `json:"style"`
															TrackingParams            string `json:"trackingParams"`
															ReelPlayerNavigationModel string `json:"reelPlayerNavigationModel"`
														} `json:"reelPlayerOverlayRenderer"`
													} `json:"overlay"`
													Params           string `json:"params"`
													SequenceProvider string `json:"sequenceProvider"`
													SequenceParams   string `json:"sequenceParams"`
													LoggingContext   struct {
														VssLoggingContext struct {
															SerializedContextData string `json:"serializedContextData"`
														} `json:"vssLoggingContext"`
														QoeLoggingContext struct {
															SerializedContextData string `json:"serializedContextData"`
														} `json:"qoeLoggingContext"`
													} `json:"loggingContext"`
												} `json:"reelWatchEndpoint"`
											} `json:"navigationEndpoint"`
											Menu struct {
												MenuRenderer struct {
													Items []struct {
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
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																GetReportFormEndpoint struct {
																	Params string `json:"params"`
																} `json:"getReportFormEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuServiceItemRenderer,omitempty"`
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
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		IgnoreNavigation bool `json:"ignoreNavigation"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																UserFeedbackEndpoint struct {
																	AdditionalDatas []struct {
																		UserFeedbackEndpointProductSpecificValueData struct {
																			Key   string `json:"key"`
																			Value string `json:"value"`
																		} `json:"userFeedbackEndpointProductSpecificValueData"`
																	} `json:"additionalDatas"`
																} `json:"userFeedbackEndpoint"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"menuNavigationItemRenderer,omitempty"`
													} `json:"items"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuRenderer"`
											} `json:"menu"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
											Style             string `json:"style"`
											VideoType         string `json:"videoType"`
											LoggingDirectives struct {
												TrackingParams string `json:"trackingParams"`
												Visibility     struct {
													Types string `json:"types"`
												} `json:"visibility"`
												EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
											} `json:"loggingDirectives"`
										} `json:"reelItemRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									Icon           struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
								} `json:"reelShelfRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
						} `json:"itemSectionRenderer,omitempty"`
						ContinuationItemRenderer struct {
							Trigger              string `json:"trigger"`
							ContinuationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ContinuationCommand struct {
									Token   string `json:"token"`
									Request string `json:"request"`
								} `json:"continuationCommand"`
							} `json:"continuationEndpoint"`
						} `json:"continuationItemRenderer,omitempty"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
					SubMenu        struct {
						SearchSubMenuRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Groups []struct {
								SearchFilterGroupRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Filters []struct {
										SearchFilterRenderer struct {
											Label struct {
												SimpleText string `json:"simpleText"`
											} `json:"label"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SearchEndpoint struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint"`
											Tooltip        string `json:"tooltip"`
											TrackingParams string `json:"trackingParams"`
										} `json:"searchFilterRenderer"`
									} `json:"filters"`
									TrackingParams string `json:"trackingParams"`
								} `json:"searchFilterGroupRenderer"`
							} `json:"groups"`
							TrackingParams string `json:"trackingParams"`
							Button         struct {
								ToggleButtonRenderer struct {
									Style struct {
										StyleType string `json:"styleType"`
									} `json:"style"`
									IsToggled   bool `json:"isToggled"`
									IsDisabled  bool `json:"isDisabled"`
									DefaultIcon struct {
										IconType string `json:"iconType"`
									} `json:"defaultIcon"`
									DefaultText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"defaultText"`
									Accessibility struct {
										Label string `json:"label"`
									} `json:"accessibility"`
									TrackingParams string `json:"trackingParams"`
									DefaultTooltip string `json:"defaultTooltip"`
									ToggledTooltip string `json:"toggledTooltip"`
									ToggledStyle   struct {
										StyleType string `json:"styleType"`
									} `json:"toggledStyle"`
									AccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibilityData"`
								} `json:"toggleButtonRenderer"`
							} `json:"button"`
						} `json:"searchSubMenuRenderer"`
					} `json:"subMenu"`
					HideBottomSeparator bool   `json:"hideBottomSeparator"`
					TargetID            string `json:"targetId"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	} `json:"contents"`
}

type RespSearchApi struct {
	OnResponseReceivedCommands []struct {
		ClickTrackingParams           string `json:"clickTrackingParams"`
		AppendContinuationItemsAction struct {
			ContinuationItems []struct {
				ItemSectionRenderer struct {
					Contents []struct {
						VideoRenderer struct {
							VideoID   string `json:"videoId"`
							Thumbnail struct {
								Thumbnails []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"title"`
							LongBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												APIURL      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseID         string `json:"browseId"`
											CanonicalBaseURL string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"longBylineText"`
							PublishedTimeText struct {
								SimpleText string `json:"simpleText"`
							} `json:"publishedTimeText"`
							LengthText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"lengthText"`
							ViewCountText struct {
								SimpleText string `json:"simpleText"`
							} `json:"viewCountText"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoID                            string `json:"videoId"`
									Params                             string `json:"params"`
									WatchEndpointSupportedOnesieConfig struct {
										HTML5PlaybackOnesieConfig struct {
											CommonConfig struct {
												URL string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							OwnerBadges []struct {
								MetadataBadgeRenderer struct {
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									Style             string `json:"style"`
									Tooltip           string `json:"tooltip"`
									TrackingParams    string `json:"trackingParams"`
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"metadataBadgeRenderer"`
							} `json:"ownerBadges"`
							OwnerText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												APIURL      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseID         string `json:"browseId"`
											CanonicalBaseURL string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"ownerText"`
							ShortBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												APIURL      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseID         string `json:"browseId"`
											CanonicalBaseURL string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"shortBylineText"`
							TrackingParams     string `json:"trackingParams"`
							ShowActionMenu     bool   `json:"showActionMenu"`
							ShortViewCountText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"shortViewCountText"`
							Menu struct {
								MenuRenderer struct {
									Items []struct {
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
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool `json:"sendPost"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SignalServiceEndpoint struct {
													Signal  string `json:"signal"`
													Actions []struct {
														ClickTrackingParams  string `json:"clickTrackingParams"`
														AddToPlaylistCommand struct {
															OpenMiniplayer      bool   `json:"openMiniplayer"`
															VideoID             string `json:"videoId"`
															ListType            string `json:"listType"`
															OnCreateListCommand struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																CreatePlaylistServiceEndpoint struct {
																	VideoIds []string `json:"videoIds"`
																	Params   string   `json:"params"`
																} `json:"createPlaylistServiceEndpoint"`
															} `json:"onCreateListCommand"`
															VideoIds []string `json:"videoIds"`
														} `json:"addToPlaylistCommand"`
													} `json:"actions"`
												} `json:"signalServiceEndpoint"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuServiceItemRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									Accessibility  struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"menuRenderer"`
							} `json:"menu"`
							ChannelThumbnailSupportedRenderers struct {
								ChannelThumbnailWithLinkRenderer struct {
									Thumbnail struct {
										Thumbnails []struct {
											URL    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												APIURL      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseID         string `json:"browseId"`
											CanonicalBaseURL string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"channelThumbnailWithLinkRenderer"`
							} `json:"channelThumbnailSupportedRenderers"`
							ThumbnailOverlays []struct {
								ThumbnailOverlayTimeStatusRenderer struct {
									Text struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Style string `json:"style"`
								} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
								ThumbnailOverlayToggleButtonRenderer struct {
									IsToggled     bool `json:"isToggled"`
									UntoggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"untoggledIcon"`
									ToggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"toggledIcon"`
									UntoggledTooltip         string `json:"untoggledTooltip"`
									ToggledTooltip           string `json:"toggledTooltip"`
									UntoggledServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										PlaylistEditEndpoint struct {
											PlaylistID string `json:"playlistId"`
											Actions    []struct {
												AddedVideoID string `json:"addedVideoId"`
												Action       string `json:"action"`
											} `json:"actions"`
										} `json:"playlistEditEndpoint"`
									} `json:"untoggledServiceEndpoint"`
									ToggledServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										PlaylistEditEndpoint struct {
											PlaylistID string `json:"playlistId"`
											Actions    []struct {
												Action         string `json:"action"`
												RemovedVideoID string `json:"removedVideoId"`
											} `json:"actions"`
										} `json:"playlistEditEndpoint"`
									} `json:"toggledServiceEndpoint"`
									UntoggledAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"untoggledAccessibility"`
									ToggledAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"toggledAccessibility"`
									TrackingParams string `json:"trackingParams"`
								} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
								ThumbnailOverlayNowPlayingRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
								ThumbnailOverlayLoadingPreviewRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
							} `json:"thumbnailOverlays"`
							DetailedMetadataSnippets []struct {
								SnippetText struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"snippetText"`
								SnippetHoverText struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"snippetHoverText"`
								MaxOneLine bool `json:"maxOneLine"`
							} `json:"detailedMetadataSnippets"`
							InlinePlaybackEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoID              string `json:"videoId"`
									PlayerParams         string `json:"playerParams"`
									PlayerExtraURLParams []struct {
										Key   string `json:"key"`
										Value string `json:"value"`
									} `json:"playerExtraUrlParams"`
									WatchEndpointSupportedOnesieConfig struct {
										HTML5PlaybackOnesieConfig struct {
											CommonConfig struct {
												URL string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"inlinePlaybackEndpoint"`
							SearchVideoResultEntityKey string `json:"searchVideoResultEntityKey"`
						} `json:"videoRenderer,omitempty"`
						ReelShelfRenderer struct {
							Title struct {
								SimpleText string `json:"simpleText"`
							} `json:"title"`
							Button struct {
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
												CommandMetadata     struct {
													WebCommandMetadata struct {
														IgnoreNavigation bool `json:"ignoreNavigation"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												UserFeedbackEndpoint struct {
													AdditionalDatas []struct {
														UserFeedbackEndpointProductSpecificValueData struct {
															Key   string `json:"key"`
															Value string `json:"value"`
														} `json:"userFeedbackEndpointProductSpecificValueData"`
													} `json:"additionalDatas"`
												} `json:"userFeedbackEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"menuNavigationItemRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									Accessibility  struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"menuRenderer"`
							} `json:"button"`
							Items []struct {
								ReelItemRenderer struct {
									VideoID  string `json:"videoId"`
									Headline struct {
										SimpleText string `json:"simpleText"`
									} `json:"headline"`
									Thumbnail struct {
										Thumbnails []struct {
											URL    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
										IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
									} `json:"thumbnail"`
									ViewCountText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"viewCountText"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ReelWatchEndpoint struct {
											VideoID      string `json:"videoId"`
											PlayerParams string `json:"playerParams"`
											Thumbnail    struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
												IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
											} `json:"thumbnail"`
											Overlay struct {
												ReelPlayerOverlayRenderer struct {
													Style                     string `json:"style"`
													TrackingParams            string `json:"trackingParams"`
													ReelPlayerNavigationModel string `json:"reelPlayerNavigationModel"`
												} `json:"reelPlayerOverlayRenderer"`
											} `json:"overlay"`
											Params           string `json:"params"`
											SequenceProvider string `json:"sequenceProvider"`
											SequenceParams   string `json:"sequenceParams"`
											LoggingContext   struct {
												VssLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"vssLoggingContext"`
												QoeLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"qoeLoggingContext"`
											} `json:"loggingContext"`
										} `json:"reelWatchEndpoint"`
									} `json:"navigationEndpoint"`
									Menu struct {
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
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														UserFeedbackEndpoint struct {
															AdditionalDatas []struct {
																UserFeedbackEndpointProductSpecificValueData struct {
																	Key   string `json:"key"`
																	Value string `json:"value"`
																} `json:"userFeedbackEndpointProductSpecificValueData"`
															} `json:"additionalDatas"`
														} `json:"userFeedbackEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuNavigationItemRenderer"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
											Accessibility  struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
										} `json:"menuRenderer"`
									} `json:"menu"`
									TrackingParams string `json:"trackingParams"`
									Accessibility  struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
									Style             string `json:"style"`
									VideoType         string `json:"videoType"`
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
									} `json:"loggingDirectives"`
								} `json:"reelItemRenderer"`
							} `json:"items"`
							TrackingParams string `json:"trackingParams"`
							Icon           struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
						} `json:"reelShelfRenderer,omitempty"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
				} `json:"itemSectionRenderer,omitempty"`
				ContinuationItemRenderer struct {
					Trigger              string `json:"trigger"`
					ContinuationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						ContinuationCommand struct {
							Token   string `json:"token"`
							Request string `json:"request"`
						} `json:"continuationCommand"`
					} `json:"continuationEndpoint"`
				} `json:"continuationItemRenderer,omitempty"`
			} `json:"continuationItems"`
			TargetID string `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedCommands"`
}

type RespUserApi struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title          string `json:"title"`
					TrackingParams string `json:"trackingParams"`
					Selected       bool   `json:"selected"`
					Content        struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										ShelfRenderer struct {
											Title struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																APIURL      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseID         string `json:"browseId"`
															Params           string `json:"params"`
															CanonicalBaseURL string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"title"`
											Endpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													Params           string `json:"params"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"endpoint"`
											Content struct {
												HorizontalListRenderer struct {
													Items []struct {
														GridVideoRenderer struct {
															VideoID   string `json:"videoId"`
															Thumbnail struct {
																Thumbnails []struct {
																	URL    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
															} `json:"thumbnail"`
															Title struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"title"`
															PublishedTimeText struct {
																SimpleText string `json:"simpleText"`
															} `json:"publishedTimeText"`
															ViewCountText struct {
																SimpleText string `json:"simpleText"`
															} `json:"viewCountText"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																WatchEndpoint struct {
																	VideoID                            string `json:"videoId"`
																	WatchEndpointSupportedOnesieConfig struct {
																		HTML5PlaybackOnesieConfig struct {
																			CommonConfig struct {
																				URL string `json:"url"`
																			} `json:"commonConfig"`
																		} `json:"html5PlaybackOnesieConfig"`
																	} `json:"watchEndpointSupportedOnesieConfig"`
																} `json:"watchEndpoint"`
															} `json:"navigationEndpoint"`
															OwnerBadges []struct {
																MetadataBadgeRenderer struct {
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	Style             string `json:"style"`
																	Tooltip           string `json:"tooltip"`
																	TrackingParams    string `json:"trackingParams"`
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"metadataBadgeRenderer"`
															} `json:"ownerBadges"`
															TrackingParams     string `json:"trackingParams"`
															ShortViewCountText struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"shortViewCountText"`
															Menu struct {
																MenuRenderer struct {
																	Items []struct {
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
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool `json:"sendPost"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignalServiceEndpoint struct {
																					Signal  string `json:"signal"`
																					Actions []struct {
																						ClickTrackingParams  string `json:"clickTrackingParams"`
																						AddToPlaylistCommand struct {
																							OpenMiniplayer      bool   `json:"openMiniplayer"`
																							VideoID             string `json:"videoId"`
																							ListType            string `json:"listType"`
																							OnCreateListCommand struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								CommandMetadata     struct {
																									WebCommandMetadata struct {
																										SendPost bool   `json:"sendPost"`
																										APIURL   string `json:"apiUrl"`
																									} `json:"webCommandMetadata"`
																								} `json:"commandMetadata"`
																								CreatePlaylistServiceEndpoint struct {
																									VideoIds []string `json:"videoIds"`
																									Params   string   `json:"params"`
																								} `json:"createPlaylistServiceEndpoint"`
																							} `json:"onCreateListCommand"`
																							VideoIds []string `json:"videoIds"`
																						} `json:"addToPlaylistCommand"`
																					} `json:"actions"`
																				} `json:"signalServiceEndpoint"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"menuServiceItemRenderer,omitempty"`
																		MenuServiceItemDownloadRenderer struct {
																			ServiceEndpoint struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				OfflineVideoEndpoint struct {
																					VideoID      string `json:"videoId"`
																					OnAddCommand struct {
																						ClickTrackingParams      string `json:"clickTrackingParams"`
																						GetDownloadActionCommand struct {
																							VideoID string `json:"videoId"`
																							Params  string `json:"params"`
																						} `json:"getDownloadActionCommand"`
																					} `json:"onAddCommand"`
																				} `json:"offlineVideoEndpoint"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"menuServiceItemDownloadRenderer,omitempty"`
																	} `json:"items"`
																	TrackingParams string `json:"trackingParams"`
																	Accessibility  struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"accessibility"`
																} `json:"menuRenderer"`
															} `json:"menu"`
															ThumbnailOverlays []struct {
																ThumbnailOverlayTimeStatusRenderer struct {
																	Text struct {
																		Accessibility struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"accessibility"`
																		SimpleText string `json:"simpleText"`
																	} `json:"text"`
																	Style string `json:"style"`
																} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
																ThumbnailOverlayToggleButtonRenderer struct {
																	IsToggled     bool `json:"isToggled"`
																	UntoggledIcon struct {
																		IconType string `json:"iconType"`
																	} `json:"untoggledIcon"`
																	ToggledIcon struct {
																		IconType string `json:"iconType"`
																	} `json:"toggledIcon"`
																	UntoggledTooltip         string `json:"untoggledTooltip"`
																	ToggledTooltip           string `json:"toggledTooltip"`
																	UntoggledServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				APIURL   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		PlaylistEditEndpoint struct {
																			PlaylistID string `json:"playlistId"`
																			Actions    []struct {
																				AddedVideoID string `json:"addedVideoId"`
																				Action       string `json:"action"`
																			} `json:"actions"`
																		} `json:"playlistEditEndpoint"`
																	} `json:"untoggledServiceEndpoint"`
																	ToggledServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				APIURL   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		PlaylistEditEndpoint struct {
																			PlaylistID string `json:"playlistId"`
																			Actions    []struct {
																				Action         string `json:"action"`
																				RemovedVideoID string `json:"removedVideoId"`
																			} `json:"actions"`
																		} `json:"playlistEditEndpoint"`
																	} `json:"toggledServiceEndpoint"`
																	UntoggledAccessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"untoggledAccessibility"`
																	ToggledAccessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"toggledAccessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
																ThumbnailOverlayNowPlayingRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
															} `json:"thumbnailOverlays"`
															RichThumbnail struct {
																MovingThumbnailRenderer struct {
																	MovingThumbnailDetails struct {
																		Thumbnails []struct {
																			URL    string `json:"url"`
																			Width  int    `json:"width"`
																			Height int    `json:"height"`
																		} `json:"thumbnails"`
																		LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
																	} `json:"movingThumbnailDetails"`
																	EnableHoveredLogging bool `json:"enableHoveredLogging"`
																	EnableOverlay        bool `json:"enableOverlay"`
																} `json:"movingThumbnailRenderer"`
															} `json:"richThumbnail"`
														} `json:"gridVideoRenderer"`
													} `json:"items"`
													TrackingParams   string `json:"trackingParams"`
													VisibleItemCount int    `json:"visibleItemCount"`
													NextButton       struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Icon       struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"nextButton"`
													PreviousButton struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Icon       struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"previousButton"`
												} `json:"horizontalListRenderer"`
											} `json:"content"`
											TrackingParams string `json:"trackingParams"`
											PlayAllButton  struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoID        string `json:"videoId"`
															PlaylistID     string `json:"playlistId"`
															LoggingContext struct {
																VssLoggingContext struct {
																	SerializedContextData string `json:"serializedContextData"`
																} `json:"vssLoggingContext"`
															} `json:"loggingContext"`
															WatchEndpointSupportedOnesieConfig struct {
																HTML5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		URL string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"playAllButton"`
										} `json:"shelfRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
							TrackingParams       string `json:"trackingParams"`
							TargetID             string `json:"targetId"`
							DisablePullToRefresh bool   `json:"disablePullToRefresh"`
						} `json:"sectionListRenderer"`
						RichGridRenderer struct {
							Contents []struct {
								RichItemRenderer struct {
									Content struct {
										VideoRenderer struct {
											VideoID   string `json:"videoId"`
											Thumbnail struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"title"`
											DescriptionSnippet struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"descriptionSnippet"`
											PublishedTimeText struct {
												SimpleText string `json:"simpleText"`
											} `json:"publishedTimeText"`
											LengthText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"lengthText"`
											ViewCountText struct {
												SimpleText string `json:"simpleText"`
											} `json:"viewCountText"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoID                            string `json:"videoId"`
													WatchEndpointSupportedOnesieConfig struct {
														HTML5PlaybackOnesieConfig struct {
															CommonConfig struct {
																URL string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams     string `json:"trackingParams"`
											ShowActionMenu     bool   `json:"showActionMenu"`
											ShortViewCountText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"shortViewCountText"`
											Menu struct {
												MenuRenderer struct {
													Items []struct {
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
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool `json:"sendPost"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignalServiceEndpoint struct {
																	Signal  string `json:"signal"`
																	Actions []struct {
																		ClickTrackingParams  string `json:"clickTrackingParams"`
																		AddToPlaylistCommand struct {
																			OpenMiniplayer      bool   `json:"openMiniplayer"`
																			VideoID             string `json:"videoId"`
																			ListType            string `json:"listType"`
																			OnCreateListCommand struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						APIURL   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				CreatePlaylistServiceEndpoint struct {
																					VideoIds []string `json:"videoIds"`
																					Params   string   `json:"params"`
																				} `json:"createPlaylistServiceEndpoint"`
																			} `json:"onCreateListCommand"`
																			VideoIds []string `json:"videoIds"`
																		} `json:"addToPlaylistCommand"`
																	} `json:"actions"`
																} `json:"signalServiceEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuServiceItemRenderer"`
													} `json:"items"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuRenderer"`
											} `json:"menu"`
											ThumbnailOverlays []struct {
												ThumbnailOverlayTimeStatusRenderer struct {
													Text struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"text"`
													Style string `json:"style"`
												} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
												ThumbnailOverlayToggleButtonRenderer struct {
													IsToggled     bool `json:"isToggled"`
													UntoggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"untoggledIcon"`
													ToggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"toggledIcon"`
													UntoggledTooltip         string `json:"untoggledTooltip"`
													ToggledTooltip           string `json:"toggledTooltip"`
													UntoggledServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistID string `json:"playlistId"`
															Actions    []struct {
																AddedVideoID string `json:"addedVideoId"`
																Action       string `json:"action"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint"`
													} `json:"untoggledServiceEndpoint"`
													ToggledServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistID string `json:"playlistId"`
															Actions    []struct {
																Action         string `json:"action"`
																RemovedVideoID string `json:"removedVideoId"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint"`
													} `json:"toggledServiceEndpoint"`
													UntoggledAccessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"untoggledAccessibility"`
													ToggledAccessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"toggledAccessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
												ThumbnailOverlayNowPlayingRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
												} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
											} `json:"thumbnailOverlays"`
											RichThumbnail struct {
												MovingThumbnailRenderer struct {
													MovingThumbnailDetails struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
														LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
													} `json:"movingThumbnailDetails"`
													EnableHoveredLogging bool `json:"enableHoveredLogging"`
													EnableOverlay        bool `json:"enableOverlay"`
												} `json:"movingThumbnailRenderer"`
											} `json:"richThumbnail"`
										} `json:"videoRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
								} `json:"richItemRenderer,omitempty"`
								ContinuationItemRenderer struct {
									Trigger              string `json:"trigger"`
									ContinuationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"continuationEndpoint"`
								} `json:"continuationItemRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								FeedFilterChipBarRenderer struct {
									Contents []struct {
										ChipCloudChipRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														APIURL   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ContinuationCommand struct {
													Token   string `json:"token"`
													Request string `json:"request"`
													Command struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														ShowReloadUICommand struct {
															TargetID string `json:"targetId"`
														} `json:"showReloadUiCommand"`
													} `json:"command"`
												} `json:"continuationCommand"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
											IsSelected     bool   `json:"isSelected"`
										} `json:"chipCloudChipRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									StyleType      string `json:"styleType"`
								} `json:"feedFilterChipBarRenderer"`
							} `json:"header"`
							TargetID string `json:"targetId"`
							Style    string `json:"style"`
						} `json:"richGridRenderer"`
					} `json:"content"`
				} `json:"tabRenderer,omitempty"`
				ExpandableTabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
				} `json:"expandableTabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		C4TabbedHeaderRenderer struct {
			ChannelID          string `json:"channelId"`
			Title              string `json:"title"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						URL         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						APIURL      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseID         string `json:"browseId"`
					CanonicalBaseURL string `json:"canonicalBaseUrl"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
			Avatar struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			Banner struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"banner"`
			SubscribeButton struct {
				SubscribeButtonRenderer struct {
					ButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"buttonText"`
					Subscribed           bool   `json:"subscribed"`
					Enabled              bool   `json:"enabled"`
					Type                 string `json:"type"`
					ChannelID            string `json:"channelId"`
					ShowPreferences      bool   `json:"showPreferences"`
					SubscribedButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"subscribedButtonText"`
					UnsubscribedButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"unsubscribedButtonText"`
					TrackingParams        string `json:"trackingParams"`
					UnsubscribeButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"unsubscribeButtonText"`
					SubscribeAccessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"subscribeAccessibility"`
					UnsubscribeAccessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"unsubscribeAccessibility"`
					NotificationPreferenceButton struct {
						SubscriptionNotificationToggleButtonRenderer struct {
							States []struct {
								StateID     int `json:"stateId"`
								NextStateID int `json:"nextStateId"`
								State       struct {
									ButtonRenderer struct {
										Style      string `json:"style"`
										Size       string `json:"size"`
										IsDisabled bool   `json:"isDisabled"`
										Icon       struct {
											IconType string `json:"iconType"`
										} `json:"icon"`
										Accessibility struct {
											Label string `json:"label"`
										} `json:"accessibility"`
										TrackingParams    string `json:"trackingParams"`
										AccessibilityData struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibilityData"`
									} `json:"buttonRenderer"`
								} `json:"state"`
							} `json:"states"`
							CurrentStateID int    `json:"currentStateId"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams    string `json:"clickTrackingParams"`
								CommandExecutorCommand struct {
									Commands []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										OpenPopupAction     struct {
											Popup struct {
												MenuPopupRenderer struct {
													Items []struct {
														MenuServiceItemRenderer struct {
															Text struct {
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																ModifyChannelNotificationPreferenceEndpoint struct {
																	Params string `json:"params"`
																} `json:"modifyChannelNotificationPreferenceEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
															IsSelected     bool   `json:"isSelected"`
														} `json:"menuServiceItemRenderer"`
													} `json:"items"`
												} `json:"menuPopupRenderer"`
											} `json:"popup"`
											PopupType string `json:"popupType"`
										} `json:"openPopupAction"`
									} `json:"commands"`
								} `json:"commandExecutorCommand"`
							} `json:"command"`
							TargetID      string `json:"targetId"`
							SecondaryIcon struct {
								IconType string `json:"iconType"`
							} `json:"secondaryIcon"`
						} `json:"subscriptionNotificationToggleButtonRenderer"`
					} `json:"notificationPreferenceButton"`
					SubscribedEntityKey  string `json:"subscribedEntityKey"`
					OnSubscribeEndpoints []struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SubscribeEndpoint struct {
							ChannelIds []string `json:"channelIds"`
							Params     string   `json:"params"`
						} `json:"subscribeEndpoint"`
					} `json:"onSubscribeEndpoints"`
					OnUnsubscribeEndpoints []struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										ConfirmDialogRenderer struct {
											TrackingParams string `json:"trackingParams"`
											DialogMessages []struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"dialogMessages"`
											ConfirmButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													ServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														UnsubscribeEndpoint struct {
															ChannelIds []string `json:"channelIds"`
															Params     string   `json:"params"`
														} `json:"unsubscribeEndpoint"`
													} `json:"serviceEndpoint"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"confirmButton"`
											CancelButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"cancelButton"`
											PrimaryIsCancel bool `json:"primaryIsCancel"`
										} `json:"confirmDialogRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"onUnsubscribeEndpoints"`
				} `json:"subscribeButtonRenderer"`
			} `json:"subscribeButton"`
			VisitTracking struct {
				RemarketingPing string `json:"remarketingPing"`
			} `json:"visitTracking"`
			SubscriberCountText struct {
				Accessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
				SimpleText string `json:"simpleText"`
			} `json:"subscriberCountText"`
			TvBanner struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"tvBanner"`
			MobileBanner struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"mobileBanner"`
			TrackingParams    string `json:"trackingParams"`
			ChannelHandleText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"channelHandleText"`
			Style           string `json:"style"`
			VideosCountText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"videosCountText"`
			Tagline struct {
				ChannelTaglineRenderer struct {
					Content      string `json:"content"`
					MaxLines     int    `json:"maxLines"`
					MoreLabel    string `json:"moreLabel"`
					MoreEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"moreEndpoint"`
					MoreIcon struct {
						IconType string `json:"iconType"`
					} `json:"moreIcon"`
				} `json:"channelTaglineRenderer"`
			} `json:"tagline"`
		} `json:"c4TabbedHeaderRenderer"`
	} `json:"header"`
	Metadata struct {
		ChannelMetadataRenderer struct {
			Title                string   `json:"title"`
			Description          string   `json:"description"`
			RssURL               string   `json:"rssUrl"`
			ChannelConversionURL string   `json:"channelConversionUrl"`
			ExternalID           string   `json:"externalId"`
			Keywords             string   `json:"keywords"`
			OwnerUrls            []string `json:"ownerUrls"`
			Avatar               struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			ChannelURL             string   `json:"channelUrl"`
			IsFamilySafe           bool     `json:"isFamilySafe"`
			AvailableCountryCodes  []string `json:"availableCountryCodes"`
			AndroidDeepLink        string   `json:"androidDeepLink"`
			AndroidAppindexingLink string   `json:"androidAppindexingLink"`
			IosAppindexingLink     string   `json:"iosAppindexingLink"`
			VanityChannelURL       string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata"`
}

type RespVideos struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title          string `json:"title"`
					TrackingParams string `json:"trackingParams"`
					Content        struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										ShelfRenderer struct {
											Title struct {
												Runs []struct {
													Text               string `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																APIURL      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseID         string `json:"browseId"`
															Params           string `json:"params"`
															CanonicalBaseURL string `json:"canonicalBaseUrl"`
														} `json:"browseEndpoint"`
													} `json:"navigationEndpoint"`
												} `json:"runs"`
											} `json:"title"`
											Endpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													Params           string `json:"params"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"endpoint"`
											Content struct {
												HorizontalListRenderer struct {
													Items []struct {
														GridVideoRenderer struct {
															VideoID   string `json:"videoId"`
															Thumbnail struct {
																Thumbnails []struct {
																	URL    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
															} `json:"thumbnail"`
															Title struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"title"`
															PublishedTimeText struct {
																SimpleText string `json:"simpleText"`
															} `json:"publishedTimeText"`
															ViewCountText struct {
																SimpleText string `json:"simpleText"`
															} `json:"viewCountText"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																WatchEndpoint struct {
																	VideoID                            string `json:"videoId"`
																	WatchEndpointSupportedOnesieConfig struct {
																		HTML5PlaybackOnesieConfig struct {
																			CommonConfig struct {
																				URL string `json:"url"`
																			} `json:"commonConfig"`
																		} `json:"html5PlaybackOnesieConfig"`
																	} `json:"watchEndpointSupportedOnesieConfig"`
																} `json:"watchEndpoint"`
															} `json:"navigationEndpoint"`
															OwnerBadges []struct {
																MetadataBadgeRenderer struct {
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	Style             string `json:"style"`
																	Tooltip           string `json:"tooltip"`
																	TrackingParams    string `json:"trackingParams"`
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"metadataBadgeRenderer"`
															} `json:"ownerBadges"`
															TrackingParams     string `json:"trackingParams"`
															ShortViewCountText struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"shortViewCountText"`
															Menu struct {
																MenuRenderer struct {
																	Items []struct {
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
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool `json:"sendPost"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignalServiceEndpoint struct {
																					Signal  string `json:"signal"`
																					Actions []struct {
																						ClickTrackingParams  string `json:"clickTrackingParams"`
																						AddToPlaylistCommand struct {
																							OpenMiniplayer      bool   `json:"openMiniplayer"`
																							VideoID             string `json:"videoId"`
																							ListType            string `json:"listType"`
																							OnCreateListCommand struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								CommandMetadata     struct {
																									WebCommandMetadata struct {
																										SendPost bool   `json:"sendPost"`
																										APIURL   string `json:"apiUrl"`
																									} `json:"webCommandMetadata"`
																								} `json:"commandMetadata"`
																								CreatePlaylistServiceEndpoint struct {
																									VideoIds []string `json:"videoIds"`
																									Params   string   `json:"params"`
																								} `json:"createPlaylistServiceEndpoint"`
																							} `json:"onCreateListCommand"`
																							VideoIds []string `json:"videoIds"`
																						} `json:"addToPlaylistCommand"`
																					} `json:"actions"`
																				} `json:"signalServiceEndpoint"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"menuServiceItemRenderer,omitempty"`
																		MenuServiceItemDownloadRenderer struct {
																			ServiceEndpoint struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				OfflineVideoEndpoint struct {
																					VideoID      string `json:"videoId"`
																					OnAddCommand struct {
																						ClickTrackingParams      string `json:"clickTrackingParams"`
																						GetDownloadActionCommand struct {
																							VideoID string `json:"videoId"`
																							Params  string `json:"params"`
																						} `json:"getDownloadActionCommand"`
																					} `json:"onAddCommand"`
																				} `json:"offlineVideoEndpoint"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"menuServiceItemDownloadRenderer,omitempty"`
																	} `json:"items"`
																	TrackingParams string `json:"trackingParams"`
																	Accessibility  struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"accessibility"`
																} `json:"menuRenderer"`
															} `json:"menu"`
															ThumbnailOverlays []struct {
																ThumbnailOverlayTimeStatusRenderer struct {
																	Text struct {
																		Accessibility struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"accessibility"`
																		SimpleText string `json:"simpleText"`
																	} `json:"text"`
																	Style string `json:"style"`
																} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
																ThumbnailOverlayToggleButtonRenderer struct {
																	IsToggled     bool `json:"isToggled"`
																	UntoggledIcon struct {
																		IconType string `json:"iconType"`
																	} `json:"untoggledIcon"`
																	ToggledIcon struct {
																		IconType string `json:"iconType"`
																	} `json:"toggledIcon"`
																	UntoggledTooltip         string `json:"untoggledTooltip"`
																	ToggledTooltip           string `json:"toggledTooltip"`
																	UntoggledServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				APIURL   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		PlaylistEditEndpoint struct {
																			PlaylistID string `json:"playlistId"`
																			Actions    []struct {
																				AddedVideoID string `json:"addedVideoId"`
																				Action       string `json:"action"`
																			} `json:"actions"`
																		} `json:"playlistEditEndpoint"`
																	} `json:"untoggledServiceEndpoint"`
																	ToggledServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				APIURL   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		PlaylistEditEndpoint struct {
																			PlaylistID string `json:"playlistId"`
																			Actions    []struct {
																				Action         string `json:"action"`
																				RemovedVideoID string `json:"removedVideoId"`
																			} `json:"actions"`
																		} `json:"playlistEditEndpoint"`
																	} `json:"toggledServiceEndpoint"`
																	UntoggledAccessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"untoggledAccessibility"`
																	ToggledAccessibility struct {
																		AccessibilityData struct {
																			Label string `json:"label"`
																		} `json:"accessibilityData"`
																	} `json:"toggledAccessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
																ThumbnailOverlayNowPlayingRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
															} `json:"thumbnailOverlays"`
															RichThumbnail struct {
																MovingThumbnailRenderer struct {
																	MovingThumbnailDetails struct {
																		Thumbnails []struct {
																			URL    string `json:"url"`
																			Width  int    `json:"width"`
																			Height int    `json:"height"`
																		} `json:"thumbnails"`
																		LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
																	} `json:"movingThumbnailDetails"`
																	EnableHoveredLogging bool `json:"enableHoveredLogging"`
																	EnableOverlay        bool `json:"enableOverlay"`
																} `json:"movingThumbnailRenderer"`
															} `json:"richThumbnail"`
														} `json:"gridVideoRenderer"`
													} `json:"items"`
													TrackingParams   string `json:"trackingParams"`
													VisibleItemCount int    `json:"visibleItemCount"`
													NextButton       struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Icon       struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"nextButton"`
													PreviousButton struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Icon       struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"previousButton"`
												} `json:"horizontalListRenderer"`
											} `json:"content"`
											TrackingParams string `json:"trackingParams"`
											PlayAllButton  struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoID        string `json:"videoId"`
															PlaylistID     string `json:"playlistId"`
															LoggingContext struct {
																VssLoggingContext struct {
																	SerializedContextData string `json:"serializedContextData"`
																} `json:"vssLoggingContext"`
															} `json:"loggingContext"`
															WatchEndpointSupportedOnesieConfig struct {
																HTML5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		URL string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"playAllButton"`
										} `json:"shelfRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
							TrackingParams       string `json:"trackingParams"`
							TargetID             string `json:"targetId"`
							DisablePullToRefresh bool   `json:"disablePullToRefresh"`
						} `json:"sectionListRenderer"`
						RichGridRenderer struct {
							Contents []struct {
								RichItemRenderer struct {
									Content struct {
										VideoRenderer struct {
											VideoID   string `json:"videoId"`
											Thumbnail struct {
												Thumbnails []struct {
													URL    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"title"`
											DescriptionSnippet struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"descriptionSnippet"`
											PublishedTimeText struct {
												SimpleText string `json:"simpleText"`
											} `json:"publishedTimeText"`
											LengthText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"lengthText"`
											ViewCountText struct {
												SimpleText string `json:"simpleText"`
											} `json:"viewCountText"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoID                            string `json:"videoId"`
													WatchEndpointSupportedOnesieConfig struct {
														HTML5PlaybackOnesieConfig struct {
															CommonConfig struct {
																URL string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams     string `json:"trackingParams"`
											ShowActionMenu     bool   `json:"showActionMenu"`
											ShortViewCountText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"shortViewCountText"`
											Menu struct {
												MenuRenderer struct {
													Items []struct {
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
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool `json:"sendPost"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignalServiceEndpoint struct {
																	Signal  string `json:"signal"`
																	Actions []struct {
																		ClickTrackingParams  string `json:"clickTrackingParams"`
																		AddToPlaylistCommand struct {
																			OpenMiniplayer      bool   `json:"openMiniplayer"`
																			VideoID             string `json:"videoId"`
																			ListType            string `json:"listType"`
																			OnCreateListCommand struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						APIURL   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				CreatePlaylistServiceEndpoint struct {
																					VideoIds []string `json:"videoIds"`
																					Params   string   `json:"params"`
																				} `json:"createPlaylistServiceEndpoint"`
																			} `json:"onCreateListCommand"`
																			VideoIds []string `json:"videoIds"`
																		} `json:"addToPlaylistCommand"`
																	} `json:"actions"`
																} `json:"signalServiceEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"menuServiceItemRenderer"`
													} `json:"items"`
													TrackingParams string `json:"trackingParams"`
													Accessibility  struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibility"`
												} `json:"menuRenderer"`
											} `json:"menu"`
											ThumbnailOverlays []struct {
												ThumbnailOverlayTimeStatusRenderer struct {
													Text struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"text"`
													Style string `json:"style"`
												} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
												ThumbnailOverlayToggleButtonRenderer struct {
													IsToggled     bool `json:"isToggled"`
													UntoggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"untoggledIcon"`
													ToggledIcon struct {
														IconType string `json:"iconType"`
													} `json:"toggledIcon"`
													UntoggledTooltip         string `json:"untoggledTooltip"`
													ToggledTooltip           string `json:"toggledTooltip"`
													UntoggledServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistID string `json:"playlistId"`
															Actions    []struct {
																AddedVideoID string `json:"addedVideoId"`
																Action       string `json:"action"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint"`
													} `json:"untoggledServiceEndpoint"`
													ToggledServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistID string `json:"playlistId"`
															Actions    []struct {
																Action         string `json:"action"`
																RemovedVideoID string `json:"removedVideoId"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint"`
													} `json:"toggledServiceEndpoint"`
													UntoggledAccessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"untoggledAccessibility"`
													ToggledAccessibility struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"toggledAccessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
												ThumbnailOverlayNowPlayingRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
												} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
											} `json:"thumbnailOverlays"`
											RichThumbnail struct {
												MovingThumbnailRenderer struct {
													MovingThumbnailDetails struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
														LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
													} `json:"movingThumbnailDetails"`
													EnableHoveredLogging bool `json:"enableHoveredLogging"`
													EnableOverlay        bool `json:"enableOverlay"`
												} `json:"movingThumbnailRenderer"`
											} `json:"richThumbnail"`
										} `json:"videoRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
								} `json:"richItemRenderer,omitempty"`
								ContinuationItemRenderer struct {
									Trigger              string `json:"trigger"`
									ContinuationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"continuationEndpoint"`
								} `json:"continuationItemRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								FeedFilterChipBarRenderer struct {
									Contents []struct {
										ChipCloudChipRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														APIURL   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ContinuationCommand struct {
													Token   string `json:"token"`
													Request string `json:"request"`
													Command struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														ShowReloadUICommand struct {
															TargetID string `json:"targetId"`
														} `json:"showReloadUiCommand"`
													} `json:"command"`
												} `json:"continuationCommand"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
											IsSelected     bool   `json:"isSelected"`
										} `json:"chipCloudChipRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									StyleType      string `json:"styleType"`
								} `json:"feedFilterChipBarRenderer"`
							} `json:"header"`
							TargetID string `json:"targetId"`
							Style    string `json:"style"`
						} `json:"richGridRenderer"`
					} `json:"content"`
				} `json:"tabRenderer,omitempty"`
				ExpandableTabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
				} `json:"expandableTabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		C4TabbedHeaderRenderer struct {
			ChannelID          string `json:"channelId"`
			Title              string `json:"title"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						URL         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						APIURL      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseID         string `json:"browseId"`
					CanonicalBaseURL string `json:"canonicalBaseUrl"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
			Avatar struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			Banner struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"banner"`
			HeaderLinks struct {
				ChannelHeaderLinksRenderer struct {
					PrimaryLinks []struct {
						NavigationEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							URLEndpoint struct {
								URL      string `json:"url"`
								Target   string `json:"target"`
								Nofollow bool   `json:"nofollow"`
							} `json:"urlEndpoint"`
						} `json:"navigationEndpoint"`
						Icon struct {
							Thumbnails []struct {
								URL string `json:"url"`
							} `json:"thumbnails"`
						} `json:"icon"`
						Title struct {
							SimpleText string `json:"simpleText"`
						} `json:"title"`
					} `json:"primaryLinks"`
					SecondaryLinks []struct {
						NavigationEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							URLEndpoint struct {
								URL      string `json:"url"`
								Target   string `json:"target"`
								Nofollow bool   `json:"nofollow"`
							} `json:"urlEndpoint"`
						} `json:"navigationEndpoint"`
						Icon struct {
							Thumbnails []struct {
								URL string `json:"url"`
							} `json:"thumbnails"`
						} `json:"icon"`
						Title struct {
							SimpleText string `json:"simpleText"`
						} `json:"title"`
					} `json:"secondaryLinks"`
				} `json:"channelHeaderLinksRenderer"`
			} `json:"headerLinks"`
			SubscribeButton struct {
				SubscribeButtonRenderer struct {
					ButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"buttonText"`
					Subscribed           bool   `json:"subscribed"`
					Enabled              bool   `json:"enabled"`
					Type                 string `json:"type"`
					ChannelID            string `json:"channelId"`
					ShowPreferences      bool   `json:"showPreferences"`
					SubscribedButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"subscribedButtonText"`
					UnsubscribedButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"unsubscribedButtonText"`
					TrackingParams        string `json:"trackingParams"`
					UnsubscribeButtonText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"unsubscribeButtonText"`
					SubscribeAccessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"subscribeAccessibility"`
					UnsubscribeAccessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"unsubscribeAccessibility"`
					NotificationPreferenceButton struct {
						SubscriptionNotificationToggleButtonRenderer struct {
							States []struct {
								StateID     int `json:"stateId"`
								NextStateID int `json:"nextStateId"`
								State       struct {
									ButtonRenderer struct {
										Style      string `json:"style"`
										Size       string `json:"size"`
										IsDisabled bool   `json:"isDisabled"`
										Icon       struct {
											IconType string `json:"iconType"`
										} `json:"icon"`
										Accessibility struct {
											Label string `json:"label"`
										} `json:"accessibility"`
										TrackingParams    string `json:"trackingParams"`
										AccessibilityData struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibilityData"`
									} `json:"buttonRenderer"`
								} `json:"state"`
							} `json:"states"`
							CurrentStateID int    `json:"currentStateId"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams    string `json:"clickTrackingParams"`
								CommandExecutorCommand struct {
									Commands []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										OpenPopupAction     struct {
											Popup struct {
												MenuPopupRenderer struct {
													Items []struct {
														MenuServiceItemRenderer struct {
															Text struct {
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																ModifyChannelNotificationPreferenceEndpoint struct {
																	Params string `json:"params"`
																} `json:"modifyChannelNotificationPreferenceEndpoint"`
															} `json:"serviceEndpoint"`
															TrackingParams string `json:"trackingParams"`
															IsSelected     bool   `json:"isSelected"`
														} `json:"menuServiceItemRenderer"`
													} `json:"items"`
												} `json:"menuPopupRenderer"`
											} `json:"popup"`
											PopupType string `json:"popupType"`
										} `json:"openPopupAction"`
									} `json:"commands"`
								} `json:"commandExecutorCommand"`
							} `json:"command"`
							TargetID      string `json:"targetId"`
							SecondaryIcon struct {
								IconType string `json:"iconType"`
							} `json:"secondaryIcon"`
						} `json:"subscriptionNotificationToggleButtonRenderer"`
					} `json:"notificationPreferenceButton"`
					SubscribedEntityKey  string `json:"subscribedEntityKey"`
					OnSubscribeEndpoints []struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SubscribeEndpoint struct {
							ChannelIds []string `json:"channelIds"`
							Params     string   `json:"params"`
						} `json:"subscribeEndpoint"`
					} `json:"onSubscribeEndpoints"`
					OnUnsubscribeEndpoints []struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										ConfirmDialogRenderer struct {
											TrackingParams string `json:"trackingParams"`
											DialogMessages []struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"dialogMessages"`
											ConfirmButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													ServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																APIURL   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														UnsubscribeEndpoint struct {
															ChannelIds []string `json:"channelIds"`
															Params     string   `json:"params"`
														} `json:"unsubscribeEndpoint"`
													} `json:"serviceEndpoint"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"confirmButton"`
											CancelButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Accessibility struct {
														Label string `json:"label"`
													} `json:"accessibility"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"cancelButton"`
											PrimaryIsCancel bool `json:"primaryIsCancel"`
										} `json:"confirmDialogRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"onUnsubscribeEndpoints"`
				} `json:"subscribeButtonRenderer"`
			} `json:"subscribeButton"`
			SubscriberCountText struct {
				Accessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
				SimpleText string `json:"simpleText"`
			} `json:"subscriberCountText"`
			TvBanner struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"tvBanner"`
			MobileBanner struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"mobileBanner"`
			TrackingParams    string `json:"trackingParams"`
			ChannelHandleText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"channelHandleText"`
			Style           string `json:"style"`
			VideosCountText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"videosCountText"`
			Tagline struct {
				ChannelTaglineRenderer struct {
					Content      string `json:"content"`
					MaxLines     int    `json:"maxLines"`
					MoreLabel    string `json:"moreLabel"`
					MoreEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"moreEndpoint"`
					MoreIcon struct {
						IconType string `json:"iconType"`
					} `json:"moreIcon"`
				} `json:"channelTaglineRenderer"`
			} `json:"tagline"`
		} `json:"c4TabbedHeaderRenderer"`
	} `json:"header"`
	Metadata struct {
		ChannelMetadataRenderer struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			RssURL      string   `json:"rssUrl"`
			ExternalID  string   `json:"externalId"`
			Keywords    string   `json:"keywords"`
			OwnerUrls   []string `json:"ownerUrls"`
			Avatar      struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			ChannelURL             string   `json:"channelUrl"`
			IsFamilySafe           bool     `json:"isFamilySafe"`
			AvailableCountryCodes  []string `json:"availableCountryCodes"`
			AndroidDeepLink        string   `json:"androidDeepLink"`
			AndroidAppindexingLink string   `json:"androidAppindexingLink"`
			IosAppindexingLink     string   `json:"iosAppindexingLink"`
			VanityChannelURL       string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata"`
}

type ExtraSearch struct {
	Keyword       string
	Sp            string
	Page          int
	MaxPage       int
	NextPageToken string
}

type ExtraSearchApi struct {
	Keyword       string
	Sp            string
	Page          int
	MaxPage       int
	NextPageToken string
}

type ExtraVideos struct {
	KeyWord  string
	Id       string
	Key      string
	UserName string
}

type ExtraUserApi struct {
	KeyWord  string
	Id       string
	Key      string
	UserName string
}

type DataUser struct {
	Id          string `bson:"_id" json:"id"`
	UserName    string `bson:"user_name" json:"user_name"`
	Description string `bson:"description" json:"description"`
	Link        string `bson:"link" json:"link"`
	Email       string `bson:"email" json:"email"`
	Followers   int    `bson:"followers" json:"followers"`
	ViewAvg     int    `bson:"view_avg" json:"view_avg"`
	Keyword     string `bson:"keyword" json:"keyword"`
}
