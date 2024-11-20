import axios from "axios";

const api = axios.create({
    baseURL : 'http://localhost:8000',
    timeout: 5000,
    headers: {
        'Content-Type' : 'application/json'
    },
    body: JSON.stringify(FormData)
})


export default api