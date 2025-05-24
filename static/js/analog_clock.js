function updateClock() {
    const now = new Date();

    const minuteAngle = (now.getMinutes() + now.getSeconds() / 60) * 6;

    const hourAngle = ((now.getHours() % 12) + now.getMinutes() / 60) * 30;

    const minuteHand = document.getElementById('minuteHand');
    const hourHand = document.getElementById('hourHand');

    minuteHand.setAttribute('transform', `rotate(${minuteAngle}, 10, 10)`);
    hourHand.setAttribute('transform', `rotate(${hourAngle}, 10, 10)`);
}

updateClock();
setInterval(updateClock, 1000);