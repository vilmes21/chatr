import React, {Component} from 'react';
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    withRouter
  } from 'react-router-dom'

class Dashboard extends Component {
     
    render() {
        const {isAuth, ...foo} = this.props;

        console.log("foo:", foo, "isAuth:", isAuth)

        if (!isAuth){
            return <Redirect to='/login' />
        }
        
        return (
            <div>
               Hello Dashboard
            </div>
        );
    }
}

export default Dashboard;
