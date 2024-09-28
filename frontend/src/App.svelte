<script>
  import { DownloadVideo } from '../wailsjs/go/main/App.js'
  let url
  let text = ""

  function handleButtonClick() {
    DownloadVideo(url).then((result) => {
      text = result
    })
  }
</script>


<main>
  
  <div class="input-box" id="input">
    <p>{text}</p>
    <input autocomplete="off" bind:value={url} class="input" id="name" type="text"/>
    <button class="btn" on:click={handleButtonClick}>Download</button>
  </div>
</main>


<style>
</style>
