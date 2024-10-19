<script>
import { GetVideoInfo } from "../wailsjs/go/controllers/App.js";
import Toast from "./components/Toast.svelte";
import "@fortawesome/fontawesome-free/css/all.min.css";
import VideoModal from "./components/VideoModal.svelte";
import ThemeSwitcher from "./components/ThemeSwitcher.svelte";
import Link from "./components/Link.svelte";

let url = "";
let errorMessage = "";
let videoResult;
let loading = false;

let showToast = false;
let showModal = false;

function isValidURL(string) {
	const regex = /^(https?:\/\/)?([\da-z.-]+)\.([a-z.]{2,6})([\/\w .-]*)*\/?$/;
	return regex.test(string);
}

function openModal() {
	showModal = true;
}
function closeModal() {
	showModal = false;
}
function openToast() {
	showToast = true;
}
function closeToast() {
	showToast = false;
}

function handleSubmit(e) {
	e.preventDefault();
	handleButtonClick();
}

function handleButtonClick() {
	if (!url) return;
	if (!isValidURL(url)) {
		errorMessage = "Error: Invalid URL";
		openToast();
		return;
	}

	loading = true;

	GetVideoInfo(url)
		.then((v) => {
			videoResult = v;
			openModal();
		})
		.catch((e) => {
			openToast();
			errorMessage = e.toString();
		})
		.finally(() => {
			loading = false;
		});
}
</script>

<style>


    main {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100dvh;
        background-color: var(--bg-color);

    }
    .title {
        color: var(--fg-color);
    }

    .title span {
        display: inline-block;
    }

    .title span::after{
        content: '';
        height: 0.1em;
        background: var(--accent);
        border-radius: 0.1em;
        display: block
    }

    .input-box {
        margin-top: 1em;
        display: flex;
        border: 0.1em solid var(--accent);
        border-radius: 0.1em;
        overflow: hidden;
        align-items: center;
        /* padding: 0.5rem; */
    }
    .input {
        border: none;
        background-color: var(--bg-color);
        color: var(--fg-color);
        width: 30rem;
        outline: none;
        flex: 1;
    }
    .btn {
        margin: 0;
        top: 0;
        right: 0;
        background-color: var(--accent);
        color: white;
        border: none;
        padding: 0.5rem 1.25rem;
        /* width: 4rem; */
        /* height: 0.5rem; */

        transition: background-color 0.3s;
    }
    .btn:hover {
        cursor: pointer;
        background-color: var(--duo-accent);
    }
    .btn:disabled{
        cursor: not-allowed;
        background-color: var(--duo-accent);
    }
    footer {
        user-select: none;
        position: absolute;
        bottom: 0;
        width: 100%;
        text-align: center;
        font-size: 12px;
        color: #787878;
    }
    :global([data-ref="accent-anchor"]) {
        color: var(--accent);
        text-decoration: none;
    }
    :global([data-ref="accent-anchor"]):hover {
        cursor: pointer;
        text-decoration: underline;
    }
    .mag {
        color: var(--accent);
        padding: 0 0.75rem;
    }
    .btn-icon {
        width: 1em;
        height: 1em;
    }

</style>

<main>
<ThemeSwitcher />

    <h1 class="title"><span>Youtube Downloader</span></h1>
    
    <form class="input-box" id="input" on:submit={handleSubmit}>
        <i class="mag fa-solid fa-magnifying-glass"></i>
        <input autocomplete="off" bind:value={url} class="input" id="name" type="text" placeholder="Enter your URL here..."/>
        <button class="btn" type="submit" on:click={handleButtonClick} disabled={loading || !url}>

                <i class="btn-icon fa-solid fa-chevron-right"></i>

        </button>
    </form>
    <footer>
        <p>Works with <Link ref="accent-anchor" href="https://github.com/yt-dlp/yt-dlp/blob/master/supportedsites.md">almost anything</Link>. Blazingly fast, I guess.</p>
        <p>Made by <Link ref="accent-anchor" href="https://github.com/nxrmqlly">@nxrmqlly</Link> with <Link ref="accent-anchor" href="https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUIcmlja3JvbGw%3D">❤</Link> using Go and Svelte</p>
    </footer>

    {#if showModal}
        <VideoModal
        videoInfo={videoResult}
            onClose={closeModal} 
        />
    {/if}

    {#if showToast}
        <Toast message={errorMessage} onClose={closeToast} />   
    {/if}
  
</main>