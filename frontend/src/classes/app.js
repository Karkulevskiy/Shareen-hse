export class App {
    constructor(model) { //Модель - массив, состоящий из блоков. Блоки описаны в blocks.js
        this.model = model
        this.$app=document.querySelector("#app")
    }

    render(){ //Каждый блок из модели отрисосываем на странице
        this.$app.textContent = ''
        this.model.forEach(block =>{
            this.$app.appendChild(block.ToObj())
        })
    }
}