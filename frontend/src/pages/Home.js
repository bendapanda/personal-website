/**
 * Ben Shirley 2024
 * This page is the main homepage of my website - it will be the first thing the user sees.
 */
import React, { useState, useEffect } from "react";

import "../styles/main.css";
import ProjectCarousel from "../components/Project";
import CommentsSection from "../components/Comments";

const { PUBLIC_URL } = process.env;

const Home = () => {
    // I was born in early January, so I'm not fussed about my age being a few days out of sync :)
    const date = new Date();

    return (
        <div id="content">
            <div id="about-me">
                <div id="vert-about-me">
                    <div id="bio" class="section">
                        <h3 className="header">About Me</h3>
                        <p>Hi there! my name is Ben Shirley and I'm a {date.getFullYear() - 2004}-year old wannabe software developer.
                            I am just about to complete my BSc in Maths and Computer Science, both of which I love.
                        </p>
                        <p>This is my personal webpage, built by me! You can check out my projects below,
                            and leave a comment for me to review later. You can also email me directly from this webpage.
                        </p>
                    </div>
                    <div id="silly-me-container">
                    <div id="silly-me">
                        <img src={`${PUBLIC_URL}/resources/ben_lying_down.png`} alt="silly me!"/>
                    </div>
                    </div>
                </div>
                <div id="photo">
                    <img src={`${PUBLIC_URL}/resources/profile_pic.jpeg`} alt="a picture of me" class="section" style={{padding: "0px"}}/> 
                </div>
            </div>
            <div id="projects" style={{marginTop: "-4%"}}>
                <h3 id="projects-header" className="header">Favorite Projects</h3>
                <ProjectCarousel />
            </div>
            <div id="comments" className="section">
                <h3 className="header">Leave a review!</h3>
                <CommentsSection/>
            </div>
    </div>
    );
};

export default Home;