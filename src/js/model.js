import { HOTLINES_URL } from "./config.js";
// import { QUOT_URL } from "./config.js";
import { getjson } from "./helpers.js";

const state = {
    country: "",
    hotline: "",
    topics: [],
    description: "",
    name:"",
    website:"",
    webChat:"",
    timezones:"",
    sms:"",
}

export const stateFilm = async function () {
    const response = await getjson(`${HOTLINES_URL}`)
    state.website

}

export const loadHotline = async function () {
 const answer = await getjson(`${HOTLINES_URL}`)
 const hotline = answer["Algeria"].raw_html
 return hotline
}

// export const loadQuotes = async function() {
//  const answer = await getjson(`${QUOTES_URL}`)
//  console.log(answer.helplines)
//  return answer.helplines
// }
