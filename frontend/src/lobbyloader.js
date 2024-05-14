import { takeButton } from "./clickhandler.js";
import img from "./assets/copy.svg"
import searchimg from "./assets/search.svg"
import {Player} from "./classes/videoplayer.js"

export function loadLobby(LobbyEvent){
    localStorage.setItem("lobby_url",LobbyEvent.URL);
    const $app = document.querySelector("#app");

    $app.innerHTML = `<form role="search" id="search-form" type="submit">
    <input placeholder="Enter the link..." class="search-form__txt" type="search">
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
    <ul class="member-list">
        <li>
            <span>You</span></li>
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
    const $form = document.getElementById("search-form");
    $form.addEventListener("submit",takeButton);

    const $sendbtn = $app.querySelector(".send-btn");
    $sendbtn.addEventListener("click",takeButton);

    configureLobby(LobbyEvent.users,LobbyEvent.chat);

    insertVideo(LobbyEvent.curVideo);

    addLobbyUrl();
}

export function configureLobby(users,chat){
    users.forEach((user) => {
        let listElem = `<li><span class="status online"><i class="fa fa-circle-o"></i></span><span>` + user + `</span></li>`;
        const $memlist = document.querySelector(".member-list");
        $memlist.insertAdjacentHTML("beforebegin",listElem);
    });
    chat.forEach(message => {
        addMessage(message)
    });
}

export function addMessage(event){
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

export function insertVideo(url){
    url = "https://vk.com/video?z=video-167127847_456279379%2Fpl_cat_trends";
    let options = {
        "width":1024,
        "height":640

    }
    let vidID="";
    let service="";
    if (url.indexOf("twitch.tv")!=-1){
        service = "twitch";
        vidID = takeTwitchChannel(url);
        options["channel"] = vidID;
        Player.player = new Twitch.Player("player", options);
        Player.status = "twitch";
        
    }
    else if (url.indexOf("youtube.com")!=-1){
        service = "youtube";
        vidID = takeYoutubeVideo(url);
        options["videoId"] = vidID;
        options["events"] = {
            'onReady':onPlayerReady
        }
        Player.player = new YT.Player('player', options);
        Player.status = "youtube";
    }
    else{
        service = "vkvideo";
        attrs = takeVKattributes(url);
        let iframe = `<iframe src="https://vk.com/video_ext.php?oid=${attrs.oid}&id=${attrs.id}&hd=2&js_api=1" width="1024" height="640" 
        allow="autoplay; encrypted-media; fullscreen; picture-in-picture;" frameborder="0" allowfullscreen ></iframe>`
        let playerdiv = document.querySelector("#player");
        playerdiv.insertAdjacentHTML("afterbegin",iframe)
        debugger
        iframe = playerdiv.childNodes[0];
        Player.player = VK.VideoPlayer(iframe);
        Player.status = "vkvideo";
                
    }
    console.log(Player.getTiming());
    Player.play();

}

function onPlayerReady(event){
    event.target.playVideo();

}

function takeTwitchChannel(url){
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

function takeVKattributes(url){
    let regex = /-?\d*[0-9]_\d*[0-9]/
    let str =  url.match(regex)[0];
    let idx = str.indexOf("_");
    attributes = {
        "oid":str.substring(0,idx),
        "id":str.substring(idx+1)
    }
    return attributes;
}

function takeYoutubeVideo(url){
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

function convertTime(time){
    let NewTime = time.substring(4,10);
    NewTime+=" " + time.substring(16,24);
    return NewTime;
}
