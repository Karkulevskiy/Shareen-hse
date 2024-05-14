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
}

export let Player = new VideoPlayer()