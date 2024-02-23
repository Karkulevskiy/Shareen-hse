//import image from "./assets/search.svg"
import { ImageBlock, ButtonBlock,InputBlock,DivBlock } from "./classes/blocks.js"
 
export const model = [
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

