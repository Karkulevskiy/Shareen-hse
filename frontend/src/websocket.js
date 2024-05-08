import { Event,NewMessageEvent,LobbyEvent } from "./classes/events";

function routeEvent(event) {

    if (event.type === undefined) {
        alert("no 'type' field in event");
    }
    if (event.status!=200){
        alert(`Unexpected websocket error(response status is ${event.status})`)
    }
    switch (event.type) {
        case "send_message":
            const messageEvent = Object.assign(new NewMessageEvent, event.payload);
            appendChatMessage(messageEvent);
            break;
        case "create_lobby":
            const newLobby = new LobbyEvent(event.payload.lobby_url);
            break;
        case "join_lobby":
            const lobby = Object.assign(new LobbyEvent,event.payload);
            break;
        case "insert_video_url":

            break;
        case "pause_video":
            break;

        default:
            alert("unsupported message type");
            break;
    }

}

function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload);
    conn.send(JSON.stringify(event));
}