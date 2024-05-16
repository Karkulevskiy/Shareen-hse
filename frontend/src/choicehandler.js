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
        if (document.getElementById("lobby-form")!=undefined){
            return;
        }
        let $app = document.querySelector("#app");

        let HTML = `<form id="lobby-form" onsubmit="event.preventDefault();" role="search">
                        <input id="search_lobby" type="text" placeholder="Write lobby link...">
                        <button id="go_lobby" type="submit" class="check-lobby__btn">Go</button>
                    </form>`
        
        $app.insertAdjacentHTML("afterbegin",HTML);
        const $form = document.getElementById("lobby-form");
        $form.classList.add("active");
        $form.addEventListener("submit",handlers.takeButton);
    }

}