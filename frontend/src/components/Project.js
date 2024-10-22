import React from "react";


const Project = (project) => {
    return (
        <div class="project fade">
            <div class="project-left">
                <h3 class="project-title">{ project.name }</h3>
                <p class="dates">{ project.started } to { project.finished }</p>
                <p class="project-desc">{ project.description }</p>
                <a href={ project.url }>Project Link</a>
            </div>
            <div class="project-right">
                <img src={ project.imageFile } alt=""/>
            </div>
        </div>

    )
}

export default Project
