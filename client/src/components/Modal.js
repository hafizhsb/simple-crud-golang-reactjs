import React from 'react'
import ReactDOM from 'react-dom'
import { Link } from 'react-router-dom'
import history from '../history'

import {
  Button,
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Row,
} from 'reactstrap';

const ModalComp = (props) => {
  return ReactDOM.createPortal(
    <Modal
      isOpen={true}
      toggle={() => history.push('/products')}
    >    
      <ModalHeader>Delete Product</ModalHeader>
      <ModalBody>
        Are you sure want to delete this product
      </ModalBody>
      <ModalFooter>
        {props.actions}
      </ModalFooter>
    </Modal>,
    document.querySelector("#modal")
  )
}

export default ModalComp