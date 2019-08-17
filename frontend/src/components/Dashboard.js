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
        unreadBadges: {},
        msgsOfChats: {}
        //{chatId: [{speakerUserId: 5, content: "hi", receiverUserId: 7}]}
    }

    componentDidMount = () => {
        //Demo fake ajax
        setTimeout(() => {
            this.setState({
                friends: fakeFriends,
                ws: this.connectWS()
            })
        }, 0)
    }

    rmBadge = userId => {
        const {unreadBadges} = this.state;
        this.setState({
            unreadBadges: {
                ...unreadBadges,
                [userId]: 0
            }
        })
    }

    addToMsgsOfChats = obj => {
        // console.log("func addToMsgsOfChats obj:") console.table(obj)

        const {msgsOfChats} = this.state;
        const {chatId, content, speakerUserId, receiverUserId} = obj;
        const arr = msgsOfChats[chatId];
        const newestMsg = {
            content,
            speakerUserId,
            receiverUserId
        };

        const arr2 = Array.isArray(arr)
            ? [
                ...arr,
                newestMsg
            ]
            : [newestMsg];

        this.setState({
            msgsOfChats: {
                ...msgsOfChats,
                [chatId]: arr2
            }
        }, () => {
            console.log("----------------")
            console.table(this.state.msgsOfChats)
        })
    }

    updateMsgBadge = data => {
        const {content, speakerUserId, chatId, receiverUserId} = data;
        const {userNowId} = this.props;

        if (userNowId === receiverUserId) {
            const {unreadBadges} = this.state;
            //int
            const unreadNum = unreadBadges[speakerUserId];
            const _unreadBadges = {}
            if (unreadNum) {
                _unreadBadges[speakerUserId] = unreadNum + 1;
            } else {
                _unreadBadges[speakerUserId] = 1;
            }

            this.setState({
                unreadBadges: {
                    ...unreadBadges,
                    ..._unreadBadges
                }
            })
        }
    }

    connectWS = () => {
        const ws = new WebSocket("ws://localhost:8081/sentence/create");

        ws.onopen = ev => {
            const {userNowId}=this.props;
            ws.send({userNowId}); 
            console.log("FIRST ws SENT: {userNowId}:", {userNowId})
        };

        ws.onmessage = (event) => {
            const data = JSON.parse(event.data)
            const {content, speakerUserId, chatId, receiverUserId} = data;
            console.log("windows compo got de server:", data)
            this.addToMsgsOfChats({content, speakerUserId, chatId, receiverUserId})
            this.updateMsgBadge(data)
        }
        return ws;
    }

    closeChat = friendId => {
        const {chatingWith, ws} = this.state;
        const stillInChat = chatingWith.filter(x => x.id !== friendId);
        this.setState({chatingWith: stillInChat})
    }

    addChat = userObj => {
        const {chatingWith} = this.state;

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

        const {friends, chatingWith, ws, msgsOfChats, unreadBadges} = this.state;

        return (
            <div>
                <h3>
                    Friends to Chat with:
                </h3>

                {friends
                    ? <Friends
                            addChat={this.addChat}
                            nameNow={nameNow}
                            friends={friends}
                            unreadBadges={unreadBadges}/>
                    : <span>loading...</span>}

                <ChatWindows
                    rmBadge={this.rmBadge}
                    unreadBadges={unreadBadges}
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