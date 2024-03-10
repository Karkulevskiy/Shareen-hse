export function div(content,options){
    const ToString = key => `${key}="${options[key]}"`;
    const Params = Object.keys(options).map(ToString).join(" ");
    return `
    <div ${Params}>${content}</div>
    `
}


export function css(styles = {}){
    if (typeof styles ==='string') {
        return styles
    }
    const ToString = key => `${key}: ${styles[key]}`
    return Object.keys(styles).map(ToString).join(';')
    
}


export function scripttag(source){
    return `<script src="${source}"></script>`
}