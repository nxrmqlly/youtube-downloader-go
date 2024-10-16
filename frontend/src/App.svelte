<script>
  import { DownloadVideo } from '../wailsjs/go/controllers/App.js'
  let url
  let text = ""
  /*
  type DownloadOptions struct {
    Path         string
    DownloadType string // "video" or "audio"
    FileType     string // e.g., "mp4", "mp3"
    Resolution   string // e.g., "720p", "1080p"
    Quality      string // e.g., "best", "worst"
  }
  */
 const downloadOptions = {
   URL: url,
   DownloadType: "video",
   FileType: "mp4",
   Resolution: "720p",
   Quality: "best",
 }

		function handleButtonClick() {
			DownloadVideo(downloadOptions).then((result) => {
				text = result;
			});
		}
</script>


<main>
  
  <div class="input-box" id="input">
    <p>{text}</p>
    <input autocomplete="off" bind:value={url} class="input" id="name" type="text"/>
    <button class="btn" on:click={handleButtonClick}>Get</button>
  </div>
  
</main>


<style>
</style>
