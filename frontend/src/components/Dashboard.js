import React, {Component} from 'react';
import fakeFriends from '../fakeData/fakeFriends'
import {BrowserRouter as Router, Route, Link, Redirect, withRouter} from 'react-router-dom'
import Friends from './Friends'
import ChatWindows from './ChatWindows'

class Dashboard extends Component {
    state = {
        friends: false,
        chatingWith: [],
        ws: false,
        msgsOfChats: {}
        //{3: [{speakerUserId: 5, content: "hi"}]}
    }

    componentDidMount = () => {
        //Demo fake ajax
        setTimeout(() => {
            this.setState({friends: fakeFriends})
        }, 0)
    }

    addToMsgsOfChats = obj => {
        const {msgsOfChats} = this.state;
        const {chatId, content, speakerUserId} = obj;
        const arr = msgsOfChats[chatId];

        const arr2 = Array.isArray(arr)
            ? [
                ...arr, {
                    content,
                    speakerUserId
                }
            ]
            : [
                {
                    content,
                    speakerUserId
                }
            ];

        this.setState({
            msgsOfChats: {
                ...msgsOfChats,
                [chatId]: arr2
            }
        }, ()=>{
            console.log("----------------")
            console.table(this.state.msgsOfChats)
        })
    }

    connectWS = () => {
        const ws = new WebSocket("ws://localhost:8081/sentence/create");

        ws.onmessage = (event) => {
            const data = JSON.parse(event.data)
            const {content, speakerUserId, chatId} = data;
            console.log("windows compo got de server:", data)
            // this.appendMsg({content, speakerUserId})
            this.addToMsgsOfChats({content, speakerUserId, chatId})
        }
        return ws;
    }

    closeChat = friendId => {
        const {chatingWith, ws} = this.state;
        const stillInChat = chatingWith.filter(x => x.id !== friendId);

        if (stillInChat.length === 0) {
            console.log("Dashboard compo FE close...")
            ws.close()
        }

        this.setState({chatingWith: stillInChat})
    }

    addChat = userObj => {
        const {chatingWith} = this.state;

        //only connect websocket when 1st chat window
        if (chatingWith.length === 0) {
            this.setState({
                ws: this.connectWS()
            })
        }

        const alreadyInChat = chatingWith.find(x => x.id === userObj.id);
        if (alreadyInChat) 
            return;
        
        this.setState({
            chatingWith: [
                ...chatingWith,
                userObj
            ]
        })
    }

    render() {
        const {isAuth, userNowId, nameNow} = this.props;

        if (!isAuth) {
            return <Redirect to='/login'/>
        }

        const {friends, chatingWith, ws, msgsOfChats} = this.state;

        return (
            <div>
                <h3>
                    Friends to Chat with:
                </h3>

                {friends
                    ? <Friends addChat={this.addChat} nameNow={nameNow} friends={friends}/>
                    : <span>loading...</span>}

                <ChatWindows
                    msgsOfChats={msgsOfChats}
                    ws={ws}
                    closeChat={this.closeChat}
                    userNowId={userNowId}
                    users={chatingWith}/>
            </div>
        );
    }
}

export default Dashboard;
