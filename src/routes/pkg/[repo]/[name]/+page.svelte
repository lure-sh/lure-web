<script>
    import { page } from '$app/stores';
    import { client } from 'twirpscript';
    import { GetPkg } from '$lib/lure.pb';
    import { LURE_WEB_API_URL } from '$env/static/public';
    import { DoubleBounce } from 'svelte-loading-spinners';

    import Header from "../../../header.svelte";

    import Icon from '@iconify/svelte';

    client.baseURL = LURE_WEB_API_URL
    client.prefix = ""

    function fullVer(pkg) {
        let ver = `${pkg.version}-${pkg.release}`;
        if (pkg.epoch != 0) {
            ver = `${pkg.epoch}:${ver}`
        }
        return ver
    }

    function objToMap(o) {
        return new Map(Object.entries(o))
    }
</script>

<svelte:head>
    <title>{$page.params.name} Package | LURE Web</title>
</svelte:head>

<Header/>

<section class="container">
    <a href="/pkgs"><Icon icon="material-symbols:arrow-back-rounded" inline=true/> Back</a>
    {#await GetPkg({name: $page.params.name, repository: $page.params.repo})}
        <center><DoubleBounce color="#375a7f" /></center>
    {:then pkg}
        <p class="title">{pkg.name}</p>
        <p class="subtitle mb-3">{fullVer(pkg)}</p>
        <a href="/pkg/{pkg.repository}/{pkg.name}/lure.sh" class="button is-primary mb-5">View lure.sh</a>
        <div class="box">
            <table class="table is-hoverable is-fullwidth">
                <tbody>
                    {#if pkg.description != ""}
                        <tr>
                            <th>Description:</th>
                            <td>{pkg.description}</td>
                        </tr>
                    {/if}
                    {#if pkg.homepage != ""}
                    <tr>
                        <th>Homepage:</th>
                        <td><a href="{pkg.homepage}">{pkg.homepage}</a></td>
                    </tr>
                    {/if}
                    {#if pkg.maintainer != ""}
                    <tr>
                        <th>Maintainer:</th>
                        <td>{pkg.maintainer}</td>
                    </tr>
                    {/if}
                    {#if pkg.licenses.length != 0}
                    <tr>
                        <th>Licenses:</th>
                        <td>
                            {#each pkg.licenses as license, index}
                                <a href="https://spdx.org/licenses/{license}.html">{license}</a>{#if index+1 < pkg.licenses.length},&nbsp;{/if}
                            {/each}
                        </td>
                    </tr>
                    {/if}
                    {#if pkg.architectures.length != 0}
                    <tr>
                        <th>Architectures:</th>
                        <td>{pkg.architectures.join(', ')}</td>
                    </tr>
                    {/if}
                    {#if pkg.conflicts.length != 0}
                    <tr>
                        <th>Conflicts:</th>
                        <td>{pkg.conflicts.join(', ')}</td>
                    </tr>
                    {/if}
                    {#if pkg.provides.length != 0}
                    <tr>
                        <th>Provides:</th>
                        <td>{pkg.provides.join(', ')}</td>
                    </tr>
                    {/if}
                    {#each [...objToMap(pkg.depends)] as [override, pkgList]}
                    <tr>    
                        <th>Depends ({override == "" ? "default" : override}):</th>
                        <td>{pkgList.entries.join(', ')}</td>
                    </tr>
                    {/each}
                    {#each [...objToMap(pkg.buildDepends)] as [override, pkgList]}
                    <tr>    
                        <th>Build Depends ({override != "" ? override : "default"}):</th>
                        <td>{pkgList.entries.join(', ')}</td>
                    </tr>
                    {/each}
                    <tr>
                        <th>Repository:</th>
                        <td>{pkg.repository}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    {:catch err}
        <div class="notification is-danger my-3">
            Error: {err.msg}
        </div>
    {/await}
</section>