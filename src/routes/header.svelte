<script>
    import { page } from '$app/stores';

    import Icon from '@iconify/svelte';

    function isActive(path) {
        return $page.route.id == path ? 'is-active' : ''
    }

    let navbarIsActive = false;
</script>

<nav class="navbar mb-5">
    <div class="navbar-brand">
        <a class="navbar-item" href="/">LURE Web</a>
        <button class="navbar-burger {navbarIsActive ? 'is-active' : ''}" aria-label="menu" on:click={() => {navbarIsActive = !navbarIsActive}}>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
        </button>
    </div>

    <div id="main-navbar" class="navbar-menu {navbarIsActive ? 'is-active' : ''}">
        <div class="navbar-end">
            <a class="navbar-item {isActive('/')}" href="/"><Icon icon="material-symbols:house" inline=true/> Home</a>
            <a class="navbar-item {isActive('/pkgs')}" href="/pkgs"><Icon icon="mdi:package-variant-closed" inline=true/>&nbsp;Packages</a>
            <a class="navbar-item {isActive('/faq')}" href="/faq"><Icon icon="mdi:question-mark-circle" inline=true/>&nbsp;FAQ</a>
        </div>
    </div>
</nav>