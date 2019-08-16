import React, {Component} from 'react';
import fakeFriends from '../fakeData/fakeFriends'
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    withRouter
  } from 'react-router-dom'
import Friends from './Friends'
import ChatWindows from './ChatWindows'

const _wshost = "ws://localhost:8081";
let ws;

const connectWS = () =>{
    ws = new WebSocket(_wshost + "/sentence/create");
    ws.onmessage = (event) => {
        const data = JSON.parse(event.data)
        const {content, speakerUserId} = data;
        console.log("windows compo got de server:", data)
        // this.appendMsg({content, speakerUserId})
    }
}

class Dashboard extends Component {
    state = {
        friends: false,
        chatingWith: []
    }

    componentDidMount=()=>{
        //Demo fake ajax
        setTimeout(()=>{
            this.setState({friends: fakeFriends})
        }, 0)
    }

    closeChat = friendId => {
        const {chatingWith}=this.state;
        const stillInChat = chatingWith.filter(x => x.id !== friendId);

        if (stillInChat.length === 0){
            console.log("Dashboard compo FE close...")
            ws.close()
        }

        this.setState({
            chatingWith: stillInChat
        })
    }

    addChat = userObj => {
        const {chatingWith}=this.state;
        const alreadyInChat = chatingWith.find(x => x.id === userObj.id);
        if (alreadyInChat) return;
        
        this.setState({
            chatingWith: [...chatingWith, userObj]
        })
    }
     
    render() {
        const {isAuth,userNowId, nameNow} = this.props;

        if (!isAuth){
            return <Redirect to='/login' />
        }

        const {friends, chatingWith}=this.state;
        
        return (
            <div>
               <h3>
               Friends to Chat with:
               </h3>

               {friends? <Friends addChat={this.addChat} nameNow={nameNow} friends={friends}/>: <span>loading...</span>}

               <ChatWindows connectWS={connectWS} ws={ws} closeChat={this.closeChat} userNowId={userNowId} users={chatingWith}/>
            </div>
        );
    }
}

export default Dashboard;
