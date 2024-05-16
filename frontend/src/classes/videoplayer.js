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
        if (this.player==""){
            return 0;
        }
        return this.player.getCurrentTime()
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