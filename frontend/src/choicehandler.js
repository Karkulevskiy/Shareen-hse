import { App } from "./classes/app.js";
import * as handlers from "./clickhandler.js";
import { DivBlock,InputBlock,ButtonBlock,ImageBlock,FormBlock, ScriptBlock } from "./classes/blocks.js";
import { model } from "./model.js";


export function choiceHandler(event){
    if (event.target.tagName=="DIV"){
        return;
    }
    const value = event.target.innerText;
    if (value==="CREATE"){
        model.length=0;
        model.push(
            new FormBlock(
                {
                    role:"search",
                    id:"search-form",
                    type:"submit"
                },
                [
                    new InputBlock(
                        "",
                        {
                            placeholder:"Enter the link...",
                            class:"search-form__txt",
                            type:"search",
                        }
                    ), 
                    new ButtonBlock(
                        "",
                        {
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
        new App(model).render();

        const $form = document.getElementById("search-form");
        $form.addEventListener("submit",handlers.takeButton);
    }
    else if (value=="JOIN"){
        model.length=0;
        model.push(
            new FormBlock({
                id:"lobby-form",
                onsubmit:"event.preventDefault();",
                role:"search"
            },
            [
                new InputBlock(
                    "",
                    {
                        id:"search_lobby",
                        type:"search",
                        placeholder:"Write lobby link..."
                    }
                ),
                new ButtonBlock(
                    "Go",
                    {
                        id:"go_lobby",
                        type:"submit",
                        class:"check-lobby__btn"
                    }
                )
            ])
        )
        new App(model).render();
        const $form = document.getElementById("lobby-form");
        $form.addEventListener("submit",handlers.takeButton);
    }

}