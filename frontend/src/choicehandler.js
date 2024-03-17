import { App } from "./classes/app.js";
import { DivBlock,InputBlock,ButtonBlock,ImageBlock,FormBlock, ScriptBlock } from "./classes/blocks.js";
import { model } from "./model.js";
import * as handlers from "./clickhandler.js";
import { parseHTMLFile,loadLobby } from "./utils.js";


export function choiceHandler(event){
    if (event.target.tagName=="DIV"){
        return;
    }
    const value = event.target.innerText;
    if (value==="CREATE"){
        loadLobby()
        // const $form = document.getElementById("search-form");
        // debugger
        // $form.addEventListener("submit",handlers.takeButton);
        // let xhr = new XMLHttpRequest();
        // xhr.open('GET', './lobby.html',true);
        // xhr.responseType="text"

        // xhr.onreadystatechange = function () {
        //     if (xhr.readyState === 4 && xhr.status === 200) {
        //         let htmlContent = xhr.response;

        //         let a = htmlContent.search(/<body>/g)
        //         let b = htmlContent.search(/<!-- Code/g)
        //         htmlContent = htmlContent.slice(a,b);
        //         htmlContent +=`</body>`

        //         document.body.innerHTML=htmlContent;
        //         const $form = document.getElementById("search-form");
        //         $form.addEventListener("submit",handlers.takeButton);
        //     }
        // }

        // xhr.send();
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