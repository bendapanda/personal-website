/**
 * Ben Shirley 2024
 * This page displays my resume in html form.
 */
import React, { useState, useEffect } from "react";

import { getCv } from "../services/api-service";

import "../styles/main.css";
import "../styles/resume.css";

const Resume = () => {
    const [htmlContent, setHtmlContent] = useState(null);

    useEffect(() => {
        // Fetch the HTML from the API. This is done only once on page load
        const fetchHtml = async () => {
            try {
                const html = await getCv(); 
                setHtmlContent(html);
            } catch (error) {
                console.error('Error fetching HTML:', error);
            }
        };

        fetchHtml();
    }, []);

    // The auto-generated html sources its image from a local directory,
    // which won't work here. So after it is loaded we must source my profile manually.
    useEffect(() => {
        const images = document.querySelectorAll("img");
        images.forEach((img, index) => {
            console.log("image found")
            img.src = `${process.env.PUBLIC_URL}/resources/profile_pic.jpeg`;  // Set dynamically
        });
    }, [htmlContent]);
 

    if (!htmlContent) {
        return <p>Loading...</p>;
    }

    // We update the inner html dangerously, which could pose a security issue.
    // I think it is safe enough as the api address is constant, but it might be worth 
    // looking into before deployment
    return (
        <div id="resume-container">
            <div id="resume-content">
                <link rel="stylesheet" type="text/css" href={`${process.env.REACT_APP_API_URL}/public/resources/cv.css`} />
                <div
                    dangerouslySetInnerHTML={{ __html: htmlContent }}
                />
            </div>
            <div id="resume-sidebar">
                <a href={`${process.env.REACT_APP_API_URL}/public/resources/cv.pdf`} target="_blank" download>View as a PDF!</a>
                <br/>
                <p>
                    Nerd Note: I am pretty proud of this section, so I am making a note of it here!
                    I write my cv in latex, because it produces pretty pdfs. However, it's a bit of a pain to 
                    always have to re-write it on the website every time I make changes. So, I have this page set
                    to automatically compile my cv from latex to html, which is then rendered here! The server also 
                    compiles my cv to a pdf, which you can view above! This means that all I have to do to update my 
                    cv is change a single file!
                </p>
            </div>
        </div>
    );
};

export default Resume;