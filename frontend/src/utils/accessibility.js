export function onAccesibilityKeydown(e) {
	if (e.key !== "Enter" && e.key !== " ") return;
	e.preventDefault();
	e.target.click();
}