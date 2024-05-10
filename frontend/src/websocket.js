import { Event,NewMessageEvent,LobbyEvent } from "./classes/events";
import { addMessage,loadLobby,insertVideo } from "./lobbyloader";

export let connection = [];

export function routeEvent(event) {

    if (event.type === undefined) {
        alert("no 'type' field in event");
    }
    if (event.status!=200){
        alert(`Unexpected websocket error(response status is ${event.status})`)
    }
    debugger;
    switch (event.type) {
        case "send_message":
            const messageEvent = Object.assign(new NewMessageEvent, event.payload);
            addMessage(messageEvent);
            break;
        case "create_lobby":
            const newLobby = new LobbyEvent(event.payload.lobby_url);
            loadLobby(newLobby);
            break;
        case "join_lobby":
            const lobby = Object.assign(new LobbyEvent,event.payload);
            loadLobby(newLobby);
            break;
        case "insert_video_url":
            insertVideo(event.payload.iframe);
            break;
        case "pause_video":

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