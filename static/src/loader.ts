import { sortButton, backButton, checkBackButton } from './button';
import { changeTableName } from './table';

/* начало и конец загрузки */
function loaderOn(): void {
    const tableJSON = document.getElementById('tableJSON') as HTMLElement;
    
    if (tableJSON) {
        tableJSON.style.visibility = 'hidden';
        tableJSON.innerHTML = '';
        changeTableName("Загрузка...");
        if(sortButton){sortButton.disabled=true;}
        if(backButton){backButton.disabled=true;}
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}

function loaderOff(): void {
    const tableJSON = document.getElementById('tableJSON') as HTMLElement;
    
    if (tableJSON) {
        tableJSON.style.visibility = 'visible';
        if(sortButton){sortButton.disabled=false;}
        checkBackButton();
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}

export { loaderOn, loaderOff };
