document.addEventListener("DOMContentLoaded", function(event) {
    document.getElementById("menu-icon").onclick = function(event) {
        toggle(document.getElementsByClassName("dropdown-content").item(0));
    };
});

function show(elem) {
    elem.style.display = 'block';
}

function hide(elem) {
    elem.style.display = 'none';
}

function toggle(elem) {
    if (window.getComputedStyle(elem).display === 'block') {
        hide(elem);
        return;
    }

    show(elem);
}