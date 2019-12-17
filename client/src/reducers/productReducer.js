import _ from 'lodash'

import {
  FETCH_PRODUCTS,
  FETCH_PRODUCT,
  CREATE_PRODUCT,
  DELETE_PRODUCT
} from '../actions/types'

export default (state = {}, action) => {
  switch (action.type) {
    case FETCH_PRODUCTS:
      return { ...state, ..._.mapKeys(action.payload, 'id')}
    case FETCH_PRODUCT:
      return { ...state, [action.payload.id]: action.payload }
    case CREATE_PRODUCT:
      return state
    case DELETE_PRODUCT:
        return _.omit(state, action.payload)
    default:
      return state
  }
}