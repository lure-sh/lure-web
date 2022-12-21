<script>
    import { page } from '$app/stores';
    import { client } from 'twirpscript';
    import { GetBuildScript } from '$lib/lure.pb';
    import { LURE_WEB_API_URL } from '$env/static/public';
    import { DoubleBounce } from 'svelte-loading-spinners';

    import Highlight from 'svelte-highlight';
    import bash from 'svelte-highlight/languages/bash';
    import agate from 'svelte-highlight/styles/agate';

    import Header from "../../../../header.svelte";

    import Icon from '@iconify/svelte';

    client.baseURL = LURE_WEB_API_URL
    client.prefix = ""
</script>

<svelte:head>
    <title>{$page.params.name} Build Script | LURE Web</title>
    <meta name="description" content="The build script for the {$page.params.name} LURE package.">
    {@html agate}
</svelte:head>

<Header/>

<section class="container">
    <a href="."><Icon icon="material-symbols:arrow-back-rounded" inline=true/> Back</a>
    {#await GetBuildScript({name: $page.params.name, repository: $page.params.repo})}
        <center><DoubleBounce color="#375a7f" /></center>
    {:then resp} 
        <p class="title">Build Script</p>
        <p class="subtitle">{$page.params.repo} / {$page.params.name}</p>
        <div class="box">
            <Highlight language={bash} code={resp.script}/>
        </div>
    {:catch err}
        <div class="notification is-danger my-3">
            Error: {err.msg}
        </div>
    {/await}
</section>

