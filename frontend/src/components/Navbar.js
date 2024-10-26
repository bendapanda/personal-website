import React from 'react';
import '../styles/Navbar.css';


const Navbar = () => {
    return (
        <ul class="navbar">
            <li><h1>Ben Shirley</h1></li>
            <li><a class="navbar-link" href="/">Home</a></li>
            <li><a class="navbar-link" href="/resume">Resume</a></li>
            <li><a class="navbar-link" href="https://github.com/bendapanda">Projects</a></li>
        </ul>
    );
};

export default Navbar;
