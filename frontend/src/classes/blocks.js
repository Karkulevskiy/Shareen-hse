import { scripttag,div } from "../utils.js";

class Block{
    constructor(value,options){
        this.value = value;
        this.options = options;
    }

    ToHTML(){
        throw new Error("Метод ToHTML не реализован в этом классе")
    }
}

export class DivBlock extends Block{
    constructor(options={},inside=[]){
        super(inside,options)
    }

    ToHTML(){
        let content = ``; 
        this.value.forEach(Block => {
            content+=Block.ToHTML();
        });
        return div(content,this.options);
    }
}

export class TextBlock extends Block{
    constructor(value,options){
        super(value,options)
    }

    ToHTML(){
        const {} = this.options
        return div(`<p>${this.value}</p>`)
    }
}

export class InputBlock extends Block{
    constructor(value,options){
        super(value,options)
    }

    ToHTML(){
        const ToString = key => `${key}="${this.options[key]}"`;
        const Params = Object.keys(this.options).map(ToString).join(" ")
        return `<input ${Params} placeholder="${this.value}">`;
    }
}

export class ButtonBlock extends Block{
    constructor(value,options,inside=[]){
        super(value,options);
        this.inside=inside;
    }

    ToHTML(){
        const ToString = key => `${key}="${this.options[key]}"`;
        const Params = Object.keys(this.options).map(ToString).join(" ");
        let HTML = `<button ${Params}>${this.value}`;
        if (this.inside.length!=0){
            this.inside.forEach(block => {
                HTML+=block.ToHTML();
            });
        }   
        const {script=''} = this.options;
        if (script!='')
            HTML+=scripttag(script);

        HTML+=`</button>`;
        return HTML;
    }
}

export class ImageBlock extends Block{
    constructor(value,options){
        super(value,options)
    }

    ToHTML(){
        const ToString = key => `${key}="${this.options[key]}"`;
        const Params = Object.keys(this.options).map(ToString).join(" ");
        return `<img src="${this.value}" ${Params}>`
    }
}

