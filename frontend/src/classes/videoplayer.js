import { sendEvent } from "../websocket";

class VideoPlayer{ //Видеоплеер
    player; //Класс предоставляемый api хостинга, к которому мы обращаемся
    status; //Хостинг, к которому мы обращаемся
    paused; //На паузе или нет
    timing; //Тайминг
    constructor(player="",status=""){
        this.player = player;
        this.status = status;
    }
    pause(){ //Поставить на паузу
        this.paused=true;
        if (this.status=="youtube"){
            this.player.pauseVideo();
        }
        else{
            if (this.status=="vkvideo"){
            }
            this.player.pause();
        }

    }
    play(){ //Начать проигрывать
        this.paused = false;
        if (this.status=="youtube"){
            this.player.playVideo();
        }
        else{
            if (this.status=="vkvideo"){
            }
            this.player.play();
        }
    }
    getTiming(){ //Выяснить тайминг текущего плеера
        if (this.player==""){
            return 0;
        }
        this.timing = this.player.getCurrentTime();
        return this.timing;
    }
    isPaused(){ //Узнать состояние плеера(на паузе или нет)
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
    rewindVideo(timing){ //Перемотать видео в плеере
        switch(this.status){
            case "youtube":
                this.player.seekTo(timing);
                break;
            case "vkvideo":
                this.player.seek(timing);
                break;
            case "twitch":
                this.player.seek(timing);
                break;
        }
    }
}

export let Player = new VideoPlayer()

class TimerChecker{ //Вспомогательный класс для реализации перемотки видео через youtube(т.к. у api youtube нет этого метода)
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