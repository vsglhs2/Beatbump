<script
    context="module"
    lang="ts"
>
</script>

<script lang="ts">
    import { browser } from "$app/environment";
    import Header from "$components/Layouts/Header.svelte";
    import { AudioPlayer } from "$lib/player";
    import { settings, type Theme } from "$stores/settings";
    import {onMount} from "svelte";
    import {APIClient} from "$lib/api";
    import {notify} from "$lib/utils";
    const themes: Theme[] = ["Dark", "Dim", "Midnight", "YTM"];

    let oauthInfo = {}
    onMount(async () => {
        try {
            const response = await APIClient.fetch(`/api/v1/deviceOauth`);
            if (!response.ok) {
                console.log(response)
            }
            oauthInfo = await response.json();
        } catch (err) {
            console.log(err)
        }
    });

    function handleStreamSelect() {
        AudioPlayer.dispatch("update:stream_type", {
            type: $settings.playback.Stream ?? "HTTP",
        });
    }

    async function handleOauthComplete() {
        try {
            if (oauthInfo == undefined || oauthInfo.deviceCode == ""){
                return
            }
            notify(`Calling server to complete OAuth flow`, "success");
            const response = await APIClient.fetch(`/api/v1/deviceOauth/${oauthInfo.deviceCode}`);
            if (!response.ok) {
                throw new Error('Failed to fetch data');
            }
            notify(`OAuth flow completed successfully`, "success");
            setTimeout(() => {
                window.location.reload();
            }, 1000)
        } catch (err) {
            notify(`OAuth flow failed`+err, "error");
            console.log(err)
        }
    }

    const updatePrefsCookie = async () => {
        await fetch("/settings/update.json", {
            body: JSON.stringify({
                "Proxy Thumbnails": $settings.network["Proxy Thumbnails"],
                Restricted: $settings.search.Restricted,
            }),
            method: "POST",
        });
    };
</script>

<Header
    title="Settings"
    url="/settings"
    desc="Configure your app settings"
