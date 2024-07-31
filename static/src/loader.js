import {sortButton, backButton, checkBackButton} from './button.js';
import { changeTableName } from './table.js';
/* начало и конец загрузки */
function loaderOn(){
    tableJSON.style.visibility='hidden';
    tableJSON.innerHTML = ''
    changeTableName("Загрузка...")
    sortButton.disabled = true
    backButton.disabled = true
}
function loaderOff(){
    tableJSON.style.visibility = 'visible';
    sortButton.disabled = false
    checkBackButton()
}

export {loaderOn, loaderOff}