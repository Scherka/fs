import { statButton, sortButton, backButton, changeRootBackward, changeSort } from './button';
import { buildNewRequest } from './request';
import './styles'
import Chart from 'chart.js/auto';
//добавление к кнопкам "назад" и "сортировка" соответствующие функции
window.addEventListener("load", function () {
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

    if (statButton) {
        statButton.addEventListener('click', ()=>
        this.window.location.href='/statistic');
    } else {
        console.error("Кнопка 'statSort' не найдена.");
    }
    buildNewRequest();
    
});
