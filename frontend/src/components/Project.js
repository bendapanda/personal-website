/**
 * This component handles the fetching and displaying of my project information.
 */

import React, { useState, useEffect } from "react";
import Slider from 'react-slick';

import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

import "../styles/project.css"
import { getProjects } from "../services/APIService"

const ProjectCarousel = () => {
    const [projects, setProjects] = useState([]);

    // method that prompts the server for a list of projects.
    const handleProjects = async () => {
        try {
            const result = await getProjects();
            setProjects(result);

        } catch(error) {
            console.error(error);
        }
    };

    useEffect(() => {
        handleProjects();
    }, []);

    const sliderSettings = {
        className: "slider",
        dots: true,
        infinite: true,
        centerMode: true,
        speed: 500,
        slidesToShow: 1,
        slidesToScroll: 1,
        variableWidth: true,
    };

    return (
        <Slider {...sliderSettings}>
            {
                projects.map((project) => {
                    return <Project key={project.Name} project={project}/>;
                })
            }
            <div style={{padding: "20px"}}>
            <div class="project">
                <img src={`${process.env.PUBLIC_URL}/resources/ben_squat_point.png`} style={{maxHeight: "101%", maxWidth: "100%"}}/>
            </div>
            </div>
        </Slider>
    );
};

const Project = ({ project }) => {
    return (
        
        <a class="project" href={project.URL}>
            <div class="project-left">
                <h3 class="project-title">{ project.Name }</h3>
                <p class="dates">{ project.Started } to { project.Finished }</p>
                <p class="project-desc">{ project.Description }</p>
            </div>
            <div class="project-right">
                <img src={ `${process.env.REACT_APP_API_URL}/public/${project.ImageFile }` } alt=""/>
            </div>
        </a>
    )
};

export default ProjectCarousel
