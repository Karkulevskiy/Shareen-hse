import errorIcon from "./assets/icons/error.svg"
import infoIcon from "./assets/icons/info.svg"
import successIcon from "./assets/icons/success.svg"

export function div(content,options){ //Функция для создания тэга div и его конвертации в HTML-строку
    const ToString = key => `${key}="${options[key]}"`;
    const Params = Object.keys(options).map(ToString).join(" ");
    return `
    <div ${Params}>${content}</div>
    `
}

export function css(styles = {}){ //Добавление стилей
    if (typeof styles ==='string') {
        return styles
    }
    const ToString = key => `${key}: ${styles[key]}`
    return Object.keys(styles).map(ToString).join(';')
    
}

export function addScript(source){         //Добавление скрипта на HTML-страницу
        const tag = document.createElement('script');
        tag.setAttribute('src',source);
        
        const $app = document.querySelector('#app');
        $app.append(tag);
}

export function loadLobby(lobbyLink = null){  //Загрузка лобби
    let app = document.querySelector("#app"); //Получаем содержимое лобби
    axios.get("lobby.html")
    .then(response => app.innerHTML=response);
}

export function sendRequest(method,URL,data=null){ //Отправка запроса на определённый URL
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

export function parseHTMLFile(content){ //Парсинг HTML-страницы
    let BodyStart = content.search(/<body>/g)
    let BodyEnd = content.search(/<!-- Code/g)
    content = content.slice(BodyStart,BodyEnd);
    content +=`</body>`
    return content 
}

export function MyAlert(text,status){ //Самописный Alert с определённым цветом(З - OK, К - ERROR , С - INFO)
    let alert = `<div class="banner ${status} hidden">
                    <img src=${findIcon(status)} class="banner-icon">
                    <div class="banner-message">${text}
                    </div>
                    <div class="banner-close"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24"
                     viewBox="0 0 24 24" class="eva eva-close-outline" fill="#ffffff"><g data-name="Layer 2">
                     <g data-name="close"><rect width="24" height="24" transform="rotate(180 12 12)" opacity="0"></rect>
                     <path d="M13.41 12l4.3-4.29a1 1 0 1 0-1.42-1.42L12 10.59l-4.29-4.3a1 1 0 0 0-1.42 1.42l4.3 4.29-4.3 4.29a1 1 0 0 0 0 1.42 1 1 0 0 0 1.42 0l4.29-4.3 4.29 4.3a1 1 0 0 0 1.42 0 1 1 0 0 0 0-1.42z"></path>
                     </g></g></svg></div>
                </div>` //html алерта
    document.body.insertAdjacentHTML("afterbegin",alert); //Появление алерта и добавление обработчика на кнопку закрытия
    const banner = document.querySelector(".banner");
    const closebtn = document.querySelector(".banner-close");
    closebtn.addEventListener("click",(()=>{
        banner.remove();
    }))
    banner.classList.replace("hidden","visible");
    setTimeout(() => {
        banner.classList.replace("visible","hidden");
        setTimeout(()=>{
            banner.remove();
        },600)
      }, 2500);
}

function findIcon(status){ //Поиск подходящей иконки для алерта
    if (status=="error"){
        return errorIcon;
    }
    else if(status=="success") return successIcon;
    else return infoIcon;
}