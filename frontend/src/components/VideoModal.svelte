<script>
import { onAccesibilityKeydown } from "../utils/accessibility.js";
import { DownloadVideo } from "../../wailsjs/go/controllers/App.js";

export let onClose;
export let videoInfo;

function onDownloadButtonClick() {
	if (selectedOption) {
		console.log(selectedOption);

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

const videoRes = [144, 240, 360, 480, 720, 1080, 1440, 2160, 4320];
const audioRes = -1;

const videoExtOptions = ["mp4", "webm", "mkv", "avi"];
const videoOptions = videoExtOptions.flatMap((ext) =>
	videoRes.map((res) => ({
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

const allOptions = [...videoOptions, ...audioOptions];

let selectedOption;
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
        background: #f6f3f4;
        padding: 25px;
        border-radius: 12px;
        width: 450px;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
        position: relative;
    }
    h2 {
        margin: 0 0 10px 0;
        font-size: 24px;
    }
    .thumbnail {
        width: 100%;
        max-height: 250px;
        border-radius: 8px;
        margin-bottom: 15px;
        object-fit: cover;
    }
    p {
        margin: 5px 0;
        font-size: 14px;
    }
    .dropdown {
        margin: 15px 0;
    }
    select {
        width: 100%;
        padding: 10px;
        border: 1px solid #ddd;
        border-radius: 6px;
        font-size: 14px;
        background-color: #f9f9f9;
        transition: border-color 0.2s;
    }
    select:focus {
        border-color: var(--accent);
        outline: none;
    }
    button.download-btn {
        background-color: var(--confirm-color);
        color: white;
        padding: 10px 20px;
        border: none;
        border-radius: 6px;
        cursor: pointer;
        font-size: 14px;
        transition: background-color 0.3s ease;
        display: block;
        width: 100%;
        text-align: center;
        margin-top: 10px;
    }
    button.download-btn:hover {
        background-color: var(--duo-confirm-color);
    }
    button.close-btn {
        position: absolute;
        top: 0.5rem;
        right: 0.5rem;
        background: transparent;
        border: none;
        font-size: 2rem;
        cursor: pointer;
        color: #888;
    }
    button.close-btn:hover {
        color: #333;
        
    }
</style>

<div class="modal" on:click={onClose} on:keydown={onAccesibilityKeydown}>
    <div class="modal-content" on:click|stopPropagation on:keydown={onAccesibilityKeydown}>
        <button class="close-btn" on:click={onClose}>&times;</button>

        <h2>{videoInfo.title}</h2>
        <img class="thumbnail" src={videoInfo.thumbnail} alt="Video thumbnail" />

        <p><strong>Uploader:</strong> {videoInfo.uploader}</p>
        <p><strong>Duration:</strong> {formatDuration(videoInfo.duration)}</p>
        <p><strong>Views:</strong> {videoInfo.view_count.toLocaleString()}</p>

        <div class="dropdown">
            <label for="formatSelect"><strong>Select Format:</strong></label>
            <select bind:value={selectedOption}>
                <option value="" disabled>Select an option</option>
                {#each allOptions as option (option.ext + option.res)}
                    <option value={option}>{option.ext} - {option.res === -1 ? "Audio" : option.res}{option.type === "audio" ?  " only" : ` @ ${videoInfo.fps}fps`}</option>
                {/each}
            </select>
        </div>

        <button class="download-btn" on:click={onDownloadButtonClick}>Download</button>
    </div>
</div>
