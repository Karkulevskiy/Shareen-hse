export function div(content,options){
    const ToString = key => `${key}="${options[key]}"`;
    const Params = Object.keys(options).map(ToString).join(" ");
    return `
    <div ${Params}>${content}</div>
    `
}

export function css(styles = {}){
    if (typeof styles ==='string') {
        return styles
    }
    const ToString = key => `${key}: ${styles[key]}`
    return Object.keys(styles).map(ToString).join(';')
    
}


export function addScript(source){
        const tag = document.createElement('script');
        tag.setAttribute('src',source);
        
        const $app = document.querySelector('#app');
        $app.append(tag);
}

export function loadLobby(lobbyLink = null){
    let app = document.querySelector("#app");
    axios.get("lobby.html")
    .then(response => app.innerHTML=response);
    // fetch("../frontend/src/lobby.html")
    // .then(response => response.text())
    // .then(text => app.innerHTML=text);
    // axios.get("http://localhost:5001/frontend/src/lobby.html")
    // .then((response)=>console.log(response.data));

    // var xhr
    // if (window.ActiveXObject)
    // {
    //  xhr = new ActiveXObject("Microsoft.XMLHTTP");
    // }
    // else if (window.XMLHttpRequest)
    // {
    //  xhr = new XMLHttpRequest();
    // } 
    //     xhr.open('GET', './lobby.html',true);
    //     xhr.responseType="text"

    //     xhr.onload = function () {
        
    //         let htmlContent = xhr.responseText;

    //         const content = parseHTMLFile(htmlContent);

    //         document.body.innerHTML=content;
    //         const $form = document.getElementById("search-form");
    //         $form.addEventListener("submit",handlers.takeButton);
    //     }
    //     xhr.send();
    //     console.log(xhr.responseText)
}

export function sendRequest(method,URL,data=null){
    var xhr
    if (window.ActiveXObject)
    {
     xhr = new ActiveXObject("Microsoft.XMLHTTP");
    }
    else if (window.XMLHttpRequest)
    {
     xhr = new XMLHttpRequest();
    } 
    xhr.open(method,URL);
    xhr.onreadystatechange = () => {
        if (xhr.status==200 && xhr.readyState==4){
            console.log("Success")
            return false;
        }
    }

    xhr.onerror = () => {
        console.log(xhr.response)
    }

    xhr.send(data);
}

export function parseHTMLFile(content){
    let BodyStart = content.search(/<body>/g)
    let BodyEnd = content.search(/<!-- Code/g)
    content = content.slice(BodyStart,BodyEnd);
    content +=`</body>`
    return content 
}

export function addMessage(person="",content,time=""){
    let userclass = ""
    if (person == ""){
        userclass="me";
    }
    if (time==""){
        let curDate = new Date();
        time+=curDate.getHours() +":" + curDate.getMinutes();
    }
    let tag = `<li class="`+userclass+`">
    <div class="name">
        <span class="">`+person+`</span>
    </div>
    <div class="message">
        <p>` + content + `</p>
        <span class="msg-time">` + time + `</span>
    </div>
    </li>`
    const $chatlist = document.querySelector("#chat-list");
    $chatlist.innerHTML += tag;
}