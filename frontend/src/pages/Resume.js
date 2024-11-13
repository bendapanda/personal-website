/**
 * Ben Shirley 2024
 * This page displays my resume in html form.
 */
import React, { useState, useEffect } from "react";

import { getCv } from "../services/APIService";

import "../styles/main.css";
import "../styles/Resume.css";

const Resume = () => {
    const [htmlContent, setHtmlContent] = useState(null);

    useEffect(() => {
        // Fetch the HTML from the API
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