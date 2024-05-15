import { App } from "./classes/app.js";
import { InputBlock,ButtonBlock,FormBlock } from "./classes/blocks.js";
import { model } from "./model.js";
import * as handlers from "./clickhandler.js";
import { sendEvent,connection } from "./websocket.js";
import { MyAlert } from "./utils.js";


function createLobby(){
    debugger;
    sendEvent("create_lobby",{"payload":{"login":localStorage.getItem("login")}});
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
        model.length=0;
        model.push(
            new FormBlock({
                id:"lobby-form",
                onsubmit:"event.preventDefault();",
                role:"search"
            },
            [
                new InputBlock(
                    "",
                    {
                        id:"search_lobby",
                        type:"search",
                        placeholder:"Write lobby link..."
                    }
                ),
                new ButtonBlock(
                    "Go",
                    {
                        id:"go_lobby",
                        type:"submit",
                        class:"check-lobby__btn"
                    }
                )
            ])
        )
        new App(model).render();
        const $form = document.getElementById("lobby-form");
        $form.addEventListener("submit",handlers.takeButton);
    }

}