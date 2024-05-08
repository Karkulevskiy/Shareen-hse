import { App } from "./classes/app.js";
import { DivBlock,InputBlock,ButtonBlock,ImageBlock,FormBlock, ScriptBlock } from "./classes/blocks.js";
import { model } from "./model.js";
import { loadLobby } from "./lobbyloader.js";
import * as handlers from "./clickhandler.js";


async function createLobby(){
    let MaxURL = 'http://localhost:5233/api/Lobby/CreateLobby?lobbyName=MaxHuy'
    // let options ={
    //     method: 'POST',
    //     headers:{
    //         'Content-Type':'text/plain; charset=utf-8',
    //         'Access-Control-Allow-Origin':"*"
    //     },
    //     body:"2134321"
    // }
    // fetch(MaxURL,options)
    // .then(response =>{
    //     if (!response.ok){
    //         console.log("Bad");
    //     };
    //     return response;
    //     }
    // )
    // .catch(error =>{
    //     console.log(error);
    // })
    // sendRequest('GET',MaxURL);
    axios
    .post(MaxURL)
    .then((data) => console.log(data));
    
    loadLobby();
    console.log("done")
}

function getLobbyParams(method, URL, data=null){
    return new Promise((resolve,reject) => {

        if (data==null){
            const lobbyName = prompt("Введите название лобби","простолобби");
            URL+=lobbyName
        }

        const xhr = new XMLHttpRequest();

        xhr.open(method,URL);

        

        xhr.onload = () => {
            if (xhr.status>=400)
                reject(xhr.response)
            else
                resolve(xhr.response)
        }

        xhr.onerror = () => {
            reject(xhr.response)
        }

        xhr.send(data)

    });
}

export function choiceHandler(event){
    console.log("hola")
    debugger;
    if (event.target.tagName=="DIV"){
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