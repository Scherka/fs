import {mainParameters, buildNewRequest} from './request.js'
const sortButton = document.getElementById("buttonSort")
const backButton = document.getElementById("buttonBack")

/* реакция на кнопку назад*/
function changeRootBackward(){
    mainParameters.curRoot = trimRoot(mainParameters.curRoot)
    buildNewRequest()
}
//блокировка кнопки назад, если пользователь пытается выйти за пределы исходной директории
function checkBackButton(){
    if (mainParameters.curRoot.length <= mainParameters.mainRoot.length){
        backButton.disabled = true
    }else{
        backButton.disabled = false
    }
}

/*реакция на кнопку сотировки */
function changeSort(){
    if (mainParameters.curSort=="asc"){
        mainParameters.curSort = mainParameters.curSort.replaceAll(mainParameters.curSort, "desc")
        sortButton.innerHTML = "Сортировать по возрастанию"
    }else{
        mainParameters.curSort = mainParameters.curSort.replaceAll(mainParameters.curSort, "asc")
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
export {sortButton, backButton,changeRootBackward, checkBackButton, changeSort}