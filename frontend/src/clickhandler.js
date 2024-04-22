import { addMessage, parseHTMLFile } from "./utils.js";
import { configureLobby, loadLobby } from "./lobbyloader.js";

function sendRequest(method, URL, data = null){
    return new Promise((resolve,reject) => {
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

export function takeButton(event){  
    const id = event.target.id;
    event.preventDefault();
    if (id =="search-form"){
        const form = event.target.elements;
        const link = form[0].value;
        const MaxURL='http://localhost:5233/api/VideoPlayer/CreateLinkToVideo?url='
        let template = MaxURL + link;

        sendRequest('POST',template,link)
        .then(data => insertVideo(data))
        .catch(err => console.log(err))
    }
    else if (id =="lobby-form"){
        console.log("pdfpaw");
        debugger;
        const lobbyLink = document.getElementById("search_lobby").value;
        const MaxURL='http://localhost:5233/api/Lobby/GetLobbyById?link=';
        let template = MaxURL+lobbyLink;

        sendRequest('GET',template,lobbyLink)
        .then((answer) => {
            console.log(answer);
            loadLobby();
        })
        .catch(err => console.log(err))
        template = `http://localhost:5233/api/Lobby/GetLobbyList`
        sendRequest('GET',template)
        .then((answer)=>{
            let info = JSON.parse(answer);
            info.lobbies.forEach(lobby =>{
                if (lobby.lobbyLink==lobbyLink){
                    let users = [];
                    lobby.users.forEach(user=>{
                        users.push(user.name);
                    })
                    configureLobby(users);
                }
            })
            debugger
            console.log(answer)
        })
    }
    else if (id=="send-chat"){
        const $input = document.querySelector(".input-wrapper");
        let text = $input.getElementsByTagName("input")[0].value;
        addMessage("",text,"");
    }
}


function insertVideo(EmbedHTML){
    var player=document.getElementById('player');
    player.innerHTML=EmbedHTML;
}