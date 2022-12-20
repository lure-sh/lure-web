<script>
    import Header from "../header.svelte";
</script>

<Header/>

<div class="container">
    <p class="title">FAQ</p>
    <br>
    <p class="subtitle" id='distro-support'>Why isn't my distro supported?</p>
    <p>
        In order to support a distro, LURE has to be able to talk to its package manager
        as well as build packages for it. Talking to the package manager is relatively simple,
        LURE just needs to be told which commands to run, but the package formats are much more
        complex. LURE uses <a href="https://github.com/goreleaser/nfpm">nFPM</a> to handle
        package formats, and nFPM supports only <code>deb</code>, <code>rpm</code>, <code>apk</code>,
        and <code>archlinux</code> package formats. When I started LURE, nFPM didn't have support for
        archlinux packages, so I added it in a PR. Despite Arch having a very simple package format, it took
        over 1k lines of code to implement, so supporting a distro with a different format, like
        <code>xbps</code> for Void Linux, is very complicated, and may take some time.
    </p>
    <hr>
    <p class="subtitle" id='flatpak-snap-appimage'>Why use LURE instead of Flatpak, Snap, or AppImage?</p>
    <p>
        LURE is not meant to solve the same problems as Flatpak, Snap, and AppImage. These are containerized
        package formats that allow the creation of a single package that works on all distros. This cross-platform package
        contains the program as well as everything else needed for that program to run. It works using
        something called a container, which can create issues where programs are slow to start, don't follow 
        your system's settings, or can't access certain parts of your system. If you want most programs to work every
        time, or you're running an old distro with old packages, you should use containzerized formats. 
        LURE, on the other hand, does not use containers. It builds the program from source automatically and
        installs it. It also doesn't have its own package format. Instead, it uses the same one as the distro
        it's running on, making LURE packages act the same way as your distro's packages. This means that
        unlike Snap and Flatpak, LURE is not a package manager. It simply uses your distro's package manager,
        which allows you to manage the packages installed by LURE even without it installed.
    </p>
    <br>
    <p>
        However, LURE also has some downsides. Because of the fact that it builds programs automatically from source,
        certain packages, especially git packages that retrieve the latest code from git, may not work on older distros,
        or distros like Debian which have old packages. Please consider your use case and whether such downsides are
        acceptable before using LURE. Also note that, like the AUR, all packages are user-submitted and are not vetted,
        so while unlikely, they may contain malicious code. It is the responsibility of the user to read the build script
        and ensure this is not the case. Malicious packages should be reported by opening an issue on the git repo containing them.
    </p>
    <hr>
    <p class="subtitle" id='adding-packages'>How do I add my own package to LURE?</p>
    <p>
        Take a look at the documentation for <a href="https://github.com/Arsen6331/lure/blob/master/docs/build-scripts.md">Build Scripts</a>
        and <a href="https://github.com/Arsen6331/lure/blob/master/docs/adding-packages.md">Adding Packages to LURE's repo</a>.
    </p>
</div>