/>
{#if browser}
    <main class="resp-content-width">
        <section>
            <span class="h5">Appearance</span>
            <div class="setting">
                <label for="theme">Theme </label>
                <div class="select">
                    <select
                        name="theme"
                        id="theme"
                        bind:value={$settings["appearance"]["Theme"]}
                    >
                        {#each themes as theme}
                            <option
                                value={theme}
                                selected={$settings["appearance"]["Theme"] === theme}
                            >{theme}</option
                            >
                        {/each}
                    </select>
                </div>
            </div>
            <div class="setting">
                <label>Immersive Queue</label>
                <input
                    type="checkbox"
                    name="immersive-queue"
                    id="immersive-queue"
                    bind:checked={$settings["appearance"]["Immersive Queue"]}
                />
                <label
                    for="immersive-queue"
                    class="switch"
                />
            </div>
        </section>
        <section>
            <span class="h5">Playback</span>
            <div class="setting">
                <label>Dedupe Automix</label>

                <input
                    name="dedupe"
                    id="dedupe"
                    type="checkbox"
                    bind:value={$settings["playback"]["Dedupe Automix"]}
                />
                <label
                    for="dedupe"
                    class="switch"
                />
            </div>
           <!-- <div class="setting">
                <label for="quality">Quality</label>
                <div class="select">
                    <select
                        name="quality"
                        disabled
                        id="quality"
                        bind:value={$settings["playback"]["Quality"]}
                    >
                        {#each ["Normal", "High"] as option}
                            <option
                                value={option}
                                selected={$settings["playback"]["Quality"] === option}
                            >{option}</option
                            >
                        {/each}
                    </select>
                </div>
            </div>-->
            <div class="setting">
                <label for="oauthClientId">
                    Oauth ClientID
                    <span class="">
                        Oauth Client ID
                    </span>
                </label>
                <div class="input-container">
                    <div class="input no-btn mb-1">
                        <input
                            type="text"
                            on:input={(e) => {
								let value = e.currentTarget.value;
                                localStorage.setItem('x-google-client-id',value);
							}}
                            value={localStorage.getItem("x-google-client-id")}
                        />
                    </div>
                </div>
            </div>
            <div class="setting">
                <label for="oauthClientSecret">
                    Oauth Client Secret
                    <span class="">
                        Oauth Client Secret
                    </span>
                </label>
                <div class="input-container">
                    <div class="input no-btn mb-1">
                        <input
                            type="text"
                            on:input={(e) => {
								let value = e.currentTarget.value;
                                localStorage.setItem('x-google-client-secret',value);
							}}
                            value={localStorage.getItem("x-google-client-secret")}
                        />
                    </div>
                </div>
            </div>
            <div class="setting">
                <label>Oauth</label>
                {#if oauthInfo.oauthStart == undefined || oauthInfo.oauthStart === ""}
                    <label>{oauthInfo.status || "Logged in"}</label>
                {:else}
                <a target="_blank" href={oauthInfo.oauthStart}>Start Oauth flow</a>
                <button
                    class="link mt-2"
                    on:click={handleOauthComplete}>Complete</button
                >
                {/if}
            </div>
            <div class="setting">
                <label for="cookieHeader"
                >Cookie header
                    <span class="">
                        Copy the Cookie header value and paste here.
                    </span>
                </label>
                <div class="input-container">
                    <div class="input no-btn mb-1">
                        <input
                            type="text"
                            on:input={(e) => {
								let value = e.currentTarget.value;
                                localStorage.setItem('x-google-cookie',value);
							}}
                            value={localStorage.getItem("x-google-cookie")}
                        />
                    </div>
                </div>
            </div>
            <div class="setting">
                <label for="invidious">
                    Enable invidious companion
                </label>
                <div class="input-container">
                    <div class="input no-btn mb-1">
                        <input
                            type="text"
                            on:input={(e) => {
								let value = e.currentTarget.value;
                                localStorage.setItem('x-use-invidious',value);
							}}
                            value={localStorage.getItem("x-use-invidious")}
                        />
                    </div>
                </div>
            </div>
           <!-- <div class="setting">
                <label>Remember Last Track</label>
                <input
                    name="lasttrack"
                    id="lasttrack"
                    type="checkbox"
                    bind:checked={$settings["playback"]["Remember Last Track"]}
                />
                <label
                    for="lasttrack"
                    class="switch"
                />
            </div>-->
           <!-- <div class="setting">
                <label for="stream">Stream </label>
                <div class="select">
                    <select
                        name="stream"
                        id="stream"
                        bind:value={$settings["playback"]["Stream"]}
                        on:change={handleStreamSelect}
                    >
                        {#each ["HTTP", "HLS"] as option}
                            <option
                                value={option}
                                selected={$settings["playback"]["Stream"] === option}
                            >{option}</option
                            >
                        {/each}
                    </select>
                </div>
            </div>-->

            <div class="setting">
                <!-- svelte-ignore a11y-label-has-associated-control -->
                <label
                >Playback Updates URL
                    <span class=""
                    >Playing a song updates the URL with the song's sharing URL.</span
                    >
                </label>
                <input
                    name="update-url"
                    id="update-url"
                    type="checkbox"
                    bind:checked={$settings["playback"]["Playback Updates URL"]}
                />
                <label
                    for="update-url"
                    class="switch"
                />
            </div>
        </section>
        <!--<section>
            <span class="h5">Network</span>
            <div class="setting">
                <label for="proxy"
                >Audio Proxy Server
                    <span class=""
                    >In order to use HLS streaming, a proxy server must be used. <br
                    />Provide the URL to your own, or you can use the default.</span
                    >
                </label>
                <div class="input-container">
                    <div class="input no-btn mb-1">
                        <input
                            type="url"
                            on:blur={(e) => {
								let value = e.currentTarget.value;

								if (!value.endsWith("/")) value = value + "/";

								if (value.match(/^https?:\/\//i)) {
									$settings["network"]["Stream Proxy Server"] = value;
								} else if (value.match(/^(.[0-9]*\.?){1,4}:[0-9]+/im)) {
									$settings["network"]["Stream Proxy Server"] = value;
								}
							}}
                            on:input={(e) => {
								let value = e.currentTarget.value;

								if (value.match(/^https?:\/\//i)) {
									$settings["network"]["Stream Proxy Server"] = value;
								} else if (value.match(/^(.[0-9]*\.?){1,4}:[0-9]+/im)) {
									$settings["network"]["Stream Proxy Server"] = value;
								}
							}}
                            placeholder="https://hls.beatbump.io/"
                            value={$settings["network"]["Stream Proxy Server"]}
                        />
                    </div>
                    <button
                        class="link mt-2"
                        on:click={() => {
							$settings["network"]["Stream Proxy Server"] =
								"https://hls.beatbump.io/";
						}}>Reset to default</button
                    >
                </div>
            </div>
            <div class="setting">
                <label
                >Proxy Audio
                    <span class="">Proxy audio streams through a server.</span>
                </label>
                <input
                    name="proxy-audio"
                    id="proxy-audio"
                    type="checkbox"
                    bind:checked={$settings["network"]["Proxy Streams"]}
                />
                <label
                    for="proxy-audio"
                    class="switch"
                />
            </div>
        </section>-->
        <section>
           <!-- <span class="h5">Search</span>
            <div class="setting">
                <label for="preserve">Preserve </label>
                <div class="select">
                    <select
                        name="preserve"
                        id="preserve"
                        bind:value={$settings["search"]["Preserve"]}
                    >
                        {#each ["Category", "Query", "Category + Query", "None"] as option}
                            <option
                                value={option}
                                selected={$settings["playback"]["Stream"] === option}
                            >{option}</option
                            >
                        {/each}
                    </select>
                </div>
            </div>-->
           <!-- <div class="setting">
                <label for="restricted"
                >Restricted Mode <span
                >Can help reduce the amount of explicit or potentially mature
						content shown. <br />(filter is not 100% accurate)</span
                ></label
                >

                <input
                    name="restricted"
                    id="restricted"
                    on:input={updatePrefsCookie}
                    type="checkbox"
                    bind:checked={$settings["search"]["Restricted"]}
                />
                <label
                    for="restricted"
                    class="switch"
                />
            </div>-->
        </section>
    </main>
{/if}

<style lang="scss">
    .input-container {
        min-width: 15ch !important;
        max-width: 32ch !important;
        width: 100%;
    }
    button {
        background: unset;
        all: unset;
        margin-top: 0.5rem;
        $link-color: rgb(245, 245, 245);
        color: rgb(245, 245, 245) !important;

        text-decoration: none;
        transition: color 0.2s;
        display: block;
        &.link {
            font-weight: 500;
        }
        &:active,
        &:focus,
        &:hover {
            background: transparent !important;
            -webkit-text-decoration: underline 0.001em solid;
            text-decoration: underline 0.001em solid;
            text-underline-offset: 0.001em;
            color: darken($link-color, 15%);
            outline: none;
        }
        &:hover {
            cursor: pointer;
        }
    }
    label {
        display: inline-flex;
        flex-direction: column;
        font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
        Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue",
        sans-serif;
        font-size: 1em;
        text-transform: none !important;
        font-variant: unset;
        gap: 0.125em;
        line-height: 1.4;
        :last-child {
            font-size: 0.875em;
            color: hsla(0, 0%, 100%, 0.7);
            line-height: 1.1;
        }

        @media screen and (min-width: 40em) {
            ~ :last-child {
                margin-left: auto;
            }
        }
    }

    section {
        display: flex;
        flex-direction: column;
        margin-block-end: 1em;

        &:not(:last-child) {
            border-bottom: 0.01em solid rgb(218 218 218 / 8.2%);
        }
    }

    .setting {
        display: inline-flex;
        color: inherit;
        vertical-align: top;
        gap: 1em;
        flex-direction: column;
        margin-block: 1em;

        &:first-of-type {
            margin-block-start: 0;
        }

        &:last-of-type {
            margin-block-end: 2em;
        }

        @media screen and (min-width: 40em) {
            align-items: center;
            flex-direction: row;
        }
    }

    .switch {
        position: relative;
        display: inline-flex;
        align-items: center;
        width: 3.8125em;
        height: 2em;
        cursor: pointer;
        overflow: hidden;
        background-color: rgb(109 109 109 / 35%);
        border-radius: 1.25rem;
        transition: background-color 0.3s;
    }

    .switch::after {
        --size: calc(2rem - (2px * 2));

        content: "";
        position: absolute;
        width: var(--size);
        height: var(--size);
        border-radius: 9999em;
        background-color: white;
        top: 50%;
        transform: translateY(-50%);
        left: 0.125em;
        transition: left 0.3s;
        box-shadow: 0 0 12px -3px rgb(0 0 0 / 38.4%);
    }

    [type="checkbox"]:checked + .switch::after {
        left: 2em;
    }

    [type="checkbox"]:checked + .switch {
        background-color: #00cd6a;
    }

    [type="checkbox"] {
        display: none;
    }
</style>
