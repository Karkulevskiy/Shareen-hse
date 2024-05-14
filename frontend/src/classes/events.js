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

export class NewLobbyEvent{
    constructor(lobbyURL){
        this.URL = lobbyURL;
    }
}

export class LobbyEvent{
    constructor(lobbyURL="",videoURL="",timing="",pause=false,users=[],chat=[]){
        this.lobbyURL = lobbyURL;
        this.videoURL = videoURL;
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