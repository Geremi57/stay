export const getjson = async function(url) {
 try{
    const response = await fetch(url)
    if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);
    const res = await response.json()
    return res;
} catch (error) {
    console.error("Error fetching:", error)
        throw error
    }
}