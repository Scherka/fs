/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./static/src/button.js":
/*!******************************!*\
  !*** ./static/src/button.js ***!
  \******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   backButton: () => (/* binding */ backButton),\n/* harmony export */   changeRootBackward: () => (/* binding */ changeRootBackward),\n/* harmony export */   changeSort: () => (/* binding */ changeSort),\n/* harmony export */   checkBackButton: () => (/* binding */ checkBackButton),\n/* harmony export */   sortButton: () => (/* binding */ sortButton)\n/* harmony export */ });\n/* harmony import */ var _request_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./request.js */ \"./static/src/request.js\");\n\nconst sortButton = document.getElementById(\"buttonSort\")\nconst backButton = document.getElementById(\"buttonBack\")\n\n/* реакция на кнопку назад*/\nfunction changeRootBackward(){\n    _request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot = trimRoot(_request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot)\n    ;(0,_request_js__WEBPACK_IMPORTED_MODULE_0__.buildNewRequest)()\n}\n//блокировка кнопки назад, если пользователь пытается выйти за пределы исходной директории\nfunction checkBackButton(){\n    if (_request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot.length <= _request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.mainRoot.length){\n        backButton.disabled = true\n    }else{\n        backButton.disabled = false\n    }\n}\n\n/*реакция на кнопку сотировки */\nfunction changeSort(){\n    if (_request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort==\"asc\"){\n        _request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort = _request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort.replaceAll(_request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort, \"desc\")\n        sortButton.innerHTML = \"Сортировать по возрастанию\"\n    }else{\n        _request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort = _request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort.replaceAll(_request_js__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curSort, \"asc\")\n        sortButton.innerHTML = \"Сортировать по убыванию\"\n    }\n    (0,_request_js__WEBPACK_IMPORTED_MODULE_0__.buildNewRequest)()}\n/* обрезание root*/\nfunction trimRoot(root){\n    let lastSlash = root.lastIndexOf(\"/\")\n    if (lastSlash == root.length-1){\n        lastSlash = root.slice(0,-1).lastIndexOf(\"/\")\n    }\n    return `${root.slice(0, lastSlash)}/`\n\n}\n\n\n//# sourceURL=webpack://fs/./static/src/button.js?");

/***/ }),

/***/ "./static/src/index.js":
/*!*****************************!*\
  !*** ./static/src/index.js ***!
  \*****************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _button_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./button.js */ \"./static/src/button.js\");\n/* harmony import */ var _request_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./request.js */ \"./static/src/request.js\");\n/* harmony import */ var _loader_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./loader.js */ \"./static/src/loader.js\");\n/* harmony import */ var _table_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./table.js */ \"./static/src/table.js\");\n\n\n\n\n\n\n\nwindow.addEventListener(\"load\", \n    function(){\n    (0,_request_js__WEBPACK_IMPORTED_MODULE_1__.buildNewRequest)()\n    document.getElementById('buttonBack').addEventListener('click', _button_js__WEBPACK_IMPORTED_MODULE_0__.changeRootBackward);\n    document.getElementById('buttonSort').addEventListener('click', _button_js__WEBPACK_IMPORTED_MODULE_0__.changeSort);\n    document.getElementsByClassName('clickableRow').addEventListener('click',(0,_table_js__WEBPACK_IMPORTED_MODULE_3__.changeRootForward)())\n    });\n\n\n\n//# sourceURL=webpack://fs/./static/src/index.js?");

/***/ }),

/***/ "./static/src/loader.js":
/*!******************************!*\
  !*** ./static/src/loader.js ***!
  \******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   loaderOff: () => (/* binding */ loaderOff),\n/* harmony export */   loaderOn: () => (/* binding */ loaderOn)\n/* harmony export */ });\n/* harmony import */ var _button_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./button.js */ \"./static/src/button.js\");\n/* harmony import */ var _table_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./table.js */ \"./static/src/table.js\");\n\n\n/* начало и конец загрузки */\nfunction loaderOn(){\n    tableJSON.style.visibility='hidden';\n    tableJSON.innerHTML = ''\n    ;(0,_table_js__WEBPACK_IMPORTED_MODULE_1__.changeTableName)(\"Загрузка...\")\n    _button_js__WEBPACK_IMPORTED_MODULE_0__.sortButton.disabled = true\n    _button_js__WEBPACK_IMPORTED_MODULE_0__.backButton.disabled = true\n}\nfunction loaderOff(){\n    tableJSON.style.visibility = 'visible';\n    _button_js__WEBPACK_IMPORTED_MODULE_0__.sortButton.disabled = false\n    ;(0,_button_js__WEBPACK_IMPORTED_MODULE_0__.checkBackButton)()\n}\n\n\n\n//# sourceURL=webpack://fs/./static/src/loader.js?");

