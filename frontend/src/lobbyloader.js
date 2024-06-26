import { takeButton } from "./clickhandler.js";
import img from "./assets/copy.svg"
import searchimg from "./assets/search.svg"
import {Player,TimeChecker} from "./classes/videoplayer.js"
import { sendEvent } from "./websocket.js";

export function loadLobby(LobbyEvent){
    Player.paused = LobbyEvent.pause;
    Player.timing = LobbyEvent.timing;
    localStorage.setItem("lobby_url",LobbyEvent.lobby_url);
    const $app = document.querySelector("#app"); //Собираем инфо о нашем лобби

    //HTML-шаблон нашего лобби
    $app.innerHTML = `<form role="search" id="search-form" type="submit">
    <input placeholder="Enter the link..." class="search-form__txt" type="url">
    <button class="search-form__btn">
    <img src=${searchimg} class="search-form__img" alt="Поиск">
    </button>
    </form>
    <div id="player" class="content">
    </div>
    <div class="right-tabs">
    <div class="window-title">
        <div class="title">
            <span>Members</span>
        </div>
        
    </div>
    <div class="member-list">
    </div>
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
    const $form = document.getElementById("search-form"); 
    $form.addEventListener("submit",takeButton);

    const $sendbtn = $app.querySelector(".send-btn");
    $sendbtn.addEventListener("click",takeButton);

    configureLobby(LobbyEvent.users,LobbyEvent.chat); //Прорисовываем всех юзеров в лобби и все сообщения

    insertVideo(LobbyEvent.video_url); //Вставляем iframe играющего видео

    addLobbyUrl(); //Добавляем снизу ссылку на лобби с возможностью копирования
}

export function configureLobby(users,chat){
    users.forEach((user) => {
        addUser(user.login);
    });
    chat.forEach(message => {
        addMessage(message)
    });
}

export function addUser(user){ //Добавление юзера при заходе в лобби 
    let userclass ="";
    if (user==localStorage.getItem("login")){
        userclass="me";
    }
    let elem = `<div class="${userclass}">
                    <span>${user}</span>
                </div>`
    const $memlist = document.querySelector(".member-list");
    $memlist.insertAdjacentHTML("beforeend",elem);
}

export function removeUser(login){ //Удаление юзера при его выходе из лобби
    const $memlist = document.querySelector(".member-list");
    $memlist.childNodes.forEach((span)=> {
        if (span.nodeName=="DIV" && span.innerText==login){
            $memlist.removeChild(span);
        }
    })
}

export function addMessage(event){ //Добавление сообщения в чат
    let userclass = ""
    if (event.login == localStorage.getItem("login")){
        userclass="me";
    }
    let tag = `<li class="`+userclass+`">
    <div class="name">
        <span class="">`+event.login+`</span>
    </div>
    <div class="message">
        <span class="msg-time">` + convertTime(String(new Date(event.time))) + `</span>
        <br>
        <p>` + event.message + `</p>
    </div>
    </li>`
    const $chatlist = document.querySelector("#chat-list");
    $chatlist.innerHTML += tag;
}

export function insertVideo(url){ //Вставка iframe
    let playerdiv = document.querySelector("#player");
    playerdiv.innerHTML="";
    let options = {
        "width":1024,
        "height":640

    }
    let vidID="";
    if (url.indexOf("twitch.tv")!=-1){ //Twitch iframe
        vidID = takeTwitchChannel(url);
        options["channel"] = vidID;
        Player.player = new Twitch.Player("player", options);
        Player.player.addEventListener(Twitch.Player.PAUSE,sendPauseState)
        Player.player.addEventListener(Twitch.Player.PLAY,sendPlayState)
        Player.status = "twitch";
        
    }
    else if (url.indexOf("youtube.com")!=-1){ //Youtube iframe
        let html = `<div id="player-yt"></div>`;
        document.querySelector("#player").insertAdjacentHTML("afterbegin",html);
        vidID = takeYoutubeVideo(url);
        options["videoId"] = vidID;
        options["events"] = {
            'onReady':onPlayerReady,
            "onStateChange":onPlayerStateChange
        }
        Player.player = new YT.Player('player-yt', options);
        Player.status = "youtube";
        TimeChecker.startCheck();

    }
    else if (url.indexOf("vk.com")!=-1){ //vk iframe
        attrs = takeVKattributes(url);
        let iframe = `<iframe src="https://vk.com/video_ext.php?oid=${attrs.oid}&id=${attrs.id}&hd=2&js_api=1" width="1024" height="640" 
        allow="autoplay; encrypted-media; fullscreen; picture-in-picture;" frameborder="0" allowfullscreen ></iframe>`
        let playerdiv = document.querySelector("#player");
        playerdiv.insertAdjacentHTML("afterbegin",iframe)
        iframe = playerdiv.childNodes[0];
        Player.player = VK.VideoPlayer(iframe);
        Player.player.on("inited",onPlayerReady)
        Player.player.on("timeupdate",rewindVideo);
        Player.player.on("paused",sendPauseState)
        Player.player.on("resumed",sendPlayState)
        Player.status = "vkvideo";
    }

}

export function sendPauseState(event){ //Обработчики VK для постановки на паузу и снятия с неё
    sendEvent("pause_video",{
        "lobby_url":localStorage.getItem("lobby_url"),
        "pause":true
    })
}

export function sendPlayState(event){
    sendEvent("pause_video",{
        "lobby_url":localStorage.getItem("lobby_url"),
        "pause":false
    })
}

function onPlayerStateChange(event){ //Обработчик ютуба для постановки на паузу и снятия с неё
    if (event.data == YT.PlayerState.PAUSED){
        sendPauseState(null);
        TimeChecker.endCheck();
    }
    else if (event.data == YT.PlayerState.BUFFERING || event.data == YT.PlayerState.PLAYING){
        sendPlayState(null);
        TimeChecker.startCheck();
    }
}

function onPlayerReady(event){ //Действия при готовности плеера
    debugger;
    Player.rewindVideo(Player.timing);
    if (Player.paused==true){
        Player.pause();
    }
    else{
        Player.play();
    }
}

export function rewindVideo(event){ //Перемотка видео 
    let old = Player.timing;
    Player.timing = event.time;
    if (Math.abs(Player.timing-old)>1){
        let timing = Player.timing;
        sendEvent("rewind_video",{
        "lobby_url":localStorage.getItem("lobby_url"),
        "timing":timing.toString()
        })
    }
}

function takeTwitchChannel(url){ //Парсинг твич-ссылки для корректного отображения iframe
    let channel = "";
    let idx = url.indexOf("twitch.tv/")+10;
    for (let i=idx;i<url.length;i++){
        if (url[i]=='/'){
            return channel;
        }
        else{
            channel+=url[i];
        }
    }
    return channel;
}

function takeVKattributes(url){ //Парсинг вк-ссылки для корректного отображения iframe
    let regex = /-?\d*[0-9]_\d*[0-9]/
    let str =  url.match(regex)[0];
    let idx = str.indexOf("_");
    attributes = {
        "oid":str.substring(0,idx),
        "id":str.substring(idx+1)
    }
    return attributes;
}

function takeYoutubeVideo(url){ //Парсинг yt-ссылки для корректного отображения iframe
    let idx = url.indexOf("watch?v=")+8;
    let vidID = "";
        for (let i=idx;i<url.length;i++){
            if (url[i]=="&"){
                return vidID;
            }
            else{
                vidID+=url[i];
            }
        }
        return vidID;
}

function addLobbyUrl(){
    const tag = `<div class="copydiv">
    <input value="${localStorage.getItem("lobby_url")}" class = "copyurl" type="text" readonly disabled>
    <input type="image" src=${img} class="copybtn" alt="Кнопка копирования">
    </div>`
    const $app = document.querySelector("#app");
    $app.insertAdjacentHTML("beforeend",tag);
    const $copyinput = document.querySelector(".copyurl");
    document.querySelector(".copybtn").addEventListener("click",function(){
        navigator.clipboard.writeText($copyinput.value);
    })
}

function convertTime(time){ //Конвертация времени для отображения времени каждого сообщения.
    let NewTime = time.substring(4,10);
    NewTime+=" " + time.substring(16,24);
    return NewTime;
}