<script>
import { onAccesibilityKeydown } from "../utils/accessibility.js";
import { DownloadVideo } from "../../wailsjs/go/controllers/App.js";
import "@fortawesome/fontawesome-free/css/all.min.css";

export let onClose;
export let videoInfo;

let selectedOption;
let currentTab = "video"; // Default tab is video

function onDownloadButtonClick() {
	if (selectedOption) {
		DownloadVideo(videoInfo, selectedOption);

		onClose();
	}
}

function formatDuration(seconds) {
	const hrs = Math.floor(seconds / 3600);
	const mins = Math.floor((seconds % 3600) / 60);
	const secs = seconds % 60;

	return `${hrs > 0 ? `${hrs}h ` : ""}${mins > 0 ? `${mins}m ` : ""}${secs > 0 ? `${secs}s` : ""}`.trim();
}

const videoRes = [144, 240, 360, 480, 720, 1080, 1440, 2160, 4320, 8640];
const audioRes = -1;

const videoExtOptions = ["mp4", "webm", "mkv", "avi"];
const videoOptions = videoExtOptions.flatMap((ext) =>
	videoRes
		.filter((res) => res <= videoInfo.height) // Only map resolutions less than or equal to video height
		.map((res) => ({
			type: "video",
			ext: ext,
			res: res,
		})),
);

const audioExtOptions = ["mp3", "m4a", "opus", "webm", "wav"];
const audioOptions = audioExtOptions.map((ext) => ({
	type: "audio",
	ext: ext,
	res: audioRes,
}));

// Switch between video and audio tab and reset the selected option
function switchTab(tab) {
	currentTab = tab;
	selectedOption = ""; // Reset the selected option when tab changes
}
</script>

<style>
    .modal {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.7);
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .modal-content {
        background: var(--duo-bg-color);
        color: var(--fg-color);
        padding: 1.3rem;
        border-radius: 1rem;
        width: 45%;
        position: relative;
        display: flex;
        flex-direction: column;
    }
    h2 {
        margin: 0;
        margin-bottom: 0.5rem;
        font-size: 1.1rem;
        text-align: start;
    }
    .thumbnail {
        width: 100%;
        max-height: 25rem;
        border-radius: 0.5rem;
        margin-bottom: 0.5rem;
        object-fit: cover;
    }
    p {
        margin: 0;
        font-size: 14px;
    }
    .dropdown {
        margin-bottom: 1rem;
    }
    select {
        width: 100%;
        padding: 0.8rem;
        border: 0.1rem solid #888888;
        border-radius: 0.3rem;
        font-size: 0.9rem;
        color: var(--fg-color);
        background-color: var(--duo-bg-color);
        transition: border-color 0.2s;
    }
    select:focus {
        border-color: var(--accent);
        outline: none;
    }
    .download-btn {
        background-color: var(--confirm-color);
    }
    .close-btn {
        background-color: var(--deny-color);
    }
    .download-btn, .close-btn {
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 6px;
        cursor: pointer;
        font-size: 0.9rem;
        transition: background-color 0.3s ease;
        display: block;
        width: 10rem;
        text-align: center;
        margin: 0;
    }
    .download-btn:hover {
        background-color: var(--duo-confirm-color);
    }
    .close-btn:hover {
        background-color: var(--duo-deny-color);
    }

    .button-bar {
        display: flex;
        justify-content: end;
        align-items: center;
        gap: 1rem;
    }

    .video-desc {
        display: flex;
        margin: 0;
        margin-bottom: 1rem;
        justify-content:baseline;
        align-items: center;
        gap: 1rem;
    }

    .tab-bar {
        display: flex;
        justify-content: start;
        gap: 1rem;
        margin-bottom: 1rem;
    }
    
    .tab {
        cursor: pointer;
        padding: 0.5rem 1rem;
        border-bottom: 2px solid transparent;
    }
    
    .tab.active {
        border-color: var(--accent);
        font-weight: bold;
    }
</style>

<div class="modal" on:click={onClose} on:keydown={onAccesibilityKeydown}>
    <div class="modal-content" on:click|stopPropagation on:keydown={onAccesibilityKeydown}>
        <h2>{videoInfo.title}</h2>
        <div class="video-desc">
            <p><i class="fa-regular fa-user"></i> {videoInfo.uploader}</p>
            <p><i class="fa-regular fa-clock"></i> {formatDuration(videoInfo.duration)}</p>
            <p><i class="fa-regular fa-eye"></i> {videoInfo.view_count.toLocaleString()}</p>
        </div>
        <img class="thumbnail" src={videoInfo.thumbnail} alt="Video thumbnail" />

        <!-- Tab Bar -->
        <div class="tab-bar">
            <div class="tab {currentTab === 'video' ? 'active' : ''}" on:keydown={onAccesibilityKeydown} on:click={() => switchTab('video')}>
                Video
            </div>
            <div class="tab {currentTab === 'audio' ? 'active' : ''}" on:keydown={onAccesibilityKeydown} on:click={() => switchTab('audio')}>
                Audio
            </div>
        </div>

        <!-- Content based on selected tab -->
        <div class="dropdown">
        {#if currentTab === "video"}
                <select bind:value={selectedOption}>
                    <option value="" disabled selected>Select a video option</option>
                    {#each videoOptions as option (option.ext + option.res)}
                        <option value={option}>
                            {option.ext} — {option.res}p {videoInfo.fps > 0 ? ` @ ${videoInfo.fps}fps` : ""}
                        </option>
                    {/each}
                </select>
   
        {:else if currentTab === "audio"}
          
                <select bind:value={selectedOption}>
                    <option value="" disabled selected>Select an audio option</option>
                    {#each audioOptions as option (option.ext + option.res)}
                        <option value={option}>
                            {option.ext} — Audio only
                        </option>
                    {/each}
                </select>
                {/if}
            </div>

        <div class="button-bar">
            <button class="download-btn" on:click={onDownloadButtonClick}>Download</button>
            <button class="close-btn" on:click={onClose}>Close</button>
        </div>
    </div>
</div>
 