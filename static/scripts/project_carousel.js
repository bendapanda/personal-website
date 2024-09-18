

function setOrder() {
    let i=0;
    let projects = document.getElementsByClassName("project");
    for(i=0;i<projects.length; i++) {
        projects[i].style.setProperty("--order", i);
    }
}

function changeProj(n) {
    let i=0;
    let projects = document.getElementsByClassName("project");
    for(i=0;i<projects.length; i++) {
        let currentOrder = projects[i].style.getPropertyValue("--order");
        let newOrder = (parseInt(currentOrder)+1) % projects.length;
        projects[i].style.setProperty("--order", newOrder);
    }
}

setOrder();