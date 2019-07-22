import React, {Component} from 'react';
import Login from './components/Login';
import './App.css';

class App extends Component {
    signup = obj => {
        console.log("received obj", obj)
    }

    render() {
        return (
            <div className="App">
                Sign up
                <Login signup={this.signup}/>
            </div>
        );
    }
}

export default App;
