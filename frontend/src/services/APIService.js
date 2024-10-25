/**
 * Ben Shirley 2024
 * handles the communication with the api enpoints published by the go server
 * 
 */


/**
 * 
 * @returns An array of project elements
 */
export const getProjects = async () => {
    const response = await fetch(`${process.env.REACT_APP_API_URL}/api/projects`);

    if (!response.ok) {
        throw new Error(`HTTP error getting projects! ${response.status}`);
    }
    
    const data = await response.json();
    return data;
}


/**
 *  Prompts the server for a list of the ids of the comments in the database 
 * @returns A list of all comment ids in the database
 */
export const getCommentIds = async () => {
    const response = await fetch(`${process.env.REACT_APP_API_URL}/api/comments/all`)

    if (!response.ok) {
        throw new Error(`HTTP error getting comments! ${response.status}`);
    }
    const data = await response.json();
    return data
}

/**
 * prompts the server for a specific comment by its id
 * @param { int } id the id of the desired comment 
 * @returns an object representation of the comment
 */
export const getComment = async (id) => {
    const url = new URL(`${process.env.REACT_APP_API_URL}/api/comments`);
    url.searchParams.append("id", id);

    const response = await fetch(url)
    if (!response.ok) {
        throw new Error(`HTTP error getting target comment! ${response.status}`);
    }
    const data = await response.json();
    return data
}
