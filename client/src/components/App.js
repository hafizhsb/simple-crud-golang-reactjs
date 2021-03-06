import React from 'react';
import componentQueries from 'react-component-queries';
import { Redirect, Router, Route, Switch } from 'react-router-dom';
import { EmptyLayout, LayoutRoute, MainLayout } from 'components/Layout';
import PageSpinner from 'components/PageSpinner';
import history from '../history'
import '../styles/reduction.scss';

const AlertPage = React.lazy(() => import('../pages/AlertPage'));
const AuthModalPage = React.lazy(() => import('../pages/AuthModalPage'));
const BadgePage = React.lazy(() => import('../pages/BadgePage'));
const ButtonGroupPage = React.lazy(() => import('../pages/ButtonGroupPage'));
const ButtonPage = React.lazy(() => import('../pages/ButtonPage'));
const CardPage = React.lazy(() => import('../pages/CardPage'));
const ChartPage = React.lazy(() => import('../pages/ChartPage'));
const DashboardPage = React.lazy(() => import('../pages/DashboardPage'));
const DropdownPage = React.lazy(() => import('../pages/DropdownPage'));
const FormPage = React.lazy(() => import('../pages/FormPage'));
const InputGroupPage = React.lazy(() => import('../pages/InputGroupPage'));
const ModalPage = React.lazy(() => import('../pages/ModalPage'));
const ProgressPage = React.lazy(() => import('../pages/ProgressPage'));
const TablePage = React.lazy(() => import('../pages/TablePage'));
const TypographyPage = React.lazy(() => import('../pages/TypographyPage'));
const WidgetPage = React.lazy(() => import('../pages/WidgetPage'));

// New component
const ProductList = React.lazy(() => import('./products/ProductList'))
const ProductCreate = React.lazy(() => import('./products/ProductCreate'))
const ProductEdit = React.lazy(() => import('./products/ProductEdit'))
const ProductDelete = React.lazy(() => import('./products/ProductDelete'))

const getBasename = () => {
  return `/${process.env.PUBLIC_URL.split('/').pop()}`;
};

class App extends React.Component {
  render() {
    return (
      <Router history={history}>
        <MainLayout breakpoint={this.props.breakpoint}>
        <React.Suspense fallback={<PageSpinner />}>
                <Route exact path="/" component={DashboardPage} />
                <Route exact path="/products" component={ProductList} />
                <Route exact path="/products/create" component={ProductCreate} />
                <Route exact path="/products/edit/:id" component={ProductEdit} />
                <Route exact path="/products/delete/:id" component={ProductDelete} />
                <Route exact path="/login-modal" component={AuthModalPage} />
                <Route exact path="/buttons" component={ButtonPage} />
                <Route exact path="/cards" component={CardPage} />
                <Route exact path="/widgets" component={WidgetPage} />
                <Route exact path="/typography" component={TypographyPage} />
                <Route exact path="/alerts" component={AlertPage} />
                <Route exact path="/tables" component={TablePage} />
                <Route exact path="/badges" component={BadgePage} />
                <Route
                  exact
                  path="/button-groups"
                  component={ButtonGroupPage}
                />
                <Route exact path="/dropdowns" component={DropdownPage} />
                <Route exact path="/progress" component={ProgressPage} />
                <Route exact path="/modals" component={ModalPage} />
                <Route exact path="/forms" component={FormPage} />
                <Route exact path="/input-groups" component={InputGroupPage} />
                <Route exact path="/charts" component={ChartPage} />
              </React.Suspense>
        </MainLayout>
      </Router>
    );
  }
}

const query = ({ width }) => {
  if (width < 575) {
    return { breakpoint: 'xs' };
  }

  if (576 < width && width < 767) {
    return { breakpoint: 'sm' };
  }

  if (768 < width && width < 991) {
    return { breakpoint: 'md' };
  }

  if (992 < width && width < 1199) {
    return { breakpoint: 'lg' };
  }

  if (width > 1200) {
    return { breakpoint: 'xl' };
  }

  return { breakpoint: 'xs' };
};

export default componentQueries(query)(App);
