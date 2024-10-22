
export const getProjects = async () => {
    const response = await fetch("http://localhost:8080/api/projects");

    if (!response.ok) {
        throw new Error(`HTTP error getting projects! ${response.status}`);
    }

    const data = await response.json();
    
    return data;
}
