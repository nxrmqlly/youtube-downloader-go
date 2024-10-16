
<script>
import { onAccesibilityKeydown } from "../utils/accessibility.js";
export let title;
export let content;
export let onClose;
</script>

<style>
    .modal {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .modal-content {
        background: #f6f3f4;
        padding: 20px;
        border-radius: 8px;
        width: 300px;
        text-align: center;
    }
</style>

<div class="modal" on:click={onClose} on:keydown={onAccesibilityKeydown}>
    <div class="modal-content" on:click|stopPropagation on:keydown={onAccesibilityKeydown}>
        <h2>{title}</h2>
        <p>{content}</p>
        <button on:click={onClose}>Close</button>
    </div>
</div>
