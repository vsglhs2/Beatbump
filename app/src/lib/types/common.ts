import type { Item } from "../types";

export type BuildMenuParams = {
	item: Item;
	dispatch: (...args: any[]) => void;
	page: string;
	idx: number;
	SITE_ORIGIN_URL: string;
};
