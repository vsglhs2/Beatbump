import type { MoodsAndGenresItem } from "$lib/types";
import { error } from "@sveltejs/kit";
import type { ParsedCarousel } from "../api/_lib/models/Carousel";
import { SERVER_DOMAIN } from "../../../env";

export const load = async ({ fetch }) => {

    const response = await fetch<{
		carousels: (ParsedCarousel & {
			items: (ParsedCarousel["items"][number] | MoodsAndGenresItem)[];
		})[];
	}>(`${SERVER_DOMAIN}/api/v1/trending`).then((res) => {
		if (!res.ok) {
			throw error(res.status, res.statusText);
		}
		return res.json();
	});
	return response;
};
