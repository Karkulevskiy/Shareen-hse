import { sendEvent } from "./websocket.js";

export function takeButton(event){  
    const id = event.target.id;
    event.preventDefault();
    if (id =="search-form"){
        const form = event.target.elements;
        let payload = {
            "video_url":form[0].value,
            "lobby_url":localStorage.getItem("lobby_url")
        }
        sendEvent("insert_video_url",JSON.stringify.payload);
    }
    else if (id =="lobby-form"){
        const lobbyLink = document.getElementById("search_lobby").value;
        let payload = {
            "login":localStorage.getItem("login"),
            "lobby_url":lobbyLink
        }
        sendEvent("join_lobby",payload);
    }
    else if (id=="send-chat"){
        const $input = document.querySelector(".input-wrapper");
        let text = $input.getElementsByTagName("input")[0].value;
        let payload = {
            "login":localStorage.getItem("login"),
            "lobby_url":localStorage.getItem("lobby_url"), //Комментарий
            "message":text
        }
        sendEvent("send_message",payload);
    }
}