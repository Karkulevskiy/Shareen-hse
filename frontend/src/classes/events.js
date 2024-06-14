export class Event { //Класс-событие, позволяющий хранить необходимую информацию под одним именем
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}


export class NewMessageEvent { //Сообщение в лобби
    constructor(message, from, sent) {
        this.message = message;
        this.login = from;
        this.time = sent;
    }
}


export class LobbyEvent{ //Вход и прогрузка лобби
    constructor(lobby_url="",video_url="",timing="",pause=false,users=[],chat=[]){
        this.lobby_url = lobby_url;
        this.video_url = video_url;
        this.timing = timing;
        this.pause = pause;
        this.users = users;
        this.chat = chat;
    }
}


export class VideoEvent{ //Смена или прогрузка видео
    constructor(pause=false,timing=""){
        this.pause=pause;
        this.timing = timing;
    }

}