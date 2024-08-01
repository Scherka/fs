import { mainParameters, buildNewRequest } from './request';

const sortButton = document.getElementById("buttonSort")as HTMLButtonElement;
const backButton = document.getElementById("buttonBack") as HTMLButtonElement;
const asc:string = 'asc';
const desc:string = 'desc';
function changeRootBackward(): void {
    mainParameters.curRoot = trimRoot(mainParameters.curRoot);
    buildNewRequest();
}

// блокировка кнопки назад, если пользователь пытается выйти за пределы исходной директории
function checkBackButton(): void {
    if (mainParameters.curRoot.length <= mainParameters.mainRoot.length) {
        if(backButton){backButton.disabled=true};
    } else {
        if(backButton){backButton.disabled=false};
    }
}

/* реакция на кнопку сотировки */
function changeSort(): void {
    if (mainParameters.curSort == asc) {
        mainParameters.curSort = desc;
        if (sortButton)sortButton.innerHTML = "Сортировать по возрастанию";
    } else {
        mainParameters.curSort = asc;
        if (sortButton) sortButton.innerHTML = "Сортировать по убыванию";
    }
    buildNewRequest();
}

/* обрезание root */
function trimRoot(root: string): string {
    let lastSlash = root.lastIndexOf("/");
    if (lastSlash == root.length - 1) {
        lastSlash = root.slice(0, -1).lastIndexOf("/");
    }
    return `${root.slice(0, lastSlash)}/`;
}

export { sortButton, backButton, changeRootBackward, checkBackButton, changeSort};
