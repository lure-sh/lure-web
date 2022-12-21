<script>
    import { client } from 'twirpscript';
    import { Search } from '$lib/lure.pb';
    import { searchReq } from '../../stores.js';
    import { LURE_WEB_API_URL } from '$env/static/public';

    import Header from "../header.svelte";
    import { DoubleBounce } from 'svelte-loading-spinners';
    import Icon from '@iconify/svelte';

    client.baseURL = LURE_WEB_API_URL
    client.prefix = ""

    let searchPromise = Search(
        $searchReq ?? {
            query: "",
            limit: BigInt(0),
            sortBy: "UNSORTED",
            filterType: "NO_FILTER",
        }
    );

    function onSubmit(e) {
        let formData = new FormData(e.target);
        let filterValue = formData.get('filter-value');
        let req = {
            query: formData.get('query'),
            limit: BigInt(0),
            sortBy: formData.get('sort-by'),
            filterType: formData.get('filter'),
            filterValue: filterValue == '' ? null : filterValue,
        }
        searchReq.set(req)
        searchPromise = Search(req);
    }

    function onFilterChange(e) {
        let filterVal = document.getElementById('filter-val');

        if (e.target.value != "NO_FILTER") {
            filterVal.classList.remove('is-hidden');
        } else {
            filterVal.classList.add('is-hidden');
        }
    }
</script>

<svelte:head>
    <title>Package Search | LURE Web</title>
</svelte:head>

<Header/>

<section class="container">
    <p class="title">Package Search</p>
    <form on:submit|preventDefault={onSubmit}>
        <div class="field">
            <div class="control">
                <input class="input" type="text" name="query" value="" placeholder="Query">
            </div>
        </div>
        <div class="columns">
            <div class="column">
                <div class="field">
                    <div class="control select is-fullwidth">
                        <select name="sort-by">
                            <option value="UNSORTED" selected>Unsorted</option>
                            <option value="NAME">Sort by Name</option>
                            <option value="VERSION">Sort by Version</option>
                            <option value="REPOSITORY">Sort by Repository</option>
                        </select>
                    </div>
                </div>
            </div>
            <div class="column">
                <div class="field">
                    <div class="control select is-fullwidth">
                        <select name="filter" on:change={onFilterChange}>
                            <option value="NO_FILTER" selected>No Filter</option>
                            <option value="IN_REPOSITORY">In Repo</option>
                            <option value="SUPPORTS_ARCH">Supports Architecture</option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
        <div class="field">
            <div class="control">
                <input id="filter-val" class="input is-hidden" type="text" value="" name="filter-value" placeholder="Filter Value">
            </div>
        </div>
        <input type="submit" class="button is-fullwidth" value="Search">
    </form>
    
    <hr>

    {#await searchPromise}
        <center><DoubleBounce color="#375a7f" /></center>
    {:then resp}
        {#each resp.packages as pkg}
            <div class="card my-5">
                <header class="card-header">
                    <p class="card-header-title">{pkg.repository} / {pkg.name}</p>
                    <div class="card-header-icon">{pkg.version}</div>
                </header>
                
                <div class="card-content">
                    {pkg.description}
                </div>

                <footer class="card-footer">
                    <a class="card-footer-item" href="/pkg/{pkg.repository}/{pkg.name}">More info <Icon icon="material-symbols:arrow-forward-rounded" inline=true/></a>
                </footer>
            </div>
        {/each}
        {#if resp.packages.length === 0}
            <p class="subtitle has-text-centered has-text-danger my-5 is-fullwidth">No Results</p>
        {/if}
    {:catch err}
        <div class="notification is-danger my-3">
            Error: {err.msg}
        </div>
    {/await}

</section>