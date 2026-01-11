import mainView from "./mainView.js"

class quoteView extends mainView {
    _data;
   _parentElement = document.querySelector(".rawHtml")

_generateMarkup () {
    return this._data.map(el => `<div><p>"${el.q}"</p><p>${el.a}</p>${el.h}</div>`).join(' ')
}
}

export default new quoteView()
