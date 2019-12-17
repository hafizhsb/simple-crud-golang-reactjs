import React from 'react'
import { connect } from 'react-redux'
import { createProduct } from '../../actions'

import Page from 'components/Page'
import { Row, Col, Card, CardHeader, CardBody } from 'reactstrap'
import ProductForm from './ProductForm'

class ProductCreate extends React.Component {
  onSubmit = formValues => {
    this.props.createProduct(formValues)
  }

  render() {
    return (
      <Page
        title="Add Products"
        breadcrumbs={[{ name: 'products/add products', active: true }]}
        className="TablePage">
        <Row>
          <Col>
            <Card>
              <CardHeader>Create Product</CardHeader>
              <CardBody>
                <ProductForm onSubmit={this.onSubmit} />
              </CardBody>
            </Card>
          </Col>
        </Row>
      </Page>
    )
  }
}

export default connect(null, { createProduct })(ProductCreate)