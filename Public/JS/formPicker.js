"use strict";
const buttons = document.querySelectorAll("div#bannerDiv ul li button");
const allFormTemplates = document.getElementsByTagName("template")[0].content;
const formDiv = document.getElementById("formDiv");
for (let button of buttons) {
    button.addEventListener("click", () => {
        formDiv.innerHTML = "";
        formDiv.appendChild(allFormTemplates.querySelector(`form#${button.innerText.replace(" ", "")}`).cloneNode(true));
    });
}
