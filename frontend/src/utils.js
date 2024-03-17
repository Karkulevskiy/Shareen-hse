import * as handlers from "./clickhandler.js";

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

export function loadLobby(){
    let xhr = new XMLHttpRequest();
        xhr.open('GET', './lobby.html',true);
        xhr.responseType="text"

        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                let htmlContent = xhr.response;

                const content = parseHTMLFile(htmlContent);

                document.body.innerHTML=content;
                const $form = document.getElementById("search-form");
                $form.addEventListener("submit",handlers.takeButton);
            }
        }
        xhr.send()
}

export function parseHTMLFile(content){
    let BodyStart = content.search(/<body>/g)
    let BodyEnd = content.search(/<!-- Code/g)
    content = content.slice(BodyStart,BodyEnd);
    content +=`</body>`
    return content 
}