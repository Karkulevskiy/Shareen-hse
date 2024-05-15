import { App } from "./classes/app.js";
import { InputBlock,ButtonBlock,FormBlock } from "./classes/blocks.js";
import { model } from "./model.js";
import * as handlers from "./clickhandler.js";
import { sendEvent,connection } from "./websocket.js";
import { MyAlert } from "./utils.js";


function createLobby(){
    debugger;
    sendEvent("create_lobby",{"login":localStorage.getItem("login")});
}

export function choiceHandler(event){
    if (event.target.tagName=="DIV"){
        return;
    }
    if (connection.length==0){
        MyAlert("You have to sign in before getting Shareen experience","info");
        return;
    }
    const value = event.target.innerText;
    if (value==="CREATE"){
        createLobby();
    }
    else if (value=="JOIN"){
        let $app = document.querySelector(".app");
        $app.insertAdjacentHTML("afterbegin",HTML);

        let HTML = `<form id="lobby-form" onsubmit="event.preventDefault();" role="search">
                        <input id="search_lobby" type="search" placeholder="Write lobby link...">
                        <button id="go_lobby" type="submit" class="check-lobby__btn">Go</button>
                    </form>`
        
        const $form = document.getElementById("lobby-form");
        $form.addEventListener("submit",handlers.takeButton);
    }

}