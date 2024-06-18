import {choiceHandler} from './choicehandler.js'
import { showSignInForm } from './authorization.js';


init_page();
const logoname = document.querySelector(".logoname")
logoname.addEventListener("click",init_page);

export function init_page(){
    const HTML = `<div id="start-choice">
                    <button type="button" class="choice__btn">Присоединиться</button>
                    <button type="button" class="choice__btn">Создать</button>
                </div>
                <p align="center" class="infotext">Shareen - бесплатный веб-сервис для совместного просмотра видео.     
                Shareen предлагает на выбор такие популярные видеосервисы, как  YouTube, Twitch, ВКВидео и др.     
                Создавайте комнаты, обменивайтесь сообщениями и наслаждайтесь совместным просмотром видео и фильмов с друзьями! </p>`

    const app = document.querySelector("#app");
    app.innerHTML = HTML;  //Вставляем наш начальный HTML
    const $buttons = document.querySelector("#start-choice"); //Добавляем обработчики событий на кнопки стартового выбора и входа
    $buttons.addEventListener("click",choiceHandler);
    const $signButton = document.getElementById("signbutton");
    $signButton.addEventListener("click",showSignInForm);
}
