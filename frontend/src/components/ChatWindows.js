import React, {Component} from 'react';
import ChatWindow from './ChatWindow';

class ChatWindows extends Component {
    render() {
        const {users,userNowId} = this.props;

        if (users.length === 0) return null;

        return users
            .map(x => {
                return <ChatWindow userNowId={userNowId} key={x.id} user={x}/>
            })
    }
}

export default ChatWindows;
