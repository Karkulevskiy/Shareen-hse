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
qq
        xhr.send(data)

    });

}

function takeButton(){
    const input = document.getElementById('input-form');
    const link = input.value;
    const MaxURL='http://localhost:5233/api/VideoPlayer/CreateLinkToVideo?url='
    let template = MaxURL + link;
    sendRequest('POST',template,link)
    .then(data => insertVideo(data))
    .catch(err => console.log(err))
}

function insertVideo(EmbedHTML){
    var player=document.getElementById('player');

    player.innerHTML=EmbedHTML
}