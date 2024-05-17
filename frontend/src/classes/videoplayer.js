import { sendEvent } from "../websocket";

class VideoPlayer{
    player;
    status;
    paused;
    timing;
    constructor(player="",status=""){
        this.player = player;
        this.status = status;
    }
    pause(){
        this.paused=true;
        if (this.status=="youtube"){
            this.player.pauseVideo();
        }
        else{
            this.player.pause();
        }

    }
    play(){
        this.paused = false;
        if (this.status=="youtube"){
            this.player.playVideo();
        }
        else{
            this.player.play();
        }
    }
    getTiming(){
        if (this.player==""){
            return 0;
        }
        this.timing = this.player.getCurrentTime();
        return this.timing;
    }
    isPaused(){
        if (this.player==""){
            return true;
        }
        switch (this.status){
            case "twitch":
                return this.player.isPaused();
            case "youtube":
                if (this.player.getPlayerState()==1){
                    return false;
                }
                else{
                    return true;
                }
            case "vkvideo":
                if (this.player.getState()=="playing"){
                    return false;
                }
                else{
                    return true;
                }
        }
    }
    rewindVideo(timing){
        switch(this.status){
            case "youtube":
                this.player.seekTo(timing);
                break;
            case "vkvideo":
            case "twitch":
                this.player.seek(timing);
                break;
        }
    }
}

export let Player = new VideoPlayer()

class TimerChecker{
    TimeChecker
    constructor(){

    }
    startCheck(){
        Player.getTiming();
        this.TimeChecker = setInterval(this.checkTime,1000); 
    }
    checkTime(){
        let pastTiming = Player.timing;
        let curTiming = Player.getTiming();
        if (Math.abs(curTiming-pastTiming)>1.5){
            sendEvent("rewind_video",{
                "lobby_url":localStorage.getItem("lobby_url"),
                "timing":curTiming.toString()
            })
        }
    
    }
    endCheck(){
        clearInterval(this.TimeChecker);
    }
}

export let TimeChecker = new TimerChecker();