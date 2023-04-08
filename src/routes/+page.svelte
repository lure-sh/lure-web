<script>
    import { onMount } from "svelte";

    import Header from "./header.svelte";
    import Footer from "./footer.svelte";

    import Highlight from 'svelte-highlight';
    import bash from 'svelte-highlight/languages/bash';
    import agate from 'svelte-highlight/styles/agate';

    let Carousel, Player, Video, DefaultUi;
    onMount(async () => {
        // Perform dynamic import of svelte-carousel because it breaks
        // if imported during SSR
        Carousel = (await import('svelte-carousel')).default;
        ({Player, Video, DefaultUi} = await import('@vime/svelte'));
    })

    let images = [
        {
            path: "/lure-arch.webm",
            caption: "Recorded on Arch Linux with LURE commit 81013ce",
        },
        {
            path: "/lure-debian.webm",
            caption: "Recorded on Debian Sid with LURE commit 81013ce",
        },
        {
            path: "/lure-fedora.webm",
            caption: "Recorded on Fedora 36 with LURE commit 81013ce",
        },
        {
            path: "/lure-alpine.webm",
            caption: "Recorded on Alpine Linux 3.16 with commit 81013ce",
        },
    ];
</script>

<svelte:head>
    <title>Home | LURE Web</title>
    <meta name="description" content="LURE Web home page">
    {@html agate}
</svelte:head>

<Header/>

<section class="container">
    <p class="title">LURE</p>
    <p class="subtitle">The user repo missing from most Linux distros</p>
    <hr>
    <div class="columns">
        <div class="column">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">Why should I use it?</p>
                </div>
                <div class="card-content">
                    LURE allows users to install software that may not be widely distributed
                    through official repositories, while still maintaining the convenience
                    of installation through repository sources. This includes features such
                    as updates and simple uninstallation. Additionally, LURE provides developers
                    with a central location for all their users to use to install their software.
                </div>
            </div>
        </div>
        <div class="column">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">How does it work?</p>
                </div>
                <div class="card-content">
                    LURE operates by abstracting package formats and package managers, enabling
                    the creation and installation of native packages automatically built from
                    PKGBUILD-like bash scripts, using the package manager already present on the system.
                    As a result, packages installed through LURE can be managed like any other package,
                    without the need for additional intervention from LURE for most operations.
                </div>
            </div>
        </div>
    </div>
    <p class="title">Installation</p>
    <p>LURE can easily be installed by running its install script:</p>
    <Highlight language={bash} code='curl https://www.arsenm.dev/lure.sh | bash'/>
    <p>
        It's also available on the AUR as <a href="https://aur.archlinux.org/packages/linux-user-repository-bin"><code>linux-user-repository-bin</code></a> and distro
        packages are provided at the <a href="https://gitea.arsenm.dev/Arsen6331/lure/releases/latest">latest Gitea release</a>.
    </p>
    <br>
    <p class="title">Examples</p>

    <!-- Use dynamically imported Carousel module -->
    <svelte:component this={Carousel}>
        {#each images as image}
            <div class="has-text-centered">
                <figure class="terminal-player mx-auto">
                    <svelte:component this={Player}>
                        <svelte:component this={Video}>
                            <source data-src="{image.path}" type="video/webm">
                        </svelte:component>
                        <svelte:component this={DefaultUi}></svelte:component>
                    </svelte:component>
                    <figcaption>{image.caption ?? ""}</figcaption>
                </figure>
            </div>
        {/each}
    </svelte:component>

</section>

<Footer/>
