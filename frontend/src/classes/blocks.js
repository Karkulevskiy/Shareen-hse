
class Block{
    constructor(value,options){
        this.value = value;
        this.options = options;
    }

    ToObj(){
        throw new Error("Метод ToObj не реализован в этом классе");
    }
}

export class ScriptBlock extends Block{
    constructor(options={},inside=[]){
        super(inside,options);
    }

    ToObj(){
        const tag = document.createElement('script');
        const SetAttr = (key) => {
            tag.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr);

        return tag;
    }
}

export class FormBlock extends Block{
    constructor(options={},inside=[]){
        super(inside,options);
    }

    ToObj(){
        var form = document.createElement("form");
        
        const SetAttr = (key) => {
            form.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr)
        
        this.value.forEach(Block =>{
            form.appendChild(Block.ToObj());
        })

        return form;

    }
}

export class DivBlock extends Block{
    constructor(options={},inside=[]){
        super(inside,options);
    }

    ToObj(){
        var div = document.createElement("div");

        this.value.forEach(Block =>{
            div.appendChild(Block.ToObj());
        })

        const SetAttr = (key) => {
            div.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr)

        return div;
        
    }
}

export class TextBlock extends Block{
    constructor(value,options){
        super(value,options)
    }

    ToObj(){
        var paragraph = document.createElement("p");
        
        const SetAttr = (key) => {
            paragraph.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr);

        paragraph.textContent(this.value)
    }
}

export class InputBlock extends Block{
    constructor(value="",options={}){
        super(value,options)
    }

    ToObj(){
        var input = document.createElement("input")

        const SetAttr = (key) => {
            input.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr);

        return input
    }
}

export class ButtonBlock extends Block{
    constructor(value="",options={},inside=[]){
        super(value,options);
        this.inside=inside;
    }

    ToObj(){
        var but = document.createElement("button")

        const SetAttr = (key) => {
            but.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr);

        but.innerText = this.value;

        this.inside.forEach(Block =>{
            but.appendChild(Block.ToObj());
        })

        return but;
    }
}

export class ImageBlock extends Block{
    constructor(value,options){
        super(value,options)
    }

    ToObj(){
        var img = document.createElement("img")
        
        img.setAttribute("src",this.value);
        const SetAttr = (key) => {
            img.setAttribute(key,this.options[key]);
        };
        Object.keys(this.options).map(SetAttr);

        return img;
    }
}

