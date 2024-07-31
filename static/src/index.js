import { changeRootBackward, changeSort } from './button.js';
import {mainParameters, buildNewRequest} from './request.js';
import { changeRootForward } from './table.js';
window.addEventListener("load", 
    function(){
    buildNewRequest()
    document.getElementById('buttonBack').addEventListener('click', changeRootBackward);
    document.getElementById('buttonSort').addEventListener('click', changeSort);
    //document.getElementsByClassName('clickableRow').addEventListener('click',changeRootForward())
    });

