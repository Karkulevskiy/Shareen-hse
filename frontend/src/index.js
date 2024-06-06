import { model } from './model.js'
import { App } from './classes/app.js'
import {choiceHandler} from './choicehandler.js'
import { showSignInForm } from './authorization.js';

new App(model).render();
const $buttons = document.querySelector("#start-choice");
$buttons.addEventListener("click",choiceHandler);
const $signButton = document.getElementById("signbutton");
$signButton.addEventListener("click",showSignInForm);
