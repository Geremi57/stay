export default class mainView{
    _data;
    render(data, render = true){
        this._data = data
        const markup = this._generateMarkup();

        if(!render) return markup

        this._data.forEach(el => {
            this._clear()
            return this._parentElement.insertAdjacentHTML('afterbegin', markup)
        })
    }
    _clear(){
        this._parentElement.innerHTML = '';
    }
}