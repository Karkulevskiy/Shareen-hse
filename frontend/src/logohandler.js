import { App } from "./classes/app.js";
import { DivBlock,ButtonBlock } from "./classes/blocks.js";
import { choiceHandler } from "./choicehandler.js";

const $logo = document.querySelector(".logoname");

$logo.addEventListener("click",logoHandler);

function logoHandler(){
    const model = [
        new DivBlock(
            {
                id:"start-choice"
            },
            [
                new ButtonBlock(
                    "JOIN",
                    {
                        type:"button",
                        class:"choice__btn"
                    },
                ),
                new ButtonBlock(
                    "CREATE",
                    {
                        type:"button",
                        class:"choice__btn"
                    },
                )
            ]
        )
    ]

    new App(model).render();
    const $buttons = document.querySelector("#start-choice");
    $buttons.addEventListener("click",choiceHandler);
}