/***/ }),

/***/ "./static/src/request.js":
/*!*******************************!*\
  !*** ./static/src/request.js ***!
  \*******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   buildNewRequest: () => (/* binding */ buildNewRequest),\n/* harmony export */   mainParameters: () => (/* binding */ mainParameters)\n/* harmony export */ });\n/* harmony import */ var _loader__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./loader */ \"./static/src/loader.js\");\n/* harmony import */ var _button__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./button */ \"./static/src/button.js\");\n/* harmony import */ var _table__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./table */ \"./static/src/table.js\");\n\n\n\nconst mistakeBox = document.getElementById(\"mistakeMessage\")\nlet mainParameters = {\n  curSort: \"asc\",\n  curRoot: \"\",\n  mainRoot: \"\"\n}\n//запрос\nfunction buildNewRequest() {\n    ;(0,_loader__WEBPACK_IMPORTED_MODULE_0__.loaderOn)();\n    let url = `/fs?sort=${mainParameters.curSort}&root=${mainParameters.curRoot}`;\n  \n    fetch(url)\n      .then(response => {\n        if (!response.ok) {\n          throw new Error(`Ошибка HTTP ${response.status}`);\n        }\n        return response.json();\n      })\n      .then(responseBody => {\n        let errorCode = responseBody['ErrorCode'];\n        if (errorCode == 0) {\n          if (mainParameters.curRoot == \"\") {\n            mainParameters.mainRoot = responseBody['Root'];\n          }\n          (0,_button__WEBPACK_IMPORTED_MODULE_1__.checkBackButton)();\n          mainParameters.curRoot = responseBody['Root'];\n          let data = responseBody['Data'];\n          (0,_table__WEBPACK_IMPORTED_MODULE_2__.tableFromJSON)(data);\n        } else {\n          mistakeBox.textContent = `Ошибка выполнения запроса: ${responseBody['ErrorMessage']}`;\n        }\n        (0,_loader__WEBPACK_IMPORTED_MODULE_0__.loaderOff)();\n      })\n      .catch(error => {\n        mistakeBox.textContent = \"Ошибка во время выполнения запроса\";\n        console.error(`Ошибка fetch:`, error);\n        (0,_loader__WEBPACK_IMPORTED_MODULE_0__.loaderOff)();\n      });\n  }\n\n\n//# sourceURL=webpack://fs/./static/src/request.js?");

/***/ }),

/***/ "./static/src/table.js":
/*!*****************************!*\
  !*** ./static/src/table.js ***!
  \*****************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   changeRootForward: () => (/* binding */ changeRootForward),\n/* harmony export */   changeTableName: () => (/* binding */ changeTableName),\n/* harmony export */   tableFromJSON: () => (/* binding */ tableFromJSON)\n/* harmony export */ });\n/* harmony import */ var _request__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./request */ \"./static/src/request.js\");\n\nconst tableName = document.getElementById(\"tableName\")\nlet tableJSON = document.getElementById(\"tableJSON\")\n/* Изменить каталог в подписи таблицы */\nfunction changeTableName(root){\n    tableName.textContent = root;\n}\n/* реакция на нажатие по строке таблицы*/\nfunction changeRootForward(){\n    let row = document.getElementsByClassName('clickableRow')\n    _request__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot = _request__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot.replaceAll(_request__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot, `${_request__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot}${row.cells[1].innerText}/`)\n    ;(0,_request__WEBPACK_IMPORTED_MODULE_0__.buildNewRequest)()\n}\n/* преобразование json в таблицу */\nfunction tableFromJSON(list) {\n    changeTableName(_request__WEBPACK_IMPORTED_MODULE_0__.mainParameters.curRoot);\n    for (let i in list) {\n        let row = document.createElement('tr');\n        if (list[i][\"Тип\"] == \"Дир\") {\n            row.className = \"color\";\n            row.classList.add(\"clickableRow\")\n            row.onclick = function() { changeRootForward(this); };\n        }\n        let cellType = document.createElement('td');\n        cellType.textContent = list[i][\"Тип\"];\n        row.appendChild(cellType);\n        let cellName = document.createElement('td');\n        cellName.textContent = list[i][\"Имя\"];\n        row.appendChild(cellName);\n        let cellSize = document.createElement('td');\n        cellSize.textContent = list[i][\"Размер\"];\n        row.appendChild(cellSize);\n        tableJSON.appendChild(row);\n    }}\n\n\n//# sourceURL=webpack://fs/./static/src/table.js?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId](module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = __webpack_require__("./static/src/index.js");
/******/ 	
/******/ })()
;