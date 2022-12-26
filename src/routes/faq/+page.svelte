<script>
    import Header from "../header.svelte";
    import Footer from "../footer.svelte";
</script>

<svelte:head>
    <title>FAQ | LURE Web</title>
    <meta name="description" content="Frequently Asked Questions">
</svelte:head>

<Header/>

<div class="container">
    <p class="title">FAQ</p>
    <br>
    <p class="subtitle" id='distro-support'>Why isn't my distro supported?</p>
    <p>
        To support a distribution, LURE must be able to communicate with its package manager and create packages for it.
        Communicating with the package manager is relatively straightforward, as LURE simply needs to be provided with the necessary
        commands. However, the package formats are more complex. LURE uses <a href="https://github.com/goreleaser/nfpm">nFPM</a> to
        handle package formats, and nFPM currently supports only deb, rpm, apk, and archlinux package formats. When developing LURE,
        support for archlinux packages was not available in nFPM, so I added it through a pull request. Despite the simplicity of the
        archlinux package format, implementing support required over 1,000 lines of code. As a result, supporting a distribution with
        a different package format, such as xbps for Void Linux, is very complex and time-consuming.
    </p>
    <hr>
    <p class="subtitle" id='flatpak-snap-appimage'>Why use LURE instead of Flatpak, Snap, or AppImage?</p>
    <p>
        LURE is not intended to address the same issues as Flatpak, Snap, and AppImage. These are containerized package formats that enable
        the creation of a single package that can be used on all distributions. This cross-platform package contains the program and everything
        else necessary for it to run, and it relies on containers to achieve this compatibility. However, containers can sometimes cause programs
        to start slowly, fail to adhere to system settings, or be unable to access certain parts of the system. If you need most programs to
        function consistently or you are using an older distribution with outdated packages, containerized formats may be the best choice. In contrast,
        LURE does not use containers. It builds the program from source and installs it automatically. It also does not have its own package format.
        Instead, it uses the same format as the distribution it is running on, so LURE packages behave like the distribution's native packages.
        This means that unlike Snap and Flatpak, LURE is not a package manager; it simply uses the distribution's package manager, allowing you to
        manage the packages installed through LURE even when LURE is not installed.
    </p>
    <br>
    <p>
        However, LURE also has some drawbacks. Since it builds programs from source, certain packages, particularly git packages that
        retrieve the latest code from git, may not work on older distributions or distributions like Debian that have outdated packages.
        Please consider your specific needs and whether these downsides are acceptable before using LURE. Additionally, similar to the AUR,
        all packages are user-submitted and not vetted, so while it is unlikely, they may contain malicious code. It is the responsibility
        of the user to review the build script to ensure this is not the case. If you come across a malicious package, please report it by
        opening an issue on the git repository containing it.
    </p>
    <hr>
    <p class="subtitle" id="handling-dependencies">How does LURE handle dependencies across distros?</p>
    <p>
        LURE manages dependencies across distributions by offering distro overrides, in which package maintainers can specify different
        variables and functions for each distribution. The most specific override is given precedence. After the overrides are resolved,
        LURE compares the resulting list of dependencies with the packages installed on the system and filters out any that are already
        installed. For the remaining dependencies, LURE checks its own repositories to see if each package is available there. If it is,
        LURE installs it from its repositories. If the package is not found in any repository, LURE passes the dependency on to the package
        manager, which handles dependency resolution and installation.
    </p>
    <hr>
    <p class="subtitle" id="testing">How can one test a LURE package to ensure it works?</p>
    <p>
        Docker is recommended for testing LURE packages on different distros. It provides a clean image of any distribution, which is very
        useful for testing as it can help catch issues that might not manifest themselves on your system. Eventually, an automated
        docker-based testing tool is planned, but in the meantime, this will need to be done manually for each distribution you're planning
        to support. To find package names for each distribution, you can use <a href="https://repology.org">repology.org</a> and
        <a href="https://pkgs.org">pkgs.org</a>. These websites maintain comprehensive databases of package repositories for various distributions.
    </p>
    <hr>
    <p class="subtitle" id='adding-packages'>How do I add my own package to LURE?</p>
    <p>
        To add your own package, please refer to the <a href="https://github.com/Arsen6331/lure/blob/master/docs/packages">package documentation</a> provided by LURE.
    </p>
</div>

<Footer/>