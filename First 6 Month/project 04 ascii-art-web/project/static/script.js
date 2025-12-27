document.addEventListener("DOMContentLoaded",()=> {
    let change_color = document.getElementById("result");
    let color_err = document.getElementById("color_err");
    const colors = ["none","red", "green", "yellow", "blue", "magenta", "cyan", "white"];
    const color = change_color.dataset.color;
    if (color === "" || color === undefined) {
        // no submission, so no error
        return;
    }
    
    let found = false;
    for (let i of colors){
        if (i=== color){
            found= true;
            break;
        }
    }
    if (found == false){
        color_err.innerHTML="you must select from this colors";
    }

    if (color==="none"){
        change_color.style.color = "white";
        return;
    }
    change_color.style.color = color;

})