import api from '../apis'
import history from '../history'

import {
  FETCH_PRODUCTS,
  FETCH_PRODUCT,
  CREATE_PRODUCT,
  EDIT_PRODUCT,
  DELETE_PRODUCT
  } from './types'

export const fetchProducts = () => async dispatch => {
  const response = await api.get('/products')
  dispatch({ type: FETCH_PRODUCTS, payload: response.data.data})
}

export const fetchProduct = (id) => async dispatch => {
  const response = await api.get(`/products/${id}`)
  console.log("data from api", response.data)
  dispatch({ type: FETCH_PRODUCT, payload: response.data.data})
}

export const createProduct = (formValue) => async dispatch => {
 
  const response = await api.post('/products', formValue)

  history.push('/products')
}

export const updateProduct = (id, formValue) => async dispatch => {
  const response = await api.put(`/products/${id}`, formValue)

  history.push('/products')
}

export const deleteProduct = (id) => async dispatch => {
  const response = await api.delete(`/products/${id}`)

  dispatch({ type: DELETE_PRODUCT, payload: id})
  history.push('/products')
}
