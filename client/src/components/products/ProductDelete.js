import React from 'react'
import Modal from '../Modal'
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'
import { fetchProduct, deleteProduct } from '../../actions'

import {
  Button,
  ModalBody,
  ModalFooter,
  ModalHeader,
  Row,
} from 'reactstrap';

class ProductDelete extends React.Component {
  renderActions() {
    return(
      <React.Fragment>
        <Button onClick={() => this.props.deleteProduct(this.props.match.params.id)} color="secondary">
          Delete
        </Button>{' '}
        <Link to="/products">
          <Button color="primary">
            Cancel
          </Button>
        </Link>
      </React.Fragment>
    )
  }

  render() {
    return (
      <div>
        <Modal 
          actions={this.renderActions()}
        />
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    product: state.products[ownProps.match.params.id]
  }
}

export default connect(mapStateToProps, { fetchProduct, deleteProduct })(ProductDelete)