import React, {Component} from 'react';
import axios from 'axios'

class ChatWindow extends Component {
    state={
        chatId: 0,

    }

    setChatId =async (userId, user2Id) => {
        const {data} = await axios.post("/chat/new", {userId, user2Id})
        console.log("after xoiso, data:", data)
    }
    
    componentDidMount(){
        const {user, userNowId} = this.props;
        this.setChatId(userNowId, user.id)
    }

    render() {
        const {user, userNowId} = this.props;

        return (
            <div key={user.id} className="chatWindow">
                <div className="chatWithTitle">{user.id}. {user.name}</div>
                <div></div>
                <div className="sendArea">
                    <form>
                        <input name="msg"/>
                        <input type="submit" value="send"/>
                    </form>
                </div>
            </div>
        )
    }
}

export default ChatWindow;
