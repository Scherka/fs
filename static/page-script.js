var mainRoot = "/home/"
const tableJSON = document.getElementById("tableJSON")
const sortButton = document.getElementById("buttonSort")
const tableName = document.getElementById("tableName")
const mistakeBox = document.getElementById("mistakeMessage")
const currentRequest = {
    sort: "asc",
    root: mainRoot,
    /*реакция на кнопку сотировки */
    changeSort(){
        if (this.sort == "asc"){
            this.sort = "desc"
            sortButton.innerHTML = "Сортировать по возрастанию";
        }else{
            this.sort = "asc"
            sortButton.innerHTML = "Сортировать по убыванию";}
        buildNewRequest()},
    /* реакция на нажатие по строке таблицы*/
    changeRootForward(row){
        const cells = row.cells;
        this.root = `${this.root}${cells[1].innerText}/`
        buildNewRequest()
    },
    /* реакция на кнопку назад*/
    changeRootBackward(){
        this.root = trimRoot(this.root)
        buildNewRequest()
    }
}

document.addEventListener("DOMContentLoaded", 
    function(){
    buildNewRequest()
    });
/* обрезание root*/
function trimRoot(root){
    var lastSlash = root.lastIndexOf("/")
    if (lastSlash == root.length-1){
        lastSlash = root.slice(0,-1).lastIndexOf("/")
    }
    return `${root.slice(0, lastSlash)}/`

    //console.log(trimmedRoot)
}


/* Изменить каталог в подписи таблицы */
function changeTableName(root){
    tableName.textContent = root;
}
/*Реакция на нажатие кнопки сортировки*/
function sortButtonReaction(){
    currentRequest.changeSort()
}
/*получение праметров из строки запроса для их проверки */
function getParameters(url){
    const urlParams = new URL(url)
    var sort = urlParams.searchParams.get('sort')
    var root = urlParams.searchParams.get('root')
    
    if ((sort !="asc" & sort !="desc") || (root == null)){
        mistakeBox.textContent = "Параметры запроса заданы неверно";
        throw new Error('Параметры запроса заданы неверно')
    }
    return [root, sort]
    }

/* запрос*/
async function buildNewRequest() {
    
    try {
      url = `http://localhost:10001/fs?sort=${currentRequest.sort}&root=${currentRequest.root}`
      changeTableName(currentRequest.root)
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`Ошибка HTTP ${response.status}`);
      }
      var data = await response.json();
      tableFromJSON(data)
    } catch (error) {
        mistakeBox.textContent = "Ошибка во время выполнения запроса";
      console.error(`Ошибка fetch:`, error);
    }
  }




/* преобразование json в таблицу */
function tableFromJSON (list){
    var html = ``;  
    for(var i in list) {
        if (list[i]["Тип"]=="Дир"){
            html += `<tr class = "color" onclick="currentRequest.changeRootForward(this)">`
        } else {
            html += `<tr>`
        }
        html += `<td>${list[i][`Тип`]}</td><td>${list[i][`Имя`]}</td><td>${list[i][`Размер`]}</td></tr>`;
        }
    tableJSON.innerHTML = html;
    }

    
