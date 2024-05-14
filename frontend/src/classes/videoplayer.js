class VideoPlayer{
    player;
    status;
    constructor(player="",status=""){
        this.player = player;
        this.status = status;
    }
    pause(){
        if (this.status=="youtube"){
            this.player.pauseVideo();
        }
        else{
            this.player.pause();
        }

    }
    play(){
        if (this.status=="youtube"){
            this.player.playVideo();
        }
        else{
            this.player.play();
        }
    }
    getTiming(){
        return this.player.getCurrentTime()
    }
    isPaused(){
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
}

export let Player = new VideoPlayer()