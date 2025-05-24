function applySystemTheme() {
    const prefersDarkScheme = window.matchMedia("(prefers-color-scheme: dark)").matches;

    if (prefersDarkScheme) {
        document.body.classList.add("dark");
        document.body.classList.remove("light");
        document.body.setAttribute("data-theme", "dark");
        updateThemeIndicator("Dark theme applied");
    } else {
        document.body.classList.add("light");
        document.body.classList.remove("dark");
        document.body.setAttribute("data-theme", "light");
        updateThemeIndicator("Light theme applied");
    }

    updateIconSources();
}

function updateThemeIndicator(theme) {
    const themeIndicator = document.getElementById("themeIndicator");
    if (themeIndicator) {
        themeIndicator.textContent = theme;
    }
}

function updateIconSources() {
    const icons = document.querySelectorAll('.tech-icon');

    icons.forEach(icon => {
        const iconName = icon.getAttribute('data-icon');
        const currentTheme = document.body.getAttribute('data-theme');

        if (currentTheme === 'dark') {
            icon.src = `https://cdn.simpleicons.org/${iconName}/white`;
        } else {
            icon.src = `https://cdn.simpleicons.org/${iconName}/black`;
        }
    });
}

window.addEventListener("load", function () {
    applySystemTheme();
});

window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", (e) => {
    applySystemTheme();
});
