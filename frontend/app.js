async function loadConfig() {
    const response = await fetch("/api/config");
    const config = await response.json();

    document.getElementById("site-name").textContent = config.site_name;

    const startDate = new Date(config.start_date);
    const today = new Date();

    const diff = today - startDate;
    const days = Math.floor(diff / (1000 * 60 * 60 * 24));

    document.getElementById("days").textContent = days;
}

loadConfig();
