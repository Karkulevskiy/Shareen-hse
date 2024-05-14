import { Event,NewMessageEvent,LobbyEvent } from "./classes/events";
import { addMessage,loadLobby,insertVideo,addUser } from "./lobbyloader";
import { MyAlert } from "./utils";
import { Player } from "./classes/videoplayer";

export let connection = [];

export function routeEvent(event) {
    if (event.type === undefined) {
        console.log("no 'type' field in event(report about that to the developer)");
        return;
    }
    if (event.type=="join_lobby" && event.status==400){
        MyAlert("Lobby with entered ID does not exist!","error");
        return;
    }
    if (event.status!=200){
        console.log(`Unexpected websocket error(response status is ${event.status})`)
        return;
    }
    switch (event.type) {
        case "send_message":
            const messageEvent = Object.assign(new NewMessageEvent, event.payload);
            addMessage(messageEvent);
            break;
        case "create_lobby":
            const newLobby = new LobbyEvent(event.payload.lobby_url,"","",true,[{"login":localStorage.getItem("login")}]);
            loadLobby(newLobby);
            break;
        case "join_lobby":
            const lobby = Object.assign(new LobbyEvent,event.payload)
            loadLobby(lobby);
            break;
        case "insert_video_url":
            insertVideo(event.payload.url);
            break;
        case "pause_video":
            if (event.payload.pause==true){
                Player.pause();
            }
            else{
                Player.play();
            }
            break;
        case "get_video_timing":
            debugger;
            let timing = Player.getTiming().toString();
            sendEvent("get_video_timing",{
                "login":event.payload.login,
                "pause":Player.isPaused(),
                "timing":timing
            })
            break;
        case "user_join_lobby":
            addUser(event.payload.login);
            break;
        case "disconnected":
            break;
        default:
            alert("unsupported message type");
            break;
    }

}

export function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload);
    connection[0].send(JSON.stringify(event));
}