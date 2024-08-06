import { loaderOn, loaderOff } from "./loader";
import { checkBackButton } from "./button";
import { tableFromJSON, changeTableName } from "./table";

const mistakeBox = document.getElementById("mistake-message") as HTMLElement;
const mainField = document.getElementById("main-field") as HTMLElement;

type MainParameters = {
  curSort: string;
  curRoot: string;
  mainRoot: string;
}

let mainParameters: MainParameters = {
  curSort: "asc",
  curRoot: "",
  mainRoot: ""
};

// запрос
function buildNewRequest(): void {
  loaderOn();
  //формирование запроса
  let url:string = `/fs?sort=${mainParameters.curSort}&root=${mainParameters.curRoot}`;
  //запрос и обработка ответа
  fetch(url)
    .then(response => {
      if (!response.ok) {
        throw new Error(`Ошибка HTTP ${response.status}`);
      }
      return response.json();
    })
    .then(responseBody => {
      var errorCode = responseBody['error_code'];
      if (mainParameters.curRoot === "") {
        mainParameters.mainRoot = responseBody['root'];
      }
      checkBackButton();
      mainParameters.curRoot = responseBody['root'];
      //переименование таблицы в соотвествии с текущей папкой
      changeTableName(mainParameters.curRoot);
      if (errorCode === 0) {
        var data = responseBody['data'];
        tableFromJSON(data);
      } else {
        if (mistakeBox) {
          mistakeBox.textContent = `Ошибка выполнения запроса: ${responseBody['ErrorMessage']}`;
        }
      }
    })
    .catch(error => {
      if (mistakeBox) {
        mistakeBox.textContent = `Ошибка во время выполнения запроса: ${error}`;
      }
      console.error(`Ошибка fetch:`, error);
    }).finally(()=>loaderOff());
}


export { mainParameters, buildNewRequest};
