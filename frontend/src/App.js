import React, {Component} from 'react';
import Login from './components/Login';
import Signup from './components/Signup';
import Dashboard from './components/Dashboard';
import ChatRoom from './components/ChatRoom';
import './App.css';
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    withRouter
  } from 'react-router-dom'

class App extends Component {
    state = {
        username: "",
        speakerId: 0,
        isAuth: false
    }

    signup = obj => {
        const {name} = obj;

        this.setState({username: name, speakerId: 3})
    }

    login = () =>{
        this.setState({isAuth: true})
    }

    logout = () =>{
        this.setState({isAuth: false})
    }

    render() {
        const {username, speakerId, isAuth} = this.state;

        const toAuthLinks = isAuth? null : 
        <span>
 <li>
                      <Link to="/login">Login</Link>
                    </li>
                    <li>
                      <Link to="/signup">Sign up</Link>
                    </li>
        </span>

        const toLogOutLink = isAuth? <li onClick={this.logout}>Log out</li>:null;

        return (
            <Router>
              <div>
                <nav>
                  <ul>
                    <li>
                      <Link to="/">{isAuth? "":"Private"} Dashboard</Link>
                    </li>

                    {toAuthLinks}
                    {toLogOutLink}                   
                  </ul>
                </nav>
        
                <Route path="/" exact component={()=>{return <Dashboard isAuth={isAuth}/>}} />

                <Route path="/login" exact component={()=>{return <Login login={this.login} isAuth={isAuth}/>}} />                
                
                <Route path="/signup" component={()=>{return <Signup signup={this.signup}/>}} />
              </div>
            </Router>
          );
        
        
    }
}

export default App;
