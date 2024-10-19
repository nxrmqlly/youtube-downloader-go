<script>
import { onMount } from "svelte";
import { onAccesibilityKeydown } from "../utils/accessibility.js";
import "@fortawesome/fontawesome-free/css/all.min.css";

let isDarkMode = window.matchMedia("(prefers-color-scheme: dark)").matches;
window
	.matchMedia("(prefers-color-scheme: dark)")
	.addEventListener("change", (e) => {
		isDarkMode = e.matches;
		toggleTheme();
	});
console.log(isDarkMode);

function toggleTheme() {
	isDarkMode = !isDarkMode;
	if (!isDarkMode) {
		document.documentElement.style.setProperty("--bg-color", "#100a0d");
		document.documentElement.style.setProperty("--duo-bg-color", "#212021");
		document.documentElement.style.setProperty("--fg-color", "#f6f3f4");
	} else {
		document.documentElement.style.setProperty("--bg-color", "#f6f3f4");
		document.documentElement.style.setProperty("--duo-bg-color", "#f6f3f4");
		document.documentElement.style.setProperty("--fg-color", "#100a0d");
	}
}

// sset default theme
onMount(() => {
	toggleTheme();
});
</script>


<style>
    .switcher {
        position: fixed;
        top: 1rem;
        right: 1rem;
        color: var(--fg-color);
        border-radius: 50%;        
        width: 4rem; 
        height: 4rem; 
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        font-size: 1.5rem;
    }

    .switcher:hover {
        opacity: 0.8;
    }
</style>

<div class="switcher" on:click={toggleTheme} on:keydown={onAccesibilityKeydown}>
    {#if isDarkMode}
        <i class="fas fa-regular fa-moon"></i>
    {:else}
        <i class="fas fa-regular fa-sun"></i>
    {/if}
</div>
