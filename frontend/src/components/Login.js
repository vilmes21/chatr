import React, {Component} from 'react';
import fakeFriends from '../fakeData/fakeFriends'
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    withRouter
  } from 'react-router-dom'
import ExistingUsers from './ExistingUsers';

class Login extends Component {
    state = {
        name: "",
        showWarning: false,
    }

    _change = (ev) => {
        const obj = {};
        const _name = ev.target.name;
        obj[_name] = ev.target.value;

        this.setState(obj);
    }

    _submit = (ev) => {
        const {login} = this.props;
        ev.preventDefault();
        const {name} = {
            ...this.state
        };

        if (fakeFriends[name]) {
            this.props.login()
        } else {
            this.setState({showWarning: true})
        }
    }

    render() {
        const {isAuth, login}=this.props;
        const {showWarning} = this.state;

        if (isAuth) {
            return <Redirect to='/' />
        }

        return (
            <div>
                <div>
                    <h2>Demo Credentials (Click)</h2>

                    <ExistingUsers users={fakeFriends} login={login}/>
                    {/* <ul>
                        <li onClick={login}>Adam</li>
                        <li onClick={login}>Bob</li>
                        <li onClick={login}>Cathy</li>
                    </ul> */}
                </div>

                <form onSubmit={this._submit}>
                    <input name="name" placeholder="Name" onChange={this._change}/> {showWarning && <div>Just use a demo name pls</div>}

                    <input type="submit" value="Log in"/>
                </form>
            </div>
        );
    }
}

export default Login;
