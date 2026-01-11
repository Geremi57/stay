import { loadHotline } from "./model.js";
import quoteView from "./view/quoteView.js";
import mainView from "./view/mainView.js";
// import {  } from "./model.js";


const rawHtml = document.querySelector(".rawHtml")
loadHotline().then(data => 
    rawHtml.innerHTML = data)
    // console.log(data))
    .catch(err => console.error(err))
// console.log(loadHotline())
