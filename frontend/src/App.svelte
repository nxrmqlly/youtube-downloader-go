<script>
import { OpenBrowser, GetVideoInfo } from "../wailsjs/go/controllers/App.js";
import { onAccesibilityKeydown } from "./utils/accessibility.js";
import VideoModal from "./components/VideoModal.svelte";
import ThemeSwitcher from "./components/ThemeSwitcher.svelte";

let showModal = false;

let modalTitle;
let modalContent;

let url;

function openInBrowser(anchorUrl) {
	OpenBrowser(anchorUrl);
}

function openModal() {
	showModal = true;
}

function closeModal() {
	showModal = false;
}

function handleButtonClick() {
	if (!url) return;

	GetVideoInfo(url)
		.then((videoInfo) => {
			console.log(videoInfo);
			if (!videoInfo.Valid) {
				modalTitle = "Error fetching video!";
				modalContent = "Please check your URL and try again.";
			} else {
				modalTitle = videoInfo.title;
				modalContent = `Duration: ${videoInfo.duration}\nUploader: ${videoInfo.uploader}\nViews : ${videoInfo.view_count}`;
			}
			openModal();
		})
		.catch((error) => {
			console.log(error);
			modalTitle = "Error fetching video!";
			modalContent = "Please check your URL and try again.";
			openModal();
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
        <p>Works with <span 
            on:keydown={onAccesibilityKeydown} 
            on:click={() => openInBrowser("https://github.com/yt-dlp/yt-dlp")}
            >almost anything</span>. Blazingly fast, I guess.</p>
        <p>Made by <span
            on:keydown={onAccesibilityKeydown}
            on:click={() => openInBrowser("https://github.com/nxrmqlly")}
            >@nxrmqlly</span> with
            
            <span
                on:keydown={onAccesibilityKeydown}
                on:click={() => openInBrowser("https://youtu.be/dQw4w9WgXcQ")}
            >❤</span> using Go and Svelte</p>
    </footer>

    {#if showModal}
        <VideoModal
            title={modalTitle} 
            content={modalContent} 
            onClose={closeModal} 
        />
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
    footer span {
        color: var(--accent);
        text-decoration: none;
    }
    footer span:hover {
        cursor: pointer;
        text-decoration: underline;
    }

</style>
