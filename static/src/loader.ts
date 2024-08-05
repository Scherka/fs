import {statButton, sortButton, backButton, checkBackButton } from './button';
import { changeTableName, tableJSON } from './table';

/* начало и конец загрузки */
function loaderOn(): void {
    //блокировка кнопок
    sortButton.disabled=true;
    backButton.disabled=true;
    statButton.disabled=true;
    //сокрытие тела таблицы
    if (tableJSON) {
        tableJSON.style.visibility = 'hidden';
        tableJSON.innerHTML = '';
        changeTableName("Загрузка...");
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}
function loaderOff(): void {
    //разблокирорвка кнопок
    sortButton.disabled=false;
    statButton.disabled=false;
    checkBackButton()
    if (tableJSON) {
        tableJSON.style.visibility = 'visible';
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}

export { loaderOn, loaderOff };
