import type { ExploreSlugResponse } from "../[slug].json/+server.js";
import { SERVER_DOMAIN } from "../../../../env";

export const load = async ({ url, params, fetch }) => {
	const response = await fetch<Awaited<ExploreSlugResponse>>(
        `${SERVER_DOMAIN}/api/v1/explore/${params.slug}`,
	).then((response) => response.json());

	const path = url.pathname;
	// console.log(routeId, params, path, sections, header, type);
	return {
		response,
		path,
	} as const;
};
