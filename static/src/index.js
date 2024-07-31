import { changeRootBackward, changeSort } from './button.js';
import {buildNewRequest} from './request.js';
window.addEventListener("load", 
    function(){
    buildNewRequest()
    document.getElementById('buttonBack').addEventListener('click', changeRootBackward);
    document.getElementById('buttonSort').addEventListener('click', changeSort);
    //document.getElementsByClassName('clickableRow').addEventListener('click',changeRootForward())
    });

