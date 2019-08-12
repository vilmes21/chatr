import React, {Component} from 'react';
import OldMsgs from './OldMsgs';
import axios from 'axios'

const _host = "ws://localhost:8081"

class ChatRoom extends Component {
    state={
        msg: "",
        oldMsgs: []
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

    _toServer = msgObj =>{
        const ws = new WebSocket(_host + "/sentence/create");
        ws.onopen = event => {
            ws.send(JSON.stringify(msgObj));
          };
        ws.onmessage = (event) => {
            console.log('just got msg back de server')
            const data = JSON.parse(event.data)
            const msg = {
                content: data.Content,
                chatSpeakerId: data.ChatSpeakerId
            }
            this.appendMsg(msg)
            console.log("APPENDING de socket server la:", event.data, "msg :", msg);
          }
    }
 
    _submit = (ev) => {
        ev.preventDefault();
        const {speakerId}=this.props;
        const {msg} = this.state;

        if (msg) {
            const newMsg ={
                content: msg,
                chatSpeakerId: speakerId
            };
            
            this._toServer(newMsg)
            this.setState({msg: ""})
        } else {
            alert("Don't send empty stuff")
        }

    }

    render() {
        const {username} = this.props;
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
