/**
 * Ben Shirley 2024
 * This is code for the navbar at the top of the page
 */
import React, {useState} from 'react';
import '../styles/navbar.css';


const Navbar = () => {
    // The only state here is whether or not the tiny me who swings from resume should be shown
    const [showHangingBen, setShowHangingBen] = useState(false);

    return (
        <ul class="navbar header">
            <li><h1>Ben Shirley</h1></li>
            <li><a class="navbar-link" href="/">Home</a></li>
            <li>
                <a id="resume-link" class="navbar-link" href="/resume"
                    onMouseEnter={() => setShowHangingBen(true)}
                    onMouseLeave={() => setShowHangingBen(false)} 
                >Resume</a>
                {
                    // If I should be shown, then add me to the dom. Animations handled automatically by css
                    showHangingBen &&
                    <img id="hanging-image" src={`${process.env.PUBLIC_URL}/resources/ben_hanging1.png`}/>
                }
                </li>
            <li><a class="navbar-link" href="https://github.com/bendapanda">Projects</a></li>
        </ul>
    );
};

export default Navbar;
