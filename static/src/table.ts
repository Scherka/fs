import { mainParameters, buildNewRequest } from "./request";

const tableName = document.getElementById("tableName") as HTMLElement;
const tableJSON = document.getElementById("tableJSON") as HTMLElement;

/* Изменить каталог в подписи таблицы */
function changeTableName(root: string): void {
    if (tableName) {
        tableName.textContent = root;
    } else {
        console.error("'tableName' не найден.");
    }
}

/* реакция на нажатие по строке таблицы */
function changeRootForward(row: HTMLTableRowElement): void {
    mainParameters.curRoot = `${mainParameters.curRoot}${row.cells[1].innerText}/`;
    buildNewRequest();
}

/* преобразование json в таблицу */
interface TableRowData {
    "Тип": string;
    "Имя": string;
    "Размер": string;
}

function tableFromJSON(list: TableRowData[]): void {
    changeTableName(mainParameters.curRoot);
    if (tableJSON) {
        tableJSON.innerHTML = ''; // Clear existing table content
        list.forEach(item => {
            let row = document.createElement('tr');
            if (item["Тип"] === "Дир") {
                row.className = "color";
                row.classList.add("clickableRow");
                row.onclick = function() { changeRootForward(row); };
            }
            let cellType = document.createElement('td');
            cellType.textContent = item["Тип"];
            row.appendChild(cellType);
            let cellName = document.createElement('td');
            cellName.textContent = item["Имя"];
            row.appendChild(cellName);
            let cellSize = document.createElement('td');
            cellSize.textContent = item["Размер"];
            row.appendChild(cellSize);
            tableJSON.appendChild(row);
        });
    } else {
        console.error("'tableJSON' не найден.");
    }
}

export { tableFromJSON, changeRootForward, changeTableName, tableJSON };
