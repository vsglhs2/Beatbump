   import type { PageServerLoad } from "./$types";
import { SERVER_DOMAIN } from "../../../env";
export const prerender = false;
export const load: PageServerLoad = async ({ fetch, url, depends }) => {
	depends("home:load");
	const params = url.searchParams.get("params");
	const data = await fetch(
		`${SERVER_DOMAIN}/api/v1/home.json${params ? `?params=${params}` : ""}`,
	).then((r) => r.json());

	return { ...data, params, path: url.pathname };
};
