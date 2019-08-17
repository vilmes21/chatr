import React, {Component} from 'react';
import ChatWindow from './ChatWindow';



class ChatWindows extends Component {
    render() {
        const {users,userNowId,ws,closeChat, msgsOfChats,unreadBadges, rmBadge} = this.props;

        if (users.length === 0) return null;

        return users
            .map(x => {
                const unreadNumFromUser = unreadBadges[x.id] || 0;
                
                return <ChatWindow
                rmBadge={rmBadge}
                 ws={ws} msgs={msgsOfChats} closeChat={closeChat} userNowId={userNowId} key={x.id} user={x} unreadNumFromUser={unreadNumFromUser}/>
            })
    }
}

export default ChatWindows;
