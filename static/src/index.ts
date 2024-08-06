import { statButton, sortButton, backButton, changeRootBackward, changeSort } from './button';
import { buildNewRequest } from './request';
import './styles'
//добавление к кнопкам "назад" и "сортировка" соответствующие функции
window.addEventListener("load", function () {
    if (backButton) {
        backButton.addEventListener('click', changeRootBackward);
    } else {
        console.error("Кнопка 'button-back' не найдена");
    }

    if (sortButton) {
        sortButton.addEventListener('click', changeSort);
    } else {
        console.error("Кнопка 'button-sort' не найдена.");
    }

    if (statButton) {
        statButton.addEventListener('click', ()=>
        this.window.location.href='/statistic');
    } else {
        console.error("Кнопка 'stat-sort' не найдена.");
    }
    buildNewRequest();
    
});
