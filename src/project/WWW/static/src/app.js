import "./base/base.scss";

import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter, Route, Switch} from 'react-router-dom';

import PageUsersList from "./components/pages/users/list/list";
import PageUserGet from "./components/pages/users/get/get";
import PageUserCreate from "./components/pages/users/create/create";

ReactDOM.render(
    <BrowserRouter>
        <Switch>
            <Route exact path={'/users'} component={PageUsersList}/>
            <Route exact path={'/user/create'} component={PageUserCreate}/>
            <Route path={'/user/:id'} component={PageUserGet}/>
        </Switch>
    </BrowserRouter>,
    document.getElementById('react-container'));