function appendText(text) {
    var p = document.createElement("p");
    p.innerHTML = text;
    document.querySelector(".texts").appendChild(p);
}

document.addEventListener("DOMContentLoaded", function() {
    window.setTimeout(function() {
        window.setTimeout(function() {
            appendText("✅ Assets is working!");
        }, 100);
    })
});

