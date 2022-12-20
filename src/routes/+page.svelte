<script>
    import { onMount } from "svelte";

    import Header from "./header.svelte";

    import Highlight from 'svelte-highlight';
    import bash from 'svelte-highlight/languages/bash';
    import agate from 'svelte-highlight/styles/agate';

    let Carousel;
    onMount(async () => {
        // Perform dynamic import of svelte-carousel because it breaks
        // if imported during SSR
        Carousel = (await import('svelte-carousel')).default;
    })

    let images = [
        {
            path: "/lure-arch.gif",
            caption: "Recorded on Arch Linux with LURE commit 81013ce",
            alt: "GIF depicting an Arch Linux terminal with LURE installing a package"
        },
        {
            path: "/lure-debian.gif",
            caption: "Recorded on Debian Sid with LURE commit 81013ce",
            alt: "GIF depicting a Debian terminal with LURE installing a package"
        },
        {
            path: "/lure-fedora.gif",
            caption: "Recorded on Fedora 36 with LURE commit 81013ce",
            alt: "GIF depicting a Fedora terminal with LURE installing a package"
        },
        {
            path: "/lure-alpine.gif",
            caption: "Recorded on Alpine Linux 3.16 with commit 81013ce",
            alt: "GIF depicting an Alpine Linux terminal with LURE installing a package"
        },
    ];
</script>

<svelte:head>
    {@html agate}
</svelte:head>

<Header/>

<section class="container">
    <p class="title">LURE</p>
    <p class="subtitle">The AUR missing from most Linux distributions</p>
    <hr>
    <div class="columns">
        <div class="column">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">Why should I use it?</p>
                </div>
                <div class="card-content">
                    LURE allows you to install your favorite software that may not be
                    popular enough to be placed into your distro's repos, while preserving
                    all the benefits of installing from repos, such as updates, easy
                    uninstalling, etc. It also allows developers to provide a single place
                    for their users to install their software.
                </div>
            </div>
        </div>
        <div class="column">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">How does it work?</p>
                </div>
                <div class="card-content">
                    LURE works by abstracting package formats and package managers. It
                    builds native packages for your native package manager using AUR 
                    PKGBUILD-like bash scripts, then installs them using that package manager.
                    This means that once LURE has successfully installed a package, it acts
                    just like any other package and can be managed without LURE.
                </div>
            </div>
        </div>
    </div>
    <p class="title">Installation</p>
    <p>LURE can easily be installed by running its install script:</p>
    <Highlight language={bash} code='curl https://www.arsenm.dev/lure.sh | bash'/>
    <p>
        It's also available on the AUR as <a href="https://aur.archlinux.org/packages/lure-bin"><code>lure-bin</code></a> and distro
        packages are provided at the <a href="https://gitea.arsenm.dev/Arsen6331/lure/releases/latest">latest Gitea release</a>.
    </p>
    <br>
    <p class="title">Examples</p>

    <!-- Use dynamically imported Carousel module -->
    <svelte:component this={Carousel}>
        {#each images as image}
            <figure class="has-text-centered">
                <img src={image.path} alt={image.alt ?? "GIF depicting LURE installing a package"}>
                <figcaption>{image.caption ?? ""}</figcaption>
            </figure>
        {/each}
    </svelte:component>

</section>
