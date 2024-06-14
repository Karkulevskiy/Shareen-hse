import { sendEvent } from "./websocket.js";

export function takeButton(event){  //Обработчик кликов на различные кнопки на сайте
    const id = event.target.id;
    event.preventDefault();
    if (id =="search-form"){ //Форма включения видео в лобби
        const form = event.target.elements;
        let payload = {
            "video_url":form[0].value, //Берём ссылку из формы и отправляем по вебсокету
            "lobby_url":localStorage.getItem("lobby_url")
        }
        sendEvent("insert_video_url",payload);
    }
    else if (id =="lobby-form"){ //Форма поиска лобби при нажатии на кнопку "Присоединиться"
        const lobbyLink = document.getElementById("search_lobby").value;
        let payload = { //Берём введеную ссылку и отправляем по вебсокету
            "login":localStorage.getItem("login"),
            "lobby_url":lobbyLink
        }
        sendEvent("join_lobby",payload);
    }
    else if (id=="send-chat"){ //Кнопка отправления сообщения в чата
        const $input = document.querySelector(".input-wrapper");
        let text = $input.getElementsByTagName("input")[0].value;
        let payload = {
            "login":localStorage.getItem("login"),
            "lobby_url":localStorage.getItem("lobby_url"), //Берём содержимое сообщения и отправляем по вебсокету
            "message":text
        }
        sendEvent("send_message",payload);
    }
}