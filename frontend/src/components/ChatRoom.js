import React, {Component} from 'react';
import OldMsgs from './OldMsgs';
import axios from 'axios'
import {
    BrowserRouter as Router,
    Route,
    Link,
    Redirect,
    withRouter
  } from 'react-router-dom'

const _host = "ws://localhost:8081";
let ws;

const _fakeChatId = Math.ceil(Math.random() * 10) % 2==0? 1: 2
console.log("_fakeChatId: ", _fakeChatId)

class ChatRoom extends Component {
    state={
        msg: "",
        oldMsgs: []
    }

    componentDidMount(){
        ws = new WebSocket(_host + "/sentence/create");
   
        ws.onmessage = (event) => {
            const data = JSON.parse(event.data)
            console.log("de server data:",data)
            const msg = {
                content: data.content,
                chatSpeakerId: data.chatSpeakerId
            }
            this.appendMsg(msg)
        }
    }

    _change = (ev) => {
        const obj = {};
        const _name = ev.target.name;
        obj[_name] = ev.target.value;

        this.setState(obj);
    }

    appendMsg = msgObj =>{
        this.setState({
            oldMsgs: [...this.state.oldMsgs, msgObj]
        })
    }

    _submit = (ev) => {
        ev.preventDefault();
        const {speakerId}=this.props;
        const {msg} = this.state;
        
        if (msg) {
            const newMsg ={
                content: msg,
                chatSpeakerId: speakerId,
                chatId: _fakeChatId
            };
            
            ws.send(JSON.stringify(newMsg));
            this.setState({msg: ""})
        } else {
            alert("Don't send empty stuff")
        }

    }

    render() {
        const {username, isAuth} = this.props;
        if (!isAuth){
            return <Redirect to='/login' />
        }
        
        const {msg, oldMsgs} = this.state;
        return (
            <div>
                <h2>{username}, say something here!</h2>
                <div className="chatMsgArea">
                    <OldMsgs oldMsgs={oldMsgs} />
                </div>
                <form onSubmit={this._submit}>
                    <input autoFocus value={msg} name="msg" placeholder="Say stuff..." onChange={this._change}/> 

                    <input type="submit" value="Send"/>
                </form>
            </div>
        );
    }
}

export default ChatRoom;
