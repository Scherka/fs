let mask = document.querySelector(".mask")
const loader = document.getElementById("loader")
let tableJSON = document.getElementById("tableJSON")
let mainRoot = ""
const sortButton = document.getElementById("buttonSort")
const backButton = document.getElementById("buttonBack")
const tableName = document.getElementById("tableName")
const mistakeBox = document.getElementById("mistakeMessage")
let curSort = "asc"
let curRoot = ""

window.addEventListener('load', ()=>{
    mask.classList.add('hide')
    mask.remove()
})
document.addEventListener("DOMContentLoaded", 
    function(){
    buildNewRequest()
    });

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
//блокировка кнопки назад, если пользователь пытается выйти за пределы исходной директории
function checkBackButton(){
    if (curRoot.length <= mainRoot.length){
        backButton.disabled = true
    }else{
        backButton.disabled = false
    }
}
//запрос
function buildNewRequest() {
    loaderOn();
    url = `/fs?sort=${curSort}&root=${curRoot}`;
  
    fetch(url)
      .then(response => {
        if (!response.ok) {
          throw new Error(`Ошибка HTTP ${response.status}`);
        }
        return response.json();
      })
      .then(responseBody => {
        let errorCode = responseBody['ErrorCode'];
        if (errorCode == 0) {
          if (curRoot == "") {
            mainRoot = responseBody['Root'];
          }
          checkBackButton();
          curRoot = responseBody['Root'];
          let data = responseBody['Data'];
          tableFromJSON(data);
        } else {
          mistakeBox.textContent = `Ошибка выполнения запроса: ${responseBody['ErrorMessage']}`;
        }
        loaderOff();
      })
      .catch(error => {
        mistakeBox.textContent = "Ошибка во время выполнения запроса";
        console.error(`Ошибка fetch:`, error);
        loaderOff();
      });
  }
  

/* начало и конец загрузки */
function loaderOn(){
    tableJSON.style.visibility='hidden';
    tableJSON.innerHTML = ''
    changeTableName("Загрузка...")
    sortButton.disabled = true
    backButton.disabled = true
}
function loaderOff(){
    tableJSON.style.visibility = 'visible';
    sortButton.disabled = false
    checkBackButton()
}

/* преобразование json в таблицу */
function tableFromJSON(list) {
    changeTableName(curRoot);
    for (let i in list) {
        let row = document.createElement('tr');
        if (list[i]["Тип"] == "Дир") {
            row.className = "color";
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


    
