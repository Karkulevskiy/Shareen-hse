export class App {
    constructor(model) {
        this.model = model
        this.$app=document.querySelector("#app")
    }

    render(){
        
        this.$app.innerHTML=''
        this.model.forEach(block =>{
            this.$app.insertAdjacentHTML('beforeend',block.ToHTML())
        })
    }
}

