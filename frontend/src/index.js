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

function takeButton(){
    const input = document.getElementById('input-form');
    const link = input.value;
    const MaxURL='?'

    sendRequest('POST',MaxURL,link)
    .then(data => insertVideo(data))
    .catch(err => console.log(err))
}

function insertVideo(EmbedHTML){
    const $player=document.querySelector('#player');

    $player.insertAdjacentHTML('afterbegin',EmbedHTML);

}