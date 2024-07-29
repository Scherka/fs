
var mainRoot = "/home/sergey/"
var mainURL = "http://localhost:10001/fs?root=/home/sergey/&sort=asc"
function serverRequest(root){
    const fetchResp = fetch('http://localhost:10001/fs?root=/home/sergey/&sort=asc', {
        method: 'GET',
        body: JSON.stringify(dataForm),
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Methods': 'GET'
        },
      });
      respons = fetchResp;
      console.log(respons);
}
function sortButtonClick() 
{
    var btn = document.getElementById("buttonSort");
    if (btn.innerHTML=="Сортировать по убыванию") {
        btn.innerHTML = "Сортировать по возрастанию";
        let url = 'http://localhost:10001/fs?root=/home/sergey/&sort=asc'
    }
    else {btn.innerHTML = "Сортировать по убыванию";
        let url = 'http://localhost:10001/fs?root=/home/sergey/&sort=desc'
    }
    serverRequest(root)
}