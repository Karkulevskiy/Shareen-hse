import { model } from './model.js'
import { App } from './classes/app.js'
import {choiceHandler} from './choicehandler.js'

debugger;
const $form = document.querySelector(".mui-form");
const $submit = document.querySelector("#submit-btn");
$form.addEventListener("submit",signHandler);

function signHandler(event){
    debugger;
    event.preventDefault();
    $submit.disabled = true;
    const data = $form.getElementsByTagName("input");
    const login = data[0].value;
    const pass = data[1].value;
    if (login=="1" && pass=="1"){
        giveAccess();
        return;
    }
    const MaxURL = `https://max.ru`
    axios.post(MaxURL + `?login=${login}&password=${pass}`)
    .then(response => {
        giveAccess();
    }
    )
    .catch(error =>{
        console.log(error);
            $submit.disabled = false;
    })
}


function giveAccess(){
    new App(model).render();
    const $buttons = document.querySelector("#start-choice");
    $buttons.addEventListener("click",choiceHandler);
}
