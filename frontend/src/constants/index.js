const consts = {
    baseUrl: "https://raw.githubusercontent.com/niumoo/bing-wallpaper/main/docs/",
    firstSrc: 202102,
    get lastSrc() {
        const date = new Date();
        let year = date.getFullYear();
        let month = date.getMonth() + 1;
        if (month < 10) {
            month = "0" + month;
        }
        return `${year}${month}`;
    },
    srcLink: function (date = "202102") {
        return `${this.baseUrl}${date.slice(0, 4)}-${date.slice(4)}.html`
    }
}

export default consts