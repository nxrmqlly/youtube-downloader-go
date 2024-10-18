<script>
import { OpenBrowser } from "../../wailsjs/go/controllers/App.js";
import { onAccesibilityKeydown } from "../utils/accessibility.js";
export let href;
export let ref;
</script>

<span data-ref={ref} on:keydown={onAccesibilityKeydown} on:click={() => OpenBrowser(href)}>
    <slot></slot>
</span>