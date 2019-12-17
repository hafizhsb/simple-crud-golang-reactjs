import React from 'react'
import _ from 'lodash'
import { connect } from 'react-redux'
import { fetchProduct, updateProduct } from '../../actions'

import Page from 'components/Page'
import { Row, Col, Card, CardHeader, CardBody } from 'reactstrap'
import ProductForm from './ProductForm'

class ProductEdit extends React.Component{
  componentDidMount() {
    this.props.fetchProduct(this.props.match.params.id)
  }

  onSubmit = (formValues) => {
    this.props.updateProduct(this.props.match.params.id, formValues)
  }

  render() {
    if (!this.props.product) {
      return <div>Loading...</div>
    }
    console.log(this.props.product)
    return (
      <Page
        title="Edit Products"
        breadcrumbs={[{ name: 'products/add products', active: true }]}
        className="TablePage">
        <Row>
          <Col>
            <Card>
              <CardHeader>Edit Product</CardHeader>
              <CardBody>
                <ProductForm
                  initialValues={_.pick(this.props.product, 'product_name', 'code', 'price', 'qty')}
                  onSubmit={this.onSubmit} />
              </CardBody>
            </Card>
          </Col>
        </Row>
      </Page>
    )
  }
}

const mapStateToProps = (state, ownProps) => {
  return {
    product: state.products[ownProps.match.params.id]
  }
}

export default connect(mapStateToProps, { fetchProduct, updateProduct })(ProductEdit)