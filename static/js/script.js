const input = document.getElementById("url");
const button = document.getElementById("shorten");
const form = document.getElementById("shorten-form");

input.addEventListener("input", () => {
    button.disabled = input.value.trim() === "";
});

form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const response = await fetch("/api/shorten", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ url: input.value.trim() }),
    });

    const data = await response.json();

    document.getElementById("result").classList.remove("hidden");

    const shortUrl = document.getElementById("short-url");
    shortUrl.textContent = data.short;
    shortUrl.href = data.short;
});