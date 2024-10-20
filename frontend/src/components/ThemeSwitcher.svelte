<script>
import { onMount } from "svelte";
import { onAccesibilityKeydown } from "../utils/accessibility.js";
import "@fortawesome/fontawesome-free/css/all.min.css";

let isDarkMode;

function setTheme(theme) {
	document.querySelector("html").setAttribute("data-theme", theme);
	localStorage.setItem("theme", theme);
}

onMount(() => {
	const savedTheme = localStorage.getItem("theme");
	if (savedTheme) {
		isDarkMode = savedTheme === "dark";
	} else {
		isDarkMode = window.matchMedia("(prefers-color-scheme: dark)").matches;
	}
	setTheme(isDarkMode ? "dark" : "light");

	// system theme changes listening
	window
		.matchMedia("(prefers-color-scheme: dark)")
		.addEventListener("change", (e) => {
			isDarkMode = e.matches;
			setTheme(isDarkMode ? "dark" : "light");
		});
});

// Toggle theme function
function toggleTheme() {
	isDarkMode = !isDarkMode;
	setTheme(isDarkMode ? "dark" : "light");
}
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
