<script>
import { GetVideoInfo } from "../wailsjs/go/controllers/App.js";
import Toast from "./components/Toast.svelte";
import VideoModal from "./components/VideoModal.svelte";
import ThemeSwitcher from "./components/ThemeSwitcher.svelte";
import Link from "./components/Link.svelte";

let url;
let errorMessage = "";
let videoResult;

let showToast = false;
let showModal = false;

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

function handleButtonClick() {
	if (!url) return;

	GetVideoInfo(url)
		.then((v) => {
			videoResult = v;
			openModal();
		})
		.catch((e) => {
			openToast();
			errorMessage = `${e}; Check your URL`;
		});
}
</script>


<main>
<ThemeSwitcher />

    <h1 class="title"><span>Youtube Downloader</span></h1>
    
    <div class="input-box" id="input">
        <input autocomplete="off" bind:value={url} class="input" id="name" type="text" placeholder="Enter your URL here..."/>
        <button class="btn" on:click={handleButtonClick}>►</button>
    </div>
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

<style>


    main {
        font-family: Arial, sans-serif;
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
    }
    .input {
        border: none;
        background-color: var(--bg-color);
        color: var(--fg-color);
        padding: 10px;
        width: 30rem;
        outline: none;
        flex: 1;
    }
    .btn {
        background-color: var(--accent);
        color: white;
        border: none;
        padding: 10px 15px;
        transition: background-color 0.3s;
    }
    .btn:hover {
        cursor: pointer;
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

</style>
