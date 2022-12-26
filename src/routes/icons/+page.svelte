<script>
    import Header from "../header.svelte";
    import Footer from "../footer.svelte";
</script>

<svelte:head>
    <title>FAQ | LURE Web</title>
    <meta name="description" content="LURE Icons">
</svelte:head>

<Header/>

<div class="container">
    <p class="title">Icons</p>
    <br>
    <div class="columns">
        <div class="column ">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">
                        Without Text
                    </p>
                </div>
                <div class="card-content has-text-centered">
                    <figure class="image is-16by9">
                        <img src="/lure-no-text.svg" alt="LURE icon without text">
                    </figure>
                </div>
                <footer class="card-footer">
                    <a class="card-footer-item" href="/lure-no-text.svg" download>Download</a>
                </footer>
            </div>
        </div>
        <div class="column ">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">
                        With Text
                    </p>
                </div>
                <div class="card-content has-text-centered">
                    <figure class="image is-16by9">
                        <img src="/lure-text.svg" alt="LURE icon with text">
                    </figure>
                </div>
                <footer class="card-footer">
                    <a class="card-footer-item" href="/lure-text.svg" download>Download</a>
                </footer>
            </div>
        </div>
        <div class="column">
            <div class="card">
                <div class="card-header">
                    <p class="card-header-title">
                        White With Text
                    </p>
                </div>
                <div class="card-content has-text-centered">
                    <figure class="image is-16by9">
                        <img src="/lure-text-white.svg" alt="White LURE icon with white text">
                    </figure>
                </div>
                <footer class="card-footer">
                    <a class="card-footer-item" href="/lure-text-white.svg" download>Download</a>
                </footer>
            </div>
        </div>
    </div>
    <p class="has-text-weight-bold">Note: The icons on this page are licensed under CC-BY-NC-SA 4.0 unless otherwise specified</p>
</div>

<Footer/>