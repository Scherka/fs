import { loaderOn, loaderOff } from "./loader.ts";
import { checkBackButton } from "./button.ts";
import { tableFromJSON, changeTableName } from "./table.ts";

const mistakeBox = document.getElementById("mistakeMessage") as HTMLElement;

interface MainParameters {
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
  var url:string = `/fs?sort=${mainParameters.curSort}&root=${mainParameters.curRoot}`;

  fetch(url)
    .then(response => {
      if (!response.ok) {
        throw new Error(`Ошибка HTTP ${response.status}`);
      }
      return response.json();
    })
    .then(responseBody => {
      var errorCode = responseBody['ErrorCode'];
      if (mainParameters.curRoot === "") {
        mainParameters.mainRoot = responseBody['Root'];
      }
      checkBackButton();
      mainParameters.curRoot = responseBody['Root'];
      changeTableName(mainParameters.curRoot);
      if (errorCode === 0) {
        var data = responseBody['Data'];
        tableFromJSON(data);
      } else {
        if (mistakeBox) {
          mistakeBox.textContent = `Ошибка выполнения запроса: ${responseBody['ErrorMessage']}`;
        }
      }
      loaderOff();
    })
    .catch(error => {
      if (mistakeBox) {
        mistakeBox.textContent = "Ошибка во время выполнения запроса";
      }
      console.error(`Ошибка fetch:`, error);
      loaderOff();
    });
}

export { mainParameters, buildNewRequest };
