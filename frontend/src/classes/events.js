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

// export class VideoEvent{
//     constructor(iframe,pause=true,timing=""){
//         let tag = document.createElement('script');
//         tag.src = 'https://www.youtube.com/iframe_api'
//         document.body.insertAdjacentHTML("beforeend",iframe);
//         this.player = new YT.Player("player",{
//             "events":{
//                 'onReady':this.onPlayerReady,
//                 'onStateChange':this.onPlayerStateChange
//         }});
//         this.pause = pause;
//         this.timing = this.player.getCurrentTime();
//         this.iframe = iframe;
//     }
//     onPlayerReady(event) {
//         // Или возобновить его
//         event.target.playVideo();
//     }
//     onPlayerStateChange(event) {
//         console.log('Player state:', event.data);
//     }

// }