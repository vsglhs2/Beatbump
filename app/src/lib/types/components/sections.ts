import type { CarouselHeader, Item, NavigationEndpoint } from "../../types";

export type Sections = [
	{
		type: string;
		header?: CarouselHeader;
		results?: Item[];
		items: [
			{
				type: string;
				thumbnail: string;
				title: string;
				browseId: string;
				videoId?: string;
				autoMixList?: string;
				subtitles: [{ text: string; navigationEndpoint: NavigationEndpoint }];
			},
		];
	},
];
