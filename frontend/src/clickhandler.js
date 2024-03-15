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
        console.log("pdfpaw")
        const lobbyLink = document.getElementById("search_lobby").value;
        const MaxURL='http://localhost:5233/api/Lobby/GetLobbyById?link=';
        const template = MaxURL+lobbyLink;

        sendRequest('POST',template)
        .then((answer) => loadLobby(answer))
        .catch(err => console.log(err))
    }
}

function loadLobby(answer){

}

function insertVideo(EmbedHTML){
    var player=document.getElementById('player');
    player.innerHTML=EmbedHTML;
}

function goHome(){

}