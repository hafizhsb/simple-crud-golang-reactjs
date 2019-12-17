import React from 'react'
import { connect } from 'react-redux'
import { fetchProducts } from '../../actions'
import { Link } from 'react-router-dom'
import Page from 'components/Page';
import { Card, CardBody, CardHeader, Col, Row, Table } from 'reactstrap';
import { ButtonGroup, Button, UncontrolledButtonDropdown, DropdownToggle, DropdownMenu, DropdownItem } from 'reactstrap'

class ProductList extends React.Component {
  componentDidMount() {
    this.props.fetchProducts()
  }

  renderButton(id) {
    return (
      <ButtonGroup>
        <UncontrolledButtonDropdown>
          <DropdownToggle caret color="primary">Select</DropdownToggle>
          <DropdownMenu>
            <DropdownItem>Detail</DropdownItem>
            <Link to={`products/edit/${id}`}><DropdownItem>Edit</DropdownItem></Link>
            <Link to={`products/delete/${id}`}><DropdownItem>Delete</DropdownItem></Link>
          </DropdownMenu>
        </UncontrolledButtonDropdown>
      </ButtonGroup>
    )
  }

  renderList() {
    let i = 1
    const nf = new Intl.NumberFormat();
    return this.props.products.map(product => {
      return (
        <tr key={product.id}>
          <th scope="row">{i++}</th>
            <td>{product.product_name}</td>
            <td>{product.code}</td>
            <td>{nf.format(product.price)}</td>
            <td>{product.qty}</td>
            <td>{this.renderButton(product.id)}</td>
        </tr>
      )
    })
  }

  render(){
    return (
      <Page
        title="Products"
        breadcrumbs={[{ name: 'products', active: true }]}
        className="TablePage"
      >
        <Row>
          <Col>
            <Link to="/products/create">
              <Button color="primary" active>
                Create
              </Button>
            </Link>
          </Col>
        </Row>
        <Row>
        <Col>
          <Card className="mb-6">
            <CardHeader>Responsive</CardHeader>
            <CardBody>
              <Table responsive>
                <thead>
                  <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Code</th>
                    <th>Price</th>
                    <th>Qty</th>
                    <th>Action</th>
                  </tr>
                </thead>
                <tbody>
                  {this.renderList()}
                </tbody>
              </Table>
            </CardBody>
          </Card>
        </Col>
      </Row>
      </Page>
    ) 
  }
}

const mapStateToProps = (state) => {
  return {
    products: Object.values(state.products)
  }
}

export default connect(mapStateToProps, { fetchProducts })(ProductList)