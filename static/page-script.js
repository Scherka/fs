document.addEventListener("DOMContentLoaded", 
    function(){
    buildNewRequest()
    });
window.addEventListener('load', function() {
        var blackout = document.getElementById('loader');
        blackout.style.display = 'none';
    
        // Показываем основное содержимое страницы
        var content = document.getElementById('content');
        content.style.display = 'block';
    });
let mainRoot = "/home/sergey/"
const loader = document.getElementById("loader")
const tableJSON = document.getElementById("tableJSON")
const sortButton = document.getElementById("buttonSort")
const tableName = document.getElementById("tableName")
const mistakeBox = document.getElementById("mistakeMessage")
let curSort = "asc"
let curRoot = mainRoot
/* реакция на нажатие по строке таблицы*/
function changeRootForward(row){
    const cells = row.cells;
    curRoot = curRoot.replaceAll(curRoot, `${curRoot}${cells[1].innerText}/`)
    buildNewRequest()
}
/* реакция на кнопку назад*/
function changeRootBackward(){
    curRoot = trimRoot(curRoot)
    buildNewRequest()
}

/*реакция на кнопку сотировки */
function changeSort(){
    if (curSort=="asc"){
        curSort = curSort.replaceAll(curSort, "desc")
        sortButton.innerHTML = "Сортировать по возрастанию"
    }else{
        curSort = curSort.replaceAll(curSort, "asc")
        sortButton.innerHTML = "Сортировать по убыванию"
    }
    buildNewRequest()}
/* обрезание root*/
function trimRoot(root){
    let lastSlash = root.lastIndexOf("/")
    if (lastSlash == root.length-1){
        lastSlash = root.slice(0,-1).lastIndexOf("/")
    }
    return `${root.slice(0, lastSlash)}/`

}


/* Изменить каталог в подписи таблицы */
function changeTableName(root){
    tableName.textContent = root;
}

/* запрос*/
async function buildNewRequest() {
    try {
      url = `/fs?sort=${curSort}&root=${curRoot}`
      changeTableName(curRoot)
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`Ошибка HTTP ${response.status}`);
      }
      let data = await response.json();
      tableFromJSON(data)

    } catch (error) {
        mistakeBox.textContent = "Ошибка во время выполнения запроса";
      console.error(`Ошибка fetch:`, error);
    }
  }


function loaderOn(){
    loader.style.width="100%"
    loader.style.height="100%"
}
function loaderOff(){
    loader.style.width="0%"
    loader.style.height="0%"
}

/* преобразование json в таблицу */
function tableFromJSON (list){
    let html = ``;  
    for(let i in list) {
        if (list[i]["Тип"]=="Дир"){
            html += `<tr class = "color" onclick="changeRootForward(this)">`
        } else {
            html += `<tr>`
        }
        html += `<td>${list[i][`Тип`]}</td><td>${list[i][`Имя`]}</td><td>${list[i][`Размер`]}</td></tr>`;
        }
    tableJSON.innerHTML = html;
    }

    
