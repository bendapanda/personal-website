
const getProjects = async () => {
    response = await fetch("http://localhost:8080/projects")

    if (!response.ok) {
        throw new Error(`HTTP error getting projects! ${response.status}`)
    }

    const data = await response.json();
    console.log(data)
    
    return [];
}