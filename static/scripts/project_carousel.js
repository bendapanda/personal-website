
/**
 * Gets all projects, and gives them an inital order property.
 */
function setOrder() {
    let i=0;
    let projects = document.getElementsByClassName("project");
    for(i=0;i<projects.length; i++) {
        projects[i].style.setProperty("--order", i);
    }
}

/**
 * When arrows are clicked, we want to increment or decrement the order of these elements.
 * Then, because the project tiles may have differing widths, we need to recalculate where they should
 * be positioned. This is stored in the --offset variable, which the css code references.
 */
function changeProj(n) {
    let PROJ_SPACING = 30;
    let i=0;
    let projects = Array.from(document.getElementsByClassName("project"));
    let runningOffset=0;
    for(i=0;i<projects.length; i++) {
        let currentOrder = projects[i].style.getPropertyValue("--order");
        let newOrder = (parseInt(currentOrder)-n) % projects.length;
        projects[i].style.setProperty("--order", newOrder)
    }

    projects.sort((a, b) => {return a.style.getPropertyValue("--order") < b.style.getPropertyValue("--order")});
    for(i=0; i<projects.length; i++) {
        projects[i].style.setProperty("--offset", runningOffset + "px");
        let elementWidth = projects[i].getBoundingClientRect().width;
        runningOffset+=elementWidth + PROJ_SPACING;
    }
}

setOrder();
changeProj(1);