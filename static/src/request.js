import { loaderOn,loaderOff } from "./loader";
import { checkBackButton } from "./button";
import { tableFromJSON } from "./table";
const mistakeBox = document.getElementById("mistakeMessage")
let mainParameters = {
  curSort: "asc",
  curRoot: "",
  mainRoot: ""
}
//запрос
function buildNewRequest() {
    loaderOn();
    let url = `/fs?sort=${mainParameters.curSort}&root=${mainParameters.curRoot}`;
  
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
          if (mainParameters.curRoot == "") {
            mainParameters.mainRoot = responseBody['Root'];
          }
          checkBackButton();
          mainParameters.curRoot = responseBody['Root'];
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
export {mainParameters, buildNewRequest}