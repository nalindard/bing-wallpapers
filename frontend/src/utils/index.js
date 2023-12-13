import { store } from "../store"

//@ts-check


export async function get(link) {
    // const res = await fetch('https://raw.githubusercontent.com/niumoo/bing-wallpaper/main/docs/2021-02.html', { cache: 'force-cache' })
    const res = await fetch(link, { cache: 'force-cache' })
    const htmlData = await res.text()
    return htmlData
}

/*
 * 
 * @param {string} htmlString 
 * @returns {[JSON]}
 */

export function getJson(htmlString) {
    let jsonData = []

    const parser = new DOMParser()
    const html_ = parser.parseFromString(htmlString, 'text/html')
    const nodeList = html_.querySelectorAll('.w3-third');
    [...nodeList].forEach(node => jsonData.push(nodeToData(node)))

    return jsonData
}

/*
 * 
 * @param {HTMLDivElement} node 
 * @returns {{
 * date: string | undefined,
 * img: string | undefined | null,
 * smallThumbnail: string | undefined | null,
 * bigThumbnail: string | undefined | null,
 * }}
 */

function nodeToData(node) {
    return {
        date: node.querySelector('p')?.innerHTML?.slice(0, 10),
        img: node.querySelector('a')?.getAttribute('href'),
        smallThumbnail: node.querySelector('.smallImg')?.getAttribute('src'),
        bigThumbnail: node.querySelector('.bigImg')?.getAttribute('src'),
    }
}

export function getSrcLink() {
    return `https://raw.githubusercontent.com/niumoo/bing-wallpaper/main/docs/${store.year}-${store.getMonth()}.html`
}