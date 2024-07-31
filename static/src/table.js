import { mainParameters, buildNewRequest} from "./request";
const tableName = document.getElementById("tableName")
let tableJSON = document.getElementById("tableJSON")
/* Изменить каталог в подписи таблицы */
function changeTableName(root){
    tableName.textContent = root;
}
/* реакция на нажатие по строке таблицы*/
function changeRootForward(){
    let row = document.getElementsByClassName('clickableRow')
    mainParameters.curRoot = mainParameters.curRoot.replaceAll(mainParameters.curRoot, `${mainParameters.curRoot}${row.cells[1].innerText}/`)
    buildNewRequest()
}
/* преобразование json в таблицу */
function tableFromJSON(list) {
    changeTableName(mainParameters.curRoot);
    for (let i in list) {
        let row = document.createElement('tr');
        if (list[i]["Тип"] == "Дир") {
            row.className = "color";
            row.classList.add("clickableRow")
            row.onclick = function() { changeRootForward(this); };
        }
        let cellType = document.createElement('td');
        cellType.textContent = list[i]["Тип"];
        row.appendChild(cellType);
        let cellName = document.createElement('td');
        cellName.textContent = list[i]["Имя"];
        row.appendChild(cellName);
        let cellSize = document.createElement('td');
        cellSize.textContent = list[i]["Размер"];
        row.appendChild(cellSize);
        tableJSON.appendChild(row);
    }}
export {tableFromJSON, changeRootForward, changeTableName}