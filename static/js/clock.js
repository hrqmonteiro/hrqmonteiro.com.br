function updateClock() {
    const now = new Date();
    const clockElement = document.getElementById('clock');
    if (clockElement) {
        clockElement.textContent = now.toLocaleTimeString('pt-br');
    }
}

updateClock();
setInterval(updateClock, 1000);