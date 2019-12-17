import React from 'react'
import { Field, reduxForm } from 'redux-form'
import { Form, FormGroup, FormText, Label, Input, Button } from 'reactstrap'

class ProductForm extends React.Component {
  renderInput({ input, meta, label, inputType}) {
    const isValid = meta.touched && !meta.error ? true : false
    const errorMessage = meta.error && meta.touched ? meta.error : ''

    return (
      <div>
        <FormGroup>
          <Label>{label}</Label>
          <Input {...input} type={inputType} valid={isValid}/>
          <FormText color="danger">{errorMessage}</FormText>
        </FormGroup>
      </div>
    )
  }

  onSubmit = formValues => {
    this.props.onSubmit(formValues)
  }

  render() {
    return (
      <Form onSubmit={this.props.handleSubmit(this.onSubmit)}>
        <Field name="product_name" component={this.renderInput} label="Enter product name" inputType="text" />
        <Field name="code" component={this.renderInput} label="Enter product code" inputType="text" />
        <Field name="price" parse={value => Number(value)} component={this.renderInput} label="Enter product price" inputType="number" min="1" />
        <Field name="qty" parse={value => Number(value)} component={this.renderInput} label="Enter product quantity" inputType="number" min="1" />
        <Button color="primary">Submit</Button>
      </Form>
    )
  }
}

const validate = (formValues) => {
  const errors = {}

  if (!formValues.product_name) {
    errors.product_name = 'You must enter product name'
  }

  if (!formValues.price) {
    errors.price = 'You must enter product price'
  }

  // if (formValues.price && formValues.price.search(/\./) != -1) {
  //   errors.price = 'You must enter price without point or thousand separator'
  // }

  if (!formValues.qty) {
    errors.qty = 'You must enter product quantity'
  }

  // if (formValues.qty && formValues.qty.search(/\./) != -1) {
  //   errors.qty = 'You must enter quantity without point or thousand separator'
  // }

  return errors
}

export default reduxForm({
  form: 'productForm',
  validate: validate
})(ProductForm)