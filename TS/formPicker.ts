const buttons: NodeListOf<HTMLButtonElement> = document.querySelectorAll("div#bannerDiv ul li button")//get all buttons

const allFormTemplates: DocumentFragment = document.getElementsByTagName("template")[0].content//get the forms from the template
const formDiv: HTMLElement  = document.getElementById("formDiv")

for(let button of buttons){//add event listener to change the form
    button.addEventListener("click",()=>{//display the form in the main page
        formDiv.innerHTML = ""
        formDiv.appendChild(allFormTemplates.querySelector(`form#${button.innerText.replace(" ","")}`).cloneNode(true))
    })
}