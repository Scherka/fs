import { sortButton, backButton, checkBackButton, disableButton,enableButton } from './button';
import { changeTableName, tableJSON } from './table';

/* начало и конец загрузки */
function loaderOn(): void {
    sortButton.disabled=true;
    backButton.disabled=true;
    if (tableJSON) {
        tableJSON.style.visibility = 'hidden';
        tableJSON.innerHTML = '';
        changeTableName("Загрузка...");
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}
function loaderOff(): void {
    if (tableJSON) {
        tableJSON.style.visibility = 'visible';
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
    sortButton.disabled=false;
    checkBackButton();
}

export { loaderOn, loaderOff };
