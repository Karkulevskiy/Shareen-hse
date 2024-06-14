import { model } from './model.js'
import { App } from './classes/app.js'
import {choiceHandler} from './choicehandler.js'
import { showSignInForm } from './authorization.js';

new App(model).render();  //Передаём модель сайта классу и он её отображает
const $buttons = document.querySelector("#start-choice"); //Добавляем обработчики событий на кнопки стартового выбора и входа
$buttons.addEventListener("click",choiceHandler);
const $signButton = document.getElementById("signbutton");
$signButton.addEventListener("click",showSignInForm);
