"use strict";
let buttons = document.querySelectorAll("div#bannerDiv ul li button");
for (let button of buttons) {
    button.addEventListener("click", () => {
        console.log(button.innerHTML);
    });
}
