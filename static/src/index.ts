import { sortButton, backButton, changeRootBackward, changeSort } from './button';
import { buildNewRequest } from './request';

window.addEventListener("load", function () {
    buildNewRequest();
    if (backButton) {
        backButton.addEventListener('click', changeRootBackward);
    } else {
        console.error("Кнопка 'buttonBack' не найдена");
    }

    if (sortButton) {
        sortButton.addEventListener('click', changeSort);
    } else {
        console.error("Кнопка 'buttonSort' не найдена.");
    }
});
