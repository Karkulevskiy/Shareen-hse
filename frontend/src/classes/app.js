export class App {
    constructor(model) {
        this.model = model
        this.$app=document.querySelector("#app")
    }

    render(){
        this.$app.textContent = ''
        this.model.forEach(block =>{
            this.$app.appendChild(block.ToObj())
        })
    }
}