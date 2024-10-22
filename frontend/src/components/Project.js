import React, { useState, useEffect } from "react";

import "../styles/project.css"

const Project = ({ project }) => {
    return (
        <div class="project fade" style={{width: 800}}>
            <div class="project-left">
                <h3 class="project-title">{ project.Name }</h3>
                <p class="dates">{ project.Started } to { project.Finished }</p>
                <p class="project-desc">{ project.Description }</p>
                <a href={ project.Url }>Project Link</a>
            </div>
            <div class="project-right">
                <img src={ `${process.env.PUBLIC_URL}${project.ImageFile }` } alt=""/>
            </div>
        </div>

    )
}

export default Project
