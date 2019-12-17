import axios from 'axios'

export default axios.create({
  baseURL: 'http://localhost:9292/api',
  headers: {'Content-Type': 'application/json'}
})