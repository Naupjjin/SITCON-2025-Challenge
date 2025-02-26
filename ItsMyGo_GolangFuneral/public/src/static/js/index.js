function openSidebar() {
    document.getElementById("sidebar").style.width = "250px";
}

function closeSidebar() {
    document.getElementById("sidebar").style.width = "0";
}
function toggleInfo(item) {
    const infoBox = item.querySelector('.info-box');
    const allInfoBoxes = document.querySelectorAll('.info-box');

    allInfoBoxes.forEach(box => {
        if (box !== infoBox) {
            box.classList.remove('active');
        }
    });


    infoBox.classList.toggle('active');
}
