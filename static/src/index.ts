import { changeRootBackward, changeSort } from './button.ts';
import { buildNewRequest } from './request.ts';

window.addEventListener("load", function () {
    buildNewRequest();

    const buttonBack = document.getElementById('buttonBack');
    const buttonSort = document.getElementById('buttonSort');

    if (buttonBack) {
        buttonBack.addEventListener('click', changeRootBackward);
    } else {
        console.error("Кнопка 'buttonBack' не найдена");
    }

    if (buttonSort) {
        buttonSort.addEventListener('click', changeSort);
    } else {
        console.error("Кнопка 'buttonSort' не найдена.");
    }
});
