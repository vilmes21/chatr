import React, {Component} from 'react';
import axios from 'axios'
import OldMsgs from './OldMsgs';



class ChatWindow extends Component {
    state = {
        msg: "",
        oldMsgs: [],
        chatId: 0
    }

    componentDidMount() {
        const {user, userNowId} = this.props;
        this.setChatId(userNowId, user.id)


    }

    _change = (ev) => {
        const obj = {};
        const _name = ev.target.name;
        obj[_name] = ev.target.value;

        this.setState(obj);
    }

    appendMsg = msgObj => {
        this.setState({
            oldMsgs: [
                ...this.state.oldMsgs,
                msgObj
            ]
        })
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
        const {user, userNowId, closeChat,ws} = this.props;
        const {chatId, msg, oldMsgs} = this.state;

        return (
            <div key={user.id} className="chatWindow">
                <div className="chatWithTitle">
                    <button
                        onClick={() => {
                        closeChat(user.id);
                        // console.log("FE close ws...");
                        // ws.close();
                    }}>x</button>
                    {user.id}. {user.name}
                    <br/>
                    chat-{chatId}
                </div>
                <div className="oldmsgArea">
                    <OldMsgs oldMsgs={oldMsgs}/>
                </div>
                <div className="sendArea">
                    <form onSubmit={this._submit}>
                        <input
                            autoFocus
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
