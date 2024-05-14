export class Event {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}

class SendMessageEvent {
    constructor(message, from) {
        this.message = message;
        this.from = from;
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
    constructor(lobbyURL="",videoURL="",Timing="",Pause=false,Users=[],Chat=[]){
        this.URL = lobbyURL;
        this.curVideo = videoURL;
        this.timing = Timing;
        this.pause = Pause;
        this.users = Users;
        this.chat = Chat;
    }
}


export class VideoEvent{
    constructor(vidID,pause=false,timing=""){
        this.vidID = vidID;
        this.pause=pause;
        this.timing = timing;
    }

}