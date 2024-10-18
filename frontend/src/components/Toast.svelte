<script>
import { onAccesibilityKeydown } from "../utils/accessibility.js";
export let message;
export let onClose;

// auto close the toast after 3 seconds
setTimeout(() => {
	onClose();
}, 3000);
</script>

<style>
.toast {
    position: fixed;
    bottom: 1rem;
    right: 1rem;
    background: red;
    color: white;
    padding: 0.5em 1em;
    border-radius: 0.5rem;
    font-size: 1rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    cursor: pointer;
    user-select:text;
}

</style>

<div class="toast" on:click={onClose} on:keydown={onAccesibilityKeydown}>
    <span>{message}</span>
</div>
