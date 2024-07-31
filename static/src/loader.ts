import { sortButton, backButton, checkBackButton } from './button.ts';
import { changeTableName } from './table.ts';

/* начало и конец загрузки */
function loaderOn(): void {
    const tableJSON = document.getElementById('tableJSON') as HTMLElement;
    
    if (tableJSON) {
        tableJSON.style.visibility = 'hidden';
        tableJSON.innerHTML = '';
        changeTableName("Загрузка...");
        if(sortButton){sortButton.setAttribute("disabled", "false");}
        if(backButton){backButton.setAttribute("disabled", "false");}
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}

function loaderOff(): void {
    const tableJSON = document.getElementById('tableJSON') as HTMLElement;
    
    if (tableJSON) {
        tableJSON.style.visibility = 'visible';
        if(sortButton){sortButton.setAttribute("disabled", "false");}
        checkBackButton();
    } else {
        console.error("Элемент 'tableJSON' не найден.");
    }
}

export { loaderOn, loaderOff };
