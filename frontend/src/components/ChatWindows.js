import React, {Component} from 'react';
import ChatWindow from './ChatWindow';



class ChatWindows extends Component {
    componentDidMount(){
        this.props.connectWS()
    }
    
    render() {
        const {users,userNowId,closeChat, ws} = this.props;

        if (users.length === 0) return null;

        return users
            .map(x => {
                return <ChatWindow closeChat={closeChat} userNowId={userNowId} key={x.id} user={x} ws={ws}/>
            })
    }
}

export default ChatWindows;
