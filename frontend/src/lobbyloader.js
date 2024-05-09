import { takeButton } from "./clickhandler.js";

export function loadLobby(LobbyEvent){
    localStorage.setItem("lobby_url",LobbyEvent.URL);
    const $app = document.querySelector("#app");
    let image = document.createElement("img");
    image.src = "assets/search.svg";
    image.class = "search-form__img";

    $app.innerHTML = `<form role="search" id="search-form" type="submit">
    <input placeholder="Enter the link..." class="search-form__txt" type="search">
    <button class="search-form__btn">
    </button>
    </form>
    <div id="player" class="content">
    </div>
    <div class="right-tabs">
    <ul class="tabs">
        <div class="title">
            <span>Members</span>
        </div>
        
    </ul>
    <ul class="tabs-container">
        <li class="active">
            <ul class="member-list">
                <li><span class="status online"><i class="fa fa-circle-o"></i></span><span>You</span></li>
            </ul>
        </li>
    </ul>
    <div></div>
    </div>
    <div class="window-wrapper">
    <div class="window-title">
        <div class="title">
            <span>Chat</span>
        </div>
    </div>
    <div class="window-area">
        
        <div class="chat-area">
            <div class="chat-list">
            <ul id="chat-list">
            </ul>    
            </div>
            <div class="input-area">
                <div class="input-wrapper">
                    <input type="text" value="">
                    <i class="fa fa-smile-o"></i>
                </div>
                <input type="button" value="Submit" class="send-btn" id="send-chat">
            </div>
        </div>
    </div>
    </div>
    </div>`;
    let $btn = $app.querySelector(".search-form__btn");
    $btn.appendChild(image);
    const $form = document.getElementById("search-form");
    $form.addEventListener("submit",takeButton);

    const $sendbtn = $app.querySelector(".send-btn");
    $sendbtn.addEventListener("click",takeButton);

    insertVideo(LobbyEvent.curVideo);

    configureLobby(LobbyEvent.users,LobbyEvent.chat);

    
}

export function configureLobby(users,chat){
    users.forEach((user) => {
        let listElem = `<li><span class="status online"><i class="fa fa-circle-o"></i></span><span>` + user + `</span></li>`;
        const $memlist = document.querySelector(".member-list");
        $memlist.innerHTML+=listElem;
    });
    chat.forEach(message => {
        addMessage(message)
    });
}

export function addMessage(event){
    let userclass = ""
    if (event.login == ""){
        userclass="me";
    }
    let tag = `<li class="`+userclass+`">
    <div class="name">
        <span class="">`+event.login+`</span>
    </div>
    <div class="message">
        <p>` + event.message + `</p>
        <span class="msg-time">` + event.time + `</span>
    </div>
    </li>`
    const $chatlist = document.querySelector("#chat-list");
    $chatlist.innerHTML += tag;
}

export function insertVideo(EmbedHTML){
    var player=document.getElementById('player'); //Комментарий
    player.innerHTML=EmbedHTML;
}