import { ButtonBlock,DivBlock, TextBlock } from "./classes/blocks.js"
 
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
                    class:"choice__btn",
                    onclick:"event.preventDefault();"
                },
            )
        ]
    ),
    new TextBlock("Shareen - бесплатный веб-сервис для совместного просмотра видео.\
     Shareen предлагает на выбор такие популярные видеосервисы, как  YouTube, Twitch, ВКВидео и др.\
     Создавайте комнаты, обменивайтесь сообщениями и наслаждайтесь совместным просмотром видео и фильмов с друзьями! ",
    {
        align:"center",
        class:"infotext"
    })

]

