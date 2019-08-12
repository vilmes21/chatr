import React, {Component} from 'react';
import Login from './components/Login';
import ChatRoom from './components/ChatRoom';
import './App.css';

class App extends Component {
    state = {
        username: "",
        speakerId: 0
    }

    signup = obj => {
        const {name} = obj;

        this.setState({username: name, speakerId: 3})
    }

    render() {
        const {username, speakerId} = this.state;

        if (username) {
            return <ChatRoom username={username} speakerId={speakerId}/>
        }

        return (
            <div className="App">
                Sign up {username}
                <Login signup={this.signup}/>
            </div>
        );
    }
}

export default App;
