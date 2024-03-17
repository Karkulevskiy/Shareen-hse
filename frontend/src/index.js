import { model } from './model.js'
import { App } from './classes/app.js'
import {choiceHandler} from './choicehandler.js'


new App(model).render();

const $buttons = document.querySelector("#start-choice");
$buttons.addEventListener("click",choiceHandler);