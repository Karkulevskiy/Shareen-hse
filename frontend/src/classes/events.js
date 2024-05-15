export class Event {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}


export class NewMessageEvent {
    constructor(message, from, sent) {
        this.message = message;
        this.login = from;
        this.time = sent;
    }
}


export class LobbyEvent{
    constructor(lobby_url="",video_url="",timing="",pause=false,users=[],chat=[]){
        debugger;
        this.lobby_url = lobby_url;
        this.video_url = video_url;
        this.timing = timing;
        this.pause = pause;
        this.users = users;
        this.chat = chat;
    }
}


export class VideoEvent{
    constructor(pause=false,timing=""){
        this.pause=pause;
        this.timing = timing;
    }

}