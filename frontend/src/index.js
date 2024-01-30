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
    const template = 'http://localhost:5233/api/VideoPlayer/CreateLinkToVideo?url=';
    let MaxURL = template + link;
    sendRequest('POST',MaxURL,link)
    .then(data => console.log(data))
    .catch(err => console.log(err))
}
