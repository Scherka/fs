import { mainParameters, buildNewRequest } from "./request";

const tableName = document.getElementById("table-name") as HTMLElement;
const tableJSON = document.getElementById("table-json") as HTMLElement;

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
type TableRowData = {
    "type": string;
    "name": string;
    "size": string;
}
//загрузка информации из json в таблицу
function tableFromJSON(list: TableRowData[]): void {
    changeTableName(mainParameters.curRoot);
    if (tableJSON) {
        tableJSON.innerHTML = ''; //очистка тела таблицы
        list.forEach(item => {
            let row = document.createElement('tr');
            if (item["type"] === "Дир") {
                row.className = "color";
                row.classList.add("clickable-row");
                row.onclick = function() { changeRootForward(row); };
            }
            //заполнение ячеек
            let cellType = document.createElement('td');
            cellType.textContent = item["type"];
            row.appendChild(cellType);
            let cellName = document.createElement('td');
            cellName.textContent = item["name"];
            row.appendChild(cellName);
            let cellSize = document.createElement('td');
            cellSize.textContent = item["size"];
            row.appendChild(cellSize);
            tableJSON.appendChild(row);
        });
    } else {
        console.error("'table-json' не найден.");
    }
}

export { tableFromJSON, changeRootForward, changeTableName, tableJSON };
