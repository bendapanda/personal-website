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
    

    return (
        <div id="content">
            <div id="about-me">
                <div id="vert-about-me">
                    <div id="bio" class="section">
                        <h3>About Me</h3>
                        <p>Hi there! my name is Ben Shirley and I'm a -year old wannabe software developer.
                            I am just about to complete my BSc in Maths and Computer Science, both of which I love.
                        </p>
                        <p>This is my personal webpage, built by me! You can check out my projects below,
                            and leave a comment for me to review later. You can also email me directly from this webpage.
                        </p>
                    </div>
                    <div id="silly-me">
                        <img src={`${PUBLIC_URL}/resources/ben_lying_down.png`} alt="silly me!"/>
                    </div>
                </div>
                <div id="photo">
                    <img src={`${PUBLIC_URL}/resources/profile_pic.jpeg`} alt="a pircure of me" class="section" style={{borderStyle: "none"}}/> 
                </div>
            </div>
            <div id="projects">
                <h3 id="projects-header">Favorite Projects</h3>
                <ProjectCarousel />
            <div/>
            <div id="comments" class="section">
                <h3>Leave a review!</h3>
                <CommentsSection/>
            </div>
        </div>
    </div>
    );
};

export default Home;