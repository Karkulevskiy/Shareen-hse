import { App } from "./classes/app.js";
import { DivBlock,InputBlock,ButtonBlock,ImageBlock } from "./classes/blocks.js";
import { model } from "./model.js";


export function Handler(event){
    const value = event.target.innerText;
    if (value==="CREATE"){
        model.length=0;
        model.push(
            new DivBlock(
                {
                    class:"search-form",
                    id:"search"
                },
                [
                    new InputBlock(
                        "Enter the link",
                        {
                            class:"search-form__txt",
                            type:"text",
                            id:"input-form"
                        }
                    ), 
                    new ButtonBlock(
                        "",
                        {
                            type:"submit",
                            onclick:"javascript: takeButton()",
                            script:"clickhandler.js",
                            class:"search-form__btn"
                        },
                        [
                            new ImageBlock(
                                "./assets/search.svg",
                                {
                                alt:"image",
                                class:"search-form__img"
                            })
                        ]
                    )
                ]
            ),
            new DivBlock(
                {
                id:"player",
                class:"content"
                }
            )
        )
    }

    new App(model).render();
}