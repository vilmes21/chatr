import React, {Component} from 'react';
import axios from 'axios'
import OldMsgs from './OldMsgs';

// const scrollToBottom = htmlId => {
    // var objDiv = document.getElementById(htmlId);
    // objDiv.scrollTop = objDiv.scrollHeight;
//   }

class ChatWindow extends Component {
    state = {
        msg: "",
        chatId: 0
    }

    componentDidMount() {
        const {user, userNowId} = this.props;
        this.setChatId(userNowId, user.id)
    }

    getChatHtmlId = () =>{
        return "chatwindowJS_" + this.state.chatId;
    }

    // componentDidUpdate(){
    //     scrollToBottom(this.getChatHtmlId())
    // }

    _change = (ev) => {
        const obj = {};
        const _name = ev.target.name;
        obj[_name] = ev.target.value;

        this.setState(obj);
    }

    _submit = (ev) => {
        ev.preventDefault();
        const {userNowId,ws} = this.props;
        const {msg, chatId} = this.state;

        if (chatId === 0) {
            alert("Wait... ChatId is still unavailable")
            return;
        }

        if (msg) {
            const newMsg = {
                msg,
                userNowId,
                chatId
            };

            ws.send(JSON.stringify(newMsg));
            this.setState({msg: ""})
        } else {
            console.log("Don't send empty stuff")
        }

    }

    setChatId = async(userId, user2Id) => {
        try {
            const {data} = await axios.post("/chat/new", {userId, user2Id})
            this.setState({chatId: data.Id})
        } catch (e) {
            console.log("setChatId func e: ", e)
        }
    }

    render() {
        const {user, userNowId,rmBadge, closeChat, msgs, unreadNumFromUser} = this.props;
        const {chatId, msg} = this.state;
        const msgsOfThisChat = msgs[chatId] || [];

        return (
            <div onClick={()=>{
                if (unreadNumFromUser >0) {
                    rmBadge(user.id)
                }
            }} className="chatWindow">
                <div className="chatWithTitle">
                    <button
                        onClick={() => {
                        closeChat(user.id);
                    }}>x</button>
                    {user.id}. {user.name}
                    <br/>
                    chat-{chatId}
                    <br />
                    {
                        unreadNumFromUser>0 && 
                        <span>Unread: {unreadNumFromUser}</span>
                    }
                </div>
                <div className="oldmsgArea" id={this.getChatHtmlId()}>
                    <OldMsgs oldMsgs={msgsOfThisChat}/>
                </div>
                <div className="sendArea">
                    <form onSubmit={this._submit}>
                        <input
                            value={msg}
                            name="msg"
                            placeholder="Say stuff..."
                            onChange={this._change}/>

                        <input type="submit" value="Send"/>
                    </form>
                </div>
            </div>
        )
    }
}

export default ChatWindow